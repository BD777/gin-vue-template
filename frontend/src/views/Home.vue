<template>
  <div style="margin-top: 50px;">
    <n-empty :description="data.emptyHint">
    </n-empty>
  </div>
</template>

<script>
import { NEmpty } from 'naive-ui'
import { reactive } from 'vue'

export default {
  name: 'HomeView',
  components: {
    NEmpty
  },
  setup () {
    const data = reactive({
      emptyHint: 'Nothing here yet.'
    })

    return {
      data
    }
  }
}
</script>