package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type _MsgRapidlyTblMgr struct {
	*_BaseMgr
}

// MsgRapidlyTblMgr open func
func MsgRapidlyTblMgr(db *gorm.DB) *_MsgRapidlyTblMgr {
	if db == nil {
		panic(fmt.Errorf("MsgRapidlyTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MsgRapidlyTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("msg_rapidly_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MsgRapidlyTblMgr) GetTableName() string {
	return "msg_rapidly_tbl"
}

// Reset 重置gorm会话
func (obj *_MsgRapidlyTblMgr) Reset() *_MsgRapidlyTblMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MsgRapidlyTblMgr) Get() (result MsgRapidlyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgRapidlyTbl{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MsgRapidlyTblMgr) Gets() (results []*MsgRapidlyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgRapidlyTbl{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MsgRapidlyTblMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(MsgRapidlyTbl{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_MsgRapidlyTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取 股票代码
func (obj *_MsgRapidlyTblMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithTag tag获取 1:上涨，0:下跌
func (obj *_MsgRapidlyTblMgr) WithTag(tag int) Option {
	return optionFunc(func(o *options) { o.query["tag"] = tag })
}

// WithKey key获取 消息类型
func (obj *_MsgRapidlyTblMgr) WithKey(key string) Option {
	return optionFunc(func(o *options) { o.query["key"] = key })
}

// WithDesc desc获取 消息描述
func (obj *_MsgRapidlyTblMgr) WithDesc(desc string) Option {
	return optionFunc(func(o *options) { o.query["desc"] = desc })
}

// WithOld old获取 之前价格
func (obj *_MsgRapidlyTblMgr) WithOld(old float64) Option {
	return optionFunc(func(o *options) { o.query["old"] = old })
}

// WithNew new获取 当前价格
func (obj *_MsgRapidlyTblMgr) WithNew(new float64) Option {
	return optionFunc(func(o *options) { o.query["new"] = new })
}

// WithPercent percent获取 百分比
func (obj *_MsgRapidlyTblMgr) WithPercent(percent float64) Option {
	return optionFunc(func(o *options) { o.query["percent"] = percent })
}

// WithOffsetPercent offset_percent获取 相对百分比
func (obj *_MsgRapidlyTblMgr) WithOffsetPercent(offsetPercent float64) Option {
	return optionFunc(func(o *options) { o.query["offset_percent"] = offsetPercent })
}

// WithDay day获取 当日0点时间戳
func (obj *_MsgRapidlyTblMgr) WithDay(day datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["day"] = day })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_MsgRapidlyTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// GetByOption 功能选项模式获取
func (obj *_MsgRapidlyTblMgr) GetByOption(opts ...Option) (result MsgRapidlyTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MsgRapidlyTbl{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MsgRapidlyTblMgr) GetByOptions(opts ...Option) (results []*MsgRapidlyTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MsgRapidlyTbl{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_MsgRapidlyTblMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]MsgRapidlyTbl, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(MsgRapidlyTbl{}).Where(options.query)
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
func (obj *_MsgRapidlyTblMgr) GetFromID(id int) (result MsgRapidlyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgRapidlyTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_MsgRapidlyTblMgr) GetBatchFromID(ids []int) (results []*MsgRapidlyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgRapidlyTbl{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 股票代码
func (obj *_MsgRapidlyTblMgr) GetFromCode(code string) (results []*MsgRapidlyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgRapidlyTbl{}).Where("`code` = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量查找 股票代码
func (obj *_MsgRapidlyTblMgr) GetBatchFromCode(codes []string) (results []*MsgRapidlyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgRapidlyTbl{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromTag 通过tag获取内容 1:上涨，0:下跌
func (obj *_MsgRapidlyTblMgr) GetFromTag(tag int) (results []*MsgRapidlyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgRapidlyTbl{}).Where("`tag` = ?", tag).Find(&results).Error

	return
}

// GetBatchFromTag 批量查找 1:上涨，0:下跌
func (obj *_MsgRapidlyTblMgr) GetBatchFromTag(tags []int) (results []*MsgRapidlyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgRapidlyTbl{}).Where("`tag` IN (?)", tags).Find(&results).Error

	return
}

// GetFromKey 通过key获取内容 消息类型
func (obj *_MsgRapidlyTblMgr) GetFromKey(key string) (results []*MsgRapidlyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgRapidlyTbl{}).Where("`key` = ?", key).Find(&results).Error

	return
}

// GetBatchFromKey 批量查找 消息类型
func (obj *_MsgRapidlyTblMgr) GetBatchFromKey(keys []string) (results []*MsgRapidlyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgRapidlyTbl{}).Where("`key` IN (?)", keys).Find(&results).Error

	return
}

// GetFromDesc 通过desc获取内容 消息描述
func (obj *_MsgRapidlyTblMgr) GetFromDesc(desc string) (results []*MsgRapidlyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgRapidlyTbl{}).Where("`desc` = ?", desc).Find(&results).Error

	return
}

// GetBatchFromDesc 批量查找 消息描述
func (obj *_MsgRapidlyTblMgr) GetBatchFromDesc(descs []string) (results []*MsgRapidlyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgRapidlyTbl{}).Where("`desc` IN (?)", descs).Find(&results).Error

	return
}

// GetFromOld 通过old获取内容 之前价格
func (obj *_MsgRapidlyTblMgr) GetFromOld(old float64) (results []*MsgRapidlyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgRapidlyTbl{}).Where("`old` = ?", old).Find(&results).Error

	return
}

// GetBatchFromOld 批量查找 之前价格
func (obj *_MsgRapidlyTblMgr) GetBatchFromOld(olds []float64) (results []*MsgRapidlyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgRapidlyTbl{}).Where("`old` IN (?)", olds).Find(&results).Error

	return
}

// GetFromNew 通过new获取内容 当前价格
func (obj *_MsgRapidlyTblMgr) GetFromNew(new float64) (results []*MsgRapidlyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgRapidlyTbl{}).Where("`new` = ?", new).Find(&results).Error

	return
}

// GetBatchFromNew 批量查找 当前价格
func (obj *_MsgRapidlyTblMgr) GetBatchFromNew(news []float64) (results []*MsgRapidlyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgRapidlyTbl{}).Where("`new` IN (?)", news).Find(&results).Error

	return
}

// GetFromPercent 通过percent获取内容 百分比
func (obj *_MsgRapidlyTblMgr) GetFromPercent(percent float64) (results []*MsgRapidlyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgRapidlyTbl{}).Where("`percent` = ?", percent).Find(&results).Error

	return
}

// GetBatchFromPercent 批量查找 百分比
func (obj *_MsgRapidlyTblMgr) GetBatchFromPercent(percents []float64) (results []*MsgRapidlyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgRapidlyTbl{}).Where("`percent` IN (?)", percents).Find(&results).Error

	return
}

// GetFromOffsetPercent 通过offset_percent获取内容 相对百分比
func (obj *_MsgRapidlyTblMgr) GetFromOffsetPercent(offsetPercent float64) (results []*MsgRapidlyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgRapidlyTbl{}).Where("`offset_percent` = ?", offsetPercent).Find(&results).Error

	return
}

// GetBatchFromOffsetPercent 批量查找 相对百分比
func (obj *_MsgRapidlyTblMgr) GetBatchFromOffsetPercent(offsetPercents []float64) (results []*MsgRapidlyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgRapidlyTbl{}).Where("`offset_percent` IN (?)", offsetPercents).Find(&results).Error

	return
}

// GetFromDay 通过day获取内容 当日0点时间戳
func (obj *_MsgRapidlyTblMgr) GetFromDay(day datatypes.Date) (results []*MsgRapidlyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgRapidlyTbl{}).Where("`day` = ?", day).Find(&results).Error

	return
}

// GetBatchFromDay 批量查找 当日0点时间戳
func (obj *_MsgRapidlyTblMgr) GetBatchFromDay(days []datatypes.Date) (results []*MsgRapidlyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgRapidlyTbl{}).Where("`day` IN (?)", days).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_MsgRapidlyTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*MsgRapidlyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgRapidlyTbl{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_MsgRapidlyTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*MsgRapidlyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgRapidlyTbl{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MsgRapidlyTblMgr) FetchByPrimaryKey(id int) (result MsgRapidlyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgRapidlyTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueIndexByCode primary or index 获取唯一内容
func (obj *_MsgRapidlyTblMgr) FetchUniqueIndexByCode(code string, tag int, day datatypes.Date) (result MsgRapidlyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgRapidlyTbl{}).Where("`code` = ? AND `tag` = ? AND `day` = ?", code, tag, day).First(&result).Error

	return
}
