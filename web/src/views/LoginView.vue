<template>
  <div class="login-container">
    <el-card class="" login-card>
      <h2>go-chat 在线聊天室</h2>
      <el-form label-width="80px">
        <el-form-item label="用户名">
          <el-input v-model="username" />
        </el-form-item>
        <el-form-item label="密码">
          <el-input v-model="password" type="password" show-password />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleLogin">登录</el-button>
          <el-button type="primary" @click="handleToRegister">没有账号？去注册</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useUserStore } from '@/stores/user'
import { useRouter } from 'vue-router'
import { ElMessage, ElSubMenu } from 'element-plus'

const username = ref('')
const password = ref('')
const userStore = useUserStore()
const router = useRouter()

const handleLogin = async () => {
  try {
    await userStore.login(username.value, password.value)
    ElMessage.success('登录成功')
    router.push('/chat')
  } catch (error) {
    ElMessage.error('登录失败，请检查用户名和密码')
  }
}
const handleToRegister = async () => {
  router.push('/register')
}
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: flex;
  align-items: center;
  height: 100vh;
}

.login-card {
  width: 350px;
  padding: 20px;
  text-align: center;
}
</style>
