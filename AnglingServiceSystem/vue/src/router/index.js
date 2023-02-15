import Vue from 'vue'
import VueRouter from 'vue-router'
import login from '../views/Login.vue'
import home from '../views/Home.vue'
import info from '../views/Info.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'login',
    component: login
  },
  {
    path:'/home',
    name:'home',
    component:home,
  },
  {
    path:'/info',
    name:'info',
    component:info,
  }
]

const router = new VueRouter({
  routes
})

export default router
