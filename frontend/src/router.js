const VueRouter = require('vue-router')

const routes = [
  { path: '/login', component: require('./views/Login.vue').default },
  { path: '/home', component: require('./views/Home.vue').default },
  { path: '/setting', component: require('./views/Setting.vue').default }
]

const router = VueRouter.createRouter({
  history: VueRouter.createWebHashHistory(),
  routes
})

export default router
