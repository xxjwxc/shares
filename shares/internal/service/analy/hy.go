package analy

import (
	"fmt"
	"shares/internal/core"
	"shares/internal/model"
	"strings"
	"time"

	"github.com/xxjwxc/public/message"
	"github.com/xxjwxc/public/myhttp"
	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/tools"
	"gorm.io/datatypes"
)

// 初始化行业,板块数据
func initHY() {
	orm := core.Dao.GetDBw()
	model.HyInfoTblMgr(orm.DB).Exec(fmt.Sprintf("TRUNCATE TABLE %v;", (&model.HyInfoTbl{}).TableName()))
	out := myhttp.OnGetJSON("https://push2.eastmoney.com/api/qt/clist/get?pn=1&pz=1000&np=1&fields=f12,f14,f62&fid=f62&fs=m:90", "") // 拉取分类
	// out = tools.ConvertToString(out, "gbk", "utf8")
	var info HyResp
	tools.JSONEncode(out, &info)
	if info.Rc == 0 {
		mgrhy := model.HyInfoTblMgr(orm.DB)
		for _, v := range info.Data.Diffs {
			if len(v.F12) == 0 {
				continue
			}
			out, _ := mgrhy.GetFromHyCode(v.F12)
			if out.ID == 0 {
				out.HyCode = v.F12
				out.Name = v.F14
				out.CreatedAt = time.Now()
				mgrhy.Save(&out)
			}
		}
	}
	initHYCode()
}

// 初始化行业,板块数据
func initHYCode() {
	orm := core.Dao.GetDBw()
	hylist, _ := model.HyInfoTblMgr(orm.DB).Gets()
	mpTmp := make(map[string][]string)
	for _, v := range hylist {
		var list []Diff
		for i := 1; i < 20; i++ {
			resp := myhttp.OnGetJSON(fmt.Sprintf("https://push2.eastmoney.com/api/qt/clist/get?pz=1000&np=1&pn=%v&fs=b:%v&fields=f12,f13,f14,f62", i, v.HyCode), "") // 拉取分类
			// resp = tools.ConvertToString(resp, "gbk", "utf8")
			var tmp HyResp
			tools.JSONEncode(resp, &tmp)
			if tmp.Rc == 0 {
				list = append(list, tmp.Data.Diffs...)
			}
			if len(list) == tmp.Data.Total {
				model.HyInfoTblMgr(orm.Where("id = ?", v.ID)).Update(model.HyInfoTblColumns.Num, tmp.Data.Total)
				break
			}
		}

		for _, v1 := range list {
			code := BuildCode(v1.F13, v1.F12)
			mpTmp[code] = append(mpTmp[code], v.Name)
		}
	}

	// 获取各自行业具体股票
	// model.HyGroupTblMgr(orm.DB).Exec(fmt.Sprintf("TRUNCATE TABLE  %v;", (&model.HyGroupTbl{}).TableName()))
	for code, v := range mpTmp {
		// model.HyGroupTblMgr(orm.DB).Save(&model.HyGroupTbl{
		// 	Code:      code,
		// 	HyName:    strings.Join(v, ","),
		// 	CreatedAt: time.Now(),
		// })
		model.SharesInfoTblMgr(orm.Where("code = ?", code)).Update(model.SharesInfoTblColumns.HyName, strings.Join(v, ","))
	}
}

func initHYDaily() {
	orm := core.Dao.GetDBw()
	hylist, _ := model.HyInfoTblMgr(orm.DB).Gets()
	for _, v := range hylist {
		list, err := GetHyDayliyEx(v.HyCode, 90)
		if err != nil {
			mylog.Error(err)
		}
		for _, v := range list {
			info, _ := model.HyDailyTblMgr(orm.DB).FetchUniqueIndexByHyCode(v.Code, v.Day0.Unix())
			if info.ID == 0 {
				info.CreatedBy = "system"
				info.CreatedAt = time.Now()
			}
			info.HyCode = v.Code
			info.HyName = v.Name
			info.Price = v.Price
			info.Percent = v.Percent
			// info.Ma5
			// info.Ma10
			// info.Ma20
			info.Day0 = v.Day0.Unix()
			info.Day0Str = tools.GetDayStr(v.Day0)
			info.Volume = v.Volume
			info.Turnover = v.Turnover
			info.TurnoverRate = v.TurnoverRate
			// info.Zljlr
			info.OpenPrice = v.OpenPrice
			info.ClosePrice = v.ClosePrice
			info.Max = v.Max
			info.Min = v.Min
			model.HyDailyTblMgr(orm.DB).Save(&info)

		}

		// ma5
		for _, v := range list {
			info, _ := model.HyDailyTblMgr(orm.DB).FetchUniqueIndexByHyCode(v.Code, v.Day0.Unix())
			if info.ID == 0 {
				continue
			}
			// ma5
			subQuery := orm.Model(&model.HyDailyTbl{}).Select("price").Where("hy_code = ? and day0 <= ?", v.Code, v.Day0.Unix()).Order("day0 desc")
			orm.Table("(?) as u", subQuery.Limit(5)).Select("avg(price)").Find(&info.Ma5)
			orm.Table("(?) as u", subQuery.Limit(10)).Select("avg(price)").Find(&info.Ma10)
			orm.Table("(?) as u", subQuery.Limit(20)).Select("avg(price)").Find(&info.Ma20)
			orm.Table("(?) as u", subQuery.Limit(60)).Select("avg(price)").Find(&info.Ma60)
			model.HyDailyTblMgr(orm.DB).Save(&info)
		}
	}
}

