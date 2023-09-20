<template>
  <div>
    <n-data-table
      :columns="data.users.columns"
      :data="data.users.data"
      :pagination="data.users.pagination"
      :loading="data.users.loading"
    />
  </div>
</template>

<script>
import $backend from '../backend'
import { useMessage, NDataTable } from 'naive-ui'
import { onMounted, reactive } from 'vue'

export default {
  name: 'AdminSetting',
  components: {
    NDataTable
  },
  setup () {
    const message = useMessage()

    const data = reactive({
      users: {
        loading: false,
        data: [],
        columns: [
          { title: 'ID', key: 'id' },
          { title: 'UserName', key: 'username' },
          { title: 'Role', key: 'role' }
        ],
        pagination: {
          page: 1,
          pageSize: 20,
          total: 0
        }
      }
    })

    async function listAdmins() {
      data.users.loading = true
      return $backend.listAdmins(data.users.pagination).then(res => {
        console.log('listAdmins resp', res)
        if (res.errcode === 0) {
          data.users.data = res.data.list
          data.users.pagination.total = res.data.pagination.total
        } else {
          message.error(res.errmsg)
        }
      }).then(() => {
        data.users.loading = false
      })
    }

    onMounted(() => {
      listAdmins()
    })

    return {
      data
    }
  }
}
</script>

