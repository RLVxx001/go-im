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