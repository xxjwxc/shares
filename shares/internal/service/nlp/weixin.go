package nlp

import (
	"time"

	"github.com/xxjwxc/public/myhttp"
	"github.com/xxjwxc/public/mylog"
)

// wxMessageOp 微信操作相关
type wxMessageOp struct {
	APIID     string
	APIKey    string
	APISecret string
}

// MessageResp response
type MessageResp struct {
	Answer string `json:"answer,omitempty"`
}

// Kv 内容信息
type Kv struct {
	Key   string // 类型(video:视频信息,img:图片信息)
	Value string // 内容
}

// getWxMessage 主入口
func getWxMessage(username, msg string) (WxAnswer, error) {
	re := &wxMessageOp{
		APIID:     "UP09unEHunIYcjM",
		APIKey:    "26jRjR3tFLD1YoP6kCX3oJl1U6TETiYC6DcvaeUyVVc",
		APISecret: "XA8YmDmwgSRXVkVhZsdkAoi4Nay7en",
	}

	return re.Aibot(username, username, msg)
}

type signInfo struct {
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Userid   string `json:"userid"`
}

type signResp struct {
	Signature string `json:"signature"`
	ExpiresIn int64  `json:"expiresIn"`
}

type aibotInfo struct {
	Signature string `json:"signature"`
	Query     string `json:"query"`
	Env       string `json:"env"` //默认是online, debug是测试环境,online是线上环境
}

type aibotResp struct {
	Answer     string `json:"answer,omitempty"`
	Status     string `json:"status,omitempty"`
	AnswerType string `json:"answer_type,omitempty"`
	Msg        []map[string]interface{}
}

type multiMsg struct {
	Multimsg []string `json:"multimsg"`
}

func (wx *wxMessageOp) sign(username, avatarID string) string {
	_tmp := GetCacheWxOpenapiSign(avatarID)
	if len(_tmp) > 0 {
		return _tmp
	}

	var sign signInfo
	sign.Username = username
	sign.Userid = avatarID
	var resp signResp
	if myhttp.SendPost(sign, &resp, "https://openai.weixin.qq.com/openapi/sign/"+wx.APISecret) == nil {
		SetCacheWxOpenapiSign(avatarID, resp.Signature, time.Duration(resp.ExpiresIn)*time.Second) // 设置过期时间
		return resp.Signature
	}

	return ""
}

// Aibot 微信 机器人
func (wx *wxMessageOp) Aibot(username, userid, query string) (WxAnswer, error) {
	mylog.Infof("username:%v, userid:%v, query:%v", username, userid, query)

	sg := wx.sign(username, userid)
	if len(sg) == 0 {
		return WxAnswer{}, nil
	}

	var resp aibotResp
	if myhttp.SendPost(aibotInfo{
		Signature: sg,
		Query:     query,
		Env:       "online", // "debug"
	}, &resp, "https://openai.weixin.qq.com/openapi/aibot/"+wx.APISecret) == nil {
		answer := WxAnswer{
			MsgType: resp.AnswerType,
			Answer:  resp.Answer,
		}
		if resp.AnswerType == "text" {
			return answer, nil
		} else if resp.AnswerType == "music" {
			answer.Title = resp.Msg[0]["song_name"].(string)
			answer.Description = resp.Answer
			answer.URL = resp.Msg[0]["music_url"].(string)
			answer.Icon = resp.Msg[0]["pic_url"].(string)
			return answer, nil
		}

		return answer, nil
	}

	return WxAnswer{}, nil
}
