<template>
  <el-header class="top-bar">
    <div class="logo" @click="goToHome">Go-Chat</div>
    <div class="right-tools">
      <!-- <div @click="ping">{{ userStore.username }}</div> -->
      <el-button type="primary" round class="sub-btn" v-if="isLoggedIn()" @click="logout"
        >登出
      </el-button>
      <el-button type="primary" round class="sub-btn" v-else @click="login">登录 </el-button>
    </div>
  </el-header>
  <div class="app-container">
    <div class="intro-container">
      <div class="content">
        <h1>Go-Chat 在线聊天室</h1>
        <p>一个基于 Go + Vue3 + Element Plus 构建的实时聊天应用，提供高效、安全的即时通讯体验。</p>
        <el-row class="features" :gutter="20">
          <el-col v-for="(item, index) in features" :key="index" class="feature" :span="12">
            <e-icon :size="20"><component :is="item.icon" /></e-icon>
            {{ item.text }}
          </el-col>
        </el-row>
        <el-button type="primary" round class="main-btn" @click="join">立即体验</el-button>
        <el-button type="primary" round class="main-btn" @click="signUp">用户注册</el-button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useAuth } from '@/composables/useAuth'
import { useUserStore } from '@/stores/user'
import { Refresh, User, Check, Edit, Grid, UploadFilled } from '@element-plus/icons-vue'
import { defineComponent } from 'vue'
import { useRouter } from 'vue-router'

const userStore = useUserStore()
const router = useRouter()
const { ping, logout } = useAuth()

const features = [
  { icon: Refresh, text: '高性能' },
  { icon: User, text: '多用户支持' },
  { icon: Check, text: '高性能' },
  { icon: Edit, text: '简洁开发' },
  { icon: Grid, text: '组件丰富' },
  { icon: UploadFilled, text: '云端部署' },
]

const isLoggedIn = (): boolean => {
  return userStore.username !== ''
}
const goToHome = () => {
  router.push('/')
}
const join = () => {
  router.push('/chat')
}
const login = () => {
  router.push('/login')
}
const signUp = () => {
  router.push('/register')
}
</script>

<style scoped>
.intro-container {
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center;
  height: 100vh;
  width: 100vw;
  background: linear-gradient(135deg, #3a1c71, #d76d77, #c37745);
  color: white;
  position: fixed;
  top: 0;
  left: 0;
}

.top-bar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  color: white;
  position: fixed;
  width: 100%;
  top: 0;
  left: 50%;
  transform: translateX(-50%);
  z-index: 1000;
}
.features {
  display: flex;
  flex-wrap: wrap;
  padding-top: 60px;
  gap: 20px;
}
.feature {
  display: flex;
  align-items: center;
  gap: 10px;
  flex: 1 1 10%;
}
.logo {
  font-size: 1.5rem;
  font-weight: bold;
  cursor: pointer;
  transition: color 0.3s;
}
.right-tools {
  display: flex;
  align-items: center;
  gap: 15px;
}
.content {
  max-width: 600px;
  padding: 20px;
}

h1 {
  font-size: 2.5rem;
  margin-bottom: 10px;
}

p {
  font-size: 1.2rem;
  opacity: 0.9;
}

ul {
  list-style: none;
  padding: 0;
  font-size: 1.2rem;
  margin: 20px 0;
}

li {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 10px;
}
.main-btn {
  background: rgba(255, 255, 255, 0.2);
  border: none;
  font-size: 1.2rem;
  padding: 12px 24px;
  transition: all 0.3s;
  margin-top: 60px;
  margin-left: 40px;
  margin-right: 40px;
}
.main-btn:hover {
  background: rgba(255, 255, 255, 0.4);
}
.sub-btn {
  background: rgba(255, 255, 255, 0.2);
  border: none;
  font-size: 1rem;
  padding: 12px 24px;
  transition: all 0.3s;
}
.sub-btn:hover {
  background: rgba(255, 255, 255, 0.4);
}
</style>
