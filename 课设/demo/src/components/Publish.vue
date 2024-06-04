<template>
  <div style="">
    <hr>
    <div style="display:flex;width:800px;height:30px;margin-bottom:-10px;margin-top:-10px;margin-left:70px">
      <router-link to="publish" style="text-decoration: none;">
        <div style="width:100px;height:50px;text-align:center;color:rgba(220, 228, 253, 0.942);font-size:20px;line-height:30px;margin-left:40px;">发表</div>
      </router-link>
      <router-link to="" style="text-decoration: none;">
        <div style="width:80px;height:50px;text-align:center;color:rgba(220, 228, 253, 0.942);font-size:20px;line-height:30px;margin-left:40px;">留言板</div>
      </router-link>
      <router-link to="" style="text-decoration: none;">
        <div style="width:80px;height:50px;text-align:center;color:rgba(220, 228, 253, 0.942);font-size:20px;line-height:30px;margin-left:40px;">与我相关</div>
      </router-link>
      <router-link to="" style="text-decoration: none;">
        <div style="width:80px;height:50px;text-align:center;color:rgba(220, 228, 253, 0.942);font-size:20px;line-height:30px;margin-left:40px;">相册</div>
      </router-link>
      <router-link to="" style="text-decoration: none;">
        <div style="width:80px;height:50px;text-align:center;color:rgba(220, 228, 253, 0.942);font-size:20px;line-height:30px;margin-left:40px;">日志</div>
      </router-link>
    </div>
    <hr>
  <div style="width:600px;height:480px;margin-left:100px;background-color:rgb(189, 184, 184);border-radius:10px;">
    <div style="height:1px;width:1px;"></div>
    <div style="font-size:20px;margin-left:250px;margin-top:20px">发表</div>
    <div style="margin-left:0px;margin-top:10px;;width:600px;height:2px;"></div>
    <textarea v-model="rcd" style="margin-top:20px;width:500px;margin-left:50px;height:300px;background-color:rgb(209, 209, 207)"></textarea>
    <button type="button" class="btn btn-primary" style="margin-left:270px;margin-top:10px" @click="publish()">发表</button>
  </div>
  </div>
</template>
<script setup>
import { ref, onMounted ,h,reactive,nextTick, isRef } from 'vue'; 
import { ElNotification,ElScrollbar } from 'element-plus'
import service from '../axios-instance'

import { useRouter } from 'vue-router' 
let rcd=reactive()
let id = ref(localStorage.getItem("id"));
let praise = reactive(0)
onMounted(()=>{
})
function publish(){
  service.post("http://localhost:8080/space/addTrends",{
    "userId":localStorage.getItem("id")-0,
    "detail":rcd,
    "praise":praise,
    "comments":[{
        "spaceId":localStorage.getItem("id")-0+4
    }
    ],
    "spaceId":localStorage.getItem("id")-0+4
  })
  .then(res=>{
    alert("发表成功")
    router.push('/person')
  })
}
</script>
<style>
.to22{
  color:rgb(111, 111, 111);
}
</style>5