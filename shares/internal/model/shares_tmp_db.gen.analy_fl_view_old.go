package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type _AnalyFlViewOldMgr struct {
	*_BaseMgr
}

// AnalyFlViewOldMgr open func
func AnalyFlViewOldMgr(db *gorm.DB) *_AnalyFlViewOldMgr {
	if db == nil {
		panic(fmt.Errorf("AnalyFlViewOldMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AnalyFlViewOldMgr{_BaseMgr: &_BaseMgr{DB: db.Table("analy_fl_view_old"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AnalyFlViewOldMgr) GetTableName() string {
	return "analy_fl_view_old"
}

// Reset 重置gorm会话
func (obj *_AnalyFlViewOldMgr) Reset() *_AnalyFlViewOldMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AnalyFlViewOldMgr) Get() (result AnalyFlViewOld, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlViewOld{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AnalyFlViewOldMgr) Gets() (results []*AnalyFlViewOld, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlViewOld{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AnalyFlViewOldMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AnalyFlViewOld{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_AnalyFlViewOldMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取 股票代码
func (obj *_AnalyFlViewOldMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithName name获取 股票名字
func (obj *_AnalyFlViewOldMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithUpName up_name获取 行业名
func (obj *_AnalyFlViewOldMgr) WithUpName(upName string) Option {
	return optionFunc(func(o *options) { o.query["up_name"] = upName })
}

// WithDay day获取 入池时间
func (obj *_AnalyFlViewOldMgr) WithDay(day datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["day"] = day })
}

// WithDayStr day_str获取 入池时间
func (obj *_AnalyFlViewOldMgr) WithDayStr(dayStr string) Option {
	return optionFunc(func(o *options) { o.query["day_str"] = dayStr })
}

// WithPrice price获取 价格
func (obj *_AnalyFlViewOldMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithNewPrice new_price获取 价格
func (obj *_AnalyFlViewOldMgr) WithNewPrice(newPrice float64) Option {
	return optionFunc(func(o *options) { o.query["new_price"] = newPrice })
}

// WithPercent percent获取 百分比
func (obj *_AnalyFlViewOldMgr) WithPercent(percent float64) Option {
	return optionFunc(func(o *options) { o.query["percent"] = percent })
}

// WithTurnoverRate turnover_rate获取 换手率
func (obj *_AnalyFlViewOldMgr) WithTurnoverRate(turnoverRate float64) Option {
	return optionFunc(func(o *options) { o.query["turnover_rate"] = turnoverRate })
}

// WithScore score获取 得分
func (obj *_AnalyFlViewOldMgr) WithScore(score int) Option {
	return optionFunc(func(o *options) { o.query["score"] = score })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_AnalyFlViewOldMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithNum num获取 总家数
func (obj *_AnalyFlViewOldMgr) WithNum(num int) Option {
	return optionFunc(func(o *options) { o.query["num"] = num })
}

// WithUp up获取 上涨家数
func (obj *_AnalyFlViewOldMgr) WithUp(up int) Option {
	return optionFunc(func(o *options) { o.query["up"] = up })
}

// WithHyName hy_name获取 行业名
func (obj *_AnalyFlViewOldMgr) WithHyName(hyName string) Option {
	return optionFunc(func(o *options) { o.query["hy_name"] = hyName })
}

// GetByOption 功能选项模式获取
func (obj *_AnalyFlViewOldMgr) GetByOption(opts ...Option) (result AnalyFlViewOld, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlViewOld{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AnalyFlViewOldMgr) GetByOptions(opts ...Option) (results []*AnalyFlViewOld, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlViewOld{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AnalyFlViewOldMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AnalyFlViewOld, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AnalyFlViewOld{}).Where(options.query)
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
func (obj *_AnalyFlViewOldMgr) GetFromID(id int) (results []*AnalyFlViewOld, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlViewOld{}).Where("`id` = ?", id).Find(&results).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_AnalyFlViewOldMgr) GetBatchFromID(ids []int) (results []*AnalyFlViewOld, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlViewOld{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 股票代码
func (obj *_AnalyFlViewOldMgr) GetFromCode(code string) (results []*AnalyFlViewOld, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlViewOld{}).Where("`code` = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量查找 股票代码
func (obj *_AnalyFlViewOldMgr) GetBatchFromCode(codes []string) (results []*AnalyFlViewOld, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlViewOld{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 股票名字
func (obj *_AnalyFlViewOldMgr) GetFromName(name string) (results []*AnalyFlViewOld, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlViewOld{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 股票名字
func (obj *_AnalyFlViewOldMgr) GetBatchFromName(names []string) (results []*AnalyFlViewOld, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlViewOld{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromUpName 通过up_name获取内容 行业名
func (obj *_AnalyFlViewOldMgr) GetFromUpName(upName string) (results []*AnalyFlViewOld, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlViewOld{}).Where("`up_name` = ?", upName).Find(&results).Error

	return
}

// GetBatchFromUpName 批量查找 行业名
func (obj *_AnalyFlViewOldMgr) GetBatchFromUpName(upNames []string) (results []*AnalyFlViewOld, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlViewOld{}).Where("`up_name` IN (?)", upNames).Find(&results).Error

	return
}

// GetFromDay 通过day获取内容 入池时间
func (obj *_AnalyFlViewOldMgr) GetFromDay(day datatypes.Date) (results []*AnalyFlViewOld, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlViewOld{}).Where("`day` = ?", day).Find(&results).Error

	return
}

// GetBatchFromDay 批量查找 入池时间
func (obj *_AnalyFlViewOldMgr) GetBatchFromDay(days []datatypes.Date) (results []*AnalyFlViewOld, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlViewOld{}).Where("`day` IN (?)", days).Find(&results).Error

	return
}

// GetFromDayStr 通过day_str获取内容 入池时间
func (obj *_AnalyFlViewOldMgr) GetFromDayStr(dayStr string) (results []*AnalyFlViewOld, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlViewOld{}).Where("`day_str` = ?", dayStr).Find(&results).Error

	return
}

// GetBatchFromDayStr 批量查找 入池时间
func (obj *_AnalyFlViewOldMgr) GetBatchFromDayStr(dayStrs []string) (results []*AnalyFlViewOld, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlViewOld{}).Where("`day_str` IN (?)", dayStrs).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 价格
func (obj *_AnalyFlViewOldMgr) GetFromPrice(price float64) (results []*AnalyFlViewOld, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlViewOld{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 价格
func (obj *_AnalyFlViewOldMgr) GetBatchFromPrice(prices []float64) (results []*AnalyFlViewOld, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlViewOld{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromNewPrice 通过new_price获取内容 价格
func (obj *_AnalyFlViewOldMgr) GetFromNewPrice(newPrice float64) (results []*AnalyFlViewOld, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlViewOld{}).Where("`new_price` = ?", newPrice).Find(&results).Error

	return
}

// GetBatchFromNewPrice 批量查找 价格
func (obj *_AnalyFlViewOldMgr) GetBatchFromNewPrice(newPrices []float64) (results []*AnalyFlViewOld, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlViewOld{}).Where("`new_price` IN (?)", newPrices).Find(&results).Error

	return
}

// GetFromPercent 通过percent获取内容 百分比
func (obj *_AnalyFlViewOldMgr) GetFromPercent(percent float64) (results []*AnalyFlViewOld, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlViewOld{}).Where("`percent` = ?", percent).Find(&results).Error

	return
}

// GetBatchFromPercent 批量查找 百分比
func (obj *_AnalyFlViewOldMgr) GetBatchFromPercent(percents []float64) (results []*AnalyFlViewOld, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlViewOld{}).Where("`percent` IN (?)", percents).Find(&results).Error

	return
}

// GetFromTurnoverRate 通过turnover_rate获取内容 换手率
func (obj *_AnalyFlViewOldMgr) GetFromTurnoverRate(turnoverRate float64) (results []*AnalyFlViewOld, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlViewOld{}).Where("`turnover_rate` = ?", turnoverRate).Find(&results).Error

	return
}

// GetBatchFromTurnoverRate 批量查找 换手率
func (obj *_AnalyFlViewOldMgr) GetBatchFromTurnoverRate(turnoverRates []float64) (results []*AnalyFlViewOld, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlViewOld{}).Where("`turnover_rate` IN (?)", turnoverRates).Find(&results).Error

	return
}

// GetFromScore 通过score获取内容 得分
func (obj *_AnalyFlViewOldMgr) GetFromScore(score int) (results []*AnalyFlViewOld, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlViewOld{}).Where("`score` = ?", score).Find(&results).Error

	return
}

// GetBatchFromScore 批量查找 得分
func (obj *_AnalyFlViewOldMgr) GetBatchFromScore(scores []int) (results []*AnalyFlViewOld, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlViewOld{}).Where("`score` IN (?)", scores).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_AnalyFlViewOldMgr) GetFromCreatedAt(createdAt time.Time) (results []*AnalyFlViewOld, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlViewOld{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_AnalyFlViewOldMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*AnalyFlViewOld, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlViewOld{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromNum 通过num获取内容 总家数
func (obj *_AnalyFlViewOldMgr) GetFromNum(num int) (results []*AnalyFlViewOld, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlViewOld{}).Where("`num` = ?", num).Find(&results).Error

	return
}

// GetBatchFromNum 批量查找 总家数
func (obj *_AnalyFlViewOldMgr) GetBatchFromNum(nums []int) (results []*AnalyFlViewOld, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlViewOld{}).Where("`num` IN (?)", nums).Find(&results).Error

	return
}

// GetFromUp 通过up获取内容 上涨家数
func (obj *_AnalyFlViewOldMgr) GetFromUp(up int) (results []*AnalyFlViewOld, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlViewOld{}).Where("`up` = ?", up).Find(&results).Error

	return
}

// GetBatchFromUp 批量查找 上涨家数
func (obj *_AnalyFlViewOldMgr) GetBatchFromUp(ups []int) (results []*AnalyFlViewOld, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlViewOld{}).Where("`up` IN (?)", ups).Find(&results).Error

	return
}

// GetFromHyName 通过hy_name获取内容 行业名
func (obj *_AnalyFlViewOldMgr) GetFromHyName(hyName string) (results []*AnalyFlViewOld, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlViewOld{}).Where("`hy_name` = ?", hyName).Find(&results).Error

	return
}

// GetBatchFromHyName 批量查找 行业名
func (obj *_AnalyFlViewOldMgr) GetBatchFromHyName(hyNames []string) (results []*AnalyFlViewOld, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyFlViewOld{}).Where("`hy_name` IN (?)", hyNames).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
