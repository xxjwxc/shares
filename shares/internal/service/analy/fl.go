package analy

import (
	"bytes"
	"fmt"
	"os/exec"
	"path"
	"shares/internal/core"
	"shares/internal/model"
	"shares/internal/service/event"
	"shares/internal/service/weixin"
	proto "shares/rpc/shares"
	"strings"
	"time"

	"github.com/axgle/mahonia"
	"github.com/xxjwxc/public/myclipboard"
	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/tools"
	wx "github.com/xxjwxc/public/weixin"
	"gorm.io/datatypes"
)

// 本地放量提醒
func watchLocalFl() {
	if !event.IsWorkDay() {
		return
	}

	for {
		getMinuteFangLiangLocal()
		countHYZLJLR()       // 计算行业
		if !continueTime() { // 3点之后是空闲时间
			break
		} else {
			time.Sleep(1 * time.Minute)
		}
	}
}

// 监听实时放量情况
func getMinuteFangLiangLocal() {
	now := time.Now()
	day0 := tools.GetUtcDay0(now)
	offset := now.Unix() - day0
	if offset > (11*60*60+30*60) && offset < 13*60*60 {
		return
	}

	orm := core.Dao.GetDBr()
	list, _ := model.AnalyFlTblMgr(orm.DB).Gets()
	for _, v := range list {
		out, err := event.GetMinute(v.Code)
		if err != nil {
			mylog.Error(err)
			continue
		}
		num := len(out.List)
		if num >= 6 { // 开始计算
			left := out.List[num-4].Vol + out.List[num-5].Vol + out.List[num-6].Vol
			right := out.List[num-1].Vol + out.List[num-2].Vol + out.List[num-3].Vol
			if left*10 < right && left > 0 && right > 5000 { // 提醒
				sharesInfo, err := event.GetShare(v.Code, true)
				if err != nil {
					mylog.Error(err)
					continue
				}
				sendflMsgLocal(sharesInfo, left, right)
			}
		}
	}

	// 低位放量
	list1, _ := model.AnalyDwflTblMgr(orm.DB).Gets()
	for _, v := range list1 {
		out, err := event.GetMinute(v.Code)
		if err != nil {
			mylog.Error(err)
			continue
		}
		num := len(out.List)
		if num >= 6 { // 开始计算
			left := out.List[num-4].Vol + out.List[num-5].Vol + out.List[num-6].Vol
			right := out.List[num-1].Vol + out.List[num-2].Vol + out.List[num-3].Vol
			if left*10 < right && left > 0 && right > 5000 { // 提醒
				sharesInfo, err := event.GetShare(v.Code, true)
				if err != nil {
					mylog.Error(err)
					continue
				}
				sendflMsgLocal(sharesInfo, left, right)
			}
		}
	}
}

// 线上放量提醒
func watchFl() {
	if !event.IsWorkDay() {
		return
	}

	go func() { // 实时监听
		for {
			getMinuteFangLiang()
			countHYZLJLR()       // 计算行业
			if !continueTime() { // 3点之后是空闲时间
				break
			} else {
				time.Sleep(1 * time.Minute)
			}
		}
	}()

	for {
		watchZTB()  // 监听涨停板板块信息
		fangLiang() // 放量

		if !continueTime() { // 3点之后是空闲时间
			break
		} else {
			time.Sleep(10 * time.Minute)
		}
	}

}

