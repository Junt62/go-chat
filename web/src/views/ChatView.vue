<template>
  <div class="chat-container">
    <el-row>
      <el-col style="width: 80px; flex: 0 0 80px">
        <div class="column-content content-1">
          <PopoverIcon title="私信" content="打开私信列表" :icon="ChatDotRound" chatPath="/chat" />
          <PopoverIcon title="创建" content="创建新的群聊" :icon="CirclePlus" chatPath="/newchat" />
          <PopoverIcon title="发现" content="进行中的群聊" :icon="Compass" chatPath="/discovery" />
          <PopoverIcon title="设置" content="管理偏好设置" :icon="Setting" chatPath="/settings" />
          <PopoverIcon
            @click="logout"
            title="退出"
            content="退出账户"
            :icon="SwitchButton"
            chatPath="/"
          />
          <ServerState style="position: fixed; bottom: 0px" />
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
import {
  User,
  ChatDotRound,
  Setting,
  SwitchButton,
  Plus,
  CirclePlus,
  Compass,
} from '@element-plus/icons-vue'
import type { Message } from '@/types/chat'
import { useAuth } from '@/composables/useAuth'
import TopBar from '@/components/TopBar.vue'
import PopoverIcon from '@/components/PopoverIcon.vue'
import FunctionBar from '@/components/FunctionBar.vue'
import ServerState from '@/components/ServerState.vue'

const { logout } = useAuth()

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
