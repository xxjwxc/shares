package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type _XqMsgDailyActiveTblMgr struct {
	*_BaseMgr
}

// XqMsgDailyActiveTblMgr open func
func XqMsgDailyActiveTblMgr(db *gorm.DB) *_XqMsgDailyActiveTblMgr {
	if db == nil {
		panic(fmt.Errorf("XqMsgDailyActiveTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_XqMsgDailyActiveTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("xq_msg_daily_active_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_XqMsgDailyActiveTblMgr) GetTableName() string {
	return "xq_msg_daily_active_tbl"
}

// Reset 重置gorm会话
func (obj *_XqMsgDailyActiveTblMgr) Reset() *_XqMsgDailyActiveTblMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_XqMsgDailyActiveTblMgr) Get() (result XqMsgDailyActiveTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyActiveTbl{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_XqMsgDailyActiveTblMgr) Gets() (results []*XqMsgDailyActiveTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyActiveTbl{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_XqMsgDailyActiveTblMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(XqMsgDailyActiveTbl{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_XqMsgDailyActiveTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取 股票代码
func (obj *_XqMsgDailyActiveTblMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithName name获取 股票名字
func (obj *_XqMsgDailyActiveTblMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDay0 day0获取 入池时间
func (obj *_XqMsgDailyActiveTblMgr) WithDay0(day0 datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["day0"] = day0 })
}

// WithDayStr day_str获取 入池时间
func (obj *_XqMsgDailyActiveTblMgr) WithDayStr(dayStr string) Option {
	return optionFunc(func(o *options) { o.query["day_str"] = dayStr })
}

// WithPrice price获取 价格
func (obj *_XqMsgDailyActiveTblMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithPercent percent获取 百分比
func (obj *_XqMsgDailyActiveTblMgr) WithPercent(percent float64) Option {
	return optionFunc(func(o *options) { o.query["percent"] = percent })
}

// WithTurnoverRate turnover_rate获取 换手率
func (obj *_XqMsgDailyActiveTblMgr) WithTurnoverRate(turnoverRate float64) Option {
	return optionFunc(func(o *options) { o.query["turnover_rate"] = turnoverRate })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_XqMsgDailyActiveTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithCount count获取 次数
func (obj *_XqMsgDailyActiveTblMgr) WithCount(count int) Option {
	return optionFunc(func(o *options) { o.query["count"] = count })
}

// GetByOption 功能选项模式获取
func (obj *_XqMsgDailyActiveTblMgr) GetByOption(opts ...Option) (result XqMsgDailyActiveTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyActiveTbl{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_XqMsgDailyActiveTblMgr) GetByOptions(opts ...Option) (results []*XqMsgDailyActiveTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyActiveTbl{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_XqMsgDailyActiveTblMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]XqMsgDailyActiveTbl, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(XqMsgDailyActiveTbl{}).Where(options.query)
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
func (obj *_XqMsgDailyActiveTblMgr) GetFromID(id int) (result XqMsgDailyActiveTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyActiveTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_XqMsgDailyActiveTblMgr) GetBatchFromID(ids []int) (results []*XqMsgDailyActiveTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyActiveTbl{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 股票代码
func (obj *_XqMsgDailyActiveTblMgr) GetFromCode(code string) (results []*XqMsgDailyActiveTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyActiveTbl{}).Where("`code` = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量查找 股票代码
func (obj *_XqMsgDailyActiveTblMgr) GetBatchFromCode(codes []string) (results []*XqMsgDailyActiveTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyActiveTbl{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 股票名字
func (obj *_XqMsgDailyActiveTblMgr) GetFromName(name string) (results []*XqMsgDailyActiveTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyActiveTbl{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 股票名字
func (obj *_XqMsgDailyActiveTblMgr) GetBatchFromName(names []string) (results []*XqMsgDailyActiveTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyActiveTbl{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromDay0 通过day0获取内容 入池时间
func (obj *_XqMsgDailyActiveTblMgr) GetFromDay0(day0 datatypes.Date) (results []*XqMsgDailyActiveTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyActiveTbl{}).Where("`day0` = ?", day0).Find(&results).Error

	return
}

// GetBatchFromDay0 批量查找 入池时间
func (obj *_XqMsgDailyActiveTblMgr) GetBatchFromDay0(day0s []datatypes.Date) (results []*XqMsgDailyActiveTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyActiveTbl{}).Where("`day0` IN (?)", day0s).Find(&results).Error

	return
}

// GetFromDayStr 通过day_str获取内容 入池时间
func (obj *_XqMsgDailyActiveTblMgr) GetFromDayStr(dayStr string) (results []*XqMsgDailyActiveTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyActiveTbl{}).Where("`day_str` = ?", dayStr).Find(&results).Error

	return
}

// GetBatchFromDayStr 批量查找 入池时间
func (obj *_XqMsgDailyActiveTblMgr) GetBatchFromDayStr(dayStrs []string) (results []*XqMsgDailyActiveTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyActiveTbl{}).Where("`day_str` IN (?)", dayStrs).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 价格
func (obj *_XqMsgDailyActiveTblMgr) GetFromPrice(price float64) (results []*XqMsgDailyActiveTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyActiveTbl{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 价格
func (obj *_XqMsgDailyActiveTblMgr) GetBatchFromPrice(prices []float64) (results []*XqMsgDailyActiveTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyActiveTbl{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromPercent 通过percent获取内容 百分比
func (obj *_XqMsgDailyActiveTblMgr) GetFromPercent(percent float64) (results []*XqMsgDailyActiveTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyActiveTbl{}).Where("`percent` = ?", percent).Find(&results).Error

	return
}

// GetBatchFromPercent 批量查找 百分比
func (obj *_XqMsgDailyActiveTblMgr) GetBatchFromPercent(percents []float64) (results []*XqMsgDailyActiveTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyActiveTbl{}).Where("`percent` IN (?)", percents).Find(&results).Error

	return
}

// GetFromTurnoverRate 通过turnover_rate获取内容 换手率
func (obj *_XqMsgDailyActiveTblMgr) GetFromTurnoverRate(turnoverRate float64) (results []*XqMsgDailyActiveTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyActiveTbl{}).Where("`turnover_rate` = ?", turnoverRate).Find(&results).Error

	return
}

// GetBatchFromTurnoverRate 批量查找 换手率
func (obj *_XqMsgDailyActiveTblMgr) GetBatchFromTurnoverRate(turnoverRates []float64) (results []*XqMsgDailyActiveTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyActiveTbl{}).Where("`turnover_rate` IN (?)", turnoverRates).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_XqMsgDailyActiveTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*XqMsgDailyActiveTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyActiveTbl{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_XqMsgDailyActiveTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*XqMsgDailyActiveTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyActiveTbl{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromCount 通过count获取内容 次数
func (obj *_XqMsgDailyActiveTblMgr) GetFromCount(count int) (results []*XqMsgDailyActiveTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyActiveTbl{}).Where("`count` = ?", count).Find(&results).Error

	return
}

// GetBatchFromCount 批量查找 次数
func (obj *_XqMsgDailyActiveTblMgr) GetBatchFromCount(counts []int) (results []*XqMsgDailyActiveTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyActiveTbl{}).Where("`count` IN (?)", counts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_XqMsgDailyActiveTblMgr) FetchByPrimaryKey(id int) (result XqMsgDailyActiveTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyActiveTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueIndexByCode primary or index 获取唯一内容
func (obj *_XqMsgDailyActiveTblMgr) FetchUniqueIndexByCode(code string, day0 datatypes.Date) (result XqMsgDailyActiveTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyActiveTbl{}).Where("`code` = ? AND `day0` = ?", code, day0).First(&result).Error

	return
}
