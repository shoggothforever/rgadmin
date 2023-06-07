<template>
  <div class="login-container">
    <el-form ref="loginFormRef" :model="loginForm" :rules="loginRules" class="login-form" autocomplete="on" label-position="left">
      <div class="title-container">
        <h3 class="title">登录</h3>
      </div>

      <el-form-item prop="username">
        <el-input ref="usernameRef" v-model="loginForm.staffCode" placeholder="用户名" name="username" type="text" tabindex="1" autocomplete="on" />
      </el-form-item>

      <el-tooltip v-model="capsTooltip" content="Caps lock is On" placement="right" manual>
        <el-form-item prop="password">
          <el-input :key="passwordType" ref="passwordRef" v-model="loginForm.password" :type="passwordType"
            placeholder="密码" name="password" tabindex="2" autocomplete="on" @keyup.native="checkCapslock"
            @blur="capsTooltip = false" @keyup.enter.native="handleLogin" />
        </el-form-item>
      </el-tooltip>

      <el-button :loading="loading" type="primary" style="width:100%;margin-bottom:30px;"
      @click.native.prevent="handleLogin">登录</el-button>
    </el-form>
  </div>
</template>

<script setup>
import { validUsername } from '@/utils/validate'
import { ref, reactive, onMounted } from 'vue'
import { useUserStore } from '../../stores/user';
import { useRouter } from 'vue-router'

const validateUsername = (rule, value, callback) => {
  if (!validUsername(value)) {
    callback(new Error('请输入正确的用户名'))
  } else {
    callback()
  }
}
const validatePassword = (rule, value, callback) => {
  if (value.length < 1) {
    callback(new Error('密码不能少于1位'))
  } else {
    callback()
  }
}

const userStore = useUserStore()
const router = useRouter()
const passwordType = ref('password')
const capsTooltip = ref(false)
const loading = ref(false)
const usernameRef = ref(null)
const passwordRef = ref(null)
const loginFormRef = ref(null)
// const redirect = ref(undefined)
const loginForm = reactive({
  staffCode: '',
  password: ''
})
const loginRules = reactive({
  staffCode: [{ required: true, validator: validateUsername, trigger: 'blur' }],
  password: [{ required: true, validator: validatePassword, trigger: 'blur' }]
})

const checkCapslock = (e) => {
  const { key } = e
  capsTooltip.value = key && key.length === 1 && key >= 'A' && key <= 'Z'
}
// const showPwd = () => {
//   if (passwordType.value === 'password') {
//     passwordType.value = ''
//   } else {
//     passwordType.value = 'password'
//   }
//   nextTick(() => {
//     passwordRef.value.focus()
//   })
// }
const handleLogin = () => {
  console.log('提交表单')
  loginFormRef.value.validate(valid => {
    if (valid) {
      loading.value = true
      userStore.login(loginForm)
        .then(() => {
          console.log('提交成功了(喜!)')
          router.push('/')
          loading.value = false
        })
        .catch(() => {
          loading.value = false
        })
    } else {
      console.log('提交失败了(悲!)')
      return false
    }
  })
}

onMounted(() => {
  if (loginForm.staffCode === '') {
    usernameRef.value.focus()
  } else if (loginForm.password === '') {
    passwordRef.value.focus()
  }
})
</script>

<!-- <style lang="scss">
/* 修复input 背景不协调 和光标变色 */
/* Detail see https://github.com/PanJiaChen/vue-element-admin/pull/927 */

$bg: #283443;
$light_gray: #fff;
$cursor: #fff;

@supports (-webkit-mask: none) and (not (cater-color: $cursor)) {
  .login-container .el-input input {
    color: $cursor;
  }
}

/* reset element-ui css */
.login-container {
  .el-input {
    display: inline-block;
    height: 47px;
    width: 85%;

    input {
      background: transparent;
      border: 0px;
      -webkit-appearance: none;
      border-radius: 0px;
      padding: 12px 5px 12px 15px;
      color: $light_gray;
      height: 47px;
      caret-color: $cursor;

      &:-webkit-autofill {
        box-shadow: 0 0 0px 1000px $bg inset !important;
        -webkit-text-fill-color: $cursor !important;
      }
    }
  }

  .el-form-item {
    border: 1px solid rgba(255, 255, 255, 0.1);
    background: rgba(0, 0, 0, 0.1);
    border-radius: 5px;
    color: #454545;
  }
}
</style> -->

<style lang="scss" scoped>
.el-input {
  height: 3em;
}
.el-button {
  height: 3em !important;
}
</style>

<style lang="scss" scoped>
$bg: #2d3a4b;
$dark_gray: #889aa4;
$light_gray: #eee;

.login-container {
  min-height: 100%;
  width: 100%;
  background-color: $bg;
  overflow: hidden;

  .login-form {
    position: relative;
    width: 520px;
    max-width: 100%;
    padding: 160px 35px 0;
    margin: 0 auto;
    overflow: hidden;
  }

  .tips {
    font-size: 14px;
    color: #fff;
    margin-bottom: 10px;

    span {
      &:first-of-type {
        margin-right: 16px;
      }
    }
  }

  .svg-container {
    padding: 6px 5px 6px 15px;
    color: $dark_gray;
    vertical-align: middle;
    width: 30px;
    display: inline-block;
  }

  .title-container {
    position: relative;

    .title {
      font-size: 26px;
      color: $light_gray;
      margin: 0px auto 40px auto;
      text-align: center;
      font-weight: bold;
    }
  }

  .show-pwd {
    position: absolute;
    right: 10px;
    top: 7px;
    font-size: 16px;
    color: $dark_gray;
    cursor: pointer;
    user-select: none;
  }

  .thirdparty-button {
    position: absolute;
    right: 0;
    bottom: 6px;
  }

  @media only screen and (max-width: 470px) {
    .thirdparty-button {
      display: none;
    }
  }
}
</style>
