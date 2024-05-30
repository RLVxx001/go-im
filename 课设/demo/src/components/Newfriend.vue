<template>
  <div style="">
      <div style="width:100px;height:50px;text-align:center;color:rgba(220, 228, 253, 0.942);font-size:20px;line-height:30px;margin-left:40px;margin-top:10px;margin-bottom:-20px">好友通知：</div>
    <hr>
  <el-scrollbar style="width:800px;height:470px;margin-top:-20px" ref="scrollbarRef" always>
      <div ref="innerRef">
        <div v-for="(tmp,index) in friendlist" :key="index">
          <div style="width:750px;height:auto;min-height:70px;background-color:rgb(189, 184, 184);margin-left:30px;margin-top:10px;border-radius:10px">
            <div style="height:50px;display:flex">
              <img :src="tmp.userResponse.img?tmp.userResponse.img:tmp.groupResponse.img" style="margin-left:10px;height:50px;width:50px;border-radius:50%;border:1px double"/>
              <div style="font-size:23px;margin-left:10px">
                {{ tmp.userResponse.account?tmp.userResponse.account:tmp.groupResponse.groupName }}
              </div>
              <div style="font-size:23px;margin-left:10px">
                <span v-if="tmp.class==0">{{ tmp.userOwner==id?'请求添加对方为好友':tmp.remarks }}</span>
                <span v-if="tmp.class==1">{{ tmp.userOwner==id?'请求加入群聊':'申请加入'+tmp.groupResponse.groupName }}</span>
                <span v-if="tmp.class==2">{{ tmp.target==id?tmp.userResponse.account+'邀请你加入':'邀请 '+tmp.target+' 加入' }}</span>
              </div>
            </div>
            <div style="display:flex">
              <div style="margin-left:160px;margin-top:-40px;font-size:15px"></div>
              <div v-if="tmp.stats==0" style="background-color:rgb(223, 219, 219);border:1px double;text-align:center;line-height:30px;margin-top:-30px;width:100px;height:30px;margin-left:300px">同意</div>
              <div v-if="tmp.stats==1" style="text-align:center;line-height:30px;margin-top:-30px;width:100px;height:30px;margin-left:300px">已拒绝</div>
              <div v-if="tmp.stats==2" style="text-align:center;line-height:30px;margin-top:-30px;width:100px;height:30px;margin-left:300px">已同意</div>
              <!-- <div :v-if="tmp.is_agr==1" style="margin-top:-30px;width:100px;height:30px;margin-left:300px">已通过</div> -->
            </div>
            <div style="margin-left:100px;font-size:15px">留言：{{tmp.msg}}</div>
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
let friendlist=reactive([])
let id = ref(localStorage.getItem("id"));
function getlist(){
  service.get('http://localhost:8080/userApplication/fids')
  .then(res=>{
    res.data.forEach(element => {
      friendlist.push(element)
    });
    console.log(friendlist)
  }).catch(err=>{

  })
}
onMounted(()=>{
  getlist()
})
</script>
<style>
.to22{
  color:rgb(223, 219, 219);
}
</style>