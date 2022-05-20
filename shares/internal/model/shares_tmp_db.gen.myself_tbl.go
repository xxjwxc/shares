package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type _MyselfTblMgr struct {
	*_BaseMgr
}

// MyselfTblMgr open func
func MyselfTblMgr(db *gorm.DB) *_MyselfTblMgr {
	if db == nil {
		panic(fmt.Errorf("MyselfTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MyselfTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("myself_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MyselfTblMgr) GetTableName() string {
	return "myself_tbl"
}

// Reset 重置gorm会话
func (obj *_MyselfTblMgr) Reset() *_MyselfTblMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MyselfTblMgr) Get() (result MyselfTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MyselfTbl{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MyselfTblMgr) Gets() (results []*MyselfTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MyselfTbl{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MyselfTblMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(MyselfTbl{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_MyselfTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取 股票代码
func (obj *_MyselfTblMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithName name获取 股票名字
func (obj *_MyselfTblMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDay day获取 入池时间
func (obj *_MyselfTblMgr) WithDay(day datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["day"] = day })
}

// WithDayStr day_str获取 入池时间
func (obj *_MyselfTblMgr) WithDayStr(dayStr string) Option {
	return optionFunc(func(o *options) { o.query["day_str"] = dayStr })
}

// WithPrice price获取 价格
func (obj *_MyselfTblMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithNewPrice new_price获取 价格
func (obj *_MyselfTblMgr) WithNewPrice(newPrice float64) Option {
	return optionFunc(func(o *options) { o.query["new_price"] = newPrice })
}

// WithPercent percent获取 百分比
func (obj *_MyselfTblMgr) WithPercent(percent float64) Option {
	return optionFunc(func(o *options) { o.query["percent"] = percent })
}

// WithTurnoverRate turnover_rate获取 换手率
func (obj *_MyselfTblMgr) WithTurnoverRate(turnoverRate float64) Option {
	return optionFunc(func(o *options) { o.query["turnover_rate"] = turnoverRate })
}

// WithNumber number获取 数量
func (obj *_MyselfTblMgr) WithNumber(number int) Option {
	return optionFunc(func(o *options) { o.query["number"] = number })
}

// WithDesc desc获取 描述
func (obj *_MyselfTblMgr) WithDesc(desc string) Option {
	return optionFunc(func(o *options) { o.query["desc"] = desc })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_MyselfTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// GetByOption 功能选项模式获取
func (obj *_MyselfTblMgr) GetByOption(opts ...Option) (result MyselfTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MyselfTbl{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MyselfTblMgr) GetByOptions(opts ...Option) (results []*MyselfTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MyselfTbl{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_MyselfTblMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]MyselfTbl, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(MyselfTbl{}).Where(options.query)
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
func (obj *_MyselfTblMgr) GetFromID(id int) (result MyselfTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MyselfTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_MyselfTblMgr) GetBatchFromID(ids []int) (results []*MyselfTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MyselfTbl{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 股票代码
func (obj *_MyselfTblMgr) GetFromCode(code string) (result MyselfTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MyselfTbl{}).Where("`code` = ?", code).First(&result).Error

	return
}

// GetBatchFromCode 批量查找 股票代码
func (obj *_MyselfTblMgr) GetBatchFromCode(codes []string) (results []*MyselfTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MyselfTbl{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 股票名字
func (obj *_MyselfTblMgr) GetFromName(name string) (results []*MyselfTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MyselfTbl{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 股票名字
func (obj *_MyselfTblMgr) GetBatchFromName(names []string) (results []*MyselfTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MyselfTbl{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromDay 通过day获取内容 入池时间
func (obj *_MyselfTblMgr) GetFromDay(day datatypes.Date) (results []*MyselfTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MyselfTbl{}).Where("`day` = ?", day).Find(&results).Error

	return
}

// GetBatchFromDay 批量查找 入池时间
func (obj *_MyselfTblMgr) GetBatchFromDay(days []datatypes.Date) (results []*MyselfTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MyselfTbl{}).Where("`day` IN (?)", days).Find(&results).Error

	return
}

// GetFromDayStr 通过day_str获取内容 入池时间
func (obj *_MyselfTblMgr) GetFromDayStr(dayStr string) (results []*MyselfTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MyselfTbl{}).Where("`day_str` = ?", dayStr).Find(&results).Error

	return
}

// GetBatchFromDayStr 批量查找 入池时间
func (obj *_MyselfTblMgr) GetBatchFromDayStr(dayStrs []string) (results []*MyselfTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MyselfTbl{}).Where("`day_str` IN (?)", dayStrs).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 价格
func (obj *_MyselfTblMgr) GetFromPrice(price float64) (results []*MyselfTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MyselfTbl{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 价格
func (obj *_MyselfTblMgr) GetBatchFromPrice(prices []float64) (results []*MyselfTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MyselfTbl{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromNewPrice 通过new_price获取内容 价格
func (obj *_MyselfTblMgr) GetFromNewPrice(newPrice float64) (results []*MyselfTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MyselfTbl{}).Where("`new_price` = ?", newPrice).Find(&results).Error

	return
}

// GetBatchFromNewPrice 批量查找 价格
func (obj *_MyselfTblMgr) GetBatchFromNewPrice(newPrices []float64) (results []*MyselfTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MyselfTbl{}).Where("`new_price` IN (?)", newPrices).Find(&results).Error

	return
}

// GetFromPercent 通过percent获取内容 百分比
func (obj *_MyselfTblMgr) GetFromPercent(percent float64) (results []*MyselfTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MyselfTbl{}).Where("`percent` = ?", percent).Find(&results).Error

	return
}

// GetBatchFromPercent 批量查找 百分比
func (obj *_MyselfTblMgr) GetBatchFromPercent(percents []float64) (results []*MyselfTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MyselfTbl{}).Where("`percent` IN (?)", percents).Find(&results).Error

	return
}

// GetFromTurnoverRate 通过turnover_rate获取内容 换手率
func (obj *_MyselfTblMgr) GetFromTurnoverRate(turnoverRate float64) (results []*MyselfTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MyselfTbl{}).Where("`turnover_rate` = ?", turnoverRate).Find(&results).Error

	return
}

// GetBatchFromTurnoverRate 批量查找 换手率
func (obj *_MyselfTblMgr) GetBatchFromTurnoverRate(turnoverRates []float64) (results []*MyselfTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MyselfTbl{}).Where("`turnover_rate` IN (?)", turnoverRates).Find(&results).Error

	return
}

// GetFromNumber 通过number获取内容 数量
func (obj *_MyselfTblMgr) GetFromNumber(number int) (results []*MyselfTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MyselfTbl{}).Where("`number` = ?", number).Find(&results).Error

	return
}

// GetBatchFromNumber 批量查找 数量
func (obj *_MyselfTblMgr) GetBatchFromNumber(numbers []int) (results []*MyselfTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MyselfTbl{}).Where("`number` IN (?)", numbers).Find(&results).Error

	return
}

// GetFromDesc 通过desc获取内容 描述
func (obj *_MyselfTblMgr) GetFromDesc(desc string) (results []*MyselfTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MyselfTbl{}).Where("`desc` = ?", desc).Find(&results).Error

	return
}

// GetBatchFromDesc 批量查找 描述
func (obj *_MyselfTblMgr) GetBatchFromDesc(descs []string) (results []*MyselfTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MyselfTbl{}).Where("`desc` IN (?)", descs).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_MyselfTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*MyselfTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MyselfTbl{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_MyselfTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*MyselfTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MyselfTbl{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MyselfTblMgr) FetchByPrimaryKey(id int) (result MyselfTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MyselfTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByCode primary or index 获取唯一内容
func (obj *_MyselfTblMgr) FetchUniqueByCode(code string) (result MyselfTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MyselfTbl{}).Where("`code` = ?", code).First(&result).Error

	return
}
