import { useAuthStore } from '../stores/auth'

const GOOGLE_CLIENT_ID = import.meta.env.VITE_GOOGLE_CLIENT_ID
const GOOGLE_REDIRECT_URI = import.meta.env.VITE_GOOGLE_REDIRECT_URI

export const initGoogleAuth = () => {
  const script = document.createElement('script')
  script.src = 'https://accounts.google.com/gsi/client'
  script.async = true
  script.defer = true
  document.head.appendChild(script)

  return new Promise((resolve) => {
    script.onload = resolve
  })
}

export const handleGoogleLogin = async () => {
  try {
    const client = google.accounts.oauth2.initTokenClient({
      client_id: GOOGLE_CLIENT_ID,
      scope: 'email profile',
      callback: async (response: any) => {
        if (response.access_token) {
          // Здесь должна быть отправка токена на ваш бэкенд для верификации
          // и получения JWT токена для вашего приложения
          const authStore = useAuthStore()
          // Временное решение - используем токен Google как JWT
          authStore.login(response.access_token)
        }
      },
    })

    client.requestAccessToken()
  } catch (error) {
    console.error('Google login error:', error)
    throw error
  }
} 