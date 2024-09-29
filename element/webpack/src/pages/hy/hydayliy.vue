<template>
  <!-- <el-switch v-model="value2" class="ml-2" @change="switchchange" /> -->

  <div id="header">
    <div class="left">
      <div class="name">{{ sharesInfo.name }}</div>
      <div class="code" > {{ sharesInfo.sampleCode }} </div>
      <div class="prise" style="font-size: 20px;" :style="{color:sharesInfo.color0}">{{ sharesInfo.price }}</div>
      <!-- <div class="zde" style="background-color: rgba(255, 0, 0, 0);" :style="{color:sharesInfo.color0}">{{ sharesInfo.priceadd }}</div> -->
      <div class="zdf" style="background-color: rgba(255, 0, 0, 0);" :style="{color:sharesInfo.color0}">{{ sharesInfo.percent }}%</div>
    </div>
    <div class="right"> 
        <!-- <div class="block" style="font-size: 14px; margin-left: 20px; padding-right: 5px;">
          <div class="row"><label>今开：</label><span class="jk rise" style="background-color: rgba(255, 0, 0, 0);" :style="{color: sharesInfo.jkc}">{{ sharesInfo.jk }}</span></div>
          <div class="row"><label>昨收：</label><span class="zs ping" style="background-color: rgba(255, 0, 0, 0);">{{ sharesInfo.zs }}</span></div>
        </div>
        <div class="block" style="font-size: 14px; margin-left: 20px; padding-right: 5px;">
          <div class="row"><label>最高：</label><span class="zg rise" style="background-color: rgba(255, 0, 0, 0);" :style="{color: sharesInfo.zgc}"> {{ sharesInfo.zg }}</span></div>
          <div class="row"><label>最低：</label><span class="zd fall" style="background-color: rgba(255, 0, 0, 0);" :style="{color: sharesInfo.zdc}">{{ sharesInfo.zd }}</span></div>
        </div>
        <div class="block" style="font-size: 14px; margin-left: 20px; padding-right: 5px;">
          <div class="row"><label>涨跌幅：</label><span class="zdf fall" style="background-color: rgba(255, 0, 0, 0);" :style="{color: sharesInfo.color0}">{{ this.sharesInfo.percent }}</span></div>
          <div class="row"><label>涨跌额：</label><span class="zde fall" style="background-color: rgba(255, 0, 0, 0);" :style="{color: sharesInfo.color0}">{{ sharesInfo.priceadd }}</span></div>
        </div>
        <div class="block" style="font-size: 14px; margin-left: 20px; padding-right: 5px;">
          <div class="row"><label>换手：</label><span class="hs" style="background-color: rgba(255, 0, 0, 0);">{{ sharesInfo.hs }}</span></div>
          <div class="row"><label>振幅：</label><span class="zf" style="background-color: rgba(255, 0, 0, 0);">{{ sharesInfo.zf }}%</span></div>
        </div>
        <div class="block" style="font-size: 14px; margin-left: 20px; padding-right: 5px;">
          <div class="row"><label>成交量：</label><span class="cjl" style="background-color: rgba(255, 0, 0, 0);">{{ sharesInfo.cjl }}</span></div>
          <div class="row"><label>成交额：</label><span class="cje" style="background-color: rgba(255, 0, 0, 0);">{{ sharesInfo.cje }}</span></div>
        </div>
        <div class="block" style="font-size: 14px; margin-left: 20px; padding-right: 5px;">
          <div class="row"><label>PE：</label><span class="np" style="background-color: rgba(255, 0, 0, 0);">{{ sharesInfo.pe }}</span></div>
          <div class="row"><label>PETTM：</label><span class="wp" style="background-color: rgba(255, 0, 0, 0);">{{ sharesInfo.pettm }}</span></div>
        </div>
        <div class="block" style="font-size: 14px; margin-left: 20px; padding-right: 5px;">
          <div class="row"><label>PB：</label><span class="szjs">{{ sharesInfo.pb }}<span></span></span></div>
          <div class="row"><label>股息率：</label><span class="xdjs">{{ sharesInfo.gxl }}%<span></span></span></div>
        </div>
        <div class="block" style="font-size: 14px; margin-left: 20px; padding-right: 5px;">
          <div class="row"><label>总市值：</label><span class="szjs">{{ sharesInfo.zsz }}<span></span></span></div>
          <div class="row"><label>流通市值：</label><span class="xdjs">{{ sharesInfo.ltsz }}<span></span></span></div>
        </div> -->
    </div>
  </div>

  <div id="root">
    <div class="left">
      <div class="switchbar0">
        <el-segmented class="switchbar" v-model="tag" :options="klineName" @change='confirmkline' block />
      </div>
      <iframe class="chart" :key="iframekey" :src="iframeUrl" frameborder="0"></iframe>
    </div>
    <div class="righ">
     <iframe  class="infoIframe" :src="iframeUrl2" frameborder="0"></iframe>
    </div>
  </div>
</template>

<script>
// import server from '@/utils/server/def.js'
import { useRoute } from 'vue-router';
import User from '@/utils/server/oauth';
import { ElMessageBox } from 'element-plus';
import Shares from '@/utils/server/shares'
// import user from '@/utils/server/oauth';
// import Server from '@/utils/server/def'

