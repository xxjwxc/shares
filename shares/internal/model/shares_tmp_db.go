package model

import (
	"time"

	"gorm.io/datatypes"
)

// AnalyCdTbl 超跌赛选器
type AnalyCdTbl struct {
	ID           int       `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Code         string    `gorm:"unique;column:code;type:varchar(255);default:null" json:"code"`           // 股票代码
	Name         string    `gorm:"column:name;type:varchar(255);default:null" json:"name"`                  // 股票名字
	Day0         int64     `gorm:"column:day0;type:bigint(20);default:null" json:"day0"`                    // 入池时间
	DayStr       string    `gorm:"column:day_str;type:varchar(255);default:null" json:"dayStr"`             // 入池时间
	Price        float64   `gorm:"column:price;type:double(11,2);default:null" json:"price"`                // 价格
	NewPrice     float64   `gorm:"column:new_price;type:double(11,2);default:null" json:"newPrice"`         // 价格
	Percent      float64   `gorm:"column:percent;type:double(16,2);default:null" json:"percent"`            // 百分比
	TurnoverRate float64   `gorm:"column:turnover_rate;type:double(16,2);default:null" json:"turnoverRate"` // 换手率
	Score        int       `gorm:"column:score;type:int(11);default:null" json:"score"`                     // 得分
	Doc          string    `gorm:"column:doc;type:varchar(1024);default:null" json:"doc"`                   // 描述
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime;default:null" json:"createdAt"`           // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *AnalyCdTbl) TableName() string {
	return "analy_cd_tbl"
}

// AnalyCdTblColumns get sql column name.获取数据库列名
var AnalyCdTblColumns = struct {
	ID           string
	Code         string
	Name         string
	Day0         string
	DayStr       string
	Price        string
	NewPrice     string
	Percent      string
	TurnoverRate string
	Score        string
	Doc          string
	CreatedAt    string
}{
	ID:           "id",
	Code:         "code",
	Name:         "name",
	Day0:         "day0",
	DayStr:       "day_str",
	Price:        "price",
	NewPrice:     "new_price",
	Percent:      "percent",
	TurnoverRate: "turnover_rate",
	Score:        "score",
	Doc:          "doc",
	CreatedAt:    "created_at",
}

// AnalyCdView VIEW
type AnalyCdView struct {
}

// TableName get sql table name.获取数据库表名
func (m *AnalyCdView) TableName() string {
	return "analy_cd_view"
}

// AnalyCdViewColumns get sql column name.获取数据库列名
var AnalyCdViewColumns = struct {
}{}

// AnalyDbszxTbl 放量赛选器
type AnalyDbszxTbl struct {
	ID           int            `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Code         string         `gorm:"unique;column:code;type:varchar(255);default:null" json:"code"`           // 股票代码
	Name         string         `gorm:"column:name;type:varchar(255);default:null" json:"name"`                  // 股票名字
	Day          datatypes.Date `gorm:"column:day;type:date;default:null" json:"day"`                            // 入池时间
	DayStr       string         `gorm:"column:day_str;type:varchar(255);default:null" json:"dayStr"`             // 入池时间
	Price        float64        `gorm:"column:price;type:double(11,2);default:null" json:"price"`                // 价格
	NewPrice     float64        `gorm:"column:new_price;type:double(11,2);default:null" json:"newPrice"`         // 价格
	Percent      float64        `gorm:"column:percent;type:double(16,2);default:null" json:"percent"`            // 百分比
	TurnoverRate float64        `gorm:"column:turnover_rate;type:double(16,2);default:null" json:"turnoverRate"` // 换手率
	CreatedAt    time.Time      `gorm:"column:created_at;type:datetime;default:null" json:"createdAt"`           // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *AnalyDbszxTbl) TableName() string {
	return "analy_dbszx_tbl"
}

// AnalyDbszxTblColumns get sql column name.获取数据库列名
var AnalyDbszxTblColumns = struct {
	ID           string
	Code         string
	Name         string
	Day          string
	DayStr       string
	Price        string
	NewPrice     string
	Percent      string
	TurnoverRate string
	CreatedAt    string
}{
	ID:           "id",
	Code:         "code",
	Name:         "name",
	Day:          "day",
	DayStr:       "day_str",
	Price:        "price",
	NewPrice:     "new_price",
	Percent:      "percent",
	TurnoverRate: "turnover_rate",
	CreatedAt:    "created_at",
}

// AnalyDwflTbl 放量赛选器
type AnalyDwflTbl struct {
	ID           int            `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Code         string         `gorm:"unique;column:code;type:varchar(255);default:null" json:"code"`           // 股票代码
	Name         string         `gorm:"column:name;type:varchar(255);default:null" json:"name"`                  // 股票名字
	Day          datatypes.Date `gorm:"column:day;type:date;default:null" json:"day"`                            // 入池时间
	DayStr       string         `gorm:"column:day_str;type:varchar(255);default:null" json:"dayStr"`             // 入池时间
	Price        float64        `gorm:"column:price;type:double(11,2);default:null" json:"price"`                // 价格
	NewPrice     float64        `gorm:"column:new_price;type:double(11,2);default:null" json:"newPrice"`         // 价格
	Percent      float64        `gorm:"column:percent;type:double(16,2);default:null" json:"percent"`            // 百分比
	TurnoverRate float64        `gorm:"column:turnover_rate;type:double(16,2);default:null" json:"turnoverRate"` // 换手率
	CreatedAt    time.Time      `gorm:"column:created_at;type:datetime;default:null" json:"createdAt"`           // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *AnalyDwflTbl) TableName() string {
	return "analy_dwfl_tbl"
}

// AnalyDwflTblColumns get sql column name.获取数据库列名
var AnalyDwflTblColumns = struct {
	ID           string
	Code         string
	Name         string
	Day          string
	DayStr       string
	Price        string
	NewPrice     string
	Percent      string
	TurnoverRate string
	CreatedAt    string
}{
	ID:           "id",
	Code:         "code",
	Name:         "name",
	Day:          "day",
	DayStr:       "day_str",
	Price:        "price",
	NewPrice:     "new_price",
	Percent:      "percent",
	TurnoverRate: "turnover_rate",
	CreatedAt:    "created_at",
}

// AnalyFlTbl 放量赛选器
type AnalyFlTbl struct {
	ID           int            `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Code         string         `gorm:"unique;column:code;type:varchar(255);default:null" json:"code"`           // 股票代码
	Name         string         `gorm:"column:name;type:varchar(255);default:null" json:"name"`                  // 股票名字
	Day          datatypes.Date `gorm:"column:day;type:date;default:null" json:"day"`                            // 入池时间
	DayStr       string         `gorm:"column:day_str;type:varchar(255);default:null" json:"dayStr"`             // 入池时间
	Price        float64        `gorm:"column:price;type:double(11,2);default:null" json:"price"`                // 价格
	NewPrice     float64        `gorm:"column:new_price;type:double(11,2);default:null" json:"newPrice"`         // 价格
	Percent      float64        `gorm:"column:percent;type:double(16,2);default:null" json:"percent"`            // 百分比
	TurnoverRate float64        `gorm:"column:turnover_rate;type:double(16,2);default:null" json:"turnoverRate"` // 换手率
	Score        int            `gorm:"column:score;type:int(11);default:null" json:"score"`                     // 得分
	CreatedAt    time.Time      `gorm:"column:created_at;type:datetime;default:null" json:"createdAt"`           // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *AnalyFlTbl) TableName() string {
	return "analy_fl_tbl"
}

// AnalyFlTblColumns get sql column name.获取数据库列名
var AnalyFlTblColumns = struct {
	ID           string
	Code         string
	Name         string
	Day          string
	DayStr       string
	Price        string
	NewPrice     string
	Percent      string
	TurnoverRate string
	Score        string
	CreatedAt    string
}{
	ID:           "id",
	Code:         "code",
	Name:         "name",
	Day:          "day",
	DayStr:       "day_str",
	Price:        "price",
	NewPrice:     "new_price",
	Percent:      "percent",
	TurnoverRate: "turnover_rate",
	Score:        "score",
	CreatedAt:    "created_at",
}

// AnalyFlView VIEW
type AnalyFlView struct {
	ID           int            `gorm:"column:id;type:int(11);not null;default:0" json:"id"`
	Code         string         `gorm:"column:code;type:varchar(255);default:null" json:"code"`                  // 股票代码
	Name         string         `gorm:"column:name;type:varchar(255);default:null" json:"name"`                  // 股票名字
	UpName       string         `gorm:"column:up_name;type:varchar(255);default:null" json:"upName"`             // 行业名
	Day          datatypes.Date `gorm:"column:day;type:date;default:null" json:"day"`                            // 入池时间
	DayStr       string         `gorm:"column:day_str;type:varchar(255);default:null" json:"dayStr"`             // 入池时间
	Price        float64        `gorm:"column:price;type:double(11,2);default:null" json:"price"`                // 价格
	NewPrice     float64        `gorm:"column:new_price;type:double(11,2);default:null" json:"newPrice"`         // 价格
	Percent      float64        `gorm:"column:percent;type:double(16,2);default:null" json:"percent"`            // 百分比
	TurnoverRate float64        `gorm:"column:turnover_rate;type:double(16,2);default:null" json:"turnoverRate"` // 换手率
	Score        int            `gorm:"column:score;type:int(11);default:null" json:"score"`                     // 得分
	CreatedAt    time.Time      `gorm:"column:created_at;type:datetime;default:null" json:"createdAt"`           // 创建时间
	Num          int            `gorm:"column:num;type:int(11);default:null" json:"num"`                         // 总家数
	Up           int            `gorm:"column:up;type:int(11);default:null" json:"up"`                           // 上涨家数
	HyName       string         `gorm:"column:hy_name;type:text;default:null" json:"hyName"`                     // 行业名
}

// TableName get sql table name.获取数据库表名
func (m *AnalyFlView) TableName() string {
	return "analy_fl_view"
}

// AnalyFlViewColumns get sql column name.获取数据库列名
var AnalyFlViewColumns = struct {
	ID           string
	Code         string
	Name         string
	UpName       string
	Day          string
	DayStr       string
	Price        string
	NewPrice     string
	Percent      string
	TurnoverRate string
	Score        string
	CreatedAt    string
	Num          string
	Up           string
	HyName       string
}{
	ID:           "id",
	Code:         "code",
	Name:         "name",
	UpName:       "up_name",
	Day:          "day",
	DayStr:       "day_str",
	Price:        "price",
	NewPrice:     "new_price",
	Percent:      "percent",
	TurnoverRate: "turnover_rate",
	Score:        "score",
	CreatedAt:    "created_at",
	Num:          "num",
	Up:           "up",
	HyName:       "hy_name",
}

// AnalyFlViewOld VIEW
type AnalyFlViewOld struct {
	ID           int            `gorm:"column:id;type:int(11);not null;default:0" json:"id"`
	Code         string         `gorm:"column:code;type:varchar(255);default:null" json:"code"`                  // 股票代码
	Name         string         `gorm:"column:name;type:varchar(255);default:null" json:"name"`                  // 股票名字
	UpName       string         `gorm:"column:up_name;type:varchar(255);default:null" json:"upName"`             // 行业名
	Day          datatypes.Date `gorm:"column:day;type:date;default:null" json:"day"`                            // 入池时间
	DayStr       string         `gorm:"column:day_str;type:varchar(255);default:null" json:"dayStr"`             // 入池时间
	Price        float64        `gorm:"column:price;type:double(11,2);default:null" json:"price"`                // 价格
	NewPrice     float64        `gorm:"column:new_price;type:double(11,2);default:null" json:"newPrice"`         // 价格
	Percent      float64        `gorm:"column:percent;type:double(16,2);default:null" json:"percent"`            // 百分比
	TurnoverRate float64        `gorm:"column:turnover_rate;type:double(16,2);default:null" json:"turnoverRate"` // 换手率
	Score        int            `gorm:"column:score;type:int(11);default:null" json:"score"`                     // 得分
	CreatedAt    time.Time      `gorm:"column:created_at;type:datetime;default:null" json:"createdAt"`           // 创建时间
	Num          int            `gorm:"column:num;type:int(11);default:null" json:"num"`                         // 总家数
	Up           int            `gorm:"column:up;type:int(11);default:null" json:"up"`                           // 上涨家数
	HyName       string         `gorm:"column:hy_name;type:text;default:null" json:"hyName"`                     // 行业名
}

// TableName get sql table name.获取数据库表名
func (m *AnalyFlViewOld) TableName() string {
	return "analy_fl_view_old"
}

// AnalyFlViewOldColumns get sql column name.获取数据库列名
var AnalyFlViewOldColumns = struct {
	ID           string
	Code         string
	Name         string
	UpName       string
	Day          string
	DayStr       string
	Price        string
	NewPrice     string
	Percent      string
	TurnoverRate string
	Score        string
	CreatedAt    string
	Num          string
	Up           string
	HyName       string
}{
	ID:           "id",
	Code:         "code",
	Name:         "name",
	UpName:       "up_name",
	Day:          "day",
	DayStr:       "day_str",
	Price:        "price",
	NewPrice:     "new_price",
	Percent:      "percent",
	TurnoverRate: "turnover_rate",
	Score:        "score",
	CreatedAt:    "created_at",
	Num:          "num",
	Up:           "up",
	HyName:       "hy_name",
}

// AnalyHpTbl 放量赛选器
type AnalyHpTbl struct {
	ID           int            `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Code         string         `gorm:"unique;column:code;type:varchar(255);default:null" json:"code"`           // 股票代码
	Name         string         `gorm:"column:name;type:varchar(255);default:null" json:"name"`                  // 股票名字
	Day          datatypes.Date `gorm:"column:day;type:date;default:null" json:"day"`                            // 入池时间
	DayStr       string         `gorm:"column:day_str;type:varchar(255);default:null" json:"dayStr"`             // 入池时间
	Price        float64        `gorm:"column:price;type:double(11,2);default:null" json:"price"`                // 价格
	NewPrice     float64        `gorm:"column:new_price;type:double(11,2);default:null" json:"newPrice"`         // 价格
	Percent      float64        `gorm:"column:percent;type:double(16,2);default:null" json:"percent"`            // 百分比
	TurnoverRate float64        `gorm:"column:turnover_rate;type:double(16,2);default:null" json:"turnoverRate"` // 换手率
	Score        float64        `gorm:"column:score;type:double(16,4);default:null" json:"score"`                // 得分
	CreatedAt    time.Time      `gorm:"column:created_at;type:datetime;default:null" json:"createdAt"`           // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *AnalyHpTbl) TableName() string {
	return "analy_hp_tbl"
}

