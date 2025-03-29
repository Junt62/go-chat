<template>
  <el-col style="width: 80px; flex: 0 0 80px">
    <div class="column-content">
      <FunctionButton title="聊天" content="打开聊天列表" :icon="ChatDotRound" path="/chat" />
      <FunctionButton title="创建" content="创建新的群聊" :icon="CirclePlus" path="/create" />
      <FunctionButton title="发现" content="进行中的群聊" :icon="Compass" path="/discovery" />
      <FunctionButton title="设置" content="管理偏好设置" :icon="Setting" path="/settings" />
      <ServerState style="position: fixed; bottom: 0px" />
    </div>
  </el-col>
</template>

<script setup lang="ts">
import { ChatDotRound, Setting, CirclePlus, Compass } from '@element-plus/icons-vue'
import FunctionButton from '@/components/FunctionButton.vue'
import ServerState from '@/components/ServerState.vue'
</script>

<style scoped>
.column-content {
  height: 100vh;
  padding: 12px;
  background-color: rgba(0, 0, 0, 0.2);
}
</style>
