package weixin

import (
	"context"
	"errors"
	"fmt"
	"path"
	"shares/internal/api"
	"shares/internal/config"
	"shares/internal/core"
	"shares/internal/model"
	"strings"
	"time"

	"rpc/common"
	proto "shares/rpc/shares"

	"github.com/xxjwxc/public/message"
	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/tools"
)

// Weixin 微信相关操作
type Weixin struct {
}

// Oauth 微信授权获取登录信息
func (w *Weixin) Oauth(ctx context.Context, req *proto.OauthReq) (*proto.OauthResp, error) {
	c, b := ctx.(*api.Context)
	if !b {
		return nil, message.GetError(message.NotVaildError)
	}

	out, err := _wx.GetWebOauth(req.Jscode)
	if err != nil {
		return nil, err
	}
	if len(out.Openid) == 0 {
		user, err := c.GetUserInfo()
		if err != nil {
			return nil, errors.New(message.InValidOp.String())
		}
		out.Openid = user.Info.Openid
	}

	var result proto.OauthResp
	result.SessionId, result.OpenId, result.OverdueTime = SaveSessionKey([]byte(tools.JSONDecode(out)))

	if len(result.SessionId) == 0 { // fail 失败
		return nil, errors.New(message.InValidOp.String())
	}

	if req.IsUpdateUser && len(out.AccessToken) > 0 {
		userinfo, _ := _wx.GetWebUserinfo(out.Openid, out.AccessToken)
		if userinfo != nil && len(userinfo.Openid) > 0 {
			info := model.WxUserinfo{
				City:      userinfo.City,
				Country:   userinfo.Country,
				AvatarURL: userinfo.Headimgurl,
				NickName:  tools.UnicodeEmojiCode(userinfo.NickName),
				Province:  userinfo.Province,
				Gender:    "女",
				Openid:    userinfo.Openid,
				Phone:     "",
				Group:     "",
				Rg:        true,
				Capacity:  config.GetMaxCapacity(),
			}
			if userinfo.Sex == 1 {
				info.Gender = "男"
			}
			updateInfo(info)
		}
	}

	//设置客户端cookie
	tm := int(result.OverdueTime)
	c.GetGinCtx().SetCookie(api.UserToken, result.OpenId, tm, "/", c.GetGinCtx().Request.Host, false, true)
	c.GetGinCtx().SetCookie(api.SessionToken, result.SessionId, tm, "/", c.GetGinCtx().Request.Host, false, true)

	return &result, nil
}

// ReLogin 是否要重新登录
func (w *Weixin) ReLogin(ctx context.Context, req *common.Empty) (*proto.ReLoginResp, error) {
	c, b := ctx.(*api.Context)
	if !b {
		return nil, message.GetError(message.NotVaildError)
	}
	user, err := c.GetUserInfo()
	if err != nil { // 需要重新登录
		return &proto.ReLoginResp{
			Relogin: true,
		}, nil
	}

	result := &proto.ReLoginResp{Relogin: false, OpenId: user.Info.Openid}
	result.SessionId, result.OpenId, result.OverdueTime = SaveSessionKey([]byte(tools.JSONDecode(WxSessionKey{Openid: user.Info.Openid})))
	//设置客户端cookie
	tm := int(result.OverdueTime)
	c.GetGinCtx().SetCookie(api.UserToken, result.OpenId, tm, "/", c.GetGinCtx().Request.Host, false, true)
	c.GetGinCtx().SetCookie(api.SessionToken, result.SessionId, tm, "/", c.GetGinCtx().Request.Host, false, true)
	return result, nil // 不需要重新登录
}

// UpdateUserInfo 	更新用户信息
func (w *Weixin) UpdateUserInfo(ctx context.Context, req *proto.WxUserinfo) (*common.Empty, error) {
	c, b := ctx.(*api.Context)
	if !b {
		return nil, message.GetError(message.NotVaildError)
	}

	sessionID := c.GetSessionToken()
	mylog.Infof("sessionid:%v", sessionID)
	session := GetSessionkey(sessionID)
	if len(session.Openid) == 0 {
		return nil, message.GetError(message.UserNameDoNotExist)
	}

	// info := model.WxUserinfo{
	// 	City:      req.City,
	// 	Country:   req.Country,
	// 	AvatarURL: req.AvatarURL,
	// 	NickName:  tools.UnicodeEmojiCode(req.NickName),
	// 	Province:  req.Province,
	// 	Gender:    req.Gender,
	// 	Language:  []byte(req.Language),
	// 	Openid:    session.Openid,
	// }

	// if len(info.AvatarURL) > 0 || len(info.NickName) > 0 {
	// 	updateInfo(info)
	// }
	return &common.Empty{}, nil
}