// AnalyHpTblColumns get sql column name.获取数据库列名
var AnalyHpTblColumns = struct {
	ID           string
	Code         string
	Name         string
	Day          string
	DayStr       string
	Price        string
	NewPrice     string
	Percent      string
	TurnoverRate string
	Score        string
	CreatedAt    string
}{
	ID:           "id",
	Code:         "code",
	Name:         "name",
	Day:          "day",
	DayStr:       "day_str",
	Price:        "price",
	NewPrice:     "new_price",
	Percent:      "percent",
	TurnoverRate: "turnover_rate",
	Score:        "score",
	CreatedAt:    "created_at",
}

// AnalyHyTbl 行业分类
type AnalyHyTbl struct {
	ID           int            `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Code         string         `gorm:"uniqueIndex:code;column:code;type:varchar(255);default:null" json:"code"` // 行业代码
	Name         string         `gorm:"column:name;type:varchar(255);default:null" json:"name"`                  // 行业名
	Percent      float64        `gorm:"column:percent;type:double(16,2);default:null" json:"percent"`            // 百分比
	TurnoverRate float64        `gorm:"column:turnover_rate;type:double(16,2);default:null" json:"turnoverRate"` // 换手率
	Num          int            `gorm:"column:num;type:int(11);default:null" json:"num"`                         // 总家数
	Up           int            `gorm:"column:up;type:int(11);default:null" json:"up"`                           // 上涨家数
	CreatedAt    time.Time      `gorm:"column:created_at;type:datetime;default:null" json:"createdAt"`           // 创建时间
	Day          datatypes.Date `gorm:"uniqueIndex:code;column:day;type:date;default:null" json:"day"`           // 入池时间
	DayStr       string         `gorm:"column:day_str;type:varchar(255);default:null" json:"dayStr"`             // 入池时间
}

// TableName get sql table name.获取数据库表名
func (m *AnalyHyTbl) TableName() string {
	return "analy_hy_tbl"
}

// AnalyHyTblColumns get sql column name.获取数据库列名
var AnalyHyTblColumns = struct {
	ID           string
	Code         string
	Name         string
	Percent      string
	TurnoverRate string
	Num          string
	Up           string
	CreatedAt    string
	Day          string
	DayStr       string
}{
	ID:           "id",
	Code:         "code",
	Name:         "name",
	Percent:      "percent",
	TurnoverRate: "turnover_rate",
	Num:          "num",
	Up:           "up",
	CreatedAt:    "created_at",
	Day:          "day",
	DayStr:       "day_str",
}

// AnalyHyView VIEW
type AnalyHyView struct {
	HyCode  string  `gorm:"column:hy_code;type:varchar(255);default:null" json:"hyCode"` // 股票代码
	HyName  string  `gorm:"column:hy_name;type:varchar(255);default:null" json:"hyName"` // 行业名
	Price1  float64 `gorm:"column:price1;type:double(21,4);default:null" json:"price1"`
	Price5  float64 `gorm:"column:price5;type:double(21,4);default:null" json:"price5"`
	Price10 float64 `gorm:"column:price10;type:double(21,4);default:null" json:"price10"`
	Price20 float64 `gorm:"column:price20;type:double(21,4);default:null" json:"price20"`
	Day0    int64   `gorm:"column:day0;type:bigint(11);default:null" json:"day0"`          // 当日0点时间戳
	Day0Str string  `gorm:"column:day0_str;type:varchar(255);default:null" json:"day0Str"` // 当日0点时间戳
}

// TableName get sql table name.获取数据库表名
func (m *AnalyHyView) TableName() string {
	return "analy_hy_view"
}

// AnalyHyViewColumns get sql column name.获取数据库列名
var AnalyHyViewColumns = struct {
	HyCode  string
	HyName  string
	Price1  string
	Price5  string
	Price10 string
	Price20 string
	Day0    string
	Day0Str string
}{
	HyCode:  "hy_code",
	HyName:  "hy_name",
	Price1:  "price1",
	Price5:  "price5",
	Price10: "price10",
	Price20: "price20",
	Day0:    "day0",
	Day0Str: "day0_str",
}

// AnalyLhbTbl 主力净流入列表
type AnalyLhbTbl struct {
	ID           int       `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Code         string    `gorm:"uniqueIndex:code;column:code;type:varchar(255);default:null" json:"code"` // 股票代码
	Name         string    `gorm:"column:name;type:varchar(255);default:null" json:"name"`                  // 股票名字
	Day0         int64     `gorm:"uniqueIndex:code;column:day0;type:bigint(11);default:null" json:"day0"`   // 当日0点时间戳
	DayStr       string    `gorm:"column:day_str;type:varchar(255);default:null" json:"dayStr"`             // 入池时间
	Jlr          float64   `gorm:"column:jlr;type:double(11,2);default:null" json:"jlr"`                    // 价格
	JgJlr        float64   `gorm:"column:jg_jlr;type:double(16,4);default:null" json:"jgJlr"`               // 机构净流入(万元)
	Price        float64   `gorm:"column:price;type:double(16,2);default:null" json:"price"`                // 龙虎榜净流入(万元)
	NewPrice     float64   `gorm:"column:new_price;type:double(11,2);default:null" json:"newPrice"`         // 价格
	TurnoverRate float64   `gorm:"column:turnover_rate;type:double(16,2);default:null" json:"turnoverRate"` // 换手率
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime;default:null" json:"createdAt"`           // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *AnalyLhbTbl) TableName() string {
	return "analy_lhb_tbl"
}

// AnalyLhbTblColumns get sql column name.获取数据库列名
var AnalyLhbTblColumns = struct {
	ID           string
	Code         string
	Name         string
	Day0         string
	DayStr       string
	Jlr          string
	JgJlr        string
	Price        string
	NewPrice     string
	TurnoverRate string
	CreatedAt    string
}{
	ID:           "id",
	Code:         "code",
	Name:         "name",
	Day0:         "day0",
	DayStr:       "day_str",
	Jlr:          "jlr",
	JgJlr:        "jg_jlr",
	Price:        "price",
	NewPrice:     "new_price",
	TurnoverRate: "turnover_rate",
	CreatedAt:    "created_at",
}

// AnalyTdxAfterView VIEW
type AnalyTdxAfterView struct {
}

// TableName get sql table name.获取数据库表名
func (m *AnalyTdxAfterView) TableName() string {
	return "analy_tdx_after_view"
}

// AnalyTdxAfterViewColumns get sql column name.获取数据库列名
var AnalyTdxAfterViewColumns = struct {
}{}

// AnalyTdxView VIEW
type AnalyTdxView struct {
}

// TableName get sql table name.获取数据库表名
func (m *AnalyTdxView) TableName() string {
	return "analy_tdx_view"
}

// AnalyTdxViewColumns get sql column name.获取数据库列名
var AnalyTdxViewColumns = struct {
}{}

// AnalyTkTbl 跳空选择器
type AnalyTkTbl struct {
	ID        int       `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Code      string    `gorm:"uniqueIndex:code;column:code;type:varchar(255);default:null" json:"code"` // 股票代码
	Name      string    `gorm:"column:name;type:varchar(255);default:null" json:"name"`                  // 股票名字
	Day0      int64     `gorm:"uniqueIndex:code;column:day0;type:bigint(11);default:null" json:"day0"`   // 当日0点时间戳
	Day0Str   string    `gorm:"column:day0_str;type:varchar(255);default:null" json:"day0Str"`           // 入池时间
	Price     float64   `gorm:"column:price;type:double(11,2);default:null" json:"price"`                // 价格
	Min       float64   `gorm:"column:min;type:double(16,2);default:null" json:"min"`                    // 跳空最低价
	Max       float64   `gorm:"column:max;type:double(16,2);default:null" json:"max"`                    // 当日最高值
	Flag      int       `gorm:"column:flag;type:int(255);default:null" json:"flag"`                      // 标记：0 默认 ，1：回补了
	Percent   float64   `gorm:"column:percent;type:double(16,2);default:null" json:"percent"`            // 百分比
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;default:null" json:"createdAt"`           // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *AnalyTkTbl) TableName() string {
	return "analy_tk_tbl"
}

// AnalyTkTblColumns get sql column name.获取数据库列名
var AnalyTkTblColumns = struct {
	ID        string
	Code      string
	Name      string
	Day0      string
	Day0Str   string
	Price     string
	Min       string
	Max       string
	Flag      string
	Percent   string
	CreatedAt string
}{
	ID:        "id",
	Code:      "code",
	Name:      "name",
	Day0:      "day0",
	Day0Str:   "day0_str",
	Price:     "price",
	Min:       "min",
	Max:       "max",
	Flag:      "flag",
	Percent:   "percent",
	CreatedAt: "created_at",
}

// AnalyUpTbl 日线多头赛选器
type AnalyUpTbl struct {
	ID           int       `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Code         string    `gorm:"unique;column:code;type:varchar(255);default:null" json:"code"`           // 股票代码
	Name         string    `gorm:"column:name;type:varchar(255);default:null" json:"name"`                  // 股票名字
	Day0         int64     `gorm:"column:day0;type:bigint(11);default:null" json:"day0"`                    // 入池时间
	DayStr       string    `gorm:"column:day_str;type:varchar(255);default:null" json:"dayStr"`             // 入池时间
	Price        float64   `gorm:"column:price;type:double(11,2);default:null" json:"price"`                // 价格
	NewPrice     float64   `gorm:"column:new_price;type:double(11,2);default:null" json:"newPrice"`         // 价格
	Percent      float64   `gorm:"column:percent;type:double(16,2);default:null" json:"percent"`            // 百分比
	TurnoverRate float64   `gorm:"column:turnover_rate;type:double(16,2);default:null" json:"turnoverRate"` // 换手率
	Doc          string    `gorm:"column:doc;type:varchar(1024);default:null" json:"doc"`                   // 描述
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime;default:null" json:"createdAt"`           // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *AnalyUpTbl) TableName() string {
	return "analy_up_tbl"
}

