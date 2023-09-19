<template>
  <n-card style="width: 300px; margin: auto;" embedded :bordered="false">
    <n-form :model="data.loginForm" label-placement="left" label-width="auto">
      <n-form-item label="用户名">
        <n-input round v-model:value="data.loginForm.username" @keydown.enter="login" />
      </n-form-item>
      <n-form-item label="密码">
        <n-input round v-model:value="data.loginForm.password" type="password" @keydown.enter="login" />
      </n-form-item>
      <n-form-item>
        <n-button round style="width: 100%;" @click="login" secondary type="primary">登录</n-button>
      </n-form-item>
    </n-form>
  </n-card>
</template>

<script>
import $backend from '../backend'
import { NForm, NFormItem, NInput, NButton, NCard, useMessage } from 'naive-ui'
import { reactive } from 'vue'

export default {
  name: 'LoginView',
  components: {
    NForm,
    NFormItem,
    NInput,
    NButton,
    NCard
  },
  setup () {
    const data = reactive({
      loginForm: {
        username: '',
        password: ''
      }
    })

    const message = useMessage()

    async function login () {
      const { username, password } = data.loginForm
      const res = await $backend.login(username, password)
      if (res.errcode === 0) {
        data.logined = true
        message.success('登录成功')
        // TODO: redirect to home page
      } else {
        message.error(`登录失败: ${res.errmsg}`)
      }
    }

    return {
      data,
      login
    }
  }
}
</script>
