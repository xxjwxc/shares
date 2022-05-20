package event

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"shares/internal/api"
	"shares/internal/core"
	"shares/internal/model"
	"strconv"
	"time"

	"github.com/mozillazg/go-pinyin"
	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/tools"
	"gorm.io/datatypes"
)

var _pinyin = pinyin.NewArgs()

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
	return GetColor(user.Info.Rg, float64(percent))
}

func FirstLetterOfPinYin(name string) string {
	result := pinyin.Pinyin(name, _pinyin)
	if len(result) == 0 {
		return ""
	}
	var out string
	for _, v := range result {
		out += fmt.Sprintf("%s", string(v[0][0]))
	}
	return out
}

func IsWorkDay() bool {
	now := time.Now()
	day0 := tools.GetUtcDay0(now) // - int64(24*60*60)
	week := tools.GetTimeWeek(day0)
	if week == 6 || week == 0 {
		return false
	}

	out, _ := model.NoTradingDayTblMgr(core.Dao.GetDBr().DB).GetFromDay(datatypes.Date(now))

	return out.ID <= 0
}
