<template>
      <div style="width:400px;height:400px;position:absolute;left:550px;top:150px">
      <form  style="color:rgba(220, 228, 253, 0.942);">
        <div class="form-group" >
          <label for="exampleInputEmail1" style="color:rgba(220, 228, 253, 0.942);">请输入账号：</label>
          <div style="display: flex;">
            <input type="text" class="form-control"  id="exampleInputEmail1" placeholder="账号" v-on:input='checkname($event)' v-bind:value='username'>
            <span v-if="ckname" style="color: red;">!!!!!</span>
          </div>
        </div>
        <div class="form-group">
          <label for="exampleInputPassword1" style="color:rgba(220, 228, 253, 0.942);">密码：</label>
          <input type="password" class="form-control" id="exampleInputPassword1" placeholder="密码" v-model="passwd">
        </div>
        <!-- <div class="form-group">
          <label for="exampleInputFile" style="color:rgba(220, 228, 253, 0.942);">File input</label>
          <input type="file" id="exampleInputFile">
          <p class="help-block"  style="color:rgba(220, 228, 253, 0.942);">Example block-level help text here.</p>
        </div> -->
        <div class="checkbox">
          <label>
            <input type="checkbox" style="color:rgba(220, 228, 253, 0.942);"> 记住我
          </label>
        </div>
        <button @click="send" class="btn btn-default">登录</button>
        <div class="form-group" style="margin-top:10px">
          <label class="exampleInputEmail1">
            没有账号？
            <router-link to="/Register" class="xx1">
              点我注册
            </router-link>
          </label>
        </div>
      </form> 
    </div> 

</template>

<script>
import TopBar from './TopBar.vue'
import SideBar from './SideBar.vue'
import Demo from './Demo.vue'
import Login from './Login.vue'
import axios from 'axios';  

import bus from "../EventBus/eventbus";
export default{
  components:{
   TopBar,
    SideBar,
    Demo,
    TopBar,
    SideBar,
    // Index,
    // Login,
  },
  data(){
    return {
      username:'',
      passwd:'',
      ckname:false,
    }
  },
  methods:{
    click()
    {
      axios.get("http://localhost:8080/login")
      .then(response =>{
        console.log("--------")
        console.log(response.data)
        
      }).catch(err=>{
        console.log("----1111----")
        console.error(err)
      })
    },
    checkname(event){
      this.username=event.target.value
      console.log(this.username)
      if(this.username!=""&&(this.username.length<5||this.username.length>20)){
        this.ckname=true
      }
      else{
        this.ckname=false
      }
    },
    send(){
      let list=[]
      list.push({key:this.passwd-0})
      this.ws.send(
      JSON.stringify({
              userTarget: this.username-0,
              userMassages:list
          }
      ));
    }
  },
  created() {
    console.log("开启socket链接-----"+'ws://')
  localStorage.setItem('token','Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTUzNTc2MjUsImlhdCI6MTcxNTI3MTIyNSwiaXNBZG1pbiI6ZmFsc2UsImlzcyI6IiIsInVzZXJJZCI6IjciLCJ1c2VybmFtZSI6Inh4MDAwNSJ9.1gBQEaI79OSpBRs99uZ1QjoHOog-Exkl3x9Z6xnYdTI')
  this.ws = new WebSocket('ws://' + 'localhost:8080' + '/usertoUser/revocation');  
  this.ws.onopen = (event) => {  
    // 当 WebSocket 连接打开时，发送认证消息  
    this.authenticate();  
  };  
  
  this.ws.onmessage = (event) => {  
    // 处理从服务器接收到的消息  
    const msg = JSON.parse(event.data);  
    console.log(msg);  
  };  
  
  // 其他 WebSocket 事件处理...  
  
  // 认证方法  
  this.authenticate = () => {  
    if (this.ws.readyState == WebSocket.OPEN && localStorage.token) {  
      console.log("发送验证信息")
      this.ws.send(  
        JSON.stringify({  
          type: 'auth', // 消息类型，用于区分是普通消息还是认证消息  
          token: localStorage.token,  
          // 其他可能需要的认证信息...  
        })  
      );  
    } 
  }
  }
}
</script>


<style>
.xx1{
  color:rgba(220, 228, 253, 0.942);
}

.xx1:hover{
  color:rgb(98, 98, 211);
}
</style>