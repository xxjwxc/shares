package analy

import (
	"shares/internal/core"
	"shares/internal/model"
	"shares/internal/service/event"

	"github.com/xxjwxc/public/mylog"
)

// 初始化macd
func initMacd() {
	// https://datacenter-web.eastmoney.com/api/data/v1/get?pageSize=1&pageNumber=1&reportName=RPT_MUTUAL_HOLDSTOCKNORTH_STA&columns=ALL&filter=(SECURITY_CODE=%22600519%22)
	var codes []string
	orm := core.Dao.GetDBr()
	model.SharesInfoTblMgr(orm.DB).Select("code").Find(&codes)
	for _, code := range codes {
		list, err := event.GetMacd(code)
		if err == nil {
			for _, v := range list {
				err = model.SharesDailyTblMgr(orm.Where("code = ? and day0_str = ?", code, v.DayStr)).Updates(map[string]interface{}{"macd": v.Macd, "dif": v.Dif, "dea": v.Dea}).Error
				if err != nil {
					mylog.Error(err)
				}
			}
		}
	}
}

// 更新macd
func updateMacd() {
	// https://datacenter-web.eastmoney.com/api/data/v1/get?pageSize=1&pageNumber=1&reportName=RPT_MUTUAL_HOLDSTOCKNORTH_STA&columns=ALL&filter=(SECURITY_CODE=%22600519%22)
	var codes []string
	orm := core.Dao.GetDBr()
	model.SharesInfoTblMgr(orm.DB).Select("code").Find(&codes)
	for _, code := range codes {
		info := event.GetDayMACD(code)
		if info != nil {
			err := model.SharesDailyTblMgr(orm.Where("code = ? and day0 = ?", code, info.Day0.Unix())).Updates(map[string]interface{}{"macd": info.Macd, "dif": info.Dif, "dea": info.Dea}).Error
			if err != nil {
				mylog.Error(err)
			}
		}
	}
}
