import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import Router from './router/router.js'
import scroll from 'vue-seamless-scroll'

createApp(App).use(Router).use(scroll).mount('#app')
