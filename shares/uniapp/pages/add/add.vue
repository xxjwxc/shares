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
				<view class="padding flex justify-between align-start bg-white" v-for="(item,index) in availableList" :key="index" @click="itemClick(item.code)">
					<text>{{item.name}}</text>
				</view>
			</view>
		</view>
		<view class="pic" v-if="info !== null">
			<iframe :src="info.img" :key="iframekey" style="width:100%;height:100%;background-color:#100C2A;" name="iframe_a"  
			frameborder="no" border="0" marginwidth="0" marginheight="0" scrolling="no" allowtransparency="yes">
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
			<form @submit="formSubmit" @reset="formReset">
				<view class="uni-form-item">
					<view class="uni-form-item right-part" v-if="info !== null">
						<view class="uni-form-item uni-column cu-item" v-if="info !== null">
							<view class="siteContent">启用提醒</view>
							<view>
								<evan-switch v-model="oldInfo.vaild" :size="20" active-color="#0000FF" inactive-color="#808080" @change="setSwitch"></evan-switch>
							</view>
						</view>
						
						<view class="uni-form-item uni-column cu-item" v-if="info !== null">
							<view class="siteContent">{{klineName}}</view>
							<view>
								<evan-switch v-model="kiline" :size="20" active-color="#aaffff" inactive-color="#00FFFF" @change="setklineSwitch"></evan-switch>
							</view>
						</view>
					</view>
					<checkbox-group class="checkbox flex" name="checkbox" @change="selectCk">
						<label class="checkboxItem flex">
							<label>
								<checkbox value="surge" :checked="oldInfo.surge" /><text>盘中急涨提醒</text>
								<checkbox value="slump" :checked="oldInfo.slump" /><text>盘中急跌提醒</text>
								<br /><br />
								<checkbox value="ai" :checked="oldInfo.ai" /><text>AI智能提醒</text>
								<checkbox value="kdj20" :checked="oldInfo.kdj20" /><text>20日线提醒</text>
								
								<br /><br />
								<checkbox value="kdj" :checked="oldInfo.kdj" /><text>日线金叉提醒</text>
								<checkbox value="public" :checked="oldInfo.public" /><text>公开到组织</text>
								
							</label>
						</label>
					</checkbox-group>
				</view>
				<view class="uni-form-item flex input-item">
					<view class="inputItem flex inputMar">
						<text class="inputTitle">估价涨到：</text>
						<input class="input" name="up" placeholder="涨到" type="text" v-model="oldInfo.up" />
					</view>
					<view class="inputItem flex">
						<text class="inputTitle">估价跌到：</text>
						<input class="input" name="down" placeholder="跌到" type="text" v-model="oldInfo.down" />
					</view>
				</view>
				<view class="uni-form-item flex input-item">
					<view class="inputItem flex inputMar">
						<text class="inputTitle inputMar">涨幅超：</text>
						<input class="input" name="upPercent" placeholder="涨幅" type="text" v-model="oldInfo.upPercent"/>%
					</view>
					<view class="inputItem flex">
						<text class="inputTitle inputMar">跌幅超：</text>
						<input class="input" name="downPercent" placeholder="跌幅" type="text"
							v-model="oldInfo.downPercent" />%
					</view>
				</view>
				<view class="inputItem1"  v-if="info.peg != ''">
					PEG指标:{{info.peg}}
				</view>
				<view class="inputItem1"  v-if="info.hy != ''">
					行业板块:{{info.hy}}
				</view>
				<view class="sendBut">
					<button form-type="submit">提交</button>
				</view>
			</form>
		</view>

	</view>
</template>

