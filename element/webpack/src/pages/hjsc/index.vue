<template>
  <div v-loading="loading" ref="hjscRef" element-loading-background="#fff">
    <div id="header">
      <div class="left">
        <div class="name">{{ data.sharesInfo.name }}</div>
        <div class="code"> {{ data.sharesInfo.sampleCode }} </div>
        <div class="prise" style="font-size: 20px;" :style="{ color: data.sharesInfo.color0 }">{{ data.sharesInfo.price
          }}</div>
        <div class="zde" style="background-color: rgba(255, 0, 0, 0);" :style="{ color: data.sharesInfo.color0 }">{{
          data.sharesInfo.priceadd }}</div>
        <div class="zdf" style="background-color: rgba(255, 0, 0, 0);" :style="{ color: data.sharesInfo.color0 }">{{
          data.sharesInfo.percent }}%</div>
      </div>
      <div class="right">
        <div class="block" style="font-size: 14px; margin-left: 20px; padding-right: 5px;">
          <div class="row"><label>今开：</label><span class="jk rise" style="background-color: rgba(255, 0, 0, 0);"
              :style="{ color: data.sharesInfo.jkc }">{{ data.sharesInfo.jk }}</span></div>
          <div class="row"><label>昨收：</label><span class="zs ping" style="background-color: rgba(255, 0, 0, 0);">{{
            data.sharesInfo.zs }}</span></div>
        </div>
        <div class="block" style="font-size: 14px; margin-left: 20px; padding-right: 5px;">
          <div class="row"><label>最高：</label><span class="zg rise" style="background-color: rgba(255, 0, 0, 0);"
              :style="{ color: data.sharesInfo.zgc }"> {{ data.sharesInfo.zg }}</span></div>
          <div class="row"><label>最低：</label><span class="zd fall" style="background-color: rgba(255, 0, 0, 0);"
              :style="{ color: data.sharesInfo.zdc }">{{ data.sharesInfo.zd }}</span></div>
        </div>
        <div class="block" style="font-size: 14px; margin-left: 20px; padding-right: 5px;">
          <div class="row"><label>涨跌幅：</label><span class="zdf fall" style="background-color: rgba(255, 0, 0, 0);"
              :style="{ color: data.sharesInfo.color0 }">{{ data.sharesInfo.percent }}</span></div>
          <div class="row"><label>涨跌额：</label><span class="zde fall" style="background-color: rgba(255, 0, 0, 0);"
              :style="{ color: data.sharesInfo.color0 }">{{ data.sharesInfo.priceadd }}</span></div>
        </div>
        <div class="block" style="font-size: 14px; margin-left: 20px; padding-right: 5px;">
          <div class="row"><label>换手：</label><span class="hs" style="background-color: rgba(255, 0, 0, 0);">{{
            data.sharesInfo.hs }}</span></div>
          <div class="row"><label>振幅：</label><span class="zf" style="background-color: rgba(255, 0, 0, 0);">{{
            data.sharesInfo.zf }}%</span></div>
        </div>
        <div class="block" style="font-size: 14px; margin-left: 20px; padding-right: 5px;">
          <div class="row"><label>成交量：</label><span class="cjl" style="background-color: rgba(255, 0, 0, 0);">{{
            data.sharesInfo.cjl }}</span></div>
          <div class="row"><label>成交额：</label><span class="cje" style="background-color: rgba(255, 0, 0, 0);">{{
            data.sharesInfo.cje }}</span></div>
        </div>
        <div class="block" style="font-size: 14px; margin-left: 20px; padding-right: 5px;">
          <div class="row"><label>PE：</label><span class="np" style="background-color: rgba(255, 0, 0, 0);">{{
            data.sharesInfo.pe }}</span></div>
          <div class="row"><label>PETTM：</label><span class="wp" style="background-color: rgba(255, 0, 0, 0);">{{
            data.sharesInfo.pettm }}</span></div>
        </div>
        <div class="block" style="font-size: 14px; margin-left: 20px; padding-right: 5px;">
          <div class="row"><label>PB：</label><span class="szjs">{{ data.sharesInfo.pb }}<span></span></span></div>
          <div class="row"><label>股息率：</label><span class="xdjs">{{ data.sharesInfo.gxl }}%<span></span></span></div>
        </div>
        <div class="block" style="font-size: 14px; margin-left: 20px; padding-right: 5px;">
          <div class="row"><label>总市值：</label><span class="szjs">{{ data.sharesInfo.zsz }}<span></span></span></div>
          <div class="row"><label>流通市值：</label><span class="xdjs">{{ data.sharesInfo.ltsz }}<span></span></span></div>
        </div>
      </div>
    </div>

    <div id="root">
      <div class="left">
        <div class="switchbar0">
          <el-segmented class="switchbar" v-model="data.tag" :options="kline" @change='confirmkline' block />
        </div>
        <iframe @load="finishLoading" class="chart" :key="data.iframekey" :src="data.iframeUrl" frameborder="0"></iframe>
      </div>
      <div class="righ">
        <Tables ref="tablesRef" @ok="setUrl" />
      </div>
    </div>
  </div>
</template>

<script setup>
// import server from '@/utils/server/def.js'
import { useRoute } from 'vue-router';
import User from '@/utils/server/oauth';
import { ElMessageBox } from 'element-plus';
import Shares from '@/utils/server/shares'
import { kline } from './index.js'
import { reactive, ref, onUnmounted } from 'vue';
import Tables from '@/components/Tables.vue'
// import Server from '@/utils/server/def'

