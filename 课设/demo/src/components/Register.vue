<template>
      <div style="width:400px;height:400px;position:absolute;left:550px;top:150px">
      <form  style="color:rgba(220, 228, 253, 0.942);">
        <div class="form-group">
          <label for="exampleInputEmail1" style="color:rgba(220, 228, 253, 0.942);">邮箱：</label>
          <input type="email" class="form-control"  id="exampleInputEmail1" placeholder="邮箱" v-model="email">
        </div>
        <div class="form-group">
          <label for="exampleInputPassword1" style="color:rgba(220, 228, 253, 0.942);">用户名：</label>
          <input type="text" class="form-control" id="exampleInputPassword1" placeholder="用户名" v-model="username">
        </div>
        <!-- <div class="form-group">
          <label for="exampleInputFile" style="color:rgba(220, 228, 253, 0.942);">File input</label>
          <input type="file" id="exampleInputFile">
          <p class="help-block"  style="color:rgba(220, 228, 253, 0.942);">Example block-level help text here.</p>
        </div> -->
        <div class="form-group">
          <label for="exampleInputPassword1" style="color:rgba(220, 228, 253, 0.942);">密码：</label>
          <input type="password" class="form-control" id="exampleInputPassword1" placeholder="密码" v-model="passwd">
        </div>
        <div class="form-group">
          <label for="exampleInputPassword1" style="color:rgba(220, 228, 253, 0.942);">确认密码</label>
          <input type="password" class="form-control" id="exampleInputPassword1" placeholder="再次输入密码" v-model="passwdes">
        </div>
        <button class="btn btn-default" @click="check">注册</button>
        <div class="form-group" style="margin-top:10px">
          <label class="exampleInputEmail1">
            已有账号？
            <router-link to="/Login" class="xx1">
              点我登录
            </router-link>
          </label>
        </div>
      </form> 
    </div> 

</template>

<script setup>  
import { ref } from 'vue';  
import TopBar from './TopBar.vue';  
import SideBar from './SideBar.vue';  
import Demo from './Demo.vue';  
import Login from './Login.vue';  
import axios from 'axios';  
  
// 响应式数据  
const email = ref('');  
const username = ref('');  
const password = ref('');  
const passwordes = ref('');  
  
// 方法  
function check(){  
  if (username.value === '' || password.value === '') {  
    alert('输入框不能为空');  
    return;  
  }  
  // 如果需要二次确认密码  
  /* if (password.value !== password2.value) {  
    alert('密码不同');  
    return;  
  } */  
  if (username.value.length < 5 || username.value.length > 20) {  
    alert('账号必须在5-20位');  
    return;  
  }  
  submit();  
};  
  
function submit() {  

    axios.post('http://localhost:8080/user/register', {  
      username: username.value,  
      password: password.value,  
      password2: passwordes.value,  
      email: email.value  
    }).then(response=>{
      console.log('--------');  
      console.log(response.data);  
      localStorage.setItem('username', response.data.username); 
      this.$router.push('login')
    }).catch (err=> {  
      console.log('----1111----');  
      console.error(err.response.data);  
    })
};  
// 在模板中可以直接使用 email, username, check, submit 等变量和方法  
</script>  


<style>
.xx1{
  color:rgba(220, 228, 253, 0.942);
}

.xx1:hover{
  color:rgb(98, 98, 211);
}
</style>