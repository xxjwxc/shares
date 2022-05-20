package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _LhbDailyTblMgr struct {
	*_BaseMgr
}

// LhbDailyTblMgr open func
func LhbDailyTblMgr(db *gorm.DB) *_LhbDailyTblMgr {
	if db == nil {
		panic(fmt.Errorf("LhbDailyTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LhbDailyTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lhb_daily_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LhbDailyTblMgr) GetTableName() string {
	return "lhb_daily_tbl"
}

// Reset 重置gorm会话
func (obj *_LhbDailyTblMgr) Reset() *_LhbDailyTblMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LhbDailyTblMgr) Get() (result LhbDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LhbDailyTbl{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LhbDailyTblMgr) Gets() (results []*LhbDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LhbDailyTbl{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LhbDailyTblMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LhbDailyTbl{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_LhbDailyTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取 股票代码
func (obj *_LhbDailyTblMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithName name获取 股票名字
func (obj *_LhbDailyTblMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDay0 day0获取 当日0点时间戳
func (obj *_LhbDailyTblMgr) WithDay0(day0 int64) Option {
	return optionFunc(func(o *options) { o.query["day0"] = day0 })
}

// WithDayStr day_str获取 入池时间
func (obj *_LhbDailyTblMgr) WithDayStr(dayStr string) Option {
	return optionFunc(func(o *options) { o.query["day_str"] = dayStr })
}

// WithPrice price获取 当前价格
func (obj *_LhbDailyTblMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithJlr jlr获取 龙虎榜净流入(万元)
func (obj *_LhbDailyTblMgr) WithJlr(jlr float64) Option {
	return optionFunc(func(o *options) { o.query["jlr"] = jlr })
}

// WithJgJlr jg_jlr获取 机构净流入(万元)
func (obj *_LhbDailyTblMgr) WithJgJlr(jgJlr float64) Option {
	return optionFunc(func(o *options) { o.query["jg_jlr"] = jgJlr })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_LhbDailyTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// GetByOption 功能选项模式获取
func (obj *_LhbDailyTblMgr) GetByOption(opts ...Option) (result LhbDailyTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LhbDailyTbl{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LhbDailyTblMgr) GetByOptions(opts ...Option) (results []*LhbDailyTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LhbDailyTbl{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LhbDailyTblMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LhbDailyTbl, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LhbDailyTbl{}).Where(options.query)
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
func (obj *_LhbDailyTblMgr) GetFromID(id int) (result LhbDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LhbDailyTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_LhbDailyTblMgr) GetBatchFromID(ids []int) (results []*LhbDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LhbDailyTbl{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 股票代码
func (obj *_LhbDailyTblMgr) GetFromCode(code string) (results []*LhbDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LhbDailyTbl{}).Where("`code` = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量查找 股票代码
func (obj *_LhbDailyTblMgr) GetBatchFromCode(codes []string) (results []*LhbDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LhbDailyTbl{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 股票名字
func (obj *_LhbDailyTblMgr) GetFromName(name string) (results []*LhbDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LhbDailyTbl{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 股票名字
func (obj *_LhbDailyTblMgr) GetBatchFromName(names []string) (results []*LhbDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LhbDailyTbl{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromDay0 通过day0获取内容 当日0点时间戳
func (obj *_LhbDailyTblMgr) GetFromDay0(day0 int64) (results []*LhbDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LhbDailyTbl{}).Where("`day0` = ?", day0).Find(&results).Error

	return
}

// GetBatchFromDay0 批量查找 当日0点时间戳
func (obj *_LhbDailyTblMgr) GetBatchFromDay0(day0s []int64) (results []*LhbDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LhbDailyTbl{}).Where("`day0` IN (?)", day0s).Find(&results).Error

	return
}

// GetFromDayStr 通过day_str获取内容 入池时间
func (obj *_LhbDailyTblMgr) GetFromDayStr(dayStr string) (results []*LhbDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LhbDailyTbl{}).Where("`day_str` = ?", dayStr).Find(&results).Error

	return
}

// GetBatchFromDayStr 批量查找 入池时间
func (obj *_LhbDailyTblMgr) GetBatchFromDayStr(dayStrs []string) (results []*LhbDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LhbDailyTbl{}).Where("`day_str` IN (?)", dayStrs).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 当前价格
func (obj *_LhbDailyTblMgr) GetFromPrice(price float64) (results []*LhbDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LhbDailyTbl{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 当前价格
func (obj *_LhbDailyTblMgr) GetBatchFromPrice(prices []float64) (results []*LhbDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LhbDailyTbl{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromJlr 通过jlr获取内容 龙虎榜净流入(万元)
func (obj *_LhbDailyTblMgr) GetFromJlr(jlr float64) (results []*LhbDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LhbDailyTbl{}).Where("`jlr` = ?", jlr).Find(&results).Error

	return
}

// GetBatchFromJlr 批量查找 龙虎榜净流入(万元)
func (obj *_LhbDailyTblMgr) GetBatchFromJlr(jlrs []float64) (results []*LhbDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LhbDailyTbl{}).Where("`jlr` IN (?)", jlrs).Find(&results).Error

	return
}

// GetFromJgJlr 通过jg_jlr获取内容 机构净流入(万元)
func (obj *_LhbDailyTblMgr) GetFromJgJlr(jgJlr float64) (results []*LhbDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LhbDailyTbl{}).Where("`jg_jlr` = ?", jgJlr).Find(&results).Error

	return
}

// GetBatchFromJgJlr 批量查找 机构净流入(万元)
func (obj *_LhbDailyTblMgr) GetBatchFromJgJlr(jgJlrs []float64) (results []*LhbDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LhbDailyTbl{}).Where("`jg_jlr` IN (?)", jgJlrs).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_LhbDailyTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*LhbDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LhbDailyTbl{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_LhbDailyTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*LhbDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LhbDailyTbl{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LhbDailyTblMgr) FetchByPrimaryKey(id int) (result LhbDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LhbDailyTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueIndexByCode primary or index 获取唯一内容
func (obj *_LhbDailyTblMgr) FetchUniqueIndexByCode(code string, day0 int64) (result LhbDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LhbDailyTbl{}).Where("`code` = ? AND `day0` = ?", code, day0).First(&result).Error

	return
}
