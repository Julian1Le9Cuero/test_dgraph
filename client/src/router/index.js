import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import BuyersPage from '../views/BuyersPage.vue'
import BuyerInfo from '../views/BuyerInfo.vue'
Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
   {
    path: '/buyers',
    name: 'BuyersPage',
    component: BuyersPage
  },
  {
    path: '/buyers/:id',
    name: 'BuyerInfo',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: BuyerInfo
    // component: () => import(/* webpackChunkName: "about" */ '../views/BuyerInfo.vue')
  }
]

const router = new VueRouter({
  routes
})

export default router
