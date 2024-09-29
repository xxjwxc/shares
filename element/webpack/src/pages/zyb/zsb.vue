<template>
  <!-- <el-switch v-model="value2" class="ml-2" @change="switchchange" /> -->

  <div id="header">
    <div class="left">
      <div class="name"> 指数榜 </div>
      <div class="zde" style="background-color: rgba(255, 0, 0, 0);color:#0000ff">基准时间:{{ startTm }}</div>
      <div class="prise" style="background-color: rgba(255, 0, 0, 0);color:#000000">显示全指:
        <el-switch v-model="showCd" @change="changeShowCd" />
      </div>
    </div>

  </div>

  <div id="root">
    <div class="left">
      <div class="switchbar0">
        <el-segmented class="switchbar" v-model="btnActive1" :options="klineName" @change='initClick1' block />
      </div>
      <div class="switchbar1">
        <div id="cjlchart" ref="cjlchart" style="height: 100%;" ></div>
      </div>
      <div class="switchbar2">
        <div id="hychart" ref="hychart" style="height: 100%;"></div>
      </div>
    </div>
    <div class="righ">
      <!-- <iframe class="infoIframe" :src="iframeUrl2" frameborder="0"></iframe> -->
      <h2 style="font-size: 20px;color: #aa00ff;">遴选指数</h2>

      <!-- <el-segmented v-model="value" :options="options" size="large" style="width: 100%;" /> -->
      <!-- @row-click="rowClick" -->
      <el-table @row-click="rowClick" :data="tsList1.list" style="width: 100%" height="400" :row-style="tableRowStyle">
        <el-table-column prop="name" label="指数名称" width="100">
          <template #default="{ row }">
            <span class="num-text">{{ row.ext }}</span>{{ row.name }}
          </template>
        </el-table-column>
        <el-table-column prop="sPrice" label="PE百分位" width="250" />
        <el-table-column prop="price" label="涨跌幅" />
      </el-table>
      <h2 style="font-size: 20px;color: #aa00ff;">行业指数</h2>
      <el-table @row-click="rowClick" :data="tsList2.list" style="width: 100%" height="400" :row-style="tableRowStyle">
        <el-table-column prop="name" label="指数名称" width="100">
          <template #default="{ row }">
            <span class="num-text">{{ row.ext }}</span>{{ row.name }}
          </template>
        </el-table-column>
        <el-table-column prop="sPrice" label="PE百分位" width="150" />
        <el-table-column prop="price" label="涨跌幅" />
      </el-table>
    </div>
  </div>
</template>

