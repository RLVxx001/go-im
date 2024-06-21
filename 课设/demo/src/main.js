import { createApp,ref } from 'vue'
import './style.css'
import App from './App.vue'
import Router from './router/router.js'
import scroll from 'vue-seamless-scroll'
import ElementPlus from 'element-plus'  
import 'element-plus/dist/index.css'
import { createPinia } from 'pinia'
import { MYGO } from './symbol'
import {ElNotification} from 'element-plus'  
import { useWsStore } from './store/user';

let mygo='back-end'

let Ws = ref(null); 

console.log("开启socket链接-----"+'ws://')//'ws://' + mygo +':18080' + '/ws'
Ws.value = new WebSocket('ws://10.200.7.84:60080/ws');  
// messageWs.value.onopen = (event) => {  
//   // 当 WebSocket 连接打开时，发送认证消息  
//   authenticate(messageWs);  
// };  


const send=(data)=>{
  // let list=[]
  // list.push({key:passwd-0})
  console.log(data)
  if(Ws.value&&localStorage.getItem('token')&&Ws.value.readyState == WebSocket.OPEN){
    Ws.value.send(
    JSON.stringify(data));
  }
  else{
    ElNotification({
        title: '发送异常',
        message:'请求失败',
        type: 'error',
      })
  }
}

const app = createApp(App)
app.use(Router)
app.use(scroll)
app.use(ElementPlus)
app.use(createPinia())
// 使用 symbol 方式
// app.provide(TEST_SYMBOL, send)
// 使用自定义字符串方式
app.provide('$Ws', send)
app.provide('$MYGO', mygo)
app.mount('#app')
const wsStore=useWsStore()

Ws.value.onmessage = (event) => {  
    // 处理从服务器接收到的消息  
    const msg = JSON.parse(event.data);
    if(msg.type=='token'){
      localStorage.removeItem('token')
      return
    }
    else if(msg.type=='err'){
      ElNotification({
        title: '异常',
        message:msg.errorMessage,
        type: 'error',
      })
      return
    }
    console.log(msg)
    wsStore.addMessage(msg.data,msg.event)
}
Ws.value.onopen = (event) => {  
    if(localStorage.getItem('token'))
    {

        Ws.value.send(JSON.stringify({
            token:localStorage.getItem('token'),
            event:''
        }))
    }
};  
