package analy

import (
	"shares/internal/core"
	"shares/internal/model"
	"shares/internal/service/event"
	"time"

	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/tools"
)

func init() {
	// initHYZLJLR()
	// countHYZLJLR(tools.GetUtcDay0(time.Now()))
	// if config.GetIsTools() {
	// 	// initHY() // 初始化板块信息
	// 	// initHYDaily() // 更新所有主力净流入
	// 	// initHYZLJLR() // 更新所有行业净流入
	// 	startTools()
	// }
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func toolsTicket() { // 只更新当天的数据
	orm := core.Dao.GetDBw()
	for {
		ticker := time.NewTicker(_timeoutTicker * time.Second * 3)
		<-ticker.C
		list, err := model.AnalyUpTblMgr(orm.DB).GetFromDay0(tools.GetUtcDay0(time.Now()))
		if err != nil {
			mylog.Error(err)
			continue
		}

		for _, v := range list {
			sharesInfo, err := event.GetShare(v.Code, true)
			if err != nil {
				continue
			}
			// v.Price = sharesInfo.Price
			// // v.Name = sharesInfo.Name
			// v.Percent = sharesInfo.Percent
			orm.Model(v).Updates(map[string]interface{}{"price": sharesInfo.Price, "percent": sharesInfo.Percent})
			// model.AnalyUpTblMgr(orm.DB).Save(v)
		}
	}
}

func initmaDaily() {
	var codes []string
	orm := core.Dao.GetDBw()
	day0 := tools.GetUtcDay0(time.Now())
	offset := time.Now().Unix() - day0
	if offset >= (14 * 60 * 60) { // 2点之前为前一天
		day0 += 24 * 60 * 60
	}

	// model.AnalyUpTblMgr(orm.Where("1 = 1")).Update("tag", 2) // 待删除
	// 保留
	malist, err := model.AnalyUpTblMgr(orm.DB).Gets()
	if err != nil {
		return
	}
	for _, v := range malist {
		sharesInfo, err := event.GetShare(v.Code, true)
		if err != nil {
			continue
		}
		list, err := model.SharesDailyTblMgr(orm.Where("day0 < ?", day0).Order("day0 desc").Limit(3)).GetFromCode(v.Code)
		if err != nil || len(list) < 3 {
			continue
		}
		// if list[0].Ma5 >= list[0].Ma10 && list[0].Ma10 >= list[0].Ma20 && sharesInfo.Price > list[0].Ma10 { // 依然是多头排列
		if sharesInfo.Price > list[0].Ma10 { // 站上10日线
			sharesInfo := event.SearchDetail(v.Code)
			v.NewPrice = sharesInfo.Price
			v.TurnoverRate = sharesInfo.TurnoverRate
			model.AnalyUpTblMgr(orm.DB).Save(&v)
		} else {
			model.AnalyUpTblMgr(orm.DB).Delete(&v)
		}
	}

	orm.Table("(SELECT code,count(*) as num FROM `shares_daily_tbl` GROUP BY code) as u").Where("num > 19").Select("code").Find(&codes)
	for _, code := range codes {
		list, err := model.SharesDailyTblMgr(orm.Where("day0 < ?", day0).Order("day0 desc").Limit(3)).GetFromCode(code)
		if err != nil {
			continue
		}
		if len(list) < 3 { // 不需要
			continue
		}
		sharesInfo, err := event.GetShare(code, true)
		if err != nil {
			continue
		}
		// 多头排列 todo:且连续2天主力净流入
		if !(list[1].Ma5 >= list[1].Ma10 && list[1].Ma10 >= list[1].Ma20) { // 上上日是非多头排列
			if list[0].Ma5 >= list[0].Ma10 && list[0].Ma10 >= list[0].Ma20 && sharesInfo.Price >= list[0].Ma5 &&
				list[0].Macd > list[1].Macd && list[0].Dea < list[0].Dif { // 上一日是多头排列(提醒)，且金叉
				if fixTurnoverRate(list[0].TurnoverRate, list[1].TurnoverRate, list[2].TurnoverRate) {
					tmp, _ := model.AnalyUpTblMgr(orm.DB).GetFromCode(code)
					if tmp.ID == 0 {
						tmp.Code = code
						tmp.Day0 = tools.GetUtcDay0(time.Now())
						tmp.DayStr = list[0].Day0Str
						tmp.CreatedAt = time.Now()
						tmp.Price = sharesInfo.Price
						tmp.TurnoverRate = list[0].TurnoverRate
					} else {
						tmp.CreatedAt = time.Now()
					}
					tmp.Name = sharesInfo.Name
					tmp.Percent = sharesInfo.Percent
					// tmp.NewPrice = sharesInfo.Price
					model.AnalyUpTblMgr(orm.DB).Save(&tmp)
				}
			}
		}
	}
}

func fixTurnoverRate(a1, b1, c1 float64) bool {
	if a1 < 1 && b1 < 1 && c1 < 1 {
		return true
	}

	a := int64(a1 * 100)
	b := int64(b1 * 100)
	c := int64(c1 * 100)
	if a > b {
		a = a ^ b
		b = a ^ b
		a = a ^ b
	}
	if b > c {
		b = b ^ c
		c = b ^ c
		b = b ^ c
	}
	if a > b {
		a = a ^ b
		b = a ^ b
		a = a ^ b
	}

	if a*15 < c*10 {
		return true
	}
	return false
}
