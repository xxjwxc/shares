package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _LogTblMgr struct {
	*_BaseMgr
}

// LogTblMgr open func
func LogTblMgr(db *gorm.DB) *_LogTblMgr {
	if db == nil {
		panic(fmt.Errorf("LogTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LogTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("log_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LogTblMgr) GetTableName() string {
	return "log_tbl"
}

// Reset 重置gorm会话
func (obj *_LogTblMgr) Reset() *_LogTblMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LogTblMgr) Get() (result LogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LogTbl{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LogTblMgr) Gets() (results []*LogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LogTbl{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LogTblMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LogTbl{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_LogTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithDesc desc获取 日志内容
func (obj *_LogTblMgr) WithDesc(desc string) Option {
	return optionFunc(func(o *options) { o.query["desc"] = desc })
}

// GetByOption 功能选项模式获取
func (obj *_LogTblMgr) GetByOption(opts ...Option) (result LogTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LogTbl{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LogTblMgr) GetByOptions(opts ...Option) (results []*LogTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LogTbl{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LogTblMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LogTbl, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LogTbl{}).Where(options.query)
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
func (obj *_LogTblMgr) GetFromID(id int) (result LogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LogTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_LogTblMgr) GetBatchFromID(ids []int) (results []*LogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LogTbl{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromDesc 通过desc获取内容 日志内容
func (obj *_LogTblMgr) GetFromDesc(desc string) (results []*LogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LogTbl{}).Where("`desc` = ?", desc).Find(&results).Error

	return
}

// GetBatchFromDesc 批量查找 日志内容
func (obj *_LogTblMgr) GetBatchFromDesc(descs []string) (results []*LogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LogTbl{}).Where("`desc` IN (?)", descs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LogTblMgr) FetchByPrimaryKey(id int) (result LogTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LogTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}
