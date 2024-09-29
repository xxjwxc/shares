import {createApp } from 'vue'

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
// import 'element-plus/theme-chalk/dark/css-vars.css'
import './index.css'
import App from './App.vue'
import router from './router/index'
// 预览组件以及样式
import VMdPreview from '@kangc/v-md-editor/lib/preview';
import '@kangc/v-md-editor/lib/style/preview.css';
// VuePress主题以及样式（这里也可以选择github主题）
import vuepressTheme from '@kangc/v-md-editor/lib/theme/vuepress.js';
import '@kangc/v-md-editor/lib/theme/style/vuepress.css';

// Prism
import Prism from 'prismjs';
// 代码高亮
import 'prismjs/components/prism-json';
// 选择使用主题
VMdPreview.use(vuepressTheme, {
  Prism,
});

// Vue.config.productionTip = false
// Vue.prototype.$fire = new Vue();
// Vue.prototype.$User = User;
// Vue.prototype.$Shares = Shares;
// Vue.prototype.$WxShares = WxShares;
// Vue.prototype.$Order = Order;

const app = createApp(App)
app.use(ElementPlus)
app.use(router)
// 引入v-md-editor预览组件
app.use(VMdPreview);
app.directive('scroll-to-bottom', {
    mounted(el) {
      el.scrollTo({
        top: el.scrollHeight - el.clientHeight,
        behavior: 'smooth' // 这里指定了平滑滚动的行为
      });
    },
    updated(el) {
      el.scrollTo({
        top: el.scrollHeight - el.clientHeight,
        behavior: 'smooth' // 这里指定了平滑滚动的行为
      });
    },
  });
app.mount('#app')
