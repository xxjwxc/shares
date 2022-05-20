/* 用户信息相关 */
import Server from './def';

const user = {}

user.Clear = clear => {
	Server.SessionID = "";
	Server.OpenID = "";
}
	
user.IsLogin = islogin => {
	console.log(Server.SessionID);

	let sessonid = Server.SessionID
	if (sessonid == "undefined" || sessonid == null || sessonid == null || sessonid == "") {
		return false;
	}

	return true;
};

// 获取用户信息

async function wxLogin(jscode) {
	console.log(jscode)
	let resp = await Server.Post("/weixin.oauth", {
		"jscode": jscode
	});

	if (resp.success) { // 设置session
		Server.SessionID = resp.success.session_id;
		Server.OpnID = resp.success.open_id;
		console.log("session:" + Server.SessionID);
	}
	if (resp.fail) { // 失败打印输出
		uni.showToast({
			title: resp.fail.error,
			icon: 'none'
		})
	}
};

user.WxLogin = async function(isBase) {
	if (!this.IsLogin()) {
		// #ifdef MP-WEIXIN
		await uni.login({
			success: res => {
				wxLogin(res.code);
			}
		});
		// #endif

		// #ifdef H5
		let code = this.getUrlCode('code');
		// uni.showToast({
		// 	title: `微信:${Server.OpenID}`
		// })
		if (code == null) {
			this.getWeChatCode(isBase)
		}else{// 获取openid / access token
			let resp =await this.Oauth(code,!isBase)
			if (resp != null){
				console.log(resp);
			}
		}
		// #endif
	};
};


user.getUrlCode = function(name) {
	return decodeURIComponent((new RegExp('[?|&]' + name + '=' + '([^&;]+?)(&|#|;|$)').exec(location.href) || [,
		''])[1].replace(/\+/g, '%20')) || null
};

//请求微信接口，用来获取code
user.getWeChatCode = function(isBase) {
	let scope = "snsapi_base";
	if (!isBase){
		scope = "snsapi_userinfo";
	}
	let local = encodeURIComponent(window.location.href); //获取当前页面地址作为回调地址
	//通过微信官方接口获取code之后，会重新刷新设置的回调地址【redirect_uri】
	window.location.href =
		"https://open.weixin.qq.com/connect/oauth2/authorize?appid=" +
		Server.Appid +
		"&redirect_uri=" +
		local +
		"&response_type=code&scope="+scope+"&state=1#wechat_redirect";
}

// 微信授权获取登录信息
user.Oauth = async function(jscode,isUpdateUser){
	let resp = await Server.Post("/weixin.oauth",{jscode:jscode,isUpdateUser:isUpdateUser});
	if (resp.success){// 设置session
		Server.SessionID = resp.success.sessionId;
		Server.OpenID  = resp.success.openId;
	}
		
	return Server.GetReturn(resp);
}

// 获取用户信息
user.GetUserInfo = async function(){
	let resp = await Server.Post("/weixin.get_user_info",{});
	return Server.GetReturn(resp);
}

// 更新用户信息
user.UpsetUserInfo = async function(key,value){
	let resp = await Server.Post("/weixin.upset_user_info",{key:key,value:value});
	return Server.GetReturn(resp);
}

// 是否要重新登录
user.ReLogin = async function(){
	let resp = await Server.Post("/weixin.re_login",{});
	if (resp.success && !resp.success.relogin){// 设置session
		Server.SessionID = resp.success.sessionId;
		Server.OpenID  = resp.success.openId;
	}
	return Server.GetReturn(resp);
}

export default user
