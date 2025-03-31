<template>
  <div>
    <h3>在线用户</h3>
    <ul>
      <li v-for="user in onlineUsers" :key="user.id">{{ user.name }}</li>
    </ul>
  </div>
</template>

<script setup lang="ts">
import { useUserStore } from '@/stores/user'
import { computed } from 'vue'

const userStore = useUserStore()
const onlineUsers = computed(() => userStore.onlineUsers)
</script>
