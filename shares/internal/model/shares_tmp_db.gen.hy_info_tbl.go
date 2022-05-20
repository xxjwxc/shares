package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _HyInfoTblMgr struct {
	*_BaseMgr
}

// HyInfoTblMgr open func
func HyInfoTblMgr(db *gorm.DB) *_HyInfoTblMgr {
	if db == nil {
		panic(fmt.Errorf("HyInfoTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_HyInfoTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("hy_info_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_HyInfoTblMgr) GetTableName() string {
	return "hy_info_tbl"
}

// Reset 重置gorm会话
func (obj *_HyInfoTblMgr) Reset() *_HyInfoTblMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_HyInfoTblMgr) Get() (result HyInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyInfoTbl{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_HyInfoTblMgr) Gets() (results []*HyInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyInfoTbl{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_HyInfoTblMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(HyInfoTbl{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_HyInfoTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithHyCode hy_code获取 行业代码代码
func (obj *_HyInfoTblMgr) WithHyCode(hyCode string) Option {
	return optionFunc(func(o *options) { o.query["hy_code"] = hyCode })
}

// WithName name获取 行业名
func (obj *_HyInfoTblMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithNum num获取 家数
func (obj *_HyInfoTblMgr) WithNum(num int) Option {
	return optionFunc(func(o *options) { o.query["num"] = num })
}

// WithPe pe获取 市盈率
func (obj *_HyInfoTblMgr) WithPe(pe float64) Option {
	return optionFunc(func(o *options) { o.query["pe"] = pe })
}

// WithPeg peg获取 PEG值
func (obj *_HyInfoTblMgr) WithPeg(peg float64) Option {
	return optionFunc(func(o *options) { o.query["peg"] = peg })
}

// WithPb pb获取 市净率
func (obj *_HyInfoTblMgr) WithPb(pb float64) Option {
	return optionFunc(func(o *options) { o.query["pb"] = pb })
}

// WithCapVag cap_vag获取 平均市值
func (obj *_HyInfoTblMgr) WithCapVag(capVag float64) Option {
	return optionFunc(func(o *options) { o.query["cap_vag"] = capVag })
}

// WithLossNum loss_num获取 亏损家数
func (obj *_HyInfoTblMgr) WithLossNum(lossNum int) Option {
	return optionFunc(func(o *options) { o.query["loss_num"] = lossNum })
}

// WithUpdateAt update_at获取 创建时间
func (obj *_HyInfoTblMgr) WithUpdateAt(updateAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_HyInfoTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// GetByOption 功能选项模式获取
func (obj *_HyInfoTblMgr) GetByOption(opts ...Option) (result HyInfoTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(HyInfoTbl{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_HyInfoTblMgr) GetByOptions(opts ...Option) (results []*HyInfoTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(HyInfoTbl{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_HyInfoTblMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]HyInfoTbl, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(HyInfoTbl{}).Where(options.query)
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
func (obj *_HyInfoTblMgr) GetFromID(id int) (result HyInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyInfoTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_HyInfoTblMgr) GetBatchFromID(ids []int) (results []*HyInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyInfoTbl{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromHyCode 通过hy_code获取内容 行业代码代码
func (obj *_HyInfoTblMgr) GetFromHyCode(hyCode string) (result HyInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyInfoTbl{}).Where("`hy_code` = ?", hyCode).First(&result).Error

	return
}

// GetBatchFromHyCode 批量查找 行业代码代码
func (obj *_HyInfoTblMgr) GetBatchFromHyCode(hyCodes []string) (results []*HyInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyInfoTbl{}).Where("`hy_code` IN (?)", hyCodes).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 行业名
func (obj *_HyInfoTblMgr) GetFromName(name string) (result HyInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyInfoTbl{}).Where("`name` = ?", name).First(&result).Error

	return
}

// GetBatchFromName 批量查找 行业名
func (obj *_HyInfoTblMgr) GetBatchFromName(names []string) (results []*HyInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyInfoTbl{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromNum 通过num获取内容 家数
func (obj *_HyInfoTblMgr) GetFromNum(num int) (results []*HyInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyInfoTbl{}).Where("`num` = ?", num).Find(&results).Error

	return
}

// GetBatchFromNum 批量查找 家数
func (obj *_HyInfoTblMgr) GetBatchFromNum(nums []int) (results []*HyInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyInfoTbl{}).Where("`num` IN (?)", nums).Find(&results).Error

	return
}

// GetFromPe 通过pe获取内容 市盈率
func (obj *_HyInfoTblMgr) GetFromPe(pe float64) (results []*HyInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyInfoTbl{}).Where("`pe` = ?", pe).Find(&results).Error

	return
}

// GetBatchFromPe 批量查找 市盈率
func (obj *_HyInfoTblMgr) GetBatchFromPe(pes []float64) (results []*HyInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyInfoTbl{}).Where("`pe` IN (?)", pes).Find(&results).Error

	return
}

// GetFromPeg 通过peg获取内容 PEG值
func (obj *_HyInfoTblMgr) GetFromPeg(peg float64) (results []*HyInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyInfoTbl{}).Where("`peg` = ?", peg).Find(&results).Error

	return
}

// GetBatchFromPeg 批量查找 PEG值
func (obj *_HyInfoTblMgr) GetBatchFromPeg(pegs []float64) (results []*HyInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyInfoTbl{}).Where("`peg` IN (?)", pegs).Find(&results).Error

	return
}

// GetFromPb 通过pb获取内容 市净率
func (obj *_HyInfoTblMgr) GetFromPb(pb float64) (results []*HyInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyInfoTbl{}).Where("`pb` = ?", pb).Find(&results).Error

	return
}

// GetBatchFromPb 批量查找 市净率
func (obj *_HyInfoTblMgr) GetBatchFromPb(pbs []float64) (results []*HyInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyInfoTbl{}).Where("`pb` IN (?)", pbs).Find(&results).Error

	return
}

// GetFromCapVag 通过cap_vag获取内容 平均市值
func (obj *_HyInfoTblMgr) GetFromCapVag(capVag float64) (results []*HyInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyInfoTbl{}).Where("`cap_vag` = ?", capVag).Find(&results).Error

	return
}

// GetBatchFromCapVag 批量查找 平均市值
func (obj *_HyInfoTblMgr) GetBatchFromCapVag(capVags []float64) (results []*HyInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyInfoTbl{}).Where("`cap_vag` IN (?)", capVags).Find(&results).Error

	return
}

// GetFromLossNum 通过loss_num获取内容 亏损家数
func (obj *_HyInfoTblMgr) GetFromLossNum(lossNum int) (results []*HyInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyInfoTbl{}).Where("`loss_num` = ?", lossNum).Find(&results).Error

	return
}

// GetBatchFromLossNum 批量查找 亏损家数
func (obj *_HyInfoTblMgr) GetBatchFromLossNum(lossNums []int) (results []*HyInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyInfoTbl{}).Where("`loss_num` IN (?)", lossNums).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容 创建时间
func (obj *_HyInfoTblMgr) GetFromUpdateAt(updateAt time.Time) (results []*HyInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyInfoTbl{}).Where("`update_at` = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量查找 创建时间
func (obj *_HyInfoTblMgr) GetBatchFromUpdateAt(updateAts []time.Time) (results []*HyInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyInfoTbl{}).Where("`update_at` IN (?)", updateAts).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_HyInfoTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*HyInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyInfoTbl{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_HyInfoTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*HyInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyInfoTbl{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_HyInfoTblMgr) FetchByPrimaryKey(id int) (result HyInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyInfoTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByHyCode primary or index 获取唯一内容
func (obj *_HyInfoTblMgr) FetchUniqueByHyCode(hyCode string) (result HyInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyInfoTbl{}).Where("`hy_code` = ?", hyCode).First(&result).Error

	return
}

// FetchUniqueByName primary or index 获取唯一内容
func (obj *_HyInfoTblMgr) FetchUniqueByName(name string) (result HyInfoTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HyInfoTbl{}).Where("`name` = ?", name).First(&result).Error

	return
}
