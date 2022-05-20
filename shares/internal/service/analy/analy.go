package analy

import (
	"context"
	"encoding/xml"
	"fmt"
	"math"
	"shares/internal/core"
	"shares/internal/model"
	"shares/internal/service/event"
	proto "shares/rpc/shares"
	"sort"
	"strings"
	"time"

	"github.com/xxjwxc/public/tools"
)

const (
	token = "xxjwxc" //设置token
)

type TextRequestBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   float64
	MsgType      string
	Content      string
	Event        string
	MediaId      string // 语音消息媒体id，可以调用获取临时素材接口拉取数据。
	Format       string // 语音格式，如amr，speex等
	MsgId        string
}

type TextResponseBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	MsgId        string
	Content      string
}

type MusicResponseBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Music        Music
}

type Music struct {
	Title        string
	Description  string
	MusicUrl     string
	HQMusicUrl   string
	ThumbMediaId string
}

type Analy struct {
}

func (a *Analy) AnalyCode(ctx context.Context, req *proto.AnalyCodeReq) (*proto.AnalyCodeResp, error) {
	resp := &proto.AnalyCodeResp{
		Msgs: AnalyCode(req.Code),
	}
	return resp, nil
}

func AnalyCode(code string) []string {
	var msg []string
	msg = append(msg, AnalyMacd(code)...)  // 分析macd
	msg = append(msg, AnalyBs(code)...)    // 分析北上
	msg = append(msg, AnalyLHB(code)...)   // 分析龙虎榜
	msg = append(msg, AnalyZLJLR(code)...) // 主力净流入
	msg = append(msg, AnalyCjl(code)...)   // 主力净流入
	msg = append(msg, AnalyUP(code)...)    // 分析趋势
	msg = append(msg, AnalySzx(code)...)   // 底部缩量十字星
	msg = append(msg, AnalyXielv(code)...) // GNN 预测走势
	msg = append(msg, AnalyMsg(code)...)   // 分析消息数据

	return msg
}

// AnalyMacd  分析macd
func AnalyMacd(code string) []string {
	day0 := tools.GetUtcDay0(time.Now())
	offset := time.Now().Unix() - day0
	if offset >= (15 * 60 * 60) {
		day0 += 24 * 60 * 60
	}

	orm := core.Dao.GetDBr()
	list, _ := model.SharesDailyTblMgr(orm.Where("day0 < ?", day0).Order("day0 desc").Limit(3)).GetFromCode(code)
	if len(list) < 3 {
		return nil
	}

	// 金叉
	var jc string
	if list[2].Dif < list[2].Dea && list[0].Dea <= list[0].Dif {
		jc = ",金叉"
	}

	if list[2].Dif > list[2].Dea && list[0].Dea >= list[0].Dif {
		jc = ",死叉"
	}

	var src []string
	// macd  向上
	if list[2].Macd <= list[1].Macd && list[1].Macd <= list[0].Macd {
		src = append(src, fmt.Sprintf("MACD 呈上升趋势%v", jc))
	}

	// macd 向下
	if list[2].Macd > list[1].Macd && list[1].Macd > list[0].Macd {
		src = append(src, fmt.Sprintf("MACD 呈下降趋势%v", jc))
	}

	// 临界点
	if (list[2].Macd > 0 && list[0].Macd < 0) || (list[2].Macd < 0 && list[0].Macd > 0) {
		src = append(src, fmt.Sprintf("MACD 在零值交叉%v", jc))
	}

	if len(src) == 0 && len(jc) > 0 {
		src = append(src, fmt.Sprintf("MACD %v", jc))
	}

	return src
}

// AnalyBs  分析北上
func AnalyBs(code string) []string {
	day0 := tools.GetUtcDay0(time.Now())

	orm := core.Dao.GetDBr()
	list, _ := model.SharesDailyTblMgr(orm.Where("day0 < ?", day0).Order("day0 desc").Limit(10)).GetFromCode(code)

	var bsjlr float64
	for _, v := range list {
		bsjlr += v.Bsjlr
	}

	var src []string
	if bsjlr != 0 {
		src = append(src, fmt.Sprintf("近10日北上净流入: %v", getSamplePrice(bsjlr)))
	}

	return src
}

// AnalyLHB  分析龙虎榜
func AnalyLHB(code string) []string {
	day0 := tools.GetUtcDay0(time.Now().AddDate(0, -1, 0))

	orm := core.Dao.GetDBr()
	list, _ := model.LhbDailyTblMgr(orm.Where("day0 > ?", day0)).GetFromCode(code)

	var jlr float64
	var jgJlr float64
	for _, v := range list {
		jlr += v.Jlr
		jgJlr += v.JgJlr
	}

	var src []string
	if jlr != 0 {
		src = append(src, fmt.Sprintf("近一月龙虎榜总金额: %v", getSamplePrice(jlr*10000)))
	}
	if jgJlr != 0 {
		src = append(src, fmt.Sprintf("近一月龙虎榜机构净流入: %v", getSamplePrice(jgJlr*10000)))
	}

	return src
}

