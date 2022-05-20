package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AnalyCdTblMgr struct {
	*_BaseMgr
}

// AnalyCdTblMgr open func
func AnalyCdTblMgr(db *gorm.DB) *_AnalyCdTblMgr {
	if db == nil {
		panic(fmt.Errorf("AnalyCdTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AnalyCdTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("analy_cd_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AnalyCdTblMgr) GetTableName() string {
	return "analy_cd_tbl"
}

// Reset 重置gorm会话
func (obj *_AnalyCdTblMgr) Reset() *_AnalyCdTblMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AnalyCdTblMgr) Get() (result AnalyCdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyCdTbl{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AnalyCdTblMgr) Gets() (results []*AnalyCdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyCdTbl{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AnalyCdTblMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AnalyCdTbl{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_AnalyCdTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取 股票代码
func (obj *_AnalyCdTblMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithName name获取 股票名字
func (obj *_AnalyCdTblMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDay0 day0获取 入池时间
func (obj *_AnalyCdTblMgr) WithDay0(day0 int64) Option {
	return optionFunc(func(o *options) { o.query["day0"] = day0 })
}

// WithDayStr day_str获取 入池时间
func (obj *_AnalyCdTblMgr) WithDayStr(dayStr string) Option {
	return optionFunc(func(o *options) { o.query["day_str"] = dayStr })
}

// WithPrice price获取 价格
func (obj *_AnalyCdTblMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithNewPrice new_price获取 价格
func (obj *_AnalyCdTblMgr) WithNewPrice(newPrice float64) Option {
	return optionFunc(func(o *options) { o.query["new_price"] = newPrice })
}

// WithPercent percent获取 百分比
func (obj *_AnalyCdTblMgr) WithPercent(percent float64) Option {
	return optionFunc(func(o *options) { o.query["percent"] = percent })
}

// WithTurnoverRate turnover_rate获取 换手率
func (obj *_AnalyCdTblMgr) WithTurnoverRate(turnoverRate float64) Option {
	return optionFunc(func(o *options) { o.query["turnover_rate"] = turnoverRate })
}

// WithScore score获取 得分
func (obj *_AnalyCdTblMgr) WithScore(score int) Option {
	return optionFunc(func(o *options) { o.query["score"] = score })
}

// WithDoc doc获取 描述
func (obj *_AnalyCdTblMgr) WithDoc(doc string) Option {
	return optionFunc(func(o *options) { o.query["doc"] = doc })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_AnalyCdTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// GetByOption 功能选项模式获取
func (obj *_AnalyCdTblMgr) GetByOption(opts ...Option) (result AnalyCdTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AnalyCdTbl{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AnalyCdTblMgr) GetByOptions(opts ...Option) (results []*AnalyCdTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AnalyCdTbl{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AnalyCdTblMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AnalyCdTbl, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AnalyCdTbl{}).Where(options.query)
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
func (obj *_AnalyCdTblMgr) GetFromID(id int) (result AnalyCdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyCdTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_AnalyCdTblMgr) GetBatchFromID(ids []int) (results []*AnalyCdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyCdTbl{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 股票代码
func (obj *_AnalyCdTblMgr) GetFromCode(code string) (result AnalyCdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyCdTbl{}).Where("`code` = ?", code).First(&result).Error

	return
}

// GetBatchFromCode 批量查找 股票代码
func (obj *_AnalyCdTblMgr) GetBatchFromCode(codes []string) (results []*AnalyCdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyCdTbl{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 股票名字
func (obj *_AnalyCdTblMgr) GetFromName(name string) (results []*AnalyCdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyCdTbl{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 股票名字
func (obj *_AnalyCdTblMgr) GetBatchFromName(names []string) (results []*AnalyCdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyCdTbl{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromDay0 通过day0获取内容 入池时间
func (obj *_AnalyCdTblMgr) GetFromDay0(day0 int64) (results []*AnalyCdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyCdTbl{}).Where("`day0` = ?", day0).Find(&results).Error

	return
}

// GetBatchFromDay0 批量查找 入池时间
func (obj *_AnalyCdTblMgr) GetBatchFromDay0(day0s []int64) (results []*AnalyCdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyCdTbl{}).Where("`day0` IN (?)", day0s).Find(&results).Error

	return
}

// GetFromDayStr 通过day_str获取内容 入池时间
func (obj *_AnalyCdTblMgr) GetFromDayStr(dayStr string) (results []*AnalyCdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyCdTbl{}).Where("`day_str` = ?", dayStr).Find(&results).Error

	return
}

// GetBatchFromDayStr 批量查找 入池时间
func (obj *_AnalyCdTblMgr) GetBatchFromDayStr(dayStrs []string) (results []*AnalyCdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyCdTbl{}).Where("`day_str` IN (?)", dayStrs).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 价格
func (obj *_AnalyCdTblMgr) GetFromPrice(price float64) (results []*AnalyCdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyCdTbl{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 价格
func (obj *_AnalyCdTblMgr) GetBatchFromPrice(prices []float64) (results []*AnalyCdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyCdTbl{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromNewPrice 通过new_price获取内容 价格
func (obj *_AnalyCdTblMgr) GetFromNewPrice(newPrice float64) (results []*AnalyCdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyCdTbl{}).Where("`new_price` = ?", newPrice).Find(&results).Error

	return
}

// GetBatchFromNewPrice 批量查找 价格
func (obj *_AnalyCdTblMgr) GetBatchFromNewPrice(newPrices []float64) (results []*AnalyCdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyCdTbl{}).Where("`new_price` IN (?)", newPrices).Find(&results).Error

	return
}

// GetFromPercent 通过percent获取内容 百分比
func (obj *_AnalyCdTblMgr) GetFromPercent(percent float64) (results []*AnalyCdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyCdTbl{}).Where("`percent` = ?", percent).Find(&results).Error

	return
}

// GetBatchFromPercent 批量查找 百分比
func (obj *_AnalyCdTblMgr) GetBatchFromPercent(percents []float64) (results []*AnalyCdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyCdTbl{}).Where("`percent` IN (?)", percents).Find(&results).Error

	return
}

// GetFromTurnoverRate 通过turnover_rate获取内容 换手率
func (obj *_AnalyCdTblMgr) GetFromTurnoverRate(turnoverRate float64) (results []*AnalyCdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyCdTbl{}).Where("`turnover_rate` = ?", turnoverRate).Find(&results).Error

	return
}

// GetBatchFromTurnoverRate 批量查找 换手率
func (obj *_AnalyCdTblMgr) GetBatchFromTurnoverRate(turnoverRates []float64) (results []*AnalyCdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyCdTbl{}).Where("`turnover_rate` IN (?)", turnoverRates).Find(&results).Error

	return
}

// GetFromScore 通过score获取内容 得分
func (obj *_AnalyCdTblMgr) GetFromScore(score int) (results []*AnalyCdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyCdTbl{}).Where("`score` = ?", score).Find(&results).Error

	return
}

// GetBatchFromScore 批量查找 得分
func (obj *_AnalyCdTblMgr) GetBatchFromScore(scores []int) (results []*AnalyCdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyCdTbl{}).Where("`score` IN (?)", scores).Find(&results).Error

	return
}

// GetFromDoc 通过doc获取内容 描述
func (obj *_AnalyCdTblMgr) GetFromDoc(doc string) (results []*AnalyCdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyCdTbl{}).Where("`doc` = ?", doc).Find(&results).Error

	return
}

// GetBatchFromDoc 批量查找 描述
func (obj *_AnalyCdTblMgr) GetBatchFromDoc(docs []string) (results []*AnalyCdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyCdTbl{}).Where("`doc` IN (?)", docs).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_AnalyCdTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*AnalyCdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyCdTbl{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_AnalyCdTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*AnalyCdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyCdTbl{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AnalyCdTblMgr) FetchByPrimaryKey(id int) (result AnalyCdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyCdTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByCode primary or index 获取唯一内容
func (obj *_AnalyCdTblMgr) FetchUniqueByCode(code string) (result AnalyCdTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyCdTbl{}).Where("`code` = ?", code).First(&result).Error

	return
}
