<template>
  <div class="dashboard">
    <header class="dashboard-header">
      <h1>Панель управления</h1>
      <!-- <div class="user-info">
        <span>{{ user?.name }}</span>
        <button @click="logout" class="logout-btn">Выйти</button>
      </div> -->
    </header>

    <main class="dashboard-content">
      <div class="dashboard-stats">
        <div class="stat-card">
          <h3>Всего заявок</h3>
          <p>{{ requestsCount }}</p>
        </div>
        <div class="stat-card">
          <h3>В обработке</h3>
          <p>{{ pendingCount }}</p>
        </div>
        <div class="stat-card">
          <h3>Завершено</h3>
          <p>{{ completedCount }}</p>
        </div>
      </div>

      <div class="requests-section">
        <h2>Мои заявки</h2>
        <button @click="showNewRequestModal = true" class="new-request-btn">
          Новая заявка
        </button>
        
        <div class="requests-list">
          <div v-if="isLoading" class="loading">Загрузка...</div>
          <div v-else-if="requests.length === 0" class="no-requests">
            У вас пока нет заявок
          </div>
          <div v-else v-for="request in requests" :key="request.id" 
               class="request-card" 
               @click="viewRequestDetails(request.id)">
            <div class="request-header">
              <h3>{{ request.title }}</h3>
              <span v-if="isAdmin" class="author">Автор: {{ request.user.name }}</span>
            </div>
            <p>{{ request.description }}</p>
            <div class="request-footer">
              <span :class="['status', request.status.toLowerCase()]">{{ getStatusText(request.status) }}</span>
              <span class="date">{{ formatDate(request.createdAt) }}</span>
            </div>
          </div>
        </div>
      </div>
    </main>

    <!-- Модальное окно новой заявки -->
    <div v-if="showNewRequestModal" class="modal">
      <div class="modal-content">
        <h2>Новая заявка</h2>
        <form @submit.prevent="submitRequest">
          <div class="form-group">
            <label for="title">Заголовок</label>
            <input
              type="text"
              id="title"
              v-model="newRequest.title"
              required
            >
          </div>

          <div class="form-group">
            <label for="description">Описание</label>
            <textarea
              id="description"
              v-model="newRequest.description"
              required
            ></textarea>
          </div>

          <div class="form-group">
            <label for="attachment">Прикрепить файл</label>
            <input
              type="file"
              id="attachment"
              @change="handleFileUpload"
            >
          </div>

          <div class="modal-actions">
            <button type="button" @click="showNewRequestModal = false">
              Отмена
            </button>
            <button type="submit" :disabled="isSubmitting">
              {{ isSubmitting ? 'Отправка...' : 'Отправить' }}
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useToast } from 'vue-toastification'
import { useAuthStore } from '@/stores/auth'
import axios from 'axios'

const router = useRouter()
const toast = useToast()
const authStore = useAuthStore()

const user = computed(() => authStore.userProfile)
const requests = ref([])
const isLoading = ref(true)
const showNewRequestModal = ref(false)
const isSubmitting = ref(false)

const newRequest = ref({
  title: '',
  description: '',
  attachment: null
})

const isAdmin = computed(() => user.value?.role === 'admin')

const requestsCount = computed(() => requests.value.length)
const pendingCount = computed(() => 
  requests.value.filter(r => r.status === 'NEW' || r.status === 'IN_PROGRESS').length
)
const completedCount = computed(() => 
  requests.value.filter(r => r.status === 'RESOLVED' || r.status === 'CLOSED').length
)

const fetchRequests = async () => {
  try {
    isLoading.value = true
    const response = await axios.get('/users/requests', {
      headers: {
        Authorization: `Bearer ${localStorage.getItem('token')}`
      }
    })
    console.log('Полученные заявки:', response.data)
    requests.value = response.data
    
  } catch (error) {
    console.error('Ошибка при загрузке заявок:', error)
    toast.error('Ошибка при загрузке заявок')
  } finally {
    isLoading.value = false
  }
}

const handleFileUpload = (event) => {
  newRequest.value.attachment = event.target.files[0]
}

const submitRequest = async () => {
  try {
    isSubmitting.value = true
    
    const formData = new FormData()
    formData.append('title', newRequest.value.title)
    formData.append('description', newRequest.value.description)
    if (newRequest.value.attachment) {
      formData.append('attachment', newRequest.value.attachment)
    }

    await axios.post('/requests/submit', formData, {
      headers: {
        Authorization: `Bearer ${localStorage.getItem('token')}`,
        'Content-Type': 'multipart/form-data'
      }
    })

    toast.success('Заявка успешно создана')
    showNewRequestModal.value = false
    newRequest.value = { title: '', description: '', attachment: null }
    await fetchRequests()
  } catch (error) {
    console.error('Ошибка при создании заявки:', error)
    toast.error('Ошибка при создании заявки')
  } finally {
    isSubmitting.value = false
  }
}