// AnalyZLJLR  主力净流入
func AnalyZLJLR(code string) []string {
	day0 := tools.GetUtcDay0(time.Now())
	offset := time.Now().Unix() - day0
	if offset >= (15 * 60 * 60) {
		day0 += 24 * 60 * 60
	}

	orm := core.Dao.GetDBr()
	list, _ := model.ZljlrDailyTblMgr(orm.Where("day0 < ?", day0).Order("day0 desc").Limit(10)).GetFromCode(code)

	var zljlr float64
	for _, v := range list {
		zljlr += v.Price
	}

	var src []string
	if zljlr != 0 {
		src = append(src, fmt.Sprintf("近10日主力净流入: %v", getSamplePrice(zljlr*10000)))
	}

	return src
}

// AnalyUP  分析趋势
func AnalyUP(code string) []string {
	orm := core.Dao.GetDBr()
	info, _ := model.AnalyUpTblMgr(orm.DB).GetFromCode(code)

	var src []string
	if info.ID != 0 {
		src = append(src, "k线呈上升趋势")
	}

	return src
}

// AnalyMsg 分析消息数据
func AnalyMsg(code string) []string {
	dayNum := 5
	analyOneCodeMsg(code, dayNum)
	day0 := tools.GetUtcDay0(time.Now())
	mp := make(map[string]int)
	var goodSentiment, badSentiment, sentiment int
	orm := core.Dao.GetDBr()
	list, _ := model.XqMsgDailyTblMgr(orm.Where("code = ? and day0 >= ? and day0 <= ?", code, day0-int64(dayNum*24*60*60), day0)).Gets()
	for _, v := range list {
		badSentiment += v.BadSentiment
		goodSentiment += v.GoodSentiment
		sentiment += v.Sentiment

		out := strings.Split(v.Tag, ",")
		for _, v1 := range out {
			tmps := strings.Split(v1, "/")
			if len(tmps) > 1 {
				mp[tmps[0]]++
			}
		}
	}

	sharesInfo, _ := event.GetShare(code, true)
	info, _ := model.APITblMgr(orm.Where("tag = ?", "tencent")).Get()

	var mpSlice MsgSortSlice
	for k, v := range mp {
		mpSlice = append(mpSlice, MsgSort{
			Msg: k,
			Sum: v,
		})
	}
	sort.Sort(mpSlice)
	var sresp []string
	for _, v := range mpSlice {
		if !strings.Contains(info.IgsTag, v.Msg) && !strings.Contains(sharesInfo.Name, v.Msg) { // 忽略关键字
			sresp = append(sresp, v.Msg)
			if len(sresp) > 10 {
				break
			}
		}
	}

	sum := float64(goodSentiment + badSentiment + sentiment)
	var src []string
	if sum > 0 {
		src = append(src, fmt.Sprintf("积极评论占比:%0.2f%%", (float64(goodSentiment)/sum)*100.0))
	}

	if len(sresp) > 0 {
		src = append(src, fmt.Sprintf("评论关键词:%v", strings.Join(sresp, ",")))
	}

	if sum > 200 {
		src = append(src, fmt.Sprintf("总筛选活跃评论数:%v条", sum))
	}

	return src
}

// AnalyCjl  分析成交量
func AnalyCjl(code string) []string {
	day0 := tools.GetUtcDay0(time.Now())
	offset := time.Now().Unix() - day0
	if offset >= (15 * 60 * 60) {
		day0 += 24 * 60 * 60
	}

	orm := core.Dao.GetDBr()
	list, _ := model.SharesDailyTblMgr(orm.Where("day0 < ?", day0).Order("day0 desc").Limit(10)).GetFromCode(code)
	if len(list) < 10 {
		return nil
	}
	var left, right float64
	for i := 5; i < 10; i++ {
		left += list[i].TurnoverRate
	}

	for i := 0; i < 5; i++ {
		right += list[i].TurnoverRate
	}
	if left == 0 || right == 0 {
		return nil
	}

	var src []string

	if left*3 < right { // 2倍以上成交量算放大
		src = append(src, "近期成交量:放巨量")
		return src
	}
	if left > right*3 { // 2倍以上成交量算缩小
		src = append(src, "近期成交量:缩巨量")
		return src
	}

	if left*2 < right { // 2倍以上成交量算放大
		src = append(src, "近期成交量:呈明显放大趋势")
		return src
	}
	if left > right*2 { // 2倍以上成交量算缩小
		src = append(src, "近期成交量:呈明显缩小趋势")
		return src
	}

	if left*1.3 < right { // 2倍以上成交量算放大
		src = append(src, "近期成交量:增多")
		return src
	}
	if left > right*1.3 { // 2倍以上成交量算缩小
		src = append(src, "近期成交量:减少")
		return src
	}

	return nil
}

