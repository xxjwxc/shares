<template>
	<!-- <el-switch v-model="value2" class="ml-2" @change="switchchange" /> -->
	<div class="switchbar0">
		<el-segmented class="switchbar" v-model="btnActive" :options="klineName" @change='confirmkline' block />
	</div>
	<div class="space">
		<div class="split">
			<div style="display: flex;">
				<div style="font-size: 20px;color: #aa00ff;flex: 1;">国债逆回购</div>
				<div style="font-size: 20px;color: #aa00ff;flex: 1;">北向资金</div>
			</div>
			<div class="row">
				<div class="switchChart" ref="chartmlf"></div>
				<div class="switchChart" ref="chartbx"></div>
			</div>
		</div>

		<div class="split">
			<div style="display: flex;">
				<div style="font-size: 20px;color: #aa00ff;flex: 1;">融资融券</div>
				<div style="font-size: 20px;color: #aa00ff;flex: 1;">
					成交额对比:{{ cjl.tm }}
					<el-segmented class="switchbar2" v-model="btnActive2" :options="klineName2" @change='confirmkline2'
						block />
				</div>
			</div>
			<div class="row">
				<div class="switchChart" ref="chartrzrq"></div>
				<div class="switchChart" ref="chart_cjl"></div>
			</div>
		</div>

		<div class="split">
			<div style="display: flex;">
				<div style="font-size: 20px;color: #aa00ff;flex: 1;">美元兑人民币</div>
			</div>
			<div class="switchChart" ref="charusdcny"></div>
		</div>
		<div class="split">
			<div class="row">
				<div class="switchChart" ref="chartm2"></div>
				<div class="switchChart" ref="chartcpi"></div>
			</div>
		</div>
		<div class="split">
			<div class="row">
				<div class="switchChart" ref="chartgdp"></div>
				<div class="switchChart" ref="chartzs"></div>
			</div>
		</div>
	</div>
</template>

<script>
// import server from '@/utils/server/def.js'
import { useRoute } from 'vue-router';
import User from '@/utils/server/oauth';
import { ElMessageBox } from 'element-plus';
import Shares from '@/utils/server/shares'
import * as echarts from 'echarts'
import { markRaw } from 'vue'
// import user from '@/utils/server/oauth';
// import Server from '@/utils/server/def'

