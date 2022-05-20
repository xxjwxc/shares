package analy

import (
	"fmt"
	"regexp"
	"shares/internal/core"
	"shares/internal/model"
	"shares/internal/service/event"
	"sort"
	"strings"
	"time"

	"github.com/xxjwxc/public/myhttp"
	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/tools"
	"gorm.io/datatypes"
)

// loop循环抓取所有
func initMsg() {
	for {
		codes := dealMsg(getCode())
		if len(codes) > 0 {
			orm := core.Dao.GetDBw()
			model.XqMsgDailyActiveTblMgr(orm.Where("day0 < ?", datatypes.Date(time.Now()))).Delete(&model.XqMsgDailyActiveTbl{})
			for _, v := range codes {
				sharesInfo := event.SearchDetail(v.Code)
				model.XqMsgDailyActiveTblMgr(orm.DB).Save(&model.XqMsgDailyActiveTbl{
					Code:         v.Code,
					Name:         sharesInfo.Name,
					Day0:         datatypes.Date(time.Now()),
					DayStr:       tools.GetDayStr(time.Now()),
					Price:        sharesInfo.Price,
					Percent:      sharesInfo.Percent,
					TurnoverRate: sharesInfo.TurnoverRate,
					CreatedAt:    time.Now(),
					Count:        v.Count,
				})
			}
			time.Sleep(10 * time.Minute)
			analyAllMsg() // 分析统计每日消息
		}
	}
}

// func Temp() {
// 	tencent := NewTencent()
// 	orm := core.Dao.GetDBw()
// 	list, _ := model.XqMsgTblMgr(orm.DB).Gets()
// 	for _, v := range list {
// 		msg := v.Msg
// 		items := tencent.Lexer(msg)
// 		v.Lexer = tencent.LexerString(items)
// 		v.Tag = tencent.LexerTag(items)
// 		model.XqMsgTblMgr(orm.DB).Save(&v)
// 	}
// 	os.Exit(1)
// }

// 诊股
func analyOneCodeMsg(code string, dayNum int) {
	dealOneMsg(code, 20)

	gmp := make(map[string]int)
	var ginfo model.XqMsgDailyTbl

	day0 := tools.GetUtcDay0(time.Now())
	for i := 0; i < dayNum; i++ {
		_analyDailyMsg(code, (day0 - int64(i*24*60*60)), gmp, &ginfo)
	}
}

// 分析
func analyAllMsg() { // 统计昨天的数据
	day0 := tools.GetUtcDay0(time.Now()) - (24 * 60 * 60)
	_analyAllMsg(day0)
	day0 = tools.GetUtcDay0(time.Now())
	_analyAllMsg(day0)
}

func _analyAllMsg(day0 int64) {
	// day0 := tools.GetUtcDay0(time.Now()) - (24 * 60 * 60)
	orm := core.Dao.GetDBw()
	var codes []string
	model.SharesInfoTblMgr(orm.DB).Select("code").Find(&codes)
	gmp := make(map[string]int)
	ginfo, _ := model.XqMsgDailyTblMgr(orm.DB).FetchUniqueIndexByCode("", day0)

	for _, v := range codes {
		_analyDailyMsg(v, day0, gmp, &ginfo)
	}

	// 全局
	ginfo.Code = ""
	ginfo.Day0 = day0
	ginfo.Day0Str = tools.GetDayStr(tools.UnixToTime(day0))
	ginfo.CreatedAt = time.Now()
	var mpSlice MsgSortSlice
	for k, v := range gmp {
		mpSlice = append(mpSlice, MsgSort{
			Msg: k,
			Sum: v,
		})
	}
	sort.Sort(mpSlice)
	var mps []string
	for _, v := range mpSlice {
		mps = append(mps, fmt.Sprintf("%v/%v", v.Msg, v.Sum))
	}
	ginfo.Tag = strings.Join(mps, ",")
	model.XqMsgDailyTblMgr(orm.DB).Save(&ginfo)
}

func getCode() []string {
	orm := core.Dao.GetDBw()
	var codes []string
	model.MsgCodeViewMgr(orm.DB).Select("DISTINCT code").Find(&codes)
	return codes
}

