<script setup lang="ts">
import { ref } from 'vue'
import { RouterLink, RouterView, useRouter } from 'vue-router'
import { useAuthStore } from './stores/auth'

const isMenuOpen = ref(false)
const authStore = useAuthStore()
const router = useRouter()

const toggleMenu = () => {
  isMenuOpen.value = !isMenuOpen.value
}

const handleLogout = () => {
  authStore.logout()
  router.push('/auth')
}

// Initialize auth state when component is mounted
authStore.initialize()
</script>

<template>
  <header>
    <div class="mobile-header">
      <button class="hamburger" :class="{ rotated: isMenuOpen }" @click="toggleMenu">☰</button>
      <div class="panel">
        <button @click="handleLogout" class="logout-btn" v-if="authStore.isAuthenticated">Выйти</button>
        <img src="./assets/avatar.jpg" alt="Avatar" class="logo" v-if="authStore.isAuthenticated"/>
      </div>
    </div>
    <div class="wrapper" :class="{ open: isMenuOpen }">
      <nav>
        <RouterLink to="/">Главная</RouterLink>
        <RouterLink to="/auth" v-if="!authStore.isAuthenticated">Авторизация</RouterLink>
        <template v-if="authStore.isAuthenticated">
          <RouterLink to="/student">Студенты</RouterLink>
          <RouterLink to="/group">Группы</RouterLink>
          <RouterLink to="/mark">Оценки</RouterLink>
          <RouterLink to="/semester">Семестры</RouterLink>
          <RouterLink to="/scholarship">Стипендии</RouterLink>
          <RouterLink to="/budget">Бюджет</RouterLink>
          <RouterLink to="/achievement">Достижения</RouterLink>
          <RouterLink to="/category">Категории достижений</RouterLink>
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
}

.mobile-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0.5rem 1rem;
}

.panel {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 4rem;
}

.logout-btn {
  background: var(--color-accent);
  color: white;
  border: none;
  padding: 0.75rem;
  border-radius: 15px;
  width: 100px;
  font-size: 1rem;
  cursor: pointer;
  transition: transform 0.3s ease;
}

.logout-btn:hover {
  transform: scale(1.1);
  font-weight: 600;
}

.hamburger {
  background: none;
  border: none;
  font-size: 2rem;
  cursor: pointer;
  margin-right: auto;
  padding: 0.5rem;
  transition: transform 0.3s ease;
  text-align: center;
  vertical-align: middle;
}

.hamburger.rotated {
  transform: rotate(90deg);
}

.wrapper {
  display: none;
  flex-direction: column;
  justify-content: space-around;
  align-items: flex-start;
  gap: 50px;
  margin-top: 0.5rem;
  position: fixed;
  top: 80px;
  left: 40px;
  background-color: var(--color-background);
  z-index: 1000;
  padding: 2rem 0rem;
  padding-bottom: 2rem;
  padding-right: 3rem;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.wrapper.open {
  display: flex;
}

nav {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  gap: 15px;
  font-size: 1rem;
  animation: fromUp 1s ease-in-out;
  margin-left: 1rem;
}

nav a.router-link-exact-active {
  color: var(--color-text);
}

nav a.router-link-exact-active:hover {
  background-color: transparent;
}

nav a {
  display: inline-block;
  padding: 0 1rem;
  border-left: 1px solid var(--color-border);
}

nav a:first-of-type {
  border: 0;
}

.none {
  display: none;
}

.logo {
  width: 64px;
  height: 64px;
  border-radius: 64px;
}

.icon {
  width: 32px;
  height: 32px;
  border-radius: 32px;
}

nav a:hover {
  transform: scale(1.1);
  font-weight: 600;
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
  .logo {
    width: 76px;
    height: 76px;
    border-radius: 80px;
  }
}

@media (max-width: 500px) {

  .mobile-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0.5rem 1rem;
  }

  .logo {
    width: 48px;
    height: 48px;
    border-radius: 50%;
  }

  #logo {
    display: none;
  }
}

@keyframes fromUp {
  from {
    opacity: 0;
    transform: translateY(-30px);
  }
  to {
    opacity: 1;
    transform: translateX(0);
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
