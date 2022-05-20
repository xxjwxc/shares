package analy

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

type XueQiuMsgResp struct {
	Status  int64           `json:"code"`
	Code    string          `json:"key"` // code
	List    []XueQiuMsgInfo `json:"list"`
	MaxPage int             `json:"maxPage"`
}

type XueQiuMsgInfo struct {
	CreatedAt   int64  `json:"created_at"`
	Description string `json:"description"` // 简短描述
	Title       string `json:"title"`       // 标题
	Text        string `json:"text"`        // 内容
	UserId      int64  `json:"user_id"`     // 用户id
}

type BaiduResp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

type BaiduLexerResp struct {
	ErrorCode int         `json:"error_code"`
	ErrorMsg  string      `json:"error_msg"`
	Text      string      `json:"text"`
	Items     []BaiduItem `json:"items"`
}

type BaiduItem struct {
	Item string `json:"item"`
	Ne   string `json:"ne"`
	Pos  string `json:"Pos"`
}

type BaiduLexerReq struct {
	Text string `json:"text"`
}

type SentimentResp struct {
	Text  string      `json:"text"`
	Items []Sentiment `json:"items"`
}

type Sentiment struct {
	Sentiment    int     `json:"sentiment"`     // 表示情感极性分类结果，0:负向，1:中性，2:正向
	Confidence   float64 `json:"confidence"`    // 表示分类的置信度，取值范围[0,1]
	PositiveProb float64 `json:"positive_prob"` // 表示属于积极类别的概率 ，取值范围[0,1]
	NegativeProb float64 `json:"negative_prob"` // 表示属于消极类别的概率，取值范围[0,1]
}

type dealMsgResp struct {
	Code  string
	Count int
}

type TencentResp struct {
	Response TencentResponse `json:"Response"`
}

type TencentResponse struct {
	Keywords  []TencentItem `json:"Keywords"`
	RequestId string        `json:"RequestId"`
}

type TencentItem struct {
	Item  string  `json:"Word"`
	Score float64 `json:"Score"`
}

/*slice 排序示例*/
type MsgSort struct {
	Msg string
	Sum int
}

type MsgSortSlice []MsgSort

func (s MsgSortSlice) Len() int           { return len(s) }
func (s MsgSortSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s MsgSortSlice) Less(i, j int) bool { return s[i].Sum > s[j].Sum }

type HyZTBResp struct {
	Rc   int   `json:"rc"`
	Data HyZTB `json:"data"`
}

type HyZTB struct { // 9% 涨停板数据
	Total int       `json:"total"`
	Diffs []DiffZTB `json:"diff"`
}

type DiffZTB struct {
	F3   float64 `json:"f3"`  // 涨跌幅
	F8   float64 `json:"f8"`  // 换手率
	F12  string  `json:"f12"` // 行业代码
	F13  int     `json:"f13"`
	F14  string  `json:"f14"`  // 行业名
	F104 int     `json:"f104"` // 上涨家数
	F105 int     `json:"f105"` // 下跌家数
}

type HyInfoEx struct {
	Code         string  `json:"code"`         // 股票代码
	Name         string  `json:"name"`         // 股票名字
	Price        float64 `json:"price"`        // 当前价格
	Percent      float64 `json:"percent"`      // 百分比
	OpenPrice    float64 `json:"openPrice"`    // 开盘价
	ClosePrice   float64 `json:"closePrice"`   // 收盘价
	Volume       float64 `json:"volume"`       // 成交量（手）
	Turnover     float64 `json:"turnover"`     // 成交额（万）
	TurnoverRate float64 `json:"turnoverRate"` // 换手率
	Max          float64 `json:"max"`          // 当日最高点
	Min          float64 `json:"min"`          // 当日最低
	Day0         time.Time
}
type HyDataEx struct {
	Code   string   `json:"code"` // 股票代码
	Name   string   `json:"name"` // 股票名字
	Klines []string `json:"klines"`
}
type Dayhyline struct {
	Rc   int      `json:"rc"`
	Data HyDataEx `json:"data"`
}

type HyPeResp struct {
	Result HyPeResult `json:"result"`
	Code   int        `json:"code"`
}

type HyPeResult struct {
	Data []HyPeData `json:"data"`
}

type HyPeData struct {
	HyName         string     `json:"BOARD_NAME"`     // 行业名字
	HyCode         string     `json:"ORIGINALCODE"`   // 行业简码
	TradeDate      tools.Time `json:"TRADE_DATE"`     // 交易日期
	PETTM          float64    `json:"PE_TTM"`         // PE(ttm)
	PE_LAR         float64    `json:"PE_LAR"`         // PE静态
	PB_MRQ         float64    `json:"PB_MRQ"`         // 市净率
	PEG_CAR        float64    `json:"PEG_CAR"`        // Peg
	MARKET_CAP_VAG float64    `json:"MARKET_CAP_VAG"` // 平均市值
	NUM            int        `json:"NUM"`            // 数量
	LOSS_COUNT     int        `json:"LOSS_COUNT"`     // 亏损家数
}
