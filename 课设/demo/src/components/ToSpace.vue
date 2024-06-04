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
  <el-scrollbar style="width:800px;height:470px;margin-top:-20px" ref="scrollbarRef" always>
      <div ref="innerRef">
        <div style="width:750px;height:auto;min-height:600px;background-color:rgb(189, 184, 184);margin-left:30px;margin-top:30px;border-radius:10px;border-top-right-radius:50px">
         <!-- <div style="box-shadow:-2px 2px 2px rgba(0, 0, 0, 0.15);height:2px;width:700px"></div> -->
          <div v-for="(tmp) in rcd[0]" style="margin-top:10px">
            <div style="box-shadow:-2px 2px 2px rgba(0, 0, 0, 0.15);height:2px;width:750px"></div>
            <div style="height:50px;display:flex;margin-top:20px">
              <img :src="tmp.user.Img" style="margin-left:10px;height:50px;width:50px;border-radius:50%;border:1px double"/>
              <div style="font-size:18px;margin-left:10px">{{tmp.user.Username}}</div>
            </div>
            <div style="display:flex">
              <div style="margin-left:200px;margin-top:-35px;font-size:15px">发表于:<div style="margin-left:20px;font-size:12px">{{tmp.tim}}</div></div>
              <router-link to="/detail" @click="Detail(tmp.trendId)">
                <button style="margin-top:-40px;width:100px;height:30px;margin-left:200px">查看详情</button>
              </router-link>
              </div>
            
            <div style="box-shadow:-2px 2px 2px rgba(0, 0, 0, 0.15);height:2px;width:750px"></div>
            <div style="margin-left:20px;font-size:18px;font-family:隶书;min-height:100px">
              {{tmp.detail}}
            </div>
            <div style="box-shadow:-2px 2px 2px rgba(0, 0, 0, 0.15);height:2px;width:750px"></div>
            <div v-for="comment in tmp.comments">
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
let rcd=reactive([])
let id = ref(localStorage.getItem("id"));
onMounted(()=>{
  service.post("http://localhost:8080/space/fidTrends",{
    "userId":localStorage.getItem("toId")-0
  }) 
  .then(res=>{
    rcd.push(res.data)
    console.log(rcd)
  })
  console.log(rcd)
})
function Detail(e){
  localStorage.setItem("ToId",e-0)

}
</script>
<style>
.to22{
  color:rgb(111, 111, 111);
}
</style>