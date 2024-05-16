<template>
   <div style="width:750px;height:auto;min-height:550px;background-color:rgb(189, 184, 184);margin-left:30px;margin-top:10px;border-radius:10px;border-top-right-radius:50px">
    <div style="margin-top:20px">  
    <div style="height:50px;display:flex">
      <img :src="this.$route.query.pht" style="margin-left:10px;height:70px;width:70px;border-radius:50%;border:1px double;"/>      
      <div style="font-size:23px;margin-left:10px">{{this.$route.query.username}}</div>     
    </div>      
    <div style="margin-left:200px;margin-top:-35px;font-size:15px">发表于:<div style="margin-left:20px;font-size:13px">{{this.$route.query.tim}}</div></div>
   </div>
    <div style="margin-top:15px">----------------------------------------------------------------------------------------------------------------------------------------------------------------</div>
  <div style="margin-left:20px;font-size:20px;font-family:隶书">
              {{this.$route.query.detail}}
            </div>
  </div>

</template>
<script setup>
import { ref, onMounted ,h,reactive,nextTick, isRef } from 'vue'; 
import { ElNotification,ElScrollbar } from 'element-plus'
import service from '../axios-instance'
let rcd=reactive([{
  "name":"xxx",
  "pht":"#",
  "detail":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
  "tim":"2024/5/14 1:11",
  "commit":[{
    "username":"xxx",
    "disc":"xasfxxxxxxxxxxx",
    "img":"#"
  }]
  },
  {
  "detail":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
  "tim":"2024/5/14 1:11",
  "commit":[{
    "username":"xxx",
    "disc":"xasfxxxxxxxxxxx",
    "img":"#"
  }]
  },
  {
  "detail":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
  "tim":"2024/5/14 1:11",
  "commit":[{
    "username":"xxx",
    "disc":"xasfxxxxxxxxxxx",
    "img":"#"
  }]
  },
  {
  "detail":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
  "tim":"2024/5/14 1:11",
  "commit":[{
    "username":"xxx",
    "disc":"xasfxxxxxxxxxxx",
    "img":"#"
  }]
  }
])
let id = ref(localStorage.getItem("id"));
</script>
<style>
.to22{
  color:rgb(111, 111, 111);
}
</style>