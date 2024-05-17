import { createRouter, createWebHashHistory } from "vue-router"
import Home from "../components/Home.vue"
import TopBar from "../components/TopBar.vue"
import Index from "../components/Index.vue"
import Login from "../components/Login.vue"
import Register from '../components/Register.vue'
import Person from "../components/Person.vue"
import Space from "../components/Space.vue"
import FriendList from "../components/FriendList.vue"
import Newfriend from "../components/Newfriend.vue"
import Detail from "../components/Detail.vue"
import GroupList from "../components/GroupList.vue"
import Finduser from "../components/Finduser.vue"
const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: "/",
      redirect: "/home",
    },
    {
      path: "/home",
      component: Home,
      meta: { requiresAuth: true }
    },
    {
      path: "/topBar",
      component: TopBar,
      meta: { requiresAuth: true }
    },
    {
      path: "/login",
      component: Login,
    },
    {
      path: "/register",
      component: Register,
    },
    {
      path: "/person",
      component: Person,
      meta: { requiresAuth: true }
    },
    {
      path: "/space",
      component: Space,
      meta: { requiresAuth: true }
    },
    {
      path: "/index",
      component: Index,
      meta: { requiresAuth: true }
    },
    {
      path: "/friendlist",
      component: FriendList,
      meta: { requiresAuth: true }
    },
    {
      path: '/advices',
      component: Home
    },
    {
      path: '/evaluate',
      component: Home
    },
    {
      path: "/newfriend",
      component: Newfriend,
    },
    {
      path: "/grouplist",
      component: GroupList,
    },
    {
      path: "/detail",
      component: Detail,
    },
    {
      path: "/finduser",
      component:Finduser,
    }
  ]
})
router.beforeEach((to, from) => {
  const token = localStorage.getItem('token')
  if (to.meta.requiresAuth && !token) {
    // return {
    //     path:'/login',
    //     query:{redirect:to.fullPath}
    // }
  }
})
export default router