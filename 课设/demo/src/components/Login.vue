<template>
      <div style="width:400px;height:400px;position:absolute;left:550px;top:150px">
      <form  style="color:rgba(220, 228, 253, 0.942);">
        <div class="form-group">
          <label for="exampleInputEmail1" style="color:rgba(220, 228, 253, 0.942);">请输入账号：</label>
          <input type="text" class="form-control"  id="exampleInputEmail1" placeholder="账号" v-model="username">
        </div>
        <div class="form-group">
          <label for="exampleInputPassword1" style="color:rgba(220, 228, 253, 0.942);">密码：</label>
          <input type="password" class="form-control" id="exampleInputPassword1" placeholder="密码" v-model="passwd">
        </div>
        <!-- <div class="form-group">
          <label for="exampleInputFile" style="color:rgba(220, 228, 253, 0.942);">File input</label>
          <input type="file" id="exampleInputFile">
          <p class="help-block"  style="color:rgba(220, 228, 253, 0.942);">Example block-level help text here.</p>
        </div> -->
        <div class="checkbox">
          <label>
            <input type="checkbox" style="color:rgba(220, 228, 253, 0.942);"> 记住我
          </label>
        </div>
        <button @click="click" class="btn btn-default">登录</button>
        <div class="form-group" style="margin-top:10px">
          <label class="exampleInputEmail1">
            没有账号？
            <router-link to="/Register" class="xx1">
              点我注册
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
  },
  data(){
    return {
      username:"",
      passwd:"",
    }
  },
  methods:{
    click()
    {
      console.log(this.username)
      console.log(this.passwd)
    }
  },
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