// AnalyUpTblColumns get sql column name.获取数据库列名
var AnalyUpTblColumns = struct {
	ID           string
	Code         string
	Name         string
	Day0         string
	DayStr       string
	Price        string
	NewPrice     string
	Percent      string
	TurnoverRate string
	Doc          string
	CreatedAt    string
}{
	ID:           "id",
	Code:         "code",
	Name:         "name",
	Day0:         "day0",
	DayStr:       "day_str",
	Price:        "price",
	NewPrice:     "new_price",
	Percent:      "percent",
	TurnoverRate: "turnover_rate",
	Doc:          "doc",
	CreatedAt:    "created_at",
}

// AnalyUpView VIEW
type AnalyUpView struct {
}

// TableName get sql table name.获取数据库表名
func (m *AnalyUpView) TableName() string {
	return "analy_up_view"
}

// AnalyUpViewColumns get sql column name.获取数据库列名
var AnalyUpViewColumns = struct {
}{}

// APITbl 微信用户信息
type APITbl struct {
	ID          int       `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	APIID       string    `gorm:"column:api_id;type:varchar(255);default:null" json:"apiId"`             // 授权应用id
	APIKey      string    `gorm:"column:api_key;type:varchar(255);default:null" json:"apiKey"`           // 授权应用key
	APISecret   string    `gorm:"column:api_secret;type:varchar(255);default:null" json:"apiSecret"`     // 秘钥/token
	Tag         string    `gorm:"column:tag;type:varchar(255);default:null" json:"tag"`                  // 标记
	AccessToken string    `gorm:"column:access_token;type:varchar(255);default:null" json:"accessToken"` // 访问令牌
	Expires     time.Time `gorm:"column:expires;type:datetime;default:null" json:"expires"`              // 过期时间
	Ignore      string    `gorm:"column:ignore;type:varchar(255);default:null" json:"ignore"`            // 忽略的词性
	IgsTag      string    `gorm:"column:igs_tag;type:text;default:null" json:"igsTag"`                   // 忽略关键词
}

// TableName get sql table name.获取数据库表名
func (m *APITbl) TableName() string {
	return "api_tbl"
}

// APITblColumns get sql column name.获取数据库列名
var APITblColumns = struct {
	ID          string
	APIID       string
	APIKey      string
	APISecret   string
	Tag         string
	AccessToken string
	Expires     string
	Ignore      string
	IgsTag      string
}{
	ID:          "id",
	APIID:       "api_id",
	APIKey:      "api_key",
	APISecret:   "api_secret",
	Tag:         "tag",
	AccessToken: "access_token",
	Expires:     "expires",
	Ignore:      "ignore",
	IgsTag:      "igs_tag",
}

// BillRefundTbl 用户退款信息
type BillRefundTbl struct {
	ID        int       `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	UserID    string    `gorm:"column:user_id;type:varchar(255);not null;default:0" json:"userId"`            // 用户id
	BillID    string    `gorm:"index:bill_id;column:bill_id;type:varchar(32);not null" json:"billId"`         // 账单id
	RefundID  string    `gorm:"column:refund_id;type:varchar(32);not null;default:0" json:"refundId"`         // 退款账单号
	RefundFee int64     `gorm:"column:refund_fee;type:bigint(20);not null;default:0" json:"refundFee"`        // 退款金额
	Desc      string    `gorm:"column:desc;type:varchar(1024);default:null" json:"desc"`                      // 订单备注
	CreatedBy string    `gorm:"column:created_by;type:varchar(255);default:null;default:''" json:"createdBy"` // 创建者
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;default:null" json:"createdAt"`                // 创建时间
	UpdatedBy string    `gorm:"column:updated_by;type:varchar(255);default:null;default:''" json:"updatedBy"` // 更新者
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;default:null" json:"updatedAt"`                // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *BillRefundTbl) TableName() string {
	return "bill_refund_tbl"
}

// BillRefundTblColumns get sql column name.获取数据库列名
var BillRefundTblColumns = struct {
	ID        string
	UserID    string
	BillID    string
	RefundID  string
	RefundFee string
	Desc      string
	CreatedBy string
	CreatedAt string
	UpdatedBy string
	UpdatedAt string
}{
	ID:        "id",
	UserID:    "user_id",
	BillID:    "bill_id",
	RefundID:  "refund_id",
	RefundFee: "refund_fee",
	Desc:      "desc",
	CreatedBy: "created_by",
	CreatedAt: "created_at",
	UpdatedBy: "updated_by",
	UpdatedAt: "updated_at",
}

// BillTbl 用户账单信息
type BillTbl struct {
	ID          int       `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	UserID      string    `gorm:"column:user_id;type:varchar(255);not null;default:0" json:"userId"`            // 用户id
	BillID      string    `gorm:"unique;column:bill_id;type:varchar(32);not null" json:"billId"`                // 账单id
	PayPlatform int       `gorm:"column:pay_platform;type:int(11);not null;default:0" json:"payPlatform"`       // 支付类型(1:微信支付)
	PayAmount   int64     `gorm:"column:pay_amount;type:bigint(20);not null;default:0" json:"payAmount"`        // 支付金额
	Status      int       `gorm:"column:status;type:int(11);not null;default:0" json:"status"`                  // 支付状态(-1:已退款,1:待支付,2:已支付,待发货,3:已取消,4:待收货,5:已完成，6:售后待评价)
	Desc        string    `gorm:"column:desc;type:varchar(1024);default:null" json:"desc"`                      // 订单备注
	CreatedBy   string    `gorm:"column:created_by;type:varchar(255);default:null;default:''" json:"createdBy"` // 创建者
	CreatedAt   time.Time `gorm:"column:created_at;type:datetime;default:null" json:"createdAt"`                // 创建时间
	UpdatedBy   string    `gorm:"column:updated_by;type:varchar(255);default:null;default:''" json:"updatedBy"` // 更新者
	UpdatedAt   time.Time `gorm:"column:updated_at;type:datetime;default:null" json:"updatedAt"`                // 更新时间
	DeletedBy   string    `gorm:"column:deleted_by;type:varchar(255);default:null;default:''" json:"deletedBy"` // 删除者
	DeletedAt   time.Time `gorm:"column:deleted_at;type:datetime;default:null" json:"deletedAt"`                // 删除时间
}

// TableName get sql table name.获取数据库表名
func (m *BillTbl) TableName() string {
	return "bill_tbl"
}

// BillTblColumns get sql column name.获取数据库列名
var BillTblColumns = struct {
	ID          string
	UserID      string
	BillID      string
	PayPlatform string
	PayAmount   string
	Status      string
	Desc        string
	CreatedBy   string
	CreatedAt   string
	UpdatedBy   string
	UpdatedAt   string
	DeletedBy   string
	DeletedAt   string
}{
	ID:          "id",
	UserID:      "user_id",
	BillID:      "bill_id",
	PayPlatform: "pay_platform",
	PayAmount:   "pay_amount",
	Status:      "status",
	Desc:        "desc",
	CreatedBy:   "created_by",
	CreatedAt:   "created_at",
	UpdatedBy:   "updated_by",
	UpdatedAt:   "updated_at",
	DeletedBy:   "deleted_by",
	DeletedAt:   "deleted_at",
}

// BsRapidlyTbl 股票急速上涨下跌提醒
type BsRapidlyTbl struct {
	ID        int       `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Day0      int64     `gorm:"unique;column:day0;type:bigint(11);default:null" json:"day0"`   // 当日0点时间戳
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;default:null" json:"createdAt"` // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *BsRapidlyTbl) TableName() string {
	return "bs_rapidly_tbl"
}

// BsRapidlyTblColumns get sql column name.获取数据库列名
var BsRapidlyTblColumns = struct {
	ID        string
	Day0      string
	CreatedAt string
}{
	ID:        "id",
	Day0:      "day0",
	CreatedAt: "created_at",
}

// CdView VIEW
type CdView struct {
	ID           int       `gorm:"column:id;type:int(11);not null;default:0" json:"id"`
	Code         string    `gorm:"column:code;type:varchar(255);default:null" json:"code"`                  // 股票代码
	Name         string    `gorm:"column:name;type:varchar(255);default:null" json:"name"`                  // 股票名字
	Day0         int64     `gorm:"column:day0;type:bigint(20);default:null" json:"day0"`                    // 入池时间
	DayStr       string    `gorm:"column:day_str;type:varchar(255);default:null" json:"dayStr"`             // 入池时间
	Price        float64   `gorm:"column:price;type:double(11,2);default:null" json:"price"`                // 价格
	NewPrice     float64   `gorm:"column:new_price;type:double(11,2);default:null" json:"newPrice"`         // 价格
	Percent      float64   `gorm:"column:percent;type:double(16,2);default:null" json:"percent"`            // 百分比
	TurnoverRate float64   `gorm:"column:turnover_rate;type:double(16,2);default:null" json:"turnoverRate"` // 换手率
	Score        int       `gorm:"column:score;type:int(11);default:null" json:"score"`                     // 得分
	Doc          string    `gorm:"column:doc;type:varchar(1024);default:null" json:"doc"`                   // 描述
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime;default:null" json:"createdAt"`           // 创建时间
	HyName       string    `gorm:"column:hy_name;type:text;default:null" json:"hyName"`                     // 行业名
}

// TableName get sql table name.获取数据库表名
func (m *CdView) TableName() string {
	return "cd_view"
}

// CdViewColumns get sql column name.获取数据库列名
var CdViewColumns = struct {
	ID           string
	Code         string
	Name         string
	Day0         string
	DayStr       string
	Price        string
	NewPrice     string
	Percent      string
	TurnoverRate string
	Score        string
	Doc          string
	CreatedAt    string
	HyName       string
}{
	ID:           "id",
	Code:         "code",
	Name:         "name",
	Day0:         "day0",
	DayStr:       "day_str",
	Price:        "price",
	NewPrice:     "new_price",
	Percent:      "percent",
	TurnoverRate: "turnover_rate",
	Score:        "score",
	Doc:          "doc",
	CreatedAt:    "created_at",
	HyName:       "hy_name",
}

// GroupListTbl 分组信息
type GroupListTbl struct {
	ID        int    `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Key       string `gorm:"unique;column:key;type:varchar(255);default:null" json:"key"`              // 搜索代码
	GroupName string `gorm:"unique;column:group_name;type:varchar(255);default:null" json:"groupName"` // 分组名
}

// TableName get sql table name.获取数据库表名
func (m *GroupListTbl) TableName() string {
	return "group_list_tbl"
}

// GroupListTblColumns get sql column name.获取数据库列名
var GroupListTblColumns = struct {
	ID        string
	Key       string
	GroupName string
}{
	ID:        "id",
	Key:       "key",
	GroupName: "group_name",
}

// GroupTbl 推荐组信息
type GroupTbl struct {
	ID        int       `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	GroupName string    `gorm:"uniqueIndex:group_name;column:group_name;type:varchar(255);default:null" json:"groupName"` // 分组名
	Code      string    `gorm:"uniqueIndex:group_name;column:code;type:varchar(255);default:null" json:"code"`            // 股票代码
	Wi        int       `gorm:"column:wi;type:int(11);default:null" json:"wi"`                                            // 推荐权重
	UserName  string    `gorm:"column:user_name;type:varchar(255);default:null" json:"userName"`                          // 推荐人
	CreatedBy string    `gorm:"column:created_by;type:varchar(255);default:null;default:''" json:"createdBy"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;default:null" json:"createdAt"`
	UpdateAt  time.Time `gorm:"column:update_at;type:datetime;default:null" json:"updateAt"`
	UpdateBy  string    `gorm:"column:update_by;type:varchar(255);default:null;default:''" json:"updateBy"`
}

// TableName get sql table name.获取数据库表名
func (m *GroupTbl) TableName() string {
	return "group_tbl"
}

// GroupTblColumns get sql column name.获取数据库列名
var GroupTblColumns = struct {
	ID        string
	GroupName string
	Code      string
	Wi        string
	UserName  string
	CreatedBy string
	CreatedAt string
	UpdateAt  string
	UpdateBy  string
}{
	ID:        "id",
	GroupName: "group_name",
	Code:      "code",
	Wi:        "wi",
	UserName:  "user_name",
	CreatedBy: "created_by",
	CreatedAt: "created_at",
	UpdateAt:  "update_at",
	UpdateBy:  "update_by",
}

