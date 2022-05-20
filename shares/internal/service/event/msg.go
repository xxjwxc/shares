package event

import (
	"fmt"
	"shares/internal/core"
	"shares/internal/model"
	"shares/internal/service/weixin"
	proto "shares/rpc/shares"
	"time"

	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/tools"
	wx "github.com/xxjwxc/public/weixin"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Msg struct {
	openid  string
	name    string
	code    string
	key     string
	price   float64
	percent float64
	desc    string
	tag     string //  标记(min,daily)
	color   string // 消息颜色
}

// 快速涨跌提醒
func surgeSlump(code string, old, new, offsetPercent float64, offsetSecond int64) error {
	tag := 1
	if !(new > old) {
		tag = -1
	}

	now := time.Now()
	day0 := tools.GetUtcDay0(now)
	offset := now.Unix() - day0
	if offset >= (12 * 60 * 60) { // 下午时间
		tag *= 2
	}

	// 提醒过，就不提醒了
	orm := core.Dao.GetDBw()
	info, _ := model.MsgRapidlyTblMgr(orm.DB).FetchUniqueIndexByCode(code, tag, datatypes.Date(time.Now()))
	if info.ID > 0 { // 已经提示过
		return nil
	}

	sharesInfo, err := GetShare(code, true)
	if err != nil {
		return nil
	}

	key := "急涨提醒"
	var desc string
	var condition model.Condition
	condition.And(model.SharesWatchTblColumns.Code, "=", code)
	condition.And(model.SharesWatchTblColumns.Vaild, "=", true)
	if new > old {
		condition.And(model.SharesWatchTblColumns.Surge, "=", true)
		desc = fmt.Sprintf("%v价格从(%v)元涨至(%v)元", getTmstr(offsetSecond), old, new)
	} else {
		condition.And(model.SharesWatchTblColumns.Slump, "=", true)
		key = "急跌提醒"
		desc = fmt.Sprintf("%v价格从(%v)元跌至(%v)元", getTmstr(offsetSecond), old, new)
	}
	where, args := condition.Get()

	infos, err := model.SharesWatchTblMgr(orm.Where(where, args...)).Gets()
	if err != nil {
		return err
	}

	// todo: 发送信息
	var msgs []*Msg
	for _, v := range infos {
		fmt.Println(v)
		msgs = append(msgs, &Msg{
			openid:  v.OpenID,
			name:    fmt.Sprintf("%v(%v)", sharesInfo.Name, code),
			code:    code,
			key:     key,
			price:   sharesInfo.Price,
			percent: sharesInfo.Percent,
			desc:    desc,
			tag:     "min",
			color:   getOpenIdColor(v.OpenID, tag),
		})
	}
	if len(msgs) > 0 {
		err = sendWXmsg(msgs)
		if err != nil {
			return err
		}
		// todo:保存到已经提醒
		model.MsgRapidlyTblMgr(orm.DB).Save(&model.MsgRapidlyTbl{
			Code:          code,
			Tag:           tag,
			Key:           key,
			Desc:          desc,
			Old:           old,
			New:           new,
			Percent:       sharesInfo.Percent,
			OffsetPercent: offsetPercent,
			Day:           datatypes.Date(time.Now()),
			CreatedAt:     time.Now(),
		})
	}

	return nil
}

// 估价涨跌+百分比
func upDown(code string, price, percent float64) error {
	sharesInfo, err := GetShare(code, true)
	if err != nil {
		return nil
	}
	/////////////////////////////////////////////////////////////////////////////////////////////

	var ids []int
	var msgs []*Msg
	orm := core.Dao.GetDBw()
	infoup, _ := model.SharesWatchTblMgr(orm.Where("up < ? and up > 0 and up_vaild = ? and vaild = ?", price, true, true)).GetFromCode(code)
	for _, v := range infoup {
		msgs = append(msgs, &Msg{
			openid:  v.OpenID,
			name:    fmt.Sprintf("%v(%v)", sharesInfo.Name, code),
			code:    code,
			key:     "上涨提醒",
			price:   price,
			percent: percent,
			desc:    fmt.Sprintf("最新价(%v)元,高于设置的(%v)元", price, v.Up),
			tag:     "min",
			color:   getOpenIdColor(v.OpenID, 1),
		})
		ids = append(ids, v.ID)
	}
	if len(msgs) > 0 {
		err = sendWXmsg(msgs)
		if err != nil {
			return err
		}
		model.SharesWatchTblMgr(orm.Where("id in ?", ids)).Updates(map[string]interface{}{"up_vaild": false, "up": 0}) // 设置已更新
	}

	/////////////////////////////////////////////////////////////////////////////////////////////

	msgs = []*Msg{}
	ids = []int{}
	infodown, _ := model.SharesWatchTblMgr(orm.Where("down > ?  and down_vaild = ? and vaild = ?", price, true, true)).GetFromCode(code)
	for _, v := range infodown {
		msgs = append(msgs, &Msg{
			openid:  v.OpenID,
			name:    fmt.Sprintf("%v(%v)", sharesInfo.Name, code),
			code:    code,
			key:     "下跌提醒",
			price:   price,
			percent: percent,
			desc:    fmt.Sprintf("最新价(%v)元,低于设置的(%v)元", price, v.Down),
			tag:     "min",
			color:   getOpenIdColor(v.OpenID, -1),
		})
		ids = append(ids, v.ID)
	}
	if len(msgs) > 0 {
		err = sendWXmsg(msgs)
		if err != nil {
			return err
		}
		model.SharesWatchTblMgr(orm.Where("id in ?", ids)).Updates(map[string]interface{}{"down_vaild": false, "down": 0}) // 设置已更新
	}

	/////////////////////////////////////////////////////////////////////////////////////////////

	msgs = []*Msg{}
	infoupto, _ := model.SharesWatchTblMgr(orm.Where("up_percent_to < ? and up_percent_to > 0 and vaild = ?", price, true)).GetFromCode(code)
	for _, v := range infoupto {
		msgs = append(msgs, &Msg{
			openid:  v.OpenID,
			name:    fmt.Sprintf("%v(%v)", sharesInfo.Name, code),
			code:    code,
			key:     "上涨提醒",
			price:   price,
			percent: percent,
			desc:    fmt.Sprintf("最新价(%v)元,高于设置的(%%%v)", price, v.UpPercent),
			tag:     "min",
			color:   getOpenIdColor(v.OpenID, 1),
		})
		v.UpPercentTo = price + ((price * v.UpPercent) * 0.01)
	}
	if len(msgs) > 0 {
		err = sendWXmsg(msgs)
		if err != nil {
			return err
		}
		for _, v := range infoupto {
			model.SharesWatchTblMgr(orm.DB).Select("up_percent_to").Updates(v) // 设置已更新
		}

	}
	msgs = []*Msg{}
	/////////////////////////////////////////////////////////////////////////////////////////////
	infodownto, _ := model.SharesWatchTblMgr(orm.Where("down_percent_to > ? and vaild = ?", price, true)).GetFromCode(code)
	for _, v := range infodownto {
		msgs = append(msgs, &Msg{
			openid:  v.OpenID,
			name:    fmt.Sprintf("%v(%v)", sharesInfo.Name, code),
			code:    code,
			key:     "下跌提醒",
			price:   price,
			percent: percent,
			desc:    fmt.Sprintf("最新价(%v)元,低于设置的(%%%v)", price, v.DownPercent),
			tag:     "min",
			color:   getOpenIdColor(v.OpenID, -1),
		})
		v.DownPercentTo = price - ((price * v.DownPercent) * 0.01)
	}
	if len(msgs) > 0 {
		err = sendWXmsg(msgs)
		if err != nil {
			return err
		}
		for _, v := range infodownto {
			model.SharesWatchTblMgr(orm.DB).Select("down_percent_to").Updates(v) // 设置已更新
		}
	}
	// msgs = []*Msg{}

	return nil
}

// 日均线实时提醒(跌破5日,10日线,站上20日线)
func maRealTime(code string, price, percent float64) error {
	sharesInfo, err := GetShare(code, true)
	if err != nil {
		return nil
	}
	day0 := tools.GetUtcDay0(time.Now())

	orm := core.Dao.GetDBw()
	list, err := model.SharesDailyTblMgr(orm.Where("day0 < ?", day0).Order("day0 desc").Limit(1)).GetFromCode(code)
	if err != nil {
		return nil
	}
	if len(list) < 1 { // 不需要
		return nil
	}
	if list[0].Price > price { // 下跌趋势
		if (list[0].Ma20 < list[0].Price) && (list[0].Ma20 > price) { // 跌破20日线
			maSend(sharesInfo, -20, list[0].Ma20, price, orm.DB)
		} else if (list[0].Ma10 < list[0].Price) && (list[0].Ma10 > price) && (price > list[0].Ma20) { // 跌破10日线
			maSend(sharesInfo, -10, list[0].Ma10, price, orm.DB)
		} else if (list[0].Ma5 < list[0].Price) && (list[0].Ma5 > price) && (price > list[0].Ma10) { // 跌破5日线
			maSend(sharesInfo, -5, list[0].Ma5, price, orm.DB)
		}
	}

	// 上涨趋势
	if list[0].Price < price {
		if (list[0].Ma20 > list[0].Price) && (list[0].Ma20 < price) { // 站上20日线
			maSend(sharesInfo, 20, list[0].Ma20, price, orm.DB)
		} else if (list[0].Ma10 > list[0].Price) && (list[0].Ma10 < price) { // 站上10日线
			maSend(sharesInfo, 10, list[0].Ma10, price, orm.DB)
		}
	}

	return nil
}

func maSend(sharesInfo *proto.SharesInfo, maNum int, old, new float64, db *gorm.DB) error {
	absMaNum := Abs(maNum)
	tt := "站上"
	if maNum < 0 {
		tt = "跌破"
	}
	info, _ := model.MsgRapidlyTblMgr(db).FetchUniqueIndexByCode(sharesInfo.Code, maNum, datatypes.Date(time.Now()))
	if info.ID == 0 { // 未提醒过
		key := "日线提醒"
		desc := fmt.Sprintf("最新价(%v)元,%v%v日线(%v)元", new, tt, absMaNum, old)
		if absMaNum == 20 { // 20日线提醒
			db = db.Where("(kdj = ? or kdj20 = ? ) and vaild = ?", true, true, true)
		} else {
			db = db.Where("kdj = ? and vaild = ?", true, true)
		}
		// todo 发消息
		var msgs []*Msg
		infomai, _ := model.SharesWatchTblMgr(db).GetFromCode(sharesInfo.Code)
		for _, v := range infomai {
			msgs = append(msgs, &Msg{
				openid:  v.OpenID,
				name:    fmt.Sprintf("%v(%v)", sharesInfo.Name, sharesInfo.Code),
				code:    sharesInfo.Code,
				key:     key,
				price:   new,
				percent: sharesInfo.Percent,
				desc:    desc,
				tag:     "daily",
				color:   getOpenIdColor(v.OpenID, maNum),
			})
		}
		if len(msgs) > 0 {
			err := sendWXmsg(msgs)
			if err != nil {
				return err
			}
		}

		// todo:保存到已经提醒
		model.MsgRapidlyTblMgr(db).Save(&model.MsgRapidlyTbl{
			Code:          sharesInfo.Code,
			Tag:           maNum,
			Key:           key,
			Desc:          desc,
			Old:           old,
			New:           new,
			Percent:       sharesInfo.Percent,
			OffsetPercent: ((old - new) / new) * 100,
			Day:           datatypes.Date(time.Now()),
			CreatedAt:     time.Now(),
		})
		return nil
	}

	return nil
}

func sendWXmsg(msgs []*Msg) error {
	if len(msgs) == 0 {
		return nil
	}
	orm := core.Dao.GetDBw()
	wxMsg := make([]wx.TempWebMsg, 0, len(msgs))
	list := make([]model.MsgTbl, 0, len(msgs))
	for _, v := range msgs {
		list = append(list, model.MsgTbl{
			OpenID:    v.openid,
			Code:      v.code,
			Key:       v.key,
			Desc:      v.desc,
			Price:     v.price,
			Percent:   v.percent,
			Day:       datatypes.Date(time.Now()),
			CreatedAt: time.Now(),
		})

		data := make(map[string]map[string]string)
		mp := make(map[string]string)
		mp["value"] = fmt.Sprintf("股票名称:%v", v.name)
		mp["color"] = v.color

		data["first"] = mp

		mp = make(map[string]string)
		mp["value"] = v.key
		mp["color"] = v.color
		data["keyword1"] = mp

		mp = make(map[string]string)
		mp["value"] = v.desc
		mp["color"] = v.color
		data["keyword2"] = mp

		// mp = make(map[string]string)
		// mp["value"] = "(无)"
		// data["keyword3"] = mp

		// mp = make(map[string]string)
		// mp["value"] = "(无)"
		// data["keyword4"] = mp

		mp = make(map[string]string)
		mp["value"] = tools.GetTimeStr(time.Now())
		data["keyword5"] = mp

		// 行业
		hy, _ := model.SharesInfoTblMgr(orm.DB).GetFromCode(v.code)
		if hy.ID > 0 && len(hy.HyName) > 0 {
			mp = make(map[string]string)
			mp["value"] = fmt.Sprintf("行业板块:%v", hy.HyName)
			data["remark"] = mp
		}

		wxMsg = append(wxMsg, wx.TempWebMsg{
			Touser:     v.openid,
			TemplateID: "t1GpGWEt4C8FAezjDqxupp0qDf--asgHfDoOq8TaECk",
			Page:       fmt.Sprintf("https://hospital.xxjwxc.cn/webshares/#/pages/add/add?scode=%v&tag=%v", v.code, v.tag),
			Data:       data,
		})
	}

	weixin.SendMsg(wxMsg)
	err := model.MsgTblMgr(orm.DB).Save(&list).Error
	if err != nil {
		mylog.Error(err)
	}

	return err
}

func getTmstr(offsetSecond int64) string {
	minute := offsetSecond / 60
	second := offsetSecond % 60
	tmp := ""
	if minute > 0 {
		tmp += fmt.Sprintf("%v分", minute)
	}
	if second > 0 {
		tmp += fmt.Sprintf("%v秒", second)
	}

	if len(tmp) > 0 {
		tmp += ","
	}
	return tmp
}

// 估价涨跌+百分比
func UPPPP() {
	// maRealTime("sz002242", 23, 5)
	// maRealTime("sh600309", 300, 5)
	// GetMinute("sh600309")
	// upDown("sh603259", 129, -10)
	// var msgs []*Msg
	// msgs = append(msgs, &Msg{
	// 	openid:  "oxFCP6hcaZTReezFSs80ZY6qNAv8",
	// 	name:    fmt.Sprintf("测试测试"),
	// 	code:    "sh600027",
	// 	key:     "上涨提醒",
	// 	price:   123,
	// 	percent: 222,
	// 	desc:    fmt.Sprintf("测试测试"),
	// 	tag:     "min",
	// 	color:   getOpenIdColor("oxFCP6hcaZTReezFSs80ZY6qNAv8", 1),
	// })
	// sendWXmsg(msgs)
	// fmt.Println(upDown("sh603259", 131, -5))
}
