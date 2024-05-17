<template>
  <div style="color:rgba(220, 228, 253, 0.942);">
    <div style="display:flex">
      <div class="List">
        <div style="width:10px;height:30px"></div>
        <el-scrollbar style="height:550px;width:200px">
          <p v-for="(item,index) in groups" 
          :key="index" style="margin-top:10px;line-height:60px;width:200px;height:60px;background-color:rgb(189, 184, 184);color:black;border-radius:12px" class="friend">
            <img src="#" style="margin-right:20px; margin-left:10px;width:50px;height:50px;border-radius:50% ;border:rgb(104, 103, 103)" @click="goindex(index)"/>
            {{ item.groupName }}12312312
          </p>
        </el-scrollbar>
      </div>
      <div style="border:1px;height:600px;width:1px;float:left"></div>
      <div>
        <div class="Message" >
          <div>
            <div>{{ item.groupName }}</div>
            <button style="float:right;margin-right:20px;margin-top:-10px;background-color:rgb(105, 105, 105);border:0px;">···</button>
          </div>
          <hr>
          <div class="Top" style="width:auto" v-if="index!=-1">
            <el-scrollbar style="width:607px;height:345px;margin-top:-10px" ref="scrollbarRef" always>
              <div ref="innerRef">
                <p v-for="(message,i) in groups[index].groupMessages" 
                :key="i" 
                 :class="getMessageClass(message.messageOwner==message.messageSender)">
                  <div v-if="message.messageOwner==message.messageSender" style="display: flex;">
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
          <textarea style="width:607px;height:100px;margin-top:30px;background-color:rgb(141, 141, 141);border:0px" v-model="message"></textarea>
          <button style="color:rgba(220, 228, 253, 0.942);background-color:#82838372;width:60px;height:30px;margin-left:500px" @click="send">发送</button>
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


let messageWs = ref(null); 
let newWs = ref(null); 
let RevocationWs = ref(null); 
let message=ref('')

var groupuser = reactive([{
  "username":"xxx",
  "img":"#"
}])

function send(){
  // let list=[]
  // list.push({key:passwd-0})
  if(messageWs.value.readyState == WebSocket.OPEN){
    messageWs.value.send(
    JSON.stringify({
            groupId: groups[index.value].id-0,
            message:message.value
        }
    ));
  }
  message.value=''
}

function createMessageWs(){
  console.log("开启socket链接-----"+'ws://')
  messageWs.value = new WebSocket('ws://' + 'localhost:8080' + '/group/sendMessage');  
  messageWs.value.onopen = (event) => {  
    // 当 WebSocket 连接打开时，发送认证消息  
    authenticate(messageWs);  
  };  
  messageWs.value.onmessage = (event) => {  
    // 处理从服务器接收到的消息  
    const msg = JSON.parse(event.data);
    if(msg.type==null||msg.type==""){//接受成功  
      console.log(msg);  
      for(let i=0;i<groups.length;i++){
        if(msg.groupId==groups[i].id){
          let IsNo=true
          groups[i].groupMessages.forEach(element => {
            if(element.messageKey==msg.messageKey){
              IsNo=true
            }
          });
          if(IsNo){
            groups[i].groupMessages.push(msg)
          }
          break
        }
      }
      gobottom()
    }
    else{//失败
      ElNotification({
        title: '发送异常',
        message: msg.errorMessage,
        type: 'error',
      })
    }
    
  };
}

function createNewWs(){
  console.log("开启socket链接-----"+'ws://')
  newWs.value = new WebSocket('ws://' + 'localhost:8080' + '/group/createGroup');  
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
  RevocationWs.value = new WebSocket('ws://' + 'localhost:8080' + '/group/revocationMessage');  
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


var groups=reactive([{}])

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
function getgroups(){
  console.log('发送请求')
   service.get('http://localhost:8080/group/fidGroup')
   .then(res=>{
    console.log(res.data)
    groups.pop()
    res.data.forEach(element => {
      groups.push(element)
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
  createMessageWs()
  createNewWs()
  createRevocationWs()
  getgroups()
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
