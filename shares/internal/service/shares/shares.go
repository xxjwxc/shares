package shares

import (
	"context"
	"fmt"
	"rpc/common"
	"shares/internal/api"
	"shares/internal/config"
	"shares/internal/core"
	"shares/internal/model"
	commonex "shares/internal/service/common"
	"shares/internal/service/event"
	proto "shares/rpc/shares"
	"strings"
	"time"

	"github.com/xxjwxc/public/message"
	"github.com/xxjwxc/public/mycache"
	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/tools"
	"gorm.io/datatypes"
)

// Shares ...
type Shares struct {
}

// GetGroup 获取分组
func (w *Shares) GetGroup(ctx context.Context, req *common.Empty) (*proto.GetGroupResp, error) {
	c, b := ctx.(*api.Context)
	if !b {
		return nil, message.GetError(message.NotVaildError)
	}
	user, err := c.GetUserInfo()
	if err != nil {
		return nil, err
	}

	var groups []string
	if len(user.Info.Group) == 0 { // 默认
		groups = append(groups, config.GetDefGroup())
	} else {
		groups = strings.Split(user.Info.Group, ",")
	}

	resp := &proto.GetGroupResp{}
	for _, v := range groups {
		orm := core.Dao.GetDBr()
		list, err := model.GroupTblMgr(orm.Order("created_at desc")).GetFromGroupName(v)
		if err == nil {
			tmp := &proto.Group{
				Name: v,
			}
			var codes []string
			for _, v1 := range list {
				codes = append(codes, v1.Code)
			}
			out, err := event.GetShares(codes, user.Info.Rg) // 获取股票信息
			if err == nil {
				for _, o := range list {
					for _, v := range out {
						if v.Code == o.Code {
							if len(o.UserName) == 0 {
								v.Attach = fmt.Sprintf("(%v)", tools.FormatTime(o.UpdateAt, "01-02"))
							} else {
								v.Attach = fmt.Sprintf("(%v,%v)", o.UserName, tools.FormatTime(o.UpdateAt, "01-02"))
							}

							tmp.List = append(tmp.List, v)
						}
					}
				}
				// tmp.List = out
			}

			resp.List = append(resp.List, tmp)
		}
	}

	return resp, nil // 获取group list
}

// GetMyGroup 获取我的组织
func (w *Shares) GetMyGroup(ctx context.Context, req *proto.CodeReq) (*proto.GetMyGroupResp, error) {
	c, b := ctx.(*api.Context)
	if !b {
		return nil, message.GetError(message.NotVaildError)
	}
	user, err := c.GetUserInfo()
	if err != nil {
		return nil, err
	}
	if len(user.Info.Group) == 0 {
		return &proto.GetMyGroupResp{}, nil
	}

	groups := strings.Split(user.Info.Group, ",")
	condition := model.Condition{}
	condition.And(model.GroupTblColumns.Code, "=", req.Code)
	condition.And(model.GroupTblColumns.GroupName, "in", groups)
	where, out := condition.Get()
	orm := core.Dao.GetDBr()
	info, _ := model.GroupTblMgr(orm.Where(where, out...).Order("created_at desc")).Get()
	mylog.Info(info)
	if info.ID == 0 {
		return &proto.GetMyGroupResp{
			Group: groups,
		}, nil
	}

	return &proto.GetMyGroupResp{
		Code:     req.Code,
		UserName: info.UserName,
		Wi:       int32(info.Wi),
		Group:    groups,
	}, nil
}

// AddCodeToGroup 添加一个股票到分组池
func (w *Shares) UpsetGroupCode(ctx context.Context, req *proto.UpsetGroupCodeReq) (*common.Empty, error) {
	c, b := ctx.(*api.Context)
	if !b {
		return nil, message.GetError(message.NotVaildError)
	}
	user, err := c.GetUserInfo()
	if err != nil {
		return nil, err
	}
	if len(user.Info.Group) == 0 {
		return nil, message.GetError(message.InValidOp)
	}
	groups := strings.Split(user.Info.Group, ",")
	isFind := false
	for _, v := range groups {
		if req.GroupName == v {
			isFind = true
		}
	}
	if !isFind {
		return nil, message.GetError(message.InValidOp)
	}
	orm := core.Dao.GetDBr()
	info, err := model.GroupTblMgr(orm.DB).FetchUniqueIndexByGroupName(req.GroupName, req.Code)
	if err != nil {
		return nil, err
	}
	if info.ID == 0 {
		info.Code = req.Code
		info.GroupName = req.GroupName
		info.CreatedAt = time.Now()
		info.CreatedBy = user.Info.Openid
		info.UpdateAt = time.Now()
		info.UpdateBy = user.Info.Openid
	} else {
		info.UpdateAt = time.Now()
		info.UpdateBy = user.Info.Openid
	}
	if len(req.UserName) > 0 {
		info.UserName = req.UserName
	}

	if req.IsAdd { // 添加
		err = model.GroupTblMgr(orm.DB).Save(&info).Error
		if err != nil {
			return nil, err
		}
	} else if info.ID > 0 { // 删除
		err = model.GroupTblMgr(orm.DB).Delete(&info).Error
		if err != nil {
			return nil, err
		}
	}

	return &common.Empty{}, nil
}

