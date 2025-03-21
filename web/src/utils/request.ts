import axios from 'axios'
import { useUserStore } from '@/stores/user'
import router from '@/router'
import { ElMessage } from 'element-plus'

const request = axios.create({
  baseURL: import.meta.env.VITE_SERVER_ADDRESS || '',
  timeout: 10000,
})

request.interceptors.request.use(
  (config) => {
    const userStore = useUserStore()
    if (userStore.token) {
      config.headers.Authorization = `Bearer ${userStore.token}`
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  },
)

request.interceptors.response.use(
  (response) => {
    return response
  },
  (error) => {
    if (error.response?.status === 401) {
      const userStore = useUserStore()
      if (userStore.token) {
        userStore.clearToken()
        router.push('/login')
        ElMessage.error('Session expired, please login again')
      }
    }
    return Promise.reject(error)
  },
)

export default request
