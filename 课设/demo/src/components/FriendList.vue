<template>
  <div style="color:rgba(220, 228, 253, 0.942);">
    <div style="display:flex">
      <div class="List">
        <div style="width:10px;height:50px"></div>
        <el-scrollbar style="height:550px;width:200px">
          <p v-for="(item,index) in usertoUsers" 
          :key="index" style="margin-top:10px;line-height:60px;width:200px;height:60px;background-color:rgb(189, 184, 184);color:black;border-radius:12px" class="friend">
            <img :src="item.ToUser.img" style="margin-right:20px; margin-left:10px;width:50px;height:50px;border-radius:50% ;border:rgb(104, 103, 103)" @click="goindex(index)"/>
            {{ item.remarks }}
          </p>
        </el-scrollbar>
      </div>
      <div style="border:1px;height:600px;width:1px;float:left"></div>
      <div>
        <div class="Message" >
          {{ index }}
          <hr>
          <div class="Top" style="width:auto" v-if="index!=-1 && usertoUsers.length!=0">
            <el-scrollbar style="width:607px;height:345px;margin-top:-10px" ref="scrollbarRef" always>
              <div ref="innerRef">
                <p v-for="(message,i) in usertoUsers[index].userMessages" 
                :key="i" 
                 :class="getMessageClass(message.userOwner==usertoUsers[index].userOwner,message.isDeleted)">
                 <div v-if="message.isDeleted" style="display: flex;">
                  {{ message.userOwner==usertoUsers[index].userOwner?'您已撤回一条消息':'对方已撤回一条消息' }}
                  <div v-if="message.standby" class="mb-4">
                    <el-button
                      type="danger"
                      text
                      @click="checkstandby(message.standby)"
                    >
                     重新编辑
                    </el-button>
                  </div>
                  <div @click="messagedelete(index,message.key,i)">
                    <img src='http://localhost:8080/static/images/close.png'/>
                  </div>
                 </div>
                 <div v-else-if="message.userOwner==usertoUsers[index].userOwner" style="display: flex;">
                      <div style="width:300px;height:1px"></div>
                    <div class="bubble">
                        <div class="message">
                          <el-popover :visible="message.visible?message.visible:false" placement="top" :width="160">
                            <div style="text-align: right; margin: 0;">
                              <el-button size="small" text @click="message.visible = false">取消</el-button>
                              <el-button size="small" type="primary" @click="revocation(index,message.key,i)">
                                撤回
                              </el-button>
                              <el-button size="small" type="primary" @click="messagedelete(index,message.key,i)">
                                删除
                              </el-button>
                            </div>
                            <template #reference>
                              <el-button @click="message.visible = true">{{ message.message }}</el-button>
                            </template>
                          </el-popover>
                        
                        </div>
                    </div>
                    <div class="avatar">
                      <img :src="user.img" class="avatar-image" style="margin-right:20px" @click="drawer2 = true"/>
                    </div>
                  </div>
                  <div v-else  style="display: flex;">
                    <div class="avatar">
                      <img :src="usertoUsers[index].ToUser.img" class="avatar-image" style="margin-left:20px" @click="drawer1 = true"/>
                    </div>
                    <div class="bubble">
                      <div class="message">
                        <el-popover :visible="message.visible?message.visible:false" placement="top" :width="160">
                          <div style="text-align: right; margin: 0;">
                            <el-button size="small" text @click="message.visible = false">取消</el-button>
                            <el-button size="small" type="primary" @click="messagedelete(index,message.key,i)">
                              删除
                            </el-button>
                          </div>
                          <template #reference>
                            <el-button @click="message.visible = true">{{ message.message }}</el-button>
                          </template>
                        </el-popover>
                      </div>
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
  <el-drawer v-model="drawer1" direction="rtl" v-if="index!=-1">
    <template #header>
      <h4><img :src="usertoUsers[index].ToUser.img"/> {{ usertoUsers[index].remarks }} </h4>
    </template>
    <template #default>
      <div>账号：{{ usertoUsers[index].ToUser.username }}</div>
      <div>账号名：{{ usertoUsers[index].ToUser.account }}</div>
      <div>备注：{{ usertoUsers[index].remarks }}</div>
      <div>其他：</div>
    </template>
    <template #footer>
      <div style="flex: auto">
        <el-button @click="drawer1 = false">取消</el-button>
        <el-button type="primary" @click="confirmClick">保存</el-button>
      </div>
    </template>
  </el-drawer>
  <el-drawer v-model="drawer2" direction="ltr">
    <template #header>
      <h4><img :src="user.img"/>  </h4>
    </template>
    <template #default>
      <div>账号：{{ user.username }}</div>
      <div>账号名：{{ user.account }}</div>
      <div>其他：</div>
    </template>
    <template #footer>
      <div style="flex: auto">
        <el-button @click="drawer2 = false">取消</el-button>
        <el-button type="primary" @click="confirmClick">保存</el-button>
      </div>
    </template>
  </el-drawer>
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
let drawer1=ref(false)
let drawer2=ref(false)