// 获取股评信息
func dealMsg(codes []string) []dealMsgResp {
	day0 := tools.GetDay0(time.Now().Unix())
	var out string
	client := myhttp.NewMyHttpClient()
	client.SendGet("https://xueqiu.com/S/SH000001", "", &out)
	var re []dealMsgResp
	baidu := NewBaidu()
	tencent := NewTencent()
	orm := core.Dao.GetDBw()
	for _, v := range codes {
		resp := getXueQiuMsg(client, v, 1)
		if len(resp.List) == 0 {
			mylog.Errorf("not find msg :%v", v)
		}

		sharesInfo := event.SearchDetail(v)
		count := 0
		var maxTm int64
		model.XqMsgTblMgr(orm.Where("code = ?", v).Order("tm desc").Limit(1)).Select("tm").Find(&maxTm) // 获取当前最新时间
		for _, v1 := range resp.List {
			v1.CreatedAt = v1.CreatedAt / 1000
			tmpday0 := tools.GetDay0(v1.CreatedAt)
			// if tmpday0.Equal(day0) {
			// 	count++
			// }
			if maxTm < v1.CreatedAt && v1.UserId > 0 && len(v1.Title) == 0 {
				if tools.GetUtf8Len(v1.Text) > 1024 {
					v1.Text = v1.Description
				}
				msg := removeHtml(v1.Text)
				if tools.GetUtf8Len(msg) > 2 {
					if tmpday0.Equal(day0) {
						count++
					}
					items := tencent.Lexer(msg)
					model.XqMsgTblMgr(orm.DB).Save(&model.XqMsgTbl{
						Code:         v,
						Tm:           v1.CreatedAt,
						Name:         sharesInfo.Name,
						Day:          tools.UnixToTime(v1.CreatedAt),
						DayStr:       tools.GetDayStr(tmpday0),
						Msg:          msg,
						Lexer:        tencent.LexerString(items),
						Tag:          tencent.LexerTag(items),
						Sentiment:    baidu.SentimentOut(msg),
						Price:        sharesInfo.Price,
						Percent:      sharesInfo.Percent,
						TurnoverRate: sharesInfo.TurnoverRate,
						CreatedAt:    time.Now(),
					})
				}
			}
		}
		if count > 15 {
			re = append(re, dealMsgResp{
				Code:  v,
				Count: count,
			})
		}
	}

	return re
}

// 获取一只股票股评信息
func dealOneMsg(code string, pages int) {
	var out string
	client := myhttp.NewMyHttpClient()
	client.SendGet("https://xueqiu.com/S/SH000001", "", &out)
	baidu := NewBaidu()
	tencent := NewTencent()
	sharesInfo := event.SearchDetail(code)
	orm := core.Dao.GetDBw()
	for i := 1; i < pages; i++ {
		resp := getXueQiuMsg(client, code, i)
		if len(resp.List) == 0 {
			mylog.Errorf("not find XueQiu msg :%v", code)
		}

		for _, v1 := range resp.List {
			v1.CreatedAt = v1.CreatedAt / 1000
			if v1.UserId > 0 && len(v1.Title) == 0 {
				if tools.GetUtf8Len(v1.Text) > 1024 {
					v1.Text = v1.Description
				}
				msg := removeHtml(v1.Text)
				if tools.GetUtf8Len(msg) > 2 {
					info, _ := model.XqMsgTblMgr(orm.DB).FetchUniqueIndexByCode(code, v1.CreatedAt)
					if info.ID == 0 { // 未找到,更新进去
						tmp, _ := model.XqMsgTblMgr(orm.DB).FetchUniqueIndexByCode(code, v1.CreatedAt)
						if tmp.ID == 0 {
							items := tencent.Lexer(msg)
							model.XqMsgTblMgr(orm.DB).Save(&model.XqMsgTbl{
								Code:         code,
								Tm:           v1.CreatedAt,
								Name:         sharesInfo.Name,
								Day:          tools.UnixToTime(v1.CreatedAt),
								DayStr:       tools.GetDayStr(tools.GetDay0(v1.CreatedAt)),
								Msg:          msg,
								Lexer:        tencent.LexerString(items),
								Tag:          tencent.LexerTag(items),
								Sentiment:    baidu.SentimentOut(msg),
								Price:        sharesInfo.Price,
								Percent:      sharesInfo.Percent,
								TurnoverRate: sharesInfo.TurnoverRate,
								CreatedAt:    time.Now(),
							})
						}

					}
				}
			}
		}
		if i >= resp.MaxPage {
			break
		}
	}
}

