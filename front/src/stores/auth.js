import { defineStore } from 'pinia'
import axios from 'axios'

const API_URL = 'http://localhost:3000'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null,
    token: localStorage.getItem('token') || null,
    isAuthenticated: !!localStorage.getItem('token')
  }),

  getters: {
    isAdmin: (state) => state.user?.role === 'admin',
    userProfile: (state) => state.user
  },

  actions: {
    async login(credentials) {
      try {
        const response = await axios.post(`${API_URL}/auth/login`, credentials)
        this.setAuth(response.data)
        return response.data
      } catch (error) {
        this.clearAuth()
        throw error
      }
    },

    async register(userData) {
      try {
        const response = await axios.post(`${API_URL}/auth/register`, userData)
        this.setAuth(response.data)
        return response.data
      } catch (error) {
        this.clearAuth()
        throw error
      }
    },

    async fetchProfile() {
      try {
        const response = await axios.get(`${API_URL}/users/profile`, {
          headers: this.getAuthHeaders()
        })
        this.user = response.data
        this.isAuthenticated = true
        return response.data
      } catch (error) {
        this.clearAuth()
        throw error
      }
    },

    setAuth(data) {
      this.token = data.token
      this.user = data.user
      this.isAuthenticated = true
      localStorage.setItem('token', data.token)
      axios.defaults.headers.common['Authorization'] = `Bearer ${data.token}`
    },

    clearAuth() {
      this.token = null
      this.user = null
      this.isAuthenticated = false
      localStorage.removeItem('token')
      delete axios.defaults.headers.common['Authorization']
    },

    getAuthHeaders() {
      return {
        Authorization: `Bearer ${this.token}`
      }
    },

    logout() {
      this.clearAuth()
    }
  }
})
