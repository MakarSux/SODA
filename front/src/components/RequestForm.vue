<template>
  <div class="request-form">
    <h2>Оставить заявку</h2>
    <form @submit.prevent="handleSubmit">
      <div class="form-group">
        <label for="name">Имя*</label>
        <input 
          type="text" 
          id="name" 
          v-model="formData.name" 
          required 
          placeholder="Введите ваше имя"
        >
      </div>

      <div class="form-group">
        <label for="contact">Контактная информация*</label>
        <input 
          type="text" 
          id="contact" 
          v-model="formData.contact" 
          required 
          placeholder="Email или телефон"
        >
      </div>

      <div class="form-group">
        <label for="service">Тип услуги*</label>
        <select id="service" v-model="formData.serviceType" required>
          <option value="">Выберите тип услуги</option>
          <option value="landing">Лендинг</option>
          <option value="shop">Интернет-магазин</option>
          <option value="corporate">Корпоративный сайт</option>
        </select>
      </div>

      <div 
        class="drop-zone"
        @drop.prevent="handleDrop"
        @dragover.prevent="handleDragOver"
        @dragleave.prevent="handleDragLeave"
        :class="{ 'drop-zone--active': isDragging }"
      >
        <div class="drop-zone__content">
          <i class="fas fa-cloud-upload-alt"></i>
          <p>Перетащите файлы сюда или кликните для выбора</p>
          <input 
            type="file" 
            multiple 
            @change="handleFileSelect" 
            class="drop-zone__input"
          >
        </div>
        <div class="drop-zone__files" v-if="files.length">
          <div v-for="(file, index) in files" :key="index" class="drop-zone__file">
            <span>{{ file.name }}</span>
            <button @click="removeFile(index)" type="button">&times;</button>
          </div>
        </div>
      </div>

      <div class="form-group">
        <label for="comments">Дополнительные пожелания</label>
        <textarea 
          id="comments" 
          v-model="formData.comments" 
          placeholder="Опишите ваши пожелания"
        ></textarea>
      </div>

      <button type="submit" class="submit-btn" :disabled="isSubmitting">
        {{ isSubmitting ? 'Отправка...' : 'Отправить заявку' }}
      </button>
    </form>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useToast } from 'vue-toastification'
import axios from 'axios'

const API_BASE_URL = 'http://localhost:3000'
const toast = useToast()
const MAX_FILE_SIZE = 5 * 1024 * 1024 // 5MB
const ALLOWED_TYPES = ['image/jpeg', 'image/png', 'application/pdf', 'application/msword', 'application/vnd.openxmlformats-officedocument.wordprocessingml.document']

const formData = reactive({
  name: '',
  contact: '',
  serviceType: '',
  comments: ''
})

const files = ref([])
const isDragging = ref(false)
const isSubmitting = ref(false)

const validateFile = (file) => {
  if (file.size > MAX_FILE_SIZE) {
    toast.error(`Файл ${file.name} слишком большой. Максимальный размер: 5MB`)
    return false
  }
  if (!ALLOWED_TYPES.includes(file.type)) {
    toast.error(`Файл ${file.name} имеет неподдерживаемый формат`)
    return false
  }
  return true
}

const handleDragOver = () => {
  isDragging.value = true
}

const handleDragLeave = () => {
  isDragging.value = false
}

const handleDrop = (e) => {
  isDragging.value = false
  const droppedFiles = [...e.dataTransfer.files]
  const validFiles = droppedFiles.filter(validateFile)
  files.value = [...files.value, ...validFiles]
}

const handleFileSelect = (e) => {
  const selectedFiles = [...e.target.files]
  const validFiles = selectedFiles.filter(validateFile)
  files.value = [...files.value, ...validFiles]
}

const removeFile = (index) => {
  files.value = files.value.filter((_, i) => i !== index)
}

const validateEmail = (email) => {
  const re = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
  return re.test(email)
}

const validatePhone = (phone) => {
  const re = /^\+?[1-9]\d{10,14}$/
  return re.test(phone)
}

const validateForm = () => {
  if (!formData.name.trim()) {
    toast.error('Пожалуйста, введите ваше имя')
    return false
  }
  
  const contact = formData.contact.trim()
  if (!contact) {
    toast.error('Пожалуйста, введите контактную информацию')
    return false
  }
  
  if (!validateEmail(contact) && !validatePhone(contact)) {
    toast.error('Пожалуйста, введите корректный email или телефон')
    return false
  }
  
  if (!formData.serviceType) {
    toast.error('Пожалуйста, выберите тип услуги')
    return false
  }
  return true
}

const handleSubmit = async () => {
  if (!validateForm()) return

  try {
    isSubmitting.value = true
    const formDataToSend = new FormData()
    
    Object.keys(formData).forEach(key => {
      formDataToSend.append(key, formData[key])
    })
    
    files.value.forEach(file => {
      formDataToSend.append('files[]', file)
    })

    // Получаем токен из localStorage (предполагая, что он там есть после авторизации)
    const token = localStorage.getItem('token')
    
    const response = await axios.post(`${API_BASE_URL}/requests/submit`, formDataToSend, {
      headers: {
        'Content-Type': 'multipart/form-data',
        'Authorization': token ? `Bearer ${token}` : undefined
      }
    })
    
    if (response.data.success) {
      toast.success('Заявка успешно отправлена!')
      
      // Очистка формы
      Object.keys(formData).forEach(key => {
        formData[key] = ''
      })
      files.value = []
    } else {
      throw new Error(response.data.message || 'Ошибка при отправке заявки')
    }
    
  } catch (error) {
    if (error.response?.status === 401) {
      toast.error('Необходимо авторизоваться для отправки заявки')
      // Здесь можно добавить перенаправление на страницу входа
      // router.push('/login')
    } else {
      toast.error(error.response?.data?.message || 'Произошла ошибка при отправке заявки')
    }
    console.error(error)
  } finally {
    isSubmitting.value = false
  }
}
</script>

<style scoped lang="scss">
.request-form {
  max-width: 600px;
  margin: 0 auto;
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

  input,
  select,
  textarea {
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

  textarea {
    min-height: 100px;
    resize: vertical;
  }
}

.drop-zone {
  border: 2px dashed #ddd;
  border-radius: 4px;
  padding: 2rem;
  text-align: center;
  transition: all 0.3s ease;
  margin-bottom: 1.5rem;
  cursor: pointer;

  &--active {
    border-color: #007bff;
    background: rgba(0, 123, 255, 0.1);
  }

  &__content {
    i {
      font-size: 2rem;
      color: #666;
      margin-bottom: 1rem;
    }

    p {
      margin: 0;
      color: #666;
    }
  }

  &__input {
    display: none;
  }

  &__files {
    margin-top: 1rem;
    text-align: left;
  }

  &__file {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0.5rem;
    background: #f8f9fa;
    border-radius: 4px;
    margin-bottom: 0.5rem;

    button {
      background: none;
      border: none;
      color: #dc3545;
      cursor: pointer;
      font-size: 1.2rem;
      padding: 0 0.5rem;

      &:hover {
        color: rgba(220, 53, 70, 0.9);
      }
    }
  }
}

.submit-btn {
  width: 100%;
  padding: 1rem;
  background: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  font-size: 1rem;
  cursor: pointer;
  transition: background-color 0.3s ease;

  &:hover {
    background: #007bff;
  }

  &:disabled {
    background: #ccc;
    cursor: not-allowed;
  }
}

@media (max-width: 768px) {
  .request-form {
    padding: 1rem;
  }
}
</style>
