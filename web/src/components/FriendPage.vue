<template>
  <div>
    <h3>我的好友</h3>
    <ul>
      <li v-for="friend in friendsList" :key="friend.id">{{ friend.name }}</li>
    </ul>
  </div>
</template>

<script setup lang="ts">
import { useUserStore } from '@/stores/user'
import { computed } from 'vue'

const userStore = useUserStore()
const friendsList = computed(() => userStore.friends)
</script>