// AnalyCjl  分析十字星
func AnalySzx(code string) []string {
	day0 := tools.GetUtcDay0(time.Now())
	orm := core.Dao.GetDBr()

	list, _ := model.SharesDailyTblMgr(orm.Where("day0 < ?", day0).Order("day0 desc").Limit(20)).GetFromCode(code)
	if len(list) < 20 { // 不需要
		return nil
	}

	sharesInfo := event.SearchDetail(code)
	// 底部缩量十字星
	// 底部(20天，只有3天低于当前价格)
	lowNumber := 0
	lowSl := 0
	for _, v := range list {
		if v.Price < sharesInfo.Price { // 底部
			lowNumber++
		}
		if v.TurnoverRate > 0 && v.TurnoverRate < sharesInfo.TurnoverRate { // 缩量
			lowSl++
		}
	}
	if lowNumber > 3 { // (20天，只有3天低于当前价格)
		return nil
	}

	if lowSl > 3 { // (20天，只有3天低于当前量能)
		return nil
	}

	// 十字星
	left := sharesInfo.OpenPrice
	right := sharesInfo.Price
	if left > right {
		left = sharesInfo.Price
		right = sharesInfo.OpenPrice
	}
	p1 := left - sharesInfo.Min
	p2 := right - left
	p3 := sharesInfo.Max - right

	var src []string
	// 百分比小于0.5
	if (p2 < sharesInfo.OpenPrice*0.005) && (p2*3 < p1) && (p2*3 < p3) { // 十字星
		src = append(src, "底部缩量十字星")
	}

	return src
}

// Xielv  斜率
func AnalyXielv(code string) []string {
	day0 := tools.GetUtcDay0(time.Now())
	orm := core.Dao.GetDBr()

	list, _ := model.SharesDailyTblMgr(orm.Where("day0 <= ?", day0).Order("day0 desc").Limit(30)).GetFromCode(code)
	if len(list) < 30 { // 不需要
		return nil
	}

	_len := len(list)
	var points []Point
	for i := len(list) - 1; i >= 0; i-- {
		points = append(points, Point{
			X: float64(_len - i),
			Y: list[i].Price,
		})
	}

	// 获取首尾直线
	off := len(list) - 1
	a := points[off].Y - points[0].Y
	b := (points[off].X - points[0].X)
	c := (points[off].X * points[0].Y) - (points[off].Y * points[0].X)
	p := Point{ // https://blog.csdn.net/qq_27278957/article/details/107962641?spm=1001.2101.3001.6661.1&utm_medium=distribute.pc_relevant_t0.none-task-blog-2%7Edefault%7ECTRLIST%7Edefault-1.no_search_link&depth_1-utm_source=distribute.pc_relevant_t0.none-task-blog-2%7Edefault%7ECTRLIST%7Edefault-1.no_search_link
		X: a / b,
		Y: c / b,
	}
	// l, p := linearRegressionLSE(points)
	// fmt.Println(l, p)

	// 计算距离
	index := 0
	var l float64
	for i, v := range points {
		tmp := math.Sqrt(math.Pow((p.X*v.X)-(1*v.Y)+p.Y, 2) / (math.Pow(v.X, 2) + math.Pow(v.Y, 2)))
		if tmp > l {
			l = tmp
			index = i
		}
	}
	// 通过直线计算距离

	_, p0 := linearRegressionLSE(points)
	price0 := 31*p0.X + p0.Y
	var jp0 float64
	for _, v := range points {
		jp0 += math.Sqrt(math.Pow((p0.X*v.X)-(1*v.Y)+p0.Y, 2) / (math.Pow(v.X, 2) + math.Pow(v.Y, 2)))
	}
	jp0 = (jp0 / float64(len(points))) * 100

	var src []string
	if l > 0.08 && index > 2 && index < 28 { // 需要做分段处理
		// _, p0 := linearRegressionLSE(points)
		// fmt.Println(p0)
		// _, p1 := linearRegressionLSE(points[0:index])
		// fmt.Println(p1)
		_, p2 := linearRegressionLSE(points[index:])
		price2 := 31*p2.X + p2.Y

		var jp2 float64
		for _, v := range points[index:] {
			jp2 += math.Sqrt(math.Pow((p2.X*v.X)-(1*v.Y)+p2.Y, 2) / (math.Pow(v.X, 2) + math.Pow(v.Y, 2)))
		}
		jp2 = (jp2 / float64(len(points[index:]))) * 100

		src = append(src, fmt.Sprintf("%v预测:%0.2f(波动率:%0.2f) , %0.2f(波动率:%0.2f)", getNextTimeDayStr(list[0].Day0), price0, jp0, price2, jp2))

	} else {
		src = append(src, fmt.Sprintf("%v预测:%0.2f(波动率:%0.2f)", getNextTimeDayStr(list[0].Day0), price0, jp0))
	}

	return src
}
