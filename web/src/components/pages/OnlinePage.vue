<template>
  <div class="online-page">
    <div class="title">在线用户：{{ onlineUsers.length }} 人</div>
    <ul>
      <div v-for="user in onlineUsers" :key="user.id" class="user-item">
        {{ user.name }}
      </div>
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
.online-page {
  width: 100%;
  height: 100%;
  padding: 10px;
  color: rgba(255, 255, 255, 0.8);
  /* background-color: rgba(0, 0, 0, 0.1); */
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
