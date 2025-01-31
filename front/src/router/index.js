import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      redirect: '/general'
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('@/views/LoginView.vue'),
      meta: { requiresGuest: true }
    },
    {
      path: '/register',
      name: 'register',
      component: () => import('@/views/RegisterView.vue'),
      meta: { requiresGuest: true }
    },
    {
      path: '/general',
      name: 'general',
      component: () => import('../views/GeneralView.vue'),
    },
    {
      path: '/about_us',
      name: 'about_us',
      component: () => import('@/views/AboutUsView.vue'),
    },
    {
      path: '/dashboard',
      name: 'dashboard',
      component: () => import('@/views/DashboardView.vue'),
      meta: { requiresAuth: true }
    },
    {
      path: '/requests/:id',
      name: 'request-details',
      component: () => import('@/views/RequestDetails.vue'),
      meta: { requiresAuth: true },
      props: true
    },
    {
      path: '/:pathMatch(.*)*',
      redirect: '/dashboard'
    }
  ]
})

// Навигационный guard для проверки авторизации
router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()
  const isAuthenticated = authStore.isAuthenticated

  // Проверяем, требует ли маршрут авторизации
  if (to.meta.requiresAuth && !isAuthenticated) {
    next('/login')
    return
  }

  // Проверяем, требует ли маршрут гостевого доступа
  if (to.meta.requiresGuest && isAuthenticated) {
    next('/dashboard')
    return
  }

  next()
})

export default router
