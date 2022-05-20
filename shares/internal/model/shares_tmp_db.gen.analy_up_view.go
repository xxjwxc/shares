package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AnalyUpViewMgr struct {
	*_BaseMgr
}

// AnalyUpViewMgr open func
func AnalyUpViewMgr(db *gorm.DB) *_AnalyUpViewMgr {
	if db == nil {
		panic(fmt.Errorf("AnalyUpViewMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AnalyUpViewMgr{_BaseMgr: &_BaseMgr{DB: db.Table("analy_up_view"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AnalyUpViewMgr) GetTableName() string {
	return "analy_up_view"
}

// Reset 重置gorm会话
func (obj *_AnalyUpViewMgr) Reset() *_AnalyUpViewMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AnalyUpViewMgr) Get() (result AnalyUpView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyUpView{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AnalyUpViewMgr) Gets() (results []*AnalyUpView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyUpView{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AnalyUpViewMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AnalyUpView{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// GetByOption 功能选项模式获取
func (obj *_AnalyUpViewMgr) GetByOption(opts ...Option) (result AnalyUpView, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AnalyUpView{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AnalyUpViewMgr) GetByOptions(opts ...Option) (results []*AnalyUpView, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AnalyUpView{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AnalyUpViewMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AnalyUpView, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AnalyUpView{}).Where(options.query)
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
