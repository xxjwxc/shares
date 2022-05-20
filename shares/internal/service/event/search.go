package event

import (
	"fmt"
	"regexp"
	"shares/internal/config"
	"shares/internal/core"
	"shares/internal/model"
	"strconv"
	"strings"
	"time"

	proto "shares/rpc/shares"

	"github.com/xxjwxc/public/mycache"
	"github.com/xxjwxc/public/tools"
)

// TrySearch 试着搜索(sh|sz|hk)
func TrySearch(code string) *proto.SharesInfo {
	aORb := regexp.MustCompile(strings.Join(config.GetExt(), "|"))
	matches := aORb.FindAllStringIndex(code, -1)
	if len(matches) > 0 { // 有包含
		out := Searchs([]string{code})
		if len(out) > 0 {
			return out[0]
		}
		return nil
	}

	// 添加头
	for _, v := range config.GetExt() {
		out := Searchs([]string{v + code})
		if len(out) > 0 {
			return out[0]
		}
	}

	// 到此没有搜索到,记录信息
	SaveLog(fmt.Sprintf("TrySearch error:%v", code))
	return nil
}

// Search 确定的搜索
func Searchs(codes []string) (resp []*proto.SharesInfo) {
	if len(codes) == 0 {
		return nil
	}
	parm := strings.Join(codes, ",s_")
	url := "http://qt.gtimg.cn/q=s_" + parm
	out := SendGet(url, "")
	out = tools.ConvertToString(out, "gbk", "utf8")
	if len(out) == 0 {
		return nil
	}
	out = strings.Replace(out, "\n", "", -1)

	// todo 分析结果
	list := strings.Split(out, ";")
	for _, v := range list {
		if len(v) < 4 {
			continue
		}
		list1 := strings.Split(v, "=")
		if len(list1) == 2 {
			if strings.HasPrefix(list1[0], "v_s_") {
				tmp := &proto.SharesInfo{
					Code: list1[0][4:],
				}
				list1[1] = strings.Trim(list1[1], "\"")
				list2 := strings.Split(list1[1], "~")
				if len(list2) > 9 {
					tmp.Ext = _extMp[list2[0]]
					tmp.Name = strings.Replace(list2[1], " ", "", -1)
					tmp.SimpleCode = list2[2]
					tmp.Price = stringToFloat(list2[3])
					tmp.Percent = stringToFloat(list2[5])
				}
				resp = append(resp, tmp)
			}
		}
	}
	return resp
}

// SearchDetail 确定的搜索(全量搜索)
func SearchDetail(code string) *proto.SharesInfoDetails {
	resp := SearchDetails([]string{code})
	if len(resp) > 0 {
		return resp[0]
	}

	return nil
}

// SearchDetails 确定的搜索(全量搜索)
func SearchDetails(codes []string) (resp []*proto.SharesInfoDetails) {
	if len(codes) == 0 {
		return nil
	}
	parm := strings.Join(codes, ",")
	url := "http://qt.gtimg.cn/q=" + parm
	out := SendGet(url, "")
	out = tools.ConvertToString(out, "gbk", "utf8")
	if len(out) == 0 {
		return nil
	}
	out = strings.Replace(out, "\n", "", -1)

	// todo 分析结果
	list := strings.Split(out, ";")
	for _, v := range list {
		if len(v) < 4 {
			continue
		}
		list1 := strings.Split(v, "=")
		if len(list1) == 2 {
			if strings.HasPrefix(list1[0], "v_") {
				tmp := &proto.SharesInfoDetails{
					Code: list1[0][2:],
				}
				list1[1] = strings.Trim(list1[1], "\"")
				list2 := strings.Split(list1[1], "~")
				if len(list2) > 45 {
					tmp.Ext = _extMp[list2[0]]
					tmp.Name = strings.Replace(list2[1], " ", "", -1)
					tmp.SimpleCode = list2[2]
					tmp.Price = stringToFloat(list2[3])
					tmp.OpenPrice = stringToFloat(list2[5])
					tmp.ClosePrice = tmp.Price
					tmp.Percent = stringToFloat(list2[32])
					tmp.Volume = stringToFloat(list2[36])
					tmp.Turnover = stringToFloat(list2[37])
					tmp.TurnoverRate = stringToFloat(list2[38])
					tmp.Pe = stringToFloat(list2[39])
					tmp.Max = stringToFloat(list2[41])
					tmp.Min = stringToFloat(list2[42])
					tmp.Pb = stringToFloat(list2[46])
					tmp.Total = stringToFloat(list2[44])
					tmp.Cap = stringToFloat(list2[45])
				}
				// 主力资金
				_tmp := getDayZLJLR(SplitCode(tmp.Code))
				if _tmp != nil {
					tmp.Zljlr = _tmp.Price
				}
				// -----
				// macd
				_tmp1 := GetDayMACD(tmp.Code)
				if _tmp1 != nil {
					tmp.Macd = _tmp1.Macd
					tmp.Dif = _tmp1.Dif
					tmp.Dea = _tmp1.Dea
				}
				//
				resp = append(resp, tmp)
			}
		}
	}
	return resp
}

