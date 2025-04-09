<template>
  <el-aside width="80px">
    <div class="content">
      <IconBtn title="聊天" content="打开聊天列表" :icon="ChatDotRound" path="/chat" />
      <IconBtn title="创建" content="创建新的聊天" :icon="CirclePlus" path="/create" />
      <IconBtn title="发现" content="其它人的聊天" :icon="Compass" path="/discovery" />
      <IconBtn title="设置" content="管理偏好设置" :icon="Setting" path="/settings" />
      <ServerState style="position: fixed; bottom: 0px" />
    </div>
  </el-aside>
</template>

<script setup lang="ts">
import { ChatDotRound, Setting, CirclePlus, Compass } from '@element-plus/icons-vue'
import IconBtn from '@/components/units/IconBtn.vue'
import ServerState from '@/components/units/ServerState.vue'
</script>

<style scoped>
.content {
  height: 100vh;
  padding: 12px;
  color: rgba(255, 255, 255, 0.8);
  background-color: rgba(0, 0, 0, 0.3);
}
</style>
