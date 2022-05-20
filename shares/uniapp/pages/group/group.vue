<template>
	<view>
		<view class="ble animation-fade">
			<!-- 顶部 -->
			<view style="margin: 20rpx 10rpx;" class="flex cul">
				<view class="flex width">
					<uni-icons type="home" size="25" @click="backHome"></uni-icons>
					<uni-search-bar :radius="100" @confirm="search" v-model="searchValue" cancelText="搜索"
						@cancel="search" @change="change" placeholder="搜索股票" style="width: 90%;">
					</uni-search-bar>
				</view>
				<view class="padding flex justify-between align-start bg-white" v-for="(item,index) in availableList"
					:key="index" @click="itemClick(item.code)">
					<text>{{item.name}}</text>
				</view>
			</view>
		</view>
		<view class="pic" v-if="info !== null">
			<iframe :src="info.img" :key="iframekey" style="width:100%;height:100%;background-color:#100C2A;"
				name="iframe_a" frameborder="no" border="0" marginwidth="0" marginheight="0" scrolling="no"
				allowtransparency="yes">
			</iframe>
		</view>

		<view class="content" v-if="info !== null">
			<view class="flex-subl flex">
				<view class="flex-title">{{info.name}}</view>
				<view class="flex-num flex">
					<view class="num-text">{{info.ext}}</view>
					<view class="flex-sub">{{info.simpleCode}}</view>
				</view>
			</view>
			<view class="flex-subr flex">
				<view class="tagb-content flex">
					<view class="tagb-left flex">
						<text class="tagb">{{info.price}}</text>
					</view>
					<view class="tagb-right flex" :style="{background:info.color}">
						<text class="tagb"><text v-if="info.percent > 0">+</text>{{info.percent}}%</text>
					</view>
				</view>
			</view>
		</view>

		<view class="formContent" v-if="info !== null">
			<view class="uni-form-item right-part">
				<view class="uni-form-item uni-column cu-item" v-if="info !== null">
					<view class="siteContent">{{klineName}}</view>
					<view>
						<evan-switch v-model="kiline" :size="20" active-color="#aaffff" inactive-color="#00FFFF"
							@change="setklineSwitch"></evan-switch>
					</view>
				</view>
			</view>
			
			<view class="inputItem1" v-if="info.hy != ''">
				行业板块:{{info.hy}}
			</view>
			
			<view class="uni-form-item flex input-item">
				<view class="inputItem flex inputMar">
					<text class="inputTitle">推荐人：</text>
					<input class="input" name="up" placeholder="推荐人" type="text" v-model="userName" />
				</view>
			</view>

			<view v-for="(gname,index) in group" :key="index">
				<text style="padding: 0rpx 0;font-size: 30rpx;">{{gname}}：</text>
				<view class="sendBut" name="aaaa">
					<button align="right" style="float:left;width: 30%;background-color: #b3b9b9;"
						@click="formSubmit(gname,false)">删除</button>
					<button align="left" style="width: 70%;background-color: #ff5500;"
						@click="formSubmit(gname,true)">添加</button>
				</view>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				nickName:"",
				userName: "",
				iframekey: 0,
				klineName: "日K",
				kiline: true,
				chineseCodes: [],
				availableList: [],
				info: null,
				searchValue: "",
				tag: "",
				checked: true,
				size: 20,
				group: []
			}
		},
		async created(options) { ////////////////// vue生命周期  页面创建时调用
			this.loadData()
		},
		onLoad(options) {
			this.iframekey = 1;
			let tag = options.tag
			if (tag != null && tag.length > 0) {
				this.tag = tag;
				if (this.tag == "min") {
					this.kiline = false
					this.klineName = "分时"
				}
			}
			let code = options.scode
			if (code != null && code.length > 0) {
				this.searchValue = code;
				this.searchCode(code)
			}
			
			this.getInfo();
		},
		destroyed() { ////////////////////////////
			console.log("destroyed --- add")
		},

		mounted() {
			// var _this=this
			// getisLogin() //是否登陆
		},

		/**
		 * watch 
		 */
		watch: {
			// 监听搜索框值的变化
			searchValue(n, o) {
				this.availableList = []
				let that = this
				if (n.length > 0 && this.isChinese(n)) {
					this.chineseCodes.forEach(item => {
						if (item.name.search(n) != -1) {
							that.availableList.push(item)
						}
					})
				} else if(/^[a-zA-Z]+$/.test(n)){ // 拼音
					n = n.toLowerCase();
					this.chineseCodes.forEach(item => {
						if (item.sName == n) {
							that.availableList.push(item)
						}
					})
				}
			}
		},
		methods: {
			async getInfo(){
				let info = await this.$User.GetUserInfo(); // 登录
				if (info != null) {
					this.userName = info.nickName
					this.nickName = info.nickName
				}
			},
			itemClick: function(e) { // 点击提醒框
				console.log("itemClick")
				console.log(e)
				this.searchCode(e)
			},
			search: function(e) { // 点击搜索
				console.log("search")
				this.searchCode(e.value)
			},
			cancel: function(e) {
				console.log(e)
				// this.info = null
			},
			isChinese: function(str) { // 是否包含中文
				if (/^[\u3220-\uFA29]+$/.test(str)) {
					return true;
				} else {
					return false;
				}
			},
			selectCk: function(e) {
				console.log('eee', e)
				console.log('多选选中的值', e.detail.value)
			},
			formSubmit: async function(group, isAdd) {
				if (this.userName.length  == 0){
					uni.showModal({
						content: "请填写推荐人",
						showCancel: false
					});
					// this.userName = this.nickName
					return
				}
				let resp = await this.$Shares.UpsetGroupCode(this.info.code, group, this.userName, isAdd);
				console.log(resp)
				if (resp) {
					uni.showModal({
						content: "操作成功",
						showCancel: false
					});
				}else{
					uni.showModal({
						content: "操作失败",
						showCancel: false
					});
				}
			},
			setklineSwitch: function(value) { // 分时，日k
				this.iframekey = this.iframekey + 1;
				this.kiline = value
				if (this.kiline == true) {
					this.tag = "daily"
					this.info.img = this.replaceAll(this.info.img, "min", "daily");
					this.klineName = "日K"
				} else {
					this.tag = "min"
					this.info.img = this.replaceAll(this.info.img, "daily", "min");
					this.klineName = "分时"
				}
				console.log(this.info.img);

				console.log(value)
			},
			replaceAll(string, search, replace) {
				return string.split(search).join(replace);
			},
			backHome: function() {
				console.log("...........................")
				uni.navigateTo({
					url: '/pages/index'
				});
			},
			/**
			 * 初始化
			 */
			async loadData() {
				let _that = this;
				uni.getStorage({
					key: 'chineseCodes',
					success: function(res) {
						_that.chineseCodes = res.data;
					},
					fail: async function(res) {
						let list = await _that.$Shares.GetAllCodeName(); // 获取所有的中文代码
						if (list) {
							_that.chineseCodes = list.list;
							uni.setStorage({
								key: 'chineseCodes',
								data: _that.chineseCodes
							});
						}
					}
				});
			},
			async searchCode(code) {
				let info = await this.$Shares.Search(code, this.tag); // 获取所有的中文代码
				if (info) {
					this.info = info.info;
					// todo:获取已经配置了的，设置更新
					let out = await this.$Shares.GetMyGroup(this.info.code);
					console.log(out)
					if(out){
						this.group = out.group
						if (this.group.length == 0) {
							uni.showModal({
								content: '请到个人中心添加加入一个组织',
								showCancel: false
							});
						}
						if (out.code.length > 0) {
							this.userName = out.userName
						}else{
							this.userName = this.nickName
						}
					}

				}
				this.availableList = []
			}
		}
	}
