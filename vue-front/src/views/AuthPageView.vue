<script setup lang="ts">
import { ref, onMounted, computed, watch } from 'vue'
import { useRouter } from 'vue-router'
import { initGoogleAuth, handleGoogleLogin } from '../services/googleAuth'
import {
  validateEmail,
  validateUsername,
  validatePassword,
  validatePasswordConfirmation,
  type ValidationError
} from '../utils/validation'
import { login, register } from '../services/authService'
import type { AuthError } from '../types/auth'

const router = useRouter()

const isLogin = ref(true)
const email = ref('')
const password = ref('')
const confirmPassword = ref('')
const error = ref('')
const validationErrors = ref<ValidationError[]>([])

onMounted(async () => {
  await initGoogleAuth()
})

const toggleForm = () => {
  isLogin.value = !isLogin.value
  error.value = ''
  validationErrors.value = []
  email.value = ''
  password.value = ''
  confirmPassword.value = ''
}

// Валидация в реальном времени
watch([email], () => {
  const emailError = validateEmail(email.value)
  updateValidationError('email', emailError)
})

watch(password, () => {
  const passwordError = validatePassword(password.value)
  updateValidationError('password', passwordError)

  // Если есть подтверждение пароля, проверяем его тоже
  if (confirmPassword.value) {
    const confirmationError = validatePasswordConfirmation(password.value, confirmPassword.value)
    updateValidationError('confirmPassword', confirmationError)
  }
})

watch(confirmPassword, () => {
  const confirmationError = validatePasswordConfirmation(password.value, confirmPassword.value)
  updateValidationError('confirmPassword', confirmationError)
})

const updateValidationError = (field: string, error: string | null) => {
  const existingErrorIndex = validationErrors.value.findIndex(e => e.field === field)

  if (error) {
    if (existingErrorIndex === -1) {
      validationErrors.value.push({ field, message: error })
    } else {
      validationErrors.value[existingErrorIndex].message = error
    }
  } else {
    if (existingErrorIndex !== -1) {
      validationErrors.value.splice(existingErrorIndex, 1)
    }
  }
}

const validateForm = (): boolean => {
  // Проверяем все поля перед отправкой
  const emailError = validateEmail(email.value)
  updateValidationError('email', emailError)

  const passwordError = validatePassword(password.value)
  updateValidationError('password', passwordError)

  if (!isLogin.value) {
    const confirmationError = validatePasswordConfirmation(password.value, confirmPassword.value)
    updateValidationError('confirmPassword', confirmationError)
  }

  return validationErrors.value.length === 0
}

const handleLogin = async () => {
  if (!validateForm()) return

  try {
    const response = await login({
      email: email.value,
      password: password.value
    })
    router.push('/')
  } catch (e) {
    const error = e as AuthError
    if (error.field) {
      validationErrors.value.push({
        field: error.field,
        message: error.message
      })
    } else {
      error.value = error.message
    }
  }
}

const handleRegister = async () => {
  if (!validateForm()) return

  try {
    const response = await register({
      email: email.value,
      password: password.value
    })
    router.push('/')
  } catch (e) {
    const error = e as AuthError
    if (error.field) {
      validationErrors.value.push({
        field: error.field,
        message: error.message
      })
    } else {
      error.value = error.message
    }
  }
}

const handleGoogleAuth = async () => {
  try {
    await handleGoogleLogin()
    router.push('/')
  } catch (e) {
    error.value = 'Ошибка авторизации через Google'
  }
}

const getFieldError = (field: string) => {
  return validationErrors.value.find(error => error.field === field)?.message
}
</script>