<script>
	export default {
		data() {
			return {
				iframekey:0,
				klineName:"日K",
				kiline:true,
				chineseCodes: [],
				availableList: [],
				info: null,
				oldInfo: {
					vaild: true,
				},
				searchValue: "",
				tag:"",
				checked: true,
				size: 20,
			}
		},
		async created(options) { ////////////////// vue生命周期  页面创建时调用
			this.loadData()
		},
		onLoad(options) {
			this.iframekey = 1;
			let tag = options.tag
			if (tag != null &&tag.length > 0){
				this.tag = tag;
				if (this.tag == "min"){
					this.kiline = false
					this.klineName = "分时"
				}
			}
			let code = options.scode
			if (code != null &&code.length > 0){
				this.searchValue = code;
				this.searchCode(code)
			}
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
				}else if(/^[a-zA-Z]+$/.test(n)){ // 拼音
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
			formSubmit: async function(e) {
				let b = false //  是否需要更新
				let data = e.detail.value
				console.log(data)
				let req = {}
				req["code"] = this.info.code
				req["price"] = this.info.price
				data.checkbox.forEach((item) => {
					req[item] = true
					b = true
				})
				req["up"] = Number(data.up)
				req["down"] = Number(data.down)
				req["upPercent"] = Number(data.upPercent)
				req["downPercent"] = Number(data.downPercent)
				req["vaild"] = this.oldInfo.vaild

				console.log(b)
				console.log(req)

				b = b || (req["up"] > 0) || (req["down"] > 0) || (req["upPercent"] > 0) || (req["downPercent"] > 0)
				console.log(b)
				if (!b) {
					uni.showModal({
						content: '请至少输入一个有效项',
						showCancel: false
					});
					return;
				}

				let resp = await this.$Shares.AddMyCode(req);
				if (resp) {
					if (!resp.status) {
						uni.showModal({
							content: resp.msg,
							showCancel: false
						});
					} else {
						uni.showModal({
							content: "添加成功",
							showCancel: false
						});
					}
				}
			},
			formReset: function(e) {
				console.log('清空数据')
			},
			setSwitch: function(value) { // 涨跌颜色设置
				console.log(value)
				this.oldInfo.vaild = value
			},
			setklineSwitch: function(value) { // 涨跌颜色设置
				this.iframekey = this.iframekey+1;
				this.kiline = value
				if (this.kiline == true){
					this.tag = "daily"
					this.info.img = this.replaceAll(this.info.img,"min","daily");
					this.klineName = "日K"
				}else{
					this.tag = "min"
					this.info.img = this.replaceAll(this.info.img,"daily","min");
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
				    success: function (res) {
				        _that.chineseCodes = res.data;
				    },
					fail: async function(res){
						// _that.oldInfo = _that.defaultOldInfo;
						let list = await _that.$Shares.GetAllCodeName(); // 获取所有的中文代码
						if (list) {
							_that.chineseCodes = list.list;
							uni.setStorage({key: 'chineseCodes',data: _that.chineseCodes});
							// list.list.forEach(item => {
							// 	_that.chineseCodes.push(item)
							// })
						}
					}
				});
			},
			async searchCode(code) {
				let info = await this.$Shares.Search(code,this.tag); // 获取所有的中文代码
				if (info) {
					this.info = info.info;
					// todo:获取已经配置了的，设置更新
					let out = await this.$Shares.GetMyCode(this.info.code);
					console.log(out)
					if (out && out.list != null) {
						this.oldInfo = out.list[0]
					} else {
						this.oldInfo = {
							vaild: true,
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

	.inputItem .inputTitle {
		padding: 20rpx 0;
		font-size: 30rpx;
	}

	.inputItem .input {
		width: 120rpx;
		height: 60rpx;
		border: #000000;
	}
	
	.inputItem1 {
		display: flex;
		flex-direction: row;
		align-items: center;//垂直居中
		width : 80%;
		margin : 20px 10%;
	}
	
	/* .inputMar {
		margin-right: 20rpx;
	} */

	.sendBut {
		width: 80%;
		height: 240rpx;
		margin: 40rpx auto calc(160upx + env(safe-area-inset-bottom) / 2);
	}
</style>
