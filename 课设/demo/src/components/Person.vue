<template>
  <div>
    <ul style="color:rgb(207, 234, 244);margin-left:150px;margin-top:80px;font-size:25px">
      <div style="display:flex;">
        <img :src="user.img" style="height:100px;width:100px;border-radius:50%"  />
        <input type="file" style="margin-left:40px;margin-top:60px;font-size:15px" />
      </div>
      <li style="margin-top:20px">账号：<input style="margin-left:50px;background-color:rgb(105, 105, 105);border:0px" type="text" v-model="user.username"/></li>
      <li style="margin-top:10px">用户名：<input style="margin-left:26px;background-color:rgb(105, 105, 105);border:0px" type="text" v-model="user.account"/></li>
      <li style="margin-top:10px">个性签名：<input style="background-color:rgb(105, 105, 105);border:0px" type="text" v-model="user.signed"/></li>
      <li style="margin-top:10px">邮箱：<input style="margin-left:50px;background-color:rgb(105, 105, 105);border:0px" type="text" v-model="user.email"/></li>
      <li style="margin-top:10px">生日：<input style="margin-left:50px;background-color:rgb(105, 105, 105);border:0px" type="text" v-model="user.birthday"/></li>
      <li style="margin-top:10px">居住地：<input style="margin-left:24px;background-color:rgb(105, 105, 105);border:0px" type="text" v-model="user.city"/></li>
    </ul>
    <div>
      <button @click="save()" style="margin-top:20px;margin-left:300px;width:110px;height:40px;font-size:20px;line-height:40px;border-radius:10px">保存</button>
    </div>
  </div>
</template>
<script setup>
import {ref,reactive,onMounted} from 'vue'
import { ElNotification } from 'element-plus'
import service from '../axios-instance'
// const user=reactive(JSON.parse(localStorage.getItem('user')))
var user = reactive(JSON.parse(localStorage.getItem('user')))
onMounted(()=>{
  
  // service.post("http://localhost:8080/user/fidUser",{"username":localStorage.getItem("username")})
  // .then(tmp=>{
  //   user.username=tmp.data.username
  //   user.account=tmp.data.account
  //   user.signed=tmp.data.signed
  //   user.email=tmp.data.email
  //   user.pht=tmp.data.img
  //   user.birthday=tmp.data.birthday
  //   user.city=tmp.data.city
  // })
})
function save()
{
  console.log(user)
  service.post("http://localhost:8080/user/update",{
    'account':user.account,
    'signed':user.signed,
    'birthday':user.birthday
  })
  .then(res=>{
    localStorage.setItem('user',JSON.stringify(user))
    ElNotification({
      title: 'Success',
      message: '修改成功',
      type: 'success',
    }) 
  }).catch(err=>{
    ElNotification({
      title: 'Error',
      message: '修改失败',
      type: 'error',
    })
  })
}
</script>