package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _HyDailyTblMgr struct {
	*_BaseMgr
}

// HyDailyTblMgr open func
func HyDailyTblMgr(db *gorm.DB) *_HyDailyTblMgr {
	if db == nil {
		panic(fmt.Errorf("HyDailyTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_HyDailyTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("hy_daily_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_HyDailyTblMgr) GetTableName() string {
	return "hy_daily_tbl"
}

// Reset 重置gorm会话
func (obj *_HyDailyTblMgr) Reset() *_HyDailyTblMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_HyDailyTblMgr) Get() (result HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_HyDailyTblMgr) Gets() (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_HyDailyTblMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_HyDailyTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithHyCode hy_code获取 股票代码
func (obj *_HyDailyTblMgr) WithHyCode(hyCode string) Option {
	return optionFunc(func(o *options) { o.query["hy_code"] = hyCode })
}

// WithHyName hy_name获取 行业名
func (obj *_HyDailyTblMgr) WithHyName(hyName string) Option {
	return optionFunc(func(o *options) { o.query["hy_name"] = hyName })
}

// WithPrice price获取 当前价格
func (obj *_HyDailyTblMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithPercent percent获取 百分比
func (obj *_HyDailyTblMgr) WithPercent(percent float64) Option {
	return optionFunc(func(o *options) { o.query["percent"] = percent })
}

// WithMa5 ma5获取 5日均线
func (obj *_HyDailyTblMgr) WithMa5(ma5 float64) Option {
	return optionFunc(func(o *options) { o.query["ma5"] = ma5 })
}

// WithMa10 ma10获取 10日均线
func (obj *_HyDailyTblMgr) WithMa10(ma10 float64) Option {
	return optionFunc(func(o *options) { o.query["ma10"] = ma10 })
}

// WithMa20 ma20获取 20日均线
func (obj *_HyDailyTblMgr) WithMa20(ma20 float64) Option {
	return optionFunc(func(o *options) { o.query["ma20"] = ma20 })
}

// WithMa60 ma60获取 60日均线
func (obj *_HyDailyTblMgr) WithMa60(ma60 float64) Option {
	return optionFunc(func(o *options) { o.query["ma60"] = ma60 })
}

// WithDay0 day0获取 当日0点时间戳
func (obj *_HyDailyTblMgr) WithDay0(day0 int64) Option {
	return optionFunc(func(o *options) { o.query["day0"] = day0 })
}

// WithDay0Str day0_str获取 当日0点时间戳
func (obj *_HyDailyTblMgr) WithDay0Str(day0Str string) Option {
	return optionFunc(func(o *options) { o.query["day0_str"] = day0Str })
}

// WithVolume volume获取 成交量（手）
func (obj *_HyDailyTblMgr) WithVolume(volume float64) Option {
	return optionFunc(func(o *options) { o.query["volume"] = volume })
}

// WithTurnover turnover获取 成交额（万）
func (obj *_HyDailyTblMgr) WithTurnover(turnover float64) Option {
	return optionFunc(func(o *options) { o.query["turnover"] = turnover })
}

// WithTurnoverRate turnover_rate获取 换手率
func (obj *_HyDailyTblMgr) WithTurnoverRate(turnoverRate float64) Option {
	return optionFunc(func(o *options) { o.query["turnover_rate"] = turnoverRate })
}

// WithZljlr zljlr获取 主力净流入
func (obj *_HyDailyTblMgr) WithZljlr(zljlr float64) Option {
	return optionFunc(func(o *options) { o.query["zljlr"] = zljlr })
}

// WithOpenPrice open_price获取 开盘价
func (obj *_HyDailyTblMgr) WithOpenPrice(openPrice float64) Option {
	return optionFunc(func(o *options) { o.query["open_price"] = openPrice })
}

// WithClosePrice close_price获取 收盘价
func (obj *_HyDailyTblMgr) WithClosePrice(closePrice float64) Option {
	return optionFunc(func(o *options) { o.query["close_price"] = closePrice })
}

// WithMax max获取 当日最高值
func (obj *_HyDailyTblMgr) WithMax(max float64) Option {
	return optionFunc(func(o *options) { o.query["max"] = max })
}

// WithMin min获取 当日最低价
func (obj *_HyDailyTblMgr) WithMin(min float64) Option {
	return optionFunc(func(o *options) { o.query["min"] = min })
}

// WithPe pe获取 市盈率
func (obj *_HyDailyTblMgr) WithPe(pe float64) Option {
	return optionFunc(func(o *options) { o.query["pe"] = pe })
}

// WithPeg peg获取 PEG值
func (obj *_HyDailyTblMgr) WithPeg(peg float64) Option {
	return optionFunc(func(o *options) { o.query["peg"] = peg })
}

// WithPb pb获取 市净率
func (obj *_HyDailyTblMgr) WithPb(pb float64) Option {
	return optionFunc(func(o *options) { o.query["pb"] = pb })
}

// WithCapVag cap_vag获取 平均市值
func (obj *_HyDailyTblMgr) WithCapVag(capVag float64) Option {
	return optionFunc(func(o *options) { o.query["cap_vag"] = capVag })
}

// WithNum num获取 家数
func (obj *_HyDailyTblMgr) WithNum(num int) Option {
	return optionFunc(func(o *options) { o.query["num"] = num })
}

// WithLossNum loss_num获取 亏损家数
func (obj *_HyDailyTblMgr) WithLossNum(lossNum int) Option {
	return optionFunc(func(o *options) { o.query["loss_num"] = lossNum })
}

// WithCreatedBy created_by获取 创建者
func (obj *_HyDailyTblMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_HyDailyTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// GetByOption 功能选项模式获取
func (obj *_HyDailyTblMgr) GetByOption(opts ...Option) (result HyDailyTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_HyDailyTblMgr) GetByOptions(opts ...Option) (results []*HyDailyTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_HyDailyTblMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]HyDailyTbl, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_HyDailyTblMgr) GetFromID(id int) (result HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_HyDailyTblMgr) GetBatchFromID(ids []int) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromHyCode 通过hy_code获取内容 股票代码
func (obj *_HyDailyTblMgr) GetFromHyCode(hyCode string) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`hy_code` = ?", hyCode).Find(&results).Error

	return
}

// GetBatchFromHyCode 批量查找 股票代码
func (obj *_HyDailyTblMgr) GetBatchFromHyCode(hyCodes []string) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`hy_code` IN (?)", hyCodes).Find(&results).Error

	return
}

// GetFromHyName 通过hy_name获取内容 行业名
func (obj *_HyDailyTblMgr) GetFromHyName(hyName string) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`hy_name` = ?", hyName).Find(&results).Error

	return
}

// GetBatchFromHyName 批量查找 行业名
func (obj *_HyDailyTblMgr) GetBatchFromHyName(hyNames []string) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`hy_name` IN (?)", hyNames).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 当前价格
func (obj *_HyDailyTblMgr) GetFromPrice(price float64) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 当前价格
func (obj *_HyDailyTblMgr) GetBatchFromPrice(prices []float64) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromPercent 通过percent获取内容 百分比
func (obj *_HyDailyTblMgr) GetFromPercent(percent float64) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`percent` = ?", percent).Find(&results).Error

	return
}

