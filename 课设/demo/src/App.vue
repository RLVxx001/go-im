<script>
// import TopBarVue from './components/TopBar.vue'
// import HelloWorld from './components/HelloWorld.vue'
import TopBar from './components/TopBar.vue'
// import Index from "./components/Index.vue"
import Home from "./components/Home.vue"
import SideBar from './components/SideBar.vue'
import Demo from './components/Demo.vue'
import Login from './components/Login.vue'
export default{
//  name:"App",
  components:{
    // Header,
    // Footer,
    TopBar,
    SideBar,
    Demo,
    // Index,
    Home,
    Login,
  }
}
</script>

<template>
  <router-view></router-view>
   
</template>

<style>

</style>
