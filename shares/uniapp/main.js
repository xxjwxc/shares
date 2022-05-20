import Vue from 'vue'
import App from './App'

import User from './utils/server/oauth'
import Shares from './utils/server/shares'



// //顶部导航
// import cuCustom from './lib/colorui/components/cu-custom.vue'
// Vue.component('cu-custom',cuCustom)
//首页
import home from './pages/home/home.vue'
Vue.component('home',home)
//添加（添加）
import add from './pages/add/add.vue'
Vue.component('add',add) 
//我的
import mine from './pages/mine/mine.vue'
Vue.component('mine',mine)
// //分组
// import group from './pages/group/group.vue'
// Vue.component('group',group)

/**
 *  因工具函数属于公司资产, 所以直接在Vue实例挂载几个常用的函数
 *  所有测试用数据均存放于根目录json.js
 *  
 *  css部分使用了App.vue下的全局样式和iconfont图标，有需要图标库的可以留言。
 *  示例使用了uni.scss下的变量, 除变量外已尽量移除特有语法,可直接替换为其他预处理器使用
 */
const msg = (title, duration=1500, mask=false, icon='none')=>{
	//统一提示方便全局修改
	if(Boolean(title) === false){
		return;
	}
	uni.showToast({
		title,
		duration,
		mask,
		icon
	});
}
const json = type=>{
	//模拟异步请求数据
	return new Promise(resolve=>{
		setTimeout(()=>{
			resolve(Json[type]);
		}, 500)
	})
}

const prePage = ()=>{
	let pages = getCurrentPages();
	let prePage = pages[pages.length - 2];
	// #ifdef H5
	return prePage;
	// #endif
	return prePage.$vm;
}


Vue.config.productionTip = false
Vue.prototype.$fire = new Vue();
Vue.prototype.$User = User;
Vue.prototype.$Shares = Shares;

App.mpType = 'app'

const app = new Vue({
    ...App
})
app.$mount()
