package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type _MsgTblMgr struct {
	*_BaseMgr
}

// MsgTblMgr open func
func MsgTblMgr(db *gorm.DB) *_MsgTblMgr {
	if db == nil {
		panic(fmt.Errorf("MsgTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MsgTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("msg_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MsgTblMgr) GetTableName() string {
	return "msg_tbl"
}

// Reset 重置gorm会话
func (obj *_MsgTblMgr) Reset() *_MsgTblMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MsgTblMgr) Get() (result MsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgTbl{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MsgTblMgr) Gets() (results []*MsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgTbl{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MsgTblMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(MsgTbl{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_MsgTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOpenID open_id获取 用户openid
func (obj *_MsgTblMgr) WithOpenID(openID string) Option {
	return optionFunc(func(o *options) { o.query["open_id"] = openID })
}

// WithCode code获取 股票代码
func (obj *_MsgTblMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithKey key获取 消息类型
func (obj *_MsgTblMgr) WithKey(key string) Option {
	return optionFunc(func(o *options) { o.query["key"] = key })
}

// WithDesc desc获取 消息描述
func (obj *_MsgTblMgr) WithDesc(desc string) Option {
	return optionFunc(func(o *options) { o.query["desc"] = desc })
}

// WithPrice price获取 当前价格
func (obj *_MsgTblMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithPercent percent获取 百分比
func (obj *_MsgTblMgr) WithPercent(percent float64) Option {
	return optionFunc(func(o *options) { o.query["percent"] = percent })
}

// WithDay day获取 当日0点时间戳
func (obj *_MsgTblMgr) WithDay(day datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["day"] = day })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_MsgTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// GetByOption 功能选项模式获取
func (obj *_MsgTblMgr) GetByOption(opts ...Option) (result MsgTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MsgTbl{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MsgTblMgr) GetByOptions(opts ...Option) (results []*MsgTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MsgTbl{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_MsgTblMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]MsgTbl, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(MsgTbl{}).Where(options.query)
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
func (obj *_MsgTblMgr) GetFromID(id int) (result MsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_MsgTblMgr) GetBatchFromID(ids []int) (results []*MsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgTbl{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOpenID 通过open_id获取内容 用户openid
func (obj *_MsgTblMgr) GetFromOpenID(openID string) (results []*MsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgTbl{}).Where("`open_id` = ?", openID).Find(&results).Error

	return
}

// GetBatchFromOpenID 批量查找 用户openid
func (obj *_MsgTblMgr) GetBatchFromOpenID(openIDs []string) (results []*MsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgTbl{}).Where("`open_id` IN (?)", openIDs).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 股票代码
func (obj *_MsgTblMgr) GetFromCode(code string) (results []*MsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgTbl{}).Where("`code` = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量查找 股票代码
func (obj *_MsgTblMgr) GetBatchFromCode(codes []string) (results []*MsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgTbl{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromKey 通过key获取内容 消息类型
func (obj *_MsgTblMgr) GetFromKey(key string) (results []*MsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgTbl{}).Where("`key` = ?", key).Find(&results).Error

	return
}

// GetBatchFromKey 批量查找 消息类型
func (obj *_MsgTblMgr) GetBatchFromKey(keys []string) (results []*MsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgTbl{}).Where("`key` IN (?)", keys).Find(&results).Error

	return
}

// GetFromDesc 通过desc获取内容 消息描述
func (obj *_MsgTblMgr) GetFromDesc(desc string) (results []*MsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgTbl{}).Where("`desc` = ?", desc).Find(&results).Error

	return
}

// GetBatchFromDesc 批量查找 消息描述
func (obj *_MsgTblMgr) GetBatchFromDesc(descs []string) (results []*MsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgTbl{}).Where("`desc` IN (?)", descs).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 当前价格
func (obj *_MsgTblMgr) GetFromPrice(price float64) (results []*MsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgTbl{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 当前价格
func (obj *_MsgTblMgr) GetBatchFromPrice(prices []float64) (results []*MsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgTbl{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromPercent 通过percent获取内容 百分比
func (obj *_MsgTblMgr) GetFromPercent(percent float64) (results []*MsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgTbl{}).Where("`percent` = ?", percent).Find(&results).Error

	return
}

// GetBatchFromPercent 批量查找 百分比
func (obj *_MsgTblMgr) GetBatchFromPercent(percents []float64) (results []*MsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgTbl{}).Where("`percent` IN (?)", percents).Find(&results).Error

	return
}

// GetFromDay 通过day获取内容 当日0点时间戳
func (obj *_MsgTblMgr) GetFromDay(day datatypes.Date) (results []*MsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgTbl{}).Where("`day` = ?", day).Find(&results).Error

	return
}

// GetBatchFromDay 批量查找 当日0点时间戳
func (obj *_MsgTblMgr) GetBatchFromDay(days []datatypes.Date) (results []*MsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgTbl{}).Where("`day` IN (?)", days).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_MsgTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*MsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgTbl{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_MsgTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*MsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgTbl{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MsgTblMgr) FetchByPrimaryKey(id int) (result MsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByCode  获取多个内容
func (obj *_MsgTblMgr) FetchIndexByCode(code string) (results []*MsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgTbl{}).Where("`code` = ?", code).Find(&results).Error

	return
}
