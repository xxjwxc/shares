package event

import (
	"encoding/json"
	"fmt"
	"shares/internal/config"
	"shares/internal/core"
	"shares/internal/model"
	"shares/rpc/shares"
	"strconv"
	"strings"
	"time"

	"github.com/xxjwxc/public/message"
	"github.com/xxjwxc/public/mycache"
	"github.com/xxjwxc/public/myhttp"
	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/mysqldb"
	"github.com/xxjwxc/public/timerDeal"
	"github.com/xxjwxc/public/tools"
	"gorm.io/datatypes"
)

func init() {
	// initDaily() // 初始化日信息
	// DayAfter()
	if config.GetIsTools() > 0 {
		return
	}
	// DayAfter()
	// initDaily()
	timerDeal.OnPeDay(9, 30, 0, OnDealDay) // 每天9点30分执行
	timerDeal.OnPeDay(8, 0, 0, maBS)       // 每天8点执行()
	timerDeal.OnPeDay(20, 0, 0, maLhb)     // 每天20点执行()
	// now := time.Now()
	// orm := core.Dao.GetDBr()
	// out, _ := model.NoTradingDayTblCopy1Mgr(orm.DB).GetFromDay(datatypes.Date(now))
	// if out.ID > 0 { // 休息时间
	// 	return
	// }

	// if now.Hour() > 9 { // 补一刀
	// 	OnDealDay()
	// }
}

func OnDealDay() { // 每天执行
	now := time.Now()
	day0 := tools.GetUtcDay0(now)
	offset := now.Unix() - day0
	week := tools.GetTimeWeek(day0)
	if week == 6 || week == 0 {
		return
	}

	if offset < (9*60*60) || offset > (15*60*60+30*60) { // 3点之后是空闲时间
		return
	}

	out, _ := model.NoTradingDayTblMgr(core.Dao.GetDBr().DB).GetFromDay(datatypes.Date(now))
	if out.ID > 0 { // 休息时间
		return
	}

	// 开始心跳更新
	go func() {
		_analy.Init() // 重置
		for {
			now = time.Now()
			day0 := tools.GetUtcDay0(now)
			offset := now.Unix() - day0
			if offset >= (15*60*60 + 1*60) { // 3点之后是空闲时间
				break
			}
			ticker := time.NewTicker(_timeoutTicker * time.Second)
			<-ticker.C
			if offset > (11*60+30)*60 && offset < (13*60*60+60) { // 盘中休息
				mylog.Infof("timer 盘中休息.....:%v", tools.GetTimeStr(time.Now()))
			} else {
				mylog.Infof("timer 盘中执行.....:%v", tools.GetTimeStr(time.Now()))
				OnDeal() //以下为定时执行的操作
			}
		}
		// TODO:每天更新下名字
		ticker := time.NewTicker(3 * time.Minute)
		<-ticker.C
		mylog.Infof("timer 执行收盘.....:%v", tools.GetTimeStr(time.Now()))
		//以下为定时执行的操作
		DayAfter()
	}()
}

func OnDeal() {
	var codes []string
	orm := core.Dao.GetDBw()
	model.SharesWatchTblMgr(orm.Where("vaild = ?", true).Group("code")).Select("code").Find(&codes)
	var outs []*shares.SharesInfo
	num := 500 // 一次处理 500 条
	for {
		if len(codes) > num {
			outs = append(outs, Searchs(codes[:num])...)
			codes = codes[num:]
		} else {
			outs = append(outs, Searchs(codes)...)
			break
		}
	}
	cache := mycache.NewCache(CacheCode)
	now := time.Now().Unix()
	for _, v := range outs { // 保存
		cache.Add(v.Code, v, CacheCodeTimeout) // 保存缓存
		_analy.Add(v, now)                     // 添加
	}
	_analy.StartAnaly(outs) // 实时分析
}

