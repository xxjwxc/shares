package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _TdxDailyTblMgr struct {
	*_BaseMgr
}

// TdxDailyTblMgr open func
func TdxDailyTblMgr(db *gorm.DB) *_TdxDailyTblMgr {
	if db == nil {
		panic(fmt.Errorf("TdxDailyTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TdxDailyTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("tdx_daily_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TdxDailyTblMgr) GetTableName() string {
	return "tdx_daily_tbl"
}

// Reset 重置gorm会话
func (obj *_TdxDailyTblMgr) Reset() *_TdxDailyTblMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TdxDailyTblMgr) Get() (result TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TdxDailyTblMgr) Gets() (results []*TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TdxDailyTblMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_TdxDailyTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取 股票代码
func (obj *_TdxDailyTblMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithName name获取 股票名字
func (obj *_TdxDailyTblMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDay0 day0获取 当日0点时间戳
func (obj *_TdxDailyTblMgr) WithDay0(day0 int64) Option {
	return optionFunc(func(o *options) { o.query["day0"] = day0 })
}

// WithDayStr day_str获取 入池时间
func (obj *_TdxDailyTblMgr) WithDayStr(dayStr string) Option {
	return optionFunc(func(o *options) { o.query["day_str"] = dayStr })
}

// WithOpen open获取 当日开盘价
func (obj *_TdxDailyTblMgr) WithOpen(open float64) Option {
	return optionFunc(func(o *options) { o.query["open"] = open })
}

// WithClose close获取 当日收盘价
func (obj *_TdxDailyTblMgr) WithClose(close float64) Option {
	return optionFunc(func(o *options) { o.query["close"] = close })
}

// WithMin min获取 当日最低价
func (obj *_TdxDailyTblMgr) WithMin(min float64) Option {
	return optionFunc(func(o *options) { o.query["min"] = min })
}

// WithMax max获取 当日最高值
func (obj *_TdxDailyTblMgr) WithMax(max float64) Option {
	return optionFunc(func(o *options) { o.query["max"] = max })
}

// WithVolume volume获取 成交量（手）
func (obj *_TdxDailyTblMgr) WithVolume(volume float64) Option {
	return optionFunc(func(o *options) { o.query["volume"] = volume })
}

// WithFyyx fyyx获取 飞鹰优选
func (obj *_TdxDailyTblMgr) WithFyyx(fyyx bool) Option {
	return optionFunc(func(o *options) { o.query["fyyx"] = fyyx })
}

// WithZlzxh zlzxh获取 主力真吸货
func (obj *_TdxDailyTblMgr) WithZlzxh(zlzxh float64) Option {
	return optionFunc(func(o *options) { o.query["zlzxh"] = zlzxh })
}

// WithFx fx获取 风险(>75 风险较大，<10 风险较小)
func (obj *_TdxDailyTblMgr) WithFx(fx float64) Option {
	return optionFunc(func(o *options) { o.query["fx"] = fx })
}

// WithSmx smx获取 生命线
func (obj *_TdxDailyTblMgr) WithSmx(smx float64) Option {
	return optionFunc(func(o *options) { o.query["smx"] = smx })
}

// WithDwzq dwzq获取 低位转强(50)
func (obj *_TdxDailyTblMgr) WithDwzq(dwzq float64) Option {
	return optionFunc(func(o *options) { o.query["dwzq"] = dwzq })
}

// WithKsls ksls获取 开始拉升(50)
func (obj *_TdxDailyTblMgr) WithKsls(ksls float64) Option {
	return optionFunc(func(o *options) { o.query["ksls"] = ksls })
}

// WithJdz jdz获取 极低涨(50)
func (obj *_TdxDailyTblMgr) WithJdz(jdz float64) Option {
	return optionFunc(func(o *options) { o.query["jdz"] = jdz })
}

// WithTqa tqa获取 唐奇安交易通道指标：-1卖，1买
func (obj *_TdxDailyTblMgr) WithTqa(tqa int) Option {
	return optionFunc(func(o *options) { o.query["tqa"] = tqa })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_TdxDailyTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdateAt update_at获取 更新时间
func (obj *_TdxDailyTblMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// GetByOption 功能选项模式获取
func (obj *_TdxDailyTblMgr) GetByOption(opts ...Option) (result TdxDailyTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TdxDailyTblMgr) GetByOptions(opts ...Option) (results []*TdxDailyTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_TdxDailyTblMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]TdxDailyTbl, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where(options.query)
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
func (obj *_TdxDailyTblMgr) GetFromID(id int) (result TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_TdxDailyTblMgr) GetBatchFromID(ids []int) (results []*TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 股票代码
func (obj *_TdxDailyTblMgr) GetFromCode(code string) (results []*TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`code` = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量查找 股票代码
func (obj *_TdxDailyTblMgr) GetBatchFromCode(codes []string) (results []*TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 股票名字
func (obj *_TdxDailyTblMgr) GetFromName(name string) (results []*TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 股票名字
func (obj *_TdxDailyTblMgr) GetBatchFromName(names []string) (results []*TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromDay0 通过day0获取内容 当日0点时间戳
func (obj *_TdxDailyTblMgr) GetFromDay0(day0 int64) (results []*TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`day0` = ?", day0).Find(&results).Error

	return
}

// GetBatchFromDay0 批量查找 当日0点时间戳
func (obj *_TdxDailyTblMgr) GetBatchFromDay0(day0s []int64) (results []*TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`day0` IN (?)", day0s).Find(&results).Error

	return
}

// GetFromDayStr 通过day_str获取内容 入池时间
func (obj *_TdxDailyTblMgr) GetFromDayStr(dayStr string) (results []*TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`day_str` = ?", dayStr).Find(&results).Error

	return
}

// GetBatchFromDayStr 批量查找 入池时间
func (obj *_TdxDailyTblMgr) GetBatchFromDayStr(dayStrs []string) (results []*TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`day_str` IN (?)", dayStrs).Find(&results).Error

	return
}

// GetFromOpen 通过open获取内容 当日开盘价
func (obj *_TdxDailyTblMgr) GetFromOpen(open float64) (results []*TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`open` = ?", open).Find(&results).Error

	return
}

// GetBatchFromOpen 批量查找 当日开盘价
func (obj *_TdxDailyTblMgr) GetBatchFromOpen(opens []float64) (results []*TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`open` IN (?)", opens).Find(&results).Error

	return
}

// GetFromClose 通过close获取内容 当日收盘价
func (obj *_TdxDailyTblMgr) GetFromClose(close float64) (results []*TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`close` = ?", close).Find(&results).Error

	return
}

// GetBatchFromClose 批量查找 当日收盘价
func (obj *_TdxDailyTblMgr) GetBatchFromClose(closes []float64) (results []*TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`close` IN (?)", closes).Find(&results).Error

	return
}

// GetFromMin 通过min获取内容 当日最低价
func (obj *_TdxDailyTblMgr) GetFromMin(min float64) (results []*TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`min` = ?", min).Find(&results).Error

	return
}

// GetBatchFromMin 批量查找 当日最低价
func (obj *_TdxDailyTblMgr) GetBatchFromMin(mins []float64) (results []*TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`min` IN (?)", mins).Find(&results).Error

	return
}

// GetFromMax 通过max获取内容 当日最高值
func (obj *_TdxDailyTblMgr) GetFromMax(max float64) (results []*TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`max` = ?", max).Find(&results).Error

	return
}

// GetBatchFromMax 批量查找 当日最高值
func (obj *_TdxDailyTblMgr) GetBatchFromMax(maxs []float64) (results []*TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`max` IN (?)", maxs).Find(&results).Error

	return
}

// GetFromVolume 通过volume获取内容 成交量（手）
func (obj *_TdxDailyTblMgr) GetFromVolume(volume float64) (results []*TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`volume` = ?", volume).Find(&results).Error

	return
}

// GetBatchFromVolume 批量查找 成交量（手）
func (obj *_TdxDailyTblMgr) GetBatchFromVolume(volumes []float64) (results []*TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`volume` IN (?)", volumes).Find(&results).Error

	return
}

// GetFromFyyx 通过fyyx获取内容 飞鹰优选
func (obj *_TdxDailyTblMgr) GetFromFyyx(fyyx bool) (results []*TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`fyyx` = ?", fyyx).Find(&results).Error

	return
}

// GetBatchFromFyyx 批量查找 飞鹰优选
func (obj *_TdxDailyTblMgr) GetBatchFromFyyx(fyyxs []bool) (results []*TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`fyyx` IN (?)", fyyxs).Find(&results).Error

	return
}

// GetFromZlzxh 通过zlzxh获取内容 主力真吸货
func (obj *_TdxDailyTblMgr) GetFromZlzxh(zlzxh float64) (results []*TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`zlzxh` = ?", zlzxh).Find(&results).Error

	return
}

// GetBatchFromZlzxh 批量查找 主力真吸货
func (obj *_TdxDailyTblMgr) GetBatchFromZlzxh(zlzxhs []float64) (results []*TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`zlzxh` IN (?)", zlzxhs).Find(&results).Error

	return
}

// GetFromFx 通过fx获取内容 风险(>75 风险较大，<10 风险较小)
func (obj *_TdxDailyTblMgr) GetFromFx(fx float64) (results []*TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`fx` = ?", fx).Find(&results).Error

	return
}

// GetBatchFromFx 批量查找 风险(>75 风险较大，<10 风险较小)
func (obj *_TdxDailyTblMgr) GetBatchFromFx(fxs []float64) (results []*TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`fx` IN (?)", fxs).Find(&results).Error

	return
}

// GetFromSmx 通过smx获取内容 生命线
func (obj *_TdxDailyTblMgr) GetFromSmx(smx float64) (results []*TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`smx` = ?", smx).Find(&results).Error

	return
}

// GetBatchFromSmx 批量查找 生命线
func (obj *_TdxDailyTblMgr) GetBatchFromSmx(smxs []float64) (results []*TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`smx` IN (?)", smxs).Find(&results).Error

	return
}

// GetFromDwzq 通过dwzq获取内容 低位转强(50)
func (obj *_TdxDailyTblMgr) GetFromDwzq(dwzq float64) (results []*TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`dwzq` = ?", dwzq).Find(&results).Error

	return
}

// GetBatchFromDwzq 批量查找 低位转强(50)
func (obj *_TdxDailyTblMgr) GetBatchFromDwzq(dwzqs []float64) (results []*TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`dwzq` IN (?)", dwzqs).Find(&results).Error

	return
}

// GetFromKsls 通过ksls获取内容 开始拉升(50)
func (obj *_TdxDailyTblMgr) GetFromKsls(ksls float64) (results []*TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`ksls` = ?", ksls).Find(&results).Error

	return
}

// GetBatchFromKsls 批量查找 开始拉升(50)
func (obj *_TdxDailyTblMgr) GetBatchFromKsls(kslss []float64) (results []*TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`ksls` IN (?)", kslss).Find(&results).Error

	return
}

// GetFromJdz 通过jdz获取内容 极低涨(50)
func (obj *_TdxDailyTblMgr) GetFromJdz(jdz float64) (results []*TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`jdz` = ?", jdz).Find(&results).Error

	return
}

// GetBatchFromJdz 批量查找 极低涨(50)
func (obj *_TdxDailyTblMgr) GetBatchFromJdz(jdzs []float64) (results []*TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`jdz` IN (?)", jdzs).Find(&results).Error

	return
}

// GetFromTqa 通过tqa获取内容 唐奇安交易通道指标：-1卖，1买
func (obj *_TdxDailyTblMgr) GetFromTqa(tqa int) (results []*TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`tqa` = ?", tqa).Find(&results).Error

	return
}

// GetBatchFromTqa 批量查找 唐奇安交易通道指标：-1卖，1买
func (obj *_TdxDailyTblMgr) GetBatchFromTqa(tqas []int) (results []*TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`tqa` IN (?)", tqas).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_TdxDailyTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_TdxDailyTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容 更新时间
func (obj *_TdxDailyTblMgr) GetFromUpdateAt(updateAt time.Time) (results []*TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`update_at` = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量查找 更新时间
func (obj *_TdxDailyTblMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`update_at` IN (?)", updateAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_TdxDailyTblMgr) FetchByPrimaryKey(id int) (result TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueIndexByCode primary or index 获取唯一内容
func (obj *_TdxDailyTblMgr) FetchUniqueIndexByCode(code string, day0 int64) (result TdxDailyTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TdxDailyTbl{}).Where("`code` = ? AND `day0` = ?", code, day0).First(&result).Error

	return
}
