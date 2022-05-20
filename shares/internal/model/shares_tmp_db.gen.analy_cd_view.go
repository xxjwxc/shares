package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AnalyCdViewMgr struct {
	*_BaseMgr
}

// AnalyCdViewMgr open func
func AnalyCdViewMgr(db *gorm.DB) *_AnalyCdViewMgr {
	if db == nil {
		panic(fmt.Errorf("AnalyCdViewMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AnalyCdViewMgr{_BaseMgr: &_BaseMgr{DB: db.Table("analy_cd_view"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AnalyCdViewMgr) GetTableName() string {
	return "analy_cd_view"
}

// Reset 重置gorm会话
func (obj *_AnalyCdViewMgr) Reset() *_AnalyCdViewMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AnalyCdViewMgr) Get() (result AnalyCdView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyCdView{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AnalyCdViewMgr) Gets() (results []*AnalyCdView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyCdView{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AnalyCdViewMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AnalyCdView{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// GetByOption 功能选项模式获取
func (obj *_AnalyCdViewMgr) GetByOption(opts ...Option) (result AnalyCdView, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AnalyCdView{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AnalyCdViewMgr) GetByOptions(opts ...Option) (results []*AnalyCdView, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AnalyCdView{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AnalyCdViewMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AnalyCdView, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AnalyCdView{}).Where(options.query)
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
