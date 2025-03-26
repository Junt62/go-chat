<template>
  <PopoverIcon title="聊天" content="查看聊天列表" :icon="ChatDotRound" chatPath="/chat" />
  <PopoverIcon title="好友" content="查看好友列表" :icon="User" chatPath="/friends" />
  <PopoverIcon title="设置" content="管理偏好设置" :icon="Setting" chatPath="/settings" />
</template>

<script setup lang="ts">
import { User, ChatDotRound, Setting } from '@element-plus/icons-vue'
import PopoverIcon from '@/components/PopoverIcon.vue'
</script>
