<script>
	import Vue from 'vue'
	
	export default {
		created() {
			// #ifdef APP-PLUS
			plus.navigator.closeSplashscreen(); 
			// #endif 
		},
		onLaunch:async function() {
			uni.getSystemInfo({
				success: function(e) {
					// #ifndef MP
					Vue.prototype.StatusBar = e.statusBarHeight;
					if (e.platform == 'android') {
						Vue.prototype.CustomBar = e.statusBarHeight + 50;
					} else {
						Vue.prototype.CustomBar = e.statusBarHeight + 45;
					};
					// #endif
		
					// #ifdef MP-WEIXIN
					Vue.prototype.StatusBar = e.statusBarHeight;
					let custom = wx.getMenuButtonBoundingClientRect();
					Vue.prototype.Custom = custom;
					Vue.prototype.CustomBar = custom.bottom + custom.top - e.statusBarHeight;
					// #endif		
		
					// #ifdef MP-ALIPAY
					Vue.prototype.StatusBar = e.statusBarHeight;
					Vue.prototype.CustomBar = e.statusBarHeight + e.titleBarHeight;
					// #endif
				}
			})
		},
		onShow:async function() {
			uni.removeStorage({
				key: 'chineseCodes',
				success() {
					console.log('删除缓存(chineseCodes)成功')
				}
			})
			console.log('App 开启')
			let relog = await this.$User.ReLogin();
			if (relog && relog.relogin){// 需要重新登录
				await this.$User.WxLogin(false);// 登录
			}
		},
		onHide: function() {
			console.log('App 关闭')
		}
	}
</script>

<style lang="scss">
	/* ColorUi */
	@import url("./lib/colorui/main.css");
	@import url("./lib/colorui/icon.css");
	
	/* 加载loding */
	.cu-load.load-modal {
		width: 0px;
		height: 0px;
		background-color: transparent;
	}
	.cu-load.load-modal::after {
		background-color: rgba(0, 0, 0, 0.3);
		width: 30px;
		height: 30px;
		border-left: 3px solid #FFF0F6;
	}
	body{
		background: #FFFFFF !important;
	}
	/* 阴影 */
	.bar-shadown{
		-webkit-box-shadow: 0 0 36px 0 rgba(43,86,112,.1) !important;
		box-shadow: 0 0 36px 0 rgba(43,86,112,.1) !important;
	}
</style>
