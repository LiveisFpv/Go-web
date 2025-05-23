<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useAuthStore } from '../stores/auth'

const authStore = useAuthStore()
const user = ref(authStore.user)

</script>

<template>
  <div class="home">

    <main class="main-content">
      <section class="welcome-section">
        <h1>Добро пожаловать!</h1>
        <p>Вы успешно авторизованы в системе.</p>
      </section>

      <section class="features-section">
        <h2>Доступные функции</h2>
        <div class="features-grid">
          <div class="feature-card">
            <h3>Профиль</h3>
            <p>Управление личными данными</p>
          </div>
          <div class="feature-card">
            <h3>Настройки</h3>
            <p>Настройка параметров системы</p>
          </div>
          <div class="feature-card">
            <h3>Статистика</h3>
            <p>Просмотр статистики использования</p>
          </div>
        </div>
      </section>
    </main>
  </div>
</template>

<style scoped>
.home {
  min-height: 100vh;
  background-color: var(--color-background);
}

.header {
  background-color: var(--color-background-soft);
  padding: 1rem 2rem;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.user-info {
  display: flex;
  justify-content: flex-end;
  align-items: center;
  gap: 1rem;
}

.user-email {
  color: var(--color-text);
  font-size: 1rem;
}

.logout-btn {
  background: none;
  border: 1px solid var(--color-border);
  color: var(--color-text);
  padding: 0.5rem 1rem;
  border-radius: 4px;
  cursor: pointer;
  transition: all 0.2s;
}

.logout-btn:hover {
  background-color: var(--color-background-soft);
  border-color: var(--color-accent);
  color: var(--color-accent);
}

.main-content {
  max-width: 1200px;
  margin: 0 auto;
  padding: 2rem;
}

.welcome-section {
  text-align: center;
  margin-bottom: 3rem;
}

.welcome-section h1 {
  font-size: 2.5rem;
  color: var(--color-heading);
  margin-bottom: 1rem;
}

.welcome-section p {
  font-size: 1.2rem;
  color: var(--color-text);
}

.features-section {
  margin-top: 2rem;
}

.features-section h2 {
  font-size: 2rem;
  color: var(--color-heading);
  margin-bottom: 2rem;
  text-align: center;
}

.features-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 2rem;
}

.feature-card {
  background-color: var(--color-background-soft);
  padding: 2rem;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  transition: transform 0.2s;
}

.feature-card:hover {
  transform: translateY(-5px);
}

.feature-card h3 {
  font-size: 1.5rem;
  color: var(--color-heading);
  margin-bottom: 1rem;
}

.feature-card p {
  color: var(--color-text);
  line-height: 1.5;
}

@media (max-width: 768px) {
  .main-content {
    padding: 1rem;
  }

  .features-grid {
    grid-template-columns: 1fr;
  }

  .welcome-section h1 {
    font-size: 2rem;
  }
}
</style>