export default {
  name: 'hydayliyView',
  props: {
    msg: String
  },
  data() {
    return {
      searchValue: "",
      iframekey: 0,
      userName: "",
      tag: 'hy',
      iframeUrl: "",
      iframeUrl2: "",
      info: null,
      klineName: [{
        label: "日K",
        value: "hy",
      }, {
        label: "分时",
        value: "hymin",
      }, {
        label: "5分钟",
        value: "hyk5",
      }, {
        label: "15分钟",
        value: "hyk15",
      }, {
        label: "30分钟",
        value: "hyk30",
      },
      {
        label: "60分钟",
        value: "hyk60",
      },
      {
        label: "复利线",
        value: "hyflkx",
      },
      {
        label: "周K",
        value: "hyk102",
      }, {
        label: "月K",
        value: "hyk103",
      }],
      userInfo: {
					avatarURL: "",
					token: 0,
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
					only249: false,
					sendArtic: false,
					isVip: false, // 是否是vip
					vipOutTime: "-", // 会员过期时间
					vipLevel: 0, // vip等级
				},
        sharesInfo:{
          name: "",
          sampleCode: "",
          price:"",
          priceadd:"",
          percent:"",
          color0: "#ff0000",
          // jk: "",// 今开
          // jkc: "",
          // zs: "",// 昨收
          // zg: "",// 最高
          // zgc:"",
          // zd: "",// 最低
          // zdc: "",
          // zde: "",// 涨跌额
          // hs: "",// 换手
          // zf: "",// 振幅
          // cjl: "",// 成交量
          // cje: "",// 成交额
          // pe: "",// 市盈率
          // pettm:  "",// 市盈率TTM
          // pb: "",// 市净率
          // zsz: "",// 总市值
          // ltsz: "",// 流通市值
          // gxl: "",// 股息率
        },
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
    if (this.interval) {
      clearInterval(this.interval);
    }
  },
  querySelector() { },
  methods: {
    confirmkline(val) {
      if (val === "flkx") {
					if (!this.userInfo.vipLevel <=2 && this.userInfo.token < 1) {
            ElMessageBox.alert('Token不足.请先扫码登录', '错误', {
                confirmButtonText: 'OK',
            })
						return
					}
			}

      this.iframekey = this.iframekey + 1;
			this.iframeUrl = this.changeURLArg(this.iframeUrl, "tag", val)
    },
    changeURLArg(url, arg, arg_val) {
				var pattern = arg + '=([^&]*)';
				var replaceText = arg + '=' + arg_val;
				if (url.match(pattern)) {
					var tmp = '/(' + arg + '=)([^&]*)/gi';
					tmp = url.replace(eval(tmp), replaceText);
					return tmp;
				} else {
					if (url.match('[\\?]')) {
						return url + '&' + replaceText;
					} else {
						return url + '?' + replaceText;
					}
				}
			},
    /**
     * 初始化
     */

    async loadData(options) {
      this.iframekey = 1;
      let tag = options?.tag
      if (tag != null && tag.length > 0) { // "日K","分时", "5分钟", "15分钟", "30分钟", "60分钟"
        this.tag = tag;
        this.klineName.forEach((i, x) => {
          if (i.key === tag) this.klinedefaultIndex = x
        })
      }
      this.userName = localStorage.getItem("user-token")
      this.userInfo =await JSON.parse(localStorage.getItem('userInfo'))

      let code = options?.scode
      if (code != null && code.length > 0) {
        this.searchValue = code;
        let info = await Shares.HySearch(this.searchValue, this.tag); // 获取所有的中文代码
				if (info) {
					this.info = info.info;
          this.iframeUrl = "https://www.xxxxxx.cn" + this.info.img.replace("/echarts/echarts.html", "/echarts1/echarts.html");
          this.iframeUrl2 = "https://www.xxxxxx.cn/webshares/#/pages/hy/hydayliy?platform=vscode&scode=" + this.searchValue;
          if (this.userName != null && this.userName.length > 0) {
            this.iframeUrl2 += "&username=" + this.userName
          }
          this.searchCode()
        }
       
      }
    },
    async searchCode() {


        this.sharesInfo.name = this.info.name;
        this.sharesInfo.sampleCode = this.info.simpleCode;
        this.sharesInfo.price =  this.info.price;
        let percent = this.info.percent;
        this.sharesInfo.color0 =this.info.color
        if(percent >0){
          // this.sharesInfo.priceadd = "+"+data[31];
          this.sharesInfo.percent = "+"+percent;
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
  --el-segmented-item-selected-color:  #000;
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
  z-index: 10;
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
    font-size: 20px;
    margin: 0 10px;
    font-weight: bold;
}
#header .left .zde, #header .left .zdf {
    font-size: 12px;
    margin-left: 5px;
}
#header .left > div {
    display: inline-block;
}

#header .right {
    position: absolute;
    right: 10px;
    top: 10px;
    z-index: 5;
    text-align: right;
}

#header .right .block {
    font-size: 14px;
    text-align: left;
    display: inline-block;
    padding-right: 10px;
    margin-top: 4px;
    margin-left: 40px;
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

#root  .left .switchbar0 {
    user-select: none;
    position: relative;
    background-color: #ea5404;
    z-index: 99;
    top: 8px;
    margin-right: 7px;
    left: 8px;
    min-height: 40px;
}

#root  .left .switchbar {
    user-select: none;
    position: relative;
    background-color: #ea5404;
    z-index: 100;
    top: 3px;
    margin-left: 2px;
    margin-right: 2px;
    max-width: 600px;
}

#root .left .chart {
    position: absolute;
    height: 100%;
    width: 100%;
    box-sizing: border-box;
    padding-top: 10px;
    left: 8px;
    padding-right: 7px;
    padding-bottom: 40px ;
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
    z-index: 20;
}

#root .righ  .infoIframe{
    position: absolute;
    height: 100%;
    width: 100%;
    top: 0;
    box-sizing: border-box;
}

</style>
