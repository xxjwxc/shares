package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _Ma20ViewMgr struct {
	*_BaseMgr
}

// Ma20ViewMgr open func
func Ma20ViewMgr(db *gorm.DB) *_Ma20ViewMgr {
	if db == nil {
		panic(fmt.Errorf("Ma20ViewMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_Ma20ViewMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ma20_view"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_Ma20ViewMgr) GetTableName() string {
	return "ma20_view"
}

// Reset 重置gorm会话
func (obj *_Ma20ViewMgr) Reset() *_Ma20ViewMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_Ma20ViewMgr) Get() (result Ma20View, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Ma20View{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_Ma20ViewMgr) Gets() (results []*Ma20View, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Ma20View{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_Ma20ViewMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Ma20View{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithCode code获取 股票代码
func (obj *_Ma20ViewMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithName name获取 股票名字
func (obj *_Ma20ViewMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithPrice price获取 当前价格
func (obj *_Ma20ViewMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithMa20 ma20获取 20日均线
func (obj *_Ma20ViewMgr) WithMa20(ma20 float64) Option {
	return optionFunc(func(o *options) { o.query["ma20"] = ma20 })
}

// WithPeg peg获取 peg估值
func (obj *_Ma20ViewMgr) WithPeg(peg float64) Option {
	return optionFunc(func(o *options) { o.query["peg"] = peg })
}

// WithPegShares peg_shares获取 估值比较
func (obj *_Ma20ViewMgr) WithPegShares(pegShares string) Option {
	return optionFunc(func(o *options) { o.query["peg_shares"] = pegShares })
}

// WithHyName hy_name获取 行业名
func (obj *_Ma20ViewMgr) WithHyName(hyName string) Option {
	return optionFunc(func(o *options) { o.query["hy_name"] = hyName })
}

// WithNickName nick_name获取 用户昵称
func (obj *_Ma20ViewMgr) WithNickName(nickName string) Option {
	return optionFunc(func(o *options) { o.query["nick_name"] = nickName })
}

// WithDay0Str day0_str获取 当日0点时间戳
func (obj *_Ma20ViewMgr) WithDay0Str(day0Str string) Option {
	return optionFunc(func(o *options) { o.query["day0_str"] = day0Str })
}

// GetByOption 功能选项模式获取
func (obj *_Ma20ViewMgr) GetByOption(opts ...Option) (result Ma20View, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Ma20View{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_Ma20ViewMgr) GetByOptions(opts ...Option) (results []*Ma20View, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Ma20View{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_Ma20ViewMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Ma20View, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Ma20View{}).Where(options.query)
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

// GetFromCode 通过code获取内容 股票代码
func (obj *_Ma20ViewMgr) GetFromCode(code string) (results []*Ma20View, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Ma20View{}).Where("`code` = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量查找 股票代码
func (obj *_Ma20ViewMgr) GetBatchFromCode(codes []string) (results []*Ma20View, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Ma20View{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 股票名字
func (obj *_Ma20ViewMgr) GetFromName(name string) (results []*Ma20View, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Ma20View{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 股票名字
func (obj *_Ma20ViewMgr) GetBatchFromName(names []string) (results []*Ma20View, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Ma20View{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 当前价格
func (obj *_Ma20ViewMgr) GetFromPrice(price float64) (results []*Ma20View, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Ma20View{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 当前价格
func (obj *_Ma20ViewMgr) GetBatchFromPrice(prices []float64) (results []*Ma20View, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Ma20View{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromMa20 通过ma20获取内容 20日均线
func (obj *_Ma20ViewMgr) GetFromMa20(ma20 float64) (results []*Ma20View, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Ma20View{}).Where("`ma20` = ?", ma20).Find(&results).Error

	return
}

// GetBatchFromMa20 批量查找 20日均线
func (obj *_Ma20ViewMgr) GetBatchFromMa20(ma20s []float64) (results []*Ma20View, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Ma20View{}).Where("`ma20` IN (?)", ma20s).Find(&results).Error

	return
}

// GetFromPeg 通过peg获取内容 peg估值
func (obj *_Ma20ViewMgr) GetFromPeg(peg float64) (results []*Ma20View, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Ma20View{}).Where("`peg` = ?", peg).Find(&results).Error

	return
}

// GetBatchFromPeg 批量查找 peg估值
func (obj *_Ma20ViewMgr) GetBatchFromPeg(pegs []float64) (results []*Ma20View, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Ma20View{}).Where("`peg` IN (?)", pegs).Find(&results).Error

	return
}

// GetFromPegShares 通过peg_shares获取内容 估值比较
func (obj *_Ma20ViewMgr) GetFromPegShares(pegShares string) (results []*Ma20View, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Ma20View{}).Where("`peg_shares` = ?", pegShares).Find(&results).Error

	return
}

// GetBatchFromPegShares 批量查找 估值比较
func (obj *_Ma20ViewMgr) GetBatchFromPegShares(pegSharess []string) (results []*Ma20View, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Ma20View{}).Where("`peg_shares` IN (?)", pegSharess).Find(&results).Error

	return
}

// GetFromHyName 通过hy_name获取内容 行业名
func (obj *_Ma20ViewMgr) GetFromHyName(hyName string) (results []*Ma20View, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Ma20View{}).Where("`hy_name` = ?", hyName).Find(&results).Error

	return
}

// GetBatchFromHyName 批量查找 行业名
func (obj *_Ma20ViewMgr) GetBatchFromHyName(hyNames []string) (results []*Ma20View, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Ma20View{}).Where("`hy_name` IN (?)", hyNames).Find(&results).Error

	return
}

// GetFromNickName 通过nick_name获取内容 用户昵称
func (obj *_Ma20ViewMgr) GetFromNickName(nickName string) (results []*Ma20View, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Ma20View{}).Where("`nick_name` = ?", nickName).Find(&results).Error

	return
}

// GetBatchFromNickName 批量查找 用户昵称
func (obj *_Ma20ViewMgr) GetBatchFromNickName(nickNames []string) (results []*Ma20View, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Ma20View{}).Where("`nick_name` IN (?)", nickNames).Find(&results).Error

	return
}

// GetFromDay0Str 通过day0_str获取内容 当日0点时间戳
func (obj *_Ma20ViewMgr) GetFromDay0Str(day0Str string) (results []*Ma20View, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Ma20View{}).Where("`day0_str` = ?", day0Str).Find(&results).Error

	return
}

// GetBatchFromDay0Str 批量查找 当日0点时间戳
func (obj *_Ma20ViewMgr) GetBatchFromDay0Str(day0Strs []string) (results []*Ma20View, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Ma20View{}).Where("`day0_str` IN (?)", day0Strs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
