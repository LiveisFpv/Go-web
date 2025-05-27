<script setup lang="ts">
import { onMounted, ref, computed } from 'vue'
import { RouterLink, RouterView, useRouter } from 'vue-router'
import { useAuthStore } from './stores/auth'

const isMenuOpen = ref(false)
const isProfileMenuOpen = ref(false)
const authStore = useAuthStore()
const router = useRouter()

const userInitial = computed(() => {
  return authStore.email ? authStore.email.charAt(0).toUpperCase() : ''
})

const toggleMenu = () => {
  isMenuOpen.value = !isMenuOpen.value
}

const toggleProfileMenu = () => {
  isProfileMenuOpen.value = !isProfileMenuOpen.value
}

const handleLogout = () => {
  authStore.logout()
  router.push('/auth')
  isProfileMenuOpen.value = false
}

onMounted(() =>{
// Initialize auth state when component is mounted
authStore.initialize()
})
</script>

<template>
  <header>
    <div class="mobile-header">
      <button class="hamburger" :class="{ rotated: isMenuOpen }" @click="toggleMenu">☰</button>
      <div class="panel" v-if="authStore.isAuthenticated">
        <RouterLink to="/profile" class="link-button">Профиль</RouterLink>
        <div class="profile-container">
          <div class="avatar-placeholder" @click="toggleProfileMenu">
            {{ userInitial }}
          </div>
          <div class="profile-menu" v-if="isProfileMenuOpen">
            <button @click="handleLogout" class="logout-btn">Выйти</button>
          </div>
        </div>
      </div>
    </div>
    <div class="wrapper" :class="{ open: isMenuOpen }">
      <nav>
        <RouterLink to="/">Главная</RouterLink>
        <RouterLink to="/about">О проекте</RouterLink>
        <RouterLink to="/auth" v-if="!authStore.isAuthenticated">Авторизация</RouterLink>
        <template v-if="authStore.isAuthenticated">
          <!-- ADMIN has access to everything -->
          <template v-if="authStore.user_role === 'ADMIN'">
            <RouterLink to="/student">Студенты</RouterLink>
            <RouterLink to="/group">Группы</RouterLink>
            <RouterLink to="/mark">Оценки</RouterLink>
            <RouterLink to="/semester">Семестры</RouterLink>
            <RouterLink to="/scholarship">Стипендии</RouterLink>
            <RouterLink to="/budget">Бюджет</RouterLink>
            <RouterLink to="/achievement">Достижения</RouterLink>
            <RouterLink to="/category">Категории достижений</RouterLink>
          </template>

          <!-- DEAN access -->
          <template v-else-if="authStore.user_role === 'DEAN'">
            <RouterLink to="/student">Студенты</RouterLink>
            <RouterLink to="/mark">Оценки</RouterLink>
            <RouterLink to="/achievement">Достижения</RouterLink>
            <RouterLink to="/category">Категории достижений</RouterLink>
          </template>

          <!-- ACCOUNTANT access -->
          <template v-else-if="authStore.user_role === 'ACCOUNTANT'">
            <RouterLink to="/budget">Бюджет</RouterLink>
            <RouterLink to="/scholarship">Стипендии</RouterLink>
          </template>

          <!-- STUDENT access -->
          <template v-else-if="authStore.user_role === 'STUDENT'">
            <RouterLink to="/achievement">Достижения</RouterLink>
            <RouterLink to="/category">Категории достижений</RouterLink>
            <RouterLink to="/scholarship">Стипендии</RouterLink>
            <RouterLink to="/mark">Оценки</RouterLink>
          </template>

          <!-- USER has access only to basic pages -->
          <template v-else>
            <!-- No additional links for basic USER role -->
          </template>
        </template>
      </nav>
    </div>
  </header>

  <footer>
    <a href="https://github.com/LiveisFpv" class="icon">
      <img src="./assets/github.png" alt="Github" class="icon animate-bounce-in" />
    </a>
  </footer>

  <RouterView />
</template>

