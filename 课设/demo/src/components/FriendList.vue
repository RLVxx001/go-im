<template>
  <div style="color:rgba(220, 228, 253, 0.942);">
    <div style="display:flex">
      <div class="List">
        <!-- <div class="find">
          <input  class="inputT" placeholder='内容'/>
          <button class="btn">搜索</button>
        </div> -->
        <div style="height:25px;font-size:17px;">
          <img src="#" style="margin-right:20px;margin-left:10px;"/>
            sdsdsds
        </div>
        <div style="height:25px;font-size:17px;">
          <img src="#" style="margin-right:20px;margin-left:10px;"/>
            sdsdsds
        </div>
      </div>
      <div>
        <div class="Message">
          <div class="Top">
          sd
          </div>
        </div>
        <div style="width:1px;backgroud-color:black"></div>
        <div class="Chat">

        </div>
      </div>
    </div>
  </div>
</template>
<script>
  export default{
    data(){
      return{
        UserList:[],
      }
    },
    methods:{
      send(){
        let list=[]
        list.push({key:this.passwd-0})
        if(this.ws.readyState == WebSocket.OPEN){
          this.ws.send(
          JSON.stringify({
                  userTarget: this.username-0,
                  userMassages:list
              }
          ));
        }
      },
      createMessageWs(){
        console.log("开启socket链接-----"+'ws://')
        localStorage.setItem('token','Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTUzNTc2MjUsImlhdCI6MTcxNTI3MTIyNSwiaXNBZG1pbiI6ZmFsc2UsImlzcyI6IiIsInVzZXJJZCI6IjciLCJ1c2VybmFtZSI6Inh4MDAwNSJ9.1gBQEaI79OSpBRs99uZ1QjoHOog-Exkl3x9Z6xnYdTI')
        this.messageWs = new WebSocket('ws://' + 'localhost:8080' + '/usertoUser/revocation');  
        this.messageWs.onopen = (event) => {  
          // 当 WebSocket 连接打开时，发送认证消息  
          this.authenticate(this.messageWs);  
        };  
        this.messageWs.onmessage = (event) => {  
          // 处理从服务器接收到的消息  
          const msg = JSON.parse(event.data);  
          console.log(msg);  
        };
      },
      createNewWs(){
        console.log("开启socket链接-----"+'ws://')
        this.newWs = new WebSocket('ws://' + 'localhost:8080' + '/usertoUser/revocation');  
        this.newWs.onopen = (event) => {  
          // 当 WebSocket 连接打开时，发送认证消息  
          this.authenticate(this.newWs);  
        };  
        
        this.newWs.onmessage = (event) => {  
          // 处理从服务器接收到的消息  
          const msg = JSON.parse(event.data);  
          console.log(msg);  
        };
      },
      createRevocationWs(){
        console.log("开启socket链接-----"+'ws://')
        this.RevocationWs = new WebSocket('ws://' + 'localhost:8080' + '/usertoUser/revocation');  
        this.RevocationWs.onopen = (event) => {  
          // 当 WebSocket 连接打开时，发送认证消息  
          this.authenticate(this.RevocationWs);  
        };  
        
        this.RevocationWs.onmessage = (event) => {  
          // 处理从服务器接收到的消息  
          const msg = JSON.parse(event.data);  
          console.log(msg);  
        };
      },
      authenticate(ws){// 认证方法 
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
      },
    },
    created() {
      localStorage.setItem('token','Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTUzNTc2MjUsImlhdCI6MTcxNTI3MTIyNSwiaXNBZG1pbiI6ZmFsc2UsImlzcyI6IiIsInVzZXJJZCI6IjciLCJ1c2VybmFtZSI6Inh4MDAwNSJ9.1gBQEaI79OSpBRs99uZ1QjoHOog-Exkl3x9Z6xnYdTI')
      this.createMessageWs()
      this.createNewWs()
      this.createRevocationWs()
    },
  }
</script>
<style>
.List{
  width:200px;
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