<template>
  <div>
    <h3>用户管理</h3>
    <el-table :data="list" border style="width: 100%">
      <el-table-column prop="id" label="ID" width="80" />
      <el-table-column prop="username" label="用户名" />
      <el-table-column prop="nickname" label="昵称" />
      <el-table-column prop="phone" label="手机号" />
      <el-table-column prop="role" label="角色" />
    </el-table>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { listUsers } from '@/api/user'

const list = ref([])

onMounted(async () => {
  const data = await listUsers({ page: 1, pageSize: 10 })
  list.value = data?.list || []
})
</script>