// GroupWatchTbl 推荐组信息
type GroupWatchTbl struct {
	ID        int       `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Code      string    `gorm:"unique;column:code;type:varchar(255);default:null" json:"code"`   // 股票代码
	Name      string    `gorm:"column:name;type:varchar(255);default:null" json:"name"`          // 股票名字
	Price     float64   `gorm:"column:price;type:double(11,2);default:null" json:"price"`        // 价格
	NewPrice  float64   `gorm:"column:new_price;type:double(11,2);default:null" json:"newPrice"` // 最新价格
	Desc      string    `gorm:"column:desc;type:varchar(1024);default:null" json:"desc"`         // 描述
	UserName  string    `gorm:"column:user_name;type:varchar(255);default:null" json:"userName"` // 推荐人
	CreatedBy string    `gorm:"column:created_by;type:varchar(255);default:null;default:''" json:"createdBy"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;default:null" json:"createdAt"` // 入池时间
}

// TableName get sql table name.获取数据库表名
func (m *GroupWatchTbl) TableName() string {
	return "group_watch_tbl"
}

// GroupWatchTblColumns get sql column name.获取数据库列名
var GroupWatchTblColumns = struct {
	ID        string
	Code      string
	Name      string
	Price     string
	NewPrice  string
	Desc      string
	UserName  string
	CreatedBy string
	CreatedAt string
}{
	ID:        "id",
	Code:      "code",
	Name:      "name",
	Price:     "price",
	NewPrice:  "new_price",
	Desc:      "desc",
	UserName:  "user_name",
	CreatedBy: "created_by",
	CreatedAt: "created_at",
}

// HyDailyTbl 日数据，每日的数据
type HyDailyTbl struct {
	ID           int       `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	HyCode       string    `gorm:"uniqueIndex:hy_code;column:hy_code;type:varchar(255);default:null" json:"hyCode"`     // 股票代码
	HyName       string    `gorm:"column:hy_name;type:varchar(255);default:null" json:"hyName"`                         // 行业名
	Price        float64   `gorm:"column:price;type:double(16,2);default:null" json:"price"`                            // 当前价格
	Percent      float64   `gorm:"column:percent;type:double(16,2);default:null" json:"percent"`                        // 百分比
	Ma5          float64   `gorm:"column:ma5;type:double(16,2);default:null" json:"ma5"`                                // 5日均线
	Ma10         float64   `gorm:"column:ma10;type:double(16,2);default:null" json:"ma10"`                              // 10日均线
	Ma20         float64   `gorm:"column:ma20;type:double(16,2);default:null" json:"ma20"`                              // 20日均线
	Ma60         float64   `gorm:"column:ma60;type:double(16,2);default:null" json:"ma60"`                              // 60日均线
	Day0         int64     `gorm:"uniqueIndex:hy_code;index:day0;column:day0;type:bigint(11);default:null" json:"day0"` // 当日0点时间戳
	Day0Str      string    `gorm:"column:day0_str;type:varchar(255);default:null" json:"day0Str"`                       // 当日0点时间戳
	Volume       float64   `gorm:"column:volume;type:double(16,2);default:null" json:"volume"`                          // 成交量（手）
	Turnover     float64   `gorm:"column:turnover;type:double(16,2);default:null" json:"turnover"`                      // 成交额（万）
	TurnoverRate float64   `gorm:"column:turnover_rate;type:double(16,2);default:null" json:"turnoverRate"`             // 换手率
	Zljlr        float64   `gorm:"column:zljlr;type:double(16,4);default:null" json:"zljlr"`                            // 主力净流入
	OpenPrice    float64   `gorm:"column:open_price;type:double(16,2);default:null" json:"openPrice"`                   // 开盘价
	ClosePrice   float64   `gorm:"column:close_price;type:double(16,2);default:null" json:"closePrice"`                 // 收盘价
	Max          float64   `gorm:"column:max;type:double(16,2);default:null" json:"max"`                                // 当日最高值
	Min          float64   `gorm:"column:min;type:double(16,2);default:null" json:"min"`                                // 当日最低价
	Pe           float64   `gorm:"column:pe;type:double(16,2);default:null" json:"pe"`                                  // 市盈率
	Peg          float64   `gorm:"column:peg;type:double(16,2);default:null" json:"peg"`                                // PEG值
	Pb           float64   `gorm:"column:pb;type:double(16,2);default:null" json:"pb"`                                  // 市净率
	CapVag       float64   `gorm:"column:cap_vag;type:double(16,2);default:null" json:"capVag"`                         // 平均市值
	Num          int       `gorm:"column:num;type:int(11);default:null" json:"num"`                                     // 家数
	LossNum      int       `gorm:"column:loss_num;type:int(11);default:null" json:"lossNum"`                            // 亏损家数
	CreatedBy    string    `gorm:"column:created_by;type:varchar(255);default:null;default:''" json:"createdBy"`        // 创建者
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime;default:null" json:"createdAt"`                       // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *HyDailyTbl) TableName() string {
	return "hy_daily_tbl"
}

// HyDailyTblColumns get sql column name.获取数据库列名
var HyDailyTblColumns = struct {
	ID           string
	HyCode       string
	HyName       string
	Price        string
	Percent      string
	Ma5          string
	Ma10         string
	Ma20         string
	Ma60         string
	Day0         string
	Day0Str      string
	Volume       string
	Turnover     string
	TurnoverRate string
	Zljlr        string
	OpenPrice    string
	ClosePrice   string
	Max          string
	Min          string
	Pe           string
	Peg          string
	Pb           string
	CapVag       string
	Num          string
	LossNum      string
	CreatedBy    string
	CreatedAt    string
}{
	ID:           "id",
	HyCode:       "hy_code",
	HyName:       "hy_name",
	Price:        "price",
	Percent:      "percent",
	Ma5:          "ma5",
	Ma10:         "ma10",
	Ma20:         "ma20",
	Ma60:         "ma60",
	Day0:         "day0",
	Day0Str:      "day0_str",
	Volume:       "volume",
	Turnover:     "turnover",
	TurnoverRate: "turnover_rate",
	Zljlr:        "zljlr",
	OpenPrice:    "open_price",
	ClosePrice:   "close_price",
	Max:          "max",
	Min:          "min",
	Pe:           "pe",
	Peg:          "peg",
	Pb:           "pb",
	CapVag:       "cap_vag",
	Num:          "num",
	LossNum:      "loss_num",
	CreatedBy:    "created_by",
	CreatedAt:    "created_at",
}

// HyInfoTbl 行业分类
type HyInfoTbl struct {
	ID        int       `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	HyCode    string    `gorm:"unique;column:hy_code;type:varchar(255);default:null" json:"hyCode"` // 行业代码代码
	Name      string    `gorm:"unique;column:name;type:varchar(255);default:null" json:"name"`      // 行业名
	Num       int       `gorm:"column:num;type:int(11);default:null" json:"num"`                    // 家数
	Pe        float64   `gorm:"column:pe;type:double(16,2);default:null" json:"pe"`                 // 市盈率
	Peg       float64   `gorm:"column:peg;type:double(16,2);default:null" json:"peg"`               // PEG值
	Pb        float64   `gorm:"column:pb;type:double(16,2);default:null" json:"pb"`                 // 市净率
	CapVag    float64   `gorm:"column:cap_vag;type:double(16,2);default:null" json:"capVag"`        // 平均市值
	LossNum   int       `gorm:"column:loss_num;type:int(11);default:null" json:"lossNum"`           // 亏损家数
	UpdateAt  time.Time `gorm:"column:update_at;type:datetime;default:null" json:"updateAt"`        // 创建时间
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;default:null" json:"createdAt"`      // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *HyInfoTbl) TableName() string {
	return "hy_info_tbl"
}

// HyInfoTblColumns get sql column name.获取数据库列名
var HyInfoTblColumns = struct {
	ID        string
	HyCode    string
	Name      string
	Num       string
	Pe        string
	Peg       string
	Pb        string
	CapVag    string
	LossNum   string
	UpdateAt  string
	CreatedAt string
}{
	ID:        "id",
	HyCode:    "hy_code",
	Name:      "name",
	Num:       "num",
	Pe:        "pe",
	Peg:       "peg",
	Pb:        "pb",
	CapVag:    "cap_vag",
	LossNum:   "loss_num",
	UpdateAt:  "update_at",
	CreatedAt: "created_at",
}

// HyUpTbl 行业分类
type HyUpTbl struct {
	ID           int            `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Code         string         `gorm:"uniqueIndex:name;column:code;type:varchar(255);default:null" json:"code"` // 行业代码
	Name         string         `gorm:"column:name;type:varchar(255);default:null" json:"name"`                  // 行业名
	Percent      float64        `gorm:"column:percent;type:double(16,2);default:null" json:"percent"`            // 百分比
	TurnoverRate float64        `gorm:"column:turnover_rate;type:double(16,2);default:null" json:"turnoverRate"` // 换手率
	Num          int            `gorm:"column:num;type:int(11);default:null" json:"num"`                         // 总家数
	Up           int            `gorm:"column:up;type:int(11);default:null" json:"up"`                           // 上涨家数
	CreatedAt    time.Time      `gorm:"column:created_at;type:datetime;default:null" json:"createdAt"`           // 创建时间
	Day          datatypes.Date `gorm:"uniqueIndex:name;column:day;type:date;default:null" json:"day"`           // 入池时间
	DayStr       string         `gorm:"column:day_str;type:varchar(255);default:null" json:"dayStr"`             // 入池时间
}

// TableName get sql table name.获取数据库表名
func (m *HyUpTbl) TableName() string {
	return "hy_up_tbl"
}

// HyUpTblColumns get sql column name.获取数据库列名
var HyUpTblColumns = struct {
	ID           string
	Code         string
	Name         string
	Percent      string
	TurnoverRate string
	Num          string
	Up           string
	CreatedAt    string
	Day          string
	DayStr       string
}{
	ID:           "id",
	Code:         "code",
	Name:         "name",
	Percent:      "percent",
	TurnoverRate: "turnover_rate",
	Num:          "num",
	Up:           "up",
	CreatedAt:    "created_at",
	Day:          "day",
	DayStr:       "day_str",
}

// LhbDailyTbl 主力净流入列表
type LhbDailyTbl struct {
	ID        int       `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Code      string    `gorm:"uniqueIndex:code;column:code;type:varchar(255);default:null" json:"code"` // 股票代码
	Name      string    `gorm:"column:name;type:varchar(255);default:null" json:"name"`                  // 股票名字
	Day0      int64     `gorm:"uniqueIndex:code;column:day0;type:bigint(11);default:null" json:"day0"`   // 当日0点时间戳
	DayStr    string    `gorm:"column:day_str;type:varchar(255);default:null" json:"dayStr"`             // 入池时间
	Price     float64   `gorm:"column:price;type:double(16,4);default:null" json:"price"`                // 当前价格
	Jlr       float64   `gorm:"column:jlr;type:double(16,4);default:null" json:"jlr"`                    // 龙虎榜净流入(万元)
	JgJlr     float64   `gorm:"column:jg_jlr;type:double(16,4);default:null" json:"jgJlr"`               // 机构净流入(万元)
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;default:null" json:"createdAt"`           // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *LhbDailyTbl) TableName() string {
	return "lhb_daily_tbl"
}

// LhbDailyTblColumns get sql column name.获取数据库列名
var LhbDailyTblColumns = struct {
	ID        string
	Code      string
	Name      string
	Day0      string
	DayStr    string
	Price     string
	Jlr       string
	JgJlr     string
	CreatedAt string
}{
	ID:        "id",
	Code:      "code",
	Name:      "name",
	Day0:      "day0",
	DayStr:    "day_str",
	Price:     "price",
	Jlr:       "jlr",
	JgJlr:     "jg_jlr",
	CreatedAt: "created_at",
}

