import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '../stores/auth'
import HomeView from '../views/HomeView.vue'

type UserRole = 'ADMIN' | 'DEAN' | 'ACCOUNTANT' | 'STUDENT' | 'USER'

// Define allowed routes for each role
const roleRoutes: Record<UserRole, string[]> = {
  ADMIN: ['/', '/about', '/student', '/group', '/mark', '/semester', '/scholarship', '/budget', '/achievement', '/category', '/profile'],
  DEAN: ['/', '/about', '/student', '/mark', '/achievement', '/category', '/profile'],
  ACCOUNTANT: ['/', '/about', '/budget', '/scholarship', '/profile'],
  STUDENT: ['/', '/about', '/achievement', '/category', '/scholarship', '/mark', '/profile'],
  USER: ['/', '/about', '/profile']
}

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
      path: '/about',
      name: 'about-page',
      component: () => import('../views/AboutPageView.vue'),
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
    },
    {
      path: '/profile',
      name: 'profile-page',
      component: () => import('../views/ProfilePageView.vue'),
      meta: { requiresAuth: true }
    }
  ]
})

router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()

  authStore.initialize()

  // Check if route requires authentication
  if (to.meta.requiresAuth && !authStore.isAuthenticated) {
    // Redirect to auth page if trying to access protected route while not authenticated
    next({ name: 'auth-page' })
    return
  }

  // If user is authenticated, check role-based access
  if (authStore.isAuthenticated) {
    const userRole = (authStore.user_role || 'USER') as UserRole
    const allowedRoutes = roleRoutes[userRole]

    // If route is not in allowed routes for user's role, redirect to home
    if (!allowedRoutes.includes(to.path)) {
      next({ name: 'home' })
      return
    }
  }

  next()
})

export default router
