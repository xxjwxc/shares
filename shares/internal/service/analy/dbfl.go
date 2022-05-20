package analy

import (
	"shares/internal/core"
	"shares/internal/model"
	"shares/internal/service/event"
	"time"

	"github.com/xxjwxc/public/tools"
	"gorm.io/datatypes"
)

// 底部放量(底部，且3倍放量)
func dwfl() {
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
	malist, err := model.AnalyDwflTblMgr(orm.DB).Gets() // 保留
	if err != nil {
		return
	}
	for _, v := range malist {
		sharesInfo := event.SearchDetail(v.Code)
		if v.Price < sharesInfo.Price && ((sharesInfo.Price - v.Price) > v.Price*0.1) { // 跌破10% 就剔除
			model.AnalyDwflTblMgr(orm.DB).Delete(&v)
		} else {
			v.Percent = sharesInfo.Percent
			v.NewPrice = sharesInfo.Price
			v.TurnoverRate = sharesInfo.TurnoverRate
			model.AnalyDwflTblMgr(orm.DB).Save(&v) // 保留
		}
	}

	model.SharesInfoTblMgr(orm.DB).Select("code").Find(&codes)
	for _, code := range codes {
		list, err := model.SharesDailyTblMgr(orm.Where("day0 < ?", day0).Order("day0 desc").Limit(30)).GetFromCode(code)
		if err != nil {
			continue
		}
		if len(list) < 30 { // 不需要
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

		// 三倍放量
		right := list[0].Turnover + list[1].Turnover + list[2].Turnover
		left := list[3].Turnover + list[4].Turnover + list[5].Turnover
		if left*3 > right { // 放大3倍量
			continue
		}
		min := getMin(list[0].Price, list[1].Price, list[2].Price, list[3].Price, list[4].Price, list[5].Price) // 最小值
		// 底部(30天，只有3天低于(6天中最低价))
		lowNumber := 0
		for _, v := range list {
			if v.Price < min { // 底部
				lowNumber++
			}
		}
		if lowNumber > 3 { // (20天，只有3天低于当前价格)
			continue
		}

		tmp, _ := model.AnalyDwflTblMgr(orm.DB).GetFromCode(code)
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
		model.AnalyDwflTblMgr(orm.DB).Save(&tmp)
	}
}
