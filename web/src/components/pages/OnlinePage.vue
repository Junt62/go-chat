<template>
  <div class="user-list">
    <div class="section-title">在线用户</div>
    <ul>
      <li v-for="user in onlineUsers" :key="user.id" class="user-item">
        {{ user.name }}
      </li>
    </ul>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import type { ShowUser } from '@/types/user'
import { useUserStore } from '@/stores/user'

const userStore = useUserStore()
const onlineUsers = computed(() => userStore.onlineUsers)

defineProps<{ tableData: ShowUser[] }>()
</script>

<style scoped>
.user-list {
  padding: 20px;
  color: white;
  background-color: rgba(0, 0, 0, 0);
  border-radius: 8px;
}

.section-title {
  font-weight: bold;
  margin-bottom: 10px;
  font-size: 1.2em;
}

.user-item {
  padding: 6px 0;
  border-bottom: 1px solid rgba(255, 255, 255, 0.2);
}

.placeholder {
  opacity: 0.6;
  font-style: italic;
}
</style>
