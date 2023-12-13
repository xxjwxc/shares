package routers

import (
	"github.com/xxjwxc/ginrpc"
)

func init() {
	ginrpc.SetVersion(1702455234)
	ginrpc.AddGenOne("Weixin.GetQrcode", "weixin.get_qrcode", []string{"post"}, []ginrpc.GenThirdParty{}, `获取微信二维码`)
	ginrpc.AddGenOne("Weixin.GetUserInfo", "weixin.get_user_info", []string{"post"}, []ginrpc.GenThirdParty{}, ``)
	ginrpc.AddGenOne("Weixin.Oauth", "weixin.oauth", []string{"post"}, []ginrpc.GenThirdParty{}, `微信授权获取登录信息`)
	ginrpc.AddGenOne("Weixin.ReLogin", "weixin.re_login", []string{"post"}, []ginrpc.GenThirdParty{}, `是否要重新登录`)
	ginrpc.AddGenOne("Weixin.UpdateUserInfo", "weixin.update_user_info", []string{"post"}, []ginrpc.GenThirdParty{}, `更新用户信息`)
	ginrpc.AddGenOne("Weixin.UpsetUserInfo", "weixin.upset_user_info", []string{"post"}, []ginrpc.GenThirdParty{}, `更新用户信息`)
	ginrpc.AddGenOne("Shares.AddGroup", "shares.add_group", []string{"post"}, []ginrpc.GenThirdParty{}, `添加一个组织`)
	ginrpc.AddGenOne("Shares.AddMyCode", "shares.add_my_code", []string{"post"}, []ginrpc.GenThirdParty{}, `给自己添加一个监听`)
	ginrpc.AddGenOne("Shares.Dayliy", "shares.dayliy", []string{"post"}, []ginrpc.GenThirdParty{}, ``)
	ginrpc.AddGenOne("Shares.DeleteMyCode", "shares.delete_my_code", []string{"post"}, []ginrpc.GenThirdParty{}, `删除一个监听`)
	ginrpc.AddGenOne("Shares.GetAllCodeName", "shares.get_all_code_name", []string{"post"}, []ginrpc.GenThirdParty{}, `获取所有中文`)
	ginrpc.AddGenOne("Shares.GetGroup", "shares.get_group", []string{"post"}, []ginrpc.GenThirdParty{}, `获取分组`)
	ginrpc.AddGenOne("Shares.GetMsg", "shares.get_msg", []string{"post"}, []ginrpc.GenThirdParty{}, `获取消息`)
	ginrpc.AddGenOne("Shares.GetMyCode", "shares.get_my_code", []string{"post"}, []ginrpc.GenThirdParty{}, ``)
	ginrpc.AddGenOne("Shares.GetMyGroup", "shares.get_my_group", []string{"post"}, []ginrpc.GenThirdParty{}, `获取我的组织`)
	ginrpc.AddGenOne("Shares.Gets", "shares.gets", []string{"post"}, []ginrpc.GenThirdParty{}, `精确查找代码`)
	ginrpc.AddGenOne("Shares.HaveNewMsg", "shares.have_new_msg", []string{"post"}, []ginrpc.GenThirdParty{}, ``)
	ginrpc.AddGenOne("Shares.Minute", "shares.minute", []string{"post"}, []ginrpc.GenThirdParty{}, `获取分时图`)
	ginrpc.AddGenOne("Shares.Search", "shares.search", []string{"post"}, []ginrpc.GenThirdParty{}, `搜索`)
	ginrpc.AddGenOne("Shares.UpsetGroupCode", "shares.upset_group_code", []string{"post"}, []ginrpc.GenThirdParty{}, ``)
	ginrpc.AddGenOne("Analy.AnalyCode", "analy.analy_code", []string{"post"}, []ginrpc.GenThirdParty{}, ``)
}
