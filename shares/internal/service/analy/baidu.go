package analy

import (
	"fmt"
	"shares/internal/core"
	"shares/internal/model"
	"strings"
	"time"

	"github.com/xxjwxc/public/myhttp"
	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/tools"
)

type baiduIfs struct {
	accessToken string
	igs         map[string]bool
	igsTag      string // 忽略关键字
}

func (b *baiduIfs) Init() {
	now := time.Now().Add(30 * time.Minute)
	orm := core.Dao.GetDBr()
	info, _ := model.APITblMgr(orm.Where("tag = ?", "baidu")).Get()
	if info.ID > 0 { // 可能有
		b.igsTag = info.IgsTag
		b.igs = make(map[string]bool)
		ignores := strings.Split(info.Ignore, ",")
		for _, v := range ignores {
			b.igs[v] = true
		}

		if info.Expires.After(now) { // 有
			b.accessToken = info.AccessToken
			return
		}
	}

	var tmp BaiduResp
	myhttp.SendGet(fmt.Sprintf("https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&client_id=%v&client_secret=%v", info.APIKey, info.APISecret), "", &tmp)
	if tmp.ExpiresIn > 0 {
		info.AccessToken = tmp.AccessToken
		info.Expires = time.Now().Add(time.Duration(tmp.ExpiresIn) * time.Second)
		model.APITblMgr(orm.DB).Save(&info)
	}

	b.accessToken = info.AccessToken
	return
}

func NewBaidu() *baiduIfs {
	b := &baiduIfs{}
	b.Init()

	return b
}

// Lexer 百度分词
func (b *baiduIfs) Lexer(txt string) []BaiduItem {
	var resp BaiduLexerResp
	err := myhttp.SendPost(BaiduLexerReq{Text: txt}, &resp, fmt.Sprintf("https://aip.baidubce.com/rpc/2.0/nlp/v1/lexer?charset=UTF-8&access_token=%v", b.accessToken))
	if err != nil {
		mylog.Error(err)
	}

	if resp.ErrorCode == 18 || resp.ErrorCode == 2 {
		time.Sleep(1 * time.Second)
		return b.Lexer(txt)
	}

	if len(resp.Items) == 0 {
		fmt.Println(txt)
	}

	return resp.Items
}

// Lexer 百度分词
func (b *baiduIfs) LexerString(txt string, items []BaiduItem) string {
	var list []string
	for _, v := range items {
		if len(v.Pos) == 0 {
			v.Pos = v.Ne
		}
		list = append(list, fmt.Sprintf("%v/%v", v.Item, v.Pos))
	}
	return strings.Join(list, ",")
}

// Lexer 百度分词
func (b *baiduIfs) LexerTag(txt string, items []BaiduItem) string {
	var list []string
	for _, v := range items {
		if len(v.Pos) == 0 {
			v.Pos = v.Ne
		}
		if !b.igs[v.Pos] && tools.GetUtf8Len(v.Item) > 1 { // 不用忽略
			if !strings.Contains(b.igsTag, v.Item) { // 忽略关键字
				list = append(list, fmt.Sprintf("%v/%v", v.Item, v.Pos))
			}
		}
	}
	return strings.Join(list, ",")
}

// Lexer 百度分词
func (b *baiduIfs) Sentiment(txt string) []Sentiment {
	var resp SentimentResp
	err := myhttp.SendPost(BaiduLexerReq{Text: txt}, &resp, fmt.Sprintf("https://aip.baidubce.com/rpc/2.0/nlp/v1/sentiment_classify?charset=UTF-8&access_token=%v", b.accessToken))
	if err != nil {
		mylog.Error(err)
	}

	return resp.Items
}

// Lexer 百度分词 0:负向，1:中性，2:正向
func (b *baiduIfs) SentimentOut(txt string) int {
	items := b.Sentiment(txt)
	if len(items) > 0 {
		return items[0].Sentiment
	}
	return 1
}
