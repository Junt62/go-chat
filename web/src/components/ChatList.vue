<template>
  <el-aside width="240px">
    <div class="content">
      <div class="title">聊天</div>
      <MessageButton title="在线列表" :icon="Switch" :is-active="activeButton === 0" @click="setActiveButton(0)" />
      <MessageButton title="我的好友" :icon="User" :is-active="activeButton === 1" @click="setActiveButton(1)" />

      <el-divider title="divider" />

      <div class="title">进行中的聊天</div>
      <MessageButton title="张三" :icon="User" :is-active="activeButton === 2" @click="setActiveButton(2)" />
      <MessageButton title="李四" :icon="User" :is-active="activeButton === 3" @click="setActiveButton(3)" />
      <MessageButton title="王五" :icon="User" :is-active="activeButton === 4" @click="setActiveButton(4)" />
    </div>
  </el-aside>
</template>

<script setup lang="ts">
import { User, Switch, Setting } from '@element-plus/icons-vue'
import MessageButton from '@/components/MsgBtn.vue'
import { ref } from 'vue'

const activeButton = ref<number>(-1)

const setActiveButton = (index: number) => {
  activeButton.value = index
}
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
