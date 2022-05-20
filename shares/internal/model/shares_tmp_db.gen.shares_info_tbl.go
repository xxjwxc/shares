package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SharesInfoTblMgr struct {
	*_BaseMgr
}

// SharesInfoTblMgr open func
func SharesInfoTblMgr(db *gorm.DB) *_SharesInfoTblMgr {
	if db == nil {
		panic(fmt.Errorf("SharesInfoTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SharesInfoTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("shares_info_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SharesInfoTblMgr) GetTableName() string {
	return "shares_info_tbl"
}

// Reset 重置gorm会话
func (obj *_SharesInfoTblMgr) Reset() *_SharesInfoTblMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SharesInfoTblMgr) Get() (result SharesInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SharesInfoTblMgr) Gets() (results []*SharesInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SharesInfoTblMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_SharesInfoTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取 股票代码
func (obj *_SharesInfoTblMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithSimpleCode simple_code获取 股票代码简写
func (obj *_SharesInfoTblMgr) WithSimpleCode(simpleCode string) Option {
	return optionFunc(func(o *options) { o.query["simple_code"] = simpleCode })
}

// WithExt ext获取 后缀
func (obj *_SharesInfoTblMgr) WithExt(ext string) Option {
	return optionFunc(func(o *options) { o.query["ext"] = ext })
}

// WithName name获取 股票名字
func (obj *_SharesInfoTblMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithSimpleName simple_name获取 股票名字拼音简写
func (obj *_SharesInfoTblMgr) WithSimpleName(simpleName string) Option {
	return optionFunc(func(o *options) { o.query["simple_name"] = simpleName })
}

// WithPrice price获取 当前价格
func (obj *_SharesInfoTblMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithPercent percent获取 百分比
func (obj *_SharesInfoTblMgr) WithPercent(percent float64) Option {
	return optionFunc(func(o *options) { o.query["percent"] = percent })
}

// WithTotal total获取 流通市值
func (obj *_SharesInfoTblMgr) WithTotal(total float64) Option {
	return optionFunc(func(o *options) { o.query["total"] = total })
}

// WithCap cap获取 总市值
func (obj *_SharesInfoTblMgr) WithCap(cap float64) Option {
	return optionFunc(func(o *options) { o.query["cap"] = cap })
}

// WithPeg peg获取 peg估值
func (obj *_SharesInfoTblMgr) WithPeg(peg float64) Option {
	return optionFunc(func(o *options) { o.query["peg"] = peg })
}

// WithPegAvg peg_avg获取 peg行业平均
func (obj *_SharesInfoTblMgr) WithPegAvg(pegAvg float64) Option {
	return optionFunc(func(o *options) { o.query["peg_avg"] = pegAvg })
}

// WithPegAve peg_ave获取 peg行业中值
func (obj *_SharesInfoTblMgr) WithPegAve(pegAve float64) Option {
	return optionFunc(func(o *options) { o.query["peg_ave"] = pegAve })
}

// WithPegShares peg_shares获取 估值比较
func (obj *_SharesInfoTblMgr) WithPegShares(pegShares string) Option {
	return optionFunc(func(o *options) { o.query["peg_shares"] = pegShares })
}

// WithHyName hy_name获取 行业名
func (obj *_SharesInfoTblMgr) WithHyName(hyName string) Option {
	return optionFunc(func(o *options) { o.query["hy_name"] = hyName })
}

// WithCreatedBy created_by获取 创建者
func (obj *_SharesInfoTblMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_SharesInfoTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// GetByOption 功能选项模式获取
func (obj *_SharesInfoTblMgr) GetByOption(opts ...Option) (result SharesInfoTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SharesInfoTblMgr) GetByOptions(opts ...Option) (results []*SharesInfoTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_SharesInfoTblMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]SharesInfoTbl, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).Where(options.query)
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
func (obj *_SharesInfoTblMgr) GetFromID(id int) (result SharesInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_SharesInfoTblMgr) GetBatchFromID(ids []int) (results []*SharesInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 股票代码
func (obj *_SharesInfoTblMgr) GetFromCode(code string) (result SharesInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).Where("`code` = ?", code).First(&result).Error

	return
}

// GetBatchFromCode 批量查找 股票代码
func (obj *_SharesInfoTblMgr) GetBatchFromCode(codes []string) (results []*SharesInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromSimpleCode 通过simple_code获取内容 股票代码简写
func (obj *_SharesInfoTblMgr) GetFromSimpleCode(simpleCode string) (results []*SharesInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).Where("`simple_code` = ?", simpleCode).Find(&results).Error

	return
}

// GetBatchFromSimpleCode 批量查找 股票代码简写
func (obj *_SharesInfoTblMgr) GetBatchFromSimpleCode(simpleCodes []string) (results []*SharesInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).Where("`simple_code` IN (?)", simpleCodes).Find(&results).Error

	return
}

// GetFromExt 通过ext获取内容 后缀
func (obj *_SharesInfoTblMgr) GetFromExt(ext string) (results []*SharesInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).Where("`ext` = ?", ext).Find(&results).Error

	return
}

// GetBatchFromExt 批量查找 后缀
func (obj *_SharesInfoTblMgr) GetBatchFromExt(exts []string) (results []*SharesInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).Where("`ext` IN (?)", exts).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 股票名字
func (obj *_SharesInfoTblMgr) GetFromName(name string) (results []*SharesInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 股票名字
func (obj *_SharesInfoTblMgr) GetBatchFromName(names []string) (results []*SharesInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromSimpleName 通过simple_name获取内容 股票名字拼音简写
func (obj *_SharesInfoTblMgr) GetFromSimpleName(simpleName string) (results []*SharesInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).Where("`simple_name` = ?", simpleName).Find(&results).Error

	return
}

// GetBatchFromSimpleName 批量查找 股票名字拼音简写
func (obj *_SharesInfoTblMgr) GetBatchFromSimpleName(simpleNames []string) (results []*SharesInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).Where("`simple_name` IN (?)", simpleNames).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 当前价格
func (obj *_SharesInfoTblMgr) GetFromPrice(price float64) (results []*SharesInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 当前价格
func (obj *_SharesInfoTblMgr) GetBatchFromPrice(prices []float64) (results []*SharesInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromPercent 通过percent获取内容 百分比
func (obj *_SharesInfoTblMgr) GetFromPercent(percent float64) (results []*SharesInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).Where("`percent` = ?", percent).Find(&results).Error

	return
}

// GetBatchFromPercent 批量查找 百分比
func (obj *_SharesInfoTblMgr) GetBatchFromPercent(percents []float64) (results []*SharesInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).Where("`percent` IN (?)", percents).Find(&results).Error

	return
}

// GetFromTotal 通过total获取内容 流通市值
func (obj *_SharesInfoTblMgr) GetFromTotal(total float64) (results []*SharesInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).Where("`total` = ?", total).Find(&results).Error

	return
}

// GetBatchFromTotal 批量查找 流通市值
func (obj *_SharesInfoTblMgr) GetBatchFromTotal(totals []float64) (results []*SharesInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).Where("`total` IN (?)", totals).Find(&results).Error

	return
}

// GetFromCap 通过cap获取内容 总市值
func (obj *_SharesInfoTblMgr) GetFromCap(cap float64) (results []*SharesInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).Where("`cap` = ?", cap).Find(&results).Error

	return
}

// GetBatchFromCap 批量查找 总市值
func (obj *_SharesInfoTblMgr) GetBatchFromCap(caps []float64) (results []*SharesInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).Where("`cap` IN (?)", caps).Find(&results).Error

	return
}

// GetFromPeg 通过peg获取内容 peg估值
func (obj *_SharesInfoTblMgr) GetFromPeg(peg float64) (results []*SharesInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).Where("`peg` = ?", peg).Find(&results).Error

	return
}

