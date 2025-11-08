import { createRouter, createWebHistory } from 'vue-router'
import LoginView from '../views/LoginView.vue'
import HomeView from '../views/HomeView.vue'

const routes = [
  { path: '/login', name: 'login', component: LoginView },
  { 
    path: '/', 
    name: 'home', 
    component: HomeView,
    meta: { requiresAuth: true }
  },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
})

// Authentification
// router.beforeEach((to, from, next) => {
//   const isLoggedIn = !!localStorage.getItem('token') // demo
//   if (to.meta.requiresAuth && !isLoggedIn) {
//     next({ name: 'login' })
//   } else if (to.name === 'login' && isLoggedIn) {
//     next({ name: 'home' })
//   } else {
//     next()
//   }
// })

export default router
