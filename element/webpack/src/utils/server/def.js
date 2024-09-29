import router from '@/router';
import axios from 'axios';
import { ElMessageBox } from 'element-plus'
import { useRoute } from 'vue-router';

const server = {}

server.Host = "/shares/api/v1";
// server.Host = "https://www.xxxxxx.cn/shares/api/v1";
server.isDebug = false;
server.Appid = "wx6314e23378f3b31a";

server.SessionID = ""; // 用户sessionid
server.OpenID = ""; // 用户openid
server.loginCount = 10;
server.UserInfo = {};
// server.IsLogin = false;

server.GetUrl = (router) => {
	return server.Host + router;
};

/*   
router：相对路由
data：请求参数
返回：成功或者失败
success:成功 调用
fail:自定义失败调用
*/
server.Post = async (router, data) => {
	return new Promise((resolve, reject) => {
		//显示加载动画
		// uni.showLoading({
		// 	title: '请稍后..',
		// 	mask: true,

		// })
        axios.post(server.GetUrl(router), data,{
            headers: {
            'Content-Type': 'application/json',
            'user-token': server.OpenID || localStorage.getItem('user-token')}
            }).then((res) => {// 成功
				if (res.status == 200) { // 判断http返回状态 
					if (res.data.state) { // 正常返回 
						resolve({
							"success": res.data.data,
							"fail": null
						})
					} else { // 错误返回
						let code = res.data.code
						switch (code) {
							// case 1018: // 用户不存在
							// 	server.SessionID = "";
							// 	server.OpenID = "";
							// 	server.SetCookie("user-token",server.SessionID)
							default:
								resolve({
									"success": null,
									"fail": res.data
								});
						}
					}
				} else {
					resolve({
						"success": null,
						"fail": {
							error: res.data
						}
					});
				}
			}).catch((err) => {// 失败
                reject(err);
                ElMessageBox.alert('服务器暂时无法访问，请稍后再试！', '错误', {
                    confirmButtonText: 'OK',
                })
            })
	})
};

server.Get = async (router, data) => {
	return new Promise((resolve) => {
        axios.get(server.GetUrl(router), data).then((res) => {// 成功
            resolve({
                "success": res,
                "fail": null
            })
        }).catch((res) => {// 失败
            resolve({
                "success": null,
                "fail": res
            });
            //  //失败的时候调用reject
            // reject('error message')
        })
	})
};

server.GetEx = async (router, data) => {
	return new Promise((resolve) => {
        axios.get(router, data).then((res) => {// 成功
            resolve({
                "success": res,
                "fail": null
            })
        }).catch((res) => {// 失败
            resolve({
                "success": null,
                "fail": res
            });
            //  //失败的时候调用reject
            // reject('error message')
        })
	})
};


server.GetReturn = (resp) => {
	if (resp.success) { // 设置session
		return resp.success
	}
	if (resp.fail) { // 失败打印输出
		switch (resp.fail.code) {
			case undefined:
            ElMessageBox.alert('服务器暂时维护中，请稍后再试！', '错误', {
                confirmButtonText: 'OK',
            })
            break
			case 2:
				break
			case 1018:
				server.SessionID = "";
				server.OpenID = "";
				server.SetCookie("user-token",server.OpenID)
				// server.WxLogin(false)
				break
			default:
                ElMessageBox.alert(resp.fail.error, '错误', {
                    confirmButtonText: 'OK',
                })
		}
	}
	return null;
};

server.SetCookie = async (key, value) => {
	if (value?.length == 0){
		localStorage.removeItem(key)
		return
	}
	localStorage.setItem(key,value)
};

//////////////////////////////////////////////////
// login
//////////////////////////////////////////////////
// async function wxLogin(jscode) {
// 	let resp = await server.Post("/weixin.oauth", {
// 		"jscode": jscode
// 	});

// 	if (resp.success) { // 设置session
// 		this.SessionID = resp.success.session_id;
// 		this.OpnID = resp.success.open_id;
// 	}
// 	if (resp.fail) { // 失败打印输出
//         ElMessageBox.alert(resp.fail.error, '错误', {
//             confirmButtonText: 'OK',
//         })
// 	}
// };

server.WxLogin = async function() {
	if (this.isDebug){
		return
	}
	if (!this.IsLogin()) {
        let deviceid = this.getUrlCode('deviceid');
        if (deviceid == null) {
            router.replace('/login');
        } else {
            router.replace('/login?deviceid=' + deviceid);
        }

	}
};

server.IsLogin = function() {
	let sessonid = this.OpenID || localStorage.getItem('user-token')
	if (sessonid == "undefined" || sessonid == null || sessonid == null || sessonid == "") {
		return false;
	}
	
	return true;
};

server.getUrlCode = function(name) {
    const route = useRoute();
    return route?.query[name];
};

//请求微信接口，用来获取code
server.getWeChatCode = function(isBase) {
	let scope = "snsapi_base";
	if (!isBase){
		scope = "snsapi_userinfo";
	}
	let local = encodeURIComponent(window.location.href); //获取当前页面地址作为回调地址
	//通过微信官方接口获取code之后，会重新刷新设置的回调地址【redirect_uri】
	window.location.href =
		"https://open.weixin.qq.com/connect/oauth2/authorize?appid=" +
		this.Appid +
		"&redirect_uri=" +
		local +
		"&response_type=code&scope="+scope+"&state=1#wechat_redirect";
}

// 微信授权获取登录信息
server.Oauth = async function(jscode,isUpdateUser){
	if (this.loginCount > 0){
		this.loginCount = this.loginCount -1
		let resp = await this.Post("/weixin.oauth",{jscode:jscode,isUpdateUser:isUpdateUser});
		if (resp.success){// 设置session
			this.SessionID = resp.success.sessionId;
			this.OpenID  = resp.success.openId;
			this.SetCookie("user-token",this.OpenID)
		}
			
		let out =  this.GetReturn(resp);
		if (out){
			localStorage.setItem('vip_level', out.level);
		}else {
			localStorage.removeItem('vip_level')
		}
		return out
	}

}
//////////////////////////////////////////////////

export default server;
