<template>
  <div class="register-container">
    <el-card class="" register-card>
      <h2>go-chat 在线聊天室</h2>
      <h4>注册</h4>
      <el-form label-width="80px">
        <el-form-item label="用户名">
          <el-input v-model="username" />
        </el-form-item>
        <el-form-item label="密码">
          <el-input v-model="password" type="password" show-password />
        </el-form-item>
        <el-form-item label="邮箱">
          <el-input v-model="email" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleRegister">注册</el-button>
          <el-button type="primary" @click="handleToLogin">已有账号？去登录</el-button>
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
const email = ref('')
const userStore = useUserStore()
const router = useRouter()

const handleRegister = async () => {
  try {
    await userStore.register(username.value, password.value, email.value)
    ElMessage.success('注册成功')
    router.push('/login')
  } catch (error) {
    ElMessage.error('注册失败')
  }
}
const handleToLogin = async () => {
  router.push('/login')
}
</script>

<style scoped>
.register-container {
  display: flex;
  justify-content: flex;
  align-items: center;
  height: 100vh;
}

.register-card {
  width: 350px;
  padding: 20px;
  text-align: center;
}
</style>
