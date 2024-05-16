<template>
  <div style="">
    <hr>
    <div style="display:flex;width:800px;height:30px;margin-bottom:-10px;margin-top:-10px;margin-left:70px">
      <router-link to="" style="text-decoration: none;">
        <div style="width:100px;height:50px;text-align:center;color:rgba(220, 228, 253, 0.942);font-size:20px;line-height:30px;margin-left:40px;">发表说说</div>
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
  <el-scrollbar style="width:800px;height:470px;margin-top:-20px" ref="scrollbarRef" always>
      <div ref="innerRef">
        <div v-for="(tmp) in rcd">
          <div style="width:750px;height:auto;min-height:250px;background-color:rgb(189, 184, 184);margin-left:30px;margin-top:30px;border-radius:10px;border-top-right-radius:50px">
            <div style="height:50px;display:flex">
              <img :src="tmp.pht" style="margin-left:10px;height:50px;width:50px;border-radius:50%;border:1px double"/>
              <div style="font-size:23px;margin-left:10px">{{tmp.name}}</div>
            </div>
            <div style="display:flex">
              <div style="margin-left:200px;margin-top:-35px;font-size:15px">发表于:<div style="margin-left:20px;font-size:13px">{{tmp.tim}}</div></div>
              <router-link :to="{path:'/detail',query:{username:tmp.name,pht:tmp.pht,detail:tmp.detail,tim:tmp.tim}}">
                <button style="margin-top:-40px;width:100px;height:30px;margin-left:200px">查看详情</button>
              </router-link>
            </div>
            <div style="margin-top:0px">----------------------------------------------------------------------------------------------------------------------------------------------------------------</div>
            <div style="margin-left:20px;font-size:20px;font-family:隶书">
              {{tmp.detail}}
            </div>
          </div>
        </div>
      </div>    
    </el-scrollbar>
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