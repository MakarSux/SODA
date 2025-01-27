<template>
  <div class="auth-form">
    <h2>Регистрация</h2>
    <form @submit.prevent="handleSubmit">
      <div class="form-group">
        <label for="name">Имя</label>
        <input
          type="text"
          id="name"
          v-model="formData.name"
          required
          placeholder="Введите ваше имя"
        >
      </div>

      <div class="form-group">
        <label for="email">Email</label>
        <input
          type="email"
          id="email"
          v-model="formData.email"
          required
          placeholder="Введите email"
        >
      </div>

      <div class="form-group">
        <label for="password">Пароль</label>
        <input
          type="password"
          id="password"
          v-model="formData.password"
          required
          placeholder="Введите пароль"
          minlength="6"
        >
      </div>

      <div class="form-group">
        <label for="confirmPassword">Подтверждение пароля</label>
        <input
          type="password"
          id="confirmPassword"
          v-model="formData.confirmPassword"
          required
          placeholder="Подтвердите пароль"
        >
      </div>

      <button type="submit" class="submit-btn" :disabled="isLoading || !isValid">
        {{ isLoading ? 'Регистрация...' : 'Зарегистрироваться' }}
      </button>

      <p class="auth-links">
        Уже есть аккаунт? <router-link to="/login">Войти</router-link>
      </p>
    </form>
  </div>
</template>

<script setup>
import { ref, reactive, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useToast } from 'vue-toastification'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const toast = useToast()
const authStore = useAuthStore()

const formData = reactive({
  name: '',
  email: '',
  password: '',
  confirmPassword: ''
})

const isLoading = ref(false)

const isValid = computed(() => {
  return formData.password === formData.confirmPassword &&
         formData.password.length >= 6 &&
         formData.name.trim() &&
         formData.email.includes('@')
})

const handleSubmit = async () => {
  if (!isValid.value) {
    toast.error('Пожалуйста, проверьте правильность заполнения формы')
    return
  }

  try {
    isLoading.value = true
    await authStore.register({
      name: formData.name,
      email: formData.email,
      password: formData.password
    })
    toast.success('Регистрация успешна!')
    router.push('/dashboard')
  } catch (error) {
    toast.error(error.response?.data?.message || 'Ошибка при регистрации')
  } finally {
    isLoading.value = false
  }
}
</script>

<style scoped lang="scss">
.auth-form {
  max-width: 400px;
  margin: 2rem auto;
  padding: 2rem;
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0,0,0,0.1);

  h2 {
    text-align: center;
    margin-bottom: 2rem;
    color: #333;
  }
}

.form-group {
  margin-bottom: 1.5rem;

  label {
    display: block;
    margin-bottom: 0.5rem;
    font-weight: 500;
    color: #333;
  }

  input {
    width: 100%;
    padding: 0.75rem;
    border: 1px solid #ddd;
    border-radius: 4px;
    font-size: 1rem;
    transition: border-color 0.3s ease;

    &:focus {
      outline: none;
      border-color: #007bff;
    }
  }
}

.submit-btn {
  width: 100%;
  padding: 0.75rem;
  background: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  font-size: 1rem;
  cursor: pointer;
  transition: background-color 0.3s ease;

  &:hover {
    background: #0056b3;
  }

  &:disabled {
    background: #ccc;
    cursor: not-allowed;
  }
}

.auth-links {
  text-align: center;
  margin-top: 1rem;
  
  a {
    color: #007bff;
    text-decoration: none;
    
    &:hover {
      text-decoration: underline;
    }
  }
}
</style>
