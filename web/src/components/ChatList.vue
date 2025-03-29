<template>
  <el-col style="width: 240px; flex: 0 0 240px">
    <div class="column-content">
      <span class="title">聊天</span>
      <el-divider title="divider" />
      <ChatButton title="在线列表" content="查看在线用户" :icon="Switch" path="/chat/online" />
      <ChatButton title="我的好友" content="查看好友列表" :icon="User" path="/chat/friend" />
      <el-divider title="divider" />
      <ChatButton title="张三" content="查看好友列表" :icon="User" path="/chat/zhangsan" />
      <ChatButton title="李四" content="查看好友列表" :icon="User" path="/chat/lisi" />
      <ChatButton title="王五" content="查看好友列表" :icon="User" path="/chat/wangwu" />
    </div>
  </el-col>
</template>

<script setup lang="ts">
import { User, Switch, Setting } from '@element-plus/icons-vue'
import ChatButton from '@/components/ChatButton.vue'
</script>

<style scoped>
.column-content {
  height: 100vh;
  padding: 10px;
  background-color: rgba(0, 0, 0, 0.15);
}
.title {
  font-size: 1.5rem;
}
.title2 {
  font-size: 1rem;
}
.el-divider {
  margin-top: 0px;
  margin-bottom: 12px;
}
</style>
