const VueRouter = require('vue-router')

const routes = [
  { path: '/login', component: require('./views/Login.vue').default },
  // TODO: setting page
]

const router = VueRouter.createRouter({
  history: VueRouter.createWebHashHistory(),
  routes
})

export default router
