package event

import (
	"time"

	"github.com/xxjwxc/public/tools"
)

const (
	CacheCode        = "shares_code_cache"
	CacheCodeTimeout = 2 * time.Hour
	_timeoutTicker   = 3
)

var _extMp = map[string]string{
	"200": "us", // 美股
	"100": "hk", // 港股
	"51":  "sz", // 深圳
	"1":   "sh", // 上海
}

type DailyInfo struct {
	Day [][]interface{} `json:"day"`
}

type DailyData struct {
	Info DailyInfo `json:"info"`
}
type DailyWWW struct {
	Code int       `json:"code"`
	Data DailyData `json:"data"`
}

type HyGroupInfo struct {
	Code string `json:"symbol"`
}

type HyData struct {
	Klines []string `json:"klines"`
}
type Daykline struct {
	Rc   int    `json:"rc"`
	Data HyData `json:"data"`
}

type HyInfo struct {
	Day0  time.Time
	Price float64
}

type MacdInfo struct {
	Day0 time.Time
	Dif  float64
	Dea  float64
	Macd float64
}

type HyResp struct {
	Rc   int        `json:"rc"`
	Data HyRespData `json:"data"`
}

type HyRespData struct {
	Total int    `json:"total"`
	Diffs []Diff `json:"diff"`
}

type Diff struct {
	F12 string  `json:"f12"`
	F13 int     `json:"f13"`
	F14 string  `json:"f14"`
	F62 float64 `json:"f62"`
}

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

type MinuteWWW struct {
	Code int        `json:"code"`
	Data MinuteData `json:"data"`
}

type MinuteData struct {
	Info struct {
		Data struct {
			Data []string `json:"data"`
		} `json:"data"`
		QT struct {
			Info []string `json:"info"`
		} `json:"qt"`
	} `json:"info"`
}

type MinuteInfo struct {
	Time  string  // 时间
	Price float64 // 当前价
	Vol   int64   // 总成交量
	Ave   float64 //  均价
}

type Minute struct {
	List     []MinuteInfo
	PrePrice float64
}

type LHBResp struct {
	Result LHBResult `json:"result"`
	Code   int       `json:"code"`
}

type LHBResult struct {
	Data []LHBData `json:"data"`
}

type LHBData struct {
	SECode          string     `json:"SECUCODE"`          // 股票代码
	TradeDate       tools.Time `json:"TRADE_DATE"`        // 日期
	Amt             float64    `json:"BILLBOARD_NET_AMT"` // 龙虎榜净买额
	OperatedeptName string     `json:"OPERATEDEPT_NAME"`  // 节点名字
	NET             float64    `json:"NET"`               // 净额
}
