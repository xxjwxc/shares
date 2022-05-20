<template>
  <view>
    <view class="mainIndex">
      <home v-if="PageCur == 'home'"></home>
      <!-- 首页 -->
      <add v-if="PageCur == 'add'"></add>
      <!-- 设备蓝牙 -->
      <mine v-if="PageCur == 'mine'"></mine>
      <!-- 我的 -->
    </view>

    <view
      class="tabbar cu-bar bg-black shadow foot"
      :style="{ zIndex: isTop ? '999999' : '99' }"
      v-if="tabbar"
    >
      <view
        class="action"
        :class="PageCur == 'home' ? 'text-white' : 'text-gray'"
        @click="NavChange"
        data-cur="home"
      >
        <view class="cuIcon-homefill"></view>
        首页
      </view>
      <view class="action text-gray add-action">
        <button
          class="cu-btn cuIcon-search bg-black shadow"
          @click="NavChange"
          data-cur="add"
        ></button>
        搜索
      </view>
      <view
        class="action"
        :class="PageCur == 'mine' ? 'text-white' : 'text-gray'"
        @click="NavChange"
        data-cur="mine"
      >
        <view class="cuIcon-my">
          <view class="cu-tag badge" v-if="badge"></view>
        </view>
        我的
      </view>
    </view>
  </view>
</template>

<script>
  export default {
    // 在App.vue中的启动方法中，去开启授权
    //  onLaunch: async function() {
    //     console.log("---------start---------")
    //     await this.$User.WxLogin(true);// 登录
    // },
    data() {
      return {
        PageCur: "home", //切换菜单
        toPageCur: "", //上次切换的菜单
        isTop: false, //底部菜单是否在最顶层
        badge: false,
        tabbar: true,
        windowHeight: "",
      };
    },
    async onLoad(options) {
      // console.log('App111 开启')
      // let relog = await this.$User.ReLogin();
      // if (relog && relog.relogin){// 需要重新登录
      // 	await this.$User.WxLogin(false);// 登录
      // }
      uni.getSystemInfo({
        success: (res) => {
          this.windowHeight = res.windowHeight;
        },
      });
      uni.onWindowResize((res) => {
        this.tabbar = false;
        if (res.size.windowHeight < this.windowHeight) {
          this.tabbar = false;
        } else {
          this.tabbar = true;
        }
        console.log("tabbar:", this.tabbar);
      });
      let out = await this.$Shares.HaveNewMsg();
      if (out) {
        this.badge = out.badge;
      }
      // await this.$User.WxLogin();// 登录
    },
    onShow: function () {
      console.log("App1 开启");
    },
    onHide: function () {
      console.log("App1 关闭");
    },

    methods: {
      NavChange: function (e) {
        console.log(e);
        //底部菜单切换
        this.PageCur = e.currentTarget.dataset.cur;
      },
    },
    watch: {
      //监听菜单变化
      PageCur: function (newVal) {
        var _this = this;
        if (newVal == "add") {
          //如果切换的蓝牙 就把底部菜单设为最顶层 避免蓝牙处弹出提示不方便切换菜单
          _this.isTop = true;
        } else {
          _this.isTop = false;
        }

        //如果上一次切换的菜单是蓝牙，但是本次切换不是蓝牙，就关闭蓝牙搜索
        if (newVal != "add" && _this.toPageCur == "add") {
        }
        _this.toPageCur = newVal; //赋值上一次切换的菜单
      },
    },
  };
</script>