// 分析一只票一天数据
func _analyDailyMsg(code string, day0 int64, gmp map[string]int, ginfo *model.XqMsgDailyTbl) {
	orm := core.Dao.GetDBw()
	list, _ := model.XqMsgTblMgr(orm.Where("code = ? and tm > ? and tm < ?", code, day0, day0+(24*60*60))).Gets()
	if len(list) == 0 {
		return
	}
	sharesInfo := event.SearchDetail(code)
	info, _ := model.XqMsgDailyTblMgr(orm.DB).FetchUniqueIndexByCode(code, day0)
	info.Code = code
	info.Name = sharesInfo.Name
	info.Day0 = day0
	info.Day0Str = tools.GetDayStr(tools.UnixToTime(day0))
	// info.Tag = ""
	info.GoodSentiment = 0
	info.BadSentiment = 0
	info.Sentiment = 0
	info.Price = sharesInfo.Price
	info.Percent = sharesInfo.Percent
	info.TurnoverRate = sharesInfo.TurnoverRate
	info.CreatedAt = time.Now()

	mp := make(map[string]int)
	for _, v := range list {
		switch v.Sentiment { // (0:负向，1:中性，2:正向)
		case 0:
			info.BadSentiment++
			ginfo.BadSentiment++
		case 1:
			info.Sentiment++
			ginfo.Sentiment++
		case 2:
			info.GoodSentiment++
			ginfo.GoodSentiment++
		}
		out := strings.Split(v.Tag, ",")
		for _, v1 := range out {
			tmps := strings.Split(v1, "/")
			if len(tmps) > 1 {
				mp[tmps[0]]++
				gmp[tmps[0]]++
			}
		}
	}

	{
		var mpSlice MsgSortSlice
		for k, v := range mp {
			mpSlice = append(mpSlice, MsgSort{
				Msg: k,
				Sum: v,
			})
		}
		sort.Sort(mpSlice)
		var mps []string
		for _, v := range mpSlice {
			mps = append(mps, fmt.Sprintf("%v/%v", v.Msg, v.Sum))
		}

		info.Tag = strings.Join(mps, ",")
		model.XqMsgDailyTblMgr(orm.DB).Save(&info)
	}
}

func getXueQiuMsg(client *myhttp.MyHttpClient, code string, page int) XueQiuMsgResp {
	var resp XueQiuMsgResp
	var times [5][0]int
	for range times {
		resp = XueQiuMsgResp{}
		err := client.SendGet(fmt.Sprintf("https://xueqiu.com/query/v1/symbol/search/status.json?count=30&comment=0&symbol=%v&hl=0&source=all&sort&page=%v&q&type=11", strings.ToUpper(code), page), "", &resp)
		if err != nil {
			mylog.Error(err)
			continue
		}
		if resp.Status == 0 {
			break
		}
		time.Sleep(time.Duration(tools.GetRandInt(1, 3)) * time.Second)
	}
	return resp
}

func removeHtml(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	// //去除STYLE
	// re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	// src = re.ReplaceAllString(src, "")

	// //去除SCRIPT
	// re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	// src = re.ReplaceAllString(src, "")

	//去除<a>
	re, _ = regexp.Compile("\\<a[\\S\\s]+?\\</a\\>")
	src = re.ReplaceAllString(src, "")

	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "")

	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "")

	src = strings.Replace(src, "&nbsp;", "", -1)
	src = tools.UnicodeEmojiCode(src) // emoji

	// re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	// src = re.ReplaceAllString(src, "")

	return strings.TrimSpace(src)
}
