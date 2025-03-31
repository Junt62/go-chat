<template>
  <el-container>
    <IconBtnGroup />

    <MsgBtnGroup :sections="sections" @button-click="handleButtonClick" />

    <div class="content">
      <OnlinePage v-if="activeSection === 'online'" :tableData="tableData" />
      <FriendPage v-else-if="activeSection === 'friend'" />
      <div v-else>选择一个好友来开始聊天</div>
    </div>
  </el-container>
</template>

<script setup lang="ts">
import IconBtnGroup from '@/components/IconBtnGroup.vue'
import MsgBtnGroup from '@/components/MsgBtnGroup.vue'
import { User, Switch, Setting } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'
import OnlinePage from '@/components/OnlinePage.vue'
import FriendPage from '@/components/FriendPage.vue'
import { ref } from 'vue'
import type { ShowUser } from '@/types/user'

const tableData: ShowUser[] = [
  {
    avatar: '',
    name: '123',
  },
  {
    avatar: '',
    name: '222',
  },
  {
    avatar: '',
    name: '333',
  },
  {
    avatar: '',
    name: '444',
  },
  {
    avatar: '',
    name: '555',
  },
]

type SectionType = 'online' | 'friend' | 'chats' | null

const activeSection = ref<SectionType>(null)

const handleButtonClick = (buttonTitle: string) => {
  if (buttonTitle === '在线列表') {
    activeSection.value = 'online'
  } else if (buttonTitle === '我的好友') {
    activeSection.value = 'friend'
  } else {
    activeSection.value = 'chats'
  }
}

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
    index: buttonIndex++,
    ...button,
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
  /* padding: 10px; */
  font-size: 1.2rem;
}
</style>
