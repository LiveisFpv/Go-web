<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useAuthStore } from '../stores/auth'
import router from '@/router'

const authStore = useAuthStore()

const navigateTo = (route: string, roles: string[]) => {
  const userRole = authStore.user_role || 'USER'
  if (roles.includes(userRole)) {
    router.push(route)
  }
}

const features = [
  {
    title: 'Студенты',
    description: 'Управление данными студентов',
    icon: '👥',
    route: '/student',
    roles: ['ADMIN', 'DEAN']
  },
  {
    title: 'Группы',
    description: 'Управление учебными группами',
    icon: '👨‍🎓',
    route: '/group',
    roles: ['ADMIN']
  },
  {
    title: 'Оценки',
    description: 'Управление успеваемостью',
    icon: '📝',
    route: '/mark',
    roles: ['ADMIN', 'DEAN', 'STUDENT']
  },
  {
    title: 'Семестры',
    description: 'Управление учебными периодами',
    icon: '📅',
    route: '/semester',
    roles: ['ADMIN']
  },
  {
    title: 'Стипендии',
    description: 'Управление стипендиальными выплатами',
    icon: '💰',
    route: '/scholarship',
    roles: ['ADMIN', 'ACCOUNTANT', 'STUDENT']
  },
  {
    title: 'Бюджет',
    description: 'Управление бюджетными средствами',
    icon: '💵',
    route: '/budget',
    roles: ['ADMIN', 'ACCOUNTANT']
  },
  {
    title: 'Достижения',
    description: 'Управление достижениями студентов',
    icon: '🏆',
    route: '/achievement',
    roles: ['ADMIN', 'DEAN', 'STUDENT']
  },
  {
    title: 'Категории',
    description: 'Управление категориями достижений',
    icon: '📋',
    route: '/category',
    roles: ['ADMIN', 'DEAN', 'STUDENT']
  }
]

const isFeatureAccessible = (roles: string[]) => {
  const userRole = authStore.user_role || 'USER'
  return roles.includes(userRole)
}
</script>

<template>
  <div class="home">
    <main class="main-content">
      <section class="welcome-section">
        <h1>Стипендиатус</h1>
        <p class="welcome-text">Эффективное управление образовательным процессом</p>
      </section>

      <section class="features-section">
        <h2>Основные модули системы</h2>
        <div class="features-grid">
          <div v-for="feature in features"
               :key="feature.route"
               class="feature-card"
               :class="{ 'disabled': !isFeatureAccessible(feature.roles) }"
               @click="navigateTo(feature.route, feature.roles)">
            <div class="feature-icon">{{ feature.icon }}</div>
            <h3>{{ feature.title }}</h3>
            <p>{{ feature.description }}</p>
            <div v-if="!isFeatureAccessible(feature.roles)" class="access-denied">
              Нет доступа
            </div>
          </div>
        </div>
      </section>

      <section class="profile-section" v-if="authStore.isAuthenticated">
        <div class="profile-card" @click="navigateTo('/profile', ['ADMIN', 'DEAN', 'ACCOUNTANT', 'STUDENT', 'USER'])">
          <div class="profile-content">
            <h2>Личный кабинет</h2>
            <p>Управление личными данными и настройками</p>
          </div>
          <div class="profile-icon">👤</div>
        </div>
      </section>
    </main>
  </div>
</template>

<style scoped>
.home {
  min-height: 100vh;
  background-color: var(--color-background);
  padding: 2rem 0;
}

.main-content {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 2rem;
}

.welcome-section {
  text-align: center;
  margin-bottom: 4rem;
  padding: 3rem 0;
  background: linear-gradient(135deg, var(--color-background-soft) 0%, var(--color-background) 100%);
  border-radius: 16px;
}

.welcome-section h1 {
  font-size: 2.8rem;
  color: var(--color-heading);
  margin-bottom: 1rem;
}

.welcome-text {
  font-size: 1.3rem;
  color: var(--color-text);
  opacity: 0.9;
}

.features-section {
  margin-bottom: 4rem;
}

.features-section h2 {
  font-size: 2rem;
  color: var(--color-heading);
  margin-bottom: 2rem;
  text-align: center;
}

.features-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 2rem;
}

.feature-card {
  background-color: var(--color-background-soft);
  padding: 2rem;
  border-radius: 12px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
  cursor: pointer;
  text-align: center;
  border: 1px solid var(--color-border);
}

.feature-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 6px 12px rgba(0, 0, 0, 0.15);
  border-color: var(--color-accent);
}

.feature-icon {
  font-size: 2.5rem;
  margin-bottom: 1rem;
}

.feature-card h3 {
  font-size: 1.5rem;
  color: var(--color-heading);
  margin-bottom: 1rem;
}

.feature-card p {
  color: var(--color-text);
  line-height: 1.6;
}

.profile-section {
  margin-top: 3rem;
}

.profile-card {
  background: linear-gradient(135deg, var(--color-accent) 0%, var(--color-accent-hover) 100%);
  padding: 2rem;
  border-radius: 12px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  cursor: pointer;
  transition: all 0.3s ease;
}

.profile-card:hover {
  transform: translateY(-3px);
  box-shadow: 0 6px 12px rgba(0, 0, 0, 0.15);
}

.profile-content {
  color: white;
}

.profile-content h2 {
  font-size: 1.8rem;
  margin-bottom: 0.5rem;
}

.profile-content p {
  opacity: 0.9;
}

.profile-icon {
  font-size: 3rem;
}

@media (max-width: 768px) {
  .main-content {
    padding: 0 1rem;
  }

  .welcome-section {
    padding: 2rem 0;
  }

  .welcome-section h1 {
    font-size: 2rem;
  }

  .features-grid {
    grid-template-columns: 1fr;
  }

  .profile-card {
    flex-direction: column;
    text-align: center;
    gap: 1rem;
  }

  .profile-icon {
    order: -1;
  }
}

.feature-card.disabled {
  opacity: 0.7;
  cursor: not-allowed;
  position: relative;
}

.feature-card.disabled:hover {
  transform: none;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  border-color: var(--color-border);
}

.access-denied {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  background-color: rgba(220, 38, 38, 0.9);
  color: white;
  padding: 0.5rem 1rem;
  border-radius: 4px;
  font-size: 0.9rem;
  opacity: 0;
  transition: opacity 0.3s ease;
}

.feature-card.disabled:hover .access-denied {
  opacity: 1;
}
</style>
