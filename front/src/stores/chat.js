import { defineStore } from 'pinia'
import axios from 'axios'

export const useChatStore = defineStore('chat', {
  state: () => ({
    messages: [],
    loading: false,
    error: null
  }),

  actions: {
    async fetchMessages(requestId) {
      this.loading = true
      this.error = null
      try {
        const response = await axios.get(`/chat/messages/${requestId}`, {
          headers: {
            Authorization: `Bearer ${localStorage.getItem('token')}`
          }
        })
        this.messages = response.data
      } catch (error) {
        this.error = error.response?.data?.error || 'Ошибка при загрузке сообщений'
        throw error
      } finally {
        this.loading = false
      }
    },

    async sendMessage(requestId, message) {
      this.loading = true
      this.error = null
      try {
        const response = await axios.post('/chat/messages', {
          requestId,
          message
        }, {
          headers: {
            Authorization: `Bearer ${localStorage.getItem('token')}`
          }
        })
        // Добавляем новое сообщение в список
        this.messages.push(response.data)
        return response.data
      } catch (error) {
        this.error = error.response?.data?.error || 'Ошибка при отправке сообщения'
        throw error
      } finally {
        this.loading = false
      }
    },

    clearMessages() {
      this.messages = []
      this.error = null
    }
  }
})
