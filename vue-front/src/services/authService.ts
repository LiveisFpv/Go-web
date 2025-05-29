import axios from 'axios'
import type { LoginRequest, RegisterRequest, GoogleAuthRequest, AuthResponse, AuthError } from '../types/auth'
import { USE_MOCKS } from '../config/mockConfig'
import { useAuthStore } from '@/stores/auth'

const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/api'

const authStore = useAuthStore()

const authApi = axios.create({
  baseURL: API_URL,
  headers: {
    'Content-Type': 'application/json'
  }
})

// Добавляем перехватчик для добавления токена к запросам
authApi.interceptors.request.use((config) => {
  const token = localStorage.getItem('token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

export const login = async (data: LoginRequest): Promise<AuthResponse> => {

  try {
    const response = await authApi.post<AuthResponse>('/auth', data)
    authStore.login(response.data.data.token)
    return response.data
  } catch (error) {
    if (axios.isAxiosError(error) && error.response?.data) {
      throw error.response.data as AuthError
    }
    throw { message: 'Ошибка при авторизации' } as AuthError
  }
}

export const register = async (data: RegisterRequest): Promise<AuthResponse> => {

  try {
    const response = await authApi.post<AuthResponse>('/register', data)
    localStorage.setItem('token', response.data.data.token)
    return response.data
  } catch (error) {
    if (axios.isAxiosError(error) && error.response?.data) {
      throw error.response.data as AuthError
    }
    throw { message: 'Ошибка при регистрации' } as AuthError
  }
}

export const googleAuth = async (data: GoogleAuthRequest): Promise<AuthResponse> => {

  try {
    const response = await authApi.post<AuthResponse>('/auth/google', data)
    localStorage.setItem('token', response.data.data.token)
    return response.data
  } catch (error) {
    if (axios.isAxiosError(error) && error.response?.data) {
      throw error.response.data as AuthError
    }
    throw { message: 'Ошибка при авторизации через Google' } as AuthError
  }
}

export const logout = async (): Promise<void> => {

  try {
    await authApi.post('/auth/logout')
  } catch (error) {
    console.error('Ошибка при выходе:', error)
  } finally {
    localStorage.removeItem('token')
  }
}
