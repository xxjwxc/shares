package analy

import (
	"shares/internal/core"
	"shares/internal/model"
	"shares/internal/service/event"
	"time"

	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/tools"
)

func tkgk() {
	if !event.IsWorkDay() {
		return
	}
	analyDaily() // 先分析
	tkDaily()    // 再拉取一次
}

// 分析跳空信息
func analyDaily() {
	orm := core.Dao.GetDBw()
	// 先删除
	list, _ := model.AnalyTkTblMgr(orm.DB).Gets() // 保留
	for _, v := range list {
		// sharesInfo := event.SearchDetail(v.Code)
		sharesInfo, err := event.GetShare(v.Code, true)
		if err != nil {
			mylog.Error(err)
			continue
		}
		if v.Price < sharesInfo.Price && ((sharesInfo.Price - v.Price) > v.Price*0.1) { // 跌破10% 就剔除
			model.AnalyTkTblMgr(orm.DB).Delete(&v)
		} else {
			model.AnalyTkTblMgr(orm.Where("id = ?", v.ID)).Update("percent", sharesInfo.Percent)
		}
	}

	var codes []string
	model.TkTblMgr(orm.DB).Select("DISTINCT code").Find(&codes)
	for _, code := range codes {
		sharesInfo := event.SearchDetail(code)

		// 判断是否回补跳空缺口
		list, _ := model.TkTblMgr(orm.Where("flag = 0")).GetFromCode(code)
		for _, v := range list {
			if v.Min > sharesInfo.Min { // 已经回补了缺口
				info, _ := model.AnalyTkTblMgr(orm.DB).FetchUniqueIndexByCode(v.Code, v.Day0)
				info.Code = v.Code
				info.Name = sharesInfo.Name
				info.Day0 = v.Day0
				info.Day0Str = v.Day0Str
				info.Price = sharesInfo.Price
				info.Min = v.Min
				info.Max = v.Max
				info.CreatedAt = time.Now()
				model.AnalyTkTblMgr(orm.DB).Save(&info)

				model.TkTblMgr(orm.Where("id = ?", v.ID)).Delete(&model.TkTbl{})
			}
		}
	}
}

// 获取跳空高开列表
func tkDaily() {
	if !event.IsWorkDay() {
		return
	}

	var codes []string
	orm := core.Dao.GetDBw()
	day0 := tools.GetUtcDay0(time.Now())
	model.SharesInfoTblMgr(orm.Where("percent > 0")).Select("code").Find(&codes)

	for _, code := range codes {
		list, err := model.SharesDailyTblMgr(orm.Where("day0 <= ?", day0).Order("day0 desc").Limit(2)).GetFromCode(code)
		if err != nil {
			continue
		}
		if len(list) < 2 { // 不需要
			continue
		}

		if list[1].Max < list[0].Min && list[1].Max > 0 && (list[0].Min-list[1].Max) > (list[0].Price*0.01) { // 跳空高开
			sharesInfo, err := event.GetShare(code, true)
			if err != nil {
				mylog.Error(err)
				continue
			}

			info, _ := model.TkTblMgr(orm.DB).FetchUniqueIndexByCode(code, day0)
			info.Code = code
			info.Name = sharesInfo.Name
			info.Day0 = day0
			info.Day0Str = tools.GetDayStr(tools.UnixToTime(day0))
			info.Price = sharesInfo.Price
			info.Min = list[1].Max
			info.Max = list[0].Min
			info.Flag = 0
			info.CreatedAt = time.Now()
			info.UpdateAt = time.Now()
			model.TkTblMgr(orm.DB).Save(&info)

		}
	}

}
