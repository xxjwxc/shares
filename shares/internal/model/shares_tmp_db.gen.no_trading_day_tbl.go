package model

import (
	"context"
	"fmt"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type _NoTradingDayTblMgr struct {
	*_BaseMgr
}

// NoTradingDayTblMgr open func
func NoTradingDayTblMgr(db *gorm.DB) *_NoTradingDayTblMgr {
	if db == nil {
		panic(fmt.Errorf("NoTradingDayTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_NoTradingDayTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("no_trading_day_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_NoTradingDayTblMgr) GetTableName() string {
	return "no_trading_day_tbl"
}

// Reset 重置gorm会话
func (obj *_NoTradingDayTblMgr) Reset() *_NoTradingDayTblMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_NoTradingDayTblMgr) Get() (result NoTradingDayTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NoTradingDayTbl{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_NoTradingDayTblMgr) Gets() (results []*NoTradingDayTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NoTradingDayTbl{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_NoTradingDayTblMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(NoTradingDayTbl{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_NoTradingDayTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithDay day获取 当日0点时间戳
func (obj *_NoTradingDayTblMgr) WithDay(day datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["day"] = day })
}

// GetByOption 功能选项模式获取
func (obj *_NoTradingDayTblMgr) GetByOption(opts ...Option) (result NoTradingDayTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(NoTradingDayTbl{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_NoTradingDayTblMgr) GetByOptions(opts ...Option) (results []*NoTradingDayTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(NoTradingDayTbl{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_NoTradingDayTblMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]NoTradingDayTbl, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(NoTradingDayTbl{}).Where(options.query)
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
func (obj *_NoTradingDayTblMgr) GetFromID(id int) (result NoTradingDayTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NoTradingDayTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_NoTradingDayTblMgr) GetBatchFromID(ids []int) (results []*NoTradingDayTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NoTradingDayTbl{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromDay 通过day获取内容 当日0点时间戳
func (obj *_NoTradingDayTblMgr) GetFromDay(day datatypes.Date) (result NoTradingDayTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NoTradingDayTbl{}).Where("`day` = ?", day).First(&result).Error

	return
}

// GetBatchFromDay 批量查找 当日0点时间戳
func (obj *_NoTradingDayTblMgr) GetBatchFromDay(days []datatypes.Date) (results []*NoTradingDayTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NoTradingDayTbl{}).Where("`day` IN (?)", days).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_NoTradingDayTblMgr) FetchByPrimaryKey(id int) (result NoTradingDayTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NoTradingDayTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByDay primary or index 获取唯一内容
func (obj *_NoTradingDayTblMgr) FetchUniqueByDay(day datatypes.Date) (result NoTradingDayTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(NoTradingDayTbl{}).Where("`day` = ?", day).First(&result).Error

	return
}
