package analy

import (
	"fmt"
	"shares/internal/config"
	"shares/internal/core"
	"shares/internal/model"
	"shares/internal/service/event"
	"shares/internal/service/weixin"
	proto "shares/rpc/shares"
	"strings"
	"time"

	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/timerDeal"
	"github.com/xxjwxc/public/tools"
	wx "github.com/xxjwxc/public/weixin"
	"gorm.io/datatypes"
)

func init() {
	if config.GetIsTools() > 0 {
		return
	}

	timerDeal.OnPeDay(8, 0, 0, ticketHy) // 初始化行业
	timerDeal.OnPeDay(9, 10, 0, maDaily) // 每天9点10分执行()
	timerDeal.OnPeDay(10, 0, 0, watchFl)
	timerDeal.OnPeDay(10, 10, 0, watchMyself) // 每日
	timerDeal.OnPeDay(14, 30, 0, ticket14)
	timerDeal.OnPeDay(16, 0, 0, ticketTJ)
}

func ticketHy() {
	if !event.IsWorkDay() {
		return
	}
	initHY() // 初始化行业
}

func ticket14() { // 时间间隔
	if !event.IsWorkDay() {
		return
	}

	tj()             // 统计一遍
	timeWatch()      // 执行watch事件
	dealWatchGroup() // 执行入池事件
	updatePeg()      // peg 计算
}

func tj() { // 统计
	hengpan() // 很盘
	// maLhb()       // 龙虎榜
	fangLiang()   // 放量
	dwfl()        // 底部放量
	chaoDie()     // 超跌
	initmaDaily() // 趋势
	dbSzx()       // 底部缩量十字星
}

func ticketTJ() { // 统计
	tj()   // 统计一遍
	tkgk() // 分析跳空
}

func dealWatchGroup() {
	if !event.IsWorkDay() {
		return
	}

	var codes []string
	orm := core.Dao.GetDBr()
	list, _ := model.GroupTblMgr(orm.DB).Gets()
	for _, v := range list { // 初始化
		codes = append(codes, v.Code)
	}

	mp := getMp(codes)
	// 开始提醒
	startWatch(mp)
}

// 开始入池提醒
func startWatch(mp map[string][]string) {
	orm := core.Dao.GetDBw()
	for code, v := range mp {
		if len(v) == 0 {
			continue
		}
		sharesInfo, _ := event.GetShare(code, true)
		gs, _ := model.GroupTblMgr(orm.DB).GetFromCode(code)
		info, _ := model.GroupWatchTblMgr(orm.DB).GetFromCode(code)
		if len(gs) > 0 { // 开始提醒
			if info.ID > 0 {
				s := strings.Split(info.Desc, ";")
				if len(v) <= len(s) { // 无更新
					continue
				}
				info.Desc = strings.Join(v, ";")
			} else {
				info = model.GroupWatchTbl{
					Code:      code,
					Name:      sharesInfo.Name,
					Price:     sharesInfo.Price,
					NewPrice:  sharesInfo.Price,
					Desc:      strings.Join(v, ";"),
					CreatedAt: time.Now(),
					UserName:  gs[0].UserName,
				}
			}
			model.GroupWatchTblMgr(orm.DB).Save(&info)
			for _, v1 := range gs { // 开始发送
				sendWatchMsg(sharesInfo, v1.GroupName, v1.UserName, v, "入池提醒")
				// model.GroupTblMgr(orm.DB).Delete(&v1) // 提醒了之后删除
			}
		}
	}
}

// 时间片轮转
func timeWatch() {
	if !event.IsWorkDay() {
		return
	}

	orm := core.Dao.GetDBw()
	list, _ := model.GroupWatchTblMgr(orm.DB).Gets()
	for _, v := range list { // 准备删除
		var msgs []string
		sharesInfo, _ := event.GetShare(v.Code, true)
		if sharesInfo.Price < v.Price && ((v.Price - sharesInfo.Price) > v.Price*0.1) { // 跌破10% 就剔除
			model.GroupWatchTblMgr(orm.DB).Delete(&v)
			msgs = append(msgs, fmt.Sprintf("当前价格(%v),跌破入池价格(%v) 的 %%10", sharesInfo.Price, v.Price))

		} else {
			// 上涨10%以上以10日线为准
			if sharesInfo.Price > v.Price && ((sharesInfo.Price - v.Price) > v.Price*0.1) { // 上涨10% 以上了，用十日线为准
				list, _ := model.SharesDailyTblMgr(orm.Where("day0 < ?", tools.GetUtcDay0(time.Now())).Order("day0 desc").Limit(1)).GetFromCode(v.Code)
				if len(list) > 0 {
					if list[0].Ma10 > 0 && sharesInfo.Price < list[0].Ma10 { // 破10日线就剔除
						model.GroupWatchTblMgr(orm.DB).Delete(&v)
						msgs = append(msgs, fmt.Sprintf("入池价格(%v),当前价格(%v),跌破10日线(%v)", v.Price, sharesInfo.Price, list[0].Ma10))
					}
				}
			}
		}

		if len(msgs) > 0 {
			gs, _ := model.GroupTblMgr(orm.DB).GetFromCode(v.Code)
			for _, v1 := range gs { // 开始发送
				sendWatchMsg(sharesInfo, v1.GroupName, v1.UserName, msgs, "剔除提醒")
				model.GroupTblMgr(orm.DB).Delete(&v1) // 删除
			}
		} else {
			v.NewPrice = sharesInfo.Price
			model.GroupWatchTblMgr(orm.DB).Save(&v) // 保留
		}
	}

}

