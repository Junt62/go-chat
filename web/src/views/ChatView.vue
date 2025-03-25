<template>
  <div class="chat-container">
    <el-row>
      <el-col style="width: 80px; flex: 0 0 80px">
        <div class="column-content content-1">
          <PopoverIcon title="最近" content="查看最近聊天" :icon="ChatDotRound" />
          <PopoverIcon title="好友" content="查看好友列表" :icon="User" />
          <PopoverIcon title="设置" content="管理偏好设置" :icon="Setting" />
        </div>
      </el-col>
      <el-col style="width: 240px; flex: 0 0 240px">
        <div class="column-content content-2">
          <i class="el-icon-star"></i>
          <p>secend column</p>
        </div>
      </el-col>
      <el-col>
        <div class="column-content content-3">
          <i class="el-icon-star"></i>
          <p>third column</p>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import UserList from '@/components/UserList.vue'
import MessageList from '@/components/MessageList.vue'
import InputBox from '@/components/InputBox.vue'
import { User, ChatDotRound, Setting } from '@element-plus/icons-vue'
import type { Message } from '@/types/chat'
import { useAuth } from '@/composables/useAuth'
import TopBar from '@/components/TopBar.vue'
import PopoverIcon from '@/components/PopoverIcon.vue'

const { logout } = useAuth()

// const users = ref<User[]>([
//   { id: '1', name: '用户1', avatar: 'https://via.placeholder.com/40', email: 'user1@qq.com' },
// ])
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

<style scoped>
.chat-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  width: 100vw;
  background: linear-gradient(135deg, #3a1c71, #d76d77, #c37745);
  position: fixed;
  top: 0;
  left: 0;
  color: white;
}
.el-row {
  width: 100%;
  flex-wrap: nowrap;
}
.column-content {
  height: 100vh;
}
.content-1 {
  padding: 12px;
  background-color: rgba(0, 0, 0, 0.2);
}
.content-2 {
  padding: 10px;
  background-color: rgba(0, 0, 0, 0.15);
}
.content-3 {
  padding: 10px;
  background-color: rgba(0, 0, 0, 0.1);
}
</style>
