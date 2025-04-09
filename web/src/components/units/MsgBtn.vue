<template>
  <div class="message-button" :class="{ active: isActive }">
    <component class="icon" :is="icon" />
    <div class="content">{{ title }}</div>
  </div>
</template>

<script setup lang="ts">
import { defineProps } from 'vue'

defineProps<{
  title: string
  icon: any
  isActive: boolean
}>()
</script>

<style scoped>
.message-button {
  display: flex;
  align-items: center;
  justify-items: center;
  border-radius: 12px;
  margin-bottom: 12px;
  padding: 4px;
  color: rgba(255, 255, 255, 0.8);
  transition: background-color 0.2s ease;
  cursor: pointer;
}
.message-button:hover {
  background-color: rgba(255, 255, 255, 0.3);
}
.message-button:active {
  background-color: rgba(255, 255, 255, 0.1);
}
.message-button.active {
  background-color: rgba(255, 255, 255, 0.2);
}
.icon {
  width: 2rem;
  height: 2rem;
  color: rgba(255, 255, 255, 0.8);
  background-color: rgba(255, 255, 255, 0.2);
  border-radius: 50%;
  padding: 6px;
}
.content {
  margin-left: 12px;
}
</style>
