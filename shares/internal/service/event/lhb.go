package event

import (
	"fmt"
	"shares/internal/core"
	"shares/internal/model"
	"strings"
	"time"

	"github.com/xxjwxc/public/myhttp"
	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/tools"
)

// 龙虎榜
func DayDailyLHB(day0 int64) error {
	day0Str := tools.GetDayStr(tools.UnixToTime(day0))
	// _, simpleCode := SplitCode(code)
	url := fmt.Sprintf("http://datacenter-web.eastmoney.com/api/data/v1/get?sortColumns=SECURITY_CODE,TRADE_DATE&sortTypes=1,-1&reportName=RPT_DAILYBILLBOARD_DETAILS&columns=SECUCODE,TRADE_DATE,BILLBOARD_NET_AMT&filter=(TRADE_DATE%%3C='%v')(TRADE_DATE%%3E='%v')", day0Str, day0Str)
	// url = "http://datacenter-web.eastmoney.com/api/data/v1/get?sortColumns=SECURITY_CODE,TRADE_DATE&sortTypes=1,-1&reportName=RPT_DAILYBILLBOARD_DETAILS&columns=SECUCODE,TRADE_DATE,BILLBOARD_NET_AMT&filter=(TRADE_DATE%3C=%272021-08-09%27)(TRADE_DATE%3E=%272021-08-09%27)"
	out := myhttp.OnGetJSON(url, "") // 拉取分类
	mylog.Info(out)
	var info LHBResp
	tools.JSONEncode(out, &info)
	if info.Code == 0 {
		orm := core.Dao.GetDBw()
		for _, v := range info.Result.Data { // 先更新数据
			var code string
			tmp := strings.Split(v.SECode, ".")
			if len(tmp) == 2 {
				code = strings.ToLower(tmp[1]) + tmp[0]
				sharesInfo, err := GetShare(code, true)
				if err != nil || len(sharesInfo.Code) == 0 {
					continue
				}

				var jgPrice float64
				// 获取机构净流入
				url = fmt.Sprintf(`http://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_BILLBOARD_DAILYDETAILSBUY&columns=OPERATEDEPT_NAME,NET&filter=(TRADE_DATE='%v')(SECURITY_CODE=%%22%v%%22)`, day0Str, tmp[0])
				out = myhttp.OnGetJSON(url, "") // 拉取分类
				var buy1 LHBResp
				tools.JSONEncode(out, &buy1)
				if buy1.Code == 0 {
					for _, v1 := range buy1.Result.Data {
						if strings.EqualFold(v1.OperatedeptName, "机构专用") {
							jgPrice += v1.NET
						}
					}
				}

				// 机构净卖出
				url = fmt.Sprintf(`http://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_BILLBOARD_DAILYDETAILSSELL&columns=OPERATEDEPT_NAME,NET&filter=(TRADE_DATE='%v')(SECURITY_CODE=%%22%v%%22)`, day0Str, tmp[0])
				out = myhttp.OnGetJSON(url, "") // 拉取分类
				var buy2 LHBResp
				tools.JSONEncode(out, &buy2)
				if buy2.Code == 0 {
					for _, v1 := range buy2.Result.Data {
						if strings.EqualFold(v1.OperatedeptName, "机构专用") {
							jgPrice += v1.NET
						}
					}
				}

				dbInfo, err := model.LhbDailyTblMgr(orm.DB).FetchUniqueIndexByCode(code, day0)
				if err != nil {
					return err
				}
				if dbInfo.ID == 0 {
					dbInfo = model.LhbDailyTbl{
						Code:      code,             // 股票代码
						Name:      sharesInfo.Name,  // 股票名字
						Price:     sharesInfo.Price, // 当前价格
						Day0:      day0,             // 当日0点时间戳
						DayStr:    day0Str,          // 入池时间
						Jlr:       v.Amt * 0.0001,   // 龙虎榜净流入(万元)
						JgJlr:     jgPrice * 0.0001, // 机构净流入(万元)
						CreatedAt: time.Now(),
					}
				} else {
					dbInfo.Jlr = v.Amt * 0.0001
					dbInfo.JgJlr = jgPrice * 0.0001
				}
				model.LhbDailyTblMgr(orm.DB).Save(&dbInfo)
			}
		}
	} else {
		mylog.Errorf("龙虎榜未找到:%v,%v", day0Str, url)
	}
	return nil
}