// Search 搜索
func (w *Shares) Search(ctx context.Context, req *proto.SearchReq) (*proto.SearchResp, error) {
	if len(req.Code) == 0 {
		return nil, message.GetError(message.ParameterInvalid)
	}

	c, b := ctx.(*api.Context)
	if !b {
		return nil, message.GetError(message.NotVaildError)
	}
	user, err := c.GetUserInfo()
	if err != nil {
		return nil, err
	}

	code := req.Code
	ishan := tools.IsHan(req.Code)
	if ishan { // 数据库查找
		condition := model.Condition{}
		condition.And(model.SharesInfoTblColumns.Name, "like", "%"+req.Code+"%")
		where, args := condition.Get()
		orm := core.Dao.GetDBr()
		out, err := model.SharesInfoTblMgr(orm.Where(where, args...)).Get()
		if err != nil {
			return nil, err
		}
		if out.ID == 0 { // 没找到
			msg := message.GetErrorMsg(message.NotFindError, "未查询到,请完整输入股票代码查询")
			return nil, msg.GetError()
		}
		code = out.Code
	}

	if len(code) == 0 {
		msg := message.GetErrorMsg(message.NotFindError, "未查询到,请完整输入股票代码查询")
		return nil, msg.GetError()
	}

	// 开始查找
	out := event.TrySearch(code)
	if out == nil {
		msg := message.GetErrorMsg(message.NotFindError, "未查询到,请完整输入股票代码查询")
		return nil, msg.GetError()
	}

	// 保存数据库
	orm := core.Dao.GetDBr()
	info, err := model.SharesInfoTblMgr(orm.DB).GetFromCode(out.Code)
	if err != nil {
		return nil, err
	}

	if info.ID == 0 { // 保存
		info.Code = out.Code
		info.SimpleCode = out.SimpleCode
		info.Ext = out.Ext
		info.Name = out.Name
		info.SimpleName = event.FirstLetterOfPinYin(out.Name)
		info.Price = float64(out.Price)
		info.Percent = float64(out.Percent)
		info.CreatedBy = user.Info.Openid
		info.CreatedAt = time.Now()
		model.SharesInfoTblMgr(orm.DB).Save(&info)
		// model.SharesListTblMgr(orm.DB).Save(&model.SharesListTbl{
		// 	Code: out.Code,
		// 	Name: out.Name,
		// })
		mycache.NewCache(_ChineseCode).Delete("map") // 清除缓存
		event.BuildOneDaily(out.Code, orm)           // 设置日均线
	}

	if len(req.Tag) == 0 {
		req.Tag = "daily"
	}

	var peg string
	if info.Peg != 0 && len(info.PegShares) > 0 { // peg 估值
		peg = fmt.Sprintf("%0.2f; 行业平均:%0.2f; 行业中值:%0.2f; 相关公司比较:%v", info.Peg, info.PegAvg, info.PegAve, info.PegShares)
	}

	hy, _ := model.SharesInfoTblMgr(orm.DB).GetFromCode(out.Code)

	// 返回信息
	return &proto.SearchResp{
		Info: &proto.SharesInfo{
			Code:       out.Code,
			SimpleCode: out.SimpleCode,
			Ext:        out.Ext,
			Name:       out.Name,
			Price:      out.Price,
			Percent:    out.Percent,
			Color:      event.GetColor(user.Info.Rg, out.Percent),
			Img:        fmt.Sprintf("/webshares/echarts/echarts.html?rg=%v&only20=%v&tag=%v&code=%v", user.Info.Rg, user.Info.Only20, req.Tag, out.Code),
			Hy:         hy.HyName,
			Peg:        peg,
		},
	}, nil // http://image.sinajs.cn/newchart/daily/n/sh000001.gif
}

