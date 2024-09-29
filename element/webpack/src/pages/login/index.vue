<template>
	<div class="web_qrcode_type_page_embedded">
		<div class="impowerBox">
			<div class="title">微信登录</div>
			<el-image class="qrcode lightBorder" :src="data.info.url" mode=""></el-image>
			<div class="info">
				<div v-if="!text" class="status status_browser js_status js_wx_default_tip normal" id="wx_default_tip">
					<p class="status_msg">使用微信扫一扫登录</p>
					<p class="status_msg">“复利备忘录”</p>
				</div>
				<div v-if="text" class="status status_succ js_status js_wx_after_scan normal" id="wx_after_scan">
					<i class="status_icon icon38_msg succ"></i>
					<div class="status_txt">
						<h4 class="status_msg">扫描成功</h4>
						<p class="status_msg">{{ text }}</p>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>

<script setup>

import server from '@/utils/server/def.js'
import { useRoute,useRouter } from 'vue-router';
import User from '@/utils/server/oauth'
import {  reactive, ref } from 'vue'
const diviceId = ref('')
const text = ref('')
const data = reactive({
	info: {},
	
	isReload: ref(true),
	// interval: null
})
const route = useRoute();
const router = useRouter();

const login = async () => {
	data.info = await User.H5Login(diviceId.value); // 登录
	if (data.info != null) {
		if (data.info.status == 1) { // 登录成功
			text.value = '欢迎，' + data.info.userName
			server.SetCookie("user-token", data.info.userId)
			localStorage.setItem("user-token", data.info.userId)
			window.parent.postMessage({
				type: 'user',
				userId: data.info.userId,
			}, '*');
			if (data.isReload) {
				setTimeout(() => {
					router.replace('zyb')
				}, 1500)
			}
			return
		}
	}
	setTimeout(() => {
		login()
	}, 3000)
}

const loadData = async () => {
	diviceId.value = route?.query?.diviceid
	if (diviceId.value != null && diviceId.value.length > 0) {
		data.isReload = false
	} else {
		diviceId.value = localStorage.getItem("divice_id")
		if (!diviceId.value) {
			diviceId.value = 'DID_' + Date.now().toString(36) + '_' + Math.random().toString(36).substr(2, 9);
			localStorage.setItem("divice_id", diviceId.value)
		}
	}
	login()
	// data.interval = setInterval(login, 3000) // todo: start timeout
}
loadData();

// onUnmounted(() => {
// 	if (data.interval) {
// 		clearInterval(data.interval);
// 	}
// })

</script>
<script>
export default {
	name: 'LoginView',
}
</script>
<style lang="scss" scoped>
.web_qrcode_type_page_embedded {
	background-color: rgb(51, 51, 51);
	padding: 100px;
	height: 100vh;
}

.impowerBox {
	line-height: 1.6;
	position: relative;
	display: inline-block;
	width: 100%;
	vertical-align: middle;
	z-index: 1;
	text-align: center;
}

.impowerBox .title {
	text-align: center;
	font-size: 40px;
	color: #fff;
}

.impowerBox .qrcode {
	height: 280px;
	margin-top: 15px;
	border: 1px solid #e2e2e2;
}

.impowerBox .info {
	width: 300px;
	margin: 0 auto;
}

.impowerBox .status.normal {
	box-sizing: border-box;
	color: #fff;
	margin-top: 30px;
	background-color: #232323;
	border-radius: 100px;
	-moz-border-radius: 100px;
	-webkit-border-radius: 100px;
	box-shadow: inset 0 10px 20px -10px #191919, 0 1px 0 0 #444;
	-moz-box-shadow: inset 0 10px 20px -10px #191919, 0 1px 0 0 #444;
	-webkit-box-shadow: inset 0 10px 20px -10px #191919, 0 1px 0 0 #444;
	padding: 0px 28px;
}

.status_msg {
	margin: 10px 5px;
}

#wx_after_scan {
	text-align: left;
	display: flex;
	align-items: center;

	.icon38_msg {
		display: inline-block;
		width: 44px;
		height: 44px;
		background-size: cover;
		background-position: no-repeat;
		background-image: url(@/assets/wx.png);
	}

	.status_txt {
		display: inline-block;
		vertical-align: middle;
	}


}
</style>