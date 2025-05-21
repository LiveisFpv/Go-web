import { ref } from 'vue'
import { defineStore } from 'pinia'

export const useAuthStore = defineStore('auth', () => {
  const isAuthenticated = ref(false)
  const token = ref<string | null>(null)

  function login(authToken: string) {
    token.value = authToken
    isAuthenticated.value = true
    localStorage.setItem('token', authToken)
  }

  function logout() {
    token.value = null
    isAuthenticated.value = false
    localStorage.removeItem('token')
  }

  function initialize() {
    const savedToken = localStorage.getItem('token')
    if (savedToken) {
      token.value = savedToken
      isAuthenticated.value = true
    }
  }

  return { isAuthenticated, token, login, logout, initialize }
}) 