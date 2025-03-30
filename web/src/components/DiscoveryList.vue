<template>
  <el-aside width="240px">
    <div class="content">
      <div class="title">聊天</div>
      <MessageButton title="在线列表" content="查看在线用户" :icon="Switch" path="/chat/online" />
      <MessageButton title="我的好友" content="查看好友列表" :icon="User" path="/chat/friend" />
      <el-divider title="divider" />
      <MessageButton title="张三" content="查看好友列表" :icon="User" path="/chat/zhangsan" />
      <MessageButton title="李四" content="查看好友列表" :icon="User" path="/chat/lisi" />
      <MessageButton title="王五" content="查看好友列表" :icon="User" path="/chat/wangwu" />
    </div>
  </el-aside>
</template>

<script setup lang="ts">
import { User, Switch, Setting } from '@element-plus/icons-vue'
import MessageButton from '@/components/MsgBtn.vue'
</script>

<style scoped>
.content {
  height: 100vh;
  padding: 10px;
  overflow-y: auto;
  background-color: rgba(0, 0, 0, 0.2);
}
.title {
  margin-bottom: 12px;
  color: rgba(255, 255, 255, 0.8);
}
</style>
