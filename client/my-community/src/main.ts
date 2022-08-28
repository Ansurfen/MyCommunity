import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import elementplus from 'element-plus'
import 'element-plus/dist/index.css'
import { createPinia } from 'pinia'
import animated from 'animate.css'

const app = createApp(App)
app.use(createPinia()).use(router).use(elementplus).use(animated).mount('#app')
