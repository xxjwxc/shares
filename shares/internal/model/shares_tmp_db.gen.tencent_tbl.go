package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _TencentTblMgr struct {
	*_BaseMgr
}

// TencentTblMgr open func
func TencentTblMgr(db *gorm.DB) *_TencentTblMgr {
	if db == nil {
		panic(fmt.Errorf("TencentTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TencentTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("tencent_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TencentTblMgr) GetTableName() string {
	return "tencent_tbl"
}

// Reset 重置gorm会话
func (obj *_TencentTblMgr) Reset() *_TencentTblMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TencentTblMgr) Get() (result TencentTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TencentTbl{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TencentTblMgr) Gets() (results []*TencentTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TencentTbl{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TencentTblMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(TencentTbl{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_TencentTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithAPIID api_id获取 授权应用id
func (obj *_TencentTblMgr) WithAPIID(apiID string) Option {
	return optionFunc(func(o *options) { o.query["api_id"] = apiID })
}

// WithAPIKey api_key获取 授权应用key
func (obj *_TencentTblMgr) WithAPIKey(apiKey string) Option {
	return optionFunc(func(o *options) { o.query["api_key"] = apiKey })
}

// WithAPISecret api_secret获取 秘钥/token
func (obj *_TencentTblMgr) WithAPISecret(apiSecret string) Option {
	return optionFunc(func(o *options) { o.query["api_secret"] = apiSecret })
}

// WithTag tag获取 标记
func (obj *_TencentTblMgr) WithTag(tag string) Option {
	return optionFunc(func(o *options) { o.query["tag"] = tag })
}

// WithIgnore ignore获取 忽略的词性
func (obj *_TencentTblMgr) WithIgnore(ignore string) Option {
	return optionFunc(func(o *options) { o.query["ignore"] = ignore })
}

// WithIgsTag igs_tag获取 忽略关键词
func (obj *_TencentTblMgr) WithIgsTag(igsTag string) Option {
	return optionFunc(func(o *options) { o.query["igs_tag"] = igsTag })
}

// GetByOption 功能选项模式获取
func (obj *_TencentTblMgr) GetByOption(opts ...Option) (result TencentTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TencentTbl{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TencentTblMgr) GetByOptions(opts ...Option) (results []*TencentTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TencentTbl{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_TencentTblMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]TencentTbl, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(TencentTbl{}).Where(options.query)
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
func (obj *_TencentTblMgr) GetFromID(id int) (result TencentTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TencentTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_TencentTblMgr) GetBatchFromID(ids []int) (results []*TencentTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TencentTbl{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromAPIID 通过api_id获取内容 授权应用id
func (obj *_TencentTblMgr) GetFromAPIID(apiID string) (results []*TencentTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TencentTbl{}).Where("`api_id` = ?", apiID).Find(&results).Error

	return
}

// GetBatchFromAPIID 批量查找 授权应用id
func (obj *_TencentTblMgr) GetBatchFromAPIID(apiIDs []string) (results []*TencentTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TencentTbl{}).Where("`api_id` IN (?)", apiIDs).Find(&results).Error

	return
}

// GetFromAPIKey 通过api_key获取内容 授权应用key
func (obj *_TencentTblMgr) GetFromAPIKey(apiKey string) (results []*TencentTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TencentTbl{}).Where("`api_key` = ?", apiKey).Find(&results).Error

	return
}

// GetBatchFromAPIKey 批量查找 授权应用key
func (obj *_TencentTblMgr) GetBatchFromAPIKey(apiKeys []string) (results []*TencentTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TencentTbl{}).Where("`api_key` IN (?)", apiKeys).Find(&results).Error

	return
}

// GetFromAPISecret 通过api_secret获取内容 秘钥/token
func (obj *_TencentTblMgr) GetFromAPISecret(apiSecret string) (results []*TencentTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TencentTbl{}).Where("`api_secret` = ?", apiSecret).Find(&results).Error

	return
}

// GetBatchFromAPISecret 批量查找 秘钥/token
func (obj *_TencentTblMgr) GetBatchFromAPISecret(apiSecrets []string) (results []*TencentTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TencentTbl{}).Where("`api_secret` IN (?)", apiSecrets).Find(&results).Error

	return
}

// GetFromTag 通过tag获取内容 标记
func (obj *_TencentTblMgr) GetFromTag(tag string) (results []*TencentTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TencentTbl{}).Where("`tag` = ?", tag).Find(&results).Error

	return
}

// GetBatchFromTag 批量查找 标记
func (obj *_TencentTblMgr) GetBatchFromTag(tags []string) (results []*TencentTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TencentTbl{}).Where("`tag` IN (?)", tags).Find(&results).Error

	return
}

// GetFromIgnore 通过ignore获取内容 忽略的词性
func (obj *_TencentTblMgr) GetFromIgnore(ignore string) (results []*TencentTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TencentTbl{}).Where("`ignore` = ?", ignore).Find(&results).Error

	return
}

// GetBatchFromIgnore 批量查找 忽略的词性
func (obj *_TencentTblMgr) GetBatchFromIgnore(ignores []string) (results []*TencentTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TencentTbl{}).Where("`ignore` IN (?)", ignores).Find(&results).Error

	return
}

// GetFromIgsTag 通过igs_tag获取内容 忽略关键词
func (obj *_TencentTblMgr) GetFromIgsTag(igsTag string) (results []*TencentTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TencentTbl{}).Where("`igs_tag` = ?", igsTag).Find(&results).Error

	return
}

// GetBatchFromIgsTag 批量查找 忽略关键词
func (obj *_TencentTblMgr) GetBatchFromIgsTag(igsTags []string) (results []*TencentTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TencentTbl{}).Where("`igs_tag` IN (?)", igsTags).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_TencentTblMgr) FetchByPrimaryKey(id int) (result TencentTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TencentTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}
