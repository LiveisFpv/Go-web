export interface ValidationError {
  field: string
  message: string
}

export const validateEmail = (email: string): string | null => {
  const emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/
  if (!email) return 'Email обязателен'
  if (!emailRegex.test(email)) return 'Некорректный формат email'
  return null
}

export const validateUsername = (username: string): string | null => {
  if (!username) return 'Имя пользователя обязательно'
  if (username.length < 3) return 'Имя пользователя должно содержать минимум 3 символа'
  if (username.length > 30) return 'Имя пользователя не должно превышать 30 символов'
  if (!/^[a-zA-Z0-9_]+$/.test(username)) {
    return 'Имя пользователя может содержать только буквы, цифры и символ подчеркивания'
  }
  return null
}

export const validatePassword = (password: string): string | null => {
  if (!password) return 'Пароль обязателен'
  if (password.length < 8) return 'Пароль должен содержать минимум 8 символов'
  if (!/[A-Z]/.test(password)) return 'Пароль должен содержать хотя бы одну заглавную букву'
  if (!/[a-z]/.test(password)) return 'Пароль должен содержать хотя бы одну строчную букву'
  if (!/[0-9]/.test(password)) return 'Пароль должен содержать хотя бы одну цифру'
  if (!/[!@#$%^&*.,]/.test(password)) return 'Пароль должен содержать хотя бы один специальный символ (!@#$%^&*.,)'
  return null
}

export const validatePasswordConfirmation = (password: string, confirmation: string): string | null => {
  if (!confirmation) return 'Подтверждение пароля обязательно'
  if (password !== confirmation) return 'Пароли не совпадают'
  return null
} 