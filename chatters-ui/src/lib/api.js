import axios from 'axios'
import { useAuthStore } from '@/stores/counter'

// We make a new api
const api = axios.create({
  baseURL: 'http://localhost:8080',
  timeout: 5000,
  // Cookies
  withCredentials: false
})

// We attach the token
api.interceptors.request.use(config => {
  const store = useAuthStore()
  const token = store.token
  if (token) {
    config.headers = config.headers || {}
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
}, err => Promise.reject(err))

export default api
