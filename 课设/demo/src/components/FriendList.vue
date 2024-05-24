<template>
  <div style="color:rgba(220, 228, 253, 0.942);">
    <div style="display:flex">
      <div class="List">
        <div style="width:10px;height:50px"></div>
        <el-scrollbar style="height:550px;width:200px">
          <p v-for="(item,index) in usertoUsers" 
          :key="index" style="margin-top:10px;line-height:60px;width:200px;height:60px;background-color:rgb(189, 184, 184);color:black;border-radius:12px" class="friend">
            <img src="#" style="margin-right:20px; margin-left:10px;width:50px;height:50px;border-radius:50% ;border:rgb(104, 103, 103)" @click="goindex(index)"/>
            {{ item.remarks }}
          </p>
        </el-scrollbar>
      </div>
      <div style="border:1px;height:600px;width:1px;float:left"></div>
      <div>
        <div class="Message" >
          {{ index }}
          <hr>
          <div class="Top" style="width:auto" v-if="index!=-1">
            <el-scrollbar style="width:607px;height:345px;margin-top:-10px" ref="scrollbarRef" always>
              <div ref="innerRef">
                <p v-for="(message,i) in usertoUsers[index].userMessages" 
                :key="i" 
                 :class="getMessageClass(message.userOwner==usertoUsers[index].userOwner)">
                  <div v-if="message.userOwner==usertoUsers[index].userOwner" style="display: flex;">
                      <div style="width:300px;height:1px"></div>
                    <div class="bubble" style="background-color:rgb(222, 221, 221);margin-right:10px;">
                      <div class="message" v-html="message.message" style=""></div>
                    </div>
                    <div class="avatar">
                      <img src="#" class="avatar-image" style="margin-right:20px" />
                    </div>
                  </div>
                  <div v-else  style="display: flex;">
                    <div class="avatar">
                      <img src="#" class="avatar-image" style="margin-left:20px"/>
                    </div>
                    <div class="bubble" style="background-color:rgb(222, 221, 221);margin-left:10px">
                      <div class="message" v-html="message.message"></div>
                    </div>
                      <div style="width:300px;height:1px"></div>
                  </div>
                </p>
              </div>
              
            </el-scrollbar>
          </div>
        </div>
        <div style="width:1px; background-color: black;"></div>
        <div class="Chat">
          <textarea style="width:607px;height:120px;margin-top:30px;background-color:rgb(141, 141, 141);border:0px" v-model="message"></textarea>
          <button style="color:rgba(220, 228, 253, 0.942);background-color:#82838372;width:60px;height:30px;margin-left:500px" @click="send">发送</button>
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import ImageViewer from "@luohc92/vue3-image-viewer";
import '@luohc92/vue3-image-viewer/dist/style.css';
import { ref, onMounted ,h,reactive,nextTick,inject,watch } from 'vue'; 
import { ElNotification,ElScrollbar } from 'element-plus'
import service from '../axios-instance'
import { useWsStore } from '../store/user';
const wsStore=useWsStore()
const $Ws: ((data) => string) | undefined = inject('$Ws')
let message=ref('')

function send(){
  $Ws && $Ws({
            userTarget: usertoUsers[index.value].userTarget-0,
            message:message.value,
            event:'/send',
            token:localStorage.getItem('token')
        })
  message.value=''
}

// 使用watch来监听userStore的userInfo变化  
watch(  
    () => wsStore.count,  
    (newUserInfo, prevUserInfo) => {  
      if(wsStore.count){
        wsStore.readAndClearMessages().then(res=>{
          console.log(res)
          res.forEach(element => {
            for(let i=0;i<usertoUsers.length;i++)
            {
              if(usertoUsers[i].id==element.usertoUserId){
                usertoUsers[i].userMessages.push(element)
                gobottom()
                break;
              }
            }
          });
        }).catch(err=>{
          console.error(err)
        })
      }
    },  
    // 可选：配置watch选项，如立即执行、深度监听等  
    { immediate: true, deep: false } // 注意：对于基本类型，通常不需要深度监听（deep: false）  
  );

//消息框样式动态选择
const getMessageClass = (isSent) => {
  return isSent ? 'message-container-right' : 'message-container-left';
};


var usertoUsers=reactive([{}])

const innerRef = ref<HTMLDivElement>()
const scrollbarRef = ref<InstanceType<typeof ElScrollbar>>()
let index=ref(0)
function goindex(val){
  console.log(val)
  index.value=val
  gobottom()
}
function gobottom(){//抵达最底部
  nextTick(() => {  
    scrollbarRef.value!.setScrollTop(20000)
  })
}
function getusers(){
  console.log('发送请求')
   service.get('http://localhost:8080/usertoUser/fid')
   .then(res=>{
    console.log(usertoUsers)
    console.log(res.data)
    usertoUsers.pop()
    res.data.forEach(element => {
      usertoUsers.push(element)
    });
    gobottom()
   }).catch(err=>{
      console.error(err)
      let data=err.response.data
      if(data.type=='token'){
        localStorage.removeItem('token')
      }
      ElNotification({
        title: 'Error',
        message: err,
        type: 'error',
      })
   })
}

onMounted(() => {
  getusers()
  goindex(0)
})
</script>
<style scoped>
.scrollbar-demo-item {
  display: flex;
  align-items: center;
  height: 50px;
  margin: 25px;
  text-align: center;
  border-radius: 4px;
  background: var(--el-color-primary-light-9);
  color: var(--el-color-primary);
  width: 100%;
}
</style>
<style>

.avatar {
  margin-left: 10px; /* 修改这里将头像放在消息框的右边 */
}
 
.avatar-image {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  object-fit: cover;
}
 
.bubble {
  color: #000;
  padding: 10px;
  border-radius: 5px;
}
.message {
  text-align: left;
  margin: 0;
}
.message-container {
  display: flex;
  align-items: center;
  margin-bottom: 10px;
}
.message-container-right {
  float:right;
  /*justify-content: flex-end;*/
}
 
.message-container-left {
  float:left;
  /*justify-content: flex-start;*/
}
.List{
  width:310px;
  height:596px;
  border-top-left-radius: 18px;
  border-bottom-left-radius: 18px;
  background-color:#666464;
}

.Message{
  width:607px;
  height:420px;
  border-top-right-radius: 18px;
  background-color: #82838372;
}

.friend:hover{
  background-color:#cdcdcda2;
}

.Chat
{
  height:176px;
  width: 607px;;
  border-bottom-right-radius: 18px;;
  background-color: #8e8f8f;
}

.find{
    height: 32px;
    width: 200px;
    background-color: #525257a2;
  
}

.bTn{
  margin-top:14px;
  color: rgb(207, 230, 244);
  font-size:17px;
  background-color: #2d9cf8;
  text-align: center;
  border-left:1px solid rgba(0, 0, 0, 0.45);
  width:50px;
  height: 34px;
  line-height: 34px;
  border-top-right-radius: 18px;
  border-bottom-right-radius: 18px
}

.inputT{
    margin-left: 210px;
    margin-top:15px;
  width:150px;
  height: 32px;
  text-indent: 1em;
   line-height: 34px;
  border: 1px solid black;
  border-top-left-radius: 17px;
  border-bottom-left-radius: 17px;
  font-size: 12px;
}
</style>