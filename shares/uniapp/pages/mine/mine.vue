<template>
	<view class="mine animation-fade">
		<!-- 顶部 -->
		<view class="logo flexlogo" :hover-class="!info ? 'logo-hover' : ''"  >
			<image class="logo-img" :src="info ? info.avatarURL : ''" @click="ReLogin"></image>
			<view class="logo-title flexlogo">
				<text class="uer-name " v-if="info != null">{{info ? info.nickName : "Hi,您未登录"}}</text>
				<text class="go-login navigat-arrow"></text>
			</view>
		</view>

		<view class="cu-list menu sm-border card-menu margin-top bar-shadown">
			<!-- 我的手机 -->
			<view class="cu-item" @click="openPhone">
				<view class="content">
					<text class="cuIcon-roundadd text-grey"></text>
					<text class="text-grey">我的手机号</text>
				</view>
				<view v-if="info != null">{{info.phone != "" ? info.phone:"请输入手机号码"}}</view>
			</view>
			<uni-popup ref="popup" type="dialog">
				<uni-popup-dialog mode="input" placeholder="请输入手机号码" title="手机号码" :duration="2000" :before-close="true"
					@close="closePhone" @confirm="confirmPhone" :value="info ?info.phone :''"></uni-popup-dialog>
			</uni-popup>

			<!-- 我的股票容量 -->
			<view class="cu-item" @click="openShare">
				<view class="content">
					<text class="cuIcon-copy text-grey"></text>
					<text class="text-grey">股票容量</text>
				</view>
				<view>{{info.capacity}}</view>
			</view>
			<uni-popup ref="sharepop" type="dialog">
				<uni-popup-dialog title="股票容量" :duration="2000" :before-close="true" @close="closeShare"
					@confirm="sure">
					<uni-number-box v-model="vModelValue" :min="info.capacity"/>
				</uni-popup-dialog>
			</uni-popup>

			<!-- 我的组织 -->
			<view class="cu-item" @click="openGroup">
				<view class="content">
					<text class="cuIcon-info text-grey"></text>
					<text class="text-grey">我的组织</text>
				</view>
				<view v-if="info != null">{{info.group}}</view>
			</view>
			<uni-popup ref="orderpop" type="dialog">
				<uni-popup-dialog mode="input" placeholder="请输入组织代码" title="加入组织" :duration="2000" :before-close="true"
					@close="closeGroup" @confirm="confirmGroup" value=""></uni-popup-dialog>
			</uni-popup>

			<!-- 股票颜色开关 -->
			<view class="cu-item">
				<view class="content">
					<text class="text-grey">股票上涨颜色</text>
				</view>
				<view>
					<evan-switch v-model="info.rg" :size="20" active-color="#EA0000" inactive-color="#007500"
						@change="setSwitch"></evan-switch>
				</view>
			</view>
			
			<!-- 只显示20日线 -->
			<view class="cu-item">
				<view class="content">
					<text class="text-grey">仅显示20日线</text>
				</view>
				<view>
					<evan-switch v-model="info.only20" :size="20" active-color="#00aaff" inactive-color="#ffffff"
						@change="setOnly20Switch"></evan-switch>
				</view>
			</view>
			

			<!-- 股票列表 -->
			<view class="cu-item card-item">
				<view class="content">
					<text class="text-grey">我的股票列表</text><text style="color:#cfcfcf;font-size:20rpx;">(左滑删除)</text>
				</view>
				

				<uni-swipe-action style="width: 100%;">
					<uni-swipe-action-item :right-options="options" @click="onClick($event, item.code)"
						v-for="item in list" :key="item.code" :code="item.code" style="margin: 8rpx 0;">
						<navigator :url="'/pages/add/add?scode='+ item.code">
							<view class="cardContent">
								<view class="flex-subl flex">
									<view class="flex-title">{{item.name}}</view>
									<view class="flex-num flex">
										<view class="num-text">{{item.ext}}</view>
										<view class="flex-sub">
											{{item.simpleCode}}
										</view>
									</view>
								</view>

								<view class="flex-subr flex">
									<view class="tagb-content flex">
										<view class="tagb-left flex">
											<text class="tagb">{{item.price}}</text>
										</view>
										<view class="tagb-right flex" :style="{background:item.color}">
											<text class="tagb"><text
													v-if="item.percent > 0">+</text>{{item.percent}}%</text>
										</view>
									</view>
								</view>
							</view>
						</navigator>
					</uni-swipe-action-item>
				</uni-swipe-action>
			</view>

			<!-- 消息列表 -->
			<view class="msgCon card-item">
				<view class="content">
					<text class="text-grey">最新消息</text>
				</view>
				<view class="msgCard" v-for="(item,index) in msgs" :key="index">
					<navigator :url="'/pages/add/add?scode='+ item.code">
					<view class="msgTitle"> {{item.now}} </view>
					<view class="msg-cardContent">
						<view class="flex-top flex">

							<view class="flex-num flex">
								<view class="num-text">{{item.ext}}</view>
								<view class="flex-title">{{item.name}}</view>
							</view>
							<view class="msg-tag">{{item.key}}</view>

						</view>

						<view class="flex-bottom flex">
							<view class="tagb-content flex">
								<view class="flex-sub">当前</view>
								<view class="tagb-left flex" >
									<text class="tagb" :style="{color:item.color}">{{item.price}}</text>
								</view>
								<view class="tagb-right flex" :style="{color:item.color}">
									<text class="tagb"><text v-if="item.percent > 0">+</text>{{item.percent}}%</text>
								</view>
							</view>
							<view style="color: #AAAAAA;font-size: 20rpx; width: 45%;">{{item.desc}}</view>
						</view>
					</view>
					</navigator>
				</view>
			</view>
		</view>

	</view>