</script>
<style scoped>
	.flex {
		display: flex;
		align-items: center;
		justify-content: center;
	}

	.cul {
		flex-direction: column;
	}

	.width {
		width: 100%;
	}

	.cu-item {
		display: flex;
		align-items: center;
		justify-content: center;
		width: 100%;
		height: 100rpx;
	}

	.right-part {
		display: flex;
		align-items: center;
		justify-content: space-between;
		width: 100%;
		height: 100rpx;
	}

	.cu-item .siteContent {
		margin-right: 40rpx;
		font-size: 34rpx;
	}

	.pic {
		width: 100%;
		height: 800rpx;
		/* margin-bottom: 24rpx; */
	}

	.pic image {
		display: block;
		width: 100%;
		height: 100%;
	}

	.content {
		display: flex;
		justify-content: space-between;
		align-item: center;
		width: 100%;
		height: 120rpx;
		padding: 0 24rpx;
		background: #000000;

	}

	.flex-subl {
		flex-direction: column;
	}

	.flex-subl .flex-title {
		margin-bottom: 15rpx;
		font-size: 32rpx;
		font-weight: bold;
		color: #FFFFFF;
	}

	.flex-subl .num-text {
		margin-right: 10rpx;
		font-size: 28rpx;
		color: #000000;
		background-color: red;
	}

	.flex-subl .flex-sub {
		font-size: 28rpx;
		color: #FFFFFF;
	}

	.flex-subr .tagb-content {
		flex-direction: row;
	}

	.flex-subr .tagb-content .tagb-left {
		font-size: 44rpx;
		color: #FFFFFF;
	}

	.flex-subr .tagb-content .tagb-right {
		padding: 10rpx 24rpx;
		margin-left: 16rpx;
		font-size: 44rpx;
		color: #FFFFFF;
		background: #1AAD19;
	}

	.checkbox {
		flex-direction: row;
		justify-content: space-evenly;
		width: 70%;
		margin: 60rpx auto;
	}

	.checkbox .checkboxItem {
		flex-direction: column;
		align-items: flex-start;
	}

	.checkbox .checkboxItem text {
		margin-right: 20rpx
	}

	.input-item {
		flex-direction: row;
	}

	.inputItem {
		flex-direction: row;
		width: 40%;
	}
	.inputItem1 {
		display: flex;
		flex-direction: row;
		align-items: center;//垂直居中
		width : 80%;
		margin : 20px 10%;
	}

	.inputItem .inputTitle {
		padding: 40rpx 0;
		font-size: 30rpx;
	}

	.inputItem .input {
		width: 150rpx;
		height: 60rpx;
		border: #000000;
	}

	/* .inputMar {
		margin-right: 20rpx;
	} */

	.sendBut {
		width: 80%;
		height: 40rpx;
		margin: 40rpx auto calc(100upx + env(safe-area-inset-bottom) / 2);
	}
</style>
