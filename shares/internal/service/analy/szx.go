package analy

import (
	"shares/internal/core"
	"shares/internal/model"
	"shares/internal/service/event"
	"time"

	"github.com/xxjwxc/public/tools"
	"gorm.io/datatypes"
)

// 底部缩量十字星
func dbSzx() {
	var codes []string
	orm := core.Dao.GetDBw()
	day0 := tools.GetUtcDay0(time.Now())
	day1Flag := true // 是否补当前的行情
	offset := time.Now().Unix() - day0

	if offset <= (14*60*60 + 30*60) { // 2点之后后一天
		return
	}

	if offset >= (14 * 60 * 60) { // 2点之后后一天
		day1Flag = true
	}

	if offset >= (15 * 60 * 60) {
		day0 += 24 * 60 * 60
		day1Flag = false
	}

	// model.AnalyFlTblMgr(orm.Where("1 = 1")).Update("tag", 2) // 待删除
	malist, err := model.AnalyDbszxTblMgr(orm.DB).Gets() // 保留
	if err != nil {
		return
	}
	for _, v := range malist {
		sharesInfo := event.SearchDetail(v.Code)
		if v.Price < sharesInfo.Price && ((sharesInfo.Price - v.Price) > v.Price*0.1) { // 跌破10% 就剔除
			model.AnalyDbszxTblMgr(orm.DB).Delete(&v)
		} else {
			v.Percent = sharesInfo.Percent
			v.NewPrice = sharesInfo.Price
			v.TurnoverRate = sharesInfo.TurnoverRate
			model.AnalyDbszxTblMgr(orm.DB).Save(&v) // 保留
		}
	}

	model.SharesInfoTblMgr(orm.DB).Select("code").Find(&codes)
	for _, code := range codes {
		list, err := model.SharesDailyTblMgr(orm.Where("day0 < ?", day0).Order("day0 desc").Limit(20)).GetFromCode(code)
		if err != nil {
			continue
		}
		if len(list) < 20 { // 不需要
			continue
		}

		sharesInfo := event.SearchDetail(code)
		if day1Flag {
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

		// 底部缩量十字星
		// 底部(20天，只有3天低于当前价格)
		lowNumber := 0
		lowSl := 0
		for _, v := range list {
			if v.Price < sharesInfo.Price { // 底部
				lowNumber++
			}
			if v.TurnoverRate > 0 && v.TurnoverRate < sharesInfo.TurnoverRate { // 缩量
				lowSl++
			}
		}
		if lowNumber > 3 { // (20天，只有3天低于当前价格)
			continue
		}

		if lowSl > 3 { // (20天，只有3天低于当前量能)
			continue
		}

		// 十字星
		left := sharesInfo.OpenPrice
		right := sharesInfo.Price
		if left > right {
			left = sharesInfo.Price
			right = sharesInfo.OpenPrice
		}
		p1 := left - sharesInfo.Min
		p2 := right - left
		p3 := sharesInfo.Max - right
		// 百分比小于0.5
		if (p2 < sharesInfo.OpenPrice*0.01) && (p2*3 < p1) && (p1 > p3) { // 十字星
			tmp, _ := model.AnalyDbszxTblMgr(orm.DB).GetFromCode(code)
			if tmp.ID == 0 {
				tmp.Code = code
				tmp.Day = datatypes.Date(time.Now())
				tmp.DayStr = list[0].Day0Str
				tmp.CreatedAt = time.Now()
			}
			tmp.Price = sharesInfo.Price
			tmp.Name = sharesInfo.Name
			tmp.Percent = sharesInfo.Percent
			tmp.TurnoverRate = list[0].TurnoverRate
			model.AnalyDbszxTblMgr(orm.DB).Save(&tmp)
		}

	}
}
