package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _APITblMgr struct {
	*_BaseMgr
}

// APITblMgr open func
func APITblMgr(db *gorm.DB) *_APITblMgr {
	if db == nil {
		panic(fmt.Errorf("APITblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_APITblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("api_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_APITblMgr) GetTableName() string {
	return "api_tbl"
}

// Reset 重置gorm会话
func (obj *_APITblMgr) Reset() *_APITblMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_APITblMgr) Get() (result APITbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(APITbl{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_APITblMgr) Gets() (results []*APITbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(APITbl{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_APITblMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(APITbl{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_APITblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithAPIID api_id获取 授权应用id
func (obj *_APITblMgr) WithAPIID(apiID string) Option {
	return optionFunc(func(o *options) { o.query["api_id"] = apiID })
}

// WithAPIKey api_key获取 授权应用key
func (obj *_APITblMgr) WithAPIKey(apiKey string) Option {
	return optionFunc(func(o *options) { o.query["api_key"] = apiKey })
}

// WithAPISecret api_secret获取 秘钥/token
func (obj *_APITblMgr) WithAPISecret(apiSecret string) Option {
	return optionFunc(func(o *options) { o.query["api_secret"] = apiSecret })
}

// WithTag tag获取 标记
func (obj *_APITblMgr) WithTag(tag string) Option {
	return optionFunc(func(o *options) { o.query["tag"] = tag })
}

// WithAccessToken access_token获取 访问令牌
func (obj *_APITblMgr) WithAccessToken(accessToken string) Option {
	return optionFunc(func(o *options) { o.query["access_token"] = accessToken })
}

// WithExpires expires获取 过期时间
func (obj *_APITblMgr) WithExpires(expires time.Time) Option {
	return optionFunc(func(o *options) { o.query["expires"] = expires })
}

// WithIgnore ignore获取 忽略的词性
func (obj *_APITblMgr) WithIgnore(ignore string) Option {
	return optionFunc(func(o *options) { o.query["ignore"] = ignore })
}

// WithIgsTag igs_tag获取 忽略关键词
func (obj *_APITblMgr) WithIgsTag(igsTag string) Option {
	return optionFunc(func(o *options) { o.query["igs_tag"] = igsTag })
}

// GetByOption 功能选项模式获取
func (obj *_APITblMgr) GetByOption(opts ...Option) (result APITbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(APITbl{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_APITblMgr) GetByOptions(opts ...Option) (results []*APITbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(APITbl{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_APITblMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]APITbl, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(APITbl{}).Where(options.query)
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
func (obj *_APITblMgr) GetFromID(id int) (result APITbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(APITbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_APITblMgr) GetBatchFromID(ids []int) (results []*APITbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(APITbl{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromAPIID 通过api_id获取内容 授权应用id
func (obj *_APITblMgr) GetFromAPIID(apiID string) (results []*APITbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(APITbl{}).Where("`api_id` = ?", apiID).Find(&results).Error

	return
}

// GetBatchFromAPIID 批量查找 授权应用id
func (obj *_APITblMgr) GetBatchFromAPIID(apiIDs []string) (results []*APITbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(APITbl{}).Where("`api_id` IN (?)", apiIDs).Find(&results).Error

	return
}

// GetFromAPIKey 通过api_key获取内容 授权应用key
func (obj *_APITblMgr) GetFromAPIKey(apiKey string) (results []*APITbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(APITbl{}).Where("`api_key` = ?", apiKey).Find(&results).Error

	return
}

// GetBatchFromAPIKey 批量查找 授权应用key
func (obj *_APITblMgr) GetBatchFromAPIKey(apiKeys []string) (results []*APITbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(APITbl{}).Where("`api_key` IN (?)", apiKeys).Find(&results).Error

	return
}

// GetFromAPISecret 通过api_secret获取内容 秘钥/token
func (obj *_APITblMgr) GetFromAPISecret(apiSecret string) (results []*APITbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(APITbl{}).Where("`api_secret` = ?", apiSecret).Find(&results).Error

	return
}

// GetBatchFromAPISecret 批量查找 秘钥/token
func (obj *_APITblMgr) GetBatchFromAPISecret(apiSecrets []string) (results []*APITbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(APITbl{}).Where("`api_secret` IN (?)", apiSecrets).Find(&results).Error

	return
}

// GetFromTag 通过tag获取内容 标记
func (obj *_APITblMgr) GetFromTag(tag string) (results []*APITbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(APITbl{}).Where("`tag` = ?", tag).Find(&results).Error

	return
}

// GetBatchFromTag 批量查找 标记
func (obj *_APITblMgr) GetBatchFromTag(tags []string) (results []*APITbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(APITbl{}).Where("`tag` IN (?)", tags).Find(&results).Error

	return
}

// GetFromAccessToken 通过access_token获取内容 访问令牌
func (obj *_APITblMgr) GetFromAccessToken(accessToken string) (results []*APITbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(APITbl{}).Where("`access_token` = ?", accessToken).Find(&results).Error

	return
}

// GetBatchFromAccessToken 批量查找 访问令牌
func (obj *_APITblMgr) GetBatchFromAccessToken(accessTokens []string) (results []*APITbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(APITbl{}).Where("`access_token` IN (?)", accessTokens).Find(&results).Error

	return
}

// GetFromExpires 通过expires获取内容 过期时间
func (obj *_APITblMgr) GetFromExpires(expires time.Time) (results []*APITbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(APITbl{}).Where("`expires` = ?", expires).Find(&results).Error

	return
}

// GetBatchFromExpires 批量查找 过期时间
func (obj *_APITblMgr) GetBatchFromExpires(expiress []time.Time) (results []*APITbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(APITbl{}).Where("`expires` IN (?)", expiress).Find(&results).Error

	return
}

// GetFromIgnore 通过ignore获取内容 忽略的词性
func (obj *_APITblMgr) GetFromIgnore(ignore string) (results []*APITbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(APITbl{}).Where("`ignore` = ?", ignore).Find(&results).Error

	return
}

// GetBatchFromIgnore 批量查找 忽略的词性
func (obj *_APITblMgr) GetBatchFromIgnore(ignores []string) (results []*APITbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(APITbl{}).Where("`ignore` IN (?)", ignores).Find(&results).Error

	return
}

// GetFromIgsTag 通过igs_tag获取内容 忽略关键词
func (obj *_APITblMgr) GetFromIgsTag(igsTag string) (results []*APITbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(APITbl{}).Where("`igs_tag` = ?", igsTag).Find(&results).Error

	return
}

// GetBatchFromIgsTag 批量查找 忽略关键词
func (obj *_APITblMgr) GetBatchFromIgsTag(igsTags []string) (results []*APITbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(APITbl{}).Where("`igs_tag` IN (?)", igsTags).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_APITblMgr) FetchByPrimaryKey(id int) (result APITbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(APITbl{}).Where("`id` = ?", id).First(&result).Error

	return
}