const formatDate = (date) => {
  return new Date(date).toLocaleDateString('ru-RU', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
}

const logout = () => {
  authStore.logout()
  router.push('/login')
}

const viewRequestDetails = (requestId) => {
  console.log('Переход к заявке с ID:', requestId)
  if (requestId) {
    router.push(`/requests/${requestId}`)
  } else {
    console.error('ID заявки не определен')
    toast.error('Ошибка: ID заявки не определен')
  }
}

const getStatusText = (status) => {
  const statusMap = {
    'NEW': 'Новая',
    'IN_PROGRESS': 'В работе',
    'RESOLVED': 'Решена',
    'CLOSED': 'Закрыта'
  }
  return statusMap[status] || status
}

onMounted(async () => {
  if (!authStore.isAuthenticated) {
    router.push('/login')
    return
  }
  await fetchRequests()
})
</script>

<style scoped lang="scss">
.dashboard {
  min-height: 100vh;
  background-color: #f5f5f5;
}

.dashboard-header {
  background: white;
  padding: 1rem 2rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);

  h1 {
    margin: 0;
    color: #333;
  }
}

.user-info {
  display: flex;
  align-items: center;
  gap: 1rem;

  .logout-btn {
    padding: 0.5rem 1rem;
    background: #dc3545;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;

    &:hover {
      background: #c82333;
    }
  }
}

.dashboard-content {
  padding: 2rem;
  max-width: 1200px;
  margin: 0 auto;
}

.dashboard-stats {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1rem;
  margin-bottom: 2rem;
}

.stat-card {
  background: white;
  padding: 1.5rem;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);

  h3 {
    margin: 0 0 0.5rem;
    color: #666;
  }

  p {
    margin: 0;
    font-size: 2rem;
    font-weight: bold;
    color: #007bff;
  }
}

.requests-section {
  background: white;
  padding: 2rem;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);

  h2 {
    margin: 0 0 1rem;
  }
}

.new-request-btn {
  padding: 0.75rem 1.5rem;
  background: #28a745;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  margin-bottom: 1rem;

  &:hover {
    background: #218838;
  }
}

.requests-list {
  display: grid;
  gap: 1rem;
}

.request-card {
  background: #fff;
  border-radius: 8px;
  padding: 1.5rem;
  margin-bottom: 1rem;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  transition: transform 0.2s, box-shadow 0.2s;
  cursor: pointer;
}

.request-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.15);
}

.request-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1rem;
}

.request-header h3 {
  margin: 0;
  color: #2c3e50;
}

.author {
  font-size: 0.875rem;
  color: #666;
  background-color: #f5f5f5;
  padding: 0.25rem 0.5rem;
  border-radius: 4px;
}

.request-card p {
  margin: 0 0 1rem 0;
  color: #666;
}

.request-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.status {
  padding: 0.25rem 0.75rem;
  border-radius: 4px;
  font-size: 0.875rem;
  font-weight: 500;
}

.status.new {
  background-color: #e3f2fd;
  color: #1976d2;
}

.status.in_progress {
  background-color: #fff3e0;
  color: #f57c00;
}

.status.resolved {
  background-color: #e8f5e9;
  color: #388e3c;
}

.status.closed {
  background-color: #efebe9;
  color: #5d4037;
}

.date {
  color: #666;
  font-size: 0.875rem;
}

.modal {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0,0,0,0.5);
  display: flex;
  align-items: center;
  justify-content: center;
}

.modal-content {
  background: white;
  padding: 2rem;
  border-radius: 8px;
  width: 100%;
  max-width: 500px;

  h2 {
    margin: 0 0 1.5rem;
  }
}

.form-group {
  margin-bottom: 1rem;

  label {
    display: block;
    margin-bottom: 0.5rem;
    font-weight: 500;
  }

  input, textarea {
    width: 100%;
    padding: 0.75rem;
    border: 1px solid #ddd;
    border-radius: 4px;
    font-size: 1rem;

    &:focus {
      outline: none;
      border-color: #007bff;
    }
  }

  textarea {
    min-height: 100px;
    resize: vertical;
  }
}

.modal-actions {
  display: flex;
  justify-content: flex-end;
  gap: 1rem;
  margin-top: 1.5rem;

  button {
    padding: 0.75rem 1.5rem;
    border: none;
    border-radius: 4px;
    cursor: pointer;

    &[type="button"] {
      background: #6c757d;
      color: white;

      &:hover {
        background: #5a6268;
      }
    }

    &[type="submit"] {
      background: #007bff;
      color: white;

      &:hover {
        background: #0056b3;
      }

      &:disabled {
        background: #ccc;
        cursor: not-allowed;
      }
    }
  }
}

.loading {
  text-align: center;
  padding: 2rem;
  color: #666;
}

.no-requests {
  text-align: center;
  padding: 2rem;
  color: #666;
  background: #f8f9fa;
  border-radius: 4px;
}
</style>
