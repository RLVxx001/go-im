<template>
  <div style="">
    <div style="color:rgb(207, 234, 244);margin-left:150px;margin-top:80px;font-size:25px">
      <div style="display:flex;">
        <img src="../assets/123.jpg" style="height:100px;width:100px;border-radius:50%"  />
      </div>
      <div style="margin-top:20px;display:flex">用户名：<div style="color:rgb(207, 234, 244);margin-left:24px;background-color:rgb(105, 105, 105);">{{res.username}}</div></div>
      <div style="margin-top:10px;display:flex">账号：<div style="color:rgb(207, 234, 244);margin-left:50px;background-color:rgb(105, 105, 105);border:0px">{{res.account}}</div></div>
      <div style="margin-top:10px;display:flex">个性签名：<div style="background-color:rgb(105, 105, 105);border:0px" >{{res.signed}}</div></div>
      <div style="margin-top:10px;display:flex">邮箱：<div style="margin-left:50px;background-color:rgb(105, 105, 105);border:0px">{{res.email}}</div></div>
      <div style="margin-top:10px;display:flex">生日：<div style="margin-left:50px;background-color:rgb(105, 105, 105);border:0px">{{res.birthday}}</div></div>
      <div style="margin-top:10px;display:flex">居住地：<div style="margin-left:24px;background-color:rgb(105, 105, 105);border:0px">{{res.city}}</div></div>
    </div>
    <div>
      <button plain @click="open" style="margin-top:10px;margin-left:300px;width:110px;height:40px;font-size:20px;line-height:40px;border-radius:10px">添加好友</button>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted ,h,reactive,nextTick, isRef } from 'vue'; 
import { ElNotification,ElScrollbar } from 'element-plus'
import service from '../axios-instance'
import { ClickOutside as vClickOutside } from 'element-plus'
import { ElMessage, ElMessageBox } from 'element-plus'
import axios from 'axios';

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
      msg.value=value
      localStorage.setItem("msg",value);
      ElMessage({
        type: 'success',
        message: "发送成功！",
      })
      axios.post("",{
        "userOwner":""
      })
      console.log(msg.value)
    })
    .catch(() => {
      ElMessage({
        type: 'info',
        message: '取消',
      })
    })
}
var user = reactive({})
    onMounted(()=>{
      axios.post("",{      })
      .then(res=>{

      })
    })
    var a = ref(localStorage.getItem("account"))
onMounted(()=>{
  console.log(a.value);
})
var res = reactive({
  "username":"Fantasy",
  "account":"2460437333",
  "signed":"xxxxxxxxxxxxxxxxxxxxxxxxxx",
  "email":"2460437333@qq.com",
  "pht":"../assets/123.jpg",
  "birthday":"2003/6/10",
  "city":"福建"
})
</script>

<style>

</style>