package analy

import (
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"shares/internal/api"
	"shares/internal/core"
	"shares/internal/model"
	"shares/internal/service/common"
	"shares/internal/service/event"
	"shares/internal/service/nlp"
	"shares/internal/service/weixin"
	proto "shares/rpc/shares"
	"sort"
	"strings"
	"time"

	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/tools"
	wx "github.com/xxjwxc/public/weixin"
	"gorm.io/datatypes"
)

func DealMsg(c *api.Context) {
	// resp := &proto.AnalyCodeResp{
	// 	Msgs: []string{""},
	// }
	body, _ := ioutil.ReadAll(c.GetGinCtx().Request.Body)
	requestBody := &TextRequestBody{}
	xml.Unmarshal(body, requestBody)
	if len(requestBody.MsgId) == 0 {
		requestBody.MsgId = fmt.Sprintf("%v_%v", requestBody.FromUserName, requestBody.CreateTime)
	}

	textResponseBody := TextResponseBody{
		ToUserName:   requestBody.FromUserName,
		FromUserName: requestBody.ToUserName,
		CreateTime:   time.Now().Unix(),
		MsgType:      "text",
		Content:      "",
	}

	orm := core.Dao.GetDBw()
	tmp, _ := model.WxMsgTblMgr(orm.DB).FetchUniqueByMsgID(requestBody.MsgId)
	if tmp.ID > 0 { // 表示已经处理过
		model.WxMsgTblMgr(orm.Where("id = ?", tmp.ID)).Update(model.WxMsgTblColumns.Status, -1) // 走客户回复消息
		textResponseBody.Content = "请稍等"
		c.GetGinCtx().XML(200, textResponseBody)
		return
	}

	tmp = model.WxMsgTbl{
		ToUserName:   requestBody.ToUserName,
		FromUserName: requestBody.FromUserName,
		Createtime:   int64(requestBody.CreateTime),
		MsgType:      requestBody.MsgType,
		MsgID:        requestBody.MsgId,
		Content:      requestBody.Content,
		CreatedAt:    time.Now(),
		Event:        requestBody.Event,
		MediaID:      requestBody.MediaId,
		Format:       requestBody.Format,
		Req:          string(body),
	}

	model.WxMsgTblMgr(orm.DB).Save(&tmp)

	if requestBody.MsgType != "text" {
		if requestBody.MsgType == "voice" { // 语音消息
			text, err := nlp.GetVoiceAsr(requestBody.FromUserName, requestBody.MediaId)
			if err != nil {
				textResponseBody.Content = "请让小潞想一想"
				model.WxMsgTblMgr(orm.Where("id = ?", tmp.ID)).Update(model.WxMsgTblColumns.Status, 1) // 回复完成
				c.GetGinCtx().XML(200, textResponseBody)
				return
			} else {
				requestBody.Content = text
				model.WxMsgTblMgr(orm.Where("id = ?", tmp.ID)).Update(model.WxMsgTblColumns.Content, requestBody.Content)
			}
		} else {
			if requestBody.MsgType == "event" {
				textResponseBody.Content = "SUCCESS"
			} else {
				textResponseBody.Content = "小潞未能明白您的意思"
			}
			model.WxMsgTblMgr(orm.Where("id = ?", tmp.ID)).Update(model.WxMsgTblColumns.Status, 1) // 回复完成
			c.GetGinCtx().XML(200, textResponseBody)
			return
		}
	}

	resp := dealWxMsg(requestBody.FromUserName, requestBody.Content)
	if resp.MsgType == "music" { // 音频
		musicResp := MusicResponseBody{
			ToUserName:   textResponseBody.ToUserName,
			FromUserName: textResponseBody.FromUserName,
			CreateTime:   textResponseBody.CreateTime,
			MsgType:      "music",
			Music: Music{
				Title:        resp.Title,
				Description:  resp.Description,
				MusicUrl:     resp.URL,
				HQMusicUrl:   resp.URL,
				ThumbMediaId: resp.Icon,
			},
		}
		c.GetGinCtx().XML(200, musicResp)

	} else {
		if len(resp.Answer) == 0 {
			resp.Answer = "小潞正在努力学习中。"
		}
		textResponseBody.Content = resp.Answer
		c.GetGinCtx().XML(200, textResponseBody)
	}

	tmp, _ = model.WxMsgTblMgr(orm.DB).FetchUniqueByMsgID(tmp.MsgID) // 先寻找消息
	if tmp.Status == -1 {                                            // 走客服消息
		var msg wx.CustomMsg
		if resp.MsgType == "music" {
			msg = wx.CustomMsg{
				Touser:  textResponseBody.ToUserName,
				Msgtype: "music",
				Music: &wx.CustomMusic{
					Title:        resp.Title,
					Description:  resp.Description,
					MusicUrl:     resp.URL,
					HQMusicUrl:   resp.URL,
					ThumbMediaId: resp.Icon,
				},
			}
		} else {
			msg = wx.CustomMsg{
				Touser:  textResponseBody.ToUserName,
				Msgtype: "text",
				Text: &wx.CustomText{
					Content: resp.Answer,
				},
			}
		}
		nlp.SendCustomMsg(msg)
	}
	model.WxMsgTblMgr(orm.Where("id = ?", tmp.ID)).Update(model.WxMsgTblColumns.Status, 1) // 回复完成
}

