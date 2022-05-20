package analy

import (
	"fmt"
	"shares/internal/service/weixin"

	wx "github.com/xxjwxc/public/weixin"
)

func initMenu() { // 自定义菜单
	var menu wx.WxMenu
	menu.Button = append(menu.Button, wx.WxMenuButton{
		Type: "view",
		Name: "首页",
		Url:  "https://hospital.xxjwxc.cn/webshares",
	})
	menu.Button = append(menu.Button, wx.WxMenuButton{
		Name: "操作",
		SubButton: []wx.SubButton{{
			Type: "view",
			Name: "录入",
			Url:  "https://hospital.xxjwxc.cn/webshares/#/pages/group/group",
		}, {
			Type: "view",
			Name: "搜索",
			Url:  "https://hospital.xxjwxc.cn/webshares/#/pages/add/add",
		}},
	})
	menu.Button = append(menu.Button, wx.WxMenuButton{
		Type: "view",
		Name: "个人中心",
		Url:  "https://hospital.xxjwxc.cn/webshares/#/pages/mine/mine",
	})
	err := weixin.InitMenu(menu)

	// err := weixin.DeleteMenu()
	fmt.Println(err)
}
