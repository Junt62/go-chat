<template>
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
        <el-button type="primary" round class="join-btn" @click="join">立即体验</el-button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { Refresh, User, Check, Edit, Grid, UploadFilled } from '@element-plus/icons-vue'
import { useRouter } from 'vue-router'

const router = useRouter()

const features = [
  { icon: Refresh, text: '高性能' },
  { icon: User, text: '多用户支持' },
  { icon: Check, text: '高性能' },
  { icon: Edit, text: '简洁开发' },
  { icon: Grid, text: '组件丰富' },
  { icon: UploadFilled, text: '云端部署' },
]

const join = () => {
  router.push('/chat')
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
.content {
  max-width: 600px;
  padding: 20px;
}
.join-btn {
  background: rgba(255, 255, 255, 0.2);
  border: none;
  font-size: 1.2rem;
  padding: 12px 24px;
  margin-top: 48px;
  transition: all 0.3s;
}
.join-btn:hover {
  background: rgba(255, 255, 255, 0.4);
}
</style>
