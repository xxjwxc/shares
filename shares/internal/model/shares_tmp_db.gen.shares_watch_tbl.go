package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SharesWatchTblMgr struct {
	*_BaseMgr
}

// SharesWatchTblMgr open func
func SharesWatchTblMgr(db *gorm.DB) *_SharesWatchTblMgr {
	if db == nil {
		panic(fmt.Errorf("SharesWatchTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SharesWatchTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("shares_watch_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SharesWatchTblMgr) GetTableName() string {
	return "shares_watch_tbl"
}

// Reset 重置gorm会话
func (obj *_SharesWatchTblMgr) Reset() *_SharesWatchTblMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SharesWatchTblMgr) Get() (result SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SharesWatchTblMgr) Gets() (results []*SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SharesWatchTblMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_SharesWatchTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取 股票代码
func (obj *_SharesWatchTblMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithOpenID open_id获取 用户openid
func (obj *_SharesWatchTblMgr) WithOpenID(openID string) Option {
	return optionFunc(func(o *options) { o.query["open_id"] = openID })
}

// WithKdj kdj获取 日线金叉提醒
func (obj *_SharesWatchTblMgr) WithKdj(kdj bool) Option {
	return optionFunc(func(o *options) { o.query["kdj"] = kdj })
}

// WithKdj20 kdj20获取 20日线提醒
func (obj *_SharesWatchTblMgr) WithKdj20(kdj20 bool) Option {
	return optionFunc(func(o *options) { o.query["kdj20"] = kdj20 })
}

// WithSurge surge获取 盘中急涨提醒
func (obj *_SharesWatchTblMgr) WithSurge(surge bool) Option {
	return optionFunc(func(o *options) { o.query["surge"] = surge })
}

// WithSlump slump获取 盘中急跌提醒
func (obj *_SharesWatchTblMgr) WithSlump(slump bool) Option {
	return optionFunc(func(o *options) { o.query["slump"] = slump })
}

// WithAi ai获取 AI智能提醒
func (obj *_SharesWatchTblMgr) WithAi(ai bool) Option {
	return optionFunc(func(o *options) { o.query["ai"] = ai })
}

// WithGroupWi group_wi获取 是否推荐给组织
func (obj *_SharesWatchTblMgr) WithGroupWi(groupWi bool) Option {
	return optionFunc(func(o *options) { o.query["group_wi"] = groupWi })
}

// WithUp up获取 估价涨到
func (obj *_SharesWatchTblMgr) WithUp(up float64) Option {
	return optionFunc(func(o *options) { o.query["up"] = up })
}

// WithUpVaild up_vaild获取 是否提醒
func (obj *_SharesWatchTblMgr) WithUpVaild(upVaild bool) Option {
	return optionFunc(func(o *options) { o.query["up_vaild"] = upVaild })
}

// WithDown down获取 估价跌到
func (obj *_SharesWatchTblMgr) WithDown(down float64) Option {
	return optionFunc(func(o *options) { o.query["down"] = down })
}

// WithDownVaild down_vaild获取 是否提醒
func (obj *_SharesWatchTblMgr) WithDownVaild(downVaild bool) Option {
	return optionFunc(func(o *options) { o.query["down_vaild"] = downVaild })
}

// WithUpPercent up_percent获取 涨幅超百分比
func (obj *_SharesWatchTblMgr) WithUpPercent(upPercent float64) Option {
	return optionFunc(func(o *options) { o.query["up_percent"] = upPercent })
}

// WithUpPercentTo up_percent_to获取 涨幅超(辅助)
func (obj *_SharesWatchTblMgr) WithUpPercentTo(upPercentTo float64) Option {
	return optionFunc(func(o *options) { o.query["up_percent_to"] = upPercentTo })
}

// WithDownPercent down_percent获取 跌幅超百分比
func (obj *_SharesWatchTblMgr) WithDownPercent(downPercent float64) Option {
	return optionFunc(func(o *options) { o.query["down_percent"] = downPercent })
}

// WithDownPercentTo down_percent_to获取 跌幅超(辅助)
func (obj *_SharesWatchTblMgr) WithDownPercentTo(downPercentTo float64) Option {
	return optionFunc(func(o *options) { o.query["down_percent_to"] = downPercentTo })
}

// WithVaild vaild获取 是否暂停提醒
func (obj *_SharesWatchTblMgr) WithVaild(vaild bool) Option {
	return optionFunc(func(o *options) { o.query["vaild"] = vaild })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_SharesWatchTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithExpiresTime expires_time获取 过期时间
func (obj *_SharesWatchTblMgr) WithExpiresTime(expiresTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["expires_time"] = expiresTime })
}

// GetByOption 功能选项模式获取
func (obj *_SharesWatchTblMgr) GetByOption(opts ...Option) (result SharesWatchTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SharesWatchTblMgr) GetByOptions(opts ...Option) (results []*SharesWatchTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_SharesWatchTblMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]SharesWatchTbl, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where(options.query)
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
func (obj *_SharesWatchTblMgr) GetFromID(id int) (result SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_SharesWatchTblMgr) GetBatchFromID(ids []int) (results []*SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 股票代码
func (obj *_SharesWatchTblMgr) GetFromCode(code string) (results []*SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`code` = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量查找 股票代码
func (obj *_SharesWatchTblMgr) GetBatchFromCode(codes []string) (results []*SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromOpenID 通过open_id获取内容 用户openid
func (obj *_SharesWatchTblMgr) GetFromOpenID(openID string) (results []*SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`open_id` = ?", openID).Find(&results).Error

	return
}

// GetBatchFromOpenID 批量查找 用户openid
func (obj *_SharesWatchTblMgr) GetBatchFromOpenID(openIDs []string) (results []*SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`open_id` IN (?)", openIDs).Find(&results).Error

	return
}

// GetFromKdj 通过kdj获取内容 日线金叉提醒
func (obj *_SharesWatchTblMgr) GetFromKdj(kdj bool) (results []*SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`kdj` = ?", kdj).Find(&results).Error

	return
}

// GetBatchFromKdj 批量查找 日线金叉提醒
func (obj *_SharesWatchTblMgr) GetBatchFromKdj(kdjs []bool) (results []*SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`kdj` IN (?)", kdjs).Find(&results).Error

	return
}

// GetFromKdj20 通过kdj20获取内容 20日线提醒
func (obj *_SharesWatchTblMgr) GetFromKdj20(kdj20 bool) (results []*SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`kdj20` = ?", kdj20).Find(&results).Error

	return
}

// GetBatchFromKdj20 批量查找 20日线提醒
func (obj *_SharesWatchTblMgr) GetBatchFromKdj20(kdj20s []bool) (results []*SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`kdj20` IN (?)", kdj20s).Find(&results).Error

	return
}

// GetFromSurge 通过surge获取内容 盘中急涨提醒
func (obj *_SharesWatchTblMgr) GetFromSurge(surge bool) (results []*SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`surge` = ?", surge).Find(&results).Error

	return
}

// GetBatchFromSurge 批量查找 盘中急涨提醒
func (obj *_SharesWatchTblMgr) GetBatchFromSurge(surges []bool) (results []*SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`surge` IN (?)", surges).Find(&results).Error

	return
}

// GetFromSlump 通过slump获取内容 盘中急跌提醒
func (obj *_SharesWatchTblMgr) GetFromSlump(slump bool) (results []*SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`slump` = ?", slump).Find(&results).Error

	return
}

// GetBatchFromSlump 批量查找 盘中急跌提醒
func (obj *_SharesWatchTblMgr) GetBatchFromSlump(slumps []bool) (results []*SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`slump` IN (?)", slumps).Find(&results).Error

	return
}

// GetFromAi 通过ai获取内容 AI智能提醒
func (obj *_SharesWatchTblMgr) GetFromAi(ai bool) (results []*SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`ai` = ?", ai).Find(&results).Error

	return
}

// GetBatchFromAi 批量查找 AI智能提醒
func (obj *_SharesWatchTblMgr) GetBatchFromAi(ais []bool) (results []*SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`ai` IN (?)", ais).Find(&results).Error

	return
}

// GetFromGroupWi 通过group_wi获取内容 是否推荐给组织
func (obj *_SharesWatchTblMgr) GetFromGroupWi(groupWi bool) (results []*SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`group_wi` = ?", groupWi).Find(&results).Error

	return
}

// GetBatchFromGroupWi 批量查找 是否推荐给组织
func (obj *_SharesWatchTblMgr) GetBatchFromGroupWi(groupWis []bool) (results []*SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`group_wi` IN (?)", groupWis).Find(&results).Error

	return
}

// GetFromUp 通过up获取内容 估价涨到
func (obj *_SharesWatchTblMgr) GetFromUp(up float64) (results []*SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`up` = ?", up).Find(&results).Error

	return
}

// GetBatchFromUp 批量查找 估价涨到
func (obj *_SharesWatchTblMgr) GetBatchFromUp(ups []float64) (results []*SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`up` IN (?)", ups).Find(&results).Error

	return
}

// GetFromUpVaild 通过up_vaild获取内容 是否提醒
func (obj *_SharesWatchTblMgr) GetFromUpVaild(upVaild bool) (results []*SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`up_vaild` = ?", upVaild).Find(&results).Error

	return
}

// GetBatchFromUpVaild 批量查找 是否提醒
func (obj *_SharesWatchTblMgr) GetBatchFromUpVaild(upVailds []bool) (results []*SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`up_vaild` IN (?)", upVailds).Find(&results).Error

	return
}

// GetFromDown 通过down获取内容 估价跌到
func (obj *_SharesWatchTblMgr) GetFromDown(down float64) (results []*SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`down` = ?", down).Find(&results).Error

	return
}

// GetBatchFromDown 批量查找 估价跌到
func (obj *_SharesWatchTblMgr) GetBatchFromDown(downs []float64) (results []*SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`down` IN (?)", downs).Find(&results).Error

	return
}

// GetFromDownVaild 通过down_vaild获取内容 是否提醒
func (obj *_SharesWatchTblMgr) GetFromDownVaild(downVaild bool) (results []*SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`down_vaild` = ?", downVaild).Find(&results).Error

	return
}

// GetBatchFromDownVaild 批量查找 是否提醒
func (obj *_SharesWatchTblMgr) GetBatchFromDownVaild(downVailds []bool) (results []*SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`down_vaild` IN (?)", downVailds).Find(&results).Error

	return
}

// GetFromUpPercent 通过up_percent获取内容 涨幅超百分比
func (obj *_SharesWatchTblMgr) GetFromUpPercent(upPercent float64) (results []*SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`up_percent` = ?", upPercent).Find(&results).Error

	return
}

// GetBatchFromUpPercent 批量查找 涨幅超百分比
func (obj *_SharesWatchTblMgr) GetBatchFromUpPercent(upPercents []float64) (results []*SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`up_percent` IN (?)", upPercents).Find(&results).Error

	return
}

// GetFromUpPercentTo 通过up_percent_to获取内容 涨幅超(辅助)
func (obj *_SharesWatchTblMgr) GetFromUpPercentTo(upPercentTo float64) (results []*SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`up_percent_to` = ?", upPercentTo).Find(&results).Error

	return
}

// GetBatchFromUpPercentTo 批量查找 涨幅超(辅助)
func (obj *_SharesWatchTblMgr) GetBatchFromUpPercentTo(upPercentTos []float64) (results []*SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`up_percent_to` IN (?)", upPercentTos).Find(&results).Error

	return
}

// GetFromDownPercent 通过down_percent获取内容 跌幅超百分比
func (obj *_SharesWatchTblMgr) GetFromDownPercent(downPercent float64) (results []*SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`down_percent` = ?", downPercent).Find(&results).Error

	return
}

// GetBatchFromDownPercent 批量查找 跌幅超百分比
func (obj *_SharesWatchTblMgr) GetBatchFromDownPercent(downPercents []float64) (results []*SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`down_percent` IN (?)", downPercents).Find(&results).Error

	return
}

// GetFromDownPercentTo 通过down_percent_to获取内容 跌幅超(辅助)
func (obj *_SharesWatchTblMgr) GetFromDownPercentTo(downPercentTo float64) (results []*SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`down_percent_to` = ?", downPercentTo).Find(&results).Error

	return
}

// GetBatchFromDownPercentTo 批量查找 跌幅超(辅助)
func (obj *_SharesWatchTblMgr) GetBatchFromDownPercentTo(downPercentTos []float64) (results []*SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`down_percent_to` IN (?)", downPercentTos).Find(&results).Error

	return
}

// GetFromVaild 通过vaild获取内容 是否暂停提醒
func (obj *_SharesWatchTblMgr) GetFromVaild(vaild bool) (results []*SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`vaild` = ?", vaild).Find(&results).Error

	return
}

// GetBatchFromVaild 批量查找 是否暂停提醒
func (obj *_SharesWatchTblMgr) GetBatchFromVaild(vailds []bool) (results []*SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`vaild` IN (?)", vailds).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_SharesWatchTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_SharesWatchTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromExpiresTime 通过expires_time获取内容 过期时间
func (obj *_SharesWatchTblMgr) GetFromExpiresTime(expiresTime time.Time) (results []*SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`expires_time` = ?", expiresTime).Find(&results).Error

	return
}

// GetBatchFromExpiresTime 批量查找 过期时间
func (obj *_SharesWatchTblMgr) GetBatchFromExpiresTime(expiresTimes []time.Time) (results []*SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`expires_time` IN (?)", expiresTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SharesWatchTblMgr) FetchByPrimaryKey(id int) (result SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueIndexByOpenCode primary or index 获取唯一内容
func (obj *_SharesWatchTblMgr) FetchUniqueIndexByOpenCode(code string, openID string) (result SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`code` = ? AND `open_id` = ?", code, openID).First(&result).Error

	return
}

// FetchIndexByCode  获取多个内容
func (obj *_SharesWatchTblMgr) FetchIndexByCode(code string) (results []*SharesWatchTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SharesWatchTbl{}).Where("`code` = ?", code).Find(&results).Error

	return
}
