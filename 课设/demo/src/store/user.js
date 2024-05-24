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
      messages: [], // 存储消息的数组  
      count: 0, // 消息数量  
      event: 0
    }),  
    actions: {  
      addMessage(message) {  
        console.log("进入paina的消息："+message)
        this.messages.push(message);  
        this.count++;  
      },  
      async readAndClearMessages() {  
        // 复制当前的消息数组，以便返回给调用者  
        console.log(this.messages)
        const messagesToRead = this.messages.slice();  
    
        // 清空数组和重置计数器  
        this.count -= messagesToRead.length; 
        this.messages.splice(0,messagesToRead.length)  
        
        // 返回一个 Promise，该 Promise 解析为被读取的消息数组  
        return messagesToRead
      },  
    },  
  });