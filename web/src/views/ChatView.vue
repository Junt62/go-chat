<template>
  <el-button type="primary" @click="handleLogout">Logout</el-button>
  <el-container>
    <el-aside width="200px">
      <UserList :users="users" />
    </el-aside>
    <el-main>
      <MessageList :messages="messages" />
      <InputBox @send="handleSend" />
    </el-main>
  </el-container>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import UserList from '@/components/UserList.vue'
import MessageList from '@/components/MessageList.vue'
import InputBox from '@/components/InputBox.vue'
import type { User } from '@/types/user'
import type { Message } from '@/types/chat'
import { useAuth } from '@/composables/useAuth'

const { logout } = useAuth()

const users = ref<User[]>([
  { id: '1', name: '用户1', avatar: 'https://via.placeholder.com/40', email: 'user1@qq.com' },
])
const messages = ref<Message[]>([])

const handleSend = (msg: string) => {
  messages.value.push({
    id: Date.now().toString(),
    sender: '我',
    content: msg,
    timestamp: new Date().toLocaleTimeString(),
  })
}

const handleLogout = async () => {
  await logout()
}
</script>