// GetQrcode 获取微信信息
func (w *Weixin) GetUserInfo(ctx context.Context, req *common.Empty) (resp *proto.GetUserInfoResp, err error) {
	c, b := ctx.(*api.Context)
	if !b {
		return nil, message.GetError(message.NotVaildError)
	}
	user, err := c.GetUserInfo()
	if err != nil {
		return nil, err
	}

	orm := core.Dao.GetDBr()
	out, err := model.WxUserinfoMgr(orm.DB).GetFromOpenid(user.Info.Openid)
	if err != nil {
		return nil, err
	}

	return &proto.GetUserInfoResp{
		Openid:    out.Openid,
		NickName:  tools.UnicodeEmojiDecode(out.NickName),
		AvatarURL: out.AvatarURL,
		Gender:    out.Gender,
		City:      out.City,
		Province:  out.Province,
		Country:   out.Country,
		Phone:     out.Phone,
		Group:     out.Group,
		Rg:        out.Rg,
		Only20:    out.Only20,
		Capacity:  int32(out.Capacity),
	}, nil
}

// UpsetUserInfo 更新用户信息
func (w *Weixin) UpsetUserInfo(ctx context.Context, req *proto.UpsetUserInfoReq) (*common.Empty, error) {
	c, b := ctx.(*api.Context)
	if !b {
		return nil, message.GetError(message.NotVaildError)
	}
	user, err := c.GetUserInfo()
	if err != nil {
		return nil, err
	}

	orm := core.Dao.GetDBr()
	out, err := model.WxUserinfoMgr(orm.DB).GetFromOpenid(user.Info.Openid)
	if err != nil {
		return nil, err
	}
	if out.ID == 0 {
		return nil, message.GetError(message.NotFindError)
	}

	switch req.Key {
	case "phone":
		out.Phone = req.Value
	case "rg":
		out.Rg = (req.Value == "true")
	case "only20":
		out.Only20 = (req.Value == "true")
	}
	err = model.WxUserinfoMgr(orm.DB).Save(&out).Error
	if err != nil {
		return nil, err
	}

	api.DeleteCacheBody(user.Info.Openid)

	return &common.Empty{}, nil
}

// GetQrcode 获取微信二维码
func (w *Weixin) GetQrcode(ctx context.Context, req *proto.GetQrcodeReq) (resp *proto.GetQrcodeResp, _err error) {
	page := fmt.Sprintf("/file/qrcode/%v_%v.jpg", strings.Replace(req.Page, "/", "_", -1), req.Scene)
	path := path.Join(tools.GetCurrentDirectory(), page)
	if !tools.CheckFileIsExist(path) {
		ret := _wx.GetShareQrcode(path, req.Scene, req.Page)
		mylog.Info(ret)
		if ret.Errcode != 0 {
			return nil, fmt.Errorf(ret.Errmsg)
		}
	}

	return &proto.GetQrcodeResp{
		Url: config.GetFileHost() + page,
	}, nil
}

// PayPlaceOrder 支付
func PayPlaceOrder(openID string, price int64, priceBody, orderID, clientIP string) (map[string]string, error) {
	msg := _wx.SmallAppUnifiedorder(openID, price, priceBody, orderID, clientIP)
	mylog.Info(msg)

	if msg.State {
		return msg.Data.(map[string]string), nil
	}

	return nil, fmt.Errorf(msg.Error)
}

// SelectOrder 支付查询接口
func SelectOrder(openID, orderID string) (int, message.MessageBody) {
	return _wx.SelectOrder(openID, orderID)
}

// RefundPay 退款
func (w *Weixin) RefundPay(billId string) error { // 退款
	orm := core.Dao.GetDBw()
	billInfo, err := model.BillTblMgr(orm.DB).GetFromBillID(billId)
	if err != nil {
		if orm.IsNotFound(err) {
			return message.GetError(message.HaveDeal)
		}
		return err
	}
	var refundBill model.BillRefundTbl
	refundBill.BillID = billInfo.BillID
	refundBill.RefundFee = billInfo.PayAmount
	refundBill.RefundID = createUnique("")
	refundBill.CreatedAt = time.Now()

	b, msg := _wx.RefundPay(billInfo.UserID, billInfo.BillID, refundBill.RefundID, refundBill.RefundFee, billInfo.PayAmount)
	mylog.Info(msg) // 退款信息
	if b {
		err := model.BillRefundTblMgr(orm.DB).Save(&refundBill).Error
		if err != nil {
			return err
		}
		// success
		return nil
	}

	return fmt.Errorf(msg.Error)
}