export default {
	name: 'hydayliyView',
	props: {
		msg: String
	},
	data() {
		return {
			userName: "",
			btnActive: 3,
			btnActive2: 1,
			cjl: {},
			yyqklines: [],
			usdcny: {}, // 美元对人民币
			isinit: false,
			isinit1: false,
			klineName2: [{
				label: "1天",
				value: 1,
			}, {
				label: "3天",
				value: 3,
			}, {
				label: "5天",
				value: 5,
			}, {
				label: "10天",
				value: 10,
			}, {
				label: "20天",
				value: 20,
			},
			{
				label: "60天",
				value: 60,
			},
			{
				label: "100天",
				value: 100,
			},
			{
				label: "300天",
				value: 300,
			}],
			klineName: [{
				label: "近1月",
				value: 1,
			}, {
				label: "近3月",
				value: 3,
			}, {
				label: "近6月",
				value: 6,
			}, {
				label: "近1年",
				value: 12,
			}, {
				label: "近2年",
				value: 24,
			},
			{
				label: "近3年",
				value: 36,
			},
			{
				label: "近5年",
				value: 60,
			},
			{
				label: "近10年",
				value: 120,
			}, {
				label: "成立以来",
				value: 1000,
			}],
			chartmlf: null,
			chartbx: null,
			chart_cjl: null,
			chartrzrq: null,
			chartm2: null,
			chartgdp: null,
			chartcpi: null,
			// chartzs: null,
			// charusdcny: null,
		}
	},
	async onLoad(options) {
		await this.loadData(options);
	},

	async mounted() {
		const { query } = useRoute();
		await this.loadData(query);
	},
	unmounted() { ////////////////////////////
	},
	querySelector() { },
	methods: {
		async confirmkline(val) {
			this.btnActive = val
			// 估值
			let out = await Shares.GetYyq(this.btnActive); // 获取联动图
			if (out != null) {
				this.yyqklines = out
				this.onloadEcharts()
				this.setEcharts()
			}
			let out8 = await Shares.GetUsdcny(this.btnActive)
			if (out8 != null) {
				this.usdcny = out8
				this.onloadklineBoll()
				// this.setklineBoll()
				//this.setkline()
			}
		},
		async confirmkline2(val) {
			let out9 = await Shares.GetShareCjl1(["sh000001", "sz399001"], this.btnActive2)
			if (out9 != null) {
				this.cjl = out9
				this.setCjl()
				//this.setkline()
			}
		},
		async setCjl() {
			var price = this.cjl.price
			var main = this.cjl.main
			this.chart_cjl.setOption({
				series: [{
					data: price,
				}, {
					data: main,
				}]
			});
		},
		/**
		 * 初始化
		 */

		async loadData(options) {

			let userName = options?.username
			if (userName != null && userName.length > 0) {
				this.userName = userName;
				localStorage.setItem("user-token", this.userName)
			} else {
				this.userName = localStorage.getItem("user-token")
			}
			// 估值
			let out = await Shares.GetYyq(this.btnActive); // 获取联动图
			if (out != null) {
				this.yyqklines = out
				this.onloadEcharts()
				this.setEcharts()
			}

			let out8 = await Shares.GetUsdcny(this.btnActive)
			if (out8 != null) {
				this.usdcny = out8
				this.onloadklineBoll()
				// this.setklineBoll()
				//this.setkline()
			}

			let out9 = await Shares.GetShareCjl1(["sh000001", "sz399001"], this.btnActive2)
			if (out9 != null) {
				this.cjl = out9
				this.onloadcjl()
				//this.setkline()
			}
			let that = this
			window.onresize = function () {
				that.chartmlf?.resize()
				that.chartbx?.resize()
				that.chartrzrq?.resize()
				that.chartm2?.resize()
				that.chartgdp?.resize()
				that.chartcpi?.resize()
				that.chartzs?.resize()
				that.chart_cjl?.resize()
				that.charusdcny?.resize()

			}

		},
		async searchCode() {
			this.sharesInfo.name = this.info.name;
			this.sharesInfo.sampleCode = this.info.simpleCode;
			this.sharesInfo.price = this.info.price;
			let percent = this.info.percent;
			this.sharesInfo.color0 = this.info.color
			if (percent > 0) {
				// this.sharesInfo.priceadd = "+"+data[31];
				this.sharesInfo.percent = "+" + percent;
			} else {
				// this.sharesInfo.priceadd = data[31];
				this.sharesInfo.percent = percent;
			}
			// this.sharesInfo.jk = data[5];
			// this.sharesInfo.zs = data[4];
			// this.sharesInfo.zg = data[33];
			// this.sharesInfo.zdf = data[32];
			// this.sharesInfo.hs =   data[38];
			// this.sharesInfo.zf =   data[43];
			// this.sharesInfo.pe =   data[53];
			// this.sharesInfo.pettm =   data[39];
			// this.sharesInfo.pb =   data[46];
			// this.sharesInfo.gxl =   data[64];
			// this.sharesInfo.zsz =   this.getSamplePrice(data[45]*100000000);
			// this.sharesInfo.ltsz =   this.getSamplePrice(data[44]*100000000);
			// this.sharesInfo.cje =   this.getSamplePrice(data[37]*10000);
			// this.sharesInfo.cjl = this.getSamplePrice(parseFloat(data[36])*0.01)

			// this.sharesInfo.zd =   data[42];
			// this.sharesInfo.zdc =await this.GetColor(this.userInfo.rg, data[42] -this.sharesInfo.zs)

			// this.sharesInfo.jkc = await this.GetColor(this.userInfo.rg, data[5]-data[4])
			// this.sharesInfo.zgc = await this.GetColor(this.userInfo.rg, data[33]-this.sharesInfo.zs)
		},
		async GetColor(rg, percent) {
			if (rg && percent >= 0) { // 默认红色
				return "#ff0000"
			}

			if (!rg && percent <= 0) { // 默认红色
				return "#ff0000"
			}

			return "#1AAD19" // 绿色
		},
		async onloadEcharts() {

			if (this.isinit == true) {
				return
			}
			this.isinit = true
			var that = this;
			this.chartmlf = markRaw(echarts.init(this.$refs.chartmlf));
			var option_chartmlf = {
				tooltip: {
					trigger: 'axis',
				},
				legend: {
					left: "0%",
					itemHeight: 5,
					itemWidth: 14,
					data: [{
						name: 'MLF',
						icon: 'rect'
					},
					{
						name: '当日净投放',
						icon: 'rect'
					},
					]
				},
				grid: [{
					top: '5%',
					left: '2%',
					right: '4%',
					bottom: '10%',
					containLabel: false
				}, {
					top: '2%',
					left: '2%',
					right: '4%',
					bottom: '10%',
					top: '40%',
					containLabel: false
				}],
				xAxis: [{
					axisTick: {
						show: false
					},
					axisLine: { // x轴是否显示
						show: false
					},
					axisLabel: {
						show: true
					},
					type: 'category',
					boundaryGap: false,
					gridIndex: 0,
					data: this.yyqklines.mlf.times
				}, {
					axisTick: {
						show: false
					},
					axisLine: { // x轴是否显示
						show: false
					},
					axisLabel: {
						show: true
					},
					type: 'category',
					boundaryGap: false,
					gridIndex: 1,
					data: this.yyqklines.mlf.times
				},],
				yAxis: [{
					type: 'value',
					scale: true,
					splitLine: {
						show: false
					},
					show: false, // y轴是否显示
					gridIndex: 0,
					// min: 0,     //设置左侧y轴最小刻度
					// max: 100     //设置左侧y轴最大刻度
				}, {
					type: 'value',
					scale: true,
					splitLine: {
						show: false
					},
					show: false, // y轴是否显示
					gridIndex: 1,
					// min: 0,     //设置左侧y轴最小刻度
					// max: 100     //设置左侧y轴最大刻度
				}],
				series: [{
					name: 'MLF',
					type: 'line',
					symbol: "none",
					data: this.yyqklines.mlf.data1,
					xAxisIndex: 0,
					yAxisIndex: 0,
					tooltip: {
						valueFormatter: function (value) {
							return value + '亿';
						}
					},
					lineStyle: { //标线的样式
						normal: {
							opacity: 0.6,
							color: "#0000ff",
							// width: 1
						}
					},
					markArea: {
						itemStyle: {
							color: 'rgba(215, 216, 217, 1.0)',
						},
						data: [
							[{
								name: '未来日期',
								xAxis: this.yyqklines.left
							},
							{
								xAxis: this.yyqklines.right
							}
							],
						]
					},
				},
				{
					name: '当日净投放',
					type: 'bar',
					symbol: "none",
					data: this.yyqklines.mlf.data2,
					xAxisIndex: 1,
					yAxisIndex: 1,
					tooltip: {
						valueFormatter: function (value) {
							return value + '亿';
						}
					},
					itemStyle: {
						normal: {
							color: function (params) {
								var colorList;
								if (params.data >= 0) {
									colorList = that.yyqklines.colorr;
								} else {
									colorList = that.yyqklines.colorg;
								}
								return colorList;
							},
						}
					},
				}
				]
			};
			this.chartmlf.setOption(option_chartmlf, true)


			this.chartbx = markRaw(echarts.init(this.$refs.chartbx));
			var option_chartbx = {
				tooltip: {
					trigger: 'axis',
				},
				legend: {
					left: "0%",
					itemHeight: 5,
					itemWidth: 14,
					data: [{
						name: '累计净流入',
						icon: 'rect'
					}, {
						name: '20日线',
						icon: 'rect'
					},
					{
						name: '当日净买入',
						icon: 'rect'
					}
					]
				},
				grid: [{
					top: '2%',
					left: '2%',
					right: '4%',
					bottom: '10%',
					containLabel: false
				}, {
					top: '2%',
					left: '2%',
					right: '4%',
					bottom: '10%',
					top: '60%',
					containLabel: false
				}],
				xAxis: [{
					axisTick: {
						show: false
					},
					axisLine: { // x轴是否显示
						show: false
					},
					axisLabel: {
						show: true
					},
					type: 'category',
					boundaryGap: false,
					gridIndex: 0,
					data: this.yyqklines.bx.times
				}, {
					axisTick: {
						show: false
					},
					axisLine: { // x轴是否显示
						show: false
					},
					axisLabel: {
						show: true
					},
					type: 'category',
					boundaryGap: false,
					gridIndex: 1,
					data: this.yyqklines.bx.times
				}],
				yAxis: [{
					type: 'value',
					scale: true,
					splitLine: {
						show: false
					},
					show: false, // y轴是否显示
					gridIndex: 0,
					// min: 0,     //设置左侧y轴最小刻度
					// max: 100     //设置左侧y轴最大刻度
				}, {
					type: 'value',
					scale: true,
					splitLine: {
						show: false
					},
					show: false, // y轴是否显示
					gridIndex: 1,
					// min: 0,     //设置左侧y轴最小刻度
					// max: 100     //设置左侧y轴最大刻度
				}],
				color: ["#ffce0a", "#adadad", '#ff0000'],
				series: [{
					name: '累计净流入',
					type: 'line',
					symbol: "none",
					data: this.yyqklines.bx.data1,
					xAxisIndex: 0,
					yAxisIndex: 0,
					tooltip: {
						valueFormatter: function (value) {
							return value.toFixed(2) + '亿';
						}
					},
					lineStyle: { //标线的样式
						normal: {
							opacity: 0.6,
							color: "#ffce0a",
							// width: 1
						}
					},

				}, {
					name: '20日线',
					type: 'line',
					symbol: "none",
					data: this.yyqklines.bx.data3,
					xAxisIndex: 0,
					yAxisIndex: 0,
					tooltip: {
						valueFormatter: function (value) {
							return value.toFixed(2) + '亿';
						}
					},
					lineStyle: { //标线的样式
						normal: {
							//opacity: 0.6,
							color: "#adadad",
							width: 1.2
						}
					},

				},
				{
					name: '当日净买入',
					type: 'bar',
					symbol: "none",
					data: this.yyqklines.bx.data2,
					xAxisIndex: 1,
					yAxisIndex: 1,
					tooltip: {
						valueFormatter: function (value) {
							return value.toFixed(2) + '亿';
						}
					},
					itemStyle: {
						normal: {
							color: function (params) {
								var colorList;
								if (params.data >= 0) {
									colorList = that.yyqklines.colorr;
								} else {
									colorList = that.yyqklines.colorg;
								}
								return colorList;
							},
						}
					},
				}
				]
			};
			this.chartbx.setOption(option_chartbx, true)

			this.chartrzrq = markRaw(echarts.init(this.$refs.chartrzrq));
			var option_chartrzrq = {
				tooltip: {
					trigger: 'axis',
				},
				color: ['#0000ff', "#55007f", "#ffda08"],
				legend: {
					left: "0%",
					itemHeight: 5,
					itemWidth: 14,
					data: [{
						name: '融资余额',
						icon: 'rect'
					},
					{
						name: '融券余额',
						icon: 'rect'
					}, {
						name: '融资净买入',
						icon: 'rect'
					},
					]
				},
				grid: [{
					top: '2%',
					left: '2%',
					right: '4%',
					bottom: '10%',
					containLabel: false
				}, {
					top: '2%',
					left: '2%',
					right: '4%',
					bottom: '10%',
					containLabel: false
				}, {
					top: '2%',
					left: '2%',
					right: '4%',
					bottom: '10%',
					top: '60%',
					containLabel: false
				}],
				xAxis: [{
					axisTick: {
						show: false
					},
					axisLine: { // x轴是否显示
						show: false
					},
					axisLabel: {
						show: true
					},
					type: 'category',
					boundaryGap: false,
					gridIndex: 0,
					data: this.yyqklines.rzrq.times
				}, {
					axisTick: {
						show: false
					},
					axisLine: { // x轴是否显示
						show: false
					},
					axisLabel: {
						show: true
					},
					type: 'category',
					boundaryGap: false,
					gridIndex: 1,
					data: this.yyqklines.rzrq.times
				}, {
					axisTick: {
						show: false
					},
					axisLine: { // x轴是否显示
						show: false
					},
					axisLabel: {
						show: true
					},
					type: 'category',
					boundaryGap: false,
					gridIndex: 2,
					data: this.yyqklines.rzrq.times
				}],
				yAxis: [{
					type: 'value',
					scale: true,
					splitLine: {
						show: false
					},
					show: false, // y轴是否显示
					gridIndex: 0,
					// min: 0,     //设置左侧y轴最小刻度
					// max: 100     //设置左侧y轴最大刻度
				}, {
					type: 'value',
					scale: true,
					splitLine: {
						show: false
					},
					show: false, // y轴是否显示
					gridIndex: 1,
					// min: 0,     //设置左侧y轴最小刻度
					// max: 100     //设置左侧y轴最大刻度
				}, {
					type: 'value',
					scale: true,
					splitLine: {
						show: false
					},
					show: false, // y轴是否显示
					gridIndex: 2,
					// min: 0,     //设置左侧y轴最小刻度
					// max: 100     //设置左侧y轴最大刻度
				}],
				series: [{
					name: '融资余额',
					type: 'line',
					symbol: "none",
					data: this.yyqklines.rzrq.data1,
					xAxisIndex: 0,
					yAxisIndex: 0,
					tooltip: {
						valueFormatter: function (value) {
							return value.toFixed(2) + '亿';
						}
					},
					lineStyle: { //标线的样式
						normal: {
							opacity: 0.6,
							color: "#0000ff",
							// width: 1
						}
					},

				}, {
					name: '融券余额',
					type: 'line',
					symbol: "none",
					data: this.yyqklines.rzrq.data2,
					xAxisIndex: 1,
					yAxisIndex: 1,
					tooltip: {
						valueFormatter: function (value) {
							return value.toFixed(2) + '亿';
						}
					},
					lineStyle: { //标线的样式
						normal: {
							opacity: 0.6,
							color: "#55007f",
							// width: 1
						}
					},

				},
				{
					name: '融资净买入',
					type: 'bar',
					symbol: "none",
					data: this.yyqklines.rzrq.data3,
					xAxisIndex: 2,
					yAxisIndex: 2,
					tooltip: {
						valueFormatter: function (value) {
							return value.toFixed(2) + '亿';
						}
					},
					itemStyle: {
						normal: {
							color: function (params) {
								var colorList;
								if (params.data >= 0) {
									colorList = that.yyqklines.colorr;
								} else {
									colorList = that.yyqklines.colorg;
								}
								return colorList;
							},
						}
					},
				}
				]
			};
			this.chartrzrq.setOption(option_chartrzrq, true)

			this.chartm2 = markRaw(echarts.init(this.$refs.chartm2));
			var option_chartm2 = {
				tooltip: {
					trigger: 'axis',
					backgroundColor: "#e69900",
					borderWidth: 0,
					formatter: function (params) {
						return params[0].axisValue
					},
				},
				color: ["#0000ff", "#e69900", '#ff0000'],
				legend: {
					left: "0%",
					itemHeight: 5,
					itemWidth: 14,
					data: [{
						name: 'M2',
						icon: 'rect'
					}, {
						name: 'M1',
						icon: 'rect'
					},
					{
						name: '新增信贷',
						icon: 'rect'
					},
					]
				},
				grid: {
					top: '2%',
					left: '2%',
					right: '2%',
					bottom: '2%',
					containLabel: false
				},
				xAxis: [{
					axisTick: {
						show: false
					},
					axisLine: { // x轴是否显示
						show: false
					},
					axisLabel: {
						show: false
					},
					type: 'category',
					boundaryGap: false,
					data: this.yyqklines.m2.times
				}, {
					axisTick: {
						show: false
					},
					axisLine: { // x轴是否显示
						show: false
					},
					axisLabel: {
						show: false
					},
					type: 'category',
					boundaryGap: false,
					data: this.yyqklines.m2.times
				}, {
					axisTick: {
						show: false
					},
					axisLine: { // x轴是否显示
						show: false
					},
					axisLabel: {
						show: true
					},
					type: 'category',
					boundaryGap: false,
					data: this.yyqklines.xzxd.times
				},],
				yAxis: [{
					type: 'value',
					scale: true,
					splitLine: {
						show: false
					},
					show: false, // y轴是否显示
					// min: 0,     //设置左侧y轴最小刻度
					// max: 100     //设置左侧y轴最大刻度
				}, {
					type: 'value',
					scale: true,
					splitLine: {
						show: false
					},
					show: false, // y轴是否显示
					// min: 0,     //设置左侧y轴最小刻度
					// max: 100     //设置左侧y轴最大刻度
				}, {
					type: 'value',
					scale: true,
					splitLine: {
						show: false
					},
					show: false, // y轴是否显示
					// min: 0,     //设置左侧y轴最小刻度
					// max: 100     //设置左侧y轴最大刻度
				}],
				series: [{
					name: 'M2',
					type: 'line',
					symbol: "none",
					data: this.yyqklines.m2.data1,
					yAxisIndex: 0,
					tooltip: {
						valueFormatter: function (value) {
							return value + '亿';
						}
					},
					lineStyle: { //标线的样式
						normal: {
							opacity: 0.6,
							color: "#0000ff",
							width: 1
						}
					},

				}, {
					name: 'M1',
					type: 'line',
					symbol: "none",
					data: this.yyqklines.m2.data2,
					yAxisIndex: 1,
					tooltip: {
						valueFormatter: function (value) {
							return value + '亿';
						}
					},
					lineStyle: { //标线的样式
						normal: {
							opacity: 0.6,
							color: "#e69900",
							width: 1
						}
					},

				},
				{
					name: '新增信贷',
					type: 'line',
					symbol: "none",
					data: this.yyqklines.xzxd.data1,
					yAxisIndex: 2,
					tooltip: {
						valueFormatter: function (value) {
							return value + '亿';
						}
					},
					lineStyle: { //标线的样式
						normal: {
							opacity: 0.6,
							color: "#ff0000",
							width: 1
						}
					},
				}
				]
			};
			this.chartm2.setOption(option_chartm2, true)

			this.chartcpi = markRaw(echarts.init(this.$refs.chartcpi));
			var option_chartcpi = {
				tooltip: {
					trigger: 'axis',
					backgroundColor: "#e69900",
					borderWidth: 0,
					formatter: function (params) {
						return params[0].axisValue
					},
				},
				color: ["#0000ff", '#ff0000', "#ffaa00", "#ff00ff"],
				legend: {
					left: "0%",
					itemHeight: 5,
					itemWidth: 14,
					data: [{
						name: 'CPI',
						icon: 'rect'
					},
					{
						name: 'PPI',
						icon: 'rect'
					}, {
						name: 'PMI(制造业)',
						icon: 'rect'
					}, {
						name: 'PMI(非制造业)',
						icon: 'rect'
					},
					]
				},
				grid: {
					top: '2%',
					left: '2%',
					right: '2%',
					bottom: '0%',
					containLabel: false
				},
				xAxis: [{
					axisTick: {
						show: false
					},
					axisLine: { // x轴是否显示
						show: false
					},
					axisLabel: {
						show: true
					},
					type: 'category',
					boundaryGap: false,
					data: this.yyqklines.cpi.times
				}],
				yAxis: [{
					type: 'value',
					scale: true,
					splitLine: {
						show: false
					},
					show: false, // y轴是否显示
					// min: 0,     //设置左侧y轴最小刻度
					// max: 100     //设置左侧y轴最大刻度
				}, {
					type: 'value',
					scale: true,
					splitLine: {
						show: false
					},
					show: false, // y轴是否显示
					// min: 0,     //设置左侧y轴最小刻度
					// max: 100     //设置左侧y轴最大刻度
				}, {
					type: 'value',
					scale: true,
					splitLine: {
						show: false
					},
					show: false, // y轴是否显示
					// min: 0,     //设置左侧y轴最小刻度
					// max: 100     //设置左侧y轴最大刻度
				}],
				series: [{
					name: 'CPI',
					type: 'line',
					symbol: "none",
					data: this.yyqklines.cpi.data1,
					yAxisIndex: 0,
					tooltip: {
						valueFormatter: function (value) {
							return value;
						}
					},
					lineStyle: { //标线的样式
						normal: {
							opacity: 0.6,
							color: "#0000ff",
							width: 1
						}
					},

				},
				{
					name: 'PPI',
					type: 'line',
					symbol: "none",
					data: this.yyqklines.ppi.data1,
					yAxisIndex: 1,
					tooltip: {
						valueFormatter: function (value) {
							return value;
						}
					},
					lineStyle: { //标线的样式
						normal: {
							opacity: 0.6,
							color: "#ff0000",
							width: 1
						}
					},
				},
				{
					name: 'PMI(制造业)',
					type: 'line',
					symbol: "none",
					data: this.yyqklines.pmi.data1,
					yAxisIndex: 2,
					tooltip: {
						valueFormatter: function (value) {
							return value;
						}
					},
					lineStyle: { //标线的样式
						normal: {
							opacity: 0.6,
							color: "#ffaa00",
							width: 1
						}
					},
				},
				{
					name: 'PMI(非制造业)',
					type: 'line',
					symbol: "none",
					data: this.yyqklines.pmi.data2,
					yAxisIndex: 2,
					tooltip: {
						valueFormatter: function (value) {
							return value;
						}
					},
					lineStyle: { //标线的样式
						normal: {
							opacity: 0.6,
							color: "#ff00ff",
							width: 1
						}
					},
				}
				]
			};
			this.chartcpi.setOption(option_chartcpi, true)

			this.chartgdp = markRaw(echarts.init(this.$refs.chartgdp));
			var option_chartgdp = {
				tooltip: {
					trigger: 'axis',
					backgroundColor: "#e69900",
					borderWidth: 0,
					formatter: function (params) {
						return params[0].axisValue
					},
				},
				color: ["#0000ff", '#ff0000'],
				legend: {
					left: "0%",
					itemHeight: 5,
					itemWidth: 14,
					data: [{
						name: 'GDP',
						icon: 'rect'
					},
					{
						name: '企业信心',
						icon: 'rect'
					}
					]
				},
				grid: {
					top: '2%',
					left: '2%',
					right: '2%',
					bottom: '0%',
					containLabel: false
				},
				xAxis: [{
					axisTick: {
						show: false
					},
					axisLine: { // x轴是否显示
						show: false
					},
					axisLabel: {
						show: true
					},
					type: 'category',
					boundaryGap: false,
					data: this.yyqklines.gdp.times
				}],
				yAxis: [{
					type: 'value',
					scale: true,
					splitLine: {
						show: false
					},
					show: false, // y轴是否显示
					// min: 0,     //设置左侧y轴最小刻度
					// max: 100     //设置左侧y轴最大刻度
				}, {
					type: 'value',
					scale: true,
					splitLine: {
						show: false
					},
					show: false, // y轴是否显示
					// min: 0,     //设置左侧y轴最小刻度
					// max: 100     //设置左侧y轴最大刻度
				}],
				series: [{
					name: 'GDP',
					type: 'line',
					symbol: "none",
					data: this.yyqklines.gdp.data1,
					yAxisIndex: 0,
					tooltip: {
						valueFormatter: function (value) {
							return value + '亿';;
						}
					},
					lineStyle: { //标线的样式
						normal: {
							opacity: 0.6,
							color: "#0000ff",
							width: 1
						}
					},
				}, {
					name: '企业信心',
					type: 'line',
					symbol: "none",
					data: this.yyqklines.qyjqzs.data1,
					yAxisIndex: 1,
					tooltip: {
						valueFormatter: function (value) {
							return value;
						}
					},
					lineStyle: { //标线的样式
						normal: {
							opacity: 0.6,
							color: "#ff0000",
							width: 1
						}
					},
				}]
			};
			this.chartgdp.setOption(option_chartgdp, true)

			this.chartzs = markRaw(echarts.init(this.$refs.chartzs));
			var option_chartzs = {
				tooltip: {
					trigger: 'axis',
					backgroundColor: "#e69900",
					borderWidth: 0,
					formatter: function (params) {
						return params[0].axisValue
					}
				},
				color: ["#0000ff", '#ff0000', "#ffaa00", "#ff00ff"],
				legend: {
					left: "0%",
					itemHeight: 5,
					itemWidth: 14,
					data: [{
						name: '工业增加值',
						icon: 'rect'
					},
					{
						name: '消费者信心',
						icon: 'rect'
					}, {
						name: '海关进出口',
						icon: 'rect'
					}, {
						name: '外商投资',
						icon: 'rect'
					},
					]
				},
				grid: {
					top: '2%',
					left: '2%',
					right: '2%',
					bottom: '0%',
					containLabel: false
				},
				xAxis: [{
					axisTick: {
						show: false
					},
					axisLine: { // x轴是否显示
						show: false
					},
					axisLabel: {
						show: true
					},
					type: 'category',
					boundaryGap: false,
					data: this.yyqklines.fdi.times
				}],
				yAxis: [{
					type: 'value',
					scale: true,
					splitLine: {
						show: false
					},
					show: false, // y轴是否显示
					// min: 0,     //设置左侧y轴最小刻度
					// max: 100     //设置左侧y轴最大刻度
				}, {
					type: 'value',
					scale: true,
					splitLine: {
						show: false
					},
					show: false, // y轴是否显示
					// min: 0,     //设置左侧y轴最小刻度
					// max: 100     //设置左侧y轴最大刻度
				}, {
					type: 'value',
					scale: true,
					splitLine: {
						show: false
					},
					show: false, // y轴是否显示
					// min: 0,     //设置左侧y轴最小刻度
					// max: 100     //设置左侧y轴最大刻度
				}, {
					type: 'value',
					scale: true,
					splitLine: {
						show: false
					},
					show: false, // y轴是否显示
					// min: 0,     //设置左侧y轴最小刻度
					// max: 100     //设置左侧y轴最大刻度
				}],
				series: [{
					name: '工业增加值',
					type: 'line',
					symbol: "none",
					data: this.yyqklines.gyzjz.data1,
					yAxisIndex: 0,
					tooltip: {
						valueFormatter: function (value) {
							return value.toFixed(2) + "%";
						}
					},
					lineStyle: { //标线的样式
						normal: {
							opacity: 0.6,
							color: "#0000ff",
							width: 1
						}
					},

				},
				{
					name: '消费者信心',
					type: 'line',
					symbol: "none",
					data: this.yyqklines.xfzxx.data1,
					yAxisIndex: 1,
					tooltip: {
						valueFormatter: function (value) {
							return value;
						}
					},
					lineStyle: { //标线的样式
						normal: {
							opacity: 0.6,
							color: "#ff0000",
							width: 1
						}
					},
				},
				{
					name: '海关进出口',
					type: 'line',
					symbol: "none",
					data: this.yyqklines.hgjck.data1,
					yAxisIndex: 2,
					tooltip: {
						valueFormatter: function (value) {
							return value.toFixed(2) + '亿$';
						}
					},
					lineStyle: { //标线的样式
						normal: {
							opacity: 0.6,
							color: "#ffaa00",
							width: 1
						}
					},
				},
				{
					name: '外商投资',
					type: 'line',
					symbol: "none",
					data: this.yyqklines.fdi.data1,
					yAxisIndex: 2,
					tooltip: {
						valueFormatter: function (value) {
							return value.toFixed(2) + '亿';
						}
					},
					lineStyle: { //标线的样式
						normal: {
							opacity: 0.6,
							color: "#ff00ff",
							width: 1
						}
					},
				}
				]
			};
			this.chartzs.setOption(option_chartzs, true)
		},
		async setEcharts() {
			var that = this;
			this.chartmlf.setOption({
				xAxis: [{
					data: this.yyqklines.mlf.times
				}, {
					data: this.yyqklines.mlf.times
				}],
				series: [{
					data: this.yyqklines.mlf.data1,
				},
				{
					data: this.yyqklines.mlf.data2,
				}
				]
			});
			this.chartbx.setOption({
				xAxis: [{
					data: this.yyqklines.bx.times
				}, {
					data: this.yyqklines.bx.times
				},],
				series: [{
					data: this.yyqklines.bx.data1,
				},
				{
					data: this.yyqklines.bx.data3,
				},
				{
					data: this.yyqklines.bx.data2,
				}
				]
			});
			this.chartrzrq.setOption({
				xAxis: [{
					data: this.yyqklines.rzrq.times
				}, {
					data: this.yyqklines.rzrq.times
				}, {
					data: this.yyqklines.rzrq.times
				}],
				series: [{
					data: this.yyqklines.rzrq.data1,
				}, {
					data: this.yyqklines.rzrq.data2,
				},
				{
					data: this.yyqklines.rzrq.data3,
				}
				]
			});
			this.chartm2.setOption({
				xAxis: [{
					data: this.yyqklines.m2.times
				}, {
					data: this.yyqklines.m2.times
				}, {
					data: this.yyqklines.xzxd.times
				},],
				series: [{
					data: this.yyqklines.m2.data1,
				}, {
					data: this.yyqklines.m2.data2,
				},
				{
					data: this.yyqklines.xzxd.data1,
				}
				]
			});
			this.chartcpi.setOption({
				xAxis: [{
					data: this.yyqklines.cpi.times
				}],
				series: [{
					data: this.yyqklines.cpi.data1,
				},
				{
					data: this.yyqklines.ppi.data1,
				},
				{
					data: this.yyqklines.pmi.data1,
				},
				{
					data: this.yyqklines.pmi.data2,
				}
				]
			});
			this.chartgdp.setOption({
				xAxis: [{
					data: this.yyqklines.gdp.times
				}],
				series: [{
					data: this.yyqklines.gdp.data1,
				}, {
					data: this.yyqklines.qyjqzs.data1,
				}]
			});
			this.chartzs.setOption({
				xAxis: [{
					data: this.yyqklines.fdi.times
				}],
				series: [{
					data: this.yyqklines.gyzjz.data1,
				},
				{
					data: this.yyqklines.xfzxx.data1,
				},
				{
					data: this.yyqklines.hgjck.data1,
				},
				{
					data: this.yyqklines.fdi.data1,
				}
				]
			});

		},
		getSamplePrice(price) {
			var tag = ""
			if (price < 0) {
				tag = "-"
				price = Math.abs(price)
			}

			if (price < 10000) { // 一万以下，直接输出数字
				return tag + Math.floor(price)
			}

			if (price < 100000000) { // 一亿以下，直接输出万级别,保留2位小数
				return tag + (price * 0.0001).toFixed(2) + "万"
			}

			// 到亿
			return tag + (price * 0.00000001).toFixed(2) + "亿"
		},
		async onloadklineBoll() {
			if (this.isinit1 == true) {
				this.setklineBoll()
				return
			}
			this.isinit1 = true
			var that = this;
			var data0 = this.splitData(this.usdcny.list)
			var upColor = this.usdcny.rColor
			var downColor = this.usdcny.gColor
			this.charusdcny = markRaw(echarts.init(this.$refs.charusdcny));
			var option_charusdcny = {
				tooltip: {
					trigger: 'axis',
					axisPointer: {
						type: 'cross'
					},
					formatter: function (params) { //修改鼠标划过显示为中文
						var out = '';
						if (params[0].componentSubType === 'candlestick') {
							out += params[0].name + '<br/>' +
								params[0].marker + '开盘:' + params[0].data[1] +
								'<br/>' +
								params[0].marker + '收盘:' + params[0].data[2] +
								'<br/>' +
								params[0].marker + '20日:' + params[0].data[6] +
								'<br/>'
						}

						return out;
					}
				},
				legend: {
					left: "0%",
					itemHeight: 5,
					itemWidth: 14,
					data: [{
						name: '日K',
						icon: 'rect'
					}, {
						name: '20日',
						icon: 'rect'
					}
					]
				},
				grid: {
					top: '0%',
					left: '2%',
					right: '2%',
					bottom: '10%',
					containLabel: false
				},
				xAxis: {
					type: 'category',
					data: data0.categoryData,
					scale: true,
					boundaryGap: true,
					axisLine: {
						show: false,
					},
					axisLabel: { //label文字设置
						show: true
					},
					axisTick: {
						show: false,
					},
					splitLine: {
						show: false,
						lineStyle: {
							color: '#3a3a3e'
						}
					},
					show: true,

				},
				yAxis: {
					scale: true,
					splitArea: {
						show: false
					},
					splitLine: {
						show: false
					},
					axisLine: {
						onZero: false
					},
					show: false, // y轴是否显示
				},
				// animation: false, //禁止动画效果
				series: [
					{
						name: '日K',
						type: 'candlestick',
						data: data0.values,
						itemStyle: {
							color: upColor,
							color0: downColor,
							borderColor: upColor,
							borderColor0: downColor
						},
					},
					{
						name: '20日',
						type: 'line',
						color: "#adadad",
						symbol: "none",
						data: that.calculateMA(1, data0.values),
						smooth: true,
						lineStyle: {
							opacity: 0.5
						}
					},
				]
			};
			this.charusdcny.setOption(option_charusdcny, true)

		},
		async setklineBoll() {
			var that = this;
			var data0 = this.splitData(this.usdcny.list)
			var upColor = this.usdcny.rColor
			var downColor = this.usdcny.gColor
			this.charusdcny.setOption({
				xAxis: {
					data: data0.categoryData,
				},
				series: [{
					data: data0.values,
					itemStyle: {
						color: upColor,
						color0: downColor,
						borderColor: upColor,
						borderColor0: downColor
					},
				},
				{
					data: that.calculateMA(1, data0.values),
				},
				]
			});
		},
		async onloadcjl() {
			var timeArr = new Array();
			timeArr.push('09:30');
			timeArr.concat(this.getNextTime('09:30', '11:30', 1, timeArr));
			timeArr.push('13:00');
			timeArr.concat(this.getNextTime('13:00', '15:00', 1, timeArr));
			//this.title = this.zlmain.tm
			var price = this.cjl.price
			var main = this.cjl.main
			var that = this;
			this.chart_cjl = markRaw(echarts.init(this.$refs.chart_cjl))
			var option_chart_cjl = {
				tooltip: { //弹框指示器
					trigger: 'axis',
				},
				color: ["#0055ff", "#b0b9ca"],
				legend: {
					left: "0%",
					itemHeight: 5,
					itemWidth: 14,
					data: [{
						name: "当前",
						icon: 'rect'
					}, {
						name: "平均",
						icon: 'rect'
					}]
				},

				grid: [{
					top: '2%',
					left: '2%',
					right: '2%',
					bottom: '12%',
					containLabel: false
				}],
				xAxis: [{
					axisTick: {
						show: false
					},
					axisLine: { // x轴是否显示
						show: false
					},
					axisLabel: {
						show: true
					},
					type: 'category',
					boundaryGap: false,
					data: timeArr
				}, {
					axisTick: {
						show: false
					},
					axisLine: { // x轴是否显示
						show: false
					},
					axisLabel: {
						show: true
					},
					type: 'category',
					boundaryGap: false,
					data: timeArr
				}],
				yAxis: [{
					axisLabel: {
						margin: 4,
					},
					type: 'value',
					scale: true,
					splitLine: {
						show: false
					},
					show: false
					// min: 0,     //设置左侧y轴最小刻度
					// max: 100     //设置左侧y轴最大刻度
				}, {
					axisLabel: {
						margin: 4,
					},
					type: 'value',
					scale: true,
					splitLine: {
						show: false
					},
					show: false
					// min: 0,     //设置左侧y轴最小刻度
					// max: 100     //设置左侧y轴最大刻度
				}],
				// animation: false, //禁止动画效果
				// backgroundColor: "#100C2A",
				// blendMode: 'source-over',
				series: [{
					name: '当前',
					type: 'line',
					symbol: "none",
					smooth: true,
					data: price,
					xAxisIndex: 0,
					yAxisIndex: 0,
					lineStyle: { //标线的样式
						normal: {
							opacity: 1,
							color: "#0055ff",
							width: 1
						}
					},
					tooltip: {
						valueFormatter: function (value) {
							return that.getSamplePrice(value);
						}
					},
				}, {
					name: '平均',
					type: 'line',
					symbol: "none",
					smooth: true,
					data: main,
					xAxisIndex: 0,
					yAxisIndex: 0,
					lineStyle: { //标线的样式
						normal: {
							opacity: 1,
							color: "#b0b9ca",
							width: 2
						}
					},
					tooltip: {
						valueFormatter: function (value) {
							return that.getSamplePrice(value);
						}
					},
					label: {
						show: true,
						formatter: function (value) {
							return that.getSamplePrice(value.data);
						}
					},
				}]
			}
			this.chart_cjl.setOption(option_chart_cjl, true)

		},
		splitData(rawData) {
			const categoryData = [];
			const values = [];
			for (var i = 0; i < rawData?.length; i++) {
				categoryData.push(rawData[i].splice(0, 1)[0]);
				values.push(rawData[i]);
			}
			return {
				categoryData: categoryData,
				values: values
			};
		},
		calculateMA(dayCount, data) {
			var result = [];
			for (var i = 0, len = data.length; i < len; i++) {
				if (data[i].length < 4 + dayCount) {
					result.push('-');
				} else {
					result.push(data[i][4 + dayCount]);
				}
			}
			return result;
		},
		addTimeStr(time, num) {
			var hour = time.split(':')[0];
			var mins = Number(time.split(':')[1]);
			var mins_un = parseInt((mins + num) / 60);
			var hour_un = parseInt((Number(hour) + mins_un) / 24);
			if (mins_un > 0) {
				if (hour_un > 0) {
					var tmpVal = ((Number(hour) + mins_un) % 24) + "";
					hour = tmpVal.length > 1 ? tmpVal : '0' + tmpVal; //判断是否是一位
				} else {
					var tmpVal = Number(hour) + mins_un + "";
					hour = tmpVal.length > 1 ? tmpVal : '0' + tmpVal;
				}
				var tmpMinsVal = ((mins + num) % 60) + "";
				mins = tmpMinsVal.length > 1 ? tmpMinsVal : 0 + tmpMinsVal; //分钟数为 取余60的数
			} else {
				var tmpMinsVal = mins + num + "";
				mins = tmpMinsVal.length > 1 ? tmpMinsVal : '0' + tmpMinsVal; //不大于整除60
			}
			return hour + ":" + mins;
		},
		//获取增加指定分钟数的 数组  如 09:30增加2分钟  结果为 ['09:31','09:32']
		getNextTime(startTime, endTIme, offset, resultArr) {
			var result = this.addTimeStr(startTime, offset);
			resultArr.push(result);
			if (result == endTIme) {
				return resultArr;
			} else {
				return this.getNextTime(result, endTIme, offset, resultArr);
			}
		},
	}
}
</script>

