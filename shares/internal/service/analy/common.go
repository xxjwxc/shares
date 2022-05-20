package analy

import (
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"shares/internal/api"
	"shares/internal/core"
	"shares/internal/model"
	"shares/internal/service/event"
	"strconv"

	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/tools"
	"gorm.io/datatypes"
)

func SaveLog(log string) {
	info := model.LogTbl{
		Desc: log,
	}
	orm := core.Dao.GetDBr()
	model.LogTblMgr(orm.DB).Save(&info)
}

//SendGet 发送get 请求 返回对象
func SendGet(url, params string) string {
	//解析这个 URL 并确保解析没有出错。
	var urls = url
	if len(params) > 0 {
		urls += "?" + params
	}
	resp, err := http.Get(urls)
	if err != nil {
		mylog.Error(err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		mylog.Error(err)
		return ""
	}

	return string(body)
}

func StringToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

func SplitCode(code string) (ext, simplecode string) {
	if len(code) < 2 {
		return
	}
	ext = code[:2]
	simplecode = code[2:]
	return
}

func BuildCode(tag int, simplecode string) (code string) {
	switch tag {
	case 0:
		return "sz" + simplecode
	case 1:
		return "sh" + simplecode
	default:
		mylog.Errorf("tag ,simplecode ,error,%v,%v", tag, simplecode)
	}
	return
}

func Abs(in int) int {
	if in < 0 {
		return -in
	}
	return in
}

func getOpenIdColor(openid string, percent int) string {
	user, err := api.GetUserFromOpenID(openid)
	if err != nil {
		user = &api.UserInfo{
			AccessToken: openid,
			Info:        model.WxUserinfo{},
		}
	}
	return event.GetColor(user.Info.Rg, float64(percent))
}

func stringToFloat(src string) float64 {
	v, _ := strconv.ParseFloat(src, 64)
	return v
}

func getMin(arrays ...float64) (f float64) {
	if len(arrays) > 0 {
		f = arrays[0]
	}
	for _, v := range arrays {
		if f > v {
			f = v
		}
	}
	return
}

func getMax(arrays ...float64) (f float64) {
	if len(arrays) > 0 {
		f = arrays[0]
	}
	for _, v := range arrays {
		if f < v {
			f = v
		}
	}
	return
}

func getSamplePrice(price float64) string {
	tag := ""
	if price < 0 {
		tag = "-"
		price = math.Abs(price)
	}

	if price < 10000 { // 一万以下，直接输出数字
		return fmt.Sprintf("%v%0.2f", tag, price)
	}

	if price < 100000000 { // 一亿以下，直接输出万级别,保留2位小数
		return fmt.Sprintf("%v%0.2f万", tag, price*0.0001)
	}

	// 到亿
	return fmt.Sprintf("%v%0.2f亿", tag, price*0.00000001)

}

func getNextTimeDayStr(day0 int64) string {
	for {
		day0 += 24 * 60 * 60
		if IsWorkDay(day0) {
			return tools.FormatTime(tools.UnixToTime(day0), "01-02")
		}
	}
}

func IsWorkDay(day0 int64) bool {
	week := tools.GetTimeWeek(day0)
	if week == 6 || week == 0 {
		return false
	}

	out, _ := model.NoTradingDayTblMgr(core.Dao.GetDBr().DB).GetFromDay(datatypes.Date(tools.UnixToTime(day0)))
	return out.ID <= 0
}

// getBackWorkDay 获取上一个工作日
func getBackWorkDay(day0 int64) int64 {
	for {
		day0 -= 24 * 60 * 60
		if IsWorkDay(day0) {
			return day0
		}
	}
}