// GetBatchFromPercent 批量查找 百分比
func (obj *_HyDailyTblMgr) GetBatchFromPercent(percents []float64) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`percent` IN (?)", percents).Find(&results).Error

	return
}

// GetFromMa5 通过ma5获取内容 5日均线
func (obj *_HyDailyTblMgr) GetFromMa5(ma5 float64) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`ma5` = ?", ma5).Find(&results).Error

	return
}

// GetBatchFromMa5 批量查找 5日均线
func (obj *_HyDailyTblMgr) GetBatchFromMa5(ma5s []float64) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`ma5` IN (?)", ma5s).Find(&results).Error

	return
}

// GetFromMa10 通过ma10获取内容 10日均线
func (obj *_HyDailyTblMgr) GetFromMa10(ma10 float64) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`ma10` = ?", ma10).Find(&results).Error

	return
}

// GetBatchFromMa10 批量查找 10日均线
func (obj *_HyDailyTblMgr) GetBatchFromMa10(ma10s []float64) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`ma10` IN (?)", ma10s).Find(&results).Error

	return
}

// GetFromMa20 通过ma20获取内容 20日均线
func (obj *_HyDailyTblMgr) GetFromMa20(ma20 float64) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`ma20` = ?", ma20).Find(&results).Error

	return
}

