<template>
  <div class="request-details">
    <div v-if="loading" class="loading">
      Загрузка...
    </div>
    <div v-else-if="error" class="error">
      {{ error }}
    </div>
    <template v-else>
      <div class="request-header">
        <h1>{{ request.title }}</h1>
        <div class="status" :class="request.status.toLowerCase()">
          {{ getStatusText(request.status) }}
        </div>
      </div>

      <div class="request-info">
        <div class="info-item">
          <strong>Автор:</strong> {{ request.user.name }}
          <span v-if="isAdmin" class="user-email">({{ request.user.email }})</span>
        </div>
        <div class="info-item">
          <strong>Создано:</strong> {{ formatDate(request.createdAt) }}
        </div>
        <div class="info-item">
          <strong>Обновлено:</strong> {{ formatDate(request.updatedAt) }}
        </div>
      </div>

      <div class="request-description">
        <h2>Описание</h2>
        <p>{{ request.description }}</p>
      </div>

      <div v-if="request.attachments && request.attachments.length > 0" class="attachments">
        <h2>Вложения</h2>
        <ul>
          <li v-for="attachment in request.attachments" :key="attachment.id">
            <a :href="getAttachmentUrl(attachment)" target="_blank">
              {{ attachment.filename }}
            </a>
            <span class="file-size">({{ formatFileSize(attachment.size) }})</span>
          </li>
        </ul>
      </div>

      <!-- Добавляем компонент чата -->
      <div class="chat-section">
        <h2>Чат по заявке</h2>
        <ChatComponent :requestId="request.id" />
      </div>

      <!-- Кнопки управления статусом (только для админа) -->
      <div v-if="isAdmin" class="admin-controls">
        <h2>Управление статусом</h2>
        <div class="status-buttons">
          <button
            v-for="status in availableStatuses"
            :key="status.value"
            :class="['status-button', status.value.toLowerCase()]"
            :disabled="request.status === status.value"
            @click="updateStatus(status.value)"
          >
            {{ status.label }}
          </button>
        </div>
      </div>
    </template>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRoute } from 'vue-router'
import { useAuthStore } from '../stores/auth'
import axios from 'axios'
import ChatComponent from '../components/ChatComponent.vue'

const route = useRoute()
const authStore = useAuthStore()

const request = ref(null)
const loading = ref(true)
const error = ref(null)

const isAdmin = computed(() => authStore.user?.role === 'admin')

const availableStatuses = [
  { value: 'NEW', label: 'Новая' },
  { value: 'IN_PROGRESS', label: 'В работе' },
  { value: 'RESOLVED', label: 'Решена' },
  { value: 'CLOSED', label: 'Закрыта' }
]

const getStatusText = (status) => {
  const statusObj = availableStatuses.find(s => s.value === status)
  return statusObj ? statusObj.label : status
}

const formatDate = (dateString) => {
  return new Date(dateString).toLocaleString('ru-RU')
}

const formatFileSize = (bytes) => {
  if (bytes === 0) return '0 Bytes'
  const k = 1024
  const sizes = ['Bytes', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}

const getAttachmentUrl = (attachment) => {
  return `/api/requests/${request.value.id}/attachments/${attachment.id}`
}

const fetchRequest = async () => {
  loading.value = true
  error.value = null
  try {
    const response = await axios.get(`/requests/${route.params.id}`, {
      headers: {
        Authorization: `Bearer ${localStorage.getItem('token')}`
      }
    })
    request.value = response.data
  } catch (err) {
    error.value = err.response?.data?.error || 'Ошибка при загрузке заявки'
  } finally {
    loading.value = false
  }
}

const updateStatus = async (newStatus) => {
  try {
    await axios.put(`/requests/${request.value.id}/status`, {
      status: newStatus
    }, {
      headers: {
        Authorization: `Bearer ${localStorage.getItem('token')}`
      }
    })
    // Обновляем данные заявки после изменения статуса
    await fetchRequest()
  } catch (err) {
    error.value = err.response?.data?.error || 'Ошибка при обновлении статуса'
  }
}

onMounted(() => {
  fetchRequest()
})
</script>

<style scoped>
.request-details {
  max-width: 1000px;
  margin: 0 auto;
  padding: 2rem;
}

.request-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
}

.request-header h1 {
  margin: 0;
  font-size: 2rem;
}

.status {
  padding: 0.5rem 1rem;
  border-radius: 4px;
  font-weight: bold;
}

.status.new { background-color: #e3f2fd; color: #1976d2; }
.status.in_progress { background-color: #fff3e0; color: #f57c00; }
.status.resolved { background-color: #e8f5e9; color: #388e3c; }
.status.closed { background-color: #efebe9; color: #5d4037; }

.request-info {
  background-color: #f5f5f5;
  padding: 1rem;
  border-radius: 4px;
  margin-bottom: 2rem;
}

.info-item {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  margin-bottom: 0.5rem;
}

.user-email {
  font-size: 0.875rem;
  color: #666;
  margin-left: 0.5rem;
}

.request-description {
  margin-bottom: 2rem;
}

.attachments {
  margin-bottom: 2rem;
}

.attachments ul {
  list-style: none;
  padding: 0;
}

.attachments li {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  margin-bottom: 0.5rem;
}

.file-size {
  color: #666;
  font-size: 0.875rem;
}

.chat-section {
  margin-bottom: 2rem;
}

.admin-controls {
  margin-top: 2rem;
}

.status-buttons {
  display: flex;
  gap: 1rem;
  flex-wrap: wrap;
}

.status-button {
  padding: 0.5rem 1rem;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-weight: bold;
  transition: opacity 0.2s;
}

.status-button:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.status-button.new { background-color: #e3f2fd; color: #1976d2; }
.status-button.in_progress { background-color: #fff3e0; color: #f57c00; }
.status-button.resolved { background-color: #e8f5e9; color: #388e3c; }
.status-button.closed { background-color: #efebe9; color: #5d4037; }

.loading, .error {
  text-align: center;
  padding: 2rem;
}

.error {
  color: #dc3545;
}
</style>
