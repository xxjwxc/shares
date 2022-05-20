package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _GroupTblMgr struct {
	*_BaseMgr
}

// GroupTblMgr open func
func GroupTblMgr(db *gorm.DB) *_GroupTblMgr {
	if db == nil {
		panic(fmt.Errorf("GroupTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_GroupTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("group_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_GroupTblMgr) GetTableName() string {
	return "group_tbl"
}

// Reset 重置gorm会话
func (obj *_GroupTblMgr) Reset() *_GroupTblMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_GroupTblMgr) Get() (result GroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupTbl{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_GroupTblMgr) Gets() (results []*GroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupTbl{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_GroupTblMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(GroupTbl{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_GroupTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithGroupName group_name获取 分组名
func (obj *_GroupTblMgr) WithGroupName(groupName string) Option {
	return optionFunc(func(o *options) { o.query["group_name"] = groupName })
}

// WithCode code获取 股票代码
func (obj *_GroupTblMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithWi wi获取 推荐权重
func (obj *_GroupTblMgr) WithWi(wi int) Option {
	return optionFunc(func(o *options) { o.query["wi"] = wi })
}

// WithUserName user_name获取 推荐人
func (obj *_GroupTblMgr) WithUserName(userName string) Option {
	return optionFunc(func(o *options) { o.query["user_name"] = userName })
}

// WithCreatedBy created_by获取
func (obj *_GroupTblMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithCreatedAt created_at获取
func (obj *_GroupTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdateAt update_at获取
func (obj *_GroupTblMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// WithUpdateBy update_by获取
func (obj *_GroupTblMgr) WithUpdateBy(updateBy string) Option {
	return optionFunc(func(o *options) { o.query["update_by"] = updateBy })
}

// GetByOption 功能选项模式获取
func (obj *_GroupTblMgr) GetByOption(opts ...Option) (result GroupTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(GroupTbl{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_GroupTblMgr) GetByOptions(opts ...Option) (results []*GroupTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(GroupTbl{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_GroupTblMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]GroupTbl, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(GroupTbl{}).Where(options.query)
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
func (obj *_GroupTblMgr) GetFromID(id int) (result GroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_GroupTblMgr) GetBatchFromID(ids []int) (results []*GroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupTbl{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromGroupName 通过group_name获取内容 分组名
func (obj *_GroupTblMgr) GetFromGroupName(groupName string) (results []*GroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupTbl{}).Where("`group_name` = ?", groupName).Find(&results).Error

	return
}

// GetBatchFromGroupName 批量查找 分组名
func (obj *_GroupTblMgr) GetBatchFromGroupName(groupNames []string) (results []*GroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupTbl{}).Where("`group_name` IN (?)", groupNames).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 股票代码
func (obj *_GroupTblMgr) GetFromCode(code string) (results []*GroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupTbl{}).Where("`code` = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量查找 股票代码
func (obj *_GroupTblMgr) GetBatchFromCode(codes []string) (results []*GroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupTbl{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromWi 通过wi获取内容 推荐权重
func (obj *_GroupTblMgr) GetFromWi(wi int) (results []*GroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupTbl{}).Where("`wi` = ?", wi).Find(&results).Error

	return
}

// GetBatchFromWi 批量查找 推荐权重
func (obj *_GroupTblMgr) GetBatchFromWi(wis []int) (results []*GroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupTbl{}).Where("`wi` IN (?)", wis).Find(&results).Error

	return
}

// GetFromUserName 通过user_name获取内容 推荐人
func (obj *_GroupTblMgr) GetFromUserName(userName string) (results []*GroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupTbl{}).Where("`user_name` = ?", userName).Find(&results).Error

	return
}

// GetBatchFromUserName 批量查找 推荐人
func (obj *_GroupTblMgr) GetBatchFromUserName(userNames []string) (results []*GroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupTbl{}).Where("`user_name` IN (?)", userNames).Find(&results).Error

	return
}

// GetFromCreatedBy 通过created_by获取内容
func (obj *_GroupTblMgr) GetFromCreatedBy(createdBy string) (results []*GroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupTbl{}).Where("`created_by` = ?", createdBy).Find(&results).Error

	return
}

// GetBatchFromCreatedBy 批量查找
func (obj *_GroupTblMgr) GetBatchFromCreatedBy(createdBys []string) (results []*GroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupTbl{}).Where("`created_by` IN (?)", createdBys).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_GroupTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*GroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupTbl{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_GroupTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*GroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupTbl{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容
func (obj *_GroupTblMgr) GetFromUpdateAt(updateAt time.Time) (results []*GroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupTbl{}).Where("`update_at` = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量查找
func (obj *_GroupTblMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*GroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupTbl{}).Where("`update_at` IN (?)", updateAts).Find(&results).Error

	return
}

// GetFromUpdateBy 通过update_by获取内容
func (obj *_GroupTblMgr) GetFromUpdateBy(updateBy string) (results []*GroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupTbl{}).Where("`update_by` = ?", updateBy).Find(&results).Error

	return
}

// GetBatchFromUpdateBy 批量查找
func (obj *_GroupTblMgr) GetBatchFromUpdateBy(updateBys []string) (results []*GroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupTbl{}).Where("`update_by` IN (?)", updateBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_GroupTblMgr) FetchByPrimaryKey(id int) (result GroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueIndexByGroupName primary or index 获取唯一内容
func (obj *_GroupTblMgr) FetchUniqueIndexByGroupName(groupName string, code string) (result GroupTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupTbl{}).Where("`group_name` = ? AND `code` = ?", groupName, code).First(&result).Error

	return
}