// DayAfter 盘后执行
func DayAfter() {
	_analy.Init() // 重置
	var codes []string
	orm := core.Dao.GetDBr()
	model.SharesInfoTblMgr(orm.DB).Select("code").Find(&codes)
	for {
		if len(codes) > 500 {
			dayAfter(codes[:500])
			codes = codes[500:]
		} else {
			dayAfter(codes)
			break
		}
	}
}

func dayAfter(codes []string) {
	orm := core.Dao.GetDBw()
	outs := SearchDetails(codes)
	if len(outs) != len(codes) {
		mylog.Error("not ssss")
	}
	day0 := tools.GetUtcDay0(time.Now())
	for _, v := range outs { // 保存
		model.SharesInfoTblMgr(orm.Where("code = ?", v.Code)).Updates(map[string]interface{}{"price": v.Price, "percent": v.Percent, "name": v.Name, "simple_name": FirstLetterOfPinYin(v.Name), "total": v.Total, "cap": v.Cap})
		info, _ := model.SharesDailyTblMgr(orm.DB).FetchUniqueIndexByCode(v.Code, day0)
		info.Code = v.Code
		info.Price = v.Price
		info.Percent = v.Percent
		info.Day0 = day0
		info.Day0Str = tools.GetDayStr(time.Now())
		info.Volume = v.Volume
		info.Turnover = v.Turnover
		info.TurnoverRate = v.TurnoverRate
		info.Max = v.Max
		info.Min = v.Min
		info.Pe = v.Pe
		info.Pb = v.Pb
		info.Total = v.Total
		info.Cap = v.Cap
		info.CreatedBy = "system"
		info.Zljlr = v.Zljlr
		info.OpenPrice = v.OpenPrice
		info.ClosePrice = v.ClosePrice
		info.Macd = v.Macd
		info.Dea = v.Dea
		info.Dif = v.Dif
		info.CreatedAt = time.Now()
		if info.Volume == 0 { // 表示停牌或者其它
			continue
		}
		model.SharesDailyTblMgr(orm.DB).Save(&info)

		tmp, _ := model.ZljlrDailyTblMgr(orm.DB).FetchUniqueIndexByCode(v.Code, day0)
		if tmp.ID == 0 {
			tmp.Code = v.Code
			tmp.Day0 = day0
			tmp.DayStr = tools.GetDayStr(time.Now())

			tmp.CreatedAt = time.Now()
		}
		tmp.Price = v.Zljlr
		orm.Save(&tmp)
		SetDailyMa(v.Code, day0, orm)
	}

	countHYZLJLR(day0) // 计算行业主力净流入
}

func initDaily() {
	var codes []string
	orm := core.Dao.GetDBw()

	// model.SharesInfoTblMgr(orm.DB).Select("code").Find(&codes)
	orm.Table("(SELECT code,count(*) as num FROM `shares_daily_tbl` GROUP BY code) as u").Where("num < 20").Select("code").Find(&codes)
	for _, v := range codes {
		BuildOneDaily(v, orm)
	}
}

// BuildOneDaily 创建一个日信息
func BuildOneDaily(code string, orm *mysqldb.MySqlDB) error {
	var count int64
	err := model.SharesDailyTblMgr(orm.Where("code = ?", code)).Count(&count).Error
	if err != nil {
		return err
	}

	if count > 20 { // 有,不需要再更新了
		return nil
	}

	// 开始获取
	out := SendGet(fmt.Sprintf("https://web.ifzq.gtimg.cn/appstock/app/fqkline/get?param=%v,day,,,30,qfq", code), "")
	out = tools.ConvertToString(out, "gbk", "utf8")
	out = strings.Replace(out, code, "info", -1)
	out = strings.Replace(out, "qfqday", "day", -1)

	var info DailyWWW
	err = json.Unmarshal([]byte(out), &info)
	if err != nil {
		mylog.Error(err)
		return err
	}

	if info.Code != 0 { // 不成功
		return message.GetError(message.InValidOp)
	}

	mgr := model.SharesDailyTblMgr(orm.DB)
	var day0 int64
	for _, v := range info.Data.Info.Day {
		if len(v) == 6 {
			tmpDay0 := tools.StrToTime(v[0].(string), "2006-01-02", time.Local).Unix()
			info, _ := mgr.Reset().FetchUniqueIndexByCode(code, tmpDay0)
			if info.ID == 0 {
				info.Code = code
				info.Day0 = tmpDay0
				info.Day0Str = v[0].(string)
				info.CreatedAt = time.Now()
			}
			info.Price = stringToFloat(v[2].(string))
			mgr.Save(&info)
			if tmpDay0 > day0 {
				day0 = tmpDay0
			}
		}
	}
	if day0 == 0 {
		mylog.Error(fmt.Sprintf("BuildOneDaily error : (%v)", code))
	}

	return SetDailyMa(code, day0, orm) // 开始匹配
}

