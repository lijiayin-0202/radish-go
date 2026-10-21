<template>
  <div>
    <h3>首页概览</h3>
    <el-row :gutter="16">
      <el-col :span="8">
        <el-card shadow="hover">
          <template #header>当前用户</template>
          <div>{{ userStore.userInfo?.username || '-' }}（{{ userStore.userInfo?.role || '-' }}）</div>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card shadow="hover">
          <template #header>后端服务</template>
          <div>gin-vue-admin :8888</div>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card shadow="hover">
          <template #header>微服务</template>
          <div>taxi :8000(HTTP) / :9000(gRPC)</div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import { useUserStore } from '@/store'

const userStore = useUserStore()
</script>
