import type { AuthResponse } from '../types/auth'

export const mockUser = {
  id: 1,
  username: 'testuser',
  email: 'test@example.com'
}

export const mockToken = 'mock-jwt-token'

export const mockAuthResponse: AuthResponse = {
  token: mockToken,
  user: mockUser
}

export const mockUsers = [
  {
    email: 'test@example.com',
    password: 'Test123!',
    username: 'testuser'
  },
  {
    email: 'admin@example.com',
    password: 'Admin123!',
    username: 'admin'
  }
] 