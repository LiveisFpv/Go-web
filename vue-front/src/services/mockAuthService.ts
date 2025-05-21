import type { LoginRequest, RegisterRequest, GoogleAuthRequest, AuthResponse, AuthError } from '../types/auth'
import { mockAuthResponse, mockUsers } from '../mocks/authMocks'
import { getMockDelay, isMockEnabled } from '../config/mockConfig'

// Имитация задержки сети
const delay = (ms: number) => new Promise(resolve => setTimeout(resolve, ms))

export const mockLogin = async (data: LoginRequest): Promise<AuthResponse> => {
  if (!isMockEnabled('auth')) {
    throw new Error('Auth mocks are disabled')
  }

  await delay(getMockDelay('auth'))

  const user = mockUsers.find(u => u.email === data.email && u.password === data.password)
  
  if (!user) {
    throw {
      message: 'Неверный email или пароль',
      field: 'email'
    } as AuthError
  }

  return mockAuthResponse
}

export const mockRegister = async (data: RegisterRequest): Promise<AuthResponse> => {
  if (!isMockEnabled('auth')) {
    throw new Error('Auth mocks are disabled')
  }

  await delay(getMockDelay('auth'))

  if (mockUsers.some(u => u.email === data.email)) {
    throw {
      message: 'Пользователь с таким email уже существует',
      field: 'email'
    } as AuthError
  }

  return mockAuthResponse
}

export const mockGoogleAuth = async (data: GoogleAuthRequest): Promise<AuthResponse> => {
  if (!isMockEnabled('google')) {
    throw new Error('Google auth mocks are disabled')
  }

  await delay(getMockDelay('google'))
  return mockAuthResponse
}

export const mockLogout = async (): Promise<void> => {
  if (!isMockEnabled('auth')) {
    throw new Error('Auth mocks are disabled')
  }

  await delay(getMockDelay('auth'))
} 