// Gets 精确查找代码
func (w *Shares) Gets(ctx context.Context, req *proto.GetsReq) (*proto.GetsResp, error) {
	c, b := ctx.(*api.Context)
	if !b {
		return nil, message.GetError(message.NotVaildError)
	}
	user, err := c.GetUserInfo()
	if err != nil {
		return nil, err
	}

	if len(req.Codes) == 0 {
		return nil, message.GetError(message.ParameterInvalid)
	}

	out, err := event.GetShares(req.Codes, user.Info.Rg) // 获取股票信息
	if err != nil {
		return nil, err
	}

	resp := &proto.GetsResp{}
	for _, v := range out {
		resp.List = append(resp.List, &proto.SimpleSharesInfo{
			Code:    v.Code,
			Price:   v.Price,
			Percent: v.Percent,
			Color:   v.Color,
		})
	}
	return resp, nil
}

// GetAllCodeName 获取所有中文
func (w *Shares) GetAllCodeName(ctx context.Context, req *common.Empty) (*proto.GetAllCodeNameResp, error) {
	tmp := []*proto.CodeNameInfo{}
	err := mycache.NewCache(_ChineseCode).Value("map", &tmp)
	if err != nil { // 没找到
		orm := core.Dao.GetDBr()
		out, err := model.SharesInfoTblMgr(orm.DB).Gets()
		if err != nil {
			return nil, err
		}

		tmp = []*proto.CodeNameInfo{}
		for _, v := range out {
			tmp = append(tmp, &proto.CodeNameInfo{
				Code:  v.Code,
				Name:  v.Name,
				SName: v.SimpleName,
			})
		}
		mycache.NewCache(_ChineseCode).Add("map", tmp, 0)
	}

	return &proto.GetAllCodeNameResp{
		List: tmp,
	}, nil
}

// AddMyCode 给自己添加一个监听
func (w *Shares) AddMyCode(ctx context.Context, req *proto.AddMyCodeReq) (*proto.AddMyCodeResp, error) {
	c, b := ctx.(*api.Context)
	if !b {
		return nil, message.GetError(message.NotVaildError)
	}
	user, err := c.GetUserInfo()
	if err != nil {
		return nil, err
	}

	orm := core.Dao.GetDBw()
	find := false
	// 添加到数据库(先查询是否有)
	info, err := model.SharesWatchTblMgr(orm.DB).FetchUniqueIndexByOpenCode(req.Code, user.Info.Openid)
	if err != nil {
		return nil, err
	}
	if info.ID > 0 {
		find = true
		// return nil, message.GetError(message.ExistedError)
	}

	var out model.WxUserinfo
	if !find {
		out, err = model.WxUserinfoMgr(orm.DB).GetFromOpenid(user.Info.Openid)
		if err != nil {
			return nil, err
		}
		if out.ID == 0 {
			return nil, message.GetError(message.NotFindError)
		}

		if out.Capacity <= 0 { // 数量不够
			return &proto.AddMyCodeResp{
				Status: false,
				Msg:    "可用容量不够,请充值或联系管理员",
			}, nil
		}
	}

	// 开始保存
	info.Code = req.Code
	info.OpenID = user.Info.Openid
	info.Kdj = req.Kdj
	info.Kdj20 = req.Kdj20
	info.Surge = req.Surge
	info.Slump = req.Slump
	info.Ai = req.Ai
	info.GroupWi = req.Public
	info.Up = req.Up
	info.Down = req.Down
	info.UpPercent = req.UpPercent
	info.DownPercent = req.DownPercent
	info.CreatedAt = time.Now()
	info.ExpiresTime = info.CreatedAt.AddDate(1, 0, 0)
	info.Vaild = req.Vaild // 默认开启
	if info.Up > 0 {
		info.UpVaild = true
	}
	if info.Down > 0 {
		info.DownVaild = true
	}
	if info.UpPercent > 0 {
		info.UpPercentTo = req.Price + ((req.Price * info.UpPercent) * 0.01)
	}

	if info.DownPercent > 0 {
		info.DownPercentTo = req.Price - ((req.Price * info.UpPercent) * 0.01)
	}

	orm.Begin()
	tx := orm.Begin()
	defer func() {
		if err != nil {
			tx.AddError(err)
		}
		orm.Commit(tx)
	}()
	err = model.SharesWatchTblMgr(tx).Save(&info).Error
	if err != nil {
		return nil, err
	}

	api.DeleteCacheBody(user.Info.Openid) // 删除cache

	if !find {
		out.Capacity -= 1
		err = model.WxUserinfoMgr(orm.Where("id = ?", out.ID)).Update(model.WxUserinfoColumns.Capacity, out.Capacity).Error
		if err != nil {
			return nil, err
		}
	}

	if req.Public { // 推荐给组织
		UpDataGroup(strings.Split(user.Info.Group, ","), req.Code, user.Info.NickName)
	}

	return &proto.AddMyCodeResp{Status: true}, nil
}

