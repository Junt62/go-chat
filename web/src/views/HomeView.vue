<template>
  <div class="app-container">
    <TopBar />

    <div class="intro-container">
      <div class="content">
        <h1>Go-Chat 在线聊天室</h1>
        <p>一个基于 Go + Vue3 + Element Plus 构建的实时聊天应用，提供高效、安全的即时通讯体验。</p>
        <ul>
          <li>
            <el-icon><ChatDotRound /></el-icon> 实时消息推送
          </li>
          <li>
            <el-icon><User /></el-icon> 多用户在线
          </li>
          <li>
            <el-icon><Check /></el-icon> 端到端加密
          </li>
        </ul>
        <el-button type="primary" round class="join-btn">立即体验</el-button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import TopBar from '@/components/TopBar.vue'
import { ChatDotRound, User, Check } from '@element-plus/icons-vue'
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
  font-family: Arial, sans-serif;
  position: fixed;
  top: 0;
  left: 0;
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

.join-btn {
  background: rgba(255, 255, 255, 0.2);
  border: none;
  font-size: 1.2rem;
  padding: 12px 24px;
  transition: all 0.3s;
}

.join-btn:hover {
  background: rgba(255, 255, 255, 0.4);
}
</style>