// 初始化行业主力流入
// func initHYZLJLR() {
// 	orm := core.Dao.GetDBr()
// 	hyList, err := model.HyTblMgr(orm.DB).Gets()
// 	if err != nil {
// 		return
// 	}
// 	mgrhy := model.HyGroupDailyTblMgr(orm.DB)
// 	for _, v := range hyList {
// 		out := myhttp.OnGetJSON(fmt.Sprintf("https://push2his.eastmoney.com/api/qt/stock/fflow/daykline/get?lmt=0&fields1=f3&fields2=f51,f52&secid=90.%v", v.HyCode), "") // 拉取分类
// 		fmt.Println(out)
// 		var info Daykline
// 		tools.JSONEncode(out, &info)
// 		if info.Rc != 0 {
// 			continue
// 		}

// 		for _, v1 := range info.Data.Klines {
// 			tmp := strings.Split(v1, ",")
// 			if len(tmp) == 2 {
// 				day0 := tools.StrToTime(tmp[0], "2006-01-02", time.Local).Unix()
// 				hyinfo, err := mgrhy.Reset().FetchUniqueIndexByCode(v.HyCode, day0)
// 				if err != nil {
// 					continue
// 				}
// 				if hyinfo.ID == 0 {
// 					hyinfo.HyCode = v.HyCode
// 					hyinfo.HyName = v.Name
// 					hyinfo.Day0 = day0
// 					hyinfo.DayStr = tmp[0]
// 					hyinfo.CreatedAt = time.Now()
// 				}
// 				hyinfo.Price = stringToFloat(tmp[1]) * 0.0001

// 				mgrhy.Save(&hyinfo)
// 			}
// 		}

// 	}

// 	// model.ZljlrDailyTblMgr(orm.Group("")).Select("day0").Find(&info.Price)
// 	// https://push2his.eastmoney.com/api/qt/stock/fflow/daykline/get?lmt=0&fields1=f3&fields2=f51,f52&secid=90.BK0477
// }

// 计算行业当天主力流入
func countHYZLJLR() {
	orm := core.Dao.GetDBw()
	day0 := tools.GetUtcDay0(time.Now())                                                                                            // 今天
	out := myhttp.OnGetJSON("https://push2.eastmoney.com/api/qt/clist/get?pn=1&pz=500&np=1&fields=f12,f14,f62&fid=f62&fs=m:90", "") // 拉取分类
	// out = tools.ConvertToString(out, "gbk", "utf8")
	var info HyResp
	tools.JSONEncode(out, &info)
	if info.Rc == 0 {
		mgrhy := model.HyDailyTblMgr(orm.DB)
		for _, v := range info.Data.Diffs {
			if len(v.F12) == 0 {
				continue
			}
			out, _ := mgrhy.Reset().FetchUniqueIndexByHyCode(v.F12, day0)
			if out.ID == 0 {
				out.HyCode = v.F12
				out.HyName = v.F14
				out.Day0 = day0
				out.Day0Str = tools.GetDayStr(tools.UnixToTime(day0))
				out.CreatedAt = time.Now()
			}
			out.Price = v.F62 * 0.0001
			mgrhy.Save(&out)
		}
	}

	// todo:其它
}