</template>

<script>
	import NavChange from '../index'
	export default {
		data() {
			return {
				info: {
					avatarURL: "",
					capacity: 0,
					city: "",
					country: "",
					gender: "",
					group: "",
					nickName: "",
					openid: "",
					phone: "",
					province: "",
					rg: true,
					only20: false,
				},
				list: [],
				msgs: [],
				interval: null,
				codes: [],
				options: [{
					text: '取消',
					style: {
						backgroundColor: '#007aff'
					}
				}, {
					text: '删除',
					style: {
						backgroundColor: '#dd524d'
					}
				}],
				vModelValue: 0,

			}
		},
		async created() { ////////////////// vue生命周期  页面创建时调用
			await this.loadData();
		},
		destroyed() { ////////////////////////////
			if (this.interval) {
				clearInterval(this.interval);
			}
		},
		mounted() {},
		methods: {
			/**
			 * 初始化
			 */
			async loadData() {
				console.log("loadDataloadData")
				await this.getInfo()

				this.getList()

				this.interval = setInterval(this.OnTimeDeal, 3000) // todo: start timeout
			},
			async getInfo(){
				let info = await this.$User.GetUserInfo(); // 登录
				if (info != null) {
					this.info = info
					this.vModelValue = info.capacity
					if (info.avatarURL.length == 0) { // 重新授权
						await this.$User.WxLogin(false); // 重新登录 
					}
				}
			},
			
			async getList(){
				// 获取列表
				let out = await this.$Shares.GetMyCode("");
				if (out) {
					this.list = out.list
					let that = this
					this.list.forEach(item => {
						that.codes.push(item.code);
					})
					this.getMsg();
				}
			},
			async getMsg() {
				let out = await this.$Shares.GetMsg();
				if (out) {
					this.msgs = out.list
				}
			},
			// 手机弹窗事件
			openPhone() {
				// 通过组件定义的ref调用uni-popup方法 ,如果传入参数 ，type 属性将失效 ，仅支持 ['top','left','bottom','right','center']
				this.$refs.popup.open('center')
			},
			closePhone() {
				// TODO 做一些其他的事情，before-close 为true的情况下，手动执行 close 才会关闭对话框
				// ...
				this.$refs.popup.close()
			},
			async confirmPhone(value) {
				if (this.info.phone == value) return

				let resp = await this.$User.UpsetUserInfo("phone", value);
				if (resp != null) {
					this.info.phone = value
					uni.showToast({
						title: "修改成功",
						icon: 'none'
					})
				}
				this.$refs.popup.close()
			},

			setSwitch: async function(value) { // 涨跌颜色设置
				console.log(value)
				let out = await this.$User.UpsetUserInfo("rg", value ? "true" : "false");
				if (out) {
					this.getMsg();
				}
			},
			
			setOnly20Switch: async function(value) { // 涨跌颜色设置
				console.log(value)
				let out = await this.$User.UpsetUserInfo("only20", value ? "true" : "false");
				if (out) {
					this.info.only20 = value;
				}
			},
			

			// 打开我的股票弹窗
			openShare() {
				// 通过组件定义的ref调用uni-popup方法 ,如果传入参数 ，type 属性将失效 ，仅支持 ['top','left','bottom','right','center']
				this.$refs.sharepop.open('center')
			},
			closeShare() {
				// TODO 做一些其他的事情，before-close 为true的情况下，手动执行 close 才会关闭对话框
				// ...
				this.$refs.sharepop.close()
			},
			
			// 打开我的组织弹窗
			openGroup() {
				this.$refs.orderpop.open('center')
			},
			closeGroup() {
				this.$refs.orderpop.close()
			},
			async confirmGroup(value) {
				this.$refs.orderpop.close()
				let out = await this.$Shares.AddGroup(value)
				if (out){
					uni.showToast({
						title: "添加成功",
						icon: 'none'
					})
					await this.getInfo()
				}
			},
			
			sure() {
				// 股票数量输入框的值
				console.log(this.vModelValue)
				this.$refs.sharepop.close()
			},
			// 滑动事件
			async onClick(e, code) {
				console.log(e.position)
				console.log(code)
				let _that = this
				if (e.position ==='right' && e.content.text === '删除'){// 删除
					uni.showModal({
						content: "删除将不可恢复,确认继续?",
						showCancel: true,
						success:async function (res) {
						        if (res.confirm) {
									let out = await _that.$Shares.DeleteMyCode(code)
									if (out) {
										uni.showToast({
											title: "删除成功",
											icon: 'none'
										})
										_that.getList()
									}
						        }// else if (res.cancel) {
						        
						    }
					});
				}
				 console.log('点击了' + (e.position === 'left' ? '左侧' : '右侧') + e.content.text + '按钮')
			},
			async OnTimeDeal() {
				if (this.codes.length == 0) {
					return
				}

				let codes = await this.$Shares.Gets(this.codes); // 获取分组信息
				if (codes != null) { // 去更新
					let _map = []
					codes.list.forEach(item => {
						_map[item.code] = item
					})

					for (let i = 0; i < this.list.length; i++) {
						let tmp = _map[this.list[i].code]
						if (tmp != null) {
							this.list[i].price = tmp.price
							this.list[i].percent = tmp.percent
							this.list[i].color = tmp.color
						}
					}
				}
			},
			async ReLogin(){// 重新登录
				this.$User.Clear();
				await this.$User.WxLogin(false);// 登录
			},
		}
	}
