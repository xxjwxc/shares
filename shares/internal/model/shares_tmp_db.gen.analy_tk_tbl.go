package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AnalyTkTblMgr struct {
	*_BaseMgr
}

// AnalyTkTblMgr open func
func AnalyTkTblMgr(db *gorm.DB) *_AnalyTkTblMgr {
	if db == nil {
		panic(fmt.Errorf("AnalyTkTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AnalyTkTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("analy_tk_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AnalyTkTblMgr) GetTableName() string {
	return "analy_tk_tbl"
}

// Reset 重置gorm会话
func (obj *_AnalyTkTblMgr) Reset() *_AnalyTkTblMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AnalyTkTblMgr) Get() (result AnalyTkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyTkTbl{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AnalyTkTblMgr) Gets() (results []*AnalyTkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyTkTbl{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AnalyTkTblMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AnalyTkTbl{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_AnalyTkTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取 股票代码
func (obj *_AnalyTkTblMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithName name获取 股票名字
func (obj *_AnalyTkTblMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDay0 day0获取 当日0点时间戳
func (obj *_AnalyTkTblMgr) WithDay0(day0 int64) Option {
	return optionFunc(func(o *options) { o.query["day0"] = day0 })
}

// WithDay0Str day0_str获取 入池时间
func (obj *_AnalyTkTblMgr) WithDay0Str(day0Str string) Option {
	return optionFunc(func(o *options) { o.query["day0_str"] = day0Str })
}

// WithPrice price获取 价格
func (obj *_AnalyTkTblMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithMin min获取 跳空最低价
func (obj *_AnalyTkTblMgr) WithMin(min float64) Option {
	return optionFunc(func(o *options) { o.query["min"] = min })
}

// WithMax max获取 当日最高值
func (obj *_AnalyTkTblMgr) WithMax(max float64) Option {
	return optionFunc(func(o *options) { o.query["max"] = max })
}

// WithFlag flag获取 标记：0 默认 ，1：回补了
func (obj *_AnalyTkTblMgr) WithFlag(flag int) Option {
	return optionFunc(func(o *options) { o.query["flag"] = flag })
}

// WithPercent percent获取 百分比
func (obj *_AnalyTkTblMgr) WithPercent(percent float64) Option {
	return optionFunc(func(o *options) { o.query["percent"] = percent })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_AnalyTkTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// GetByOption 功能选项模式获取
func (obj *_AnalyTkTblMgr) GetByOption(opts ...Option) (result AnalyTkTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AnalyTkTbl{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AnalyTkTblMgr) GetByOptions(opts ...Option) (results []*AnalyTkTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AnalyTkTbl{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AnalyTkTblMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AnalyTkTbl, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AnalyTkTbl{}).Where(options.query)
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
func (obj *_AnalyTkTblMgr) GetFromID(id int) (result AnalyTkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyTkTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_AnalyTkTblMgr) GetBatchFromID(ids []int) (results []*AnalyTkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyTkTbl{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 股票代码
func (obj *_AnalyTkTblMgr) GetFromCode(code string) (results []*AnalyTkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyTkTbl{}).Where("`code` = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量查找 股票代码
func (obj *_AnalyTkTblMgr) GetBatchFromCode(codes []string) (results []*AnalyTkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyTkTbl{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 股票名字
func (obj *_AnalyTkTblMgr) GetFromName(name string) (results []*AnalyTkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyTkTbl{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 股票名字
func (obj *_AnalyTkTblMgr) GetBatchFromName(names []string) (results []*AnalyTkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyTkTbl{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromDay0 通过day0获取内容 当日0点时间戳
func (obj *_AnalyTkTblMgr) GetFromDay0(day0 int64) (results []*AnalyTkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyTkTbl{}).Where("`day0` = ?", day0).Find(&results).Error

	return
}

// GetBatchFromDay0 批量查找 当日0点时间戳
func (obj *_AnalyTkTblMgr) GetBatchFromDay0(day0s []int64) (results []*AnalyTkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyTkTbl{}).Where("`day0` IN (?)", day0s).Find(&results).Error

	return
}

// GetFromDay0Str 通过day0_str获取内容 入池时间
func (obj *_AnalyTkTblMgr) GetFromDay0Str(day0Str string) (results []*AnalyTkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyTkTbl{}).Where("`day0_str` = ?", day0Str).Find(&results).Error

	return
}

// GetBatchFromDay0Str 批量查找 入池时间
func (obj *_AnalyTkTblMgr) GetBatchFromDay0Str(day0Strs []string) (results []*AnalyTkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyTkTbl{}).Where("`day0_str` IN (?)", day0Strs).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 价格
func (obj *_AnalyTkTblMgr) GetFromPrice(price float64) (results []*AnalyTkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyTkTbl{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 价格
func (obj *_AnalyTkTblMgr) GetBatchFromPrice(prices []float64) (results []*AnalyTkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyTkTbl{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromMin 通过min获取内容 跳空最低价
func (obj *_AnalyTkTblMgr) GetFromMin(min float64) (results []*AnalyTkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyTkTbl{}).Where("`min` = ?", min).Find(&results).Error

	return
}

// GetBatchFromMin 批量查找 跳空最低价
func (obj *_AnalyTkTblMgr) GetBatchFromMin(mins []float64) (results []*AnalyTkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyTkTbl{}).Where("`min` IN (?)", mins).Find(&results).Error

	return
}

// GetFromMax 通过max获取内容 当日最高值
func (obj *_AnalyTkTblMgr) GetFromMax(max float64) (results []*AnalyTkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyTkTbl{}).Where("`max` = ?", max).Find(&results).Error

	return
}

// GetBatchFromMax 批量查找 当日最高值
func (obj *_AnalyTkTblMgr) GetBatchFromMax(maxs []float64) (results []*AnalyTkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyTkTbl{}).Where("`max` IN (?)", maxs).Find(&results).Error

	return
}

// GetFromFlag 通过flag获取内容 标记：0 默认 ，1：回补了
func (obj *_AnalyTkTblMgr) GetFromFlag(flag int) (results []*AnalyTkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyTkTbl{}).Where("`flag` = ?", flag).Find(&results).Error

	return
}

// GetBatchFromFlag 批量查找 标记：0 默认 ，1：回补了
func (obj *_AnalyTkTblMgr) GetBatchFromFlag(flags []int) (results []*AnalyTkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyTkTbl{}).Where("`flag` IN (?)", flags).Find(&results).Error

	return
}

// GetFromPercent 通过percent获取内容 百分比
func (obj *_AnalyTkTblMgr) GetFromPercent(percent float64) (results []*AnalyTkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyTkTbl{}).Where("`percent` = ?", percent).Find(&results).Error

	return
}

// GetBatchFromPercent 批量查找 百分比
func (obj *_AnalyTkTblMgr) GetBatchFromPercent(percents []float64) (results []*AnalyTkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyTkTbl{}).Where("`percent` IN (?)", percents).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_AnalyTkTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*AnalyTkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyTkTbl{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_AnalyTkTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*AnalyTkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyTkTbl{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AnalyTkTblMgr) FetchByPrimaryKey(id int) (result AnalyTkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyTkTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueIndexByCode primary or index 获取唯一内容
func (obj *_AnalyTkTblMgr) FetchUniqueIndexByCode(code string, day0 int64) (result AnalyTkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyTkTbl{}).Where("`code` = ? AND `day0` = ?", code, day0).First(&result).Error

	return
}
