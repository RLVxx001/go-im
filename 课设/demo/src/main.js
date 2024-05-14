import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import Router from './router/router.js'
import scroll from 'vue-seamless-scroll'
import ElementPlus from 'element-plus'  
import 'element-plus/dist/index.css'
import { createPinia } from 'pinia'

createApp(App).use(Router).use(scroll).use(ElementPlus).use(createPinia()).mount('#app')
