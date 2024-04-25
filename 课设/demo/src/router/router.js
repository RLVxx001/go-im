import { createRouter, createWebHashHistory } from "vue-router"
import Home from "../components/Home.vue"
import TopBar from "../components/TopBar.vue"
import Index from "../components/Index.vue"
import Login from "../components/Login.vue"
import Register from '../components/Register.vue'
const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: "/",
      redirect: "/Home",
    },
    {
      path: "/Home",
      component: Home,
    },
    {
      path: "/TopBar",
      component: TopBar,
    },

    {
      path: "/Login",
      component: Login,
    },
    {
      path: "/Register",
      component: Register,
    }
  ]
})

export default router