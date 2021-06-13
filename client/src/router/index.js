import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Buyers from '../views/Buyers.vue'
import BuyersInfo from '../views/BuyersInfo.vue'
Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
   {
    path: '/buyers',
    name: 'Buyers',
    component: Buyers
  },
  {
    path: '/buyers/:id',
    name: 'BuyersInfo',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: BuyersInfo
    // component: () => import(/* webpackChunkName: "about" */ '../views/BuyerInfo.vue')
  }
]

const router = new VueRouter({
  routes
})

export default router