// ma日均线

func SetDailyMa(code string, day0 int64, orm *mysqldb.MySqlDB) error {

	setDailyMa(code, day0, orm)
	setDailyMa(code, day0-(24*60*60), orm)
	return setDailyMa(code, day0-(2*(24*60*60)), orm)
}

func setDailyMa(code string, day0 int64, orm *mysqldb.MySqlDB) error {
	if day0 == 0 {
		return nil
	}

	info, err := model.SharesDailyTblMgr(orm.DB).FetchUniqueIndexByCode(code, day0)
	if err != nil {
		return err
	}
	if info.ID == 0 {
		return nil
	}

	// ma5
	subQuery := orm.Model(&model.SharesDailyTbl{}).Select("price").Where("code = ? and day0 <= ?", code, day0).Order("day0 desc")
	orm.Table("(?) as u", subQuery.Limit(5)).Select("avg(price)").Find(&info.Ma5)
	orm.Table("(?) as u", subQuery.Limit(10)).Select("avg(price)").Find(&info.Ma10)
	orm.Table("(?) as u", subQuery.Limit(20)).Select("avg(price)").Find(&info.Ma20)
	if info.Vol == 0 && info.Volume > 0 {
		var vol5 float64
		subQuery := orm.Model(&model.SharesDailyTbl{}).Select("volume").Where("code = ? and day0 < ?", code, day0).Order("day0 desc")
		orm.Table("(?) as u", subQuery.Limit(5)).Select("COALESCE(avg(volume),0) as avgg").Find(&vol5)
		if vol5 > 0 {
			info.Vol = info.Volume / vol5
		}
	}

	return orm.Save(&info).Error
}

// 计算行业当天主力流入
func countHYZLJLR(day0 int64) {
	orm := core.Dao.GetDBw()
	day00 := tools.GetUtcDay0(time.Now())
	if day0 == day00 { // 今天
		out := myhttp.OnGetJSON("https://push2.eastmoney.com/api/qt/clist/get?pn=1&pz=1000&np=1&fields=f12,f14,f62&fid=f62&fs=m:90", "") // 拉取分类
		// out = tools.ConvertToString(out, "gbk", "utf8")
		var info HyResp
		tools.JSONEncode(out, &info)
		if info.Rc == 0 {
			mgrhy := model.HyDailyTblMgr(orm.DB)

			for _, v := range info.Data.Diffs {
				if len(v.F12) == 0 {
					continue
				}
				out, _ := mgrhy.Reset().FetchUniqueIndexByHyCode(v.F12, day0)
				if out.ID == 0 {
					out.HyCode = v.F12
					out.HyName = v.F14
					out.Day0 = day0
					out.Day0Str = tools.GetDayStr(tools.UnixToTime(day0))
					out.CreatedAt = time.Now()
				}
				out.Zljlr = v.F62 * 0.0001
				mgrhy.Save(&out)
			}
		}
	}
}

