package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _NoTdxTblMgr struct {
	*_BaseMgr
}

// NoTdxTblMgr open func
func NoTdxTblMgr(db *gorm.DB) *_NoTdxTblMgr {
	if db == nil {
		panic(fmt.Errorf("NoTdxTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_NoTdxTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("no_tdx_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_NoTdxTblMgr) GetTableName() string {
	return "no_tdx_tbl"
}

// Reset 重置gorm会话
func (obj *_NoTdxTblMgr) Reset() *_NoTdxTblMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_NoTdxTblMgr) Get() (result NoTdxTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NoTdxTbl{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_NoTdxTblMgr) Gets() (results []*NoTdxTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NoTdxTbl{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_NoTdxTblMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(NoTdxTbl{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_NoTdxTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取 股票代码
func (obj *_NoTdxTblMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_NoTdxTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// GetByOption 功能选项模式获取
func (obj *_NoTdxTblMgr) GetByOption(opts ...Option) (result NoTdxTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(NoTdxTbl{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_NoTdxTblMgr) GetByOptions(opts ...Option) (results []*NoTdxTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(NoTdxTbl{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_NoTdxTblMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]NoTdxTbl, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(NoTdxTbl{}).Where(options.query)
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
func (obj *_NoTdxTblMgr) GetFromID(id int) (result NoTdxTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NoTdxTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_NoTdxTblMgr) GetBatchFromID(ids []int) (results []*NoTdxTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NoTdxTbl{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 股票代码
func (obj *_NoTdxTblMgr) GetFromCode(code string) (result NoTdxTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NoTdxTbl{}).Where("`code` = ?", code).First(&result).Error

	return
}

// GetBatchFromCode 批量查找 股票代码
func (obj *_NoTdxTblMgr) GetBatchFromCode(codes []string) (results []*NoTdxTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NoTdxTbl{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_NoTdxTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*NoTdxTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NoTdxTbl{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_NoTdxTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*NoTdxTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NoTdxTbl{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_NoTdxTblMgr) FetchByPrimaryKey(id int) (result NoTdxTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NoTdxTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByCode primary or index 获取唯一内容
func (obj *_NoTdxTblMgr) FetchUniqueByCode(code string) (result NoTdxTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NoTdxTbl{}).Where("`code` = ?", code).First(&result).Error

	return
}
