package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _WxMsgTblMgr struct {
	*_BaseMgr
}

// WxMsgTblMgr open func
func WxMsgTblMgr(db *gorm.DB) *_WxMsgTblMgr {
	if db == nil {
		panic(fmt.Errorf("WxMsgTblMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WxMsgTblMgr{_BaseMgr: &_BaseMgr{DB: db.Table("wx_msg_tbl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WxMsgTblMgr) GetTableName() string {
	return "wx_msg_tbl"
}

// Reset 重置gorm会话
func (obj *_WxMsgTblMgr) Reset() *_WxMsgTblMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_WxMsgTblMgr) Get() (result WxMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxMsgTbl{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WxMsgTblMgr) Gets() (results []*WxMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxMsgTbl{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_WxMsgTblMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(WxMsgTbl{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_WxMsgTblMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithToUserName to_user_name获取 接收方帐号（收到的OpenID）
func (obj *_WxMsgTblMgr) WithToUserName(toUserName string) Option {
	return optionFunc(func(o *options) { o.query["to_user_name"] = toUserName })
}

// WithFromUserName from_user_name获取 开发者微信号
func (obj *_WxMsgTblMgr) WithFromUserName(fromUserName string) Option {
	return optionFunc(func(o *options) { o.query["from_user_name"] = fromUserName })
}

// WithCreatetime createTime获取 消息创建时间 （整型）
func (obj *_WxMsgTblMgr) WithCreatetime(createtime int64) Option {
	return optionFunc(func(o *options) { o.query["createTime"] = createtime })
}

// WithMsgType msg_type获取 消息类型，文本为text
func (obj *_WxMsgTblMgr) WithMsgType(msgType string) Option {
	return optionFunc(func(o *options) { o.query["msg_type"] = msgType })
}

// WithMsgID msg_id获取 消息id
func (obj *_WxMsgTblMgr) WithMsgID(msgID string) Option {
	return optionFunc(func(o *options) { o.query["msg_id"] = msgID })
}

// WithContent content获取 回复的消息内容（换行：在content中能够换行，微信客户端就支持换行显示）
func (obj *_WxMsgTblMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_WxMsgTblMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithEvent event获取 事件内容
func (obj *_WxMsgTblMgr) WithEvent(event string) Option {
	return optionFunc(func(o *options) { o.query["event"] = event })
}

// WithMediaID media_id获取 语音消息媒体id，可以调用获取临时素材接口拉取数据。
func (obj *_WxMsgTblMgr) WithMediaID(mediaID string) Option {
	return optionFunc(func(o *options) { o.query["media_id"] = mediaID })
}

// WithFormat format获取 语音格式，如amr，speex等
func (obj *_WxMsgTblMgr) WithFormat(format string) Option {
	return optionFunc(func(o *options) { o.query["format"] = format })
}

// WithReq req获取 请求参数
func (obj *_WxMsgTblMgr) WithReq(req string) Option {
	return optionFunc(func(o *options) { o.query["req"] = req })
}

// WithStatus status获取 当前状态(0:已创建,1:已回复,-1:连接已断开)
func (obj *_WxMsgTblMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// GetByOption 功能选项模式获取
func (obj *_WxMsgTblMgr) GetByOption(opts ...Option) (result WxMsgTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WxMsgTbl{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WxMsgTblMgr) GetByOptions(opts ...Option) (results []*WxMsgTbl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WxMsgTbl{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_WxMsgTblMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]WxMsgTbl, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(WxMsgTbl{}).Where(options.query)
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
func (obj *_WxMsgTblMgr) GetFromID(id int) (result WxMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxMsgTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_WxMsgTblMgr) GetBatchFromID(ids []int) (results []*WxMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxMsgTbl{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromToUserName 通过to_user_name获取内容 接收方帐号（收到的OpenID）
func (obj *_WxMsgTblMgr) GetFromToUserName(toUserName string) (results []*WxMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxMsgTbl{}).Where("`to_user_name` = ?", toUserName).Find(&results).Error

	return
}

// GetBatchFromToUserName 批量查找 接收方帐号（收到的OpenID）
func (obj *_WxMsgTblMgr) GetBatchFromToUserName(toUserNames []string) (results []*WxMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxMsgTbl{}).Where("`to_user_name` IN (?)", toUserNames).Find(&results).Error

	return
}

// GetFromFromUserName 通过from_user_name获取内容 开发者微信号
func (obj *_WxMsgTblMgr) GetFromFromUserName(fromUserName string) (results []*WxMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxMsgTbl{}).Where("`from_user_name` = ?", fromUserName).Find(&results).Error

	return
}

// GetBatchFromFromUserName 批量查找 开发者微信号
func (obj *_WxMsgTblMgr) GetBatchFromFromUserName(fromUserNames []string) (results []*WxMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxMsgTbl{}).Where("`from_user_name` IN (?)", fromUserNames).Find(&results).Error

	return
}

// GetFromCreatetime 通过createTime获取内容 消息创建时间 （整型）
func (obj *_WxMsgTblMgr) GetFromCreatetime(createtime int64) (results []*WxMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxMsgTbl{}).Where("`createTime` = ?", createtime).Find(&results).Error

	return
}

// GetBatchFromCreatetime 批量查找 消息创建时间 （整型）
func (obj *_WxMsgTblMgr) GetBatchFromCreatetime(createtimes []int64) (results []*WxMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxMsgTbl{}).Where("`createTime` IN (?)", createtimes).Find(&results).Error

	return
}

// GetFromMsgType 通过msg_type获取内容 消息类型，文本为text
func (obj *_WxMsgTblMgr) GetFromMsgType(msgType string) (results []*WxMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxMsgTbl{}).Where("`msg_type` = ?", msgType).Find(&results).Error

	return
}

// GetBatchFromMsgType 批量查找 消息类型，文本为text
func (obj *_WxMsgTblMgr) GetBatchFromMsgType(msgTypes []string) (results []*WxMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxMsgTbl{}).Where("`msg_type` IN (?)", msgTypes).Find(&results).Error

	return
}

// GetFromMsgID 通过msg_id获取内容 消息id
func (obj *_WxMsgTblMgr) GetFromMsgID(msgID string) (result WxMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxMsgTbl{}).Where("`msg_id` = ?", msgID).First(&result).Error

	return
}

// GetBatchFromMsgID 批量查找 消息id
func (obj *_WxMsgTblMgr) GetBatchFromMsgID(msgIDs []string) (results []*WxMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxMsgTbl{}).Where("`msg_id` IN (?)", msgIDs).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容 回复的消息内容（换行：在content中能够换行，微信客户端就支持换行显示）
func (obj *_WxMsgTblMgr) GetFromContent(content string) (results []*WxMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxMsgTbl{}).Where("`content` = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量查找 回复的消息内容（换行：在content中能够换行，微信客户端就支持换行显示）
func (obj *_WxMsgTblMgr) GetBatchFromContent(contents []string) (results []*WxMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxMsgTbl{}).Where("`content` IN (?)", contents).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_WxMsgTblMgr) GetFromCreatedAt(createdAt time.Time) (results []*WxMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxMsgTbl{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_WxMsgTblMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*WxMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxMsgTbl{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromEvent 通过event获取内容 事件内容
func (obj *_WxMsgTblMgr) GetFromEvent(event string) (results []*WxMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxMsgTbl{}).Where("`event` = ?", event).Find(&results).Error

	return
}

// GetBatchFromEvent 批量查找 事件内容
func (obj *_WxMsgTblMgr) GetBatchFromEvent(events []string) (results []*WxMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxMsgTbl{}).Where("`event` IN (?)", events).Find(&results).Error

	return
}

// GetFromMediaID 通过media_id获取内容 语音消息媒体id，可以调用获取临时素材接口拉取数据。
func (obj *_WxMsgTblMgr) GetFromMediaID(mediaID string) (results []*WxMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxMsgTbl{}).Where("`media_id` = ?", mediaID).Find(&results).Error

	return
}

// GetBatchFromMediaID 批量查找 语音消息媒体id，可以调用获取临时素材接口拉取数据。
func (obj *_WxMsgTblMgr) GetBatchFromMediaID(mediaIDs []string) (results []*WxMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxMsgTbl{}).Where("`media_id` IN (?)", mediaIDs).Find(&results).Error

	return
}

// GetFromFormat 通过format获取内容 语音格式，如amr，speex等
func (obj *_WxMsgTblMgr) GetFromFormat(format string) (results []*WxMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxMsgTbl{}).Where("`format` = ?", format).Find(&results).Error

	return
}

// GetBatchFromFormat 批量查找 语音格式，如amr，speex等
func (obj *_WxMsgTblMgr) GetBatchFromFormat(formats []string) (results []*WxMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxMsgTbl{}).Where("`format` IN (?)", formats).Find(&results).Error

	return
}

// GetFromReq 通过req获取内容 请求参数
func (obj *_WxMsgTblMgr) GetFromReq(req string) (results []*WxMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxMsgTbl{}).Where("`req` = ?", req).Find(&results).Error

	return
}

// GetBatchFromReq 批量查找 请求参数
func (obj *_WxMsgTblMgr) GetBatchFromReq(reqs []string) (results []*WxMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxMsgTbl{}).Where("`req` IN (?)", reqs).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 当前状态(0:已创建,1:已回复,-1:连接已断开)
func (obj *_WxMsgTblMgr) GetFromStatus(status int) (results []*WxMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxMsgTbl{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 当前状态(0:已创建,1:已回复,-1:连接已断开)
func (obj *_WxMsgTblMgr) GetBatchFromStatus(statuss []int) (results []*WxMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxMsgTbl{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_WxMsgTblMgr) FetchByPrimaryKey(id int) (result WxMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxMsgTbl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByMsgID primary or index 获取唯一内容
func (obj *_WxMsgTblMgr) FetchUniqueByMsgID(msgID string) (result WxMsgTbl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WxMsgTbl{}).Where("`msg_id` = ?", msgID).First(&result).Error

	return
}
