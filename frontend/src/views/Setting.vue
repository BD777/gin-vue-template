<template>
  <div>
    <n-tabs type="line" animated>
      <n-tab-pane name="admin" tab="Admin">
        <AdminSetting />
      </n-tab-pane>
    </n-tabs>
  </div>
</template>

<script>
import { NTabs, NTabPane } from 'naive-ui'
import { reactive } from 'vue'
import AdminSetting from '@/components/AdminSetting.vue'

export default {
  name: 'SettingView',
  components: {
    AdminSetting,
    NTabs,
    NTabPane
  },
  setup () {
    const data = reactive({
    })

    return {
      data
    }
  }
}
</script>