// AddMyCode 给自己添加一个监听
func (w *Shares) GetMyCode(ctx context.Context, req *proto.GetMyCodeReq) (*proto.GetMyCodeResp, error) {
	c, b := ctx.(*api.Context)
	if !b {
		return nil, message.GetError(message.NotVaildError)
	}
	user, err := c.GetUserInfo()
	if err != nil {
		return nil, err
	}

	var list []*model.SharesWatchTbl
	orm := core.Dao.GetDBr()
	if len(req.Code) > 0 {
		info, err := model.SharesWatchTblMgr(orm.DB).FetchUniqueIndexByOpenCode(req.Code, user.Info.Openid)
		if err != nil {
			return nil, err
		}
		if info.ID > 0 {
			list = append(list, &info)
		}
	} else {
		list, err = model.SharesWatchTblMgr(orm.Order("id desc")).GetFromOpenID(user.Info.Openid)
		if err != nil {
			return nil, err
		}
	}
	var codes []string
	for _, v := range list {
		codes = append(codes, v.Code)
	}
	out, err := event.GetShares(codes, user.Info.Rg) // 获取股票信息
	if err != nil {
		return nil, err
	}
	_map := make(map[string]*proto.SharesInfo)
	for _, v := range out {
		_map[v.Code] = v
	}

	// output
	resp := &proto.GetMyCodeResp{}
	for _, v := range list {
		tmp := &proto.AddMyCodeReq{
			Code:        v.Code,
			Kdj:         v.Kdj,
			Kdj20:       v.Kdj20,
			Surge:       v.Surge,
			Slump:       v.Slump,
			Ai:          v.Ai,
			Public:      v.GroupWi,
			Up:          v.Up,
			Down:        v.Down,
			UpPercent:   v.UpPercent,
			DownPercent: v.DownPercent,
			Vaild:       v.Vaild,
		}
		if v, ok := _map[v.Code]; ok {
			tmp.Price = v.Price
			tmp.SimpleCode = v.SimpleCode
			tmp.Ext = v.Ext
			tmp.Name = v.Name
			tmp.Percent = v.Percent
			tmp.Color = v.Color
		}

		resp.List = append(resp.List, tmp)
	}

	return resp, nil
}

// DeleteMyCode 删除一个监听
func (w *Shares) DeleteMyCode(ctx context.Context, req *proto.DeleteMyCodeReq) (*common.Empty, error) {
	c, b := ctx.(*api.Context)
	if !b {
		return nil, message.GetError(message.NotVaildError)
	}
	user, err := c.GetUserInfo()
	if err != nil {
		return nil, err
	}
	if len(req.Code) == 0 {
		return nil, message.GetError(message.ParameterInvalid)
	}

	orm := core.Dao.GetDBr()

	info, err := model.SharesWatchTblMgr(orm.DB).FetchUniqueIndexByOpenCode(req.Code, user.Info.Openid)
	if err != nil {
		return nil, err
	}
	if info.ID == 0 {
		return nil, message.GetError(message.NotFindError)
	}
	err = model.SharesWatchTblMgr(orm.DB).Delete(&info).Error
	if err != nil {
		return nil, err
	}

	return &common.Empty{}, nil
}

