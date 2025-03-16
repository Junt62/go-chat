import { defineStore } from 'pinia'
import { login as apiLogin, logout as apiLogout, ping as apiPing } from '@/api/auth'

export const useUserStore = defineStore('user', {
  state: () => ({
    token: localStorage.getItem('token') || '',
    username: '',
  }),
  actions: {
    async register(username: string, password: string, email: string) {
      try {
        const res = await apiLogin({ username, password })
        this.token = res.data.token
        this.username = username
        localStorage.setItem('token', res.data.token)
      } catch (error) {
        throw error
      }
    },
    async login(username: string, password: string) {
      try {
        const res = await apiLogin({ username, password })
        this.token = res.data.token
        this.username = username
        localStorage.setItem('token', res.data.token)
      } catch (error) {
        throw error
      }
    },
    async logout() {
      await apiLogout()
      this.token = ''
      this.username = ''
      localStorage.removeItem('token')
    },
    async ping() {
      try {
        const res = await apiPing()
        console.log(res)
      } catch (error) {
        throw error
      }
    },
  },
})