</script>

<style lang="scss" scoped>
	.flex {
		display: flex;
		justify-content: center;
		align-items: center;
	}

	button .cu-tag {
		position: absolute;
		top: 8upx;
		right: 8upx;
	}

	@font-face {
		font-family: texticons;
		font-weight: normal;
		font-style: normal;
	}

	.flexlogo {
		display: flex;
	}

	page {
		background-color: #f8f8f8;
	}

	.center {
		flex-direction: column;
	}

	.logo {
		width: 750upx;
		height: 360upx;
		padding: 36upx;
		box-sizing: border-box;
		background: linear-gradient(135deg, rgba(35, 35, 35, 0.9), rgba(35, 35, 35, 0.7));
		;
		flex-direction: row;
		align-items: center;

	}

	.logo-hover {
		opacity: 0.8;
	}

	.logo-img {
		width: 150upx;
		height: 150upx;
		border-radius: 150upx;
		background: #000000;
	}

	.logo-title {
		height: 150upx;
		flex: 1;
		align-items: center;
		justify-content: space-between;
		flex-direction: row;
		margin-left: 20upx;
	}

	.uer-name {
		height: 60upx;
		line-height: 60upx;
		font-size: 38upx;
		color: #FFFFFF;
	}

	.go-login.navigat-arrow {
		font-size: 38upx;
		color: #FFFFFF;
	}

	.login-title {
		height: 150upx;
		align-items: self-start;
		justify-content: center;
		flex-direction: column;
		margin-left: 20upx;
	}

	.center-list {
		background-color: #FFFFFF;
		margin-top: 20upx;
		width: 750upx;
		flex-direction: column;
	}

	.center-list-item {
		height: 90upx;
		width: 750upx;
		box-sizing: border-box;
		flex-direction: row;
		padding: 0upx 20upx;
	}

	.border-bottom {
		border-bottom-width: 1upx;
		border-color: #c8c7cc;
		border-bottom-style: solid;
	}

	.list-icon {
		width: 40upx;
		height: 90upx;
		line-height: 90upx;
		font-size: 34upx;
		color: #4cd964;
		text-align: center;
		font-family: texticons;
		margin-right: 20upx;
	}

	.list-text {
		height: 90upx;
		line-height: 90upx;
		font-size: 34upx;
		color: #555;
		flex: 1;
		text-align: left;
	}

	.navigat-arrow {
		height: 90upx;
		width: 40upx;
		line-height: 90upx;
		font-size: 34upx;
		color: #555;
		text-align: right;
		font-family: texticons;
	}

	.msgCon {
		padding: 0 15px;
	}

	.msgCon .msgCard {
		margin-top: 10px;
	}

	.cu-list {
		margin-bottom: 164rpx;
	}

	.cu-list .card-item {
		margin: 10px auto;
		padding-top: 10px;
		flex-direction: column;
		align-items: flex-start;
	}

	.cardContent {
		display: flex;
		justify-content: space-between;
		align-items: center;
		width: 100%;
		height: 140rpx;
		background: #000000;
		margin: 0 4rpx;
		padding: 0 24rpx;
	}

	.flex-subl {
		flex-direction: column;
	}

	.flex-subl .flex-title {
		font-size: 32rpx;
		font-weight: bold;
		color: #FFFFFF;
		margin-bottom: 15rpx;
	}

	.flex-subl .num-text {
		background-color: red;
		color: #000000;
		font-size: 28rpx;
		margin-right: 10rpx;
	}

	.flex-subl .flex-sub {
		font-size: 28rpx;
		color: #FFFFFF;
	}

	.flex-subr .tagb-content {
		flex-direction: row;
	}

	.flex-subr .tagb-content .tagb-left {
		color: #FFFFFF;
		font-size: 44rpx;
	}

	.flex-subr .tagb-content .tagb-right {
		color: #FFFFFF;
		font-size: 44rpx;
		background: #1AAD19;
		padding: 10rpx 24rpx;
		margin-left: 16rpx;
	}

	.msg-cardContent {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: flex-start;
		width: 100%;
		height: auto;
		background: #000000;
		margin: 0 4rpx;
		padding: 10rpx 0;
	}

	.msg-tag {
		border: 1rpx solid #555555;
		color: #AAAAAA;
	}

	.flex-top {
		padding: 0 24rpx;
		margin: 20rpx 0;
		width: 100%;
		justify-content: space-between;
	}

	.flex-top .flex-title {
		font-size: 32rpx;
		font-weight: bold;
		color: #FFFFFF;
	}

	.flex-top .num-text {
		background-color: red;
		color: #000000;
		font-size: 28rpx;
		margin-right: 10rpx;
	}

	.flex-top .flex-sub {
		font-size: 28rpx;
		color: #FFFFFF;
	}

	.flex-subr .tagb-content {
		flex-direction: row;
	}

	.flex-subr .tagb-content .tagb-left {
		color: #FFFFFF;
		font-size: 44rpx;
	}

	.flex-subr .tagb-content .tagb-right {
		color: #FFFFFF;
		font-size: 44rpx;
		background: #1AAD19;
		padding: 10rpx 24rpx;
		margin-left: 16rpx;
	}

	.flex-bottom {
		margin-bottom: 10rpx;
		padding: 6rpx 24rpx;
		width: 100%;
		justify-content: space-between;
	}

	.flex-bottom .tagb-content {
		flex-direction: row;
	}

	.flex-bottom .tagb-content .tagb-left {
		color: #DC143C;
		font-size: 30rpx;
	}

	.flex-bottom .tagb-content .tagb-right {
		color: #DC143C;
		font-size: 30rpx;
		margin-left: 16rpx;
	}

	.flex-bottom .flex-sub {
		font-size: 24rpx;
		color: #FFFFFF;
		margin-right: 18rpx;
	}
</style>
