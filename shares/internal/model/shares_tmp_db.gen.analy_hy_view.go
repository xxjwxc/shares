package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AnalyHyViewMgr struct {
	*_BaseMgr
}

// AnalyHyViewMgr open func
func AnalyHyViewMgr(db *gorm.DB) *_AnalyHyViewMgr {
	if db == nil {
		panic(fmt.Errorf("AnalyHyViewMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AnalyHyViewMgr{_BaseMgr: &_BaseMgr{DB: db.Table("analy_hy_view"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AnalyHyViewMgr) GetTableName() string {
	return "analy_hy_view"
}

// Reset 重置gorm会话
func (obj *_AnalyHyViewMgr) Reset() *_AnalyHyViewMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AnalyHyViewMgr) Get() (result AnalyHyView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyView{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AnalyHyViewMgr) Gets() (results []*AnalyHyView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyView{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AnalyHyViewMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AnalyHyView{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithHyCode hy_code获取 股票代码
func (obj *_AnalyHyViewMgr) WithHyCode(hyCode string) Option {
	return optionFunc(func(o *options) { o.query["hy_code"] = hyCode })
}

// WithHyName hy_name获取 行业名
func (obj *_AnalyHyViewMgr) WithHyName(hyName string) Option {
	return optionFunc(func(o *options) { o.query["hy_name"] = hyName })
}

// WithPrice1 price1获取
func (obj *_AnalyHyViewMgr) WithPrice1(price1 float64) Option {
	return optionFunc(func(o *options) { o.query["price1"] = price1 })
}

// WithPrice5 price5获取
func (obj *_AnalyHyViewMgr) WithPrice5(price5 float64) Option {
	return optionFunc(func(o *options) { o.query["price5"] = price5 })
}

// WithPrice10 price10获取
func (obj *_AnalyHyViewMgr) WithPrice10(price10 float64) Option {
	return optionFunc(func(o *options) { o.query["price10"] = price10 })
}

// WithPrice20 price20获取
func (obj *_AnalyHyViewMgr) WithPrice20(price20 float64) Option {
	return optionFunc(func(o *options) { o.query["price20"] = price20 })
}

// WithDay0 day0获取 当日0点时间戳
func (obj *_AnalyHyViewMgr) WithDay0(day0 int64) Option {
	return optionFunc(func(o *options) { o.query["day0"] = day0 })
}

// WithDay0Str day0_str获取 当日0点时间戳
func (obj *_AnalyHyViewMgr) WithDay0Str(day0Str string) Option {
	return optionFunc(func(o *options) { o.query["day0_str"] = day0Str })
}

// GetByOption 功能选项模式获取
func (obj *_AnalyHyViewMgr) GetByOption(opts ...Option) (result AnalyHyView, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyView{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AnalyHyViewMgr) GetByOptions(opts ...Option) (results []*AnalyHyView, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyView{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AnalyHyViewMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AnalyHyView, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AnalyHyView{}).Where(options.query)
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

// GetFromHyCode 通过hy_code获取内容 股票代码
func (obj *_AnalyHyViewMgr) GetFromHyCode(hyCode string) (results []*AnalyHyView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyView{}).Where("`hy_code` = ?", hyCode).Find(&results).Error

	return
}

// GetBatchFromHyCode 批量查找 股票代码
func (obj *_AnalyHyViewMgr) GetBatchFromHyCode(hyCodes []string) (results []*AnalyHyView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyView{}).Where("`hy_code` IN (?)", hyCodes).Find(&results).Error

	return
}

// GetFromHyName 通过hy_name获取内容 行业名
func (obj *_AnalyHyViewMgr) GetFromHyName(hyName string) (results []*AnalyHyView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyView{}).Where("`hy_name` = ?", hyName).Find(&results).Error

	return
}

// GetBatchFromHyName 批量查找 行业名
func (obj *_AnalyHyViewMgr) GetBatchFromHyName(hyNames []string) (results []*AnalyHyView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyView{}).Where("`hy_name` IN (?)", hyNames).Find(&results).Error

	return
}

// GetFromPrice1 通过price1获取内容
func (obj *_AnalyHyViewMgr) GetFromPrice1(price1 float64) (results []*AnalyHyView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyView{}).Where("`price1` = ?", price1).Find(&results).Error

	return
}

// GetBatchFromPrice1 批量查找
func (obj *_AnalyHyViewMgr) GetBatchFromPrice1(price1s []float64) (results []*AnalyHyView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyView{}).Where("`price1` IN (?)", price1s).Find(&results).Error

	return
}

// GetFromPrice5 通过price5获取内容
func (obj *_AnalyHyViewMgr) GetFromPrice5(price5 float64) (results []*AnalyHyView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyView{}).Where("`price5` = ?", price5).Find(&results).Error

	return
}

// GetBatchFromPrice5 批量查找
func (obj *_AnalyHyViewMgr) GetBatchFromPrice5(price5s []float64) (results []*AnalyHyView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyView{}).Where("`price5` IN (?)", price5s).Find(&results).Error

	return
}

// GetFromPrice10 通过price10获取内容
func (obj *_AnalyHyViewMgr) GetFromPrice10(price10 float64) (results []*AnalyHyView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyView{}).Where("`price10` = ?", price10).Find(&results).Error

	return
}

// GetBatchFromPrice10 批量查找
func (obj *_AnalyHyViewMgr) GetBatchFromPrice10(price10s []float64) (results []*AnalyHyView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyView{}).Where("`price10` IN (?)", price10s).Find(&results).Error

	return
}

// GetFromPrice20 通过price20获取内容
func (obj *_AnalyHyViewMgr) GetFromPrice20(price20 float64) (results []*AnalyHyView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyView{}).Where("`price20` = ?", price20).Find(&results).Error

	return
}

// GetBatchFromPrice20 批量查找
func (obj *_AnalyHyViewMgr) GetBatchFromPrice20(price20s []float64) (results []*AnalyHyView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyView{}).Where("`price20` IN (?)", price20s).Find(&results).Error

	return
}

// GetFromDay0 通过day0获取内容 当日0点时间戳
func (obj *_AnalyHyViewMgr) GetFromDay0(day0 int64) (results []*AnalyHyView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyView{}).Where("`day0` = ?", day0).Find(&results).Error

	return
}

// GetBatchFromDay0 批量查找 当日0点时间戳
func (obj *_AnalyHyViewMgr) GetBatchFromDay0(day0s []int64) (results []*AnalyHyView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyView{}).Where("`day0` IN (?)", day0s).Find(&results).Error

	return
}

// GetFromDay0Str 通过day0_str获取内容 当日0点时间戳
func (obj *_AnalyHyViewMgr) GetFromDay0Str(day0Str string) (results []*AnalyHyView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyView{}).Where("`day0_str` = ?", day0Str).Find(&results).Error

	return
}

// GetBatchFromDay0Str 批量查找 当日0点时间戳
func (obj *_AnalyHyViewMgr) GetBatchFromDay0Str(day0Strs []string) (results []*AnalyHyView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AnalyHyView{}).Where("`day0_str` IN (?)", day0Strs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
