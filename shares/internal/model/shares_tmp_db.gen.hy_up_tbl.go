package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type _HyUpTblMgr struct {
	*_BaseMgr
}

// HyUpTblMgr open func
func HyUpTblMgr(db *gorm.DB) *_HyUpTblMgr {
	if db == nil {
		panic(fmt.Errorf("HyUpTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_HyUpTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("hy_up_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_HyUpTblMgr) GetTableName() string {
	return "hy_up_tbl"
}

// Reset 重置gorm会话
func (obj *_HyUpTblMgr) Reset() *_HyUpTblMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_HyUpTblMgr) Get() (result HyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyUpTbl{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_HyUpTblMgr) Gets() (results []*HyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyUpTbl{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_HyUpTblMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(HyUpTbl{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_HyUpTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取 行业代码
func (obj *_HyUpTblMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithName name获取 行业名
func (obj *_HyUpTblMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithPercent percent获取 百分比
func (obj *_HyUpTblMgr) WithPercent(percent float64) Option {
	return optionFunc(func(o *options) { o.query["percent"] = percent })
}

// WithTurnoverRate turnover_rate获取 换手率
func (obj *_HyUpTblMgr) WithTurnoverRate(turnoverRate float64) Option {
	return optionFunc(func(o *options) { o.query["turnover_rate"] = turnoverRate })
}

// WithNum num获取 总家数
func (obj *_HyUpTblMgr) WithNum(num int) Option {
	return optionFunc(func(o *options) { o.query["num"] = num })
}

// WithUp up获取 上涨家数
func (obj *_HyUpTblMgr) WithUp(up int) Option {
	return optionFunc(func(o *options) { o.query["up"] = up })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_HyUpTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithDay day获取 入池时间
func (obj *_HyUpTblMgr) WithDay(day datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["day"] = day })
}

// WithDayStr day_str获取 入池时间
func (obj *_HyUpTblMgr) WithDayStr(dayStr string) Option {
	return optionFunc(func(o *options) { o.query["day_str"] = dayStr })
}

// GetByOption 功能选项模式获取
func (obj *_HyUpTblMgr) GetByOption(opts ...Option) (result HyUpTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(HyUpTbl{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_HyUpTblMgr) GetByOptions(opts ...Option) (results []*HyUpTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(HyUpTbl{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_HyUpTblMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]HyUpTbl, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(HyUpTbl{}).Where(options.query)
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
func (obj *_HyUpTblMgr) GetFromID(id int) (result HyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyUpTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_HyUpTblMgr) GetBatchFromID(ids []int) (results []*HyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyUpTbl{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 行业代码
func (obj *_HyUpTblMgr) GetFromCode(code string) (results []*HyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyUpTbl{}).Where("`code` = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量查找 行业代码
func (obj *_HyUpTblMgr) GetBatchFromCode(codes []string) (results []*HyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyUpTbl{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 行业名
func (obj *_HyUpTblMgr) GetFromName(name string) (results []*HyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyUpTbl{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 行业名
func (obj *_HyUpTblMgr) GetBatchFromName(names []string) (results []*HyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyUpTbl{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromPercent 通过percent获取内容 百分比
func (obj *_HyUpTblMgr) GetFromPercent(percent float64) (results []*HyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyUpTbl{}).Where("`percent` = ?", percent).Find(&results).Error

	return
}

// GetBatchFromPercent 批量查找 百分比
func (obj *_HyUpTblMgr) GetBatchFromPercent(percents []float64) (results []*HyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyUpTbl{}).Where("`percent` IN (?)", percents).Find(&results).Error

	return
}

// GetFromTurnoverRate 通过turnover_rate获取内容 换手率
func (obj *_HyUpTblMgr) GetFromTurnoverRate(turnoverRate float64) (results []*HyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyUpTbl{}).Where("`turnover_rate` = ?", turnoverRate).Find(&results).Error

	return
}

// GetBatchFromTurnoverRate 批量查找 换手率
func (obj *_HyUpTblMgr) GetBatchFromTurnoverRate(turnoverRates []float64) (results []*HyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyUpTbl{}).Where("`turnover_rate` IN (?)", turnoverRates).Find(&results).Error

	return
}

// GetFromNum 通过num获取内容 总家数
func (obj *_HyUpTblMgr) GetFromNum(num int) (results []*HyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyUpTbl{}).Where("`num` = ?", num).Find(&results).Error

	return
}

// GetBatchFromNum 批量查找 总家数
func (obj *_HyUpTblMgr) GetBatchFromNum(nums []int) (results []*HyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyUpTbl{}).Where("`num` IN (?)", nums).Find(&results).Error

	return
}

// GetFromUp 通过up获取内容 上涨家数
func (obj *_HyUpTblMgr) GetFromUp(up int) (results []*HyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyUpTbl{}).Where("`up` = ?", up).Find(&results).Error

	return
}

// GetBatchFromUp 批量查找 上涨家数
func (obj *_HyUpTblMgr) GetBatchFromUp(ups []int) (results []*HyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyUpTbl{}).Where("`up` IN (?)", ups).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_HyUpTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*HyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyUpTbl{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_HyUpTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*HyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyUpTbl{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromDay 通过day获取内容 入池时间
func (obj *_HyUpTblMgr) GetFromDay(day datatypes.Date) (results []*HyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyUpTbl{}).Where("`day` = ?", day).Find(&results).Error

	return
}

// GetBatchFromDay 批量查找 入池时间
func (obj *_HyUpTblMgr) GetBatchFromDay(days []datatypes.Date) (results []*HyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyUpTbl{}).Where("`day` IN (?)", days).Find(&results).Error

	return
}

// GetFromDayStr 通过day_str获取内容 入池时间
func (obj *_HyUpTblMgr) GetFromDayStr(dayStr string) (results []*HyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyUpTbl{}).Where("`day_str` = ?", dayStr).Find(&results).Error

	return
}

// GetBatchFromDayStr 批量查找 入池时间
func (obj *_HyUpTblMgr) GetBatchFromDayStr(dayStrs []string) (results []*HyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyUpTbl{}).Where("`day_str` IN (?)", dayStrs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_HyUpTblMgr) FetchByPrimaryKey(id int) (result HyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyUpTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueIndexByName primary or index 获取唯一内容
func (obj *_HyUpTblMgr) FetchUniqueIndexByName(code string, day datatypes.Date) (result HyUpTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyUpTbl{}).Where("`code` = ? AND `day` = ?", code, day).First(&result).Error

	return
}
