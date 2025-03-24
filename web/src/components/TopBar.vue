<!-- src/components/TopBar.vue -->
<template>
  <el-header class="top-bar">
    <div class="logo" @click="goToHome">Go-Chat</div>
    <div class="menu">
      <div class="menu-item" @click="goToChat">聊天</div>
      <div class="menu-item" @click="goToSetting">设置</div>
      <div class="menu-item" @click="goToAbout">关于</div>
    </div>
    <!-- <el-menu mode="horizontal" :ellipsis="false" class="menu">
      <el-menu-item index="1">首页</el-menu-item>
      <el-menu-item index="2">消息</el-menu-item>
      <el-menu-item index="3">好友</el-menu-item>
      <el-menu-item index="4">设置</el-menu-item>
    </el-menu> -->
    <div class="right-tools">
      <p>{{ userStore.username }}</p>
      <!-- <el-input v-model="search" placeholder="搜索..." class="search-box">
        <template #prefix
          ><el-icon><Search /></el-icon
        ></template>
      </el-input> -->
      <el-switch v-model="darkMode" class="dark-mode-switch" @change="toggleDarkMode">
        <template #active-action
          ><el-icon><Moon /></el-icon
        ></template>
        <template #inactive-action
          ><el-icon><Sunny /></el-icon
        ></template>
      </el-switch>
    </div>
  </el-header>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { Search, Moon, Sunny } from '@element-plus/icons-vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'

const search = ref<string>('')
const darkMode = ref<boolean>(false)
const router = useRouter()
const userStore = useUserStore()

const toggleDarkMode = (): void => {
  darkMode.value = !darkMode.value
  document.body.classList.toggle('dark-mode', darkMode.value)
}

const goToHome = () => {
  router.push('/')
}
const goToChat = () => {
  router.push('/chat')
}
const goToAbout = () => {
  router.push('/about')
}
const goToSetting = () => {
  router.push('/setting')
}
</script>

<style scoped>
.top-bar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 10px 20px;
  /* background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  -webkit-backdrop-filter: blur(10px);
  border-bottom: 1px solid rgba(255, 255, 255, 0.2); */
  color: white;
  position: fixed;
  width: 1024px;
  top: 0;
  left: 50%;
  transform: translateX(-50%);
  z-index: 1000;
}

.el-menu {
  background: transparent !important;
  border: none;
  /* box-shadow: 0 4px 6px rgba(0, 0, 0, 0.5); */
}

.el-menu-item {
  color: white !important;
}

.el-menu-item:hover {
  background-color: rgba(255, 255, 255, 0.1) !important; /* 更改为你喜欢的颜色 */
}

.logo {
  font-size: 1.5rem;
  font-weight: bold;
  cursor: pointer;
  transition: color 0.3s;
}

.logo:hover {
  color: rgba(255, 255, 255, 0.8);
}

.menu {
  display: flex;
  gap: 20px;
  justify-content: flex-start;
  flex: 1;
  padding-left: 36px;
  padding-top: 4px;
}

.menu-item {
  font-size: 0.9rem;
  font-weight: 600;
  cursor: pointer;
  padding: 12px;
  transition:
    color 0.3s,
    transform 0.3s;
  color: rgba(255, 255, 255, 0.8);
}

.menu-item:hover {
  color: rgba(255, 255, 255, 0.6); /* hover 状态下变色 */
}

.right-tools {
  display: flex;
  align-items: center;
  gap: 15px;
}

.search-box {
  width: 180px;
}

.dark-mode-switch {
  --el-switch-on-color: #222;
  --el-switch-off-color: #f7c948;
}
</style>
