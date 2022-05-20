package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _MsgCodeViewMgr struct {
	*_BaseMgr
}

// MsgCodeViewMgr open func
func MsgCodeViewMgr(db *gorm.DB) *_MsgCodeViewMgr {
	if db == nil {
		panic(fmt.Errorf("MsgCodeViewMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MsgCodeViewMgr{_BaseMgr: &_BaseMgr{DB: db.Table("msg_code_view"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MsgCodeViewMgr) GetTableName() string {
	return "msg_code_view"
}

// Reset 重置gorm会话
func (obj *_MsgCodeViewMgr) Reset() *_MsgCodeViewMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MsgCodeViewMgr) Get() (result MsgCodeView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgCodeView{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MsgCodeViewMgr) Gets() (results []*MsgCodeView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgCodeView{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MsgCodeViewMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(MsgCodeView{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithCode code获取
func (obj *_MsgCodeViewMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithTp tp获取
func (obj *_MsgCodeViewMgr) WithTp(tp string) Option {
	return optionFunc(func(o *options) { o.query["tp"] = tp })
}

// GetByOption 功能选项模式获取
func (obj *_MsgCodeViewMgr) GetByOption(opts ...Option) (result MsgCodeView, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MsgCodeView{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MsgCodeViewMgr) GetByOptions(opts ...Option) (results []*MsgCodeView, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MsgCodeView{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_MsgCodeViewMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]MsgCodeView, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(MsgCodeView{}).Where(options.query)
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

// GetFromCode 通过code获取内容
func (obj *_MsgCodeViewMgr) GetFromCode(code string) (results []*MsgCodeView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgCodeView{}).Where("`code` = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量查找
func (obj *_MsgCodeViewMgr) GetBatchFromCode(codes []string) (results []*MsgCodeView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgCodeView{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromTp 通过tp获取内容
func (obj *_MsgCodeViewMgr) GetFromTp(tp string) (results []*MsgCodeView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgCodeView{}).Where("`tp` = ?", tp).Find(&results).Error

	return
}

// GetBatchFromTp 批量查找
func (obj *_MsgCodeViewMgr) GetBatchFromTp(tps []string) (results []*MsgCodeView, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MsgCodeView{}).Where("`tp` IN (?)", tps).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
