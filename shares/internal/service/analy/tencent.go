package analy

import (
	"fmt"
	"shares/internal/core"
	"shares/internal/model"
	"strings"
	"time"

	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/tools"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	nlp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/nlp/v20190408"
)

type TencentIfs struct {
	client *nlp.Client
	igsTag string // 忽略关键字
}

func (b *TencentIfs) Init() {
	orm := core.Dao.GetDBr()
	info, _ := model.APITblMgr(orm.Where("tag = ?", "tencent")).Get()
	if info.ID > 0 { // 可能有
		b.igsTag = info.IgsTag
	}

	credential := common.NewCredential(
		info.APISecret,
		info.APIKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "nlp.tencentcloudapi.com"
	client, err := nlp.NewClient(credential, "ap-guangzhou", cpf)
	if err != nil {
		mylog.Error(err)
	}

	b.client = client
	return
}

func NewTencent() *TencentIfs {
	b := &TencentIfs{}
	b.Init()

	return b
}

// Lexer 腾讯分词
func (b *TencentIfs) Lexer(txt string) []TencentItem {
	if b.client == nil {
		return []TencentItem{}
	}
	request := nlp.NewKeywordsExtractionRequest()
	request.Text = common.StringPtr(txt)
	response, err := b.client.KeywordsExtraction(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		mylog.Errorf("An API error has returned: %s", err)
		time.Sleep(1 * time.Second)
		return b.Lexer(txt)
	}
	if err != nil {
		mylog.Error(err)
	}

	var tmp TencentResp
	tools.JSONEncode(response.ToJsonString(), &tmp)
	// mylog.Infof("%s", response.ToJsonString())

	return tmp.Response.Keywords
}

// Lexer 分词
func (b *TencentIfs) LexerString(items []TencentItem) string {
	var list []string
	for _, v := range items {
		list = append(list, fmt.Sprintf("%v/%0.2f", v.Item, v.Score))
	}
	return strings.Join(list, ",")
}

// Lexer 分词
func (b *TencentIfs) LexerTag(items []TencentItem) string {
	var list []string
	for _, v := range items {
		if v.Score > 0.7 { // 不用忽略
			if !strings.Contains(b.igsTag, v.Item) { // 忽略关键字
				list = append(list, fmt.Sprintf("%v/%0.2f", v.Item, v.Score))
			}
		}
	}
	return strings.Join(list, ",")
}
