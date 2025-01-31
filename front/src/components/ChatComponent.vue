<template>
  <div class="chat-container">
    <div class="chat-messages" ref="messagesContainer">
      <div v-if="chatStore.loading" class="loading">
        Загрузка сообщений...
      </div>
      <div v-else-if="chatStore.error" class="error">
        {{ chatStore.error }}
      </div>
      <template v-else>
        <div v-for="message in chatStore.messages" :key="message.id" 
             :class="['message', { 'message-own': message.userId === authStore.user?.id }]">
          <div class="message-header">
            <span class="message-author">{{ message.user.name }}</span>
            <span class="message-time">{{ formatDate(message.createdAt) }}</span>
          </div>
          <div class="message-content">
            {{ message.message }}
          </div>
        </div>
      </template>
    </div>
    
    <div class="chat-input">
      <form @submit.prevent="sendMessage" class="input-form">
        <input
          v-model="newMessage"
          type="text"
          placeholder="Введите сообщение..."
          :disabled="chatStore.loading"
        />
        <button type="submit" :disabled="chatStore.loading || !newMessage.trim()">
          Отправить
        </button>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, watch } from 'vue'
import { useChatStore } from '../stores/chat'
import { useAuthStore } from '../stores/auth'
import { storeToRefs } from 'pinia'

const props = defineProps({
  requestId: {
    type: [Number, String],
    required: true
  }
})

const chatStore = useChatStore()
const authStore = useAuthStore()
const messagesContainer = ref(null)
const newMessage = ref('')

// Интервал обновления сообщений (каждые 5 секунд)
let updateInterval

const formatDate = (dateString) => {
  const date = new Date(dateString)
  return date.toLocaleString('ru-RU', {
    hour: '2-digit',
    minute: '2-digit',
    day: '2-digit',
    month: '2-digit',
    year: '2-digit'
  })
}

const scrollToBottom = () => {
  if (messagesContainer.value) {
    messagesContainer.value.scrollTop = messagesContainer.value.scrollHeight
  }
}

const sendMessage = async () => {
  if (!newMessage.value.trim()) return
  
  try {
    await chatStore.sendMessage(props.requestId, newMessage.value.trim())
    newMessage.value = ''
    scrollToBottom()
  } catch (error) {
    console.error('Ошибка при отправке сообщения:', error)
  }
}

// Загрузка сообщений и установка интервала обновления
const setupChat = async () => {
  try {
    await chatStore.fetchMessages(props.requestId)
    scrollToBottom()
    
    // Установка интервала обновления
    updateInterval = setInterval(async () => {
      await chatStore.fetchMessages(props.requestId)
    }, 5000)
  } catch (error) {
    console.error('Ошибка при загрузке сообщений:', error)
  }
}

// Следим за изменением requestId
watch(() => props.requestId, (newId) => {
  if (newId) {
    chatStore.clearMessages()
    setupChat()
  }
})

onMounted(() => {
  setupChat()
})

onUnmounted(() => {
  if (updateInterval) {
    clearInterval(updateInterval)
  }
  chatStore.clearMessages()
})
</script>

<style scoped>
.chat-container {
  display: flex;
  flex-direction: column;
  height: 400px;
  border: 1px solid #ddd;
  border-radius: 8px;
  background-color: #fff;
}

.chat-messages {
  flex: 1;
  overflow-y: auto;
  padding: 1rem;
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.message {
  max-width: 70%;
  padding: 0.75rem;
  border-radius: 8px;
  background-color: #f0f2f5;
  align-self: flex-start;
}

.message-own {
  background-color: #0084ff;
  color: white;
  align-self: flex-end;
}

.message-header {
  display: flex;
  justify-content: space-between;
  margin-bottom: 0.5rem;
  font-size: 0.875rem;
}

.message-own .message-header {
  color: rgba(255, 255, 255, 0.8);
}

.message-author {
  font-weight: bold;
}

.message-time {
  font-size: 0.75rem;
}

.message-content {
  word-break: break-word;
}

.chat-input {
  padding: 1rem;
  border-top: 1px solid #ddd;
}

.input-form {
  display: flex;
  gap: 0.5rem;
}

input {
  flex: 1;
  padding: 0.5rem;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 1rem;
}

button {
  padding: 0.5rem 1rem;
  background-color: #0084ff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 1rem;
}

button:disabled {
  background-color: #ccc;
  cursor: not-allowed;
}

.loading, .error {
  text-align: center;
  padding: 1rem;
  color: #666;
}

.error {
  color: #dc3545;
}
</style>