<style lang="scss">
* {
	box-sizing: border-box;
}

.el-segmented {
	--el-segmented-item-selected-color: #000;
	--el-segmented-item-selected-bg-color: #aa55ff;
	--el-segmented-border-radius-base: 16px;
	--el-segmented-color: #fff;
	--el-segmented-bg-color: rgb(255, 119, 39);
	--el-segmented-item-hover-color: var(--el-text-color-primary);
	--el-segmented-item-hover-bg-color: #aa55ffb2;
	--el-segmented-item-active-bg-color: #aa55ffb2;
	--el-segmented-padding: 0px;
	// font-weight: bold;
}

body {
	// background-color: #eeeeee;
}

.switchbar0 {
	user-select: none;
	position: relative;
	background-color: #ea5404;
	// z-index: 99;
	// top: 8px;
	margin-right: 7px;
	left: 8px;
	min-height: 40px;

}

.switchbar {
	user-select: none;
	position: relative;
	background-color: #ea5404;
	top: 4px;
	max-width: 800px;
}

.switchbar2 {
	user-select: none;
	position: relative;
	background-color: #a39f9d;
	top: 3px;
}

.space {
	user-select: none;
	position: relative;
	// background-color: #15c3da;
	top: 3px;
	// margin-right: 7px;
	left: 0px;
	margin-left: 8px;
	padding-right: 7px;
	height: 100vh;
	display: flex;
	flex-direction: column;
	// width: 100%;
}

.split {
	flex: 1;

	.switchChart {
		display: flex;
		flex-direction: column;
		height: 30vh;
		width: 100vw;
	}

	.row {
		display: flex;
		justify-content: space-between;

	}
}
</style>
