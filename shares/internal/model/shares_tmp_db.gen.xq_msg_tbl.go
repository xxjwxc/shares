package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _XqMsgTblMgr struct {
	*_BaseMgr
}

// XqMsgTblMgr open func
func XqMsgTblMgr(db *gorm.DB) *_XqMsgTblMgr {
	if db == nil {
		panic(fmt.Errorf("XqMsgTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_XqMsgTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("xq_msg_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_XqMsgTblMgr) GetTableName() string {
	return "xq_msg_tbl"
}

// Reset 重置gorm会话
func (obj *_XqMsgTblMgr) Reset() *_XqMsgTblMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_XqMsgTblMgr) Get() (result XqMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgTbl{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_XqMsgTblMgr) Gets() (results []*XqMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgTbl{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_XqMsgTblMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(XqMsgTbl{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_XqMsgTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取 股票代码
func (obj *_XqMsgTblMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithTm tm获取 评论时间
func (obj *_XqMsgTblMgr) WithTm(tm int64) Option {
	return optionFunc(func(o *options) { o.query["tm"] = tm })
}

// WithName name获取 股票名字
func (obj *_XqMsgTblMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDay day获取 入池时间
func (obj *_XqMsgTblMgr) WithDay(day time.Time) Option {
	return optionFunc(func(o *options) { o.query["day"] = day })
}

// WithDayStr day_str获取 入池时间
func (obj *_XqMsgTblMgr) WithDayStr(dayStr string) Option {
	return optionFunc(func(o *options) { o.query["day_str"] = dayStr })
}

// WithMsg msg获取 消息内容
func (obj *_XqMsgTblMgr) WithMsg(msg string) Option {
	return optionFunc(func(o *options) { o.query["msg"] = msg })
}

// WithLexer lexer获取 分词结果
func (obj *_XqMsgTblMgr) WithLexer(lexer string) Option {
	return optionFunc(func(o *options) { o.query["lexer"] = lexer })
}

// WithTag tag获取 分词关键字
func (obj *_XqMsgTblMgr) WithTag(tag string) Option {
	return optionFunc(func(o *options) { o.query["tag"] = tag })
}

// WithSentiment sentiment获取 情绪(0:负向，1:中性，2:正向)
func (obj *_XqMsgTblMgr) WithSentiment(sentiment int) Option {
	return optionFunc(func(o *options) { o.query["sentiment"] = sentiment })
}

// WithPrice price获取 价格
func (obj *_XqMsgTblMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithPercent percent获取 百分比
func (obj *_XqMsgTblMgr) WithPercent(percent float64) Option {
	return optionFunc(func(o *options) { o.query["percent"] = percent })
}

// WithTurnoverRate turnover_rate获取 换手率
func (obj *_XqMsgTblMgr) WithTurnoverRate(turnoverRate float64) Option {
	return optionFunc(func(o *options) { o.query["turnover_rate"] = turnoverRate })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_XqMsgTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// GetByOption 功能选项模式获取
func (obj *_XqMsgTblMgr) GetByOption(opts ...Option) (result XqMsgTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(XqMsgTbl{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_XqMsgTblMgr) GetByOptions(opts ...Option) (results []*XqMsgTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(XqMsgTbl{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_XqMsgTblMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]XqMsgTbl, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(XqMsgTbl{}).Where(options.query)
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
func (obj *_XqMsgTblMgr) GetFromID(id int) (result XqMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_XqMsgTblMgr) GetBatchFromID(ids []int) (results []*XqMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgTbl{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 股票代码
func (obj *_XqMsgTblMgr) GetFromCode(code string) (results []*XqMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgTbl{}).Where("`code` = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量查找 股票代码
func (obj *_XqMsgTblMgr) GetBatchFromCode(codes []string) (results []*XqMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgTbl{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromTm 通过tm获取内容 评论时间
func (obj *_XqMsgTblMgr) GetFromTm(tm int64) (results []*XqMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgTbl{}).Where("`tm` = ?", tm).Find(&results).Error

	return
}

// GetBatchFromTm 批量查找 评论时间
func (obj *_XqMsgTblMgr) GetBatchFromTm(tms []int64) (results []*XqMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgTbl{}).Where("`tm` IN (?)", tms).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 股票名字
func (obj *_XqMsgTblMgr) GetFromName(name string) (results []*XqMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgTbl{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 股票名字
func (obj *_XqMsgTblMgr) GetBatchFromName(names []string) (results []*XqMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgTbl{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromDay 通过day获取内容 入池时间
func (obj *_XqMsgTblMgr) GetFromDay(day time.Time) (results []*XqMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgTbl{}).Where("`day` = ?", day).Find(&results).Error

	return
}

// GetBatchFromDay 批量查找 入池时间
func (obj *_XqMsgTblMgr) GetBatchFromDay(days []time.Time) (results []*XqMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgTbl{}).Where("`day` IN (?)", days).Find(&results).Error

	return
}

// GetFromDayStr 通过day_str获取内容 入池时间
func (obj *_XqMsgTblMgr) GetFromDayStr(dayStr string) (results []*XqMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgTbl{}).Where("`day_str` = ?", dayStr).Find(&results).Error

	return
}

// GetBatchFromDayStr 批量查找 入池时间
func (obj *_XqMsgTblMgr) GetBatchFromDayStr(dayStrs []string) (results []*XqMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgTbl{}).Where("`day_str` IN (?)", dayStrs).Find(&results).Error

	return
}

// GetFromMsg 通过msg获取内容 消息内容
func (obj *_XqMsgTblMgr) GetFromMsg(msg string) (results []*XqMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgTbl{}).Where("`msg` = ?", msg).Find(&results).Error

	return
}

// GetBatchFromMsg 批量查找 消息内容
func (obj *_XqMsgTblMgr) GetBatchFromMsg(msgs []string) (results []*XqMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgTbl{}).Where("`msg` IN (?)", msgs).Find(&results).Error

	return
}

// GetFromLexer 通过lexer获取内容 分词结果
func (obj *_XqMsgTblMgr) GetFromLexer(lexer string) (results []*XqMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgTbl{}).Where("`lexer` = ?", lexer).Find(&results).Error

	return
}

// GetBatchFromLexer 批量查找 分词结果
func (obj *_XqMsgTblMgr) GetBatchFromLexer(lexers []string) (results []*XqMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgTbl{}).Where("`lexer` IN (?)", lexers).Find(&results).Error

	return
}

// GetFromTag 通过tag获取内容 分词关键字
func (obj *_XqMsgTblMgr) GetFromTag(tag string) (results []*XqMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgTbl{}).Where("`tag` = ?", tag).Find(&results).Error

	return
}

// GetBatchFromTag 批量查找 分词关键字
func (obj *_XqMsgTblMgr) GetBatchFromTag(tags []string) (results []*XqMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgTbl{}).Where("`tag` IN (?)", tags).Find(&results).Error

	return
}

// GetFromSentiment 通过sentiment获取内容 情绪(0:负向，1:中性，2:正向)
func (obj *_XqMsgTblMgr) GetFromSentiment(sentiment int) (results []*XqMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgTbl{}).Where("`sentiment` = ?", sentiment).Find(&results).Error

	return
}

// GetBatchFromSentiment 批量查找 情绪(0:负向，1:中性，2:正向)
func (obj *_XqMsgTblMgr) GetBatchFromSentiment(sentiments []int) (results []*XqMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgTbl{}).Where("`sentiment` IN (?)", sentiments).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 价格
func (obj *_XqMsgTblMgr) GetFromPrice(price float64) (results []*XqMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgTbl{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 价格
func (obj *_XqMsgTblMgr) GetBatchFromPrice(prices []float64) (results []*XqMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgTbl{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromPercent 通过percent获取内容 百分比
func (obj *_XqMsgTblMgr) GetFromPercent(percent float64) (results []*XqMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgTbl{}).Where("`percent` = ?", percent).Find(&results).Error

	return
}

// GetBatchFromPercent 批量查找 百分比
func (obj *_XqMsgTblMgr) GetBatchFromPercent(percents []float64) (results []*XqMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgTbl{}).Where("`percent` IN (?)", percents).Find(&results).Error

	return
}

// GetFromTurnoverRate 通过turnover_rate获取内容 换手率
func (obj *_XqMsgTblMgr) GetFromTurnoverRate(turnoverRate float64) (results []*XqMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgTbl{}).Where("`turnover_rate` = ?", turnoverRate).Find(&results).Error

	return
}

// GetBatchFromTurnoverRate 批量查找 换手率
func (obj *_XqMsgTblMgr) GetBatchFromTurnoverRate(turnoverRates []float64) (results []*XqMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgTbl{}).Where("`turnover_rate` IN (?)", turnoverRates).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_XqMsgTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*XqMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgTbl{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_XqMsgTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*XqMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgTbl{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_XqMsgTblMgr) FetchByPrimaryKey(id int) (result XqMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueIndexByCode primary or index 获取唯一内容
func (obj *_XqMsgTblMgr) FetchUniqueIndexByCode(code string, tm int64) (result XqMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(XqMsgTbl{}).Where("`code` = ? AND `tm` = ?", code, tm).First(&result).Error

	return
}
