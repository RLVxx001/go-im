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
import Publish from "../components/Publish.vue"
import ToSpace from "../components/ToSpace.vue"
import Message from "../components/Message.vue"
import NewMessage from "../components/NewMessage.vue"
import ToMessage from "../components/ToMessage.vue"
import ToNewMessage from "../components/ToNewMessage.vue"
import PhotoGraph from "../components/PhotoGraph.vue"
import ToPhotoGraph from "../components/ToPhotoGraph.vue"
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
      component: Home,
      meta: { requiresAuth: true }
    },
    {
      path: '/evaluate',
      component: Home,
      meta: { requiresAuth: true }
    },
    {
      path: "/newfriend",
      component: Newfriend,
      meta: { requiresAuth: true }
    },
    {
      path: "/grouplist",
      component: GroupList,
      meta: { requiresAuth: true }
    },
    {
      path: "/detail",
      component: Detail,
      meta: { requiresAuth: true }
    },
    {
      path: "/finduser",
      component: Finduser,
      meta: { requiresAuth: true }
    },
    {
      path: "/publish",
      component: Publish
    },
    {
      path: "/toSpace",
      component: ToSpace,
    },
    {
      path: "/message",
      component: Message,
    },
    {
      path: "/newMessage",
      component: NewMessage
    },
    {
      path: "/toMessage",
      component: ToMessage,
    },
    {
      path: "/toNewMessage",
      component: ToNewMessage,
    },
    {
      path: "/photoGraph",
      component: PhotoGraph
    },
    {
      path: "/toPhotoGraph",
      component: ToPhotoGraph,
    }
  ]
})
router.beforeEach((to, from) => {
  const token = localStorage.getItem('token')
  if (to.meta.requiresAuth && !token) {
    return {
      path: '/login',
      query: { redirect: to.fullPath }
    }
  }
})
export default router