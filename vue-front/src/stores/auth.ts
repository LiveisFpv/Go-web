import { ref } from 'vue'
import { defineStore } from 'pinia'
import { userService } from '../services/userService'

export const useAuthStore = defineStore('auth', () => {
  const isAuthenticated = ref(false)
  const token = ref<string | null>(null)
  const user_role = ref<string | null>(null)
  const email = ref<string | null>(null)

  function decodeJwt(token: string) {
    try {
      const base64Url = token.split('.')[1]
      const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/')
      const jsonPayload = decodeURIComponent(atob(base64).split('').map(function(c) {
        return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2)
      }).join(''))
      return JSON.parse(jsonPayload)
    } catch (e) {
      console.error('Error decoding JWT:', e)
      return null
    }
  }

  async function login(authToken: string) {
    token.value = authToken
    isAuthenticated.value = true
    const decodedToken = decodeJwt(authToken)
    email.value = decodedToken?.email || null
    localStorage.setItem('token', authToken)
    
    if (email.value) {
      console.log('Decoded email:', email.value)
      localStorage.setItem('email', email.value)
      try {
        const userData = await userService.getUserbyEmail(email.value)
        user_role.value = userData.user_role
        localStorage.setItem('user_role', userData.user_role || 'USER')
      } catch (error) {
        console.error('Error fetching user role:', error)
      }
    }
  }

  function logout() {
    token.value = null
    isAuthenticated.value = false
    email.value = null
    user_role.value = null
    localStorage.removeItem('token')
    localStorage.removeItem('user_role')
    localStorage.removeItem('email')
  }

  async function initialize() {
    const savedToken = localStorage.getItem('token')
    const savedUserRole = localStorage.getItem('user_role')
    const savedEmail = localStorage.getItem('email')

    if (savedToken) {
      token.value = savedToken
      isAuthenticated.value = true
      const decodedToken = decodeJwt(savedToken)
      email.value = decodedToken?.email || savedEmail || null
      
      if (email.value && !savedUserRole) {
        try {
          const userData = await userService.getUserbyEmail(email.value)
          user_role.value = userData.user_role
          localStorage.setItem('user_role', userData.user_role || 'USER')
        } catch (error) {
          console.error('Error fetching user role:', error)
        }
      } else {
        user_role.value = savedUserRole
      }
    }
  }

  return { isAuthenticated, token, email, user_role, login, logout, initialize }
})