// GetMsg 获取消息
func (w *Shares) GetMsg(ctx context.Context, req *common.Empty) (*proto.GetMsgResp, error) {
	c, b := ctx.(*api.Context)
	if !b {
		return nil, message.GetError(message.NotVaildError)
	}
	user, err := c.GetUserInfo()
	if err != nil {
		return nil, err
	}
	orm := core.Dao.GetDBr()
	list, err := model.MsgTblMgr(orm.Order("id desc").Limit(10)).GetFromOpenID(user.Info.Openid)
	if err != nil {
		return nil, err
	}

	resp := &proto.GetMsgResp{}
	for _, v := range list {
		info, err := event.GetShare(v.Code, user.Info.Rg)
		if err != nil {
			return nil, err
		}
		tag := "daily"
		if v.Key == "急涨提醒" || v.Key == "急跌提醒" {
			tag = "min"
		}
		resp.List = append(resp.List, &proto.Msg{
			Code:       v.Code,
			SimpleCode: info.SimpleCode,
			Ext:        info.Ext,
			Name:       info.Name,
			Price:      v.Price,
			Key:        v.Key,
			Desc:       v.Desc,
			Percent:    v.Percent,
			Color:      event.GetColor(user.Info.Rg, v.Percent),
			Now:        tools.GetTimeStr(v.CreatedAt),
			Tag:        tag,
		})
	}

	return resp, nil
}

// AddMyCode 是否有新消息
func (w *Shares) HaveNewMsg(ctx context.Context, req *common.Empty) (*proto.HaveNewMsgResp, error) {
	c, b := ctx.(*api.Context)
	if !b {
		return nil, message.GetError(message.NotVaildError)
	}
	user, err := c.GetUserInfo()
	if err != nil {
		return nil, err
	}
	var count int64
	orm := core.Dao.GetDBr()
	err = model.MsgTblMgr(orm.Where("day = ? and open_id = ?", datatypes.Date(time.Now()), user.Info.Openid).Order("id desc").Limit(10)).Count(&count).Error
	if err != nil {
		return nil, err
	}

	return &proto.HaveNewMsgResp{
		Badge: count > 0,
	}, nil
}

// AddGroup 添加一个组织
func (w *Shares) AddGroup(ctx context.Context, req *proto.AddGroupReq) (*common.Empty, error) {
	c, b := ctx.(*api.Context)
	if !b {
		return nil, message.GetError(message.NotVaildError)
	}
	user, err := c.GetUserInfo()
	if err != nil {
		return nil, err
	}

	_, err = commonex.AddGroup(user.Info.Openid, req.Key)
	if err != nil {
		return nil, err
	}
	// orm := core.Dao.GetDBw()
	// out, err := model.GroupListTblMgr(orm.DB).GetFromKey(req.Key)
	// if err != nil {
	// 	return nil, err
	// }
	// if out.ID == 0 {
	// 	return nil, message.GetError(message.NotFindError)
	// }
	// list := strings.Split(user.Info.Group, ",")
	// for _, v := range list {
	// 	if strings.EqualFold(v, out.GroupName) { // 已经存在
	// 		return &common.Empty{}, nil
	// 	}
	// }
	// list = append(list, out.GroupName)
	// tmp := list[:0]
	// for _, v := range list {
	// 	if len(v) > 0 {
	// 		tmp = append(tmp, v)
	// 	}
	// }

	// model.WxUserinfoMgr(orm.Where("openid=?", user.Info.Openid)).Update("group", strings.Join(tmp, ","))
	// api.DeleteCacheBody(user.Info.Openid) // 清除缓存
	return &common.Empty{}, nil
}

// Minute 获取分时图
func (w *Shares) Minute(ctx *api.Context, req *proto.CodeReq) (*MinuteOut, error) {
	// 开始获取
	out, err := event.GetMinute(req.Code)
	if err != nil {
		return nil, err
	}
	resp := &MinuteOut{
		Yestclose: out.PrePrice,
	}

	for _, v := range out.List {
		resp.Data = append(resp.Data, []interface{}{v.Time, v.Price, v.Ave, v.Vol})
	}

	resp.Ref = true
	now := time.Now()
	day0 := tools.GetUtcDay0(now)
	offset := now.Unix() - day0
	week := tools.GetTimeWeek(day0)
	if week == 6 || week == 0 {
		resp.Ref = false
	}

	if resp.Ref && (offset < (9*60*60) || offset > (15*60*60+30*60)) { // 3点之后是空闲时间
		resp.Ref = false
	}

	resp.Ref = true

	return resp, nil
}

// Minute 获取日k图
func (w *Shares) Dayliy(ctx *api.Context, req *proto.CodeReq) ([][]interface{}, error) {
	// 开始获取
	out, err := event.GetDayliy(req.Code)
	if err != nil {
		return nil, err
	}
	return out, nil
}
