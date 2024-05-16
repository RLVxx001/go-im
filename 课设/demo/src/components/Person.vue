<template>
  <div>
    <ul style="color:rgb(207, 234, 244);margin-left:150px;margin-top:80px;font-size:25px">
      <div style="display:flex;">
        <img src="../assets/123.jpg" style="height:100px;width:100px;border-radius:50%"  />
        <input type="file" style="margin-left:40px;margin-top:60px;font-size:15px" />
      </div>
      <li style="margin-top:20px">用户名：<input style="margin-left:24px;background-color:rgb(105, 105, 105);border:0px" type="text" :value="user.username"/></li>
      <li style="margin-top:10px">账号：<input style="margin-left:50px;background-color:rgb(105, 105, 105);border:0px" type="text" :value="user.account"/></li>
      <li style="margin-top:10px">个性签名：<input style="background-color:rgb(105, 105, 105);border:0px" type="text" :value="user.signed"/></li>
      <li style="margin-top:10px">邮箱：<input style="margin-left:50px;background-color:rgb(105, 105, 105);border:0px" type="text" :value="user.email"/></li>
      <li style="margin-top:10px">生日：<input style="margin-left:50px;background-color:rgb(105, 105, 105);border:0px" type="text" :value="user.birthday"/></li>
      <li style="margin-top:10px">居住地：<input style="margin-left:24px;background-color:rgb(105, 105, 105);border:0px" type="text" :value="user.city"/></li>
    </ul>
    <div>
      <button style="margin-top:20px;margin-left:300px;width:110px;height:40px;font-size:20px;line-height:40px;border-radius:10px">保存</button>
    </div>
  </div>
</template>
<script setup>
import {ref,reactive} from 'vue'
// const user=reactive(JSON.parse(localStorage.getItem('user')))
var user = reactive({
  "username":"Fantasy",
  "account":"2460437333",
  "signed":"xxxxxxxxxxxxxxxxxxxxxxxxxx",
  "email":"2460437333@qq.com",
  "pht":"../assets/123.jpg",
  "birthday":"2003/6/10",
  "city":"福建"
})
</script>