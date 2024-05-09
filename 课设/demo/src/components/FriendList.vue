<template>
  <div style="color:rgba(220, 228, 253, 0.942);">
    <div style="display:flex">
      <div class="List">
        <!-- <div class="find">
          <input  class="inputT" placeholder='内容'/>
          <button class="btn">搜索</button>
        </div> -->
        <div style="height:25px;font-size:17px;">
          <img src="#" style="margin-right:20px;margin-left:10px;"/>
            sdsdsds
        </div>
        <div style="height:25px;font-size:17px;">
          <img src="#" style="margin-right:20px;margin-left:10px;"/>
            sdsdsds
        </div>
      </div>
      <div>
        <div class="Message">
          <div class="Top">
          sd
          </div>
        </div>
        <div style="width:1px;backgroud-color:black"></div>
        <div class="Chat">

        </div>
      </div>
    </div>
  </div>
</template>
<script>
  export default{
    data(){
      return{
        UserList:[],
      }
    }
  }
</script>
<style>
.List{
  width:200px;
  height:596px;
  border-top-left-radius: 18px;
  border-bottom-left-radius: 18px;
  background-color: #b9b8b8a2;
}

.Message{
  width:607px;
  height:420px;
  border-top-right-radius: 18px;
  background-color: #a7a7aca2;
}



.Chat
{
  height:176px;
  background-color: #c8c8c9a2;
  width: 607px;;
  border-bottom-right-radius: 18px;;
}

.find{
    height: 32px;
    width: 200px;
    background-color: #525257a2;
  
}

.bTn{
  margin-top:14px;
  color: rgb(207, 230, 244);
  font-size:17px;
  background-color: #2d9cf8;
  text-align: center;
  border-left:1px solid rgba(0, 0, 0, 0.45);
  width:50px;
  height: 34px;
  line-height: 34px;
  border-top-right-radius: 18px;
  border-bottom-right-radius: 18px
}

.inputT{
    margin-left: 210px;
    margin-top:15px;
  width:150px;
  height: 32px;
  text-indent: 1em;
   line-height: 34px;
  border: 1px solid black;
  border-top-left-radius: 17px;
  border-bottom-left-radius: 17px;
  font-size: 12px;
}
</style>