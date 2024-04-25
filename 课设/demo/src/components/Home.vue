<template>
  <!-- <TopBar></TopBar> -->
  <div style="background-color: #525257a2;width:100%">
  <TopBar></TopBar>
  <div style="display:flex">
    <SideBar></SideBar>
      <router-view></router-view>
  </div>    
  </div>
</template>
<script>
import TopBar from './TopBar.vue'
import SideBar from './SideBar.vue'
import Demo from './Demo.vue'
import Login from './Login.vue'
export default{
  components:{
   TopBar,
    SideBar,
    Demo,
    // Index,
    Login,
  }
}
</script>
