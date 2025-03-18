import { defineStore } from 'pinia'

export const useUserStore = defineStore('user', {
  state: () => ({
    token: localStorage.getItem('token') || '',
    username: '',
  }),
  actions: {
    async setToken(token: string, username: string) {
      this.token = token
      this.username = username
      localStorage.setItem('token', token)
    },
    async logout() {
      this.token = ''
      this.username = ''
      localStorage.removeItem('token')
    },
  },
})
