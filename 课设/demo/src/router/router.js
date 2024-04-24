import { createRouter, createWebHashHistory } from "vue-router"
import Home from "../components/Home.vue"
import TopBar from "../components/TopBar.vue"
import Index from "../components/Index.vue"
import Login from "../components/Login.vue"

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: "/",
      redirect: "/Login",
    },
    {
      path: "/Home",
      component: Home,
      children: [
        {
          path: "/Index",
          component: Index,
          },
        
      ]
    },
    {
      path: "/TopBar",
      component: TopBar,
    },
    
    {
      path: "/Login",
      component: Login,
    }
  ]
})

export default router