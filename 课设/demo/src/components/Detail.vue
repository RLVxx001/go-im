<template>
   <el-scrollbar style="width:750px;max-height:550px;background-color:rgb(189, 184, 184);margin-left:30px;margin-top:10px;border-radius:10px;border-top-right-radius:50px">
    <div style="margin-top:20px">  
    <div style="height:50px;display:flex">
      <img :src="rcd[index].img" style="margin-left:10px;height:70px;width:70px;border-radius:50%;border:1px double;"/>      
      <div style="font-size:23px;margin-left:10px">{{rcd[index].name}}</div>     
    </div>      
    <div style="margin-left:200px;margin-top:-35px;font-size:20px">发表于:<div style="margin-left:20px;font-size:13px;font-size:17px;">{{rcd[index].tim}}</div></div>
   </div>
    <div style="margin-top:35px">----------------------------------------------------------------------------------------------------------------------------------------------------------------</div>
  <div style="margin-left:140px;font-size:20px;font-family:隶书;min-height:350px;width:450px">
        {{rcd[index].detail}}
  </div>
  {{rcd[index].is_praise}}
  {{tmp[0]}}
  <img id="show" v-if="rcd[index].is_praise==0" src="../assets/praise.png" style="width:40px;height:40px;margin-left:600px" @click="praise()"/>
  <img id="show" v-else src="../assets/praise.png" style="width:40px;height:40px;margin-left:600px" @click="praise()"/>
  <!-- <div style="height:10px;width:1px"></div> -->
  <div style="margin-top:5px;">----------------------------------------------------------------------------------------------------------------------------------------------------------------</div>
    <div style="">
      <div style="font-size:20px;margin-left:20px;margin-top:10px">评论详情</div>
      <div style="margin-left:40px;font-size:15px;">
        <div v-for="res in rcd[index].commit">
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
onMounted(()=>{
  for(let i=0;i<rcd.length;i++)
  {
    if(rcd[i].id==p)
    {
      index=i;
      return;
    }
  }
})

// var show = document.getElementById("show");
function praise(){
  // axois.post("http://localhost:8080/user/login",{
  //   "username":localStorage.getItem("name"),
  //   "password":localStorage.getItem("password"),
  // }).then(res=>{
    rcd[index].name="yyy";
    rcd[index].is_praise=1-rcd[index].is_praise;
    tmp.value[0].name="xxx";
    console.log(rcd[index].is_praise)
    console.log(tmp.value)
  // })
}
let tmp = ref([{
  "name":"zzz"
}]);
let rcd=reactive([{
  "id":1,
  "name":"xxx",
  "pht":"#",
  "detail":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
  "tim":"2024/5/14 1:11",
  "is_praise":1,
  "commit":[{
    "username":"xxx",
    "disc":"xasfxxxxxxxxxxx",
    "img":"#",
    "tim":"2024/5/14 1:11",
    },
    {
    "username":"xxx",
  "tim":"2024/5/14 1:11",
    "disc":"xasfxxxxxxxxxxxxasfxxxxxxxxxxxxasfxxxxxxxxxxxxasfxxxxxxxxxxxxasfxxxxxxxxxxxxasfxxxxxxxxxxxxasfxxxxxxxxxxxxasfxxxxxxxxxxxxasfxxxxxxxxxxxxasfxxxxxxxxxxxxasfxxxxxxxxxxxxasfxxxxxxxxxxxxasfxxxasfxxxxxxxxxxxxxxxxxxxxxasfxxxxxxxxxxxxasfxxxxxxxxxxx",
    "img":"#"
    },
    {
    "username":"xxx",
  "tim":"2024/5/14 1:11",
    "disc":"xasfxxxxxxxxxxx",
    "img":"#"
    },
    {
    "username":"xxx",
  "tim":"2024/5/14 1:11",
    "disc":"xasfxxxxxxxxxxx",
    "img":"#"
    },
  ]
  },
  {
  "id":2,
  "detail":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
  "tim":"2024/5/14 1:11",
  "commit":[{
    "username":"xxx",
    "disc":"xasfxxxxxxxxxxx",
    "img":"#",
  "tim":"2024/5/14 1:11",
    },
    {
    "username":"xxx",
  "tim":"2024/5/14 1:11",
    "disc":"xasfxxxxxxxxxxx",
    "img":"#"
    },
    {
    "username":"xxx",
  "tim":"2024/5/14 1:11",
    "disc":"xasfxxxxxxxxxxx",
    "img":"#"
    },
    {
    "username":"xxx",
  "tim":"2024/5/14 1:11",
    "disc":"xasfxxxxxxxxxxx",
    "img":"#"
    },
  ]
  },
  {
  "id":3,
  "detail":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
  "tim":"2024/5/14 1:11",
"commit":[{
    "username":"xxx",
    "disc":"xasfxxxxxxxxxxx",
    "img":"#",
  "tim":"2024/5/14 1:11",
    },
    {
    "username":"xxx",
  "tim":"2024/5/14 1:11",
    "disc":"xasfxxxxxxxxxxx",
    "img":"#"
    },
    {
    "username":"xxx",
  "tim":"2024/5/14 1:11",
    "disc":"xasfxxxxxxxxxxx",
    "img":"#"
    },
    {
    "username":"xxx",
  "tim":"2024/5/14 1:11",
    "disc":"xasfxxxxxxxxxxx",
    "img":"#"
    },
  ]
  },
  {
  "id":4,
  "detail":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
  "tim":"2024/5/14 1:11",
  "commit":[{
    "username":"xxx",
    "disc":"xasfxxxxxxxxxxx",
    "img":"#",
  "tim":"2024/5/14 1:11",
    },
    {
    "username":"xxx",
  "tim":"2024/5/14 1:11",
    "disc":"xasfxxxxxxxxxxx",
    "img":"#"
    },
    {
    "username":"xxx",
  "tim":"2024/5/14 1:11",
    "disc":"xasfxxxxxxxxxxx",
    "img":"#"
    },
    {
    "username":"xxx",
  "tim":"2024/5/14 1:11",
    "disc":"xasfxxxxxxxxxxx",
    "img":"#"
    },
  ]
  }
])
let id = ref(localStorage.getItem("id"));
</script>
<style>
.to22{
  color:rgb(111, 111, 111);
}
</style>