// SearchZLJLR 主力资金净流入
func SearchZLJLR(ext, simplecode string) (resp []*HyInfo) {
	switch ext {
	case "sh":
		ext = "1"
	case "sz":
		ext = "0"
	}
	url := fmt.Sprintf("https://push2his.eastmoney.com/api/qt/stock/fflow/daykline/get?fields1=f1&fields2=f51,f52&secid=%v.%v", ext, simplecode)
	out := SendGet(url, "")
	if len(out) == 0 {
		return nil
	}

	var tmp Daykline
	tools.JSONEncode(out, &tmp)
	// if tmp.Rc != 0 {
	// 	return
	// }

	for _, v := range tmp.Data.Klines {
		tmp := strings.Split(v, ",")
		if len(tmp) == 2 {
			resp = append(resp, &HyInfo{
				Day0:  tools.StrToTime(tmp[0], "2006-01-02", time.Local),
				Price: stringToFloat(tmp[1]) * 0.0001,
			})
		}
	}
	return
}

// getDayZLJLR 当日主力资金净流入
func getDayZLJLR(ext, simplecode string) (resp *HyInfo) {
	switch ext {
	case "sh":
		ext = "1"
	case "sz":
		ext = "0"
	}
	url := fmt.Sprintf("https://push2.eastmoney.com/api/qt/ulist.np/get?secids=%v.%v&fields=f12,f13,f14,f62", ext, simplecode)
	out := SendGet(url, "")
	if len(out) == 0 {
		return nil
	}

	var tmp HyResp
	tools.JSONEncode(out, &tmp)
	if tmp.Rc == 0 && len(tmp.Data.Diffs) > 0 {
		return &HyInfo{
			Day0:  time.Now(),
			Price: tmp.Data.Diffs[0].F62 * 0.0001,
		}
	}

	return
}

// GetDayMACD 当日macd
func GetDayMACD(code string) (resp *MacdInfo) {
	list, err := GetMacd(code)
	if err == nil {
		index := len(list) - 1
		if index >= 0 {
			return &MacdInfo{
				Day0: tools.StrToTime(list[index].DayStr, "2006-01-02", time.Local),
				Macd: list[index].Macd,
				Dif:  list[index].Dif,
				Dea:  list[index].Dea,
			}
		}
	}

	return nil
}

func stringToFloat(src string) float64 {
	v, _ := strconv.ParseFloat(src, 64)
	return v
}

func GetShare(code string, rg bool) (*proto.SharesInfo, error) {
	tmp := &proto.SharesInfo{}
	err := mycache.NewCache(CacheCode).Value(code, tmp)
	if err == nil {
		return tmp, nil
	}

	// 开盘期间实时获取
	day0 := tools.GetUtcDay0(time.Now())
	offset := time.Now().Unix() - day0
	week := tools.GetTimeWeek(day0)
	if (week > 0 && week < 6) && (offset >= (9*60*60+15*60) && offset <= (15*60*60+1*60)) { // 实时获取
		outs := Searchs([]string{code})
		if len(outs) > 0 {
			v := outs[0]
			return &proto.SharesInfo{
				Code:       v.Code,
				SimpleCode: v.SimpleCode,
				Ext:        v.Ext,
				Name:       v.Name,
				Price:      v.Price,
				Percent:    v.Percent,
				Color:      GetColor(rg, v.Percent),
			}, nil
		}
	}
	// ----------------end

	orm := core.Dao.GetDBr()
	v, _ := model.SharesInfoTblMgr(orm.DB).GetFromCode(code)
	return &proto.SharesInfo{
		Code:       v.Code,
		SimpleCode: v.SimpleCode,
		Ext:        v.Ext,
		Name:       v.Name,
		Price:      v.Price,
		Percent:    v.Percent,
		Color:      GetColor(rg, v.Percent),
	}, nil
}

func GetShares(codes []string, rg bool) (out []*proto.SharesInfo, err error) {
	var remain []string
	for _, code := range codes { // 先走缓存拿
		tmp := &proto.SharesInfo{}
		err := mycache.NewCache(CacheCode).Value(code, tmp)
		if err != nil {
			remain = append(remain, code)
		} else {
			tmp.Color = GetColor(rg, tmp.Percent)
			out = append(out, tmp)
		}
	}

	// 开盘期间实时获取
	day0 := tools.GetUtcDay0(time.Now())
	offset := time.Now().Unix() - day0
	week := tools.GetTimeWeek(day0)
	if (week > 0 && week < 6) && (offset >= (9*60*60+15*60) && offset <= (15*60*60+1*60)) { // 实时获取
		outs := Searchs(remain)
		for _, v := range outs {
			out = append(out, &proto.SharesInfo{
				Code:       v.Code,
				SimpleCode: v.SimpleCode,
				Ext:        v.Ext,
				Name:       v.Name,
				Price:      v.Price,
				Percent:    v.Percent,
				Color:      GetColor(rg, v.Percent),
			})
		}

		return out, nil
	}
	// ----------------end

	orm := core.Dao.GetDBr()
	list, _ := model.SharesInfoTblMgr(orm.DB).GetBatchFromCode(remain)
	for _, v := range list {
		out = append(out, &proto.SharesInfo{
			Code:       v.Code,
			SimpleCode: v.SimpleCode,
			Ext:        v.Ext,
			Name:       v.Name,
			Price:      v.Price,
			Percent:    v.Percent,
			Color:      GetColor(rg, v.Percent),
		})
	}

	return out, nil
}

func GetColor(rg bool, percent float64) string {
	if rg && percent >= 0 { // 默认红色
		return "#ff0000"
	}

	if !rg && percent <= 0 { // 默认红色
		return "#ff0000"
	}

	return "#1AAD19" // 绿色
}
