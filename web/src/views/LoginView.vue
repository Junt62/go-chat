<template>
  <div class="login-container">
    <el-card class="login-card">
      <template #header>
        <div class="card-header">
          <span>Go-Chat: Sign In</span>
        </div>
      </template>

      <el-form ref="loginDataRef" :model="loginData" :rules="rules" label-width="80px">
        <el-form-item label="Username" prop="username">
          <el-input v-model="loginData.username" />
        </el-form-item>
        <el-form-item label="Password" prop="password">
          <el-input v-model="loginData.password" type="password" show-password />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleLogin">Sign In</el-button>
          <el-button type="info" @click="goToRegister">No account？Sign Up</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { useUserStore } from '@/stores/user'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import type { FormRules, FormInstance } from 'element-plus'
import type { LoginData } from '@/api/types'
import { useAuth } from '@/composables/useAuth'

const { login } = useAuth()

const router = useRouter()

const loginDataRef = ref<FormInstance>()
const loginData = reactive<LoginData>({
  username: '',
  password: '',
})
const rules = reactive<FormRules<LoginData>>({
  username: [
    { required: true, message: 'Please enter your username', trigger: 'blur' },
    { min: 3, message: 'Username must be at least 3 characters', trigger: 'blur' },
  ],
  password: [
    { required: true, message: 'Please enter your password', trigger: 'blur' },
    { min: 6, message: 'Password must be at least 6 characters', trigger: 'blur' },
  ],
})

const handleLogin = async () => {
  if (!loginDataRef.value) return
  try {
    await loginDataRef.value.validate()
    await login(loginData)
  } catch (error) {
    ElMessage.error('Please correct the errors before submitting')
  }
}
const goToRegister = async () => {
  router.push('/register')
}
</script>

<style scoped>
.card-header {
  font-size: 24px;
  text-align: center;
}
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}
.login-card {
  width: auto;
  padding: 20px;
  text-align: center;
}
</style>
