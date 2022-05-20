<template>
	<view class="home animation-fade">
		<!-- LOGO -->
		<view class="logo">
			<image src="../../static/watch_icon.png">
		</view>
		<scroll-view scroll-y class="page">
			<!-- 我的组织列表 -->
			<view class="animation-fade" v-for="(item,index) in list" :key="index">
				<view class="cu-card article">
					<view class="cu-item bar-shadown">
						<view class="title">
							<view class="text-cut">
								<text class="text-black">{{item.name}}</text>
							</view>
						</view>
						<view v-for="(item1,index) in item.list" :key="index">
							<navigator :url="'/pages/add/add?scode='+ item1.code" >
								<view class="content" >
									<view class="flex-subl flex">
										<view class="flex-title">{{item1.name}}</view>
										<view class="flex-num flex">
											<view class="num-text">{{item1.ext}}</view>
											<view class="flex-sub">{{item1.simpleCode}}</view>
										</view>
										<view class="flex-sub1">{{item1.attach}}</view>
									</view>
									<view class="flex-subr flex">
										<view class="tagb-content flex">
											<view class="tagb-left flex">
												<text class="tagb">{{item1.price}}</text>
											</view>
											<view class="tagb-right flex" :style="{background:item1.color}">
												<text class="tagb"><text v-if="item1.percent > 0">+</text>{{item1.percent}}%</text>
											</view>
										</view>
									</view>
								</view>
							</navigator>
						</view>
					
					</view>
				</view>
			</view>
		<view class="ring"></view>
		</scroll-view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				list: [],
				interval: null,
				codes: [],
			}
		},
		// updated:function(){////////////////
		// 	console.log('updated')
		// },
		async created() { ////////////////// vue生命周期  页面创建时调用
			await this.loadData();
		},
		destroyed() { ////////////////////////////
			if (this.interval) {
				clearInterval(this.interval);
			}
		},
		methods: {
			/**
			 * 初始化
			 */
			async loadData() {
				let group = await this.$Shares.GetGroup(); // 获取分组信息
				if (group != null) {
					this.list = group.list
					let that = this
					this.list.forEach(item => {
						item.list.forEach(item1 => {
							that.codes.push(item1.code);	
						})
					})
					this.interval = setInterval(this.OnTimeDeal, 3000) // todo: start timeout
				}else{
					setTimeout(this.loadData,1000)
				}
			},

			async OnTimeDeal() {
				if (this.codes.length == 0){
					return
				}
				
				let codes = await this.$Shares.Gets(this.codes); // 获取分组信息
				if (codes != null) {// 去更新
				let _map = []
				codes.list.forEach(item => {
					_map[item.code] = item
				})
				
				 for(let i=0;i<this.list.length;i++){
					 for (let j= 0;j<this.list[i].list.length;j++){
						let tmp = _map[this.list[i].list[j].code]
						 if (tmp != null){
							 this.list[i].list[j].price = tmp.price
							 this.list[i].list[j].percent = tmp.percent
							 this.list[i].list[j].color = tmp.color
						 }
					 }
				 }
				}
			},
		},
		// mounted() {
		// 	console.log("home:------------------")
		// },
	}
</script>

<style scoped>
	.flex {
		display: flex;
		justify-content: center;
		align-items: center;
	}

	.logo {
		display: flex;
		align-items: center;
		justify-content: center;
		height: 480rpx;
	}

	.logo image {
		width: 200rpx;
		height: 200rpx;
	}

	.cu-item {
		/* background: lightblue; */
	}
	
	.ring {
		margin-bottom: calc(160upx + env(safe-area-inset-bottom) / 2);
	}

	.content {
		display: flex;
		justify-content: space-between;
		align-item: center;
		width: 100%;
		height: 140rpx;
		background: #000000;
		margin: 24rpx 4rpx;
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
	
	.flex-subl .flex-sub1 {
		font-size: 20rpx;
		color: #8799a3;
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
</style>
