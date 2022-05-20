package nlp

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"shares/internal/config"
	"shares/internal/core"
	"shares/internal/model"
	"strings"
	"time"
	"unicode"

	asr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/asr/v20190614"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/tools"
	wx "github.com/xxjwxc/public/weixin"
)

type PostFileResp struct {
	Errcode string `json:"errcode"`
	Type    string `json:"type"`
	MediaId string `json:"media_id"`
}

type WxAnswer struct {
	MsgType     string
	Answer      string
	Title       string // 标题
	Description string // 描述
	URL         string // 链接
	Icon        string // 缩略图
}

var _wx wx.WxTools

func init() {
	info := config.GetWxInfo()
	t, err := wx.New(wx.WxInfo{
		AppID:     info.AppID,
		AppSecret: info.AppSecret,
		APIKey:    info.APIKey,
		MchID:     info.MchID,
		NotifyURL: info.NotifyURL,
		ShearURL:  info.ShearURL,
	})
	if err != nil {
		mylog.Fatal(err)
	}
	_wx = t
}

func NewTextAnswer(msg string) WxAnswer {
	return WxAnswer{
		MsgType: "text",
		Answer:  msg,
	}
}

func GetAnswer(username, msg string) (WxAnswer, error) {
	out, err := getWxMessage(username, msg)
	if out.MsgType == "music" {
		res, err := http.Get(out.Icon)
		if err != nil {
			return out, err
		}
		//defer res.Close()

		accessToken, err := _wx.GetAccessToken()
		if err != nil {
			return out, err
		}
		tmp, err := PostFile("icon.jpg", fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/media/upload?type=image&access_token=%v", accessToken), res.Body)
		if err != nil {
			return out, err
		}
		out.Icon = tmp.MediaId
	}

	return out, err
}

//PostFile 模拟客戶端文件上传
//fieldname注意与服务器端保持一致
func PostFile(filename, targetURL string, fh io.Reader) (result PostFileResp, e error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		e = err
		return
	}

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		e = err
		return
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targetURL, contentType, bodyBuf)
	if err != nil {
		e = err
		return
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		e = err
		return
	}
	tools.JSONEncode(string(respBody), &result)
	return
}

func GetVoiceAsr(username, mediaId string) (string, error) {
	accessToken, err := _wx.GetAccessToken()
	if err != nil {
		return "", err
	}

	resp, err := http.Get(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/media/get?media_id=%v&access_token=%v", mediaId, accessToken))
	if err != nil {
		mylog.Error(err)
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		mylog.Error(err)
		return "", err
	}

	b64 := tools.Base64Encode(body)

	orm := core.Dao.GetDBr()
	info, _ := model.APITblMgr(orm.Where("tag = ?", "tencent")).Get()
	if info.ID == 0 { // 可能有
		return "", err
	}

	credential := common.NewCredential(
		info.APISecret,
		info.APIKey,
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "asr.tencentcloudapi.com"
	client, _ := asr.NewClient(credential, "", cpf)
	request := asr.NewCreateRecTaskRequest()

	request.EngineModelType = common.StringPtr("16k_zh")
	request.ChannelNum = common.Uint64Ptr(1)
	request.ResTextFormat = common.Uint64Ptr(0)
	request.SourceType = common.Uint64Ptr(1)
	request.Data = common.StringPtr(b64)

	response, err := client.CreateRecTask(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return "", err
	}
	if err != nil {
		return "", err
	}
	var vResp TencentVoiceResp
	tools.JSONEncode(response.ToJsonString(), &vResp)

	var text string
	var times [5][0]int
	for range times { // 轮训获取任务结果
		request1 := asr.NewDescribeTaskStatusRequest()
		request1.TaskId = common.Uint64Ptr(uint64(vResp.Response.Data.TaskId))
		response1, err := client.DescribeTaskStatus(request1)
		if _, ok := err.(*errors.TencentCloudSDKError); ok {
			fmt.Printf("An API error has returned: %s", err)
			return "", err
		}
		if err != nil {
			return "", err
		}

		fmt.Println(response1.ToJsonString())
		var tmp TencentVoiceResp
		tools.JSONEncode(response1.ToJsonString(), &tmp)
		if tmp.Response.Data.Status > 1 { // 表示已经结束
			tmp.Response.Data.Result = strings.Replace(tmp.Response.Data.Result, "\n", "", -1)
			list := strings.Split(tmp.Response.Data.Result, "  ")
			if len(list) > 1 {
				text = trimString([]rune(tools.DbcToSbc(list[1])))
			}
			break
		}
		time.Sleep(1 * time.Second)
	}

	return text, nil
}

// 去掉首尾符号加空格
func trimString(r []rune) string {
	begin := 0
	end := len(r)
	for i := begin; i < len(r); i++ {
		if r[i] == ' ' || unicode.IsPunct(r[i]) {
			begin++
		} else {
			break
		}
	}
	for i := end; i > 0; i-- {
		if r[i-1] == ' ' || unicode.IsPunct(r[i-1]) {
			end--
		} else {
			break
		}
	}

	if begin >= end { // 空
		return ""
	}

	if (end - begin) == 1 {
		tmp := r[begin:end]
		if unicode.IsPunct(tmp[0]) {
			return ""
		}
	}

	return string(r[begin:end])
}

// SendCustomMsg 发送客服消息
func SendCustomMsg(msg wx.CustomMsg) error {
	return _wx.SendCustomMsg(msg)
}
