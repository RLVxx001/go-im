<template>
  <div style="color:rgba(220, 228, 253, 0.942);">
    <div style="display:flex">
      <div class="List">
        <el-scrollbar height="400px">
          <p v-for="(item,index) in usertoUsers" 
          :key="index" class="scrollbar-demo-item" >
            <img src="#" style="margin-right:20px; margin-left:10px; " @click="goindex(index)"/>
            {{ item.remarks }}{{ item.userTarget }}
          </p>
        </el-scrollbar>
      </div>
      <div>
        <div class="Message" >
          {{ index }}
          <div class="Top" v-if="index!=-1">
            <el-scrollbar height="400px"  ref="scrollbarRef" always>
              <div ref="innerRef">
                <p v-for="(message,i) in usertoUsers[index].userMessages" 
                :key="i" class="scrollbar-demo-item"
                 :class="getMessageClass(message.isSent)" @scroll="scroll">
                  <div v-if="message.isSent" style="display: flex;">
                    <div class="bubble">
                      <div class="message" v-html="message.message"></div>
                    </div>
                    <div class="avatar">
                      <img src="#" class="avatar-image"/>
                    </div>
                  </div>
                  <div v-else  style="display: flex;">
                    <div class="avatar">
                      <img src="#" class="avatar-image"/>
                    </div>
                    <div class="bubble">
                      <div class="message" v-html="message.message"></div>
                    </div>
                  </div>
                </p>
              </div>
              
            </el-scrollbar>
          </div>
        </div>
        <div style="width:1px; background-color: black;"></div>
        <div class="Chat">

        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import ImageViewer from "@luohc92/vue3-image-viewer";
import '@luohc92/vue3-image-viewer/dist/style.css';
import { ref, onMounted ,h,reactive,nextTick } from 'vue'; 
import { ElNotification,ElScrollbar } from 'element-plus'
import service from '../axios-instance'


const messageWs = ref(null); 
const newWs = ref(null); 
const RevocationWs = ref(null); 

function send(){
  // let list=[]
  // list.push({key:passwd-0})
  // if(messageWs.value.readyState == WebSocket.OPEN){
  //   messageWs.value.send(
  //   JSON.stringify({
  //           userTarget: username-0,
  //           userMassages:list
  //       }
  //   ));
  // }
}

function createMessageWs(){
  console.log("开启socket链接-----"+'ws://')
  messageWs.value = new WebSocket('ws://' + 'localhost:8080' + '/usertoUser/send');  
  messageWs.value.onopen = (event) => {  
    // 当 WebSocket 连接打开时，发送认证消息  
    authenticate(messageWs);  
  };  
  messageWs.value.onmessage = (event) => {  
    // 处理从服务器接收到的消息  
    const msg = JSON.parse(event.data);  
    console.log(msg);  
  };
}

function createNewWs(){
  console.log("开启socket链接-----"+'ws://')
  newWs.value = new WebSocket('ws://' + 'localhost:8080' + '/usertoUser');  
  newWs.value.onopen = (event) => {  
    // 当 WebSocket 连接打开时，发送认证消息  
    authenticate(newWs);  
  };  
  
  newWs.value.onmessage = (event) => {  
    // 处理从服务器接收到的消息  
    const msg = JSON.parse(event.data);  
    console.log(msg);  
  };
}


function createRevocationWs(){
  console.log("开启socket链接-----"+'ws://')
  RevocationWs.value = new WebSocket('ws://' + 'localhost:8080' + '/usertoUser/revocation');  
  RevocationWs.value.onopen = (event) => {  
    // 当 WebSocket 连接打开时，发送认证消息  
    authenticate(RevocationWs);  
  };  
  
  RevocationWs.value.onmessage = (event) => {  
    // 处理从服务器接收到的消息  
    const msg = JSON.parse(event.data);  
    console.log(msg);  
  };
}

function authenticate(ws){// 认证方法 
  if (ws.value.readyState == WebSocket.OPEN && localStorage.token) {  
    console.log("发送验证信息")
    ws.value.send(  
      JSON.stringify({  
        type: 'auth', // 消息类型，用于区分是普通消息还是认证消息  
        token: localStorage.token,  
        // 其他可能需要的认证信息...  
      })  
    );  
  }
}



//消息框样式动态选择
const getMessageClass = (isSent) => {
  return isSent ? 'message-container-right' : 'message-container-left';
};


let usertoUsers=reactive([
  {remarks:'1',userTarget:'2',userMessages:[
    {message:'您的选择错误啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊'},{message:'11111',isSent:true},{message:'11111',isSent:true},{message:'11111'},{message:'11111'},
    {message:'11111'},{message:'11111'},{message:'11111'},{message:'11111'},{message:'11111'},
    {message:'11111'},{message:'11111'},{message:'11111'},{message:'11111'},{message:'11111'},
    {message:'11111'},{message:'11111'},{message:'11111'},{message:'11111'},{message:'11111'},
    {message:'11111'},{message:'11111'},{message:'11111'},{message:'11111'},{message:'11111'},
    {message:'11111'},{message:'11111'},{message:'11111'},{message:'11111'},{message:'11111'},
    {message:'11111'},{message:'11111'},{message:'11111'},{message:'11111'},{message:'11111'},
    {message:'11111'},{message:'11111'},{message:'11111'},{message:'您的选择错误啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊',isSent:true},{message:'您的选择错误啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊'},
  ]},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
  {remarks:'1',userTarget:'2'},
])

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
    scrollbarRef.value!.setScrollTop(innerRef.value!.clientHeight - 350)
  })
}
function getusers(){
  console.log('发送请求')
   service.get('http://localhost:8080/usertoUser/fid')
   .then(res=>{
    console.log(res.data)
    usertoUsers=res.data
    console.log(usertoUsers)
    index.value=0
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
  // createMessageWs()
  // createNewWs()
  // createRevocationWs()
  // getusers()
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
  width: 50%;
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
  background-color: #e8e8e8;
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
  justify-content: flex-end;
}
 
.message-container-left {
  justify-content: flex-start;
}
.List{
  width:100px;
  height:596px;
  border-top-left-radius: 18px;
  border-bottom-left-radius: 18px;
  background-color: #b9b8b8a2;
}

.Message{
  width:607px;
  height:420px;
  border-top-right-radius: 18px;
  background-color: #a7a7aca2;
}



.Chat
{
  height:176px;
  background-color: #c8c8c9a2;
  width: 607px;;
  border-bottom-right-radius: 18px;;
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