// LogTbl 日志
type LogTbl struct {
	ID   int    `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Desc string `gorm:"column:desc;type:varchar(255);default:null" json:"desc"` // 日志内容
}

// TableName get sql table name.获取数据库表名
func (m *LogTbl) TableName() string {
	return "log_tbl"
}

// LogTblColumns get sql column name.获取数据库列名
var LogTblColumns = struct {
	ID   string
	Desc string
}{
	ID:   "id",
	Desc: "desc",
}

// Ma20View VIEW
type Ma20View struct {
	Code      string  `gorm:"column:code;type:varchar(255);default:null" json:"code"`             // 股票代码
	Name      string  `gorm:"column:name;type:varchar(255);default:null" json:"name"`             // 股票名字
	Price     float64 `gorm:"column:price;type:double(16,2);default:null" json:"price"`           // 当前价格
	Ma20      float64 `gorm:"column:ma20;type:double(16,2);default:null" json:"ma20"`             // 20日均线
	Peg       float64 `gorm:"column:peg;type:double(16,3);default:null" json:"peg"`               // peg估值
	PegShares string  `gorm:"column:peg_shares;type:varchar(1024);default:null" json:"pegShares"` // 估值比较
	HyName    string  `gorm:"column:hy_name;type:text;default:null" json:"hyName"`                // 行业名
	NickName  string  `gorm:"column:nick_name;type:varchar(255);default:null" json:"nickName"`    // 用户昵称
	Day0Str   string  `gorm:"column:day0_str;type:varchar(255);default:null" json:"day0Str"`      // 当日0点时间戳
}

// TableName get sql table name.获取数据库表名
func (m *Ma20View) TableName() string {
	return "ma20_view"
}

// Ma20ViewColumns get sql column name.获取数据库列名
var Ma20ViewColumns = struct {
	Code      string
	Name      string
	Price     string
	Ma20      string
	Peg       string
	PegShares string
	HyName    string
	NickName  string
	Day0Str   string
}{
	Code:      "code",
	Name:      "name",
	Price:     "price",
	Ma20:      "ma20",
	Peg:       "peg",
	PegShares: "peg_shares",
	HyName:    "hy_name",
	NickName:  "nick_name",
	Day0Str:   "day0_str",
}

// MaCdTbl 超跌赛选器
type MaCdTbl struct {
	ID           int            `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Code         string         `gorm:"unique;column:code;type:varchar(255);default:null" json:"code"`           // 股票代码
	Name         string         `gorm:"column:name;type:varchar(255);default:null" json:"name"`                  // 股票名字
	Day          datatypes.Date `gorm:"column:day;type:date;default:null" json:"day"`                            // 入池时间
	DayStr       string         `gorm:"column:day_str;type:varchar(255);default:null" json:"dayStr"`             // 入池时间
	Price        float64        `gorm:"column:price;type:double(11,2);default:null" json:"price"`                // 价格
	NewPrice     float64        `gorm:"column:new_price;type:double(11,2);default:null" json:"newPrice"`         // 价格
	Percent      float64        `gorm:"column:percent;type:double(16,2);default:null" json:"percent"`            // 百分比
	TurnoverRate float64        `gorm:"column:turnover_rate;type:double(16,2);default:null" json:"turnoverRate"` // 换手率
	Score        int            `gorm:"column:score;type:int(11);default:null" json:"score"`                     // 得分
	CreatedAt    time.Time      `gorm:"column:created_at;type:datetime;default:null" json:"createdAt"`           // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *MaCdTbl) TableName() string {
	return "ma_cd_tbl"
}

// MaCdTblColumns get sql column name.获取数据库列名
var MaCdTblColumns = struct {
	ID           string
	Code         string
	Name         string
	Day          string
	DayStr       string
	Price        string
	NewPrice     string
	Percent      string
	TurnoverRate string
	Score        string
	CreatedAt    string
}{
	ID:           "id",
	Code:         "code",
	Name:         "name",
	Day:          "day",
	DayStr:       "day_str",
	Price:        "price",
	NewPrice:     "new_price",
	Percent:      "percent",
	TurnoverRate: "turnover_rate",
	Score:        "score",
	CreatedAt:    "created_at",
}

// MaFlTbl 放量赛选器
type MaFlTbl struct {
	ID           int            `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Code         string         `gorm:"unique;column:code;type:varchar(255);default:null" json:"code"`           // 股票代码
	Name         string         `gorm:"column:name;type:varchar(255);default:null" json:"name"`                  // 股票名字
	Day          datatypes.Date `gorm:"column:day;type:date;default:null" json:"day"`                            // 入池时间
	DayStr       string         `gorm:"column:day_str;type:varchar(255);default:null" json:"dayStr"`             // 入池时间
	Price        float64        `gorm:"column:price;type:double(11,2);default:null" json:"price"`                // 价格
	NewPrice     float64        `gorm:"column:new_price;type:double(11,2);default:null" json:"newPrice"`         // 价格
	Percent      float64        `gorm:"column:percent;type:double(16,2);default:null" json:"percent"`            // 百分比
	TurnoverRate float64        `gorm:"column:turnover_rate;type:double(16,2);default:null" json:"turnoverRate"` // 换手率
	Score        int            `gorm:"column:score;type:int(11);default:null" json:"score"`                     // 得分
	CreatedAt    time.Time      `gorm:"column:created_at;type:datetime;default:null" json:"createdAt"`           // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *MaFlTbl) TableName() string {
	return "ma_fl_tbl"
}

// MaFlTblColumns get sql column name.获取数据库列名
var MaFlTblColumns = struct {
	ID           string
	Code         string
	Name         string
	Day          string
	DayStr       string
	Price        string
	NewPrice     string
	Percent      string
	TurnoverRate string
	Score        string
	CreatedAt    string
}{
	ID:           "id",
	Code:         "code",
	Name:         "name",
	Day:          "day",
	DayStr:       "day_str",
	Price:        "price",
	NewPrice:     "new_price",
	Percent:      "percent",
	TurnoverRate: "turnover_rate",
	Score:        "score",
	CreatedAt:    "created_at",
}

// MaLhbTbl 主力净流入列表
type MaLhbTbl struct {
	ID           int       `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Code         string    `gorm:"uniqueIndex:code;column:code;type:varchar(255);default:null" json:"code"` // 股票代码
	Name         string    `gorm:"column:name;type:varchar(255);default:null" json:"name"`                  // 股票名字
	Day0         int64     `gorm:"uniqueIndex:code;column:day0;type:bigint(11);default:null" json:"day0"`   // 当日0点时间戳
	DayStr       string    `gorm:"column:day_str;type:varchar(255);default:null" json:"dayStr"`             // 入池时间
	Jlr          float64   `gorm:"column:jlr;type:double(11,2);default:null" json:"jlr"`                    // 价格
	JgJlr        float64   `gorm:"column:jg_jlr;type:double(16,4);default:null" json:"jgJlr"`               // 机构净流入(万元)
	Price        float64   `gorm:"column:price;type:double(16,4);default:null" json:"price"`                // 龙虎榜净流入(万元)
	NewPrice     float64   `gorm:"column:new_price;type:double(11,2);default:null" json:"newPrice"`         // 价格
	TurnoverRate float64   `gorm:"column:turnover_rate;type:double(16,2);default:null" json:"turnoverRate"` // 换手率
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime;default:null" json:"createdAt"`           // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *MaLhbTbl) TableName() string {
	return "ma_lhb_tbl"
}

// MaLhbTblColumns get sql column name.获取数据库列名
var MaLhbTblColumns = struct {
	ID           string
	Code         string
	Name         string
	Day0         string
	DayStr       string
	Jlr          string
	JgJlr        string
	Price        string
	NewPrice     string
	TurnoverRate string
	CreatedAt    string
}{
	ID:           "id",
	Code:         "code",
	Name:         "name",
	Day0:         "day0",
	DayStr:       "day_str",
	Jlr:          "jlr",
	JgJlr:        "jg_jlr",
	Price:        "price",
	NewPrice:     "new_price",
	TurnoverRate: "turnover_rate",
	CreatedAt:    "created_at",
}

// MaUpTbl 日线多头赛选器
type MaUpTbl struct {
	ID           int            `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Code         string         `gorm:"unique;column:code;type:varchar(255);default:null" json:"code"`           // 股票代码
	Name         string         `gorm:"column:name;type:varchar(255);default:null" json:"name"`                  // 股票名字
	Day          datatypes.Date `gorm:"column:day;type:date;default:null" json:"day"`                            // 入池时间
	DayStr       string         `gorm:"column:day_str;type:varchar(255);default:null" json:"dayStr"`             // 入池时间
	Price        float64        `gorm:"column:price;type:double(11,2);default:null" json:"price"`                // 价格
	NewPrice     float64        `gorm:"column:new_price;type:double(11,2);default:null" json:"newPrice"`         // 价格
	Percent      float64        `gorm:"column:percent;type:double(16,2);default:null" json:"percent"`            // 百分比
	TurnoverRate float64        `gorm:"column:turnover_rate;type:double(16,2);default:null" json:"turnoverRate"` // 换手率
	CreatedAt    time.Time      `gorm:"column:created_at;type:datetime;default:null" json:"createdAt"`           // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *MaUpTbl) TableName() string {
	return "ma_up_tbl"
}

// MaUpTblColumns get sql column name.获取数据库列名
var MaUpTblColumns = struct {
	ID           string
	Code         string
	Name         string
	Day          string
	DayStr       string
	Price        string
	NewPrice     string
	Percent      string
	TurnoverRate string
	CreatedAt    string
}{
	ID:           "id",
	Code:         "code",
	Name:         "name",
	Day:          "day",
	DayStr:       "day_str",
	Price:        "price",
	NewPrice:     "new_price",
	Percent:      "percent",
	TurnoverRate: "turnover_rate",
	CreatedAt:    "created_at",
}

// MsgCodeView VIEW
type MsgCodeView struct {
	Code string `gorm:"column:code;type:varchar(255);default:null" json:"code"`
	Tp   string `gorm:"column:tp;type:varchar(5);not null;default:''" json:"tp"`
}

// TableName get sql table name.获取数据库表名
func (m *MsgCodeView) TableName() string {
	return "msg_code_view"
}

// MsgCodeViewColumns get sql column name.获取数据库列名
var MsgCodeViewColumns = struct {
	Code string
	Tp   string
}{
	Code: "code",
	Tp:   "tp",
}

// MsgRapidlyTbl 股票急速上涨下跌提醒
type MsgRapidlyTbl struct {
	ID            int            `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Code          string         `gorm:"uniqueIndex:code;column:code;type:varchar(255);default:null" json:"code"`   // 股票代码
	Tag           int            `gorm:"uniqueIndex:code;column:tag;type:int(11);default:null" json:"tag"`          // 1:上涨，0:下跌
	Key           string         `gorm:"column:key;type:varchar(255);default:null" json:"key"`                      // 消息类型
	Desc          string         `gorm:"column:desc;type:varchar(255);default:null" json:"desc"`                    // 消息描述
	Old           float64        `gorm:"column:old;type:double(10,2);default:null" json:"old"`                      // 之前价格
	New           float64        `gorm:"column:new;type:double(10,2);default:null" json:"new"`                      // 当前价格
	Percent       float64        `gorm:"column:percent;type:double(10,2);default:null" json:"percent"`              // 百分比
	OffsetPercent float64        `gorm:"column:offset_percent;type:double(10,2);default:null" json:"offsetPercent"` // 相对百分比
	Day           datatypes.Date `gorm:"uniqueIndex:code;column:day;type:date;default:null" json:"day"`             // 当日0点时间戳
	CreatedAt     time.Time      `gorm:"column:created_at;type:datetime;default:null" json:"createdAt"`             // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *MsgRapidlyTbl) TableName() string {
	return "msg_rapidly_tbl"
}

// MsgRapidlyTblColumns get sql column name.获取数据库列名
var MsgRapidlyTblColumns = struct {
	ID            string
	Code          string
	Tag           string
	Key           string
	Desc          string
	Old           string
	New           string
	Percent       string
	OffsetPercent string
	Day           string
	CreatedAt     string
}{
	ID:            "id",
	Code:          "code",
	Tag:           "tag",
	Key:           "key",
	Desc:          "desc",
	Old:           "old",
	New:           "new",
	Percent:       "percent",
	OffsetPercent: "offset_percent",
	Day:           "day",
	CreatedAt:     "created_at",
}