// 监听涨停板板块信息
func watchZTB() {
	// orm := core.Dao.GetDBw()
	// out := myhttp.OnGetJSON("http://push2.eastmoney.com/api/qt/clist/get?fid=f3&po=1&pz=500&pn=1&np=1&fltt=2&invt=2&fs=m:0+t:6+f:!2,m:0+t:13+f:!2,m:0+t:80+f:!2,m:1+t:2+f:!2,m:1+t:23+f:!2,m:0+t:7+f:!2,m:1+t:3+f:!2&fields=f3,f12,f13,f14", "") // 拉取分类
	// var info HyZTBResp
	// tools.JSONEncode(out, &info)
	// mp := make(map[string]int)
	// if info.Rc == 0 {
	// 	for _, v := range info.Data.Diffs {
	// 		if v.F3 > 9 { // 大于9%就放上去
	// 			code := BuildCode(v.F13, v.F12)
	// 			info, _ := model.HyGroupTblMgr(orm.DB).GetFromCode(code)
	// 			list := strings.Split(info.HyName, ",")
	// 			for _, v1 := range list {
	// 				mp[v1]++
	// 			}
	// 		}
	// 	}
	// }

	out := myhttp.OnGetJSON("http://10.push2.eastmoney.com/api/qt/clist/get?pn=1&pz=200&po=1&np=1&fltt=2&invt=2&fid=f3&fs=m:90+t:2+f:!50&fields=f3,f8,f12,f14,f104,f105", "") // 拉取分类
	var info HyZTBResp
	tools.JSONEncode(out, &info)
	if info.Rc != 0 {
		return
	}

	orm := core.Dao.GetDBw()
	model.HyUpTblMgr(orm.Where("day = ?", datatypes.Date(time.Now()))).Delete(&model.HyUpTbl{})
	model.AnalyHyTblMgr(orm.DB).Exec(fmt.Sprintf("TRUNCATE TABLE  %v;", (&model.AnalyHyTbl{}).TableName()))
	for _, v := range info.Data.Diffs {
		info, _ := model.HyUpTblMgr(orm.DB).FetchUniqueIndexByName(v.F12, datatypes.Date(time.Now()))
		if info.ID == 0 {
			info.Code = v.F12
			info.Name = v.F14
			info.CreatedAt = time.Now()
			info.Day = datatypes.Date(time.Now())
			info.DayStr = tools.GetDayStr(time.Now())
		}
		info.Percent = v.F3
		info.TurnoverRate = v.F8
		info.Num = v.F104 + v.F105
		info.Up = v.F104
		model.HyUpTblMgr(orm.DB).Save(&info)

		if v.F3 >= 2 { // 更新当天
			model.AnalyHyTblMgr(orm.DB).Save(&model.AnalyHyTbl{
				Code:         v.F12,
				Name:         v.F14,
				Percent:      v.F3,
				TurnoverRate: v.F8,
				Num:          v.F104 + v.F105,
				Up:           v.F104,
				CreatedAt:    time.Now(),
				Day:          datatypes.Date(time.Now()),
				DayStr:       tools.GetDayStr(time.Now()),
			})
		}
	}
}

// 获取行业分时原始数据
func GetHyDayliyEx(code string, count int) ([]HyInfoEx, error) {
	out := SendGet(fmt.Sprintf("https://46.push2his.eastmoney.com/api/qt/stock/kline/get?secid=90.%v&fields1=f1,f2,f3&fields2=f51,f52,f53,f54,f55,f56,f57,f59,f61&klt=101&fqt=1&end=20500101&lmt=%v", code, count), "")
	if len(out) == 0 {
		return nil, message.GetError(message.NotFindError)
	}

	var info Dayhyline
	tools.JSONEncode(out, &info)

	if info.Rc != 0 { // 不成功
		return nil, message.GetError(message.InValidOp)
	}

	var output []HyInfoEx

	for _, v := range info.Data.Klines {
		tmp := strings.Split(v, ",")
		if len(tmp) >= 9 {
			output = append(output, HyInfoEx{
				Code:         info.Data.Code,
				Name:         info.Data.Name,
				Price:        stringToFloat(tmp[2]),
				OpenPrice:    stringToFloat(tmp[1]),
				ClosePrice:   stringToFloat(tmp[2]),
				Percent:      stringToFloat(tmp[7]),
				Volume:       stringToFloat(tmp[5]),
				Turnover:     stringToFloat(tmp[6]),
				TurnoverRate: stringToFloat(tmp[8]),
				Max:          stringToFloat(tmp[3]),
				Min:          stringToFloat(tmp[4]),
				Day0:         tools.StrToTime(tmp[0], "2006-01-02", nil),
			})
		}
	}

	return output, nil
}
