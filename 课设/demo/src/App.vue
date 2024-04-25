<script>
// import TopBarVue from './components/TopBar.vue'
// import HelloWorld from './components/HelloWorld.vue'
import TopBar from './components/TopBar.vue'
// import Index from "./components/Index.vue"
import Home from "./components/Home.vue"
import SideBar from './components/SideBar.vue'
import Demo from './components/Demo.vue'
import Login from './components/Login.vue'
import Register from './components/Register.vue'
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
    Register,
  }
}
</script>

<template>
  <div style="background-color: #525257a2;width:100%">
  <TopBar></TopBar>
  <div style="display:flex">
    <SideBar></SideBar>
      <router-view></router-view>
  </div>    
  </div>
</template>

<style>

</style>
