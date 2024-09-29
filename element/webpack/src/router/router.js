
const routes = [
    { path: '/',component: () => import('@/pages/home/home.vue') },
    { path: '/add',  component: () => import('@/pages/add/index.vue') },
    { path: '/login',  component: () => import('@/pages/login/index.vue') },
    { path: '/zyb',  component: () => import('@/pages/zyb/index.vue') },
    { path: '/zsb',  component: () => import('@/pages/zyb/zsb.vue') },
    { path: '/hjsc',  component: () => import('@/pages/hjsc/index.vue') },
    { path: '/hydayliy',  component: () => import('@/pages/hy/hydayliy.vue') },
    { path: '/yyq',  component: () => import('@/pages/yyq/yyq.vue') },
    { path: '/chat',  component: () => import('@/pages/chat/index.vue') },
  ]

export default routes