package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _NoPegTblMgr struct {
	*_BaseMgr
}

// NoPegTblMgr open func
func NoPegTblMgr(db *gorm.DB) *_NoPegTblMgr {
	if db == nil {
		panic(fmt.Errorf("NoPegTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_NoPegTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("no_peg_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_NoPegTblMgr) GetTableName() string {
	return "no_peg_tbl"
}

// Reset 重置gorm会话
func (obj *_NoPegTblMgr) Reset() *_NoPegTblMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_NoPegTblMgr) Get() (result NoPegTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NoPegTbl{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_NoPegTblMgr) Gets() (results []*NoPegTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NoPegTbl{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_NoPegTblMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(NoPegTbl{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_NoPegTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取 股票代码
func (obj *_NoPegTblMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_NoPegTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// GetByOption 功能选项模式获取
func (obj *_NoPegTblMgr) GetByOption(opts ...Option) (result NoPegTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(NoPegTbl{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_NoPegTblMgr) GetByOptions(opts ...Option) (results []*NoPegTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(NoPegTbl{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_NoPegTblMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]NoPegTbl, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(NoPegTbl{}).Where(options.query)
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
func (obj *_NoPegTblMgr) GetFromID(id int) (result NoPegTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NoPegTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_NoPegTblMgr) GetBatchFromID(ids []int) (results []*NoPegTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NoPegTbl{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 股票代码
func (obj *_NoPegTblMgr) GetFromCode(code string) (result NoPegTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NoPegTbl{}).Where("`code` = ?", code).First(&result).Error

	return
}

// GetBatchFromCode 批量查找 股票代码
func (obj *_NoPegTblMgr) GetBatchFromCode(codes []string) (results []*NoPegTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NoPegTbl{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_NoPegTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*NoPegTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NoPegTbl{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_NoPegTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*NoPegTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NoPegTbl{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_NoPegTblMgr) FetchByPrimaryKey(id int) (result NoPegTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NoPegTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByCode primary or index 获取唯一内容
func (obj *_NoPegTblMgr) FetchUniqueByCode(code string) (result NoPegTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NoPegTbl{}).Where("`code` = ?", code).First(&result).Error

	return
}