// MsgTbl 推送消息
type MsgTbl struct {
	ID        int            `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	OpenID    string         `gorm:"column:open_id;type:varchar(255);default:null" json:"openId"`       // 用户openid
	Code      string         `gorm:"index:code;column:code;type:varchar(255);default:null" json:"code"` // 股票代码
	Key       string         `gorm:"column:key;type:varchar(255);default:null" json:"key"`              // 消息类型
	Desc      string         `gorm:"column:desc;type:varchar(1024);default:null" json:"desc"`           // 消息描述
	Price     float64        `gorm:"column:price;type:double(10,2);default:null" json:"price"`          // 当前价格
	Percent   float64        `gorm:"column:percent;type:double(10,2);default:null" json:"percent"`      // 百分比
	Day       datatypes.Date `gorm:"column:day;type:date;default:null" json:"day"`                      // 当日0点时间戳
	CreatedAt time.Time      `gorm:"column:created_at;type:datetime;default:null" json:"createdAt"`     // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *MsgTbl) TableName() string {
	return "msg_tbl"
}

// MsgTblColumns get sql column name.获取数据库列名
var MsgTblColumns = struct {
	ID        string
	OpenID    string
	Code      string
	Key       string
	Desc      string
	Price     string
	Percent   string
	Day       string
	CreatedAt string
}{
	ID:        "id",
	OpenID:    "open_id",
	Code:      "code",
	Key:       "key",
	Desc:      "desc",
	Price:     "price",
	Percent:   "percent",
	Day:       "day",
	CreatedAt: "created_at",
}

// MyselfTbl 放量赛选器
type MyselfTbl struct {
	ID           int            `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Code         string         `gorm:"unique;column:code;type:varchar(255);default:null" json:"code"`           // 股票代码
	Name         string         `gorm:"column:name;type:varchar(255);default:null" json:"name"`                  // 股票名字
	Day          datatypes.Date `gorm:"column:day;type:date;default:null" json:"day"`                            // 入池时间
	DayStr       string         `gorm:"column:day_str;type:varchar(255);default:null" json:"dayStr"`             // 入池时间
	Price        float64        `gorm:"column:price;type:double(11,2);default:null" json:"price"`                // 价格
	NewPrice     float64        `gorm:"column:new_price;type:double(11,2);default:null" json:"newPrice"`         // 价格
	Percent      float64        `gorm:"column:percent;type:double(16,2);default:null" json:"percent"`            // 百分比
	TurnoverRate float64        `gorm:"column:turnover_rate;type:double(16,2);default:null" json:"turnoverRate"` // 换手率
	Number       int            `gorm:"column:number;type:int(11);default:null" json:"number"`                   // 数量
	Desc         string         `gorm:"column:desc;type:varchar(1024);default:null" json:"desc"`                 // 描述
	CreatedAt    time.Time      `gorm:"column:created_at;type:datetime;default:null" json:"createdAt"`           // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *MyselfTbl) TableName() string {
	return "myself_tbl"
}

// MyselfTblColumns get sql column name.获取数据库列名
var MyselfTblColumns = struct {
	ID           string
	Code         string
	Name         string
	Day          string
	DayStr       string
	Price        string
	NewPrice     string
	Percent      string
	TurnoverRate string
	Number       string
	Desc         string
	CreatedAt    string
}{
	ID:           "id",
	Code:         "code",
	Name:         "name",
	Day:          "day",
	DayStr:       "day_str",
	Price:        "price",
	NewPrice:     "new_price",
	Percent:      "percent",
	TurnoverRate: "turnover_rate",
	Number:       "number",
	Desc:         "desc",
	CreatedAt:    "created_at",
}

// NoPegTbl peg剔除列表
type NoPegTbl struct {
	ID        int       `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Code      string    `gorm:"unique;column:code;type:varchar(255);default:null" json:"code"` // 股票代码
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;default:null" json:"createdAt"` // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *NoPegTbl) TableName() string {
	return "no_peg_tbl"
}

// NoPegTblColumns get sql column name.获取数据库列名
var NoPegTblColumns = struct {
	ID        string
	Code      string
	CreatedAt string
}{
	ID:        "id",
	Code:      "code",
	CreatedAt: "created_at",
}

// NoTdxTbl 通达信剔除列表
type NoTdxTbl struct {
	ID        int       `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Code      string    `gorm:"unique;column:code;type:varchar(255);default:null" json:"code"` // 股票代码
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;default:null" json:"createdAt"` // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *NoTdxTbl) TableName() string {
	return "no_tdx_tbl"
}

// NoTdxTblColumns get sql column name.获取数据库列名
var NoTdxTblColumns = struct {
	ID        string
	Code      string
	CreatedAt string
}{
	ID:        "id",
	Code:      "code",
	CreatedAt: "created_at",
}

// NoTradingDayTbl 非交易日数据
type NoTradingDayTbl struct {
	ID  int            `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Day datatypes.Date `gorm:"unique;column:day;type:date;default:null" json:"day"` // 当日0点时间戳
}

// TableName get sql table name.获取数据库表名
func (m *NoTradingDayTbl) TableName() string {
	return "no_trading_day_tbl"
}

// NoTradingDayTblColumns get sql column name.获取数据库列名
var NoTradingDayTblColumns = struct {
	ID  string
	Day string
}{
	ID:  "id",
	Day: "day",
}

// SharesDailyTbl 日数据，每日的数据
type SharesDailyTbl struct {
	ID           int       `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Code         string    `gorm:"uniqueIndex:code;column:code;type:varchar(255);default:null" json:"code"`          // 股票代码
	Price        float64   `gorm:"column:price;type:double(16,2);default:null" json:"price"`                         // 当前价格
	Percent      float64   `gorm:"column:percent;type:double(16,2);default:null" json:"percent"`                     // 百分比
	Ma5          float64   `gorm:"column:ma5;type:double(16,2);default:null" json:"ma5"`                             // 5日均线
	Ma10         float64   `gorm:"column:ma10;type:double(16,2);default:null" json:"ma10"`                           // 10日均线
	Ma20         float64   `gorm:"column:ma20;type:double(16,2);default:null" json:"ma20"`                           // 20日均线
	Ma30         float64   `gorm:"column:ma30;type:double(16,2);default:null" json:"ma30"`                           // 20日均线
	Ma60         float64   `gorm:"column:ma60;type:double(16,2);default:null" json:"ma60"`                           // 20日均线
	Day0         int64     `gorm:"uniqueIndex:code;index:day0;column:day0;type:bigint(11);default:null" json:"day0"` // 当日0点时间戳
	Day0Str      string    `gorm:"column:day0_str;type:varchar(255);default:null" json:"day0Str"`                    // 当日0点时间戳
	Volume       float64   `gorm:"column:volume;type:double(16,2);default:null" json:"volume"`                       // 成交量（手）
	Turnover     float64   `gorm:"column:turnover;type:double(16,2);default:null" json:"turnover"`                   // 成交额（万）
	TurnoverRate float64   `gorm:"column:turnover_rate;type:double(16,2);default:null" json:"turnoverRate"`          // 换手率
	Vol          float64   `gorm:"column:vol;type:double(16,2);default:null" json:"vol"`                             // 成交量占比 (volume / volume 前5日平均)
	Pe           float64   `gorm:"column:pe;type:double(16,2);default:null" json:"pe"`                               // 市盈率
	Pb           float64   `gorm:"column:pb;type:double(16,2);default:null" json:"pb"`                               // 市净率
	Total        float64   `gorm:"column:total;type:double(16,2);default:null" json:"total"`                         // 流通市值
	Cap          float64   `gorm:"column:cap;type:double(16,2);default:null" json:"cap"`                             // 总市值
	Zljlr        float64   `gorm:"column:zljlr;type:double(16,4);default:null" json:"zljlr"`                         // 主力净流入
	OpenPrice    float64   `gorm:"column:open_price;type:double(16,2);default:null" json:"openPrice"`                // 开盘价
	ClosePrice   float64   `gorm:"column:close_price;type:double(16,2);default:null" json:"closePrice"`              // 收盘价
	Bscg         float64   `gorm:"column:bscg;type:double(16,4);default:null" json:"bscg"`                           // 北上持股数
	Bsjlr        float64   `gorm:"column:bsjlr;type:double(16,4);default:null" json:"bsjlr"`                         // 北上净流入
	BsPercent    float64   `gorm:"column:bs_percent;type:double(16,4);default:null" json:"bsPercent"`                // 北上持股占比
	Macd         float64   `gorm:"column:macd;type:double(16,3);default:null" json:"macd"`                           // macd 指标
	Dif          float64   `gorm:"column:dif;type:double(16,3);default:null" json:"dif"`                             // DIF快线
	Dea          float64   `gorm:"column:dea;type:double(16,3);default:null" json:"dea"`                             // Dea 慢线
	Max          float64   `gorm:"column:max;type:double(16,2);default:null" json:"max"`                             // 当日最高值
	Min          float64   `gorm:"column:min;type:double(16,2);default:null" json:"min"`                             // 当日最低价
	Peg          float64   `gorm:"column:peg;type:double(16,2);default:null" json:"peg"`                             // peg估值
	CreatedBy    string    `gorm:"column:created_by;type:varchar(255);default:null;default:''" json:"createdBy"`     // 创建者
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime;default:null" json:"createdAt"`                    // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *SharesDailyTbl) TableName() string {
	return "shares_daily_tbl"
}

// SharesDailyTblColumns get sql column name.获取数据库列名
var SharesDailyTblColumns = struct {
	ID           string
	Code         string
	Price        string
	Percent      string
	Ma5          string
	Ma10         string
	Ma20         string
	Ma30         string
	Ma60         string
	Day0         string
	Day0Str      string
	Volume       string
	Turnover     string
	TurnoverRate string
	Vol          string
	Pe           string
	Pb           string
	Total        string
	Cap          string
	Zljlr        string
	OpenPrice    string
	ClosePrice   string
	Bscg         string
	Bsjlr        string
	BsPercent    string
	Macd         string
	Dif          string
	Dea          string
	Max          string
	Min          string
	Peg          string
	CreatedBy    string
	CreatedAt    string
}{
	ID:           "id",
	Code:         "code",
	Price:        "price",
	Percent:      "percent",
	Ma5:          "ma5",
	Ma10:         "ma10",
	Ma20:         "ma20",
	Ma30:         "ma30",
	Ma60:         "ma60",
	Day0:         "day0",
	Day0Str:      "day0_str",
	Volume:       "volume",
	Turnover:     "turnover",
	TurnoverRate: "turnover_rate",
	Vol:          "vol",
	Pe:           "pe",
	Pb:           "pb",
	Total:        "total",
	Cap:          "cap",
	Zljlr:        "zljlr",
	OpenPrice:    "open_price",
	ClosePrice:   "close_price",
	Bscg:         "bscg",
	Bsjlr:        "bsjlr",
	BsPercent:    "bs_percent",
	Macd:         "macd",
	Dif:          "dif",
	Dea:          "dea",
	Max:          "max",
	Min:          "min",
	Peg:          "peg",
	CreatedBy:    "created_by",
	CreatedAt:    "created_at",
}

