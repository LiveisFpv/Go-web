<script setup lang="ts">
import { onMounted, ref } from 'vue'
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
        <button @click="handleLogout" class="logout-btn">Выйти</button>
        <RouterLink to="/profile" class="link-button">Профиль</RouterLink>
        <img src="./assets/avatar.jpg" alt="Avatar" class="logo"/>
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
  background: var(--color-accent);
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
  background: var(--color-accent);
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

.logo {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  object-fit: cover;
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
  .logo {
    width: 48px;
    height: 48px;
    border-radius: 50%;
  }
}

@media (max-width: 500px) {
  .mobile-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0.3rem 0.8rem;
  }

  .logo {
    width: 36px;
    height: 36px;
    border-radius: 50%;
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
