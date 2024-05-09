<template>
    <div class="side">
        <router-link to="/person" style="text-decoration: none;">
          <div style="margin-top:30px;margin-bottom:50px;">
            <img style="width:120px;height:120px;border-radius:50%" alt="" src="../assets/123.jpg"/>
          </div>
          </router-link>
        <!-- <span class="slice"></span> -->
        <router-link to="/person" style="text-decoration: none;">
           <div class="to">个人信息</div>
        </router-link>
        <!-- <span class="slice"></span> -->
        <router-link to="/space" style="text-decoration: none;">
           <div class="to">个人空间</div>
        </router-link>
        <!-- <span class="slice"></span> -->
           <router-link to="/friendlist" style="text-decoration: none;">
        <div class="to">好友列表</div>
           </router-link>
            <router-link to="/advices" style="text-decoration: none;">
        <div class="to">群聊列表</div>
           </router-link>
        <!-- <span class="slice"></span> -->
        <router-link to="/evaluate" style="text-decoration: none;">
           <div class="to">退出登录</div>
        </router-link>
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

.pht{
  size:60px;
}

.slice{
  line-height: 40px;
  font-size: 20px;
  margin-left: 10px;
}

.side{
  margin-top:20px;
  text-align: center;
  color:#222226;
  display: block;
  height: 600px;
  width: 200px;
  margin-bottom: 20px;;
  border:2px double rgb(138, 137, 137);
  background-color: #525257a2;
  border-radius:18px;
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
  position: relative;
  color:rgba(220, 228, 253, 0.942);
  width:100%;
  height:40px;
  font-size:20px;
  margin-top:20px;
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

.to::before{
  content: '';
            height: 4px;
            background: rgb(137, 203, 233);
            /* 伪元素默认样式 display: inline;所以需要转成inline-block宽高才会生效 */
            display: inline-block;
            /* 通过定位使下划线在最低层 */
            position: absolute;
            border-radius: 18px;;
            bottom: -6px;
            width: 0;
            /* 加上一个过渡效果，使之丝滑一些 */
            transition: width 0.36s;
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
  color:rgb(57, 144, 220)
}

.to:hover::before{
  width: 100px;
}
.inputType:focus{
  outline: none;   
  border: 1px solid rgb(230, 32, 213);
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


