<template>
    <div class="side">
        <router-link to="/unusedOrder" style="text-decoration: none;">
            <div class="to">首页</div>
        </router-link>
        <!-- <span class="slice"></span> -->
        <router-link to="/usedOrder" style="text-decoration: none;">
           <div class="to">个人空间</div>
        </router-link>
        <!-- <span class="slice"></span> -->
        <router-link to="/usedOrder" style="text-decoration: none;">
           <div class="to">导航三</div>
        </router-link>
        <!-- <span class="slice"></span> -->
           <router-link to="/advices" style="text-decoration: none;">
        <div class="to">导航四</div>
           </router-link>
        <!-- <span class="slice"></span> -->
        <router-link to="/evaluate" style="text-decoration: none;">
           <div class="to">导航五</div>
        </router-link>
       <!-- <div class="search">
         <div contenteditable class="inputType" placeholder='请输入文字'></div>
        <div class="button">搜索</div>
       </div> -->
       <router-link to="/Login" style="text-decoration: none;">
       <div class="to">
          登录
       </div>
      </router-link>
      <router-link to="/Register" style="text-decoration: none;">
       <div class="to">
          注册
       </div>
      </router-link>
      <!-- <router-view></router-view> -->
   </div>
</template>
<script>
export default {
    name:'TopBar',
}
</script>
<style>

router-link{
  text-decoration: none;
}

.slice{
  line-height: 40px;
  font-size: 20px;
  margin-left: 10px;
}

.side{
  text-align: center;
  color:#222226;
  display: block;
  height: 800px;
  width: 150px;
  border:1px double linen;
  background-color: #525257a2;
}

.login{
  color:white;
  width:100px;
  height:40px;
  text-align: center;
  margin-left: 50px;
  line-height: 40px;
}

.login:hover{
  color:rgb(61, 129, 188)
}

.register{
  color:white;
  width:100px;
  height:40px;
  text-align: center;
  line-height: 40px;
}

.register:hover{
  color:rgb(61, 129, 188)
}

.title-left-image{
    height: 48px;
    width:120px
}

.to{
  color:rgba(220, 228, 253, 0.942);
  width:100%;
  height:60px;
  font-size:20px;
  margin-top:10px;
  text-align: center;
  line-height: 40px;
  /*background: -webkit-linear-gradient(135deg,
    #77a0dd,
    #4c6da7 25%,
                    #39bedf 50%,
                    #597ed4 55%,
                    #4a8fbd 60%,
                    #8b2ce0 80%,
                    #af83f0 95%,
                    #c759e0);
            -webkit-text-fill-color: transparent;
            -webkit-background-clip: text;
            -webkit-background-size: 200% 100%;
            -webkit-animation: flowCss 12s infinite linear;*/

}
/*
@-webkit-keyframes flowCss {
            0% {
                background-position: 0 0;
            }

            100% {
                background-position: -400% 0;
            }
        }

*/
.to:hover{
  /*-webkit-animation: flowCss 4s infinite linear;*/
  color:rgb(61, 129, 188)
}

.search{
    margin-left: 50px;
    height: 32px;
    width: 500px;
    background-color:#fff;

    display: flex;
}
.inputType{
  width:400px;
  height: 30px;
  text-indent: 1em;
   line-height: 34px;
  background-color: #fff;
  border: 1px solid black;
  border-top-left-radius: 17px;
  border-bottom-left-radius: 17px;
  font-size: 12px;
}
.inputType:focus{
  outline: none;   
  border: 1px solid rgb(230, 32, 213);
}
.button{
  color: #fff;
  text-align: center;
  background: rgba(204, 51, 0, 1);
  border-left:1px solid rgba(0, 0, 0, 0.45);
  width:100px;
  height: 32px;
  line-height: 34px;
  border-top-right-radius: 18px;
  border-bottom-right-radius: 18px
}
.title-right{
     width: 380px;
    height: 48px;
    display: flex;
}
.title-right-image{
    width: 40px;
    height: 40px;
    border: 1px solid black;
    border-radius: 20px;
    margin: auto ;
    margin-left: 40px;
}
</style>


