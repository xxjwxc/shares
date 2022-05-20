package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _ZljlrDailyTblMgr struct {
	*_BaseMgr
}

// ZljlrDailyTblMgr open func
func ZljlrDailyTblMgr(db *gorm.DB) *_ZljlrDailyTblMgr {
	if db == nil {
		panic(fmt.Errorf("ZljlrDailyTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ZljlrDailyTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("zljlr_daily_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ZljlrDailyTblMgr) GetTableName() string {
	return "zljlr_daily_tbl"
}

// Reset 重置gorm会话
func (obj *_ZljlrDailyTblMgr) Reset() *_ZljlrDailyTblMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ZljlrDailyTblMgr) Get() (result ZljlrDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZljlrDailyTbl{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ZljlrDailyTblMgr) Gets() (results []*ZljlrDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZljlrDailyTbl{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ZljlrDailyTblMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(ZljlrDailyTbl{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ZljlrDailyTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取 股票代码
func (obj *_ZljlrDailyTblMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithDay0 day0获取 当日0点时间戳
func (obj *_ZljlrDailyTblMgr) WithDay0(day0 int64) Option {
	return optionFunc(func(o *options) { o.query["day0"] = day0 })
}

// WithDayStr day_str获取 入池时间
func (obj *_ZljlrDailyTblMgr) WithDayStr(dayStr string) Option {
	return optionFunc(func(o *options) { o.query["day_str"] = dayStr })
}

// WithPrice price获取 当前价格(万元)
func (obj *_ZljlrDailyTblMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_ZljlrDailyTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// GetByOption 功能选项模式获取
func (obj *_ZljlrDailyTblMgr) GetByOption(opts ...Option) (result ZljlrDailyTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ZljlrDailyTbl{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ZljlrDailyTblMgr) GetByOptions(opts ...Option) (results []*ZljlrDailyTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ZljlrDailyTbl{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_ZljlrDailyTblMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]ZljlrDailyTbl, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(ZljlrDailyTbl{}).Where(options.query)
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
func (obj *_ZljlrDailyTblMgr) GetFromID(id int) (result ZljlrDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZljlrDailyTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_ZljlrDailyTblMgr) GetBatchFromID(ids []int) (results []*ZljlrDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZljlrDailyTbl{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 股票代码
func (obj *_ZljlrDailyTblMgr) GetFromCode(code string) (results []*ZljlrDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZljlrDailyTbl{}).Where("`code` = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量查找 股票代码
func (obj *_ZljlrDailyTblMgr) GetBatchFromCode(codes []string) (results []*ZljlrDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZljlrDailyTbl{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromDay0 通过day0获取内容 当日0点时间戳
func (obj *_ZljlrDailyTblMgr) GetFromDay0(day0 int64) (results []*ZljlrDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZljlrDailyTbl{}).Where("`day0` = ?", day0).Find(&results).Error

	return
}

// GetBatchFromDay0 批量查找 当日0点时间戳
func (obj *_ZljlrDailyTblMgr) GetBatchFromDay0(day0s []int64) (results []*ZljlrDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZljlrDailyTbl{}).Where("`day0` IN (?)", day0s).Find(&results).Error

	return
}

// GetFromDayStr 通过day_str获取内容 入池时间
func (obj *_ZljlrDailyTblMgr) GetFromDayStr(dayStr string) (results []*ZljlrDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZljlrDailyTbl{}).Where("`day_str` = ?", dayStr).Find(&results).Error

	return
}

// GetBatchFromDayStr 批量查找 入池时间
func (obj *_ZljlrDailyTblMgr) GetBatchFromDayStr(dayStrs []string) (results []*ZljlrDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZljlrDailyTbl{}).Where("`day_str` IN (?)", dayStrs).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 当前价格(万元)
func (obj *_ZljlrDailyTblMgr) GetFromPrice(price float64) (results []*ZljlrDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZljlrDailyTbl{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 当前价格(万元)
func (obj *_ZljlrDailyTblMgr) GetBatchFromPrice(prices []float64) (results []*ZljlrDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZljlrDailyTbl{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_ZljlrDailyTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*ZljlrDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZljlrDailyTbl{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_ZljlrDailyTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*ZljlrDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZljlrDailyTbl{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ZljlrDailyTblMgr) FetchByPrimaryKey(id int) (result ZljlrDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZljlrDailyTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueIndexByCode primary or index 获取唯一内容
func (obj *_ZljlrDailyTblMgr) FetchUniqueIndexByCode(code string, day0 int64) (result ZljlrDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ZljlrDailyTbl{}).Where("`code` = ? AND `day0` = ?", code, day0).First(&result).Error

	return
}