func getMp(codes []string) map[string][]string {
	mp := make(map[string][]string)
	orm := core.Dao.GetDBr()
	for _, code := range codes { // 初始化
		mp[code] = []string{}
	}

	// 龙虎榜
	lhbs, _ := model.AnalyLhbTblMgr(orm.DB).GetBatchFromCode(codes)
	for _, v := range lhbs {
		mp[v.Code] = append(mp[v.Code], fmt.Sprintf("龙虎榜(%v) 总金额:%v 机构净流入:%v", v.DayStr, getSamplePrice(v.Jlr*10000), getSamplePrice(v.JgJlr)))
	}

	// 放量
	fls, _ := model.AnalyFlTblMgr(orm.DB).GetBatchFromCode(codes)
	for _, v := range fls {
		mp[v.Code] = append(mp[v.Code], fmt.Sprintf("放巨量(%v) :%v%%", tools.GetDayStr(time.Time(v.Day)), v.TurnoverRate))
	}

	// 超跌
	cds, _ := model.AnalyCdTblMgr(orm.DB).GetBatchFromCode(codes)
	for _, v := range cds {
		mp[v.Code] = append(mp[v.Code], fmt.Sprintf("超跌反弹(%v)", tools.GetDayStr(tools.GetDay0(v.Day0))))
	}

	// 趋势
	ups, _ := model.AnalyUpTblMgr(orm.DB).GetBatchFromCode(codes)
	for _, v := range ups {
		mp[v.Code] = append(mp[v.Code], fmt.Sprintf("趋势向上(%v)", tools.GetDayStr(tools.GetDay0(v.Day0))))
	}

	// 底部缩量十字星
	szxs, _ := model.AnalyDbszxTblMgr(orm.DB).GetBatchFromCode(codes)
	for _, v := range szxs {
		mp[v.Code] = append(mp[v.Code], fmt.Sprintf("底部缩量十字星(%v)", tools.GetDayStr(time.Time(v.Day))))
	}

	// 低位放量
	dwfls, _ := model.AnalyDwflTblMgr(orm.DB).GetBatchFromCode(codes)
	for _, v := range dwfls {
		mp[v.Code] = append(mp[v.Code], fmt.Sprintf("低位放量(%v)", tools.GetDayStr(time.Time(v.Day))))
	}

	// 跳空回补
	tkhb, _ := model.AnalyTkTblMgr(orm.DB).GetBatchFromCode(codes)
	for _, v := range tkhb {
		mp[v.Code] = append(mp[v.Code], fmt.Sprintf("跳空回补(%v):(%v)", v.Min, v.Day0Str))
	}

	return mp
}

func sendWatchMsg(sharesInfo *proto.SharesInfo, groupName, username string, msgs []string, key string) error {
	if len(msgs) == 0 {
		return nil
	}

	wxMsg := make([]wx.TempWebMsg, 0, len(msgs))
	list := make([]model.MsgTbl, 0, len(msgs))
	orm := core.Dao.GetDBr()
	users, _ := model.WxUserinfoMgr(orm.Where("`group` like ?", fmt.Sprintf("%%%v%%", groupName))).Gets()
	//users, _ := model.WxUserinfoMgr(orm.Where("openid = 'oxFCP6hcaZTReezFSs80ZY6qNAv8'")).Gets()
	for _, v := range users {
		list = append(list, model.MsgTbl{
			OpenID:    v.Openid,
			Code:      sharesInfo.Code,
			Key:       "组织提醒",
			Desc:      strings.Join(msgs, "\n"),
			Price:     sharesInfo.Price,
			Percent:   sharesInfo.Percent,
			Day:       datatypes.Date(time.Now()),
			CreatedAt: time.Now(),
		})

		data := make(map[string]map[string]string)
		mp := make(map[string]string)
		mp["value"] = fmt.Sprintf("股票名称:%v(%v)", sharesInfo.Name, sharesInfo.Code)
		data["first"] = mp

		mp = make(map[string]string)
		mp["value"] = fmt.Sprintf("%v(%v)", groupName, username)
		data["keyword1"] = mp

		mp = make(map[string]string)
		mp["value"] = key
		data["keyword2"] = mp

		mp = make(map[string]string)
		mp["value"] = tools.GetTimeStr(time.Now())
		data["keyword3"] = mp

		mp = make(map[string]string)
		mp["value"] = strings.Join(msgs, "\n")
		data["remark"] = mp

		wxMsg = append(wxMsg, wx.TempWebMsg{
			Touser:     v.Openid,
			TemplateID: "z8VrA-QQRjXLIRAkz54lSPvHuaWU2mIzyeWuSfQBRms",
			Page:       fmt.Sprintf("https://hospital.xxjwxc.cn/webshares/#/pages/add/add?scode=%v&tag=%v", sharesInfo.Code, "daily"),
			Data:       data,
		})
	}

	weixin.SendMsg(wxMsg)
	err := model.MsgTblMgr(orm.DB).Save(&list).Error
	if err != nil {
		mylog.Error(err)
	}

	return err
}
