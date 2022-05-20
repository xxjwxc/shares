package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _BillTblMgr struct {
	*_BaseMgr
}

// BillTblMgr open func
func BillTblMgr(db *gorm.DB) *_BillTblMgr {
	if db == nil {
		panic(fmt.Errorf("BillTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_BillTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("bill_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_BillTblMgr) GetTableName() string {
	return "bill_tbl"
}

// Reset 重置gorm会话
func (obj *_BillTblMgr) Reset() *_BillTblMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_BillTblMgr) Get() (result BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BillTbl{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_BillTblMgr) Gets() (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BillTbl{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_BillTblMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(BillTbl{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_BillTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUserID user_id获取 用户id
func (obj *_BillTblMgr) WithUserID(userID string) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithBillID bill_id获取 账单id
func (obj *_BillTblMgr) WithBillID(billID string) Option {
	return optionFunc(func(o *options) { o.query["bill_id"] = billID })
}

// WithPayPlatform pay_platform获取 支付类型(1:微信支付)
func (obj *_BillTblMgr) WithPayPlatform(payPlatform int) Option {
	return optionFunc(func(o *options) { o.query["pay_platform"] = payPlatform })
}

// WithPayAmount pay_amount获取 支付金额
func (obj *_BillTblMgr) WithPayAmount(payAmount int64) Option {
	return optionFunc(func(o *options) { o.query["pay_amount"] = payAmount })
}

// WithStatus status获取 支付状态(-1:已退款,1:待支付,2:已支付,待发货,3:已取消,4:待收货,5:已完成，6:售后待评价)
func (obj *_BillTblMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithDesc desc获取 订单备注
func (obj *_BillTblMgr) WithDesc(desc string) Option {
	return optionFunc(func(o *options) { o.query["desc"] = desc })
}

// WithCreatedBy created_by获取 创建者
func (obj *_BillTblMgr) WithCreatedBy(createdBy string) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_BillTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedBy updated_by获取 更新者
func (obj *_BillTblMgr) WithUpdatedBy(updatedBy string) Option {
	return optionFunc(func(o *options) { o.query["updated_by"] = updatedBy })
}

// WithUpdatedAt updated_at获取 更新时间
func (obj *_BillTblMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedBy deleted_by获取 删除者
func (obj *_BillTblMgr) WithDeletedBy(deletedBy string) Option {
	return optionFunc(func(o *options) { o.query["deleted_by"] = deletedBy })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_BillTblMgr) WithDeletedAt(deletedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// GetByOption 功能选项模式获取
func (obj *_BillTblMgr) GetByOption(opts ...Option) (result BillTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(BillTbl{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_BillTblMgr) GetByOptions(opts ...Option) (results []*BillTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(BillTbl{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_BillTblMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]BillTbl, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(BillTbl{}).Where(options.query)
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
func (obj *_BillTblMgr) GetFromID(id int) (result BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BillTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_BillTblMgr) GetBatchFromID(ids []int) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BillTbl{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 用户id
func (obj *_BillTblMgr) GetFromUserID(userID string) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BillTbl{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找 用户id
func (obj *_BillTblMgr) GetBatchFromUserID(userIDs []string) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BillTbl{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromBillID 通过bill_id获取内容 账单id
func (obj *_BillTblMgr) GetFromBillID(billID string) (result BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BillTbl{}).Where("`bill_id` = ?", billID).First(&result).Error

	return
}

// GetBatchFromBillID 批量查找 账单id
func (obj *_BillTblMgr) GetBatchFromBillID(billIDs []string) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BillTbl{}).Where("`bill_id` IN (?)", billIDs).Find(&results).Error

	return
}

// GetFromPayPlatform 通过pay_platform获取内容 支付类型(1:微信支付)
func (obj *_BillTblMgr) GetFromPayPlatform(payPlatform int) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BillTbl{}).Where("`pay_platform` = ?", payPlatform).Find(&results).Error

	return
}

// GetBatchFromPayPlatform 批量查找 支付类型(1:微信支付)
func (obj *_BillTblMgr) GetBatchFromPayPlatform(payPlatforms []int) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BillTbl{}).Where("`pay_platform` IN (?)", payPlatforms).Find(&results).Error

	return
}

// GetFromPayAmount 通过pay_amount获取内容 支付金额
func (obj *_BillTblMgr) GetFromPayAmount(payAmount int64) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BillTbl{}).Where("`pay_amount` = ?", payAmount).Find(&results).Error

	return
}

// GetBatchFromPayAmount 批量查找 支付金额
func (obj *_BillTblMgr) GetBatchFromPayAmount(payAmounts []int64) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BillTbl{}).Where("`pay_amount` IN (?)", payAmounts).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 支付状态(-1:已退款,1:待支付,2:已支付,待发货,3:已取消,4:待收货,5:已完成，6:售后待评价)
func (obj *_BillTblMgr) GetFromStatus(status int) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BillTbl{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 支付状态(-1:已退款,1:待支付,2:已支付,待发货,3:已取消,4:待收货,5:已完成，6:售后待评价)
func (obj *_BillTblMgr) GetBatchFromStatus(statuss []int) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BillTbl{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromDesc 通过desc获取内容 订单备注
func (obj *_BillTblMgr) GetFromDesc(desc string) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BillTbl{}).Where("`desc` = ?", desc).Find(&results).Error

	return
}

// GetBatchFromDesc 批量查找 订单备注
func (obj *_BillTblMgr) GetBatchFromDesc(descs []string) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BillTbl{}).Where("`desc` IN (?)", descs).Find(&results).Error

	return
}

// GetFromCreatedBy 通过created_by获取内容 创建者
func (obj *_BillTblMgr) GetFromCreatedBy(createdBy string) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BillTbl{}).Where("`created_by` = ?", createdBy).Find(&results).Error

	return
}

// GetBatchFromCreatedBy 批量查找 创建者
func (obj *_BillTblMgr) GetBatchFromCreatedBy(createdBys []string) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BillTbl{}).Where("`created_by` IN (?)", createdBys).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_BillTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BillTbl{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_BillTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BillTbl{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedBy 通过updated_by获取内容 更新者
func (obj *_BillTblMgr) GetFromUpdatedBy(updatedBy string) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BillTbl{}).Where("`updated_by` = ?", updatedBy).Find(&results).Error

	return
}

// GetBatchFromUpdatedBy 批量查找 更新者
func (obj *_BillTblMgr) GetBatchFromUpdatedBy(updatedBys []string) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BillTbl{}).Where("`updated_by` IN (?)", updatedBys).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 更新时间
func (obj *_BillTblMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BillTbl{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找 更新时间
func (obj *_BillTblMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BillTbl{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedBy 通过deleted_by获取内容 删除者
func (obj *_BillTblMgr) GetFromDeletedBy(deletedBy string) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BillTbl{}).Where("`deleted_by` = ?", deletedBy).Find(&results).Error

	return
}

// GetBatchFromDeletedBy 批量查找 删除者
func (obj *_BillTblMgr) GetBatchFromDeletedBy(deletedBys []string) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BillTbl{}).Where("`deleted_by` IN (?)", deletedBys).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_BillTblMgr) GetFromDeletedAt(deletedAt time.Time) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BillTbl{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_BillTblMgr) GetBatchFromDeletedAt(deletedAts []time.Time) (results []*BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BillTbl{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_BillTblMgr) FetchByPrimaryKey(id int) (result BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BillTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByBillID primary or index 获取唯一内容
func (obj *_BillTblMgr) FetchUniqueByBillID(billID string) (result BillTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BillTbl{}).Where("`bill_id` = ?", billID).First(&result).Error

	return
}
