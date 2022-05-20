package analy

import (
	"fmt"
	"shares/internal/core"
	"shares/internal/model"
	"shares/internal/service/event"

	"github.com/xxjwxc/public/tools"
	"gorm.io/datatypes"
)

// 北上
func initDayDailyLHB() {
	var day0s []int64
	orm := core.Dao.GetDBr()
	model.SharesDailyTblMgr(orm.Where("code = ?", "sh000001").Order("day0")).Select("DISTINCT day0").Find(&day0s)
	for _, v := range day0s {
		fmt.Println(v)
		event.DayDailyLHB(v)
	}
}

// 龙虎榜
func maLhb() {
	// 更新龙虎榜数据
	orm := core.Dao.GetDBw()

	var day0 int64
	model.LhbDailyTblMgr(orm.Order("day0 desc").Limit(1)).Select("day0").Find(&day0)

	out, _ := model.NoTradingDayTblMgr(core.Dao.GetDBr().DB).GetFromDay(datatypes.Date(tools.UnixToTime(day0)))
	if out.ID > 0 { // 休息时间
		return
	}

	event.DayDailyLHB(day0)

	// 更新新的
	list1, _ := model.LhbDailyTblMgr(orm.DB).GetFromDay0(day0)
	for _, v := range list1 {
		if v.JgJlr > 0 && v.Jlr > 0 { // 机构净流入才进
			info, _ := model.AnalyLhbTblMgr(orm.DB).FetchUniqueIndexByCode(v.Code, day0)
			if info.ID == 0 {
				info.Code = v.Code
				info.Name = v.Name
				info.Day0 = v.Day0
				info.DayStr = v.DayStr
				info.Price = v.Price
				info.JgJlr = v.JgJlr
				info.Jlr = v.Jlr
				info.CreatedAt = v.CreatedAt
				model.AnalyLhbTblMgr(orm.DB).Save(&info)
			}
		}
	}

	// 删除老的
	list, _ := model.AnalyLhbTblMgr(orm.DB).Gets()
	for _, v := range list {
		sharesInfo := event.SearchDetail(v.Code)
		if v.Price > sharesInfo.Price && ((v.Price - sharesInfo.Price) > v.Price*0.1) { // 未站上10日线
			model.AnalyLhbTblMgr(orm.DB).Delete(&v) // 删除
		} else {
			v.NewPrice = sharesInfo.Price
			v.TurnoverRate = sharesInfo.TurnoverRate
			model.AnalyLhbTblMgr(orm.DB).Save(&v)
		}
	}

}
