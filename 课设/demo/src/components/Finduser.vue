<template>
  <div style="">
    <div style="color:rgb(207, 234, 244);margin-left:150px;margin-top:80px;font-size:25px">
      <div style="display:flex;">
        <img :src="res.pht" style="height:100px;width:100px;border-radius:50%"  />
      </div>
      <div style="margin-top:20px;display:flex">用户名：<div style="color:rgb(207, 234, 244);margin-left:24px;background-color:rgb(105, 105, 105);">{{res.username}}</div></div>
      <div style="margin-top:10px;display:flex">账号：<div style="color:rgb(207, 234, 244);margin-left:50px;background-color:rgb(105, 105, 105);border:0px">{{res.account}}</div></div>
      <div style="margin-top:10px;display:flex">个性签名：<div style="background-color:rgb(105, 105, 105);border:0px" >{{res.signed}}</div></div>
      <div style="margin-top:10px;display:flex">生日：<div style="margin-left:50px;background-color:rgb(105, 105, 105);border:0px">{{res.birthday}}</div></div>
      <div style="margin-top:10px;display:flex">居住地：<div style="margin-left:24px;background-color:rgb(105, 105, 105);border:0px">{{res.city}}</div></div>
    </div>
    <div>
      <div style="display:flex">
        <button plain @click="open" style="margin-top:10px;margin-left:250px;width:110px;height:40px;font-size:20px;line-height:40px;border-radius:10px">添加好友</button>
        <button plain @click="toSpace" style="margin-top:10px;margin-left:30px;width:110px;height:40px;font-size:20px;line-height:40px;border-radius:10px">访问空间</button>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted ,h,reactive,nextTick, isRef } from 'vue'; 
import { ElNotification,ElScrollbar } from 'element-plus'
import { ClickOutside as vClickOutside } from 'element-plus'
import { ElMessage, ElMessageBox } from 'element-plus'
import service from '../axios-instance'
import { useRouter } from 'vue-router' 
import { useUserStore } from '../store/user';
const userStore=useUserStore()
const router = useRouter()  
const msg = ref('12345678');
const popoverRef = ref()
const open = () => {
  ElMessageBox.prompt('请输入验证信息', '验证', {
    confirmButtonText: 'OK',
    cancelButtonText: 'Cancel',
    inputPattern:
      /(?:\.[\w!#$%&'*+/=?^_`{|}~-]+)*/,
    inputErrorMessage: 'Invalid Email',
  })
    .then(({ value }) => {
      
      service.post('http://localhost:8080/userApplication',{
        'userOwner':JSON.parse(localStorage.getItem('user')).id-0,
        'class':0,
        'target':res.userId-0,
        'text':value,
        'stats':0
      }).then(res=>{
        ElMessage({
          type: 'success',
          message: "发送成功！",
        })
      }).catch(err=>{
        console.error(err)
        ElMessage({
          type: 'error',
          message: err.response.data.errorMessage,
        })
      })
    })
    .catch(() => {
      ElMessage({
        type: 'info',
        message: '取消',
      })
    })
}

var user = reactive({})
    var a = ref(localStorage.getItem("account"))
var res = reactive({})
onMounted(()=>{
  console.log(a.value);
  service.post("http://localhost:8080/user/fidUser",{"username":a.value})
  .then(tmp=>{
    res.username=tmp.data.username
    res.account=tmp.data.account
    res.signed=tmp.data.signed
    res.email=tmp.data.email
    res.pht=tmp.data.img
    res.birthday=tmp.data.birthday
    res.city=tmp.data.city
    res.userId=tmp.data.userId
  })
  .catch(err=>{
    alert("未找到该用户")
    router.push('/person')
  })
})

function toSpace(){
  localStorage.setItem("toId",res.userId)
  router.push("/toSpace")
}
</script>

<style>

</style>