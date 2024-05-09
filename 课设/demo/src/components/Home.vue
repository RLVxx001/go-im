<template>
  <div class="xx2">
    欢迎来到xxx聊天室
    {{color}}
  </div>
</template>
<script>
import bus from "../EventBus/eventbus.js";
  import App from "../App.vue"
  
  export default{
    inject: ["color"],
    data(){
      return{
        val:"123"
      }
    },
  components:{
    // Header,
    // Footer,
    // App,
  },
  created() {
    console.log(this.val);
    bus.on("countChange", (v) => {
      this.val = v;
    });
  },
}
</script>
<style>
.xx2{
  font-size: 50px;
  width:500px;
  height:300px;
  
  color:rgba(220, 228, 253, 0.942);
}
</style>