// 北上
func maBS() {
	now := time.Now()
	day0 := tools.GetUtcDay0(now)
	week := tools.GetTimeWeek(day0)
	if week == 6 || week == 0 {
		return
	}

	out, _ := model.NoTradingDayTblMgr(core.Dao.GetDBr().DB).GetFromDay(datatypes.Date(now))
	if out.ID > 0 { // 休息时间
		return
	}

	info, _ := model.BsRapidlyTblMgr(core.Dao.GetDBw().DB).FetchUniqueByDay0(day0)
	if info.ID > 0 {
		return
	}

	var codes []string
	orm := core.Dao.GetDBr()
	model.SharesInfoTblMgr(orm.DB).Select("code").Find(&codes)
	for _, code := range codes {
		CountBSJLR(code, 1)
	}
	model.BsRapidlyTblMgr(core.Dao.GetDBw().DB).Save(&model.BsRapidlyTbl{
		Day0:      day0,
		CreatedAt: time.Now(),
	})

}

// 计算北上净流入(page = 0 重新计算)

func CountBSJLR(code string, page int) error {
	_, simpleCode := SplitCode(code)
	url := fmt.Sprintf("https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_MUTUAL_HOLDSTOCKNORTH_STA&columns=CLOSE_PRICE,TRADE_DATE,HOLD_SHARES,A_SHARES_RATIO&filter=(SECURITY_CODE=%v)", simpleCode)
	if page > 0 {
		url = fmt.Sprintf("https://datacenter-web.eastmoney.com/api/data/v1/get?pageSize=%v&reportName=RPT_MUTUAL_HOLDSTOCKNORTH_STA&columns=CLOSE_PRICE,TRADE_DATE,HOLD_SHARES,A_SHARES_RATIO&filter=(SECURITY_CODE=%v)", page, simpleCode)
	}
	out := myhttp.OnGetJSON(url, "") // 拉取分类
	mylog.Info(out)
	var info BSResp
	tools.JSONEncode(out, &info)
	if info.Code == 0 {
		orm := core.Dao.GetDBw()
		for _, v := range info.Result.Data { // 先更新数据
			err := model.SharesDailyTblMgr(orm.Where("`code` = ? AND `day0` = ?", code, v.TradeDate.Unix())).Updates(map[string]interface{}{
				model.SharesDailyTblColumns.Bscg:      v.Bscg,    // 北上持股数
				model.SharesDailyTblColumns.BsPercent: v.Percent, // 北上持股占比
			}).Error
			if err != nil {
				mylog.Error(err)
			}
		}

		// 计算净流入
		for _, v := range info.Result.Data {
			info, err := model.SharesDailyTblMgr(orm.Where("`code` = ? AND `day0` < ?", code, v.TradeDate.Unix()).Order("day0 desc").Limit(1)).Get()
			if err != nil {
				mylog.Error(err)
			}
			if info.ID > 0 && info.Bscg > 0 {
				bsjlr := (v.Bscg - info.Bscg) * v.ClosePrice
				model.SharesDailyTblMgr(orm.Where("`code` = ? AND `day0` = ?", code, v.TradeDate.Unix())).Update(model.SharesDailyTblColumns.Bsjlr, bsjlr)
			}
		}
	} else {
		if info.Code == 9201 {
			return message.GetError(message.NotFindError)
		}
	}
	return nil
}

// 获取分时原始数据

func GetMinute(code string) (*Minute, error) {
	// 开始获取
	out := SendGet(fmt.Sprintf("https://web.ifzq.gtimg.cn/appstock/app/minute/query?code=%v", code), "")
	out = tools.ConvertToString(out, "gbk", "utf8")
	out = strings.Replace(out, code, "info", -1)
	var info MinuteWWW
	err := json.Unmarshal([]byte(out), &info)
	if err != nil {
		mylog.Error(code)
		mylog.Error(out)
		mylog.Error(err)
		return nil, err
	}

	var sumPrice float64
	var sumVol int64
	var tmp Minute
	if info.Code == 0 {
		if len(info.Data.Info.QT.Info) > 4 {
			tmp.PrePrice = stringToFloat(info.Data.Info.QT.Info[4])
		}
		// index := len(info.Data.Info.Data.Data) - 1
		for _, v := range info.Data.Info.Data.Data {
			out := strings.Split(v, " ")
			if len(out) == 4 {
				minfo := MinuteInfo{
					Time:  out[0],
					Price: stringToFloat(out[1]),
					Vol:   int64(StringToInt(out[2])),
				}
				minfo.Vol = minfo.Vol - sumVol
				sumVol += minfo.Vol // 总成交量
				sumPrice += minfo.Price * float64(minfo.Vol)
				if sumVol > 0 {
					minfo.Ave = Decimal(sumPrice / float64(sumVol))
				}
				tmp.List = append(tmp.List, minfo)
			}
		}
	}

	return &tmp, nil
}

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

// 获取分时原始数据

func GetDayliy(code string) ([][]interface{}, error) {
	out := SendGet(fmt.Sprintf("https://web.ifzq.gtimg.cn/appstock/app/fqkline/get?param=%v,day,,,320,qfq", code), "")
	out = tools.ConvertToString(out, "gbk", "utf8")
	out = strings.Replace(out, code, "info", -1)
	out = strings.Replace(out, "qfqday", "day", -1)

	var info DailyWWW
	err := json.Unmarshal([]byte(out), &info)
	if err != nil {
		mylog.Error(err)
		return nil, err
	}

	if info.Code != 0 { // 不成功
		return nil, message.GetError(message.InValidOp)
	}
	var output [][]interface{}
	for _, tmp := range info.Data.Info.Day {
		if len(tmp) == 6 {
			output = append(output, []interface{}{tmp[0], stringToFloat(tmp[1].(string)), stringToFloat(tmp[2].(string)), stringToFloat(tmp[4].(string)), stringToFloat(tmp[3].(string)), int64(stringToFloat(tmp[5].(string)))})
		}
	}

	return output, nil
}

func GetMacd(code string) ([]Macd, error) {
	kdata, err := GetDayliy(code)
	if err != nil {
		return nil, err
	}
	return calcMACD(12, 26, 9, kdata, 2), nil
}

// 龙虎榜
func maLhb() {
	now := time.Now()
	day0 := tools.GetUtcDay0(now) // - int64(24*60*60)
	week := tools.GetTimeWeek(day0)
	if week == 6 || week == 0 {
		return
	}

	out, _ := model.NoTradingDayTblMgr(core.Dao.GetDBr().DB).GetFromDay(datatypes.Date(now))
	if out.ID > 0 { // 休息时间
		return
	}

	DayDailyLHB(day0)

	// 更新龙虎榜数据
	orm := core.Dao.GetDBw()

	// 更新新的
	list1, _ := model.LhbDailyTblMgr(orm.DB).GetFromDay0(day0)
	for _, v := range list1 {
		if v.JgJlr > 0 && v.Jlr > 0 { // 机构净流入才进
			info, _ := model.AnalyLhbTblMgr(orm.DB).FetchUniqueIndexByCode(v.Code, day0)
			info.Code = v.Code
			info.Name = v.Name
			info.Day0 = v.Day0
			info.DayStr = v.DayStr
			info.Price = v.Price
			info.JgJlr = v.JgJlr
			info.Jlr = v.Jlr
			info.CreatedAt = v.CreatedAt
			model.AnalyLhbTblMgr(orm.DB).Save(&info)
		}
	}

	// 删除老的
	list, _ := model.AnalyLhbTblMgr(orm.DB).Gets()
	for _, v := range list {
		//sharesInfo, err := GetShare(v.Code, true)
		dailys, _ := model.SharesDailyTblMgr(orm.Where("day0 < ?", day0).Order("day0 desc").Limit(1)).GetFromCode(v.Code)
		if len(dailys) == 0 {
			continue
		}
		if dailys[0].Price < dailys[0].Ma10 { // 未站上10日线
			model.AnalyLhbTblMgr(orm.DB).Delete(&v) // 删除
		} else {
			sharesInfo := SearchDetail(v.Code)
			v.NewPrice = sharesInfo.Price
			v.TurnoverRate = sharesInfo.TurnoverRate
			model.AnalyLhbTblMgr(orm.DB).Save(&v)
		}
	}

}
