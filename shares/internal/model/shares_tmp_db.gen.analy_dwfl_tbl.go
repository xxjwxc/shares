package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type _AnalyDwflTblMgr struct {
	*_BaseMgr
}

// AnalyDwflTblMgr open func
func AnalyDwflTblMgr(db *gorm.DB) *_AnalyDwflTblMgr {
	if db == nil {
		panic(fmt.Errorf("AnalyDwflTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AnalyDwflTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("analy_dwfl_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AnalyDwflTblMgr) GetTableName() string {
	return "analy_dwfl_tbl"
}

// Reset 重置gorm会话
func (obj *_AnalyDwflTblMgr) Reset() *_AnalyDwflTblMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AnalyDwflTblMgr) Get() (result AnalyDwflTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyDwflTbl{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AnalyDwflTblMgr) Gets() (results []*AnalyDwflTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyDwflTbl{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AnalyDwflTblMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AnalyDwflTbl{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_AnalyDwflTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取 股票代码
func (obj *_AnalyDwflTblMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithName name获取 股票名字
func (obj *_AnalyDwflTblMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDay day获取 入池时间
func (obj *_AnalyDwflTblMgr) WithDay(day datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["day"] = day })
}

// WithDayStr day_str获取 入池时间
func (obj *_AnalyDwflTblMgr) WithDayStr(dayStr string) Option {
	return optionFunc(func(o *options) { o.query["day_str"] = dayStr })
}

// WithPrice price获取 价格
func (obj *_AnalyDwflTblMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithNewPrice new_price获取 价格
func (obj *_AnalyDwflTblMgr) WithNewPrice(newPrice float64) Option {
	return optionFunc(func(o *options) { o.query["new_price"] = newPrice })
}

// WithPercent percent获取 百分比
func (obj *_AnalyDwflTblMgr) WithPercent(percent float64) Option {
	return optionFunc(func(o *options) { o.query["percent"] = percent })
}

// WithTurnoverRate turnover_rate获取 换手率
func (obj *_AnalyDwflTblMgr) WithTurnoverRate(turnoverRate float64) Option {
	return optionFunc(func(o *options) { o.query["turnover_rate"] = turnoverRate })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_AnalyDwflTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// GetByOption 功能选项模式获取
func (obj *_AnalyDwflTblMgr) GetByOption(opts ...Option) (result AnalyDwflTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AnalyDwflTbl{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AnalyDwflTblMgr) GetByOptions(opts ...Option) (results []*AnalyDwflTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AnalyDwflTbl{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AnalyDwflTblMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AnalyDwflTbl, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AnalyDwflTbl{}).Where(options.query)
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
func (obj *_AnalyDwflTblMgr) GetFromID(id int) (result AnalyDwflTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyDwflTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_AnalyDwflTblMgr) GetBatchFromID(ids []int) (results []*AnalyDwflTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyDwflTbl{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 股票代码
func (obj *_AnalyDwflTblMgr) GetFromCode(code string) (result AnalyDwflTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyDwflTbl{}).Where("`code` = ?", code).First(&result).Error

	return
}

// GetBatchFromCode 批量查找 股票代码
func (obj *_AnalyDwflTblMgr) GetBatchFromCode(codes []string) (results []*AnalyDwflTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyDwflTbl{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 股票名字
func (obj *_AnalyDwflTblMgr) GetFromName(name string) (results []*AnalyDwflTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyDwflTbl{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 股票名字
func (obj *_AnalyDwflTblMgr) GetBatchFromName(names []string) (results []*AnalyDwflTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyDwflTbl{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromDay 通过day获取内容 入池时间
func (obj *_AnalyDwflTblMgr) GetFromDay(day datatypes.Date) (results []*AnalyDwflTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyDwflTbl{}).Where("`day` = ?", day).Find(&results).Error

	return
}

// GetBatchFromDay 批量查找 入池时间
func (obj *_AnalyDwflTblMgr) GetBatchFromDay(days []datatypes.Date) (results []*AnalyDwflTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyDwflTbl{}).Where("`day` IN (?)", days).Find(&results).Error

	return
}

// GetFromDayStr 通过day_str获取内容 入池时间
func (obj *_AnalyDwflTblMgr) GetFromDayStr(dayStr string) (results []*AnalyDwflTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyDwflTbl{}).Where("`day_str` = ?", dayStr).Find(&results).Error

	return
}

// GetBatchFromDayStr 批量查找 入池时间
func (obj *_AnalyDwflTblMgr) GetBatchFromDayStr(dayStrs []string) (results []*AnalyDwflTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyDwflTbl{}).Where("`day_str` IN (?)", dayStrs).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 价格
func (obj *_AnalyDwflTblMgr) GetFromPrice(price float64) (results []*AnalyDwflTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyDwflTbl{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 价格
func (obj *_AnalyDwflTblMgr) GetBatchFromPrice(prices []float64) (results []*AnalyDwflTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyDwflTbl{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromNewPrice 通过new_price获取内容 价格
func (obj *_AnalyDwflTblMgr) GetFromNewPrice(newPrice float64) (results []*AnalyDwflTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyDwflTbl{}).Where("`new_price` = ?", newPrice).Find(&results).Error

	return
}

// GetBatchFromNewPrice 批量查找 价格
func (obj *_AnalyDwflTblMgr) GetBatchFromNewPrice(newPrices []float64) (results []*AnalyDwflTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyDwflTbl{}).Where("`new_price` IN (?)", newPrices).Find(&results).Error

	return
}

// GetFromPercent 通过percent获取内容 百分比
func (obj *_AnalyDwflTblMgr) GetFromPercent(percent float64) (results []*AnalyDwflTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyDwflTbl{}).Where("`percent` = ?", percent).Find(&results).Error

	return
}

// GetBatchFromPercent 批量查找 百分比
func (obj *_AnalyDwflTblMgr) GetBatchFromPercent(percents []float64) (results []*AnalyDwflTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyDwflTbl{}).Where("`percent` IN (?)", percents).Find(&results).Error

	return
}

// GetFromTurnoverRate 通过turnover_rate获取内容 换手率
func (obj *_AnalyDwflTblMgr) GetFromTurnoverRate(turnoverRate float64) (results []*AnalyDwflTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyDwflTbl{}).Where("`turnover_rate` = ?", turnoverRate).Find(&results).Error

	return
}

// GetBatchFromTurnoverRate 批量查找 换手率
func (obj *_AnalyDwflTblMgr) GetBatchFromTurnoverRate(turnoverRates []float64) (results []*AnalyDwflTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyDwflTbl{}).Where("`turnover_rate` IN (?)", turnoverRates).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_AnalyDwflTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*AnalyDwflTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyDwflTbl{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_AnalyDwflTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*AnalyDwflTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyDwflTbl{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AnalyDwflTblMgr) FetchByPrimaryKey(id int) (result AnalyDwflTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyDwflTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByCode primary or index 获取唯一内容
func (obj *_AnalyDwflTblMgr) FetchUniqueByCode(code string) (result AnalyDwflTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyDwflTbl{}).Where("`code` = ?", code).First(&result).Error

	return
}
