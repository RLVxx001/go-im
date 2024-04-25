<template>
      <div style="width:400px;height:400px;position:absolute;left:550px;top:150px">
      <form  style="color:rgba(220, 228, 253, 0.942);">
        <div class="form-group">
          <label for="exampleInputEmail1" style="color:rgba(220, 228, 253, 0.942);">邮箱：</label>
          <input type="email" class="form-control"  id="exampleInputEmail1" placeholder="邮箱">
        </div>
        <div class="form-group">
          <label for="exampleInputPassword1" style="color:rgba(220, 228, 253, 0.942);">用户名：</label>
          <input type="password" class="form-control" id="exampleInputPassword1" placeholder="用户名">
        </div>
        <!-- <div class="form-group">
          <label for="exampleInputFile" style="color:rgba(220, 228, 253, 0.942);">File input</label>
          <input type="file" id="exampleInputFile">
          <p class="help-block"  style="color:rgba(220, 228, 253, 0.942);">Example block-level help text here.</p>
        </div> -->
        <div class="form-group">
          <label for="exampleInputPassword1" style="color:rgba(220, 228, 253, 0.942);">密码：</label>
          <input type="password" class="form-control" id="exampleInputPassword1" placeholder="密码">
        </div>
        <div class="form-group">
          <label for="exampleInputPassword1" style="color:rgba(220, 228, 253, 0.942);">确认密码</label>
          <input type="password" class="form-control" id="exampleInputPassword1" placeholder="再次输入密码">
        </div>
        <button type="submit" class="btn btn-default">注册</button>
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

<script>
import TopBar from './TopBar.vue'
import SideBar from './SideBar.vue'
import Demo from './Demo.vue'
import Login from './Login.vue'
export default{
  components:{
   TopBar,
    SideBar,
    Demo,
    TopBar,
    SideBar,
    // Index,
    // Login,
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