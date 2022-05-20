package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _GroupWatchTblMgr struct {
	*_BaseMgr
}

// GroupWatchTblMgr open func
func GroupWatchTblMgr(db *gorm.DB) *_GroupWatchTblMgr {
	if db == nil {
		panic(fmt.Errorf("GroupWatchTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_GroupWatchTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("group_watch_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_GroupWatchTblMgr) GetTableName() string {
	return "group_watch_tbl"
}

// Reset 重置gorm会话
func (obj *_GroupWatchTblMgr) Reset() *_GroupWatchTblMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_GroupWatchTblMgr) Get() (result GroupWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupWatchTbl{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_GroupWatchTblMgr) Gets() (results []*GroupWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupWatchTbl{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_GroupWatchTblMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(GroupWatchTbl{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_GroupWatchTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取 股票代码
func (obj *_GroupWatchTblMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithName name获取 股票名字
func (obj *_GroupWatchTblMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithPrice price获取 价格
func (obj *_GroupWatchTblMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithNewPrice new_price获取 最新价格
func (obj *_GroupWatchTblMgr) WithNewPrice(newPrice float64) Option {
	return optionFunc(func(o *options) { o.query["new_price"] = newPrice })
}

// WithDesc desc获取 描述
func (obj *_GroupWatchTblMgr) WithDesc(desc string) Option {
	return optionFunc(func(o *options) { o.query["desc"] = desc })
}

// WithUserName user_name获取 推荐人
func (obj *_GroupWatchTblMgr) WithUserName(userName string) Option {
	return optionFunc(func(o *options) { o.query["user_name"] = userName })
}

// WithCreatedBy created_by获取
func (obj *_GroupWatchTblMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithCreatedAt created_at获取 入池时间
func (obj *_GroupWatchTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// GetByOption 功能选项模式获取
func (obj *_GroupWatchTblMgr) GetByOption(opts ...Option) (result GroupWatchTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(GroupWatchTbl{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_GroupWatchTblMgr) GetByOptions(opts ...Option) (results []*GroupWatchTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(GroupWatchTbl{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_GroupWatchTblMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]GroupWatchTbl, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(GroupWatchTbl{}).Where(options.query)
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
func (obj *_GroupWatchTblMgr) GetFromID(id int) (result GroupWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupWatchTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_GroupWatchTblMgr) GetBatchFromID(ids []int) (results []*GroupWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupWatchTbl{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 股票代码
func (obj *_GroupWatchTblMgr) GetFromCode(code string) (result GroupWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupWatchTbl{}).Where("`code` = ?", code).First(&result).Error

	return
}

// GetBatchFromCode 批量查找 股票代码
func (obj *_GroupWatchTblMgr) GetBatchFromCode(codes []string) (results []*GroupWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupWatchTbl{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 股票名字
func (obj *_GroupWatchTblMgr) GetFromName(name string) (results []*GroupWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupWatchTbl{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 股票名字
func (obj *_GroupWatchTblMgr) GetBatchFromName(names []string) (results []*GroupWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupWatchTbl{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 价格
func (obj *_GroupWatchTblMgr) GetFromPrice(price float64) (results []*GroupWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupWatchTbl{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 价格
func (obj *_GroupWatchTblMgr) GetBatchFromPrice(prices []float64) (results []*GroupWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupWatchTbl{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromNewPrice 通过new_price获取内容 最新价格
func (obj *_GroupWatchTblMgr) GetFromNewPrice(newPrice float64) (results []*GroupWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupWatchTbl{}).Where("`new_price` = ?", newPrice).Find(&results).Error

	return
}

// GetBatchFromNewPrice 批量查找 最新价格
func (obj *_GroupWatchTblMgr) GetBatchFromNewPrice(newPrices []float64) (results []*GroupWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupWatchTbl{}).Where("`new_price` IN (?)", newPrices).Find(&results).Error

	return
}

// GetFromDesc 通过desc获取内容 描述
func (obj *_GroupWatchTblMgr) GetFromDesc(desc string) (results []*GroupWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupWatchTbl{}).Where("`desc` = ?", desc).Find(&results).Error

	return
}

// GetBatchFromDesc 批量查找 描述
func (obj *_GroupWatchTblMgr) GetBatchFromDesc(descs []string) (results []*GroupWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupWatchTbl{}).Where("`desc` IN (?)", descs).Find(&results).Error

	return
}

// GetFromUserName 通过user_name获取内容 推荐人
func (obj *_GroupWatchTblMgr) GetFromUserName(userName string) (results []*GroupWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupWatchTbl{}).Where("`user_name` = ?", userName).Find(&results).Error

	return
}

// GetBatchFromUserName 批量查找 推荐人
func (obj *_GroupWatchTblMgr) GetBatchFromUserName(userNames []string) (results []*GroupWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupWatchTbl{}).Where("`user_name` IN (?)", userNames).Find(&results).Error

	return
}

// GetFromCreatedBy 通过created_by获取内容
func (obj *_GroupWatchTblMgr) GetFromCreatedBy(createdBy string) (results []*GroupWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupWatchTbl{}).Where("`created_by` = ?", createdBy).Find(&results).Error

	return
}

// GetBatchFromCreatedBy 批量查找
func (obj *_GroupWatchTblMgr) GetBatchFromCreatedBy(createdBys []string) (results []*GroupWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupWatchTbl{}).Where("`created_by` IN (?)", createdBys).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 入池时间
func (obj *_GroupWatchTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*GroupWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupWatchTbl{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 入池时间
func (obj *_GroupWatchTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*GroupWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupWatchTbl{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_GroupWatchTblMgr) FetchByPrimaryKey(id int) (result GroupWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupWatchTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByCode primary or index 获取唯一内容
func (obj *_GroupWatchTblMgr) FetchUniqueByCode(code string) (result GroupWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupWatchTbl{}).Where("`code` = ?", code).First(&result).Error

	return
}