// GetBatchFromMa20 批量查找 20日均线
func (obj *_HyDailyTblMgr) GetBatchFromMa20(ma20s []float64) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`ma20` IN (?)", ma20s).Find(&results).Error

	return
}

// GetFromMa60 通过ma60获取内容 60日均线
func (obj *_HyDailyTblMgr) GetFromMa60(ma60 float64) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`ma60` = ?", ma60).Find(&results).Error

	return
}

// GetBatchFromMa60 批量查找 60日均线
func (obj *_HyDailyTblMgr) GetBatchFromMa60(ma60s []float64) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`ma60` IN (?)", ma60s).Find(&results).Error

	return
}

// GetFromDay0 通过day0获取内容 当日0点时间戳
func (obj *_HyDailyTblMgr) GetFromDay0(day0 int64) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`day0` = ?", day0).Find(&results).Error

	return
}

// GetBatchFromDay0 批量查找 当日0点时间戳
func (obj *_HyDailyTblMgr) GetBatchFromDay0(day0s []int64) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`day0` IN (?)", day0s).Find(&results).Error

	return
}

// GetFromDay0Str 通过day0_str获取内容 当日0点时间戳
func (obj *_HyDailyTblMgr) GetFromDay0Str(day0Str string) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`day0_str` = ?", day0Str).Find(&results).Error

	return
}

// GetBatchFromDay0Str 批量查找 当日0点时间戳
func (obj *_HyDailyTblMgr) GetBatchFromDay0Str(day0Strs []string) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`day0_str` IN (?)", day0Strs).Find(&results).Error

	return
}

// GetFromVolume 通过volume获取内容 成交量（手）
func (obj *_HyDailyTblMgr) GetFromVolume(volume float64) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`volume` = ?", volume).Find(&results).Error

	return
}

// GetBatchFromVolume 批量查找 成交量（手）
func (obj *_HyDailyTblMgr) GetBatchFromVolume(volumes []float64) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`volume` IN (?)", volumes).Find(&results).Error

	return
}

// GetFromTurnover 通过turnover获取内容 成交额（万）
func (obj *_HyDailyTblMgr) GetFromTurnover(turnover float64) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`turnover` = ?", turnover).Find(&results).Error

	return
}

// GetBatchFromTurnover 批量查找 成交额（万）
func (obj *_HyDailyTblMgr) GetBatchFromTurnover(turnovers []float64) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`turnover` IN (?)", turnovers).Find(&results).Error

	return
}

// GetFromTurnoverRate 通过turnover_rate获取内容 换手率
func (obj *_HyDailyTblMgr) GetFromTurnoverRate(turnoverRate float64) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`turnover_rate` = ?", turnoverRate).Find(&results).Error

	return
}

// GetBatchFromTurnoverRate 批量查找 换手率
func (obj *_HyDailyTblMgr) GetBatchFromTurnoverRate(turnoverRates []float64) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`turnover_rate` IN (?)", turnoverRates).Find(&results).Error

	return
}

// GetFromZljlr 通过zljlr获取内容 主力净流入
func (obj *_HyDailyTblMgr) GetFromZljlr(zljlr float64) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`zljlr` = ?", zljlr).Find(&results).Error

	return
}

// GetBatchFromZljlr 批量查找 主力净流入
func (obj *_HyDailyTblMgr) GetBatchFromZljlr(zljlrs []float64) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`zljlr` IN (?)", zljlrs).Find(&results).Error

	return
}

// GetFromOpenPrice 通过open_price获取内容 开盘价
func (obj *_HyDailyTblMgr) GetFromOpenPrice(openPrice float64) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`open_price` = ?", openPrice).Find(&results).Error

	return
}

// GetBatchFromOpenPrice 批量查找 开盘价
func (obj *_HyDailyTblMgr) GetBatchFromOpenPrice(openPrices []float64) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`open_price` IN (?)", openPrices).Find(&results).Error

	return
}

// GetFromClosePrice 通过close_price获取内容 收盘价
func (obj *_HyDailyTblMgr) GetFromClosePrice(closePrice float64) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`close_price` = ?", closePrice).Find(&results).Error

	return
}

// GetBatchFromClosePrice 批量查找 收盘价
func (obj *_HyDailyTblMgr) GetBatchFromClosePrice(closePrices []float64) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`close_price` IN (?)", closePrices).Find(&results).Error

	return
}

