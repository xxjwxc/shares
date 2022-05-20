package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _CdViewMgr struct {
	*_BaseMgr
}

// CdViewMgr open func
func CdViewMgr(db *gorm.DB) *_CdViewMgr {
	if db == nil {
		panic(fmt.Errorf("CdViewMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CdViewMgr{_BaseMgr: &_BaseMgr{DB: db.Table("cd_view"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CdViewMgr) GetTableName() string {
	return "cd_view"
}

// Reset 重置gorm会话
func (obj *_CdViewMgr) Reset() *_CdViewMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CdViewMgr) Get() (result CdView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdView{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CdViewMgr) Gets() (results []*CdView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdView{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CdViewMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(CdView{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_CdViewMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取 股票代码
func (obj *_CdViewMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithName name获取 股票名字
func (obj *_CdViewMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDay0 day0获取 入池时间
func (obj *_CdViewMgr) WithDay0(day0 int64) Option {
	return optionFunc(func(o *options) { o.query["day0"] = day0 })
}

// WithDayStr day_str获取 入池时间
func (obj *_CdViewMgr) WithDayStr(dayStr string) Option {
	return optionFunc(func(o *options) { o.query["day_str"] = dayStr })
}

// WithPrice price获取 价格
func (obj *_CdViewMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithNewPrice new_price获取 价格
func (obj *_CdViewMgr) WithNewPrice(newPrice float64) Option {
	return optionFunc(func(o *options) { o.query["new_price"] = newPrice })
}

// WithPercent percent获取 百分比
func (obj *_CdViewMgr) WithPercent(percent float64) Option {
	return optionFunc(func(o *options) { o.query["percent"] = percent })
}

// WithTurnoverRate turnover_rate获取 换手率
func (obj *_CdViewMgr) WithTurnoverRate(turnoverRate float64) Option {
	return optionFunc(func(o *options) { o.query["turnover_rate"] = turnoverRate })
}

// WithScore score获取 得分
func (obj *_CdViewMgr) WithScore(score int) Option {
	return optionFunc(func(o *options) { o.query["score"] = score })
}

// WithDoc doc获取 描述
func (obj *_CdViewMgr) WithDoc(doc string) Option {
	return optionFunc(func(o *options) { o.query["doc"] = doc })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_CdViewMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithHyName hy_name获取 行业名
func (obj *_CdViewMgr) WithHyName(hyName string) Option {
	return optionFunc(func(o *options) { o.query["hy_name"] = hyName })
}

// GetByOption 功能选项模式获取
func (obj *_CdViewMgr) GetByOption(opts ...Option) (result CdView, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdView{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CdViewMgr) GetByOptions(opts ...Option) (results []*CdView, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CdView{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_CdViewMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]CdView, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(CdView{}).Where(options.query)
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
func (obj *_CdViewMgr) GetFromID(id int) (results []*CdView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdView{}).Where("`id` = ?", id).Find(&results).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_CdViewMgr) GetBatchFromID(ids []int) (results []*CdView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdView{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 股票代码
func (obj *_CdViewMgr) GetFromCode(code string) (results []*CdView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdView{}).Where("`code` = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量查找 股票代码
func (obj *_CdViewMgr) GetBatchFromCode(codes []string) (results []*CdView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdView{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 股票名字
func (obj *_CdViewMgr) GetFromName(name string) (results []*CdView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdView{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 股票名字
func (obj *_CdViewMgr) GetBatchFromName(names []string) (results []*CdView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdView{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromDay0 通过day0获取内容 入池时间
func (obj *_CdViewMgr) GetFromDay0(day0 int64) (results []*CdView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdView{}).Where("`day0` = ?", day0).Find(&results).Error

	return
}

// GetBatchFromDay0 批量查找 入池时间
func (obj *_CdViewMgr) GetBatchFromDay0(day0s []int64) (results []*CdView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdView{}).Where("`day0` IN (?)", day0s).Find(&results).Error

	return
}

// GetFromDayStr 通过day_str获取内容 入池时间
func (obj *_CdViewMgr) GetFromDayStr(dayStr string) (results []*CdView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdView{}).Where("`day_str` = ?", dayStr).Find(&results).Error

	return
}

// GetBatchFromDayStr 批量查找 入池时间
func (obj *_CdViewMgr) GetBatchFromDayStr(dayStrs []string) (results []*CdView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdView{}).Where("`day_str` IN (?)", dayStrs).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 价格
func (obj *_CdViewMgr) GetFromPrice(price float64) (results []*CdView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdView{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 价格
func (obj *_CdViewMgr) GetBatchFromPrice(prices []float64) (results []*CdView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdView{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromNewPrice 通过new_price获取内容 价格
func (obj *_CdViewMgr) GetFromNewPrice(newPrice float64) (results []*CdView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdView{}).Where("`new_price` = ?", newPrice).Find(&results).Error

	return
}

// GetBatchFromNewPrice 批量查找 价格
func (obj *_CdViewMgr) GetBatchFromNewPrice(newPrices []float64) (results []*CdView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdView{}).Where("`new_price` IN (?)", newPrices).Find(&results).Error

	return
}

// GetFromPercent 通过percent获取内容 百分比
func (obj *_CdViewMgr) GetFromPercent(percent float64) (results []*CdView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdView{}).Where("`percent` = ?", percent).Find(&results).Error

	return
}

// GetBatchFromPercent 批量查找 百分比
func (obj *_CdViewMgr) GetBatchFromPercent(percents []float64) (results []*CdView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdView{}).Where("`percent` IN (?)", percents).Find(&results).Error

	return
}

// GetFromTurnoverRate 通过turnover_rate获取内容 换手率
func (obj *_CdViewMgr) GetFromTurnoverRate(turnoverRate float64) (results []*CdView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdView{}).Where("`turnover_rate` = ?", turnoverRate).Find(&results).Error

	return
}

// GetBatchFromTurnoverRate 批量查找 换手率
func (obj *_CdViewMgr) GetBatchFromTurnoverRate(turnoverRates []float64) (results []*CdView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdView{}).Where("`turnover_rate` IN (?)", turnoverRates).Find(&results).Error

	return
}

// GetFromScore 通过score获取内容 得分
func (obj *_CdViewMgr) GetFromScore(score int) (results []*CdView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdView{}).Where("`score` = ?", score).Find(&results).Error

	return
}

// GetBatchFromScore 批量查找 得分
func (obj *_CdViewMgr) GetBatchFromScore(scores []int) (results []*CdView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdView{}).Where("`score` IN (?)", scores).Find(&results).Error

	return
}

// GetFromDoc 通过doc获取内容 描述
func (obj *_CdViewMgr) GetFromDoc(doc string) (results []*CdView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdView{}).Where("`doc` = ?", doc).Find(&results).Error

	return
}

// GetBatchFromDoc 批量查找 描述
func (obj *_CdViewMgr) GetBatchFromDoc(docs []string) (results []*CdView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdView{}).Where("`doc` IN (?)", docs).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_CdViewMgr) GetFromCreatedAt(createdAt time.Time) (results []*CdView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdView{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_CdViewMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*CdView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdView{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromHyName 通过hy_name获取内容 行业名
func (obj *_CdViewMgr) GetFromHyName(hyName string) (results []*CdView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdView{}).Where("`hy_name` = ?", hyName).Find(&results).Error

	return
}

// GetBatchFromHyName 批量查找 行业名
func (obj *_CdViewMgr) GetBatchFromHyName(hyNames []string) (results []*CdView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CdView{}).Where("`hy_name` IN (?)", hyNames).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
