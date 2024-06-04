<template>
   <el-scrollbar style="width:750px;max-height:550px;background-color:rgb(189, 184, 184);margin-left:30px;margin-top:10px;border-radius:10px;border-top-right-radius:50px" v-if="rcd&&rcd[0]">
    <div style="margin-top:20px">  
    <div style="height:50px;display:flex">
      <img :src="rcd[0].user.img" style="margin-left:10px;height:70px;width:70px;border-radius:50%;border:1px double;"/>      
      <div style="font-size:23px;margin-left:10px">{{rcd[0].user.username}}</div>     
    </div>      
    <div style="margin-left:200px;margin-top:-35px;font-size:20px">发表于:<div style="margin-left:20px;font-size:13px;font-size:17px;">{{rcd[index].tim}}</div></div>
   </div>
    <div style="margin-top:35px">----------------------------------------------------------------------------------------------------------------------------------------------------------------</div>
  <div style="margin-left:140px;font-size:20px;font-family:隶书;min-height:350px;width:450px">
        {{rcd[0].detail}}
  </div>
  <!-- {{this.$route.query.Id}} -->
  <img id="show" v-if="rcd[0].is_praise==0" src="../assets/praise.png" style="width:40px;height:40px;margin-left:600px" @click="praise()"/>
  <img id="show" v-else src="../assets/praise.png" style="width:40px;height:40px;margin-left:600px" @click="praise()"/>
  <!-- <div style="height:10px;width:1px"></div> -->
  <div style="margin-top:5px;">----------------------------------------------------------------------------------------------------------------------------------------------------------------</div>
    <div style="">
      <div style="font-size:20px;margin-left:20px;margin-top:10px">评论详情</div>
      <div style="margin-left:40px;font-size:15px;">
        <div v-for="res in rcd[0].commit">
          <div style="margin-left:-30px">
          <hr style="margin-left:0px;color:red;background-color:red"/>
          </div>
        <div style="margin-top:20px">
          <div style="display:flex">
            <img :src="res.pht" style="margin-left:10px;height:40px;width:40px;border-radius:50%;border:1px double;"/>      
            <div style="font-size:23px;margin-left:10px;font-size:15px">{{res.username}}</div>   
            <div style=""></div>
          </div>
          <div style="font-size:15px;margin-left:70px;height:auto;width:430px;word-break:break-all">
            {{res.disc}}
          </div>
          
              <div style="margin-left:70px;font-size:13px;margin-top:20px">
                {{res.tim}}
              </div>
        </div>
          <!-- <div style="width:1px;height:20px"></div> -->
        </div>
      </div>
    </div>
  </el-scrollbar>

</template>
<script setup>
import axois from 'axios'
import { ref, onMounted ,h,reactive,nextTick, isRef } from 'vue'; 
import { ElNotification,ElScrollbar } from 'element-plus'
import service from '../axios-instance'
import { useRoute } from 'vue-router' 
let route=useRoute()
var index=0;
var p=route.query.id;
let trend=reactive()
onMounted(()=>{
  service.post("http://localhost:8080/space/fidTrend",{
    "TrendId":localStorage.getItem("ToId")-0
  })
  .then(res=>{
    rcd.push(res.data)
    console.log(rcd[0])
  })
  .catch(err=>{
    console.log(err)  
  })
})

let rcd=reactive([])
let id = ref(localStorage.getItem("id"));
</script>
<style>
.to22{
  color:rgb(111, 111, 111);
}
</style>