import {createApp} from 'vue'
import App from './App.vue'
import './style.css'
import 'element-plus/theme-chalk/src/message.scss'
import 'element-plus/theme-chalk/src/index.scss'
import router from '@/routes'

const app = createApp(App)
app.use(router)
app.mount('#app')
