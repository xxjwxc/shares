package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AnalyTdxAfterViewMgr struct {
	*_BaseMgr
}

// AnalyTdxAfterViewMgr open func
func AnalyTdxAfterViewMgr(db *gorm.DB) *_AnalyTdxAfterViewMgr {
	if db == nil {
		panic(fmt.Errorf("AnalyTdxAfterViewMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AnalyTdxAfterViewMgr{_BaseMgr: &_BaseMgr{DB: db.Table("analy_tdx_after_view"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AnalyTdxAfterViewMgr) GetTableName() string {
	return "analy_tdx_after_view"
}

// Reset 重置gorm会话
func (obj *_AnalyTdxAfterViewMgr) Reset() *_AnalyTdxAfterViewMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AnalyTdxAfterViewMgr) Get() (result AnalyTdxAfterView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyTdxAfterView{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AnalyTdxAfterViewMgr) Gets() (results []*AnalyTdxAfterView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyTdxAfterView{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AnalyTdxAfterViewMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AnalyTdxAfterView{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// GetByOption 功能选项模式获取
func (obj *_AnalyTdxAfterViewMgr) GetByOption(opts ...Option) (result AnalyTdxAfterView, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AnalyTdxAfterView{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AnalyTdxAfterViewMgr) GetByOptions(opts ...Option) (results []*AnalyTdxAfterView, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AnalyTdxAfterView{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AnalyTdxAfterViewMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AnalyTdxAfterView, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AnalyTdxAfterView{}).Where(options.query)
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

//////////////////////////primary index case ////////////////////////////////////////////