// SharesInfoTbl 实时股票信息
type SharesInfoTbl struct {
	ID         int       `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Code       string    `gorm:"unique;column:code;type:varchar(255);default:null" json:"code"`                // 股票代码
	SimpleCode string    `gorm:"column:simple_code;type:varchar(255);default:null" json:"simpleCode"`          // 股票代码简写
	Ext        string    `gorm:"column:ext;type:varchar(255);default:null" json:"ext"`                         // 后缀
	Name       string    `gorm:"column:name;type:varchar(255);default:null" json:"name"`                       // 股票名字
	SimpleName string    `gorm:"column:simple_name;type:varchar(255);default:null" json:"simpleName"`          // 股票名字拼音简写
	Price      float64   `gorm:"column:price;type:double(16,2);default:null" json:"price"`                     // 当前价格
	Percent    float64   `gorm:"column:percent;type:double(16,2);default:null" json:"percent"`                 // 百分比
	Total      float64   `gorm:"column:total;type:double(16,2);default:null" json:"total"`                     // 流通市值
	Cap        float64   `gorm:"column:cap;type:double(16,2);default:null" json:"cap"`                         // 总市值
	Peg        float64   `gorm:"column:peg;type:double(16,3);default:null" json:"peg"`                         // peg估值
	PegAvg     float64   `gorm:"column:peg_avg;type:double(16,3);default:null" json:"pegAvg"`                  // peg行业平均
	PegAve     float64   `gorm:"column:peg_ave;type:double(16,3);default:null" json:"pegAve"`                  // peg行业中值
	PegShares  string    `gorm:"column:peg_shares;type:varchar(1024);default:null" json:"pegShares"`           // 估值比较
	HyName     string    `gorm:"column:hy_name;type:text;default:null" json:"hyName"`                          // 行业名
	CreatedBy  string    `gorm:"column:created_by;type:varchar(255);default:null;default:''" json:"createdBy"` // 创建者
	CreatedAt  time.Time `gorm:"column:created_at;type:datetime;default:null" json:"createdAt"`                // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *SharesInfoTbl) TableName() string {
	return "shares_info_tbl"
}

// SharesInfoTblColumns get sql column name.获取数据库列名
var SharesInfoTblColumns = struct {
	ID         string
	Code       string
	SimpleCode string
	Ext        string
	Name       string
	SimpleName string
	Price      string
	Percent    string
	Total      string
	Cap        string
	Peg        string
	PegAvg     string
	PegAve     string
	PegShares  string
	HyName     string
	CreatedBy  string
	CreatedAt  string
}{
	ID:         "id",
	Code:       "code",
	SimpleCode: "simple_code",
	Ext:        "ext",
	Name:       "name",
	SimpleName: "simple_name",
	Price:      "price",
	Percent:    "percent",
	Total:      "total",
	Cap:        "cap",
	Peg:        "peg",
	PegAvg:     "peg_avg",
	PegAve:     "peg_ave",
	PegShares:  "peg_shares",
	HyName:     "hy_name",
	CreatedBy:  "created_by",
	CreatedAt:  "created_at",
}

// SharesWatchTbl 实时监听信息
type SharesWatchTbl struct {
	ID            int       `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Code          string    `gorm:"uniqueIndex:open_code;index:code;column:code;type:varchar(255);default:null" json:"code"` // 股票代码
	OpenID        string    `gorm:"uniqueIndex:open_code;column:open_id;type:varchar(255);default:null" json:"openId"`       // 用户openid
	Kdj           bool      `gorm:"column:kdj;type:tinyint(1);default:null" json:"kdj"`                                      // 日线金叉提醒
	Kdj20         bool      `gorm:"column:kdj20;type:tinyint(1);default:null" json:"kdj20"`                                  // 20日线提醒
	Surge         bool      `gorm:"column:surge;type:tinyint(1);default:null" json:"surge"`                                  // 盘中急涨提醒
	Slump         bool      `gorm:"column:slump;type:tinyint(1);default:null" json:"slump"`                                  // 盘中急跌提醒
	Ai            bool      `gorm:"column:ai;type:tinyint(1);default:null" json:"ai"`                                        // AI智能提醒
	GroupWi       bool      `gorm:"column:group_wi;type:tinyint(1);default:null" json:"groupWi"`                             // 是否推荐给组织
	Up            float64   `gorm:"column:up;type:double(11,2);default:null" json:"up"`                                      // 估价涨到
	UpVaild       bool      `gorm:"column:up_vaild;type:tinyint(1);default:null" json:"upVaild"`                             // 是否提醒
	Down          float64   `gorm:"column:down;type:double(11,2);default:null" json:"down"`                                  // 估价跌到
	DownVaild     bool      `gorm:"column:down_vaild;type:tinyint(1);default:null" json:"downVaild"`                         // 是否提醒
	UpPercent     float64   `gorm:"column:up_percent;type:double(11,2);default:null" json:"upPercent"`                       // 涨幅超百分比
	UpPercentTo   float64   `gorm:"column:up_percent_to;type:double(11,2);default:null" json:"upPercentTo"`                  // 涨幅超(辅助)
	DownPercent   float64   `gorm:"column:down_percent;type:double(11,2);default:null" json:"downPercent"`                   // 跌幅超百分比
	DownPercentTo float64   `gorm:"column:down_percent_to;type:double(11,2);default:null" json:"downPercentTo"`              // 跌幅超(辅助)
	Vaild         bool      `gorm:"column:vaild;type:tinyint(1);default:null" json:"vaild"`                                  // 是否暂停提醒
	CreatedAt     time.Time `gorm:"column:created_at;type:datetime;default:null" json:"createdAt"`                           // 创建时间
	ExpiresTime   time.Time `gorm:"column:expires_time;type:datetime;default:null" json:"expiresTime"`                       // 过期时间
}

// TableName get sql table name.获取数据库表名
func (m *SharesWatchTbl) TableName() string {
	return "shares_watch_tbl"
}

// SharesWatchTblColumns get sql column name.获取数据库列名
var SharesWatchTblColumns = struct {
	ID            string
	Code          string
	OpenID        string
	Kdj           string
	Kdj20         string
	Surge         string
	Slump         string
	Ai            string
	GroupWi       string
	Up            string
	UpVaild       string
	Down          string
	DownVaild     string
	UpPercent     string
	UpPercentTo   string
	DownPercent   string
	DownPercentTo string
	Vaild         string
	CreatedAt     string
	ExpiresTime   string
}{
	ID:            "id",
	Code:          "code",
	OpenID:        "open_id",
	Kdj:           "kdj",
	Kdj20:         "kdj20",
	Surge:         "surge",
	Slump:         "slump",
	Ai:            "ai",
	GroupWi:       "group_wi",
	Up:            "up",
	UpVaild:       "up_vaild",
	Down:          "down",
	DownVaild:     "down_vaild",
	UpPercent:     "up_percent",
	UpPercentTo:   "up_percent_to",
	DownPercent:   "down_percent",
	DownPercentTo: "down_percent_to",
	Vaild:         "vaild",
	CreatedAt:     "created_at",
	ExpiresTime:   "expires_time",
}

// TdxDailyTbl 主力净流入列表
type TdxDailyTbl struct {
	ID        int       `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Code      string    `gorm:"uniqueIndex:code;column:code;type:varchar(255);default:null" json:"code"` // 股票代码
	Name      string    `gorm:"column:name;type:varchar(255);default:null" json:"name"`                  // 股票名字
	Day0      int64     `gorm:"uniqueIndex:code;column:day0;type:bigint(11);default:null" json:"day0"`   // 当日0点时间戳
	DayStr    string    `gorm:"column:day_str;type:varchar(255);default:null" json:"dayStr"`             // 入池时间
	Open      float64   `gorm:"column:open;type:double(16,2);default:null" json:"open"`                  // 当日开盘价
	Close     float64   `gorm:"column:close;type:double(16,2);default:null" json:"close"`                // 当日收盘价
	Min       float64   `gorm:"column:min;type:double(16,2);default:null" json:"min"`                    // 当日最低价
	Max       float64   `gorm:"column:max;type:double(16,2);default:null" json:"max"`                    // 当日最高值
	Volume    float64   `gorm:"column:volume;type:double(16,2);default:null" json:"volume"`              // 成交量（手）
	Fyyx      bool      `gorm:"column:fyyx;type:tinyint(1);default:null" json:"fyyx"`                    // 飞鹰优选
	Zlzxh     float64   `gorm:"column:zlzxh;type:double(16,2);default:null" json:"zlzxh"`                // 主力真吸货
	Fx        float64   `gorm:"column:fx;type:double(16,2);default:null" json:"fx"`                      // 风险(>75 风险较大，<10 风险较小)
	Smx       float64   `gorm:"column:smx;type:double(16,2);default:null" json:"smx"`                    // 生命线
	Dwzq      float64   `gorm:"column:dwzq;type:double(16,2);default:null" json:"dwzq"`                  // 低位转强(50)
	Ksls      float64   `gorm:"column:ksls;type:double(16,2);default:null" json:"ksls"`                  // 开始拉升(50)
	Jdz       float64   `gorm:"column:jdz;type:double(16,2);default:null" json:"jdz"`                    // 极低涨(50)
	Tqa       int       `gorm:"column:tqa;type:int(11);default:null" json:"tqa"`                         // 唐奇安交易通道指标：-1卖，1买
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;default:null" json:"createdAt"`           // 创建时间
	UpdateAt  time.Time `gorm:"column:update_at;type:datetime;default:null" json:"updateAt"`             // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *TdxDailyTbl) TableName() string {
	return "tdx_daily_tbl"
}

// TdxDailyTblColumns get sql column name.获取数据库列名
var TdxDailyTblColumns = struct {
	ID        string
	Code      string
	Name      string
	Day0      string
	DayStr    string
	Open      string
	Close     string
	Min       string
	Max       string
	Volume    string
	Fyyx      string
	Zlzxh     string
	Fx        string
	Smx       string
	Dwzq      string
	Ksls      string
	Jdz       string
	Tqa       string
	CreatedAt string
	UpdateAt  string
}{
	ID:        "id",
	Code:      "code",
	Name:      "name",
	Day0:      "day0",
	DayStr:    "day_str",
	Open:      "open",
	Close:     "close",
	Min:       "min",
	Max:       "max",
	Volume:    "volume",
	Fyyx:      "fyyx",
	Zlzxh:     "zlzxh",
	Fx:        "fx",
	Smx:       "smx",
	Dwzq:      "dwzq",
	Ksls:      "ksls",
	Jdz:       "jdz",
	Tqa:       "tqa",
	CreatedAt: "created_at",
	UpdateAt:  "update_at",
}

// TencentTbl 微信用户信息
type TencentTbl struct {
	ID        int    `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	APIID     string `gorm:"column:api_id;type:varchar(255);default:null" json:"apiId"`         // 授权应用id
	APIKey    string `gorm:"column:api_key;type:varchar(255);default:null" json:"apiKey"`       // 授权应用key
	APISecret string `gorm:"column:api_secret;type:varchar(255);default:null" json:"apiSecret"` // 秘钥/token
	Tag       string `gorm:"column:tag;type:varchar(255);default:null" json:"tag"`              // 标记
	Ignore    string `gorm:"column:ignore;type:varchar(255);default:null" json:"ignore"`        // 忽略的词性
	IgsTag    string `gorm:"column:igs_tag;type:text;default:null" json:"igsTag"`               // 忽略关键词
}

// TableName get sql table name.获取数据库表名
func (m *TencentTbl) TableName() string {
	return "tencent_tbl"
}

// TencentTblColumns get sql column name.获取数据库列名
var TencentTblColumns = struct {
	ID        string
	APIID     string
	APIKey    string
	APISecret string
	Tag       string
	Ignore    string
	IgsTag    string
}{
	ID:        "id",
	APIID:     "api_id",
	APIKey:    "api_key",
	APISecret: "api_secret",
	Tag:       "tag",
	Ignore:    "ignore",
	IgsTag:    "igs_tag",
}

// TkTbl 跳空选择器
type TkTbl struct {
	ID        int       `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Code      string    `gorm:"uniqueIndex:code;column:code;type:varchar(255);default:null" json:"code"` // 股票代码
	Name      string    `gorm:"column:name;type:varchar(255);default:null" json:"name"`                  // 股票名字
	Day0      int64     `gorm:"uniqueIndex:code;column:day0;type:bigint(11);default:null" json:"day0"`   // 当日0点时间戳
	Day0Str   string    `gorm:"column:day0_str;type:varchar(255);default:null" json:"day0Str"`           // 入池时间
	Price     float64   `gorm:"column:price;type:double(11,2);default:null" json:"price"`                // 价格
	Min       float64   `gorm:"column:min;type:double(16,2);default:null" json:"min"`                    // 跳空最低价
	Max       float64   `gorm:"column:max;type:double(16,2);default:null" json:"max"`                    // 当日最高值
	Flag      int       `gorm:"column:flag;type:int(255);default:null" json:"flag"`                      // 标记：0 默认 ，1：回补了
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;default:null" json:"createdAt"`           // 创建时间
	UpdateAt  time.Time `gorm:"column:update_at;type:datetime;default:null" json:"updateAt"`             // 回补时间,,,回补时间1
}

// TableName get sql table name.获取数据库表名
func (m *TkTbl) TableName() string {
	return "tk_tbl"
}

// TkTblColumns get sql column name.获取数据库列名
var TkTblColumns = struct {
	ID        string
	Code      string
	Name      string
	Day0      string
	Day0Str   string
	Price     string
	Min       string
	Max       string
	Flag      string
	CreatedAt string
	UpdateAt  string
}{
	ID:        "id",
	Code:      "code",
	Name:      "name",
	Day0:      "day0",
	Day0Str:   "day0_str",
	Price:     "price",
	Min:       "min",
	Max:       "max",
	Flag:      "flag",
	CreatedAt: "created_at",
	UpdateAt:  "update_at",
}