// 放量(放量超过前一天3倍，且成交量>10%)
func fangLiang() {
	var codes []string
	orm := core.Dao.GetDBw()
	day0 := tools.GetUtcDay0(time.Now())
	day1Flag := true // 是否补当前的行情
	offset := time.Now().Unix() - day0
	if offset >= (14 * 60 * 60) { // 2点之后后一天
		day1Flag = true
	}

	if offset >= (15 * 60 * 60) {
		day0 += 24 * 60 * 60
		day1Flag = false
	}

	// model.AnalyFlTblMgr(orm.Where("1 = 1")).Update("tag", 2) // 待删除
	malist, err := model.AnalyFlTblMgr(orm.DB).Gets() // 保留
	if err != nil {
		return
	}
	for _, v := range malist {
		sharesInfo := event.SearchDetail(v.Code)
		if v.Price < sharesInfo.Price && ((sharesInfo.Price - v.Price) > v.Price*0.1) { // 跌破10% 就剔除
			model.AnalyFlTblMgr(orm.DB).Delete(&v)
		} else {
			v.Percent = sharesInfo.Percent
			v.NewPrice = sharesInfo.Price
			v.TurnoverRate = sharesInfo.TurnoverRate
			model.AnalyFlTblMgr(orm.DB).Save(&v) // 保留
		}
	}

	model.SharesInfoTblMgr(orm.DB).Select("code").Find(&codes)
	for _, code := range codes {
		list, err := model.SharesDailyTblMgr(orm.Where("day0 < ?", day0).Order("day0 desc").Limit(2)).GetFromCode(code)
		if err != nil {
			continue
		}
		if len(list) < 2 { // 不需要
			continue
		}
		if day1Flag {
			sharesInfo := event.SearchDetail(code)
			if sharesInfo != nil {
				list = append([]*model.SharesDailyTbl{{
					Code:         code,
					Price:        sharesInfo.Price,
					Percent:      sharesInfo.Percent,
					Volume:       sharesInfo.Volume,
					Turnover:     sharesInfo.Turnover,
					TurnoverRate: sharesInfo.TurnoverRate,
					Max:          sharesInfo.Max,
					Min:          sharesInfo.Min,
					Pe:           sharesInfo.Pe,
					Pb:           sharesInfo.Pb,
					Total:        sharesInfo.Total,
					Cap:          sharesInfo.Cap,
					Zljlr:        sharesInfo.Zljlr,
					Macd:         sharesInfo.Macd,
					Dif:          sharesInfo.Dif,
					Dea:          sharesInfo.Dea,
					CreatedAt:    time.Now(),
				}}, list...)
			}
		}

		// 放量超过前一天3倍，且成交量>10%
		if list[0].TurnoverRate >= 10 && list[0].TurnoverRate >= (list[1].TurnoverRate*3) {
			sharesInfo, err := event.GetShare(code, true)
			if err != nil {
				continue
			}
			// if len(sharesInfo.Code) == 0 { // 删除
			// 	model.SharesDailyTblMgr(orm.Where("code = ?", code)).Delete(&model.SharesDailyTbl{})
			// 	continue
			// }
			tmp, _ := model.AnalyFlTblMgr(orm.DB).GetFromCode(code)
			if tmp.ID == 0 {
				tmp.Code = code
				tmp.Day = datatypes.Date(time.Now())
				tmp.DayStr = list[0].Day0Str
				tmp.CreatedAt = time.Now()
				tmp.Price = sharesInfo.Price
			}
			tmp.Name = sharesInfo.Name
			tmp.Percent = sharesInfo.Percent
			tmp.TurnoverRate = list[0].TurnoverRate
			tmp.Score = 100 // 100分
			model.AnalyFlTblMgr(orm.DB).Save(&tmp)
		}
	}
}

// 监听实时放量情况
func getMinuteFangLiang() {
	now := time.Now()
	day0 := tools.GetUtcDay0(now)
	offset := now.Unix() - day0
	if offset > (11*60*60+30*60) && offset < 13*60*60 {
		return
	}

	orm := core.Dao.GetDBr()
	list, _ := model.AnalyFlTblMgr(orm.DB).Gets()
	for _, v := range list {
		out, err := event.GetMinute(v.Code)
		if err != nil {
			mylog.Error(err)
			continue
		}
		num := len(out.List)
		if num >= 6 { // 开始计算
			left := out.List[num-4].Vol + out.List[num-5].Vol + out.List[num-6].Vol
			right := out.List[num-1].Vol + out.List[num-2].Vol + out.List[num-3].Vol
			if left*10 < right && left > 0 && right > 10000 { // 提醒
				sharesInfo, err := event.GetShare(v.Code, true)
				if err != nil {
					mylog.Error(err)
					continue
				}
				sendflMsg(sharesInfo, left, right)
			}
		}
	}

	// 低位放量
	list1, _ := model.AnalyDwflTblMgr(orm.DB).Gets()
	for _, v := range list1 {
		out, err := event.GetMinute(v.Code)
		if err != nil {
			mylog.Error(err)
			continue
		}
		num := len(out.List)
		if num >= 6 { // 开始计算
			left := out.List[num-4].Vol + out.List[num-5].Vol + out.List[num-6].Vol
			right := out.List[num-1].Vol + out.List[num-2].Vol + out.List[num-3].Vol
			if left*10 < right && left > 0 && right > 10000 { // 提醒
				sharesInfo, err := event.GetShare(v.Code, true)
				if err != nil {
					mylog.Error(err)
					continue
				}
				sendflMsg(sharesInfo, left, right)
			}
		}
	}
}

