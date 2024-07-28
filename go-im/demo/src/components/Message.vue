<template>
  <div style="">
    <hr>
    <div style="display:flex;width:800px;height:30px;margin-bottom:-10px;margin-top:-10px;margin-left:70px">
      <router-link to="/space" style="text-decoration: none;">
        <div class="to1" style="margin-top:-7px">说说</div>
      </router-link>
      <router-link to="/newMessage" style="text-decoration: none;">
        <div class="to1" style="margin-top:-7px">留言</div>
      </router-link>
      <router-link to="" style="text-decoration: none;">
        <div class="to1" style="margin-top:-7px">与我相关</div>
      </router-link>
      <router-link to="/photoGraph" style="text-decoration: none;">
        <div class="to1" style="margin-top:-7px">相册</div>
      </router-link>
      <router-link to="" style="text-decoration: none;">
        <div class="to1" style="margin-top:-7px">日志</div>
      </router-link>
    </div>
    <hr>
  <el-scrollbar style="width:800px;height:470px;margin-top:-20px" ref="scrollbarRef" always>
      <div ref="innerRef">
         <!-- <div style="box-shadow:-2px 2px 2px rgba(0, 0, 0, 0.15);height:2px;width:700px"></div> -->
          <div v-for="(tmp) in rcd[0]" style="margin-top:10px">
            <div style="width:750px;height:auto;min-height:200px;background-color:rgb(189, 184, 184);margin-left:30px;margin-top:30px;border-radius:10px;border-top-right-radius:50px">
            <!-- <div style="box-shadow:-2px 2px 2px rgba(0, 0, 0, 0.15);height:2px;width:750px"></div> -->
            <div style="height:50px;display:flex;margin-top:20px">
              <img :src="tmp.user.Img" style="margin-left:10px;height:50px;width:50px;border-radius:50%;border:1px double"/>
              <div style="font-size:18px;margin-left:10px">{{tmp.user.Username}}</div>
            </div>
            <div style="display:flex">
              <div style="margin-left:200px;margin-top:-35px;font-size:15px">留言于:<div style="margin-left:20px;font-size:12px">{{tmp.tim}}</div></div>
              <!-- <router-link to="/detail" @click="Detail(tmp.trendId)">
                <button style="margin-top:-40px;width:100px;height:30px;margin-left:200px">查看详情</button>
              </router-link> -->
              <button style="margin-top:-40px;width:100px;height:30px;margin-left:200px" @click="del(tmp.Id)">删除留言</button>
              </div>
            
            <div style="box-shadow:-2px 2px 2px rgba(0, 0, 0, 0.15);height:2px;width:750px"></div>
            <div style="margin-left:20px;font-size:18px;font-family:隶书;min-height:100px;margin-top:20px">
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
import { ref, onMounted ,h,reactive,nextTick, isRef,inject } from 'vue'; 
import { ElNotification,ElScrollbar } from 'element-plus'
import service from '../axios-instance'
import { useRouter } from 'vue-router';
const $MYGO = inject('$MYGO', '');
let rcd=reactive([])
let id = ref(localStorage.getItem("id"));
 const router = useRouter()
onMounted(()=>{
  service.post($MYGO+'/space/fidMessage',{
    "userId":localStorage.getItem("id")-0
  }) 
  .then(res=>{
    rcd.push(res.data)
    console.log(rcd)
  })
  console.log(rcd)
})
function del(e){
  service.post($MYGO+'/space/delMessage',{
    "messageId":e
  })
  .then(res=>{
    rcd.splice(0,rcd.length)
    service.post($MYGO+'/space/fidMessage',{
    "userId":localStorage.getItem("id")-0
    }) 
    .then(res=>{
      rcd.push(res.data)
      console.log(rcd)
    })
    router.push("/message")
  })
}
</script>
<style>
.to22{
  color:rgb(111, 111, 111);
}
</style>