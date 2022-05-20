package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SharesDailyTblMgr struct {
	*_BaseMgr
}

// SharesDailyTblMgr open func
func SharesDailyTblMgr(db *gorm.DB) *_SharesDailyTblMgr {
	if db == nil {
		panic(fmt.Errorf("SharesDailyTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SharesDailyTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("shares_daily_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SharesDailyTblMgr) GetTableName() string {
	return "shares_daily_tbl"
}

// Reset 重置gorm会话
func (obj *_SharesDailyTblMgr) Reset() *_SharesDailyTblMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SharesDailyTblMgr) Get() (result SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SharesDailyTblMgr) Gets() (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SharesDailyTblMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_SharesDailyTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取 股票代码
func (obj *_SharesDailyTblMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithPrice price获取 当前价格
func (obj *_SharesDailyTblMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithPercent percent获取 百分比
func (obj *_SharesDailyTblMgr) WithPercent(percent float64) Option {
	return optionFunc(func(o *options) { o.query["percent"] = percent })
}

// WithMa5 ma5获取 5日均线
func (obj *_SharesDailyTblMgr) WithMa5(ma5 float64) Option {
	return optionFunc(func(o *options) { o.query["ma5"] = ma5 })
}

// WithMa10 ma10获取 10日均线
func (obj *_SharesDailyTblMgr) WithMa10(ma10 float64) Option {
	return optionFunc(func(o *options) { o.query["ma10"] = ma10 })
}

// WithMa20 ma20获取 20日均线
func (obj *_SharesDailyTblMgr) WithMa20(ma20 float64) Option {
	return optionFunc(func(o *options) { o.query["ma20"] = ma20 })
}

// WithMa30 ma30获取 20日均线
func (obj *_SharesDailyTblMgr) WithMa30(ma30 float64) Option {
	return optionFunc(func(o *options) { o.query["ma30"] = ma30 })
}

// WithMa60 ma60获取 20日均线
func (obj *_SharesDailyTblMgr) WithMa60(ma60 float64) Option {
	return optionFunc(func(o *options) { o.query["ma60"] = ma60 })
}

// WithDay0 day0获取 当日0点时间戳
func (obj *_SharesDailyTblMgr) WithDay0(day0 int64) Option {
	return optionFunc(func(o *options) { o.query["day0"] = day0 })
}

// WithDay0Str day0_str获取 当日0点时间戳
func (obj *_SharesDailyTblMgr) WithDay0Str(day0Str string) Option {
	return optionFunc(func(o *options) { o.query["day0_str"] = day0Str })
}

// WithVolume volume获取 成交量（手）
func (obj *_SharesDailyTblMgr) WithVolume(volume float64) Option {
	return optionFunc(func(o *options) { o.query["volume"] = volume })
}

// WithTurnover turnover获取 成交额（万）
func (obj *_SharesDailyTblMgr) WithTurnover(turnover float64) Option {
	return optionFunc(func(o *options) { o.query["turnover"] = turnover })
}

// WithTurnoverRate turnover_rate获取 换手率
func (obj *_SharesDailyTblMgr) WithTurnoverRate(turnoverRate float64) Option {
	return optionFunc(func(o *options) { o.query["turnover_rate"] = turnoverRate })
}

// WithVol vol获取 成交量占比 (volume / volume 前5日平均)
func (obj *_SharesDailyTblMgr) WithVol(vol float64) Option {
	return optionFunc(func(o *options) { o.query["vol"] = vol })
}

// WithPe pe获取 市盈率
func (obj *_SharesDailyTblMgr) WithPe(pe float64) Option {
	return optionFunc(func(o *options) { o.query["pe"] = pe })
}

// WithPb pb获取 市净率
func (obj *_SharesDailyTblMgr) WithPb(pb float64) Option {
	return optionFunc(func(o *options) { o.query["pb"] = pb })
}

// WithTotal total获取 流通市值
func (obj *_SharesDailyTblMgr) WithTotal(total float64) Option {
	return optionFunc(func(o *options) { o.query["total"] = total })
}

// WithCap cap获取 总市值
func (obj *_SharesDailyTblMgr) WithCap(cap float64) Option {
	return optionFunc(func(o *options) { o.query["cap"] = cap })
}

// WithZljlr zljlr获取 主力净流入
func (obj *_SharesDailyTblMgr) WithZljlr(zljlr float64) Option {
	return optionFunc(func(o *options) { o.query["zljlr"] = zljlr })
}

// WithOpenPrice open_price获取 开盘价
func (obj *_SharesDailyTblMgr) WithOpenPrice(openPrice float64) Option {
	return optionFunc(func(o *options) { o.query["open_price"] = openPrice })
}

// WithClosePrice close_price获取 收盘价
func (obj *_SharesDailyTblMgr) WithClosePrice(closePrice float64) Option {
	return optionFunc(func(o *options) { o.query["close_price"] = closePrice })
}

// WithBscg bscg获取 北上持股数
func (obj *_SharesDailyTblMgr) WithBscg(bscg float64) Option {
	return optionFunc(func(o *options) { o.query["bscg"] = bscg })
}

// WithBsjlr bsjlr获取 北上净流入
func (obj *_SharesDailyTblMgr) WithBsjlr(bsjlr float64) Option {
	return optionFunc(func(o *options) { o.query["bsjlr"] = bsjlr })
}

// WithBsPercent bs_percent获取 北上持股占比
func (obj *_SharesDailyTblMgr) WithBsPercent(bsPercent float64) Option {
	return optionFunc(func(o *options) { o.query["bs_percent"] = bsPercent })
}

// WithMacd macd获取 macd 指标
func (obj *_SharesDailyTblMgr) WithMacd(macd float64) Option {
	return optionFunc(func(o *options) { o.query["macd"] = macd })
}

// WithDif dif获取 DIF快线
func (obj *_SharesDailyTblMgr) WithDif(dif float64) Option {
	return optionFunc(func(o *options) { o.query["dif"] = dif })
}

// WithDea dea获取 Dea 慢线
func (obj *_SharesDailyTblMgr) WithDea(dea float64) Option {
	return optionFunc(func(o *options) { o.query["dea"] = dea })
}

// WithMax max获取 当日最高值
func (obj *_SharesDailyTblMgr) WithMax(max float64) Option {
	return optionFunc(func(o *options) { o.query["max"] = max })
}

// WithMin min获取 当日最低价
func (obj *_SharesDailyTblMgr) WithMin(min float64) Option {
	return optionFunc(func(o *options) { o.query["min"] = min })
}

// WithPeg peg获取 peg估值
func (obj *_SharesDailyTblMgr) WithPeg(peg float64) Option {
	return optionFunc(func(o *options) { o.query["peg"] = peg })
}

// WithCreatedBy created_by获取 创建者
func (obj *_SharesDailyTblMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_SharesDailyTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// GetByOption 功能选项模式获取
func (obj *_SharesDailyTblMgr) GetByOption(opts ...Option) (result SharesDailyTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SharesDailyTblMgr) GetByOptions(opts ...Option) (results []*SharesDailyTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_SharesDailyTblMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]SharesDailyTbl, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where(options.query)
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
func (obj *_SharesDailyTblMgr) GetFromID(id int) (result SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_SharesDailyTblMgr) GetBatchFromID(ids []int) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 股票代码
func (obj *_SharesDailyTblMgr) GetFromCode(code string) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`code` = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量查找 股票代码
func (obj *_SharesDailyTblMgr) GetBatchFromCode(codes []string) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 当前价格
func (obj *_SharesDailyTblMgr) GetFromPrice(price float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 当前价格
func (obj *_SharesDailyTblMgr) GetBatchFromPrice(prices []float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromPercent 通过percent获取内容 百分比
func (obj *_SharesDailyTblMgr) GetFromPercent(percent float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`percent` = ?", percent).Find(&results).Error

	return
}

// GetBatchFromPercent 批量查找 百分比
func (obj *_SharesDailyTblMgr) GetBatchFromPercent(percents []float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`percent` IN (?)", percents).Find(&results).Error

	return
}

// GetFromMa5 通过ma5获取内容 5日均线
func (obj *_SharesDailyTblMgr) GetFromMa5(ma5 float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`ma5` = ?", ma5).Find(&results).Error

	return
}

// GetBatchFromMa5 批量查找 5日均线
func (obj *_SharesDailyTblMgr) GetBatchFromMa5(ma5s []float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`ma5` IN (?)", ma5s).Find(&results).Error

	return
}

// GetFromMa10 通过ma10获取内容 10日均线
func (obj *_SharesDailyTblMgr) GetFromMa10(ma10 float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`ma10` = ?", ma10).Find(&results).Error

	return
}

// GetBatchFromMa10 批量查找 10日均线
func (obj *_SharesDailyTblMgr) GetBatchFromMa10(ma10s []float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`ma10` IN (?)", ma10s).Find(&results).Error

	return
}

// GetFromMa20 通过ma20获取内容 20日均线
func (obj *_SharesDailyTblMgr) GetFromMa20(ma20 float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`ma20` = ?", ma20).Find(&results).Error

	return
}

// GetBatchFromMa20 批量查找 20日均线
func (obj *_SharesDailyTblMgr) GetBatchFromMa20(ma20s []float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`ma20` IN (?)", ma20s).Find(&results).Error

	return
}

// GetFromMa30 通过ma30获取内容 20日均线
func (obj *_SharesDailyTblMgr) GetFromMa30(ma30 float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`ma30` = ?", ma30).Find(&results).Error

	return
}

// GetBatchFromMa30 批量查找 20日均线
func (obj *_SharesDailyTblMgr) GetBatchFromMa30(ma30s []float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`ma30` IN (?)", ma30s).Find(&results).Error

	return
}

// GetFromMa60 通过ma60获取内容 20日均线
func (obj *_SharesDailyTblMgr) GetFromMa60(ma60 float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`ma60` = ?", ma60).Find(&results).Error

	return
}

// GetBatchFromMa60 批量查找 20日均线
func (obj *_SharesDailyTblMgr) GetBatchFromMa60(ma60s []float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`ma60` IN (?)", ma60s).Find(&results).Error

	return
}

// GetFromDay0 通过day0获取内容 当日0点时间戳
func (obj *_SharesDailyTblMgr) GetFromDay0(day0 int64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`day0` = ?", day0).Find(&results).Error

	return
}

// GetBatchFromDay0 批量查找 当日0点时间戳
func (obj *_SharesDailyTblMgr) GetBatchFromDay0(day0s []int64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`day0` IN (?)", day0s).Find(&results).Error

	return
}

// GetFromDay0Str 通过day0_str获取内容 当日0点时间戳
func (obj *_SharesDailyTblMgr) GetFromDay0Str(day0Str string) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`day0_str` = ?", day0Str).Find(&results).Error

	return
}

// GetBatchFromDay0Str 批量查找 当日0点时间戳
func (obj *_SharesDailyTblMgr) GetBatchFromDay0Str(day0Strs []string) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`day0_str` IN (?)", day0Strs).Find(&results).Error

	return
}

// GetFromVolume 通过volume获取内容 成交量（手）
func (obj *_SharesDailyTblMgr) GetFromVolume(volume float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`volume` = ?", volume).Find(&results).Error

	return
}

// GetBatchFromVolume 批量查找 成交量（手）
func (obj *_SharesDailyTblMgr) GetBatchFromVolume(volumes []float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`volume` IN (?)", volumes).Find(&results).Error

	return
}

// GetFromTurnover 通过turnover获取内容 成交额（万）
func (obj *_SharesDailyTblMgr) GetFromTurnover(turnover float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`turnover` = ?", turnover).Find(&results).Error

	return
}

// GetBatchFromTurnover 批量查找 成交额（万）
func (obj *_SharesDailyTblMgr) GetBatchFromTurnover(turnovers []float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`turnover` IN (?)", turnovers).Find(&results).Error

	return
}

// GetFromTurnoverRate 通过turnover_rate获取内容 换手率
func (obj *_SharesDailyTblMgr) GetFromTurnoverRate(turnoverRate float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`turnover_rate` = ?", turnoverRate).Find(&results).Error

	return
}

// GetBatchFromTurnoverRate 批量查找 换手率
func (obj *_SharesDailyTblMgr) GetBatchFromTurnoverRate(turnoverRates []float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`turnover_rate` IN (?)", turnoverRates).Find(&results).Error

	return
}

// GetFromVol 通过vol获取内容 成交量占比 (volume / volume 前5日平均)
func (obj *_SharesDailyTblMgr) GetFromVol(vol float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`vol` = ?", vol).Find(&results).Error

	return
}

// GetBatchFromVol 批量查找 成交量占比 (volume / volume 前5日平均)
func (obj *_SharesDailyTblMgr) GetBatchFromVol(vols []float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`vol` IN (?)", vols).Find(&results).Error

	return
}

// GetFromPe 通过pe获取内容 市盈率
func (obj *_SharesDailyTblMgr) GetFromPe(pe float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`pe` = ?", pe).Find(&results).Error

	return
}

// GetBatchFromPe 批量查找 市盈率
func (obj *_SharesDailyTblMgr) GetBatchFromPe(pes []float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`pe` IN (?)", pes).Find(&results).Error

	return
}

// GetFromPb 通过pb获取内容 市净率
func (obj *_SharesDailyTblMgr) GetFromPb(pb float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`pb` = ?", pb).Find(&results).Error

	return
}

// GetBatchFromPb 批量查找 市净率
func (obj *_SharesDailyTblMgr) GetBatchFromPb(pbs []float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`pb` IN (?)", pbs).Find(&results).Error

	return
}

// GetFromTotal 通过total获取内容 流通市值
func (obj *_SharesDailyTblMgr) GetFromTotal(total float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`total` = ?", total).Find(&results).Error

	return
}

// GetBatchFromTotal 批量查找 流通市值
func (obj *_SharesDailyTblMgr) GetBatchFromTotal(totals []float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`total` IN (?)", totals).Find(&results).Error

	return
}

// GetFromCap 通过cap获取内容 总市值
func (obj *_SharesDailyTblMgr) GetFromCap(cap float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`cap` = ?", cap).Find(&results).Error

	return
}

// GetBatchFromCap 批量查找 总市值
func (obj *_SharesDailyTblMgr) GetBatchFromCap(caps []float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`cap` IN (?)", caps).Find(&results).Error

	return
}

// GetFromZljlr 通过zljlr获取内容 主力净流入
func (obj *_SharesDailyTblMgr) GetFromZljlr(zljlr float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`zljlr` = ?", zljlr).Find(&results).Error

	return
}

// GetBatchFromZljlr 批量查找 主力净流入
func (obj *_SharesDailyTblMgr) GetBatchFromZljlr(zljlrs []float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`zljlr` IN (?)", zljlrs).Find(&results).Error

	return
}

// GetFromOpenPrice 通过open_price获取内容 开盘价
func (obj *_SharesDailyTblMgr) GetFromOpenPrice(openPrice float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`open_price` = ?", openPrice).Find(&results).Error

	return
}

// GetBatchFromOpenPrice 批量查找 开盘价
func (obj *_SharesDailyTblMgr) GetBatchFromOpenPrice(openPrices []float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`open_price` IN (?)", openPrices).Find(&results).Error

	return
}

// GetFromClosePrice 通过close_price获取内容 收盘价
func (obj *_SharesDailyTblMgr) GetFromClosePrice(closePrice float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`close_price` = ?", closePrice).Find(&results).Error

	return
}

// GetBatchFromClosePrice 批量查找 收盘价
func (obj *_SharesDailyTblMgr) GetBatchFromClosePrice(closePrices []float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`close_price` IN (?)", closePrices).Find(&results).Error

	return
}

// GetFromBscg 通过bscg获取内容 北上持股数
func (obj *_SharesDailyTblMgr) GetFromBscg(bscg float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`bscg` = ?", bscg).Find(&results).Error

	return
}

// GetBatchFromBscg 批量查找 北上持股数
func (obj *_SharesDailyTblMgr) GetBatchFromBscg(bscgs []float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`bscg` IN (?)", bscgs).Find(&results).Error

	return
}

// GetFromBsjlr 通过bsjlr获取内容 北上净流入
func (obj *_SharesDailyTblMgr) GetFromBsjlr(bsjlr float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`bsjlr` = ?", bsjlr).Find(&results).Error

	return
}

// GetBatchFromBsjlr 批量查找 北上净流入
func (obj *_SharesDailyTblMgr) GetBatchFromBsjlr(bsjlrs []float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`bsjlr` IN (?)", bsjlrs).Find(&results).Error

	return
}

// GetFromBsPercent 通过bs_percent获取内容 北上持股占比
func (obj *_SharesDailyTblMgr) GetFromBsPercent(bsPercent float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`bs_percent` = ?", bsPercent).Find(&results).Error

	return
}

// GetBatchFromBsPercent 批量查找 北上持股占比
func (obj *_SharesDailyTblMgr) GetBatchFromBsPercent(bsPercents []float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`bs_percent` IN (?)", bsPercents).Find(&results).Error

	return
}

// GetFromMacd 通过macd获取内容 macd 指标
func (obj *_SharesDailyTblMgr) GetFromMacd(macd float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`macd` = ?", macd).Find(&results).Error

	return
}

// GetBatchFromMacd 批量查找 macd 指标
func (obj *_SharesDailyTblMgr) GetBatchFromMacd(macds []float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`macd` IN (?)", macds).Find(&results).Error

	return
}

// GetFromDif 通过dif获取内容 DIF快线
func (obj *_SharesDailyTblMgr) GetFromDif(dif float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`dif` = ?", dif).Find(&results).Error

	return
}

// GetBatchFromDif 批量查找 DIF快线
func (obj *_SharesDailyTblMgr) GetBatchFromDif(difs []float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`dif` IN (?)", difs).Find(&results).Error

	return
}

// GetFromDea 通过dea获取内容 Dea 慢线
func (obj *_SharesDailyTblMgr) GetFromDea(dea float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`dea` = ?", dea).Find(&results).Error

	return
}

// GetBatchFromDea 批量查找 Dea 慢线
func (obj *_SharesDailyTblMgr) GetBatchFromDea(deas []float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`dea` IN (?)", deas).Find(&results).Error

	return
}

// GetFromMax 通过max获取内容 当日最高值
func (obj *_SharesDailyTblMgr) GetFromMax(max float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`max` = ?", max).Find(&results).Error

	return
}

// GetBatchFromMax 批量查找 当日最高值
func (obj *_SharesDailyTblMgr) GetBatchFromMax(maxs []float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`max` IN (?)", maxs).Find(&results).Error

	return
}

// GetFromMin 通过min获取内容 当日最低价
func (obj *_SharesDailyTblMgr) GetFromMin(min float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`min` = ?", min).Find(&results).Error

	return
}

// GetBatchFromMin 批量查找 当日最低价
func (obj *_SharesDailyTblMgr) GetBatchFromMin(mins []float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`min` IN (?)", mins).Find(&results).Error

	return
}

// GetFromPeg 通过peg获取内容 peg估值
func (obj *_SharesDailyTblMgr) GetFromPeg(peg float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`peg` = ?", peg).Find(&results).Error

	return
}

// GetBatchFromPeg 批量查找 peg估值
func (obj *_SharesDailyTblMgr) GetBatchFromPeg(pegs []float64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`peg` IN (?)", pegs).Find(&results).Error

	return
}

// GetFromCreatedBy 通过created_by获取内容 创建者
func (obj *_SharesDailyTblMgr) GetFromCreatedBy(createdBy string) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`created_by` = ?", createdBy).Find(&results).Error

	return
}

// GetBatchFromCreatedBy 批量查找 创建者
func (obj *_SharesDailyTblMgr) GetBatchFromCreatedBy(createdBys []string) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`created_by` IN (?)", createdBys).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_SharesDailyTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_SharesDailyTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SharesDailyTblMgr) FetchByPrimaryKey(id int) (result SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueIndexByCode primary or index 获取唯一内容
func (obj *_SharesDailyTblMgr) FetchUniqueIndexByCode(code string, day0 int64) (result SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`code` = ? AND `day0` = ?", code, day0).First(&result).Error

	return
}

// FetchIndexByDay0  获取多个内容
func (obj *_SharesDailyTblMgr) FetchIndexByDay0(day0 int64) (results []*SharesDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesDailyTbl{}).Where("`day0` = ?", day0).Find(&results).Error

	return
}
