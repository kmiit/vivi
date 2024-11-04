import { createRouter, createWebHistory } from 'vue-router'
import FilePage from '@/views/FilePage.vue'
import HomePage from '@/views/HomePage.vue'
import SelfPage from '@/views/SelfPage.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomePage
    },
    {
      path: '/files',
      name: 'filepage',
      component: FilePage
    },
    {
      path: '/self',
      name: 'selfpage',
      component: SelfPage
    }
  ]
})

export default router