<template>
  <div class="auth">
    <div class="auth-container">
      <h1>{{ isLogin ? 'Авторизация' : 'Регистрация' }}</h1>
      <form @submit.prevent="isLogin ? handleLogin() : handleRegister()" class="auth-form">
        <div class="form-group">
          <label for="email">Email</label>
          <input
            type="email"
            id="email"
            v-model="email"
            required
            placeholder="Введите email"
            :class="{ 'error-input': getFieldError('email') }"
          />
          <span v-if="getFieldError('email')" class="field-error">
            {{ getFieldError('email') }}
          </span>
        </div>
        <div class="form-group">
          <label for="password">Пароль</label>
          <input
            type="password"
            id="password"
            v-model="password"
            required
            placeholder="Введите пароль"
            :class="{ 'error-input': getFieldError('password') }"
          />
          <span v-if="getFieldError('password')" class="field-error">
            {{ getFieldError('password') }}
          </span>
        </div>
        <div class="form-group" v-if="!isLogin">
          <label for="confirmPassword">Подтвердите пароль</label>
          <input
            type="password"
            id="confirmPassword"
            v-model="confirmPassword"
            required
            placeholder="Повторите пароль"
            :class="{ 'error-input': getFieldError('confirmPassword') }"
          />
          <span v-if="getFieldError('confirmPassword')" class="field-error">
            {{ getFieldError('confirmPassword') }}
          </span>
        </div>
        <div v-if="error" class="error">{{ error }}</div>
        <button type="submit" class="submit-btn">
          {{ isLogin ? 'Войти' : 'Зарегистрироваться' }}
        </button>
        <div class="divider">
          <span>или</span>
        </div>
        <button type="button" @click="handleGoogleAuth" class="google-btn">
          <img src="../assets/google.svg" alt="Google" class="google-icon" />
          Войти через Google
        </button>
        <button type="button" @click="toggleForm" class="toggle-btn">
          {{ isLogin ? 'Создать аккаунт' : 'Уже есть аккаунт? Войти' }}
        </button>
        <p id="text" v-if="!isLogin">By registering, you agree to our <a href="#">Terms of Service</a> and <a href="#">Privacy Policy</a>.</p>
      </form>
    </div>
  </div>
</template>

<style scoped>
.auth {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 80vh;
}

.auth-container {
  background: var(--color-background-soft);
  padding: 2rem;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  width: 100%;
  max-width: 400px;
}

.auth-form {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

label {
  font-weight: 500;
}

input,select {
  padding: 0.5rem;
  border: 1px solid var(--color-border);
  border-radius: 4px;
  font-size: 1rem;
}

input.error-input {
  border-color: #dc2626;
}

.field-error {
  color: #dc2626;
  font-size: 0.75rem;
  margin-top: -0.25rem;
}

.submit-btn {
  background: var(--color-accent);
  color: white;
  border: none;
  padding: 0.75rem;
  border-radius: 4px;
  font-size: 1rem;
  cursor: pointer;
  transition: opacity 0.2s;
}

.submit-btn:hover {
  opacity: 0.9;
}

.google-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
  background: white;
  border: 1px solid var(--color-border);
  padding: 0.75rem;
  border-radius: 4px;
  font-size: 1rem;
  cursor: pointer;
  transition: background-color 0.2s;
}

.google-btn:hover {
  background-color: #f8f8f8;
}

.google-icon {
  width: 24px;
  height: 24px;
}

.divider {
  display: flex;
  align-items: center;
  text-align: center;
  color: var(--color-text);
  margin: 0.5rem 0;
}

.divider::before,
.divider::after {
  content: '';
  flex: 1;
  border-bottom: 1px solid var(--color-border);
}

.divider span {
  padding: 0 1rem;
}

.toggle-btn {
  background: none;
  border: none;
  color: var(--color-accent);
  padding: 0.5rem;
  font-size: 0.875rem;
  cursor: pointer;
  text-decoration: underline;
}

.toggle-btn:hover {
  opacity: 0.8;
}

.error {
  color: #dc2626;
  font-size: 0.875rem;
}

#text{
  font-size: 0.75rem;
  text-align: center;
}
</style>
