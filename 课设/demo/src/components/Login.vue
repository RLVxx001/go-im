<!-- sign_up -->
<template>
    <div
        v-if="loginModel === '注册'"
        class="container a-container"
        id="a-container"
    >
        <form class="form" id="a-form" method="" action="">
            <h2 class="form_title title">欢迎使用FileHub</h2>
            <div class="form__icons">
                <img class="form__icon" src=" " />
                <img class="form__icon" src=" " />
                <img class="form__icon" src=" " />
            </div>
            <!-- <span class="form__span">or use email for registration</span> -->
            <input class="form__input" type="text" placeholder="用户名" />
            <input class="form__input" type="password" placeholder="密码" />
            <input class="form__input" type="text" placeholder="Token" />
            <div class="login-info">
                <a class="form__link" @click="forgetPassword">忘记密码了？</a>
                &nbsp;&nbsp;&nbsp;&nbsp;
                <a class="form__link" @click="switchModel('token')"
                    >获取Token教程</a
                >
                &nbsp;&nbsp;&nbsp;&nbsp;
                <a class="form__link" @click="switchModel('登陆')"
                    >使用账号登陆
                </a>
            </div>
            <button class="form__button button submit">
                注&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;册
            </button>
        </form>
    </div>
    <div
        v-else-if="loginModel === '登陆' || loginModel === 'token'"
        class="container b-container"
        id="b-container"
    >
        <form class="form" id="b-form" method="" action="">
            <h2 class="form_title title">欢迎登陆FileHub</h2>
            <div class="form__icons">
                <img class="form__icon" src=" " />
                <img class="form__icon" src=" " />
                <img class="form__icon" src=" " />
            </div>
            <!-- <span class="form__span">或者使用Token登陆</span> -->
            <input
                v-if="loginModel === '登陆'"
                class="form__input"
                type="text"
                placeholder="用户名"
            />
            <input
                v-if="loginModel === '登陆'"
                class="form__input"
                type="password"
                placeholder="密码"
            />
            <input
                v-if="loginModel === 'token'"
                class="form__input"
                type="password"
                placeholder="Token"
            />
            <div class="login-info">
                <a class="form__link" @click="forgetPassword">忘记密码了？</a>
                &nbsp;&nbsp;&nbsp;&nbsp;
                <template v-if="loginModel === '登陆'">
                    <a class="form__link" @click="switchModel('token')"
                        >使用Token登陆
                    </a>
                    &nbsp;&nbsp;&nbsp;&nbsp;
                </template>
                <template v-if="loginModel === 'token'">
                    <a class="form__link" @click="switchModel('登陆')"
                        >使用账号登陆
                    </a>
                    &nbsp;&nbsp;&nbsp;&nbsp;
                </template>
                <a class="form__link" @click="switchModel('注册')"
                    >注册新账户</a
                >
            </div>
            <button class="form__button button submit">
                登&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;陆
            </button>
        </form>
    </div>
</template>
 
<script setup lang="ts">
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
 
// 忘记密码提示
const forgetPassword = () => {
  ElMessage({
    message: '请联系微信或者QQ：648133599',
    type: 'success',
  })
}
 
// 模式：登陆，token，注册
let loginModel = ref('登陆')
 
const switchModel = (model: string) => {
    loginModel.value = model
}
</script>
 
<style scoped>
@import '@/style/login.css';
 
.login-info {
    margin-top: 10px;
}

</style>