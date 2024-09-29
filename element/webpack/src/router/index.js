import { createRouter, createWebHashHistory } from 'vue-router'
import User from '@/utils/server/oauth';
import routes from './router'
const router = createRouter({
  // history: createWebHistory("/pc"),
  history: createWebHashHistory("/pc"),
  routes,
})


router.beforeEach(async (to, from, next) => {
  // console.log(to)
  // console.log(from)
  if (to.path === '/login') return next()

  let name = to.query?.username  || ''
  if (name != null && name.length > 0) {
    localStorage.setItem("user-token", name)
  } else {
    let logout = to.query?.logout || ''
    if (logout == 'true') {
      localStorage.clear()
    }else {
      name = localStorage.getItem("user-token")
    }
  }

  if (name?.length > 0) {
    let info = await User.GetUserInfo()
    if (info != null) {
      localStorage.setItem("userInfo", JSON.stringify(info))
    } 
  }
  
  next()
})

export default router