<template>
  <div style="min-height: 100vh;">
    <n-message-provider>
      <n-dialog-provider>
        <n-layout has-sider style="min-height: 100vh;">
          <n-layout-sider
            bordered
            :width="200"
            style="min-height: 100vh;"
          >
            <n-menu
              v-model:value="data.checkedMenu"
              :options="menuOptions"
              @update:value="onMenuUpdate"
            />
          </n-layout-sider>
          <n-layout>
            <div style="margin: 15px;">
              <router-view />
            </div>
          </n-layout>
        </n-layout>
      </n-dialog-provider>
    </n-message-provider>
  </div>
</template>

<script>
import $backend from './backend'
import { NMenu, NLayout, NLayoutSider, NMessageProvider, NIcon, NDialogProvider } from "naive-ui"
import { h, reactive, computed, onMounted } from 'vue'
import { LogIn, IdCardOutline } from '@vicons/ionicons5'

function renderIcon (icon) {
  return () => h(NIcon, null, { default: () => h(icon) })
}

const menuOptions = [
  { label: '管理员', key: 'admin', icon: renderIcon(IdCardOutline) },
]

const loginOptions = [
  { label: '登录', key: 'login', icon: renderIcon(LogIn) }
]

export default {
  name: 'App',
  components: {
    NMenu,
    NLayout,
    NLayoutSider,
    NMessageProvider,
    NDialogProvider
  },
  setup () {
    const data = reactive({
      logined: false,
      checkedMenu: '',
      authedPages: {}
    })

    const options = computed(() => {
      if (!data.logined) return loginOptions
      let resp = []
      menuOptions.forEach(opt => {
        let nopt = opt
        if (opt.children) {
          let children = []
          opt.children.forEach(copt => {
            if (copt.key in data.authedPages) children.push(copt)
          })
          if(children.length > 0) {
            nopt.children = children
            resp.push(nopt)
          }
        }
        else {
          if (opt.key in data.authedPages) resp.push(nopt)
        }
      })
      return resp
    })

    async function getPageAuth () {
      return $backend.getPageAuth().then(res => {
        console.log('getPageAuth res', res)
        res.data.forEach(key => data.authedPages[key] = true)
      })
    }

    async function getLoginInfo() {
      return $backend.getLoginInfo().then(async res => {
        if (res.errcode === 0) {
          await getPageAuth()
          data.logined = true
        }
      })
    }

    onMounted(async () => {
      await getLoginInfo()
    })

    return {
      data,
      menuOptions: options
    }
  },
  methods: {
    onMenuUpdate (key) {
      this.$router.replace({
        path: key
      })
    }
  },
  mounted () {
    this.checkedMenu = this.$route.path.substr(1)
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
</style>
