package routers

import (
	"github.com/xxjwxc/ginrpc"
)

func init() {
	ginrpc.SetVersion(1640864803)
	ginrpc.AddGenOne("Weixin.GetQrcode", "weixin.get_qrcode", []string{"post"})
	ginrpc.AddGenOne("Weixin.GetUserInfo", "weixin.get_user_info", []string{"post"})
	ginrpc.AddGenOne("Weixin.Oauth", "weixin.oauth", []string{"post"})
	ginrpc.AddGenOne("Weixin.ReLogin", "weixin.re_login", []string{"post"})
	ginrpc.AddGenOne("Weixin.UpdateUserInfo", "weixin.update_user_info", []string{"post"})
	ginrpc.AddGenOne("Weixin.UpsetUserInfo", "weixin.upset_user_info", []string{"post"})
	ginrpc.AddGenOne("Shares.AddGroup", "shares.add_group", []string{"post"})
	ginrpc.AddGenOne("Shares.AddMyCode", "shares.add_my_code", []string{"post"})
	ginrpc.AddGenOne("Shares.Dayliy", "shares.dayliy", []string{"post"})
	ginrpc.AddGenOne("Shares.DeleteMyCode", "shares.delete_my_code", []string{"post"})
	ginrpc.AddGenOne("Shares.GetAllCodeName", "shares.get_all_code_name", []string{"post"})
	ginrpc.AddGenOne("Shares.GetGroup", "shares.get_group", []string{"post"})
	ginrpc.AddGenOne("Shares.GetMsg", "shares.get_msg", []string{"post"})
	ginrpc.AddGenOne("Shares.GetMyCode", "shares.get_my_code", []string{"post"})
	ginrpc.AddGenOne("Shares.GetMyGroup", "shares.get_my_group", []string{"post"})
	ginrpc.AddGenOne("Shares.Gets", "shares.gets", []string{"post"})
	ginrpc.AddGenOne("Shares.HaveNewMsg", "shares.have_new_msg", []string{"post"})
	ginrpc.AddGenOne("Shares.Minute", "shares.minute", []string{"post"})
	ginrpc.AddGenOne("Shares.Search", "shares.search", []string{"post"})
	ginrpc.AddGenOne("Shares.UpsetGroupCode", "shares.upset_group_code", []string{"post"})
	ginrpc.AddGenOne("Analy.AnalyCode", "analy.analy_code", []string{"post"})
}
