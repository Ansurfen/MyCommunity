import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'home',
    // meta: {
    //   title: "首页",
    // },
    component: HomeView
  },
  {
    path: '/about',
    name: 'about',
    component: () => import('../views/AboutView.vue')
  },
  {
    path: '/404',
    name: 'NoPage',
    component: () => import('../views/NoPage.vue')
  },
  {
    path: '/:pathMatch(.*)',
    redirect: '/404',
  },
  // {
  //   path: '/user',
  //   name: 'user',
  //   component: HomeView,
  //   // redirect: '/404',
  //   children: [
  //     {
  //       path: '/user/login',
  //       name: '/login',
  //       component: () => import('../views/user/Login.vue')
  //     }
  //   ]
  // },
  {
    path: '/user/login',
    name: 'user/login',
    component: () => import('../views/user/Login.vue')
  },
  {
    path: '/user/register',
    name: 'user/register',
    component: () => import('../views/user/Register.vue')
  },
  {
    path: '/user/find',
    name: 'user/find',
    component: () => import('../views/user/Find.vue')
  },
  {
    path: '/community/search/:str',
    name: 'community/search',
    component: () => import('../views/community/Search.vue')
  },
  {
    path: '/community/main/:name',
    name: 'community/main',
    component: () => import('../views/community/Main.vue')
  },
  {
    path: '/community/detail/:id',
    name: 'community/detail',
    component: () => import('../views/community/Detail.vue')
  },
  {
    path: '/community/edit',
    name: 'community/edit',
    component: () => import('../views/community/Edit.vue')
  },
  {
    path: '/backed',
    name: 'backed',
    component: () => import('../views/backed/Common.vue')
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
