<template>
  <el-container>
    <IconBar />

    <MsgBtnGroup :sections="sections" />

    <div class="content">创建一个新的聊天</div>
  </el-container>
</template>

<script setup lang="ts">
import IconBar from '@/components/IconBtnGroup.vue'
import MsgBtnGroup from '@/components/MsgBtnGroup.vue'
import { User, Switch, Setting } from '@element-plus/icons-vue'

let buttonIndex = 0

const sections = [
  {
    title: '聊天',
    buttons: [
      { title: '在线列表', icon: Switch },
      { title: '我的好友', icon: User },
    ],
  },
  {
    title: '进行中的聊天',
    buttons: [
      { title: '张三', icon: User },
      { title: '李四', icon: User },
      { title: '王五', icon: User },
    ],
  },
].map((section) => ({
  title: section.title,
  buttons: section.buttons.map((button) => ({
    ...button,
    index: buttonIndex++,
  })),
}))
</script>

<style scoped>
.el-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  width: 100vw;
  background: linear-gradient(135deg, #3a1c71, #d76d77, #c37745);
  position: fixed;
  top: 0;
  left: 0;
  color: rgba(255, 255, 255, 0.8);
}
.content {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  width: 100vw;
  background-color: rgba(0, 0, 0, 0.1);
  padding: 10px;
  font-size: 1.2rem;
}
</style>
