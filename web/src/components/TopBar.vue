<!-- src/components/TopBar.vue -->
<template>
  <el-header class="top-bar">
    <div class="logo" @click="goHome">Go-Chat</div>
    <div class="right-tools">
      <el-button type="primary" round class="main-btn" @click="join">立即体验</el-button>
    </div>
  </el-header>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'

const router = useRouter()

const goHome = () => {
  router.push('/')
}
const join = () => {
  router.push('/chat')
}
</script>

<style scoped>
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
.main-btn {
  background: rgba(255, 255, 255, 0.2);
  border: none;
  font-size: 1rem;
  padding: 12px 24px;
  transition: all 0.3s;
}
.main-btn:hover {
  background: rgba(255, 255, 255, 0.4);
}
</style>
