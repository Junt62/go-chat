<template>
  <div>
    <h2>Chat ID:{{ chatId }}</h2>
    <p>聊天窗口内容...</p>
  </div>
</template>

<script setup lang="ts">
defineProps<{ chatId: string }>()
</script>
