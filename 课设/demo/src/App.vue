<script  lang="ts" setup>  
import { ref,provide,onMounted,inject  } from 'vue';  
import { TEST_SYMBOL } from '@/symbol'
import TopBar from './components/TopBar.vue';  
import Home from './components/Home.vue';  
import SideBar from './components/SideBar.vue';  
import Demo from './components/Demo.vue';  
import bus from "./EventBus/eventbus"; // 确保这个 eventbus 适用于你的项目  
import service from './axios-instance'
import { ElNotification } from 'element-plus'
import { useUserStore } from './store/user';
import { useRouter } from 'vue-router' 
const router = useRouter()  
const userStore=useUserStore()
// 在 <script setup> 中，组件默认是局部注册的，所以不需要在 components 对象中声明  
onMounted(()=>{
  const $Test2: ((data) => string) | undefined = inject('$Test')
  $Test2 && $Test2({'da':'da','b':1})

  // createMessageWs()
})
// 提供 color 属性  
provide('color', 'red');

let messageWs = ref(null); 
function createMessageWs(){
  console.log("开启socket链接-----"+'ws://')
  messageWs.value = new WebSocket('ws://' + 'localhost:8080' + '/ws');  
  // messageWs.value.onopen = (event) => {  
  //   // 当 WebSocket 连接打开时，发送认证消息  
  //   authenticate(messageWs);  
  // };  
  messageWs.value.onmessage = (event) => {  
    // 处理从服务器接收到的消息  
    const msg = JSON.parse(event.data);
    console.log(msg)
    }
};
function send(datas){
  // let list=[]
  // list.push({key:passwd-0})
  if(localStorage.getItem('token')&&messageWs.value.readyState == WebSocket.OPEN){
    messageWs.value.send(
    JSON.stringify({
            userTarget: 2,
            message:'message.value',
            token:localStorage.getItem('token')
        }
    ));
  }
  else{
    ElNotification({
        title: '发送异常',
        type: 'error',
      })
  }
}
</script>

<template>
  <div @click="send">Send</div>
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
