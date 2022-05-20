package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type _AnalyFlTblMgr struct {
	*_BaseMgr
}

// AnalyFlTblMgr open func
func AnalyFlTblMgr(db *gorm.DB) *_AnalyFlTblMgr {
	if db == nil {
		panic(fmt.Errorf("AnalyFlTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AnalyFlTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("analy_fl_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AnalyFlTblMgr) GetTableName() string {
	return "analy_fl_tbl"
}

// Reset 重置gorm会话
func (obj *_AnalyFlTblMgr) Reset() *_AnalyFlTblMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AnalyFlTblMgr) Get() (result AnalyFlTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlTbl{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AnalyFlTblMgr) Gets() (results []*AnalyFlTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlTbl{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AnalyFlTblMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AnalyFlTbl{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_AnalyFlTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取 股票代码
func (obj *_AnalyFlTblMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithName name获取 股票名字
func (obj *_AnalyFlTblMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDay day获取 入池时间
func (obj *_AnalyFlTblMgr) WithDay(day datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["day"] = day })
}

// WithDayStr day_str获取 入池时间
func (obj *_AnalyFlTblMgr) WithDayStr(dayStr string) Option {
	return optionFunc(func(o *options) { o.query["day_str"] = dayStr })
}

// WithPrice price获取 价格
func (obj *_AnalyFlTblMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithNewPrice new_price获取 价格
func (obj *_AnalyFlTblMgr) WithNewPrice(newPrice float64) Option {
	return optionFunc(func(o *options) { o.query["new_price"] = newPrice })
}

// WithPercent percent获取 百分比
func (obj *_AnalyFlTblMgr) WithPercent(percent float64) Option {
	return optionFunc(func(o *options) { o.query["percent"] = percent })
}

// WithTurnoverRate turnover_rate获取 换手率
func (obj *_AnalyFlTblMgr) WithTurnoverRate(turnoverRate float64) Option {
	return optionFunc(func(o *options) { o.query["turnover_rate"] = turnoverRate })
}

// WithScore score获取 得分
func (obj *_AnalyFlTblMgr) WithScore(score int) Option {
	return optionFunc(func(o *options) { o.query["score"] = score })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_AnalyFlTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// GetByOption 功能选项模式获取
func (obj *_AnalyFlTblMgr) GetByOption(opts ...Option) (result AnalyFlTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlTbl{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AnalyFlTblMgr) GetByOptions(opts ...Option) (results []*AnalyFlTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlTbl{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AnalyFlTblMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AnalyFlTbl, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AnalyFlTbl{}).Where(options.query)
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
func (obj *_AnalyFlTblMgr) GetFromID(id int) (result AnalyFlTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_AnalyFlTblMgr) GetBatchFromID(ids []int) (results []*AnalyFlTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlTbl{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 股票代码
func (obj *_AnalyFlTblMgr) GetFromCode(code string) (result AnalyFlTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlTbl{}).Where("`code` = ?", code).First(&result).Error

	return
}

// GetBatchFromCode 批量查找 股票代码
func (obj *_AnalyFlTblMgr) GetBatchFromCode(codes []string) (results []*AnalyFlTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlTbl{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 股票名字
func (obj *_AnalyFlTblMgr) GetFromName(name string) (results []*AnalyFlTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlTbl{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 股票名字
func (obj *_AnalyFlTblMgr) GetBatchFromName(names []string) (results []*AnalyFlTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlTbl{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromDay 通过day获取内容 入池时间
func (obj *_AnalyFlTblMgr) GetFromDay(day datatypes.Date) (results []*AnalyFlTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlTbl{}).Where("`day` = ?", day).Find(&results).Error

	return
}

// GetBatchFromDay 批量查找 入池时间
func (obj *_AnalyFlTblMgr) GetBatchFromDay(days []datatypes.Date) (results []*AnalyFlTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlTbl{}).Where("`day` IN (?)", days).Find(&results).Error

	return
}

// GetFromDayStr 通过day_str获取内容 入池时间
func (obj *_AnalyFlTblMgr) GetFromDayStr(dayStr string) (results []*AnalyFlTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlTbl{}).Where("`day_str` = ?", dayStr).Find(&results).Error

	return
}

// GetBatchFromDayStr 批量查找 入池时间
func (obj *_AnalyFlTblMgr) GetBatchFromDayStr(dayStrs []string) (results []*AnalyFlTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlTbl{}).Where("`day_str` IN (?)", dayStrs).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 价格
func (obj *_AnalyFlTblMgr) GetFromPrice(price float64) (results []*AnalyFlTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlTbl{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 价格
func (obj *_AnalyFlTblMgr) GetBatchFromPrice(prices []float64) (results []*AnalyFlTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlTbl{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromNewPrice 通过new_price获取内容 价格
func (obj *_AnalyFlTblMgr) GetFromNewPrice(newPrice float64) (results []*AnalyFlTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlTbl{}).Where("`new_price` = ?", newPrice).Find(&results).Error

	return
}

// GetBatchFromNewPrice 批量查找 价格
func (obj *_AnalyFlTblMgr) GetBatchFromNewPrice(newPrices []float64) (results []*AnalyFlTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlTbl{}).Where("`new_price` IN (?)", newPrices).Find(&results).Error

	return
}

// GetFromPercent 通过percent获取内容 百分比
func (obj *_AnalyFlTblMgr) GetFromPercent(percent float64) (results []*AnalyFlTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlTbl{}).Where("`percent` = ?", percent).Find(&results).Error

	return
}

// GetBatchFromPercent 批量查找 百分比
func (obj *_AnalyFlTblMgr) GetBatchFromPercent(percents []float64) (results []*AnalyFlTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlTbl{}).Where("`percent` IN (?)", percents).Find(&results).Error

	return
}

// GetFromTurnoverRate 通过turnover_rate获取内容 换手率
func (obj *_AnalyFlTblMgr) GetFromTurnoverRate(turnoverRate float64) (results []*AnalyFlTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlTbl{}).Where("`turnover_rate` = ?", turnoverRate).Find(&results).Error

	return
}

// GetBatchFromTurnoverRate 批量查找 换手率
func (obj *_AnalyFlTblMgr) GetBatchFromTurnoverRate(turnoverRates []float64) (results []*AnalyFlTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlTbl{}).Where("`turnover_rate` IN (?)", turnoverRates).Find(&results).Error

	return
}

// GetFromScore 通过score获取内容 得分
func (obj *_AnalyFlTblMgr) GetFromScore(score int) (results []*AnalyFlTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlTbl{}).Where("`score` = ?", score).Find(&results).Error

	return
}

// GetBatchFromScore 批量查找 得分
func (obj *_AnalyFlTblMgr) GetBatchFromScore(scores []int) (results []*AnalyFlTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlTbl{}).Where("`score` IN (?)", scores).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_AnalyFlTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*AnalyFlTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlTbl{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_AnalyFlTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*AnalyFlTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlTbl{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AnalyFlTblMgr) FetchByPrimaryKey(id int) (result AnalyFlTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByCode primary or index 获取唯一内容
func (obj *_AnalyFlTblMgr) FetchUniqueByCode(code string) (result AnalyFlTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlTbl{}).Where("`code` = ?", code).First(&result).Error

	return
}
