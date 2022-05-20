package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _XqMsgDailyTblMgr struct {
	*_BaseMgr
}

// XqMsgDailyTblMgr open func
func XqMsgDailyTblMgr(db *gorm.DB) *_XqMsgDailyTblMgr {
	if db == nil {
		panic(fmt.Errorf("XqMsgDailyTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_XqMsgDailyTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("xq_msg_daily_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_XqMsgDailyTblMgr) GetTableName() string {
	return "xq_msg_daily_tbl"
}

// Reset 重置gorm会话
func (obj *_XqMsgDailyTblMgr) Reset() *_XqMsgDailyTblMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_XqMsgDailyTblMgr) Get() (result XqMsgDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyTbl{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_XqMsgDailyTblMgr) Gets() (results []*XqMsgDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyTbl{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_XqMsgDailyTblMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(XqMsgDailyTbl{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_XqMsgDailyTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取 股票代码
func (obj *_XqMsgDailyTblMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithName name获取 股票名字
func (obj *_XqMsgDailyTblMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDay0 day0获取 入池时间
func (obj *_XqMsgDailyTblMgr) WithDay0(day0 int64) Option {
	return optionFunc(func(o *options) { o.query["day0"] = day0 })
}

// WithDay0Str day0_str获取 入池时间
func (obj *_XqMsgDailyTblMgr) WithDay0Str(day0Str string) Option {
	return optionFunc(func(o *options) { o.query["day0_str"] = day0Str })
}

// WithTag tag获取 分词关键字
func (obj *_XqMsgDailyTblMgr) WithTag(tag string) Option {
	return optionFunc(func(o *options) { o.query["tag"] = tag })
}

// WithGoodSentiment good_sentiment获取 正面情绪数量
func (obj *_XqMsgDailyTblMgr) WithGoodSentiment(goodSentiment int) Option {
	return optionFunc(func(o *options) { o.query["good_sentiment"] = goodSentiment })
}

// WithBadSentiment bad_sentiment获取 负面情绪数量
func (obj *_XqMsgDailyTblMgr) WithBadSentiment(badSentiment int) Option {
	return optionFunc(func(o *options) { o.query["bad_sentiment"] = badSentiment })
}

// WithSentiment sentiment获取 中性情绪数量
func (obj *_XqMsgDailyTblMgr) WithSentiment(sentiment int) Option {
	return optionFunc(func(o *options) { o.query["sentiment"] = sentiment })
}

// WithPrice price获取 价格
func (obj *_XqMsgDailyTblMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithPercent percent获取 百分比
func (obj *_XqMsgDailyTblMgr) WithPercent(percent float64) Option {
	return optionFunc(func(o *options) { o.query["percent"] = percent })
}

// WithTurnoverRate turnover_rate获取 换手率
func (obj *_XqMsgDailyTblMgr) WithTurnoverRate(turnoverRate float64) Option {
	return optionFunc(func(o *options) { o.query["turnover_rate"] = turnoverRate })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_XqMsgDailyTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// GetByOption 功能选项模式获取
func (obj *_XqMsgDailyTblMgr) GetByOption(opts ...Option) (result XqMsgDailyTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyTbl{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_XqMsgDailyTblMgr) GetByOptions(opts ...Option) (results []*XqMsgDailyTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyTbl{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_XqMsgDailyTblMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]XqMsgDailyTbl, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(XqMsgDailyTbl{}).Where(options.query)
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
func (obj *_XqMsgDailyTblMgr) GetFromID(id int) (result XqMsgDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_XqMsgDailyTblMgr) GetBatchFromID(ids []int) (results []*XqMsgDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyTbl{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 股票代码
func (obj *_XqMsgDailyTblMgr) GetFromCode(code string) (results []*XqMsgDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyTbl{}).Where("`code` = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量查找 股票代码
func (obj *_XqMsgDailyTblMgr) GetBatchFromCode(codes []string) (results []*XqMsgDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyTbl{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 股票名字
func (obj *_XqMsgDailyTblMgr) GetFromName(name string) (results []*XqMsgDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyTbl{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 股票名字
func (obj *_XqMsgDailyTblMgr) GetBatchFromName(names []string) (results []*XqMsgDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyTbl{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromDay0 通过day0获取内容 入池时间
func (obj *_XqMsgDailyTblMgr) GetFromDay0(day0 int64) (results []*XqMsgDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyTbl{}).Where("`day0` = ?", day0).Find(&results).Error

	return
}

// GetBatchFromDay0 批量查找 入池时间
func (obj *_XqMsgDailyTblMgr) GetBatchFromDay0(day0s []int64) (results []*XqMsgDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyTbl{}).Where("`day0` IN (?)", day0s).Find(&results).Error

	return
}

// GetFromDay0Str 通过day0_str获取内容 入池时间
func (obj *_XqMsgDailyTblMgr) GetFromDay0Str(day0Str string) (results []*XqMsgDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyTbl{}).Where("`day0_str` = ?", day0Str).Find(&results).Error

	return
}

// GetBatchFromDay0Str 批量查找 入池时间
func (obj *_XqMsgDailyTblMgr) GetBatchFromDay0Str(day0Strs []string) (results []*XqMsgDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyTbl{}).Where("`day0_str` IN (?)", day0Strs).Find(&results).Error

	return
}

// GetFromTag 通过tag获取内容 分词关键字
func (obj *_XqMsgDailyTblMgr) GetFromTag(tag string) (results []*XqMsgDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyTbl{}).Where("`tag` = ?", tag).Find(&results).Error

	return
}

// GetBatchFromTag 批量查找 分词关键字
func (obj *_XqMsgDailyTblMgr) GetBatchFromTag(tags []string) (results []*XqMsgDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyTbl{}).Where("`tag` IN (?)", tags).Find(&results).Error

	return
}

// GetFromGoodSentiment 通过good_sentiment获取内容 正面情绪数量
func (obj *_XqMsgDailyTblMgr) GetFromGoodSentiment(goodSentiment int) (results []*XqMsgDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyTbl{}).Where("`good_sentiment` = ?", goodSentiment).Find(&results).Error

	return
}

// GetBatchFromGoodSentiment 批量查找 正面情绪数量
func (obj *_XqMsgDailyTblMgr) GetBatchFromGoodSentiment(goodSentiments []int) (results []*XqMsgDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyTbl{}).Where("`good_sentiment` IN (?)", goodSentiments).Find(&results).Error

	return
}

// GetFromBadSentiment 通过bad_sentiment获取内容 负面情绪数量
func (obj *_XqMsgDailyTblMgr) GetFromBadSentiment(badSentiment int) (results []*XqMsgDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyTbl{}).Where("`bad_sentiment` = ?", badSentiment).Find(&results).Error

	return
}

// GetBatchFromBadSentiment 批量查找 负面情绪数量
func (obj *_XqMsgDailyTblMgr) GetBatchFromBadSentiment(badSentiments []int) (results []*XqMsgDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyTbl{}).Where("`bad_sentiment` IN (?)", badSentiments).Find(&results).Error

	return
}

// GetFromSentiment 通过sentiment获取内容 中性情绪数量
func (obj *_XqMsgDailyTblMgr) GetFromSentiment(sentiment int) (results []*XqMsgDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyTbl{}).Where("`sentiment` = ?", sentiment).Find(&results).Error

	return
}

// GetBatchFromSentiment 批量查找 中性情绪数量
func (obj *_XqMsgDailyTblMgr) GetBatchFromSentiment(sentiments []int) (results []*XqMsgDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyTbl{}).Where("`sentiment` IN (?)", sentiments).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 价格
func (obj *_XqMsgDailyTblMgr) GetFromPrice(price float64) (results []*XqMsgDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyTbl{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 价格
func (obj *_XqMsgDailyTblMgr) GetBatchFromPrice(prices []float64) (results []*XqMsgDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyTbl{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromPercent 通过percent获取内容 百分比
func (obj *_XqMsgDailyTblMgr) GetFromPercent(percent float64) (results []*XqMsgDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyTbl{}).Where("`percent` = ?", percent).Find(&results).Error

	return
}

// GetBatchFromPercent 批量查找 百分比
func (obj *_XqMsgDailyTblMgr) GetBatchFromPercent(percents []float64) (results []*XqMsgDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyTbl{}).Where("`percent` IN (?)", percents).Find(&results).Error

	return
}

// GetFromTurnoverRate 通过turnover_rate获取内容 换手率
func (obj *_XqMsgDailyTblMgr) GetFromTurnoverRate(turnoverRate float64) (results []*XqMsgDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyTbl{}).Where("`turnover_rate` = ?", turnoverRate).Find(&results).Error

	return
}

// GetBatchFromTurnoverRate 批量查找 换手率
func (obj *_XqMsgDailyTblMgr) GetBatchFromTurnoverRate(turnoverRates []float64) (results []*XqMsgDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyTbl{}).Where("`turnover_rate` IN (?)", turnoverRates).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_XqMsgDailyTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*XqMsgDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyTbl{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_XqMsgDailyTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*XqMsgDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyTbl{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_XqMsgDailyTblMgr) FetchByPrimaryKey(id int) (result XqMsgDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueIndexByCode primary or index 获取唯一内容
func (obj *_XqMsgDailyTblMgr) FetchUniqueIndexByCode(code string, day0 int64) (result XqMsgDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgDailyTbl{}).Where("`code` = ? AND `day0` = ?", code, day0).First(&result).Error

	return
}
