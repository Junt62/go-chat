import { apiLogin, apiRegister } from '@/api/auth'
import type { LoginData, RegisterData } from '@/api/types'
import router from '@/router'
import { useUserStore } from '@/stores/user'
import { ElMessage } from 'element-plus'

export const useAuth = () => {
  const userStore = useUserStore()

  const register = async (data: RegisterData) => {
    try {
      const response = await apiRegister(data)

      if (response.status === 200) {
        ElMessage.success('Sign up successful')
        router.push('/login')
      } else {
        ElMessage.error('Sign up failed, please try again')
      }
    } catch (error) {
      ElMessage.error('Please check the form for errors')
    }
  }

  const login = async (data: LoginData) => {
    try {
      const response = await apiLogin(data)

      if (response.data.token) {
        userStore.setToken(response.data.token, data.username)
        ElMessage.success('Login successful')
        router.push('/chat')
      } else {
        ElMessage.error('Incorrect username or password. Please try again.')
      }
    } catch (error) {
      ElMessage.error('Please check the form for errors')
    }
  }
  const logout = async () => {
    userStore.clearToken()
    ElMessage.success('Logged out successfully')
    router.push('/login')
  }

  return { register, login, logout }
}
