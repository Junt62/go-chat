import { defineStore } from 'pinia'

const EXPIRATION_TIME = 20 * 60 * 60 * 1000

export const useUserStore = defineStore('user', {
  state: () => ({
    token: localStorage.getItem('token') || '',
    username: '',
    tokenSetTime: localStorage.getItem('tokenSetTime')
      ? new Date(localStorage.getItem('tokenSetTime') as string)
      : null,
    onlineUsers: [
      { id: 1, name: 'Alice' },
      { id: 2, name: 'Bob' },
      { id: 3, name: 'Charlie' },
    ],
    friends: [
      { id: 2, name: 'Bob' },
      { id: 4, name: 'David' },
    ],
  }),
  actions: {
    setToken(token: string, username: string) {
      const currentTime = new Date()
      this.tokenSetTime = currentTime
      this.token = token
      this.username = username
      localStorage.setItem('token', token)
      localStorage.setItem('tokenSetTime', currentTime.toISOString())
    },
    clearToken() {
      this.token = ''
      this.username = ''
      this.tokenSetTime = null
      localStorage.removeItem('token')
      localStorage.removeItem('tokenSetTime')
    },
    isTokenExpired() {
      if (!this.tokenSetTime) {
        return true
      }
      const currentTime = new Date()
      const timestamp1 = this.tokenSetTime.getTime()
      const timestamp2 = currentTime.getTime()
      const timeDifference = timestamp2 - timestamp1
      return timeDifference > EXPIRATION_TIME
    },
  },
})