func WxTokenSignature(c *api.Context) {
	fmt.Println(c.GetGinCtx().Request.Method)
	timestamp := c.GetGinCtx().Query("timestamp")
	nonce := c.GetGinCtx().Query("nonce")
	signature := c.GetGinCtx().Query("signature")
	echostr := c.GetGinCtx().Query("echostr")
	signatureGen := makeSignature(timestamp, nonce)

	if signatureGen != signature {
		return
	}

	if strings.EqualFold(c.GetGinCtx().Request.Method, "get") { // 校验
		c.GetGinCtx().String(http.StatusOK, echostr)
	} else {
		DealMsg(c)
	}
}

func makeSignature(timestamp, nonce string) string { //本地计算signature
	si := []string{token, timestamp, nonce}
	sort.Strings(si)            //字典序排序
	str := strings.Join(si, "") //组合字符串
	s := sha1.New()             //返回一个新的使用SHA1校验的hash.Hash接口
	io.WriteString(s, str)      //WriteString函数将字符串数组str中的内容写入到s中
	return fmt.Sprintf("%x", s.Sum(nil))
}

func dealWxMsg(openID, text string) nlp.WxAnswer {
	if len(text) == 0 {
		return nlp.NewTextAnswer("小潞未能理解您的意思,请稍后再试")
	}

	if strings.EqualFold(text, "yy") { // 回复YY绑定监听(一分钟内)
		SetCacheWxMsgCode(openID, "yy", 1*time.Minute)
		return nlp.NewTextAnswer("请输入要绑定的股票代码或名称")
	}

	orm := core.Dao.GetDBr()
	// 判断是否是绑定组织
	ginfo, _ := model.GroupListTblMgr(orm.DB).GetFromKey(text)
	if ginfo.ID > 0 { // 组织
		isAdd, err := common.AddGroup(openID, text)
		if err != nil {
			return nlp.NewTextAnswer(err.Error())
		}

		if isAdd {
			return nlp.NewTextAnswer(fmt.Sprintf("成功加入组织:%v", ginfo.GroupName))
		} else {
			return nlp.NewTextAnswer(fmt.Sprintf("您已退出组织:%v", ginfo.GroupName))
		}
	}
	// -----------------

	// 绑定监听
	code := text
	ishan := tools.IsHan(code)
	if ishan { // 数据库查找
		condition := model.Condition{}
		condition.And(model.SharesInfoTblColumns.Name, "like", "%"+code+"%")
		where, args := condition.Get()
		orm := core.Dao.GetDBr()
		out, err := model.SharesInfoTblMgr(orm.Where(where, args...)).Get()
		if err != nil {
			return nlp.NewTextAnswer(err.Error())
		}
		if out.ID != 0 {
			code = out.Code
		} else { // 没找到
			tmp, err := nlp.GetAnswer(openID, code)
			if err != nil {
				return nlp.NewTextAnswer(err.Error())
			}
			return tmp
		}
	}

	// 开始查找
	sharesInfo := event.TrySearch(code)
	if sharesInfo == nil {
		return nlp.NewTextAnswer("未查询到,请完整输入股票代码查询")
	} else {
		code = sharesInfo.Code
	}

	// 判断是否绑定
	if GetCacheWxMsgCode(openID) == "yy" { // 绑定
		// 添加到数据库(先查询是否有)
		info, err := model.SharesWatchTblMgr(orm.DB).FetchUniqueIndexByOpenCode(code, openID)
		if err != nil {
			return nlp.NewTextAnswer(err.Error())
		}
		find := (info.ID != 0)

		var wxUserInfo model.WxUserinfo
		if !find {
			wxUserInfo, err = model.WxUserinfoMgr(orm.DB).GetFromOpenid(openID)
			if err != nil {
				return nlp.NewTextAnswer(err.Error())
			}
			if wxUserInfo.ID == 0 {
				return nlp.NewTextAnswer("未查询到用户信息，请检查是否有关注公众号")
			}

			if wxUserInfo.Capacity <= 0 { // 数量不够
				return nlp.NewTextAnswer("可用容量不够,请充值或联系管理员")
			}
		}
		info.Code = code
		info.OpenID = openID
		// info.Kdj = true
		info.Kdj20 = true
		info.Surge = true
		info.Slump = true
		info.Ai = true
		info.CreatedAt = time.Now()
		info.ExpiresTime = info.CreatedAt.AddDate(1, 0, 0)
		info.Vaild = true // 默认开启
		model.SharesWatchTblMgr(orm.DB).Save(&info)
		api.DeleteCacheBody(openID) // 删除cache

		if !find {
			wxUserInfo.Capacity -= 1
			model.WxUserinfoMgr(orm.Where("id = ?", wxUserInfo.ID)).Update(model.WxUserinfoColumns.Capacity, wxUserInfo.Capacity)
		}
		DeleteCacheWxMsgCode(openID) // 删除缓存
		return nlp.NewTextAnswer(fmt.Sprintf("小潞已为您开启(%v)提醒:20日线提醒,盘中急涨急跌提醒,AI智能提醒", sharesInfo.Name))
	}

	// 诊股
	go func() {
		msgs := AnalyCode(code)
		err := sendWXmsg(sharesInfo, openID, msgs)
		if err != nil {
			mylog.Error(err)
		}
	}()

	return nlp.NewTextAnswer("小潞正在疯狂计算中,请稍等片刻...")
}

