<template>
  <div>
    <ul>
      <li>用户名：{{user.username}}</li>
      <li>个人名：{{user.account}}</li>
      <li>邮箱：{{user.email}}</li>
      <li>头像：</li>
    </ul>
  </div>
</template>
<script setup>
import {ref,reactive} from 'vue'
const user=reactive(JSON.parse(localStorage.getItem('user')))
</script>