<style scoped>
header {
  line-height: 1.5;
  max-height: 100vh;
  box-shadow: 4px 4px 6px rgba(0, 0, 0, 0.1);
  width: 100%;
  border-radius: 20px;
  padding: 0.5rem;
  margin-bottom: 2rem;
}

.mobile-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0.3rem 0.8rem;
}

.panel {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 4rem;
}

.link-button{
  text-decoration: none;
  align-items: center;
  justify-self: center;
  text-align: center;
  vertical-align: middle;
  background: linear-gradient(135deg, var(--color-accent) 0%, var(--color-accent-hover) 100%);
  color: white;
  border: none;
  padding: 0.4rem;
  border-radius: 8px;
  width: 90px;
  font-size: 0.9rem;
  cursor: pointer;
  transition: transform 0.2s ease;
}

.logout-btn {
  background: linear-gradient(135deg, var(--color-accent) 0%, var(--color-accent-hover) 100%);
  color: white;
  border: none;
  padding: 0.55rem;
  border-radius: 8px;
  width: 90px;
  font-size: 0.9rem;
  cursor: pointer;
  transition: transform 0.2s ease;
}

.logout-btn:hover, .link-button:hover {
  transform: scale(1.05);
  font-weight: 500;
}

.hamburger {
  background: none;
  border: none;
  font-size: 1.5rem;
  cursor: pointer;
  margin-right: auto;
  padding: 0.3rem;
  transition: transform 0.2s ease;
  text-align: center;
  vertical-align: middle;
}

.hamburger.rotated {
  transform: rotate(90deg);
}

.wrapper {
  display: none;
  flex-direction: column;
  position: absolute;
  top: 95px;
  left: 3rem;
  background-color: var(--color-background);
  z-index: 1000;
  padding: 0.5rem;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  min-width: 200px;
}

.wrapper.open {
  display: flex;
  animation: fadeIn 0.2s ease-out;
}

nav {
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
}

nav a {
  display: block;
  padding: 0.5rem 1rem;
  text-decoration: none;
  color: var(--color-text);
  transition: all 0.2s ease;
  border-radius: 4px;
}

nav a:hover {
  background-color: var(--color-background-soft);
  color: var(--color-accent);
}

nav a.router-link-exact-active {
  color: var(--color-accent);
  font-weight: 500;
}

.none {
  display: none;
}

.avatar-placeholder {
  width: 40px;
  height: 40px;
  background: linear-gradient(135deg, var(--color-accent) 0%, var(--color-accent-hover) 100%);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
  font-weight: 500;
  color: white;
  border: 2px solid rgba(255, 255, 255, 0.3);
  cursor: pointer;
  transition: transform 0.2s ease;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.avatar-placeholder:hover {
  transform: scale(1.05);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.profile-container {
  position: relative;
}

.profile-menu {
  position: absolute;
  top: 100%;
  right: -2rem;
  background-color: var(--color-background);
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  padding: 0.5rem;
  margin-top: 0.5rem;
  z-index: 1000;
  animation: fadeIn 0.2s ease-out;
}

.logo {
  cursor: pointer;
  transition: transform 0.2s ease;
}

.logo:hover {
  transform: scale(1.05);
}

.icon {
  width: 32px;
  height: 32px;
  border-radius: 32px;
}

.animate-bounce-in {
  animation: bounce-in 0.6s ease-out;
}

footer {
  position: fixed;
  right: 0;
  bottom: 0;
  width: 100%;
  padding: 1rem;
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  background-color: var(--color-background-soft);
  z-index: 100;
  animation: appear 1s ease-in-out;
}

@keyframes bounce-in {
  0% {
    transform: scale(0.5);
    opacity: 0;
  }
  100% {
    transform: scale(1);
    opacity: 1;
  }
}

@media (min-width: 768px) {
  .avatar-placeholder {
    width: 48px;
    height: 48px;
    font-size: 24px;
  }
}

@media (max-width: 500px) {
  .mobile-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0.3rem 0.8rem;
  }

  .avatar-placeholder {
    width: 36px;
    height: 36px;
    font-size: 18px;
  }

  #logo {
    display: none;
  }
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(-5px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes appear {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}

</style>
