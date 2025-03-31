<template>
  <el-table :data="tableData" :default-sort="{ prop: 'date', order: 'descending' }">
    <el-table-column prop="avatar" label="头像" width="60" />
    <el-table-column prop="name" label="名称" width="120" />
  </el-table>
</template>

<script setup lang="ts">
import { useUserStore } from '@/stores/user'
import { computed } from 'vue'
import type { ShowUser } from '@/types/user'

const userStore = useUserStore()
const onlineUsers = computed(() => userStore.onlineUsers)

defineProps<{ tableData: ShowUser[] }>()
</script>

<style scoped>
.el-table {
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0);
  /* color: rgba(255, 255, 255, 0.8); */
}
.el-table th {
  background-color: rgba(255, 255, 255, 0.1); /* 可选：让表头半透明 */
}
.el-table tr {
  background-color: rgba(255, 255, 255, 0.1); /* 可选：让表头半透明 */
}
</style>
