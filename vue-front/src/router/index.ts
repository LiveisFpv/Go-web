import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '../stores/auth'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
      meta: { requiresAuth: false }
    },
    {
      path: '/auth',
      name: 'auth-page',
      component: () => import('../views/AuthPageView.vue'),
      meta: { requiresAuth: false }
    },
    {
      path: '/student',
      name: 'student-page',
      component: () => import('../views/StudentPageView.vue'),
      meta: { requiresAuth: true }
    },
    {
      path: '/group',
      name: 'group-page',
      component: () => import('../views/GroupPageView.vue'),
      meta: { requiresAuth: true }
    },
    {
      path: '/mark',
      name: 'mark-page',
      component: () => import('../views/MarkPageView.vue'),
      meta: { requiresAuth: true }
    },
    {
      path: '/semester',
      name: 'semester-page',
      component: () => import('../views/SemesterPageView.vue'),
      meta: { requiresAuth: true }
    },
    {
      path: '/scholarship',
      name: 'scholarship-page',
      component: () => import('../views/ScholarshipPageView.vue'),
      meta: { requiresAuth: true }
    },
    {
      path: '/budget',
      name: 'budget-page',
      component: () => import('../views/BudgetPageView.vue'),
      meta: { requiresAuth: true }
    },
    {
      path: '/achievement',
      name: 'achievement-page',
      component: () => import('../views/AchievementPageView.vue'),
      meta: { requiresAuth: true }
    },
    {
      path: '/category',
      name: 'category-page',
      component: () => import('../views/CategoryPageView.vue'),
      meta: { requiresAuth: true }
    }
  ]
})

router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()
  
  // Initialize auth state
  authStore.initialize()

  if (to.meta.requiresAuth && !authStore.isAuthenticated) {
    // Redirect to auth page if trying to access protected route while not authenticated
    next({ name: 'auth-page' })
  } else {
    next()
  }
})

export default router
