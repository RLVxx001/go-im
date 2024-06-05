import { defineStore } from "pinia";

export const useUserStore=defineStore('user',{
    state:()=>{
        return {
            username:'',
            token:localStorage.getItem("token")
        }
    },
    actions:{
        logout(){
            localStorage.removeItem("token")
            this.token=''
            this.username=''
        }
    }
})
export const useWsStore = defineStore('ws', {  
    state: () => ({  
      Frientmessages: [], // 存储消息的数组  
      Frientmessagecount: 0, // 消息数量  
      Frientrevocations: [], // 存储消息的数组  
      Frientrevocationcount: 0, // 消息数量
      Frientusers: [], // 存储消息的数组  
      Frientusercount: 0, // 消息数量
      Frientcount:0,

      Groupmessages: [], // 存储消息的数组  
      Groupmessagecount: 0, // 消息数量  
      Grouprevocations: [], // 存储消息的数组  
      Grouprevocationcount: 0, // 消息数量
      Groupusers: [], 
      Groupusercount: 0, 
      Groupcount:0,
      event: -1
    }),  
    actions: {  
      addMessage(message,event) {  
        console.log("event:"+event)
        if(event=='/usertoUser/revocation')
        {
            if(this.event==0)
            {
                this.Frientcount=0
                this.Frientrevocationcount++
                this.Frientrevocations.push(message)
            }
            else
            {
                this.Frientcount++
                this.Frientrevocationcount=0
                this.Frientrevocations.splice(0,this.Frientrevocations.length)
            }
        }
        else if(event=='/usertoUser/send')
        {
            if(this.event==0)
            {
                this.Frientcount=0
                this.Frientmessagecount++
                this.Frientmessages.push(message)
            }
            else
            {
                this.Frientcount++
                this.Frientmessagecount=0
                this.Frientmessages.splice(0,this.Frientmessages.length)
            }
        }
        else if(event=="/usertoUser")
        {
            if(this.event==0)
                {
                    this.Frientcount=0
                    this.Frientusercount++
                    this.Frientusers.push(message)
                }
                else
                {
                    this.Frientcount++
                    this.Frientusercount=0
                    this.Frientusers.splice(0,this.Frientusers.length)
                } 
        }
        else if(event=="/group/sendMessage")
        {
            if(this.event==1)
            {
                this.Groupcount=0
                this.Groupmessagecount++
                this.Groupmessages.push(message)
            }
            else
            {
                this.Groupcount++
                this.Groupmessagecount=0
                this.Groupmessages.splice(0,this.Groupmessages.length)
            }
        }
        else if(event=="/group/revocationMessage")
        {
            if(this.event==1)
            {
                this.Groupcount=0
                this.Grouprevocationcount++
                this.Grouprevocations.push(message)
            }
            else
            {
                this.Groupcount++
                this.Grouprevocationcount=0
                this.Grouprevocations.splice(0,this.Grouprevocations.length)
            }
        }
        else if(event=="/group/createGroup")
        {
            if(this.event==1)
            {
                this.Groupcount=0
                this.Groupusercount++
                this.Groupusers.push(message)
            }
            else
            {
                this.Groupcount++
                this.Groupusercount=0
                this.Groupusers.splice(0,this.Groupusers.length)
            }
        }
        console.log("进入paina的消息："+message)
      },  
      
      async readFrientUsers(){
        // 复制当前的消息数组，以便返回给调用者  
        console.log(this.Frientusers)
        const messagesToRead = this.Frientusers.slice();  
    
        // 清空数组和重置计数器  
        this.Frientusercount -= messagesToRead.length; 
        this.Frientusers.splice(0,messagesToRead.length)  
        
        // 返回一个 Promise，该 Promise 解析为被读取的消息数组  
        return messagesToRead
      },
      async readFrientRevocations() {  
        // 复制当前的消息数组，以便返回给调用者  
        console.log(this.Frientrevocations)
        const messagesToRead = this.Frientrevocations.slice();  
    
        // 清空数组和重置计数器  
        this.Frientrevocationcount -= messagesToRead.length; 
        this.Frientrevocations.splice(0,messagesToRead.length)  
        
        // 返回一个 Promise，该 Promise 解析为被读取的消息数组  
        return messagesToRead
      }, 
      async readFrientMessages() {  
        // 复制当前的消息数组，以便返回给调用者  
        console.log(this.Frientmessages)
        const messagesToRead = this.Frientmessages.slice();  
    
        // 清空数组和重置计数器  
        this.Frientmessagecount -= messagesToRead.length; 
        this.Frientmessages.splice(0,messagesToRead.length)  
        
        // 返回一个 Promise，该 Promise 解析为被读取的消息数组  
        return messagesToRead
      }, 
      async readGroupMessages() {  
        // 复制当前的消息数组，以便返回给调用者  
        console.log(this.Groupmessages)
        const messagesToRead = this.Groupmessages.slice();  
    
        // 清空数组和重置计数器  
        this.Groupmessagecount -= messagesToRead.length; 
        this.Groupmessages.splice(0,messagesToRead.length)  
        
        // 返回一个 Promise，该 Promise 解析为被读取的消息数组  
        return messagesToRead
      },  
      async readGroupRevocations() {  
        // 复制当前的消息数组，以便返回给调用者  
        console.log(this.Groupmessages)
        const messagesToRead = this.Grouprevocations.slice();  
    
        // 清空数组和重置计数器  
        this.Grouprevocationcount -= messagesToRead.length; 
        this.Grouprevocations.splice(0,messagesToRead.length)  
        
        // 返回一个 Promise，该 Promise 解析为被读取的消息数组  
        return messagesToRead
      },  
      async readGroupUsers() {  
        // 复制当前的消息数组，以便返回给调用者  
        console.log(this.Frientusers)
        const messagesToRead = this.Groupusers.slice();  
    
        // 清空数组和重置计数器  
        this.Groupusercount -= messagesToRead.length; 
        this.Groupusers.splice(0,messagesToRead.length)  
        
        // 返回一个 Promise，该 Promise 解析为被读取的消息数组  
        return messagesToRead
      }, 

    },  
  });