package analy

import (
	"shares/internal/core"
	"shares/internal/model"
	"shares/internal/service/event"
	"strings"
	"time"

	"github.com/xxjwxc/public/tools"
	"gorm.io/datatypes"
)

// 放量(放量超过前一天3倍，且成交量>10%)
func watchMyself() {
	if !event.IsWorkDay() {
		return
	}
	for {
		dealMyself()         // 放量
		if !continueTime() { // 3点之后是空闲时间
			break
		} else {
			time.Sleep(10 * time.Minute)
		}
	}

	deleteMyself() // 剔除一次
}

func dealMyself() {
	var codes, tmp []string
	orm := core.Dao.GetDBw()
	model.AnalyLhbTblMgr(orm.DB).Select("code").Find(&tmp) // 龙虎榜
	codes = append(codes, tmp...)
	model.AnalyFlTblMgr(orm.DB).Select("code").Find(&tmp) // 放量
	codes = append(codes, tmp...)
	model.AnalyCdTblMgr(orm.DB).Select("code").Find(&tmp) // 超跌
	codes = append(codes, tmp...)
	model.AnalyUpTblMgr(orm.DB).Select("code").Find(&tmp) // 趋势
	codes = append(codes, tmp...)
	model.AnalyDbszxTblMgr(orm.DB).Select("code").Find(&tmp) // 底部缩量十字星
	codes = append(codes, tmp...)
	model.AnalyDwflTblMgr(orm.DB).Select("code").Find(&tmp) // 低位放量
	codes = append(codes, tmp...)
	model.AnalyTkTblMgr(orm.DB).GetBatchFromCode(codes) // 跳空回补
	codes = append(codes, tmp...)

	mp := getMp(codes)
	for code, v := range mp {
		if len(v) > 2 { // 多个
			sharesInfo := event.SearchDetail(code)
			info, _ := model.MyselfTblMgr(orm.DB).GetFromCode(code)
			if info.ID == 0 {
				info.Code = code
				info.Day = datatypes.Date(time.Now())
				info.DayStr = tools.GetDayStr(time.Now())
				info.CreatedAt = time.Now()
				info.Price = sharesInfo.Price
			}
			info.Name = sharesInfo.Name
			info.Percent = sharesInfo.Percent
			info.NewPrice = sharesInfo.Price
			info.TurnoverRate = sharesInfo.TurnoverRate
			info.Desc = strings.Join(v, ";") // 100分
			info.Number = len(v)
			model.MyselfTblMgr(orm.DB).Save(&info)
		}
	}
}

func deleteMyself() {
	orm := core.Dao.GetDBw()
	malist, err := model.MyselfTblMgr(orm.DB).Gets() // 保留
	if err != nil {
		return
	}
	for _, v := range malist {
		sharesInfo := event.SearchDetail(v.Code)
		if v.Price < sharesInfo.Price && ((sharesInfo.Price - v.Price) > v.Price*0.1) { // 跌破10% 就剔除
			model.MyselfTblMgr(orm.DB).Delete(&v)
		} else {
			v.Percent = sharesInfo.Percent
			v.NewPrice = sharesInfo.Price
			v.TurnoverRate = sharesInfo.TurnoverRate
			model.MyselfTblMgr(orm.DB).Save(&v) // 保留
		}
	}
}