<script>
// import server from '@/utils/server/def.js'
import { useRoute, useRouter } from 'vue-router';
// import router from '@/router';
// import { ElMessageBox } from 'element-plus';
import * as echarts from 'echarts'
import Shares from '@/utils/server/shares'
import { markRaw } from 'vue'
// import user from '@/utils/server/oauth';
// import Server from '@/utils/server/def'
let router
export default {
  name: 'ZybView',
  props: {
    msg: String
  },
  data() {
    return {
      searchValue: "",
      iframekey: 0,
      userName: "",
      info: null,
      value: 1,
      options: [
        {
          label: '2024-06-28',
          value: 1,
        },
        {
          label: '5',
          value: 5,
        },
        {
          label: '10',
          value: 10,
        },
        {
          label: 20,
          value: 20,
        },
        {
          label: 100,
          value: 100,
        },
      ],
      klineName: [{
        label: "近1日",
        value: 1,
      }, {
        label: "近3日",
        value: 3,
      }, {
        label: "近5日",
        value: 5,
      }, {
        label: "近10日",
        value: 10,
      }, {
        label: "近20日",
        value: 20,
      },
      {
        label: "近30日",
        value: 30,
      },
      {
        label: "近60日",
        value: 60,
      },
      {
        label: "近100日",
        value: 100,
      },
      {
        label: "近200日",
        value: 200,
      },
      {
        label: "近500日",
        value: 500,
      }],
      btnActive1: 20,
      showCd: true,
      list: [],
      times: [],
      // tags: [],
      // tags1: [],
      legendData: [],
      seriesData: [],
      names: [],
      cjlseries: [],
      cjlmin: 0,
      cjlmax: 100,
      tsList1: [],
      tsList2: [],
      startTm: "",
      color: ['#dd4444', '#fec42c', '#9cb9c8', '#ffff00', '#5555ff', '#ff007f',
        '#aaffff', '#55007f', '#aaaa00', '#aaff00', '#555500', '#ff00ff',
        '#00ff7f', '#55ff7f'
      ],
      hychart: null,
      cjlchart: null,
    }
  },
  async onLoad(options) {
    await this.loadData(options);
  },

  async mounted() {
    const { query } = useRoute();
    router = useRouter()

    await this.loadData(query);
  },
  unmounted() { ////////////////////////////
    if (this.interval) {
      clearInterval(this.interval);
    }
  },
  querySelector() { },
  methods: {
    rowClick(val) {
      var url = `add?scode=${val.code}`
      const userid = this.userName;
      if (userid?.length > 0) {
        url += `&username=${userid}`
      }
      router.push(url)
    },
    initClick1(val) {
      this.loadData(null)
    },
    changeShowCd() {
      this.isinit1 = false
      this.list = []
      this.loadData();
      // -------- 分享
    },
    tableRowStyle({ row }) {
      return {
        color: row.color,
      }

    },

    async loadData(options) {
      let username = options?.username
      if (username != null && username.length > 0) {
        this.userName = username
        localStorage.setItem("user-token", username)
      }

      let group = await Shares.GetZybZsKline(this.btnActive1, this.showCd); // 获取分组信息
      if (group != null) {
        this.list = group.list
        this.times = group.times
        // this.tags = group.tags
        // this.tags1 = group.tags1
        this.legendData = []
        this.seriesData = []
        this.names = []
        this.cjlseries = []
        this.cjlmax = group.cjlmax
        this.cjlmin = group.cjlmin
        this.tsList1 = group.tsList1
        this.tsList2 = group.tsList2
        this.startTm = group.startTm
        // if (this.hyMmaddTag == 1) {
        // 	this.hyMmadd = this.tsList1
        // } else {
        // 	this.hyMmadd = this.tsList2
        // }

        let that = this;
        let i = 0;
        this.list.forEach(item => {
          that.legendData.push({
            name: item.name,
            icon: 'rect'
          });
          that.seriesData.push({
            name: item.name,
            type: 'line',
            symbol: "none",
            sampling: 'lttb',
            data: item.data,
            lineStyle: { //标线的样式
              normal: {
                opacity: 1,
                color: this.color[i],
                width: 1
              }
            },
            endLabel: {
              show: true,
              offset: [-40, -5],
              fontSize: 10,
              formatter: function (params) {
                return params.seriesName;
              }
            },
          })

          that.names.push(item.name);
          that.cjlseries.push({
            name: item.name,
            code: item.code,
            type: 'scatter',
            itemStyle: {
              opacity: 0.8,
              shadowBlur: 10,
              shadowOffsetX: 0,
              shadowOffsetY: 0,
              shadowColor: 'rgba(0,0,0,0.3)'
            },
            data: [item.cjl],
          })
          i++;
        })
        // Var aa=new Array(); //定义一维数组

        this.onloadZybEcharts()
        this.setZybEcharts()
      }
    },
    onloadZybEcharts() {
      if (this.isinit1 == true) {
        return
      }
      this.isinit1 = true
      var that = this;

      var option1 = {
        color: that.color,
        legend: {
          top: 0,
          data: that.names,
          textStyle: {
            fontSize: 4
          },
          show: false
        },
        // grid: {
        // 	left: '10%',
        // 	right: 150,
        // 	top: '18%',
        // 	bottom: '10%'
        // },
        grid: {
          top: '10%',
          left: '4%',
          right: '8%',
          bottom: '4%',
          containLabel: false
        },
        label: {
          show: true,
          position: 'inside',
          // color: '#666666',
          fontSize: 10,
          formatter: function (param) {
            return param.seriesName;
          },
          rich: {
            a: {
              color: '#2E6BE6',
              align: 'center'
            }
          }
        },
        tooltip: {
          // extraCssText: 'width:100rex;height:100rex;',
          backgroundColor: 'rgba(255,255,255,0.7)',
          formatter: function (param) {
            var value = param.value;
            return param.seriesName + '</br>' +
              'PE百分位：' + value[0].toFixed(2) + '%</br>' +
              '涨幅：' + value[1].toFixed(2) + '%</br>' +
              '成交量：' + value[2].toFixed(2) + '亿</br>' +
              '融资：' + value[3].toFixed(2) + '亿</br>'
          }
        },
        xAxis: {
          type: 'value',
          name: 'PE百分位',
          min: 'dataMin', // 最小值
          // nameGap: 16,
          nameTextStyle: {
            fontSize: 8
          },
          "axisLine": {
            "show": false //隐藏x轴线
          },
          "axisTick": {
            "show": false //隐藏x轴刻度
          },
          // max: 31,
          splitLine: {
            show: false
          },
          axisLabel: {
								formatter: function(value, index) {
									return value + '%';
								}
							},
        },
        yAxis: {
          type: 'value',
          name: '涨幅',
          min: 'dataMin', // 最小值
          // nameLocation: 'end',
          // nameGap: 20,
          nameTextStyle: {
            fontSize: 8
          },
          "axisLine": {
            "show": false //隐藏x轴线
          },
          "axisTick": {
            "show": false //隐藏x轴刻度
          },
          splitLine: {
            show: false
          },
          axisLabel: {
								formatter: function(value, index) {
									return value + '%';
								}
							},
        },
        visualMap: [{
          show: false,
          dimension: 2,
          min: that.cjlmin,
          max: that.cjlmax,
          calculable: true,
          inRange: {
            symbolSize: [0, 100]
          }
        }],
        series: that.cjlseries
      };
      this.cjlchart = markRaw(echarts.init(this.$refs.cjlchart));
      document.getElementById('cjlchart').removeAttribute('_echarts_instance_');
      this.cjlchart.setOption(option1)
      this.cjlchart.on("click", function (params) {
        var now = new Date()
        if (that.clickIndex == params.seriesIndex) {
          var timeDiff = Math.abs(now - that.clickTm)
          if (timeDiff < 1000) {
            var code = ""
            for (var i = 0; i < that.cjlseries.length; i++) {
              if (that.cjlseries[i].name == params.seriesName) {
                code = that.cjlseries[i].code
                break
              }
            }
            router.push("add?scode=" + code)
          }
        }

        that.clickTm = now;
        that.clickIndex = params.seriesIndex
      })


      this.hychart = markRaw(echarts.init(this.$refs.hychart));
      var option = {
        tooltip: {
          trigger: 'axis',
          backgroundColor: "#e69900",
          borderWidth: 0,
          formatter: function (params) {
            return params[0].axisValue
          },
          
        },
        legend: {
          top: "0%",
          itemHeight: 3,
          itemWidth: 5,
          data: that.legendData,
          show:false,
        },
        grid: {
          left: '2%',
          right: '8%',
          bottom: '4%',
          containLabel: false
        },
        xAxis: {
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
          data: that.times
        },
        yAxis: {
          axisTick: {
            inside: true
          },
          axisLabel: {
            margin: 4,
            inside: true,
            formatter: function(value, index) {
									return value + '%';
								}
          },
          type: 'value',
          scale: true,
          splitLine: {
            show: false
          },
          left: '0%',
        },
        series: that.seriesData,
      };
      document.getElementById('hychart').removeAttribute('_echarts_instance_');
      this.hychart.setOption(option, true)


      // window.addEventListener('resize', ()=>{
      //   this.cjlchart.resize()
      //   this.cjlchart.setOption(option1)
      //   this.hychart.resize()
      //   this.hychart.setOption(option)
      // });

      // window.addEventListener('resize', this.hychart.resize);
      // // 为窗口的resize事件绑定处理函数
      // window.onresize = function() {
      // 		// 让echarts图表实例调用resize方法，以适应新窗口大小
      //     that.hychart?.resize();
      // 		that.cjlchart?.resize();
      // };
    },
    setZybEcharts() {
      var that = this;
      this.hychart.setOption({
        legend: {
          data: that.legendData,
        },

        xAxis: {
          data: that.times
        },
        series: that.seriesData,
      });
      this.cjlchart.setOption({
        legend: {
          data: that.names,
        },
        visualMap: [{
          min: that.cjlmin,
          max: that.cjlmax,
        }],
        series: that.cjlseries
      });
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

.num-text {
  background-color: #eb5a66;
  color: #ffffff;
  font-size: 11px;
  height: 15px;
  width: 15px;
  margin-right: 5px;
  padding: 1px;
  // align-items: center;
  // justify-content: center;
  // display: flex;
}


body {
  height: 100vh;
  overflow: hidden;
  display: flex;
  flex-direction: column;
  // background-color: #252526;
}

#header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: #eeeeee;
  padding: 10px;

  margin-right: -2px;
}

#headerInfo {
  width: 50%;
  display: flex;
  flex-wrap: wrap;
}

.headerInfoItem {
  width: 12.5%;
}

// #kline {
//   padding: 1rem;
//   background-color: pink;
// }

// .klineItem {
//   padding: 0.5rem;
// }


#header .left {
  position: relative;
  z-index: 5;
  float: left;
  line-height: 40px;
}

