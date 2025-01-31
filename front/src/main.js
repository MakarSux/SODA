import './assets/main.css'
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router'
import axios from 'axios'
import Toast from 'vue-toastification'
import 'vue-toastification/dist/index.css'
import '@fortawesome/fontawesome-free/css/all.css'

// Настройка базового URL для axios
axios.defaults.baseURL = import.meta.env.VITE_API_URL || 'http://localhost:3000'

const app = createApp(App)

// Использование Pinia для управления состоянием
const pinia = createPinia()
app.use(pinia)
app.use(router)

const options = {
    position: "top-right",
    timeout: 3000,
    closeOnClick: true,
    pauseOnFocusLoss: true,
    pauseOnHover: true,
    draggable: true,
    draggablePercent: 0.6,
    showCloseButtonOnHover: false,
    hideProgressBar: false,
    closeButton: "button",
    icon: true,
    rtl: false
}

app.use(Toast, options)

app.mount('#app')