func sendWXmsg(sharesInfo *proto.SharesInfo, openID string, msgs []string) error {
	if len(msgs) == 0 {
		return nil
	}

	orm := core.Dao.GetDBr()
	msg := model.MsgTbl{
		OpenID:    openID,
		Code:      sharesInfo.Code,
		Key:       "诊股",
		Desc:      strings.Join(msgs, "."),
		Price:     sharesInfo.Price,
		Percent:   sharesInfo.Percent,
		Day:       datatypes.Date(time.Now()),
		CreatedAt: time.Now(),
	}

	// 行业
	// hy, _ := model.HyGroupTblMgr(orm.DB).GetFromCode(sharesInfo.Code)
	// if hy.ID > 0 && len(hy.HyName) > 0 {
	// 	msgs = append(msgs, fmt.Sprintf("行业板块:%v", hy.HyName))
	// }

	var wxMsg []wx.TempWebMsg
	data := make(map[string]map[string]string)
	// mp := make(map[string]string)
	// mp["value"] = "你的分析任务已经运行完成"
	// //mp["color"] = v.color
	// data["first"] = mp

	mp := make(map[string]string)
	mp["value"] = fmt.Sprintf("%v(%v)", sharesInfo.Name, sharesInfo.Code)
	// mp["color"] = v.color
	data["keyword1"] = mp

	mp = make(map[string]string)
	mp["value"] = strings.Join(msgs, "\n")
	// mp["color"] = v.color
	data["keyword2"] = mp

	mp = make(map[string]string)
	mp["value"] = tools.GetTimeStr(time.Now())
	data["remark"] = mp

	wxMsg = append(wxMsg, wx.TempWebMsg{
		Touser:     openID,
		TemplateID: "6-QRNrivuHF56H8wzJyRmhfQQmO0fRmNZFrFH-c_wQE",
		Page:       fmt.Sprintf("https://hospital.xxjwxc.cn/webshares/#/pages/add/add?scode=%v&tag=%v", sharesInfo.Code, "daily"),
		Data:       data,
	})

	weixin.SendMsg(wxMsg)
	err := model.MsgTblMgr(orm.DB).Save(&msg).Error
	if err != nil {
		mylog.Error(err)
	}

	return err
}
