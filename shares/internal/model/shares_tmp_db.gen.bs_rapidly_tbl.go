package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _BsRapidlyTblMgr struct {
	*_BaseMgr
}

// BsRapidlyTblMgr open func
func BsRapidlyTblMgr(db *gorm.DB) *_BsRapidlyTblMgr {
	if db == nil {
		panic(fmt.Errorf("BsRapidlyTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_BsRapidlyTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("bs_rapidly_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_BsRapidlyTblMgr) GetTableName() string {
	return "bs_rapidly_tbl"
}

// Reset 重置gorm会话
func (obj *_BsRapidlyTblMgr) Reset() *_BsRapidlyTblMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_BsRapidlyTblMgr) Get() (result BsRapidlyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BsRapidlyTbl{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_BsRapidlyTblMgr) Gets() (results []*BsRapidlyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BsRapidlyTbl{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_BsRapidlyTblMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(BsRapidlyTbl{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_BsRapidlyTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithDay0 day0获取 当日0点时间戳
func (obj *_BsRapidlyTblMgr) WithDay0(day0 int64) Option {
	return optionFunc(func(o *options) { o.query["day0"] = day0 })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_BsRapidlyTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// GetByOption 功能选项模式获取
func (obj *_BsRapidlyTblMgr) GetByOption(opts ...Option) (result BsRapidlyTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(BsRapidlyTbl{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_BsRapidlyTblMgr) GetByOptions(opts ...Option) (results []*BsRapidlyTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(BsRapidlyTbl{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_BsRapidlyTblMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]BsRapidlyTbl, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(BsRapidlyTbl{}).Where(options.query)
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
func (obj *_BsRapidlyTblMgr) GetFromID(id int) (result BsRapidlyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BsRapidlyTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_BsRapidlyTblMgr) GetBatchFromID(ids []int) (results []*BsRapidlyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BsRapidlyTbl{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromDay0 通过day0获取内容 当日0点时间戳
func (obj *_BsRapidlyTblMgr) GetFromDay0(day0 int64) (result BsRapidlyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BsRapidlyTbl{}).Where("`day0` = ?", day0).First(&result).Error

	return
}

// GetBatchFromDay0 批量查找 当日0点时间戳
func (obj *_BsRapidlyTblMgr) GetBatchFromDay0(day0s []int64) (results []*BsRapidlyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BsRapidlyTbl{}).Where("`day0` IN (?)", day0s).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_BsRapidlyTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*BsRapidlyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BsRapidlyTbl{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_BsRapidlyTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*BsRapidlyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BsRapidlyTbl{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_BsRapidlyTblMgr) FetchByPrimaryKey(id int) (result BsRapidlyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BsRapidlyTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByDay0 primary or index 获取唯一内容
func (obj *_BsRapidlyTblMgr) FetchUniqueByDay0(day0 int64) (result BsRapidlyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BsRapidlyTbl{}).Where("`day0` = ?", day0).First(&result).Error

	return
}
