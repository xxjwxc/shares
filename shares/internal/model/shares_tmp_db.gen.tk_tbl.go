package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _TkTblMgr struct {
	*_BaseMgr
}

// TkTblMgr open func
func TkTblMgr(db *gorm.DB) *_TkTblMgr {
	if db == nil {
		panic(fmt.Errorf("TkTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TkTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("tk_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TkTblMgr) GetTableName() string {
	return "tk_tbl"
}

// Reset 重置gorm会话
func (obj *_TkTblMgr) Reset() *_TkTblMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TkTblMgr) Get() (result TkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TkTbl{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TkTblMgr) Gets() (results []*TkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TkTbl{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TkTblMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(TkTbl{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_TkTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取 股票代码
func (obj *_TkTblMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithName name获取 股票名字
func (obj *_TkTblMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDay0 day0获取 当日0点时间戳
func (obj *_TkTblMgr) WithDay0(day0 int64) Option {
	return optionFunc(func(o *options) { o.query["day0"] = day0 })
}

// WithDay0Str day0_str获取 入池时间
func (obj *_TkTblMgr) WithDay0Str(day0Str string) Option {
	return optionFunc(func(o *options) { o.query["day0_str"] = day0Str })
}

// WithPrice price获取 价格
func (obj *_TkTblMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithMin min获取 跳空最低价
func (obj *_TkTblMgr) WithMin(min float64) Option {
	return optionFunc(func(o *options) { o.query["min"] = min })
}

// WithMax max获取 当日最高值
func (obj *_TkTblMgr) WithMax(max float64) Option {
	return optionFunc(func(o *options) { o.query["max"] = max })
}

// WithFlag flag获取 标记：0 默认 ，1：回补了
func (obj *_TkTblMgr) WithFlag(flag int) Option {
	return optionFunc(func(o *options) { o.query["flag"] = flag })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_TkTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdateAt update_at获取 回补时间
//
//
//回补时间1
func (obj *_TkTblMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// GetByOption 功能选项模式获取
func (obj *_TkTblMgr) GetByOption(opts ...Option) (result TkTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TkTbl{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TkTblMgr) GetByOptions(opts ...Option) (results []*TkTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TkTbl{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_TkTblMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]TkTbl, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(TkTbl{}).Where(options.query)
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
func (obj *_TkTblMgr) GetFromID(id int) (result TkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TkTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_TkTblMgr) GetBatchFromID(ids []int) (results []*TkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TkTbl{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 股票代码
func (obj *_TkTblMgr) GetFromCode(code string) (results []*TkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TkTbl{}).Where("`code` = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量查找 股票代码
func (obj *_TkTblMgr) GetBatchFromCode(codes []string) (results []*TkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TkTbl{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 股票名字
func (obj *_TkTblMgr) GetFromName(name string) (results []*TkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TkTbl{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 股票名字
func (obj *_TkTblMgr) GetBatchFromName(names []string) (results []*TkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TkTbl{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromDay0 通过day0获取内容 当日0点时间戳
func (obj *_TkTblMgr) GetFromDay0(day0 int64) (results []*TkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TkTbl{}).Where("`day0` = ?", day0).Find(&results).Error

	return
}

// GetBatchFromDay0 批量查找 当日0点时间戳
func (obj *_TkTblMgr) GetBatchFromDay0(day0s []int64) (results []*TkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TkTbl{}).Where("`day0` IN (?)", day0s).Find(&results).Error

	return
}

// GetFromDay0Str 通过day0_str获取内容 入池时间
func (obj *_TkTblMgr) GetFromDay0Str(day0Str string) (results []*TkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TkTbl{}).Where("`day0_str` = ?", day0Str).Find(&results).Error

	return
}

// GetBatchFromDay0Str 批量查找 入池时间
func (obj *_TkTblMgr) GetBatchFromDay0Str(day0Strs []string) (results []*TkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TkTbl{}).Where("`day0_str` IN (?)", day0Strs).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 价格
func (obj *_TkTblMgr) GetFromPrice(price float64) (results []*TkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TkTbl{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 价格
func (obj *_TkTblMgr) GetBatchFromPrice(prices []float64) (results []*TkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TkTbl{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromMin 通过min获取内容 跳空最低价
func (obj *_TkTblMgr) GetFromMin(min float64) (results []*TkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TkTbl{}).Where("`min` = ?", min).Find(&results).Error

	return
}

// GetBatchFromMin 批量查找 跳空最低价
func (obj *_TkTblMgr) GetBatchFromMin(mins []float64) (results []*TkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TkTbl{}).Where("`min` IN (?)", mins).Find(&results).Error

	return
}

// GetFromMax 通过max获取内容 当日最高值
func (obj *_TkTblMgr) GetFromMax(max float64) (results []*TkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TkTbl{}).Where("`max` = ?", max).Find(&results).Error

	return
}

// GetBatchFromMax 批量查找 当日最高值
func (obj *_TkTblMgr) GetBatchFromMax(maxs []float64) (results []*TkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TkTbl{}).Where("`max` IN (?)", maxs).Find(&results).Error

	return
}

// GetFromFlag 通过flag获取内容 标记：0 默认 ，1：回补了
func (obj *_TkTblMgr) GetFromFlag(flag int) (results []*TkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TkTbl{}).Where("`flag` = ?", flag).Find(&results).Error

	return
}

// GetBatchFromFlag 批量查找 标记：0 默认 ，1：回补了
func (obj *_TkTblMgr) GetBatchFromFlag(flags []int) (results []*TkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TkTbl{}).Where("`flag` IN (?)", flags).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_TkTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*TkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TkTbl{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_TkTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*TkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TkTbl{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容 回补时间
//
//
//回补时间1
func (obj *_TkTblMgr) GetFromUpdateAt(updateAt time.Time) (results []*TkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TkTbl{}).Where("`update_at` = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量查找 回补时间
//
//
//回补时间1
func (obj *_TkTblMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*TkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TkTbl{}).Where("`update_at` IN (?)", updateAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_TkTblMgr) FetchByPrimaryKey(id int) (result TkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TkTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueIndexByCode primary or index 获取唯一内容
func (obj *_TkTblMgr) FetchUniqueIndexByCode(code string, day0 int64) (result TkTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TkTbl{}).Where("`code` = ? AND `day0` = ?", code, day0).First(&result).Error

	return
}
