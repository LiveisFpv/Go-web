// Общий флаг для включения/отключения моков
export const USE_MOCKS = false

// Настройки для отдельных моков
export const MOCK_SETTINGS = {
  // Моки авторизации
  auth: {
    enabled: true,
    delay: 0
  },
  // Моки Google авторизации
  google: {
    enabled: true,
    delay: 0
  },
  // Моки пользовательских данных
  user: {
    enabled: true,
    delay: 0
  }
}

// Вспомогательная функция для получения задержки
export const getMockDelay = (mockType: keyof typeof MOCK_SETTINGS): number => {
  return MOCK_SETTINGS[mockType].enabled ? MOCK_SETTINGS[mockType].delay : 0
}

// Вспомогательная функция для проверки включен ли мок
export const isMockEnabled = (mockType: keyof typeof MOCK_SETTINGS): boolean => {
  return USE_MOCKS && MOCK_SETTINGS[mockType].enabled
}
