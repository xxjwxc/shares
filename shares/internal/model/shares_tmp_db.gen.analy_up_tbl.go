package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AnalyUpTblMgr struct {
	*_BaseMgr
}

// AnalyUpTblMgr open func
func AnalyUpTblMgr(db *gorm.DB) *_AnalyUpTblMgr {
	if db == nil {
		panic(fmt.Errorf("AnalyUpTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AnalyUpTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("analy_up_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AnalyUpTblMgr) GetTableName() string {
	return "analy_up_tbl"
}

// Reset 重置gorm会话
func (obj *_AnalyUpTblMgr) Reset() *_AnalyUpTblMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AnalyUpTblMgr) Get() (result AnalyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyUpTbl{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AnalyUpTblMgr) Gets() (results []*AnalyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyUpTbl{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AnalyUpTblMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AnalyUpTbl{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_AnalyUpTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取 股票代码
func (obj *_AnalyUpTblMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithName name获取 股票名字
func (obj *_AnalyUpTblMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDay0 day0获取 入池时间
func (obj *_AnalyUpTblMgr) WithDay0(day0 int64) Option {
	return optionFunc(func(o *options) { o.query["day0"] = day0 })
}

// WithDayStr day_str获取 入池时间
func (obj *_AnalyUpTblMgr) WithDayStr(dayStr string) Option {
	return optionFunc(func(o *options) { o.query["day_str"] = dayStr })
}

// WithPrice price获取 价格
func (obj *_AnalyUpTblMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithNewPrice new_price获取 价格
func (obj *_AnalyUpTblMgr) WithNewPrice(newPrice float64) Option {
	return optionFunc(func(o *options) { o.query["new_price"] = newPrice })
}

// WithPercent percent获取 百分比
func (obj *_AnalyUpTblMgr) WithPercent(percent float64) Option {
	return optionFunc(func(o *options) { o.query["percent"] = percent })
}

// WithTurnoverRate turnover_rate获取 换手率
func (obj *_AnalyUpTblMgr) WithTurnoverRate(turnoverRate float64) Option {
	return optionFunc(func(o *options) { o.query["turnover_rate"] = turnoverRate })
}

// WithDoc doc获取 描述
func (obj *_AnalyUpTblMgr) WithDoc(doc string) Option {
	return optionFunc(func(o *options) { o.query["doc"] = doc })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_AnalyUpTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// GetByOption 功能选项模式获取
func (obj *_AnalyUpTblMgr) GetByOption(opts ...Option) (result AnalyUpTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AnalyUpTbl{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AnalyUpTblMgr) GetByOptions(opts ...Option) (results []*AnalyUpTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AnalyUpTbl{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AnalyUpTblMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AnalyUpTbl, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AnalyUpTbl{}).Where(options.query)
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
func (obj *_AnalyUpTblMgr) GetFromID(id int) (result AnalyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyUpTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_AnalyUpTblMgr) GetBatchFromID(ids []int) (results []*AnalyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyUpTbl{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 股票代码
func (obj *_AnalyUpTblMgr) GetFromCode(code string) (result AnalyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyUpTbl{}).Where("`code` = ?", code).First(&result).Error

	return
}

// GetBatchFromCode 批量查找 股票代码
func (obj *_AnalyUpTblMgr) GetBatchFromCode(codes []string) (results []*AnalyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyUpTbl{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 股票名字
func (obj *_AnalyUpTblMgr) GetFromName(name string) (results []*AnalyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyUpTbl{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 股票名字
func (obj *_AnalyUpTblMgr) GetBatchFromName(names []string) (results []*AnalyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyUpTbl{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromDay0 通过day0获取内容 入池时间
func (obj *_AnalyUpTblMgr) GetFromDay0(day0 int64) (results []*AnalyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyUpTbl{}).Where("`day0` = ?", day0).Find(&results).Error

	return
}

// GetBatchFromDay0 批量查找 入池时间
func (obj *_AnalyUpTblMgr) GetBatchFromDay0(day0s []int64) (results []*AnalyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyUpTbl{}).Where("`day0` IN (?)", day0s).Find(&results).Error

	return
}

// GetFromDayStr 通过day_str获取内容 入池时间
func (obj *_AnalyUpTblMgr) GetFromDayStr(dayStr string) (results []*AnalyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyUpTbl{}).Where("`day_str` = ?", dayStr).Find(&results).Error

	return
}

// GetBatchFromDayStr 批量查找 入池时间
func (obj *_AnalyUpTblMgr) GetBatchFromDayStr(dayStrs []string) (results []*AnalyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyUpTbl{}).Where("`day_str` IN (?)", dayStrs).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 价格
func (obj *_AnalyUpTblMgr) GetFromPrice(price float64) (results []*AnalyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyUpTbl{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 价格
func (obj *_AnalyUpTblMgr) GetBatchFromPrice(prices []float64) (results []*AnalyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyUpTbl{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromNewPrice 通过new_price获取内容 价格
func (obj *_AnalyUpTblMgr) GetFromNewPrice(newPrice float64) (results []*AnalyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyUpTbl{}).Where("`new_price` = ?", newPrice).Find(&results).Error

	return
}

// GetBatchFromNewPrice 批量查找 价格
func (obj *_AnalyUpTblMgr) GetBatchFromNewPrice(newPrices []float64) (results []*AnalyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyUpTbl{}).Where("`new_price` IN (?)", newPrices).Find(&results).Error

	return
}

// GetFromPercent 通过percent获取内容 百分比
func (obj *_AnalyUpTblMgr) GetFromPercent(percent float64) (results []*AnalyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyUpTbl{}).Where("`percent` = ?", percent).Find(&results).Error

	return
}

// GetBatchFromPercent 批量查找 百分比
func (obj *_AnalyUpTblMgr) GetBatchFromPercent(percents []float64) (results []*AnalyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyUpTbl{}).Where("`percent` IN (?)", percents).Find(&results).Error

	return
}

// GetFromTurnoverRate 通过turnover_rate获取内容 换手率
func (obj *_AnalyUpTblMgr) GetFromTurnoverRate(turnoverRate float64) (results []*AnalyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyUpTbl{}).Where("`turnover_rate` = ?", turnoverRate).Find(&results).Error

	return
}

// GetBatchFromTurnoverRate 批量查找 换手率
func (obj *_AnalyUpTblMgr) GetBatchFromTurnoverRate(turnoverRates []float64) (results []*AnalyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyUpTbl{}).Where("`turnover_rate` IN (?)", turnoverRates).Find(&results).Error

	return
}

// GetFromDoc 通过doc获取内容 描述
func (obj *_AnalyUpTblMgr) GetFromDoc(doc string) (results []*AnalyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyUpTbl{}).Where("`doc` = ?", doc).Find(&results).Error

	return
}

// GetBatchFromDoc 批量查找 描述
func (obj *_AnalyUpTblMgr) GetBatchFromDoc(docs []string) (results []*AnalyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyUpTbl{}).Where("`doc` IN (?)", docs).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_AnalyUpTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*AnalyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyUpTbl{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_AnalyUpTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*AnalyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyUpTbl{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AnalyUpTblMgr) FetchByPrimaryKey(id int) (result AnalyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyUpTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByCode primary or index 获取唯一内容
func (obj *_AnalyUpTblMgr) FetchUniqueByCode(code string) (result AnalyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyUpTbl{}).Where("`code` = ?", code).First(&result).Error

	return
}
