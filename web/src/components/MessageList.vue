<script setup lang="ts">
import { defineProps } from 'vue'
import type { Message } from '@/types/chat'

defineProps<{
  messages: Message[]
}>()
</script>

<template>
  <el-scrollbar height="400px">
    <el-card v-for="msg in messages" :key="msg.id">
      <span>{{ msg.sender }}:{{ msg.content }}</span>
      <span style="float: right; font-size: 12px; color: gray">{{ msg.timestamp }}</span>
    </el-card>
  </el-scrollbar>
</template>
