package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _GroupListTblMgr struct {
	*_BaseMgr
}

// GroupListTblMgr open func
func GroupListTblMgr(db *gorm.DB) *_GroupListTblMgr {
	if db == nil {
		panic(fmt.Errorf("GroupListTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_GroupListTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("group_list_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_GroupListTblMgr) GetTableName() string {
	return "group_list_tbl"
}

// Reset 重置gorm会话
func (obj *_GroupListTblMgr) Reset() *_GroupListTblMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_GroupListTblMgr) Get() (result GroupListTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupListTbl{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_GroupListTblMgr) Gets() (results []*GroupListTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupListTbl{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_GroupListTblMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(GroupListTbl{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_GroupListTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithKey key获取 搜索代码
func (obj *_GroupListTblMgr) WithKey(key string) Option {
	return optionFunc(func(o *options) { o.query["key"] = key })
}

// WithGroupName group_name获取 分组名
func (obj *_GroupListTblMgr) WithGroupName(groupName string) Option {
	return optionFunc(func(o *options) { o.query["group_name"] = groupName })
}

// GetByOption 功能选项模式获取
func (obj *_GroupListTblMgr) GetByOption(opts ...Option) (result GroupListTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(GroupListTbl{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_GroupListTblMgr) GetByOptions(opts ...Option) (results []*GroupListTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(GroupListTbl{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_GroupListTblMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]GroupListTbl, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(GroupListTbl{}).Where(options.query)
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
func (obj *_GroupListTblMgr) GetFromID(id int) (result GroupListTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupListTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_GroupListTblMgr) GetBatchFromID(ids []int) (results []*GroupListTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupListTbl{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromKey 通过key获取内容 搜索代码
func (obj *_GroupListTblMgr) GetFromKey(key string) (result GroupListTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupListTbl{}).Where("`key` = ?", key).First(&result).Error

	return
}

// GetBatchFromKey 批量查找 搜索代码
func (obj *_GroupListTblMgr) GetBatchFromKey(keys []string) (results []*GroupListTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupListTbl{}).Where("`key` IN (?)", keys).Find(&results).Error

	return
}

// GetFromGroupName 通过group_name获取内容 分组名
func (obj *_GroupListTblMgr) GetFromGroupName(groupName string) (result GroupListTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupListTbl{}).Where("`group_name` = ?", groupName).First(&result).Error

	return
}

// GetBatchFromGroupName 批量查找 分组名
func (obj *_GroupListTblMgr) GetBatchFromGroupName(groupNames []string) (results []*GroupListTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupListTbl{}).Where("`group_name` IN (?)", groupNames).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_GroupListTblMgr) FetchByPrimaryKey(id int) (result GroupListTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupListTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByKey primary or index 获取唯一内容
func (obj *_GroupListTblMgr) FetchUniqueByKey(key string) (result GroupListTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupListTbl{}).Where("`key` = ?", key).First(&result).Error

	return
}

// FetchUniqueByGroupName primary or index 获取唯一内容
func (obj *_GroupListTblMgr) FetchUniqueByGroupName(groupName string) (result GroupListTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GroupListTbl{}).Where("`group_name` = ?", groupName).First(&result).Error

	return
}
