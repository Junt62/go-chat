<script setup lang="ts">
import { defineProps } from 'vue'
import type { User } from '@/types/user'

defineProps<{ users: User[] }>()
</script>

<template>
  <el-card>
    <el-avatar v-for="user in users" :key="user.id" :src="user.avatar" class="user-avatar" />
  </el-card>
</template>

<style scoped>
.user-avatar {
  margin-right: 10px;
}
</style>