const data = reactive({
  searchValue: "",
  iframekey: 0,
  userName: "",
  tag: 'daily',
  iframeUrl: "",
  iframeUrl2: "",
  info: null,
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
  sharesInfo: {
    name: "",
    sampleCode: "",
    price: "",
    priceadd: "",
    percent: "",
    color0: "#ff0000",
    jk: "",// 今开
    jkc: "",
    zs: "",// 昨收
    zg: "",// 最高
    zgc: "",
    zd: "",// 最低
    zdc: "",
    zde: "",// 涨跌额
    hs: "",// 换手
    zf: "",// 振幅
    cjl: "",// 成交量
    cje: "",// 成交额
    pe: "",// 市盈率
    pettm: "",// 市盈率TTM
    pb: "",// 市净率
    zsz: "",// 总市值
    ltsz: "",// 流通市值
    gxl: "",// 股息率
  },
})
const loading = ref(false)
const time = ref()
const hjscRef = ref()
const { query } = useRoute();
const tablesRef = ref()
const loadData = async (options) => {
  loading.value = true
  data.iframekey = 1;
  let tag = options?.tag
  if (tag != null && tag.length > 0) { // "日K","分时", "5分钟", "15分钟", "30分钟", "60分钟"
    data.tag = tag;
    data.klineName.forEach((i, x) => {
      if (i.key === tag) data.klinedefaultIndex = x
    })
  }
  data.userName = localStorage.getItem("user-token")
  let info =await JSON.parse(localStorage.getItem('userInfo'))
  if (info != null) {
    data.userInfo = info
    if (info.isVip) {
      tablesRef.value?.init(async code => {
        if (code) {
          setUrl(code)
        }
      })
    } else {
      hjscRef.value.style.filter = "blur(10px)";
      ElMessageBox.confirm(
        '仅对L2会员及以上用户开放',
        '错误',
        {
          confirmButtonText: '确定',
          type: 'error',
          showCancelButton: false,
          closeOnClickModal: false,
          closeOnPressEscape: false,
          showClose: false
        }
      )
        .then(() => {
          location.href = 'https://www.xxxxxx.cn/webshares/#/pages/vip/level?platform=vscode'
        })
    }
  }
}
const finishLoading = () =>{
  setTimeout(()=>loading.value = false,500)
}


const setUrl = async (code) => {
  data.searchValue = code;
  let info = await Shares.Search(data.searchValue, data.tag); // 获取所有的中文代码
  if (info) {
    data.info = info.info;
    data.iframeUrl = "https://www.xxxxxx.cn" + data.info.img.replace("/echarts/echarts.html", "/echarts1/echarts.html");
  }
  clearTimeout(time.value)
  searchCode()
}

const confirmkline = (val) => {
  if (val === "flkx") {
    if (!data.userInfo.vipLevel <= 2 && data.userInfo.token < 1) {
      ElMessageBox.alert('Token不足.请先扫码登录', '错误', {
        confirmButtonText: 'OK',
      })
      return
    }
  }
  data.iframekey = data.iframekey + 1;
  data.iframeUrl = changeURLArg(data.iframeUrl, "tag", val)
}

const changeURLArg = (url, arg, arg_val) => {
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
}

const searchCode = async () => {
  let out = await User.GetEx("https://qt.gtimg.cn/q=" + data.searchValue + "&fmt=json")
  if (out != null) {
    let info = out.data[data.searchValue]
    if (info == null) {
      ElMessageBox.alert('没有查询到相应信息。', '错误', {
        confirmButtonText: 'OK',
      })
      return
    }

    data.sharesInfo.name = info[1];
    data.sharesInfo.sampleCode = info[2];
    data.sharesInfo.price = info[3];
    let percent = parseFloat(info[32]);
    data.sharesInfo.color0 = GetColor(data.userInfo.rg, percent)
    if (percent > 0) {
      data.sharesInfo.priceadd = "+" + info[31];
      data.sharesInfo.percent = "+" + info[32];
    } else {
      data.sharesInfo.priceadd = info[31];
      data.sharesInfo.percent = info[32];
    }
    data.sharesInfo.jk = info[5];
    data.sharesInfo.zs = info[4];
    data.sharesInfo.zg = info[33];
    data.sharesInfo.zdf = info[32];
    data.sharesInfo.hs = info[38];
    data.sharesInfo.zf = info[43];
    data.sharesInfo.pe = info[53];
    data.sharesInfo.pettm = info[39];
    data.sharesInfo.pb = info[46];
    data.sharesInfo.gxl = info[64];
    data.sharesInfo.zsz = getSamplePrice(info[45] * 100000000);
    data.sharesInfo.ltsz = getSamplePrice(info[44] * 100000000);
    data.sharesInfo.cje = getSamplePrice(info[37] * 10000);
    data.sharesInfo.cjl = getSamplePrice(parseFloat(info[36]) * 0.01)

    data.sharesInfo.zd = info[42];
    data.sharesInfo.zdc = GetColor(data.userInfo.rg, info[42] - data.sharesInfo.zs)

    data.sharesInfo.jkc = GetColor(data.userInfo.rg, info[5] - info[4])
    data.sharesInfo.zgc = GetColor(data.userInfo.rg, info[33] - data.sharesInfo.zs)
    time.value = setTimeout(() => {
      searchCode()
    }, 1500)
  }
}

const GetColor = (rg, percent) => {
  if (rg && percent >= 0) { // 默认红色
    return "#ff0000"
  }

  if (!rg && percent <= 0) { // 默认红色
    return "#ff0000"
  }

  return "#1AAD19" // 绿色
}
const getSamplePrice = (price) => {
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
}

loadData(query);

onUnmounted(() => {
  if (data.interval) clearInterval(data.interval);
})
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
  font-size: 20px;
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
  top: 10px;
  z-index: 10;
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

#root .left .switchbar {
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