func sendflMsg(sharesInfo *proto.SharesInfo, left, right int64) {
	var wxMsg []wx.TempWebMsg
	data := make(map[string]map[string]string)
	// mp := make(map[string]string)
	// mp["value"] = "你的分析任务已经运行完成"
	// //mp["color"] = v.color
	// data["first"] = mp

	mp := make(map[string]string)
	mp["value"] = "分时放量提醒" //
	// mp["color"] = v.color
	data["first"] = mp

	mp = make(map[string]string)
	mp["value"] = "分时放量提醒"
	// mp["color"] = v.color
	data["keyword1"] = mp

	mp = make(map[string]string)
	mp["value"] = fmt.Sprintf("%v(%v)", sharesInfo.Name, sharesInfo.Code)
	// mp["color"] = v.color
	data["keyword2"] = mp

	mp = make(map[string]string)
	mp["value"] = fmt.Sprintf("%v(%v%%)", sharesInfo.Price, sharesInfo.Percent)
	data["keyword3"] = mp

	mp = make(map[string]string)
	mp["value"] = tools.GetTimeStr(time.Now())
	data["keyword5"] = mp

	mp = make(map[string]string)
	mp["value"] = fmt.Sprintf("成交量由%v放大到%v", getSamplePrice(float64(left)), getSamplePrice(float64(right)))
	data["remark"] = mp

	wxMsg = append(wxMsg, wx.TempWebMsg{
		Touser:     "oxFCP6hcaZTReezFSs80ZY6qNAv8",
		TemplateID: "reXjLLFWHN61C3wpmTm-brJoLuDwiSpnca66XmRVJVw",
		Page:       fmt.Sprintf("https://hospital.xxjwxc.cn/webshares/#/pages/add/add?scode=%v&tag=%v", sharesInfo.Code, "min"),
		Data:       data,
	})

	weixin.SendMsg(wxMsg)
}

func continueTime() bool {
	now := time.Now()
	day0 := tools.GetUtcDay0(now)
	offset := now.Unix() - day0

	return offset <= (15 * 60 * 60) // 3点之后是空闲时间
}

func sendflMsgLocal(sharesInfo *proto.SharesInfo, left, right int64) {
	var msgs []string
	msgs = append(msgs, fmt.Sprintf("股票名字:%v(%v)", sharesInfo.Name, sharesInfo.Code))
	msgs = append(msgs, fmt.Sprintf("当前价格:%v(%v%%)", sharesInfo.Price, sharesInfo.Percent))
	msgs = append(msgs, fmt.Sprintf("成交量由%v放大到%v", getSamplePrice(float64(left)), getSamplePrice(float64(right))))
	msgs = append(msgs, fmt.Sprintf("https://hospital.xxjwxc.cn/webshares/#/pages/add/add?scode=%v&tag=%v", sharesInfo.Code, "min"))
	myclipboard.Set(strings.Join(msgs, "\r\n"))
	datapath := "cd " + path.Join(tools.GetCurrentDirectory(), "/tools") + "&MyCheck.exe"
	cmd := exec.Command("cmd", "/C", datapath)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		enc := mahonia.NewDecoder("gbk")
		str := enc.ConvertString(stderr.String())
		fmt.Printf("%s\n", str)
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		fmt.Println("Result: " + out.String())
		mylog.Error(err)
	}
}