function send(){
  $Ws && $Ws({
            userTarget: usertoUsers[index.value].userTarget-0,
            message:message.value,
            event:'/usertoUser/send',
            token:localStorage.getItem('token')
        })
  message.value=''
}
function messagedelete(i,key,j){
  usertoUsers[i].userMessages[j].visible=false
  let userMessages=[{key:key-0}]
  service.post('http://localhost:8080/usertoUser/deleteMessage',{
    'userTarget': usertoUsers[i].userTarget-0,
    'userMessages':userMessages,
  }).then(res=>{
    usertoUsers[i].userMessages.splice(j,1)
  }).catch(err=>{
    console.error(err)
    ElNotification({
        title: 'Error',
        message: err.response.data.errorMessage,
        type: 'error',
      })
  })
}
function revocation(i,key,j){
  usertoUsers[i].userMessages[j].visible=false
  let userMessages=[{key:key-0}]
  $Ws && $Ws({
            userTarget: usertoUsers[i].userTarget-0,
            userMessages:userMessages,
            event:'/usertoUser/revocation',
            token:localStorage.getItem('token')
        })
  
}
var user = reactive(JSON.parse(localStorage.getItem('user')))
// 使用watch来监听userStore的userInfo变化  
watch(  
    () => wsStore.Frientmessagecount,  
    (newUserInfo, prevUserInfo) => {  
      if(wsStore.Frientmessagecount){
        wsStore.readFrientMessages().then(res=>{
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
watch(  
    () => wsStore.Frientusercount,  
    (newUserInfo, prevUserInfo) => {  
      if(wsStore.Frientusercount){
        wsStore.readFrientUsers().then(res=>{
          console.log(res)
          res.forEach(element => {
            let p=0;
            for(let i=0;i<usertoUsers.length;i++)
            {
              if(usertoUsers[i].id==element.id)
              {
                usertoUsers.splice(i,1)
                if(index.value>i)
                {
                  break
                }
                else if(index.value==i)
                {
                  p=1;
                  break
                }
                else
                {
                  index.value++
                  break
                }
              }
            }
            service.post('http://localhost:8080/usertoUser/fid',{'userTarget':element.userTarget})
            .then(res=>{
              console.log(res.data)
              element=res.data
              usertoUsers.unshift(element)
              goindex(0)
            })
            
          });
        }).catch(err=>{
          console.error(err)
        })
      }
    },  
    // 可选：配置watch选项，如立即执行、深度监听等  
    { immediate: true, deep: false } // 注意：对于基本类型，通常不需要深度监听（deep: false）  
);
watch(  
    () => wsStore.Frientrevocationcount,  
    (newUserInfo, prevUserInfo) => {  
      if(wsStore.Frientrevocationcount){
        wsStore.readFrientRevocations().then(res=>{

          res.forEach(element => {
            let p=0;
            for(let i=0;i<usertoUsers.length;i++)
            {
              if(usertoUsers[i].id==element.usertoUserId)
              {
                let l=0,r=usertoUsers[i].userMessages.length-1
                while(l<=r)
                {
                  let mid=(l+r)/2
                  if((l+r)%2!=0)
                  {
                    mid=(l+r-1)/2
                  }
                  if(usertoUsers[i].userMessages[mid].key<element.key)
                  {
                    l=mid+1
                  }
                  else if(usertoUsers[i].userMessages[mid].key>element.key)
                  {
                    r=mid-1
                  }
                  else
                  {
                    console.log('撤回了')
                    console.log(mid)
                    if(usertoUsers[i].userMessages[mid].userOwner==usertoUsers[i].userOwner){
                      usertoUsers[i].userMessages[mid].standby=usertoUsers[i].userMessages[mid].message
                    }
                    usertoUsers[i].userMessages[mid].message=''
                    usertoUsers[i].userMessages[mid].isDeleted=true
                    break
                  }
                }
                break
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
const getMessageClass = (isSent,isDeleted) => {
  if(isDeleted){
    return 'message-container-centre'
  }
  return isSent ? 'message-container-right' : 'message-container-left';
};


var usertoUsers=reactive([])

const innerRef = ref<HTMLDivElement>()
const scrollbarRef = ref<InstanceType<typeof ElScrollbar>>()
let index=ref(-1)
function goindex(val){
  index.value=val
  gobottom()
}
function gobottom(){//抵达最底部
  if(usertoUsers.length==0||index.value==-1)
  {
    return
  }
  console.log(usertoUsers)
  nextTick(() => {  
    scrollbarRef.value!.setScrollTop(20000)
  })
}
function getusers(){
  console.log('发送请求')
   service.get('http://localhost:8080/usertoUser/fids')
   .then(res=>{
    console.log(usertoUsers)
    console.log(res.data)
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
function checkstandby(st){
  message.value+=st
}
onMounted(() => {
  wsStore.event=0
  getusers()

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
.message-container-centre{
  float: left;
  margin-left: 40%;
  
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