export interface LoginRequest {
  email: string
  password: string
}

export interface RegisterRequest {
  email: string
  password: string
}

export interface GoogleAuthRequest {
  token: string
}

export interface AuthResponse {
  data: {
    token: string
  }
}

export interface AuthError {
  message: string
  field?: string
  value?: string
}
