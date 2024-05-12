<script setup>  
import { provide } from 'vue';  
import TopBar from './components/TopBar.vue';  
import Home from './components/Home.vue';  
import SideBar from './components/SideBar.vue';  
import Demo from './components/Demo.vue';  
import Login from './components/Login.vue';  
import Register from './components/Register.vue';  
import bus from "./EventBus/eventbus"; // 确保这个 eventbus 适用于你的项目  
  
// 在 <script setup> 中，组件默认是局部注册的，所以不需要在 components 对象中声明  
  
// 提供 color 属性  
provide('color', 'red');
</script>

<template>
  <div style="background-color: #525257a2;width:100%">

    <div style="height:10px;"></div>
    <div>
      <TopBar></TopBar>
    </div>
    <div style="display:flex;margin-top:-15px;">
    <div style="display:flex;margin-left:250px;">
      <SideBar></SideBar>
      <div style="width:810px;height:600px;background-color:#525257a2;margin-left:5px;margin-top:20px;border-radius:18px;border:2px double rgb(138, 137, 137);">
        <router-view></router-view>
      </div>
    </div>    
    </div>
  </div>
</template>

<style>

</style>
