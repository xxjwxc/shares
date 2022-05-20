package analy

import (
	"fmt"
	"shares/internal/core"
	"shares/internal/model"
	"shares/internal/service/event"

	"github.com/xxjwxc/public/message"
	"github.com/xxjwxc/public/myhttp"
	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/tools"
)

type BSResp struct {
	Result BSResult `json:"result"`
	Code   int      `json:"code"`
}

type BSResult struct {
	Data []BSData `json:"data"`
}

type BSData struct {
	ClosePrice float64    `json:"CLOSE_PRICE"`    // 收盘价
	TradeDate  tools.Time `json:"TRADE_DATE"`     // 交易日期
	Percent    float64    `json:"A_SHARES_RATIO"` // 北上持股占比
	Bscg       float64    `json:"HOLD_SHARES"`    // 北上持股数
}

// 初始化北上数据
func updateBS() {
	// https://datacenter-web.eastmoney.com/api/data/v1/get?pageSize=1&pageNumber=1&reportName=RPT_MUTUAL_HOLDSTOCKNORTH_STA&columns=ALL&filter=(SECURITY_CODE=%22600519%22)
	var codes []string
	orm := core.Dao.GetDBr()
	model.SharesInfoTblMgr(orm.DB).Select("code").Find(&codes)
	for _, code := range codes {
		CountBSJLR(code, 1)
	}
}

// 初始化北上数据
func initBS() {
	// https://datacenter-web.eastmoney.com/api/data/v1/get?pageSize=1&pageNumber=1&reportName=RPT_MUTUAL_HOLDSTOCKNORTH_STA&columns=ALL&filter=(SECURITY_CODE=%22600519%22)
	var codes []string
	orm := core.Dao.GetDBr()
	model.SharesInfoTblMgr(orm.DB).Select("code").Find(&codes)
	for _, code := range codes {
		event.CountBSJLR(code, 0)
	}
}

// 计算北上净流入(page = 0 重新计算)
func CountBSJLR(code string, page int) error {
	_, simpleCode := SplitCode(code)
	url := fmt.Sprintf("https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_MUTUAL_HOLDSTOCKNORTH_STA&columns=CLOSE_PRICE,TRADE_DATE,HOLD_SHARES,A_SHARES_RATIO&filter=(SECURITY_CODE=%v)", simpleCode)
	if page > 0 {
		url = fmt.Sprintf("https://datacenter-web.eastmoney.com/api/data/v1/get?pageSize=%v&reportName=RPT_MUTUAL_HOLDSTOCKNORTH_STA&columns=CLOSE_PRICE,TRADE_DATE,HOLD_SHARES,A_SHARES_RATIO&filter=(SECURITY_CODE=%v)", page, simpleCode)
	}
	out := myhttp.OnGetJSON(url, "") // 拉取分类
	mylog.Info(out)
	var info BSResp
	tools.JSONEncode(out, &info)
	if info.Code == 0 {
		orm := core.Dao.GetDBw()
		for _, v := range info.Result.Data { // 先更新数据
			err := model.SharesDailyTblMgr(orm.Where("`code` = ? AND `day0` = ?", code, v.TradeDate.Unix())).Updates(map[string]interface{}{
				model.SharesDailyTblColumns.Bscg:      v.Bscg,    // 北上持股数
				model.SharesDailyTblColumns.BsPercent: v.Percent, // 北上持股占比
			}).Error
			if err != nil {
				mylog.Error(err)
			}
		}

		// 计算净流入
		for _, v := range info.Result.Data {
			info, err := model.SharesDailyTblMgr(orm.Where("`code` = ? AND `day0` < ?", code, v.TradeDate.Unix()).Order("day0 desc").Limit(1)).Get()
			if err != nil {
				mylog.Error(err)
			}
			if info.ID > 0 && info.Bscg > 0 {
				bsjlr := (v.Bscg - info.Bscg) * v.ClosePrice
				model.SharesDailyTblMgr(orm.Where("`code` = ? AND `day0` = ?", code, v.TradeDate.Unix())).Update(model.SharesDailyTblColumns.Bsjlr, bsjlr)
			}
		}
	} else {
		if info.Code == 9201 {
			return message.GetError(message.NotFindError)
		}
	}
	return nil
}
