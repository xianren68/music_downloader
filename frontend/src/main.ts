import {createApp} from 'vue'
import App from './App.vue'
import './style.css'
import 'element-plus/theme-chalk/src/message.scss'
import 'element-plus/theme-chalk/src/index.scss'
import router from '@/routes'
import {createPinia} from 'pinia'
import "@/events"

const app = createApp(App)
const pinia = createPinia()
app.use(router)
app.use(pinia)
app.mount('#app')
