import { apiLogin, apiRegister, apiLogout, apiPing } from '@/api/auth'
import type { LoginData, RegisterData } from '@/api/types'
import router from '@/router'
import { useUserStore } from '@/stores/user'
import { ElMessage } from 'element-plus'
import { useI18n } from 'vue-i18n'

export const useAuth = () => {
  const userStore = useUserStore()
  const { t } = useI18n()

  const register = async (data: RegisterData) => {
    try {
      const response = await apiRegister(data)

      if (response.status === 200) {
        ElMessage.success(t('RegisterSuccess'))
        router.push('/login')
      } else {
        ElMessage.error(t('RegisterFailed'))
      }
    } catch (error) {
      ElMessage.error(t('RegisterErrorAuth'))
    }
  }

  const login = async (data: LoginData) => {
    try {
      const response = await apiLogin(data)

      if (response.data.token) {
        userStore.setToken(response.data.token, data.username)
        ElMessage.success(t('LoginSuccess'))
        router.push('/chat')
      } else {
        ElMessage.error(t('LoginFailed'))
      }
    } catch (error) {
      ElMessage.error(t('LoginErrorAuth'))
    }
  }

  const logout = async () => {
    apiLogout()
    userStore.clearToken()
    ElMessage.success(t('LogoutSuccess'))
    router.push('/')
  }

  const ping = async () => {
    try {
      const res = await apiPing()
      if (res.data) {
        ElMessage.success(res.data)
      }
    } catch (error) {
      console.error(error)
      ElMessage.error(t('Error'))
    }
  }

  return { register, login, logout, ping }
}