// WxMsgTbl 推送消息
type WxMsgTbl struct {
	ID           int       `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	ToUserName   string    `gorm:"column:to_user_name;type:varchar(255);default:null" json:"toUserName"`     // 接收方帐号（收到的OpenID）
	FromUserName string    `gorm:"column:from_user_name;type:varchar(255);default:null" json:"fromUserName"` // 开发者微信号
	Createtime   int64     `gorm:"column:createTime;type:bigint(11);default:null" json:"createTime"`         // 消息创建时间 （整型）
	MsgType      string    `gorm:"column:msg_type;type:varchar(255);default:null" json:"msgType"`            // 消息类型，文本为text
	MsgID        string    `gorm:"unique;column:msg_id;type:varchar(255);default:null" json:"msgId"`         // 消息id
	Content      string    `gorm:"column:content;type:varchar(255);default:null" json:"content"`             // 回复的消息内容（换行：在content中能够换行，微信客户端就支持换行显示）
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime;default:null" json:"createdAt"`            // 创建时间
	Event        string    `gorm:"column:event;type:varchar(255);default:null" json:"event"`                 // 事件内容
	MediaID      string    `gorm:"column:media_id;type:varchar(255);default:null" json:"mediaId"`            // 语音消息媒体id，可以调用获取临时素材接口拉取数据。
	Format       string    `gorm:"column:format;type:varchar(255);default:null" json:"format"`               // 语音格式，如amr，speex等
	Req          string    `gorm:"column:req;type:varchar(1024);default:null" json:"req"`                    // 请求参数
	Status       int       `gorm:"column:status;type:int(11);default:null" json:"status"`                    // 当前状态(0:已创建,1:已回复,-1:连接已断开)
}

// TableName get sql table name.获取数据库表名
func (m *WxMsgTbl) TableName() string {
	return "wx_msg_tbl"
}

// WxMsgTblColumns get sql column name.获取数据库列名
var WxMsgTblColumns = struct {
	ID           string
	ToUserName   string
	FromUserName string
	Createtime   string
	MsgType      string
	MsgID        string
	Content      string
	CreatedAt    string
	Event        string
	MediaID      string
	Format       string
	Req          string
	Status       string
}{
	ID:           "id",
	ToUserName:   "to_user_name",
	FromUserName: "from_user_name",
	Createtime:   "createTime",
	MsgType:      "msg_type",
	MsgID:        "msg_id",
	Content:      "content",
	CreatedAt:    "created_at",
	Event:        "event",
	MediaID:      "media_id",
	Format:       "format",
	Req:          "req",
	Status:       "status",
}

// WxUserinfo 微信用户信息
type WxUserinfo struct {
	ID        int    `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Openid    string `gorm:"unique;column:openid;type:varchar(255);not null" json:"openid"`     // 微信用户唯一标识符,,微信用户唯一标识符
	NickName  string `gorm:"column:nick_name;type:varchar(255);default:null" json:"nickName"`   // 用户昵称
	AvatarURL string `gorm:"column:avatar_url;type:varchar(255);default:null" json:"avatarUrl"` // 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
	Gender    string `gorm:"column:gender;type:varchar(16);default:null" json:"gender"`         // 用户的性别
	City      string `gorm:"column:city;type:varchar(255);default:null" json:"city"`            // 用户所在城市
	Province  string `gorm:"column:province;type:varchar(255);default:null" json:"province"`    // 用户所在省份
	Country   string `gorm:"column:country;type:varchar(255);default:null" json:"country"`      // 用户所在国家
	Language  []byte `gorm:"column:language;type:varbinary(16);default:null" json:"language"`   // 用户的语言，简体中文为zh_CN
	Phone     string `gorm:"column:phone;type:varchar(255);default:null" json:"phone"`          // 用户绑定的手机号
	Group     string `gorm:"column:group;type:varchar(255);default:null" json:"group"`          // 分组
	Rg        bool   `gorm:"column:rg;type:tinyint(1);default:null" json:"rg"`                  // 是否涨红颜色
	Vip       bool   `gorm:"column:vip;type:tinyint(1);default:null" json:"vip"`                // 是否vip
	Capacity  int    `gorm:"column:capacity;type:int(11);default:null" json:"capacity"`         // 股票容量
	Only20    bool   `gorm:"column:only20;type:tinyint(1);default:null" json:"only20"`          // 只显示20日线
}

// TableName get sql table name.获取数据库表名
func (m *WxUserinfo) TableName() string {
	return "wx_userinfo"
}

// WxUserinfoColumns get sql column name.获取数据库列名
var WxUserinfoColumns = struct {
	ID        string
	Openid    string
	NickName  string
	AvatarURL string
	Gender    string
	City      string
	Province  string
	Country   string
	Language  string
	Phone     string
	Group     string
	Rg        string
	Vip       string
	Capacity  string
	Only20    string
}{
	ID:        "id",
	Openid:    "openid",
	NickName:  "nick_name",
	AvatarURL: "avatar_url",
	Gender:    "gender",
	City:      "city",
	Province:  "province",
	Country:   "country",
	Language:  "language",
	Phone:     "phone",
	Group:     "group",
	Rg:        "rg",
	Vip:       "vip",
	Capacity:  "capacity",
	Only20:    "only20",
}

// XqMsgDailyActiveTbl 当日评论活跃股票
type XqMsgDailyActiveTbl struct {
	ID           int            `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Code         string         `gorm:"uniqueIndex:code;column:code;type:varchar(255);default:null" json:"code"` // 股票代码
	Name         string         `gorm:"column:name;type:varchar(255);default:null" json:"name"`                  // 股票名字
	Day0         datatypes.Date `gorm:"uniqueIndex:code;column:day0;type:date;default:null" json:"day0"`         // 入池时间
	DayStr       string         `gorm:"column:day_str;type:varchar(255);default:null" json:"dayStr"`             // 入池时间
	Price        float64        `gorm:"column:price;type:double(11,2);default:null" json:"price"`                // 价格
	Percent      float64        `gorm:"column:percent;type:double(16,2);default:null" json:"percent"`            // 百分比
	TurnoverRate float64        `gorm:"column:turnover_rate;type:double(16,2);default:null" json:"turnoverRate"` // 换手率
	CreatedAt    time.Time      `gorm:"column:created_at;type:datetime;default:null" json:"createdAt"`           // 创建时间
	Count        int            `gorm:"column:count;type:int(11);default:null" json:"count"`                     // 次数
}

// TableName get sql table name.获取数据库表名
func (m *XqMsgDailyActiveTbl) TableName() string {
	return "xq_msg_daily_active_tbl"
}

// XqMsgDailyActiveTblColumns get sql column name.获取数据库列名
var XqMsgDailyActiveTblColumns = struct {
	ID           string
	Code         string
	Name         string
	Day0         string
	DayStr       string
	Price        string
	Percent      string
	TurnoverRate string
	CreatedAt    string
	Count        string
}{
	ID:           "id",
	Code:         "code",
	Name:         "name",
	Day0:         "day0",
	DayStr:       "day_str",
	Price:        "price",
	Percent:      "percent",
	TurnoverRate: "turnover_rate",
	CreatedAt:    "created_at",
	Count:        "count",
}

// XqMsgDailyTbl 股评
type XqMsgDailyTbl struct {
	ID            int       `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Code          string    `gorm:"uniqueIndex:code;column:code;type:varchar(255);default:null" json:"code"` // 股票代码
	Name          string    `gorm:"column:name;type:varchar(255);default:null" json:"name"`                  // 股票名字
	Day0          int64     `gorm:"uniqueIndex:code;column:day0;type:bigint(20);default:null" json:"day0"`   // 入池时间
	Day0Str       string    `gorm:"column:day0_str;type:varchar(255);default:null" json:"day0Str"`           // 入池时间
	Tag           string    `gorm:"column:tag;type:text;default:null" json:"tag"`                            // 分词关键字
	GoodSentiment int       `gorm:"column:good_sentiment;type:int(11);default:null" json:"goodSentiment"`    // 正面情绪数量
	BadSentiment  int       `gorm:"column:bad_sentiment;type:int(11);default:null" json:"badSentiment"`      // 负面情绪数量
	Sentiment     int       `gorm:"column:sentiment;type:int(11);default:null" json:"sentiment"`             // 中性情绪数量
	Price         float64   `gorm:"column:price;type:double(11,2);default:null" json:"price"`                // 价格
	Percent       float64   `gorm:"column:percent;type:double(16,2);default:null" json:"percent"`            // 百分比
	TurnoverRate  float64   `gorm:"column:turnover_rate;type:double(16,2);default:null" json:"turnoverRate"` // 换手率
	CreatedAt     time.Time `gorm:"column:created_at;type:datetime;default:null" json:"createdAt"`           // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *XqMsgDailyTbl) TableName() string {
	return "xq_msg_daily_tbl"
}

// XqMsgDailyTblColumns get sql column name.获取数据库列名
var XqMsgDailyTblColumns = struct {
	ID            string
	Code          string
	Name          string
	Day0          string
	Day0Str       string
	Tag           string
	GoodSentiment string
	BadSentiment  string
	Sentiment     string
	Price         string
	Percent       string
	TurnoverRate  string
	CreatedAt     string
}{
	ID:            "id",
	Code:          "code",
	Name:          "name",
	Day0:          "day0",
	Day0Str:       "day0_str",
	Tag:           "tag",
	GoodSentiment: "good_sentiment",
	BadSentiment:  "bad_sentiment",
	Sentiment:     "sentiment",
	Price:         "price",
	Percent:       "percent",
	TurnoverRate:  "turnover_rate",
	CreatedAt:     "created_at",
}

// XqMsgTbl 股评
type XqMsgTbl struct {
	ID           int       `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Code         string    `gorm:"uniqueIndex:code;column:code;type:varchar(255);default:null" json:"code"` // 股票代码
	Tm           int64     `gorm:"uniqueIndex:code;column:tm;type:bigint(20);default:null" json:"tm"`       // 评论时间
	Name         string    `gorm:"column:name;type:varchar(255);default:null" json:"name"`                  // 股票名字
	Day          time.Time `gorm:"column:day;type:datetime;default:null" json:"day"`                        // 入池时间
	DayStr       string    `gorm:"column:day_str;type:varchar(255);default:null" json:"dayStr"`             // 入池时间
	Msg          string    `gorm:"column:msg;type:text;default:null" json:"msg"`                            // 消息内容
	Lexer        string    `gorm:"column:lexer;type:text;default:null" json:"lexer"`                        // 分词结果
	Tag          string    `gorm:"column:tag;type:text;default:null" json:"tag"`                            // 分词关键字
	Sentiment    int       `gorm:"column:sentiment;type:int(11);default:null" json:"sentiment"`             // 情绪(0:负向，1:中性，2:正向)
	Price        float64   `gorm:"column:price;type:double(11,2);default:null" json:"price"`                // 价格
	Percent      float64   `gorm:"column:percent;type:double(16,2);default:null" json:"percent"`            // 百分比
	TurnoverRate float64   `gorm:"column:turnover_rate;type:double(16,2);default:null" json:"turnoverRate"` // 换手率
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime;default:null" json:"createdAt"`           // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *XqMsgTbl) TableName() string {
	return "xq_msg_tbl"
}

// XqMsgTblColumns get sql column name.获取数据库列名
var XqMsgTblColumns = struct {
	ID           string
	Code         string
	Tm           string
	Name         string
	Day          string
	DayStr       string
	Msg          string
	Lexer        string
	Tag          string
	Sentiment    string
	Price        string
	Percent      string
	TurnoverRate string
	CreatedAt    string
}{
	ID:           "id",
	Code:         "code",
	Tm:           "tm",
	Name:         "name",
	Day:          "day",
	DayStr:       "day_str",
	Msg:          "msg",
	Lexer:        "lexer",
	Tag:          "tag",
	Sentiment:    "sentiment",
	Price:        "price",
	Percent:      "percent",
	TurnoverRate: "turnover_rate",
	CreatedAt:    "created_at",
}

// ZljlrDailyTbl 主力净流入列表
type ZljlrDailyTbl struct {
	ID        int       `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Code      string    `gorm:"uniqueIndex:code;column:code;type:varchar(255);default:null" json:"code"` // 股票代码
	Day0      int64     `gorm:"uniqueIndex:code;column:day0;type:bigint(11);default:null" json:"day0"`   // 当日0点时间戳
	DayStr    string    `gorm:"column:day_str;type:varchar(255);default:null" json:"dayStr"`             // 入池时间
	Price     float64   `gorm:"column:price;type:double(16,4);default:null" json:"price"`                // 当前价格(万元)
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;default:null" json:"createdAt"`           // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *ZljlrDailyTbl) TableName() string {
	return "zljlr_daily_tbl"
}

// ZljlrDailyTblColumns get sql column name.获取数据库列名
var ZljlrDailyTblColumns = struct {
	ID        string
	Code      string
	Day0      string
	DayStr    string
	Price     string
	CreatedAt string
}{
	ID:        "id",
	Code:      "code",
	Day0:      "day0",
	DayStr:    "day_str",
	Price:     "price",
	CreatedAt: "created_at",
}
