import { defineStore } from 'pinia'
import { login as loginApi, getUserInfo as getUserInfoApi } from '@/api/user'

export const useUserStore = defineStore('user', {
  state: () => ({
    token: localStorage.getItem('radish-token') || '',
    userInfo: null
  }),
  getters: {
    isLogin: (state) => Boolean(state.token)
  },
  actions: {
    async login(username, password) {
      const data = await loginApi({ username, password })
      this.token = data.token
      this.userInfo = data.user
      localStorage.setItem('radish-token', data.token)
      return data
    },
    async fetchUserInfo() {
      this.userInfo = await getUserInfoApi()
      return this.userInfo
    },
    logout() {
      this.token = ''
      this.userInfo = null
      localStorage.removeItem('radish-token')
    }
  }
})
