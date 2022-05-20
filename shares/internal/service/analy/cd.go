package analy

import (
	"shares/internal/core"
	"shares/internal/model"
	"shares/internal/service/event"
	"time"

	"github.com/xxjwxc/public/tools"
)

// 超跌反弹(20天最低点,倒数3天收阳)
func chaoDie() {
	var codes []string
	orm := core.Dao.GetDBw()
	day0 := tools.GetUtcDay0(time.Now())
	day1Flag := false // 是否补当前的行情
	offset := time.Now().Unix() - day0
	if offset >= (14 * 60 * 60) { // 2点之后后一天
		// day0 += 24 * 60 * 60
		day1Flag = true
	}

	if offset >= (15 * 60 * 60) {
		day0 += 24 * 60 * 60
		day1Flag = false
	}

	malist, err := model.AnalyCdTblMgr(orm.DB).Gets() // 保留
	if err != nil {
		return
	}
	for _, v := range malist {
		sharesInfo := event.SearchDetail(v.Code)
		if v.Price < sharesInfo.Price && ((sharesInfo.Price - v.Price) > v.Price*0.1) { // 跌破10% 就剔除
			model.AnalyCdTblMgr(orm.DB).Delete(&v)
		} else {
			v.NewPrice = sharesInfo.Price
			v.TurnoverRate = sharesInfo.TurnoverRate
			model.AnalyCdTblMgr(orm.DB).Save(&v) // 保留
		}

		// list, _ := model.SharesDailyTblMgr(orm.Where("day0 < ?", day0).Order("day0 desc").Limit(3)).GetFromCode(v.Code)
		// if len(list) == 0 {
		// 	continue
		// }
		// if list[0].Price > list[0].Ma10 { // 站上10日线
		// 	sharesInfo := event.SearchDetail(v.Code)
		// 	v.NewPrice = sharesInfo.Price
		// 	v.TurnoverRate = sharesInfo.TurnoverRate
		// 	model.AnalyCdTblMgr(orm.DB).Save(&v) // 保留
		// } else {
		// 	model.AnalyCdTblMgr(orm.DB).Delete(&v)
		// }
	}

	orm.Table("(SELECT code,count(*) as num FROM `shares_daily_tbl` GROUP BY code) as u").Where("num > 19").Select("code").Find(&codes)
	for _, code := range codes {
		list, err := model.SharesDailyTblMgr(orm.Where("day0 < ?", day0).Order("day0 desc").Limit(20)).GetFromCode(code)
		if err != nil {
			continue
		}
		if len(list) < 20 { // 不需要
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

		// 倒数三天(收阳线)
		if list[0].Price >= list[1].Price && list[1].Price >= list[2].Price &&
			list[0].OpenPrice <= list[0].ClosePrice && list[1].OpenPrice <= list[1].ClosePrice && list[2].OpenPrice <= list[2].ClosePrice {
			// 前4天有近20天最低点
			var arrays []float64
			for _, v := range list {
				arrays = append(arrays, v.Price)
			}

			min0 := getMin(list[0].Price, list[1].Price, list[2].Price, list[3].Price)
			min1 := getMin(arrays...)
			if min0 == min1 { // 找到了
				sharesInfo, err := event.GetShare(code, true)
				if err != nil {
					continue
				}
				// if len(sharesInfo.Code) == 0 { // 删除
				// 	model.SharesDailyTblMgr(orm.Where("code = ?", code)).Delete(&model.SharesDailyTbl{})
				// 	continue
				// }
				tmp, _ := model.AnalyCdTblMgr(orm.DB).GetFromCode(code)
				if tmp.ID == 0 {
					tmp.Code = code
					tmp.Day0 = tools.GetUtcDay0(time.Now())
					tmp.DayStr = list[0].Day0Str
					tmp.CreatedAt = time.Now()
				} else {
					tmp.CreatedAt = time.Now()
				}
				tmp.Price = sharesInfo.Price
				tmp.Name = sharesInfo.Name
				tmp.Percent = sharesInfo.Percent
				tmp.TurnoverRate = list[0].TurnoverRate
				tmp.Score = 100 // 100分
				model.AnalyCdTblMgr(orm.DB).Save(&tmp)
			}
		}
	}
}

// 初始化 vol
func initVol() {
	orm := core.Dao.GetDBw()
	id := 0
	for {
		list, _ := model.SharesDailyTblMgr(orm.Where("id > ? and volume > 0", id).Order("id asc").Limit(100)).Gets()
		if len(list) == 0 {
			break
		}

		for _, v := range list {
			var vol5 float64
			subQuery := orm.Model(&model.SharesDailyTbl{}).Select("volume").Where("code = ? and day0 < ?", v.Code, v.Day0).Order("day0 desc")
			orm.Table("(?) as u", subQuery.Limit(5)).Select("avg(volume)").Find(&vol5)
			if vol5 > 0 {
				model.SharesDailyTblMgr(orm.Where("id = ?", v.ID)).Update(model.SharesDailyTblColumns.Vol, (v.Volume / vol5))
			}
			if id < v.ID {
				id = v.ID
			}
		}
	}
}
