package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type _AnalyHyTblMgr struct {
	*_BaseMgr
}

// AnalyHyTblMgr open func
func AnalyHyTblMgr(db *gorm.DB) *_AnalyHyTblMgr {
	if db == nil {
		panic(fmt.Errorf("AnalyHyTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AnalyHyTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("analy_hy_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AnalyHyTblMgr) GetTableName() string {
	return "analy_hy_tbl"
}

// Reset 重置gorm会话
func (obj *_AnalyHyTblMgr) Reset() *_AnalyHyTblMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AnalyHyTblMgr) Get() (result AnalyHyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyTbl{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AnalyHyTblMgr) Gets() (results []*AnalyHyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyTbl{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AnalyHyTblMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AnalyHyTbl{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_AnalyHyTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取 行业代码
func (obj *_AnalyHyTblMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithName name获取 行业名
func (obj *_AnalyHyTblMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithPercent percent获取 百分比
func (obj *_AnalyHyTblMgr) WithPercent(percent float64) Option {
	return optionFunc(func(o *options) { o.query["percent"] = percent })
}

// WithTurnoverRate turnover_rate获取 换手率
func (obj *_AnalyHyTblMgr) WithTurnoverRate(turnoverRate float64) Option {
	return optionFunc(func(o *options) { o.query["turnover_rate"] = turnoverRate })
}

// WithNum num获取 总家数
func (obj *_AnalyHyTblMgr) WithNum(num int) Option {
	return optionFunc(func(o *options) { o.query["num"] = num })
}

// WithUp up获取 上涨家数
func (obj *_AnalyHyTblMgr) WithUp(up int) Option {
	return optionFunc(func(o *options) { o.query["up"] = up })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_AnalyHyTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithDay day获取 入池时间
func (obj *_AnalyHyTblMgr) WithDay(day datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["day"] = day })
}

// WithDayStr day_str获取 入池时间
func (obj *_AnalyHyTblMgr) WithDayStr(dayStr string) Option {
	return optionFunc(func(o *options) { o.query["day_str"] = dayStr })
}

// GetByOption 功能选项模式获取
func (obj *_AnalyHyTblMgr) GetByOption(opts ...Option) (result AnalyHyTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyTbl{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AnalyHyTblMgr) GetByOptions(opts ...Option) (results []*AnalyHyTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyTbl{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AnalyHyTblMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AnalyHyTbl, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AnalyHyTbl{}).Where(options.query)
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
func (obj *_AnalyHyTblMgr) GetFromID(id int) (result AnalyHyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_AnalyHyTblMgr) GetBatchFromID(ids []int) (results []*AnalyHyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyTbl{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 行业代码
func (obj *_AnalyHyTblMgr) GetFromCode(code string) (results []*AnalyHyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyTbl{}).Where("`code` = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量查找 行业代码
func (obj *_AnalyHyTblMgr) GetBatchFromCode(codes []string) (results []*AnalyHyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyTbl{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 行业名
func (obj *_AnalyHyTblMgr) GetFromName(name string) (results []*AnalyHyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyTbl{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 行业名
func (obj *_AnalyHyTblMgr) GetBatchFromName(names []string) (results []*AnalyHyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyTbl{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromPercent 通过percent获取内容 百分比
func (obj *_AnalyHyTblMgr) GetFromPercent(percent float64) (results []*AnalyHyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyTbl{}).Where("`percent` = ?", percent).Find(&results).Error

	return
}

// GetBatchFromPercent 批量查找 百分比
func (obj *_AnalyHyTblMgr) GetBatchFromPercent(percents []float64) (results []*AnalyHyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyTbl{}).Where("`percent` IN (?)", percents).Find(&results).Error

	return
}

// GetFromTurnoverRate 通过turnover_rate获取内容 换手率
func (obj *_AnalyHyTblMgr) GetFromTurnoverRate(turnoverRate float64) (results []*AnalyHyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyTbl{}).Where("`turnover_rate` = ?", turnoverRate).Find(&results).Error

	return
}

// GetBatchFromTurnoverRate 批量查找 换手率
func (obj *_AnalyHyTblMgr) GetBatchFromTurnoverRate(turnoverRates []float64) (results []*AnalyHyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyTbl{}).Where("`turnover_rate` IN (?)", turnoverRates).Find(&results).Error

	return
}

// GetFromNum 通过num获取内容 总家数
func (obj *_AnalyHyTblMgr) GetFromNum(num int) (results []*AnalyHyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyTbl{}).Where("`num` = ?", num).Find(&results).Error

	return
}

// GetBatchFromNum 批量查找 总家数
func (obj *_AnalyHyTblMgr) GetBatchFromNum(nums []int) (results []*AnalyHyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyTbl{}).Where("`num` IN (?)", nums).Find(&results).Error

	return
}

// GetFromUp 通过up获取内容 上涨家数
func (obj *_AnalyHyTblMgr) GetFromUp(up int) (results []*AnalyHyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyTbl{}).Where("`up` = ?", up).Find(&results).Error

	return
}

// GetBatchFromUp 批量查找 上涨家数
func (obj *_AnalyHyTblMgr) GetBatchFromUp(ups []int) (results []*AnalyHyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyTbl{}).Where("`up` IN (?)", ups).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_AnalyHyTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*AnalyHyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyTbl{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_AnalyHyTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*AnalyHyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyTbl{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromDay 通过day获取内容 入池时间
func (obj *_AnalyHyTblMgr) GetFromDay(day datatypes.Date) (results []*AnalyHyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyTbl{}).Where("`day` = ?", day).Find(&results).Error

	return
}

// GetBatchFromDay 批量查找 入池时间
func (obj *_AnalyHyTblMgr) GetBatchFromDay(days []datatypes.Date) (results []*AnalyHyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyTbl{}).Where("`day` IN (?)", days).Find(&results).Error

	return
}

// GetFromDayStr 通过day_str获取内容 入池时间
func (obj *_AnalyHyTblMgr) GetFromDayStr(dayStr string) (results []*AnalyHyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyTbl{}).Where("`day_str` = ?", dayStr).Find(&results).Error

	return
}

// GetBatchFromDayStr 批量查找 入池时间
func (obj *_AnalyHyTblMgr) GetBatchFromDayStr(dayStrs []string) (results []*AnalyHyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyTbl{}).Where("`day_str` IN (?)", dayStrs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AnalyHyTblMgr) FetchByPrimaryKey(id int) (result AnalyHyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueIndexByCode primary or index 获取唯一内容
func (obj *_AnalyHyTblMgr) FetchUniqueIndexByCode(code string, day datatypes.Date) (result AnalyHyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyTbl{}).Where("`code` = ? AND `day` = ?", code, day).First(&result).Error

	return
}