#header .left .name {
  font-size: 20px;
  color: #474747;
  font-weight: bold;
}

#header .left .code {
  font-size: 16px;
  font-weight: bold;
  margin: 0 10px;
}

#header .left .prise {
  position: relative;
  font-size: 15px;
  margin: 0 10px;
  font-weight: bold;
}

#header .left .zde,
#header .left .zdf {
  font-size: 12px;
  margin-left: 5px;
}

#header .left>div {
  display: inline-block;
}

#header .right {
  position: absolute;
  right: 10px;
  left: 20vh;
  top: 10px;
  z-index: 0;
  text-align: right;
}

#header .right .block {
  font-size: 14px;
  text-align: left;
  // display: inline-block;
  padding-right: 10px;
  margin-top: 4px;
  margin-left: 5px;
  min-width: 80px;
}

#header .right .block .row {
  line-height: 25px;
  // color: #fff;
}

#echartsIframe {
  flex: 1;
}


#root {
  display: flex;
  // margin-top: 10px;

  width: 100%;
  height: 100%;
  min-width: 100vh;
  min-height: 100vh;
  overflow-x: auto;
}

#root .left .switchbar0 {
  user-select: none;
  position: relative;
  background-color: #ea5404;
  z-index: 99;
  top: 8px;
  margin-right: 7px;
  left: 8px;
  min-height: 40px;
}

#root .left .switchbar1 {
  user-select: none;
  position: relative;
  z-index: 99;
  top: 8px;
  margin-right: 7px;
  left: 8px;
  height: 65%;
  width: 100%
}


#root .left .switchbar2 {
  user-select: none;
  position: relative;
  z-index: 99;
  top: 8px;
  margin-right: 7px;
  left: 8px;
  height: 30%;
  width: 100%
}


#root .left .switchbar {
  user-select: none;
  position: relative;
  background-color: #ea5404;
  z-index: 100;
  top: 3px;
  margin-left: 2px;
  margin-right: 2px;
  max-width: 1000px;
}

#root .left .chart {
  position: absolute;
  height: 100%;
  width: 100%;
  box-sizing: border-box;
  padding-top: 10px;
  left: 8px;
  padding-right: 7px;
  padding-bottom: 40px;
}

#root .left {
  position: absolute;
  top: 66px;
  left: 0;
  bottom: 0;
  right: 410px;
}

#root .righ {
  position: absolute;
  width: 400px;
  top: 74px;
  right: 6px;
  bottom: 10px;
}

#root .righ .infoIframe {
  position: absolute;
  height: 100%;
  width: 100%;
  top: 0;
  box-sizing: border-box;
}
</style>