// GetFromMax 通过max获取内容 当日最高值
func (obj *_HyDailyTblMgr) GetFromMax(max float64) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`max` = ?", max).Find(&results).Error

	return
}

// GetBatchFromMax 批量查找 当日最高值
func (obj *_HyDailyTblMgr) GetBatchFromMax(maxs []float64) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`max` IN (?)", maxs).Find(&results).Error

	return
}

// GetFromMin 通过min获取内容 当日最低价
func (obj *_HyDailyTblMgr) GetFromMin(min float64) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`min` = ?", min).Find(&results).Error

	return
}

// GetBatchFromMin 批量查找 当日最低价
func (obj *_HyDailyTblMgr) GetBatchFromMin(mins []float64) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`min` IN (?)", mins).Find(&results).Error

	return
}

// GetFromPe 通过pe获取内容 市盈率
func (obj *_HyDailyTblMgr) GetFromPe(pe float64) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`pe` = ?", pe).Find(&results).Error

	return
}

// GetBatchFromPe 批量查找 市盈率
func (obj *_HyDailyTblMgr) GetBatchFromPe(pes []float64) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`pe` IN (?)", pes).Find(&results).Error

	return
}

// GetFromPeg 通过peg获取内容 PEG值
func (obj *_HyDailyTblMgr) GetFromPeg(peg float64) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`peg` = ?", peg).Find(&results).Error

	return
}

// GetBatchFromPeg 批量查找 PEG值
func (obj *_HyDailyTblMgr) GetBatchFromPeg(pegs []float64) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`peg` IN (?)", pegs).Find(&results).Error

	return
}

// GetFromPb 通过pb获取内容 市净率
func (obj *_HyDailyTblMgr) GetFromPb(pb float64) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`pb` = ?", pb).Find(&results).Error

	return
}

// GetBatchFromPb 批量查找 市净率
func (obj *_HyDailyTblMgr) GetBatchFromPb(pbs []float64) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`pb` IN (?)", pbs).Find(&results).Error

	return
}

// GetFromCapVag 通过cap_vag获取内容 平均市值
func (obj *_HyDailyTblMgr) GetFromCapVag(capVag float64) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`cap_vag` = ?", capVag).Find(&results).Error

	return
}

// GetBatchFromCapVag 批量查找 平均市值
func (obj *_HyDailyTblMgr) GetBatchFromCapVag(capVags []float64) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`cap_vag` IN (?)", capVags).Find(&results).Error

	return
}

// GetFromNum 通过num获取内容 家数
func (obj *_HyDailyTblMgr) GetFromNum(num int) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`num` = ?", num).Find(&results).Error

	return
}

// GetBatchFromNum 批量查找 家数
func (obj *_HyDailyTblMgr) GetBatchFromNum(nums []int) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`num` IN (?)", nums).Find(&results).Error

	return
}

// GetFromLossNum 通过loss_num获取内容 亏损家数
func (obj *_HyDailyTblMgr) GetFromLossNum(lossNum int) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`loss_num` = ?", lossNum).Find(&results).Error

	return
}

// GetBatchFromLossNum 批量查找 亏损家数
func (obj *_HyDailyTblMgr) GetBatchFromLossNum(lossNums []int) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`loss_num` IN (?)", lossNums).Find(&results).Error

	return
}

// GetFromCreatedBy 通过created_by获取内容 创建者
func (obj *_HyDailyTblMgr) GetFromCreatedBy(createdBy string) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`created_by` = ?", createdBy).Find(&results).Error

	return
}

// GetBatchFromCreatedBy 批量查找 创建者
func (obj *_HyDailyTblMgr) GetBatchFromCreatedBy(createdBys []string) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`created_by` IN (?)", createdBys).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_HyDailyTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_HyDailyTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_HyDailyTblMgr) FetchByPrimaryKey(id int) (result HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueIndexByHyCode primary or index 获取唯一内容
func (obj *_HyDailyTblMgr) FetchUniqueIndexByHyCode(hyCode string, day0 int64) (result HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`hy_code` = ? AND `day0` = ?", hyCode, day0).First(&result).Error

	return
}

// FetchIndexByDay0  获取多个内容
func (obj *_HyDailyTblMgr) FetchIndexByDay0(day0 int64) (results []*HyDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyDailyTbl{}).Where("`day0` = ?", day0).Find(&results).Error

	return
}
