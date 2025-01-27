<template>
  <nav class="navbar">
    <div class="navbar-brand">
      <router-link to="/" class="logo">
        Система заявок
      </router-link>
    </div>

    <div class="navbar-menu">
      <template v-if="isAuthenticated">
        <router-link to="/dashboard" class="nav-link">
          Панель управления
        </router-link>
        <div class="user-menu">
          <span class="username">{{ user?.name }}</span>
          <button @click="logout" class="logout-btn">
            Выйти
          </button>
        </div>
      </template>
      <template v-else>
        <router-link to="/login" class="nav-link">
          Вход
        </router-link>
        <router-link to="/register" class="nav-link">
          Регистрация
        </router-link>
      </template>
    </div>
  </nav>
</template>

<script setup>
import { computed } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const isAuthenticated = computed(() => authStore.isAuthenticated)
const user = computed(() => authStore.userProfile)

const logout = () => {
  authStore.logout()
  router.push('/login')
}
</script>

<style scoped lang="scss">
.navbar {
  background: white;
  padding: 1rem 2rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.navbar-brand {
  .logo {
    font-size: 1.5rem;
    font-weight: bold;
    color: #007bff;
    text-decoration: none;
    
    &:hover {
      color: #0056b3;
    }
  }
}

.navbar-menu {
  display: flex;
  align-items: center;
  gap: 1.5rem;
}

.nav-link {
  color: #333;
  text-decoration: none;
  font-weight: 500;
  
  &:hover {
    color: #007bff;
  }

  &.router-link-active {
    color: #007bff;
  }
}

.user-menu {
  display: flex;
  align-items: center;
  gap: 1rem;

  .username {
    font-weight: 500;
    color: #333;
  }

  .logout-btn {
    padding: 0.5rem 1rem;
    background: #dc3545;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    font-weight: 500;

    &:hover {
      background: #c82333;
    }
  }
}
</style>