// GetBatchFromPeg 批量查找 peg估值
func (obj *_SharesInfoTblMgr) GetBatchFromPeg(pegs []float64) (results []*SharesInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).Where("`peg` IN (?)", pegs).Find(&results).Error

	return
}

// GetFromPegAvg 通过peg_avg获取内容 peg行业平均
func (obj *_SharesInfoTblMgr) GetFromPegAvg(pegAvg float64) (results []*SharesInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).Where("`peg_avg` = ?", pegAvg).Find(&results).Error

	return
}

// GetBatchFromPegAvg 批量查找 peg行业平均
func (obj *_SharesInfoTblMgr) GetBatchFromPegAvg(pegAvgs []float64) (results []*SharesInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).Where("`peg_avg` IN (?)", pegAvgs).Find(&results).Error

	return
}

// GetFromPegAve 通过peg_ave获取内容 peg行业中值
func (obj *_SharesInfoTblMgr) GetFromPegAve(pegAve float64) (results []*SharesInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).Where("`peg_ave` = ?", pegAve).Find(&results).Error

	return
}

// GetBatchFromPegAve 批量查找 peg行业中值
func (obj *_SharesInfoTblMgr) GetBatchFromPegAve(pegAves []float64) (results []*SharesInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).Where("`peg_ave` IN (?)", pegAves).Find(&results).Error

	return
}

// GetFromPegShares 通过peg_shares获取内容 估值比较
func (obj *_SharesInfoTblMgr) GetFromPegShares(pegShares string) (results []*SharesInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).Where("`peg_shares` = ?", pegShares).Find(&results).Error

	return
}

// GetBatchFromPegShares 批量查找 估值比较
func (obj *_SharesInfoTblMgr) GetBatchFromPegShares(pegSharess []string) (results []*SharesInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).Where("`peg_shares` IN (?)", pegSharess).Find(&results).Error

	return
}

// GetFromHyName 通过hy_name获取内容 行业名
func (obj *_SharesInfoTblMgr) GetFromHyName(hyName string) (results []*SharesInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).Where("`hy_name` = ?", hyName).Find(&results).Error

	return
}

// GetBatchFromHyName 批量查找 行业名
func (obj *_SharesInfoTblMgr) GetBatchFromHyName(hyNames []string) (results []*SharesInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).Where("`hy_name` IN (?)", hyNames).Find(&results).Error

	return
}

// GetFromCreatedBy 通过created_by获取内容 创建者
func (obj *_SharesInfoTblMgr) GetFromCreatedBy(createdBy string) (results []*SharesInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).Where("`created_by` = ?", createdBy).Find(&results).Error

	return
}

// GetBatchFromCreatedBy 批量查找 创建者
func (obj *_SharesInfoTblMgr) GetBatchFromCreatedBy(createdBys []string) (results []*SharesInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).Where("`created_by` IN (?)", createdBys).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_SharesInfoTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*SharesInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_SharesInfoTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*SharesInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SharesInfoTblMgr) FetchByPrimaryKey(id int) (result SharesInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByCode primary or index 获取唯一内容
func (obj *_SharesInfoTblMgr) FetchUniqueByCode(code string) (result SharesInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesInfoTbl{}).Where("`code` = ?", code).First(&result).Error

	return
}
