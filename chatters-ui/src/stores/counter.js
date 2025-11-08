import { defineStore } from 'pinia'
import { ref } from 'vue'
import api from '@/lib/api'
import router from '@/router/index'

export const useAuthStore = defineStore('auth', () => {
  const user = ref(JSON.parse(localStorage.getItem('user') || 'null'))
  const token = ref(localStorage.getItem('token') || null)
  const loading = ref(false)
  const error = ref(null)

  function setSession(u, t) {
    user.value = u
    token.value = t
    if (u) localStorage.setItem('user', JSON.stringify(u)); else localStorage.removeItem('user')
    if (t) localStorage.setItem('token', t); else localStorage.removeItem('token')
  }

  async function login(username, password) {
    loading.value = true
    error.value = null
    try {
      const res = await api.post('/login', { username, password })
      const u = res.data.user
      const t = res.data.token
      setSession(u, t)
      await router.push({ name: 'home' })
      return u
    } catch (err) {
      error.value = err.response?.data?.error || err.message || 'Login failed'
      throw err
    } finally {
      loading.value = false
    }
  }

  // For cookies we clear them server side
  function logout() {
    setSession(null, null)
    router.push({ name: 'login' })
  }

  return { user, token, loading, error, login, logout, setSession }
})
