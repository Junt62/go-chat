<template>
  <div class="register-container">
    <TopBar />

    <el-card class="register-card">
      <template #header>
        <div class="card-header">
          <span>Go-Chat: Sign Up</span>
        </div>
      </template>

      <el-form ref="ruleFormRef" :model="ruleForm" :rules="rules" label-width="80px">
        <el-form-item label="Username" prop="username">
          <el-input v-model="ruleForm.username" />
        </el-form-item>
        <el-form-item label="Password" prop="password">
          <el-input v-model="ruleForm.password" type="password" show-password />
        </el-form-item>
        <el-form-item label="Email" prop="email">
          <el-input v-model="ruleForm.email" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleRegister">Sign Up</el-button>
          <el-button type="info" @click="goToLogin">Have an account？Sign In</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import type { FormRules, FormInstance } from 'element-plus'
import type { RegisterData } from '@/api/types'
import { useAuth } from '@/composables/useAuth'
import TopBar from '@/components/TopBar.vue'

const router = useRouter()
const { register } = useAuth()
const ruleFormRef = ref<FormInstance>()

const ruleForm = reactive<RegisterData>({
  username: '',
  password: '',
  email: '',
})

const rules = reactive<FormRules<RegisterData>>({
  username: [
    { required: true, message: 'Please input username', trigger: 'blur' },
    { min: 3, message: 'Username must be at least 3 characters', trigger: 'blur' },
  ],
  password: [
    { required: true, message: 'Please input password', trigger: 'blur' },
    { min: 6, message: 'Password must be at least 6 characters', trigger: 'blur' },
  ],
  email: [
    { required: true, message: 'Please input email', trigger: 'blur' },
    { type: 'email', message: 'Please input a valid email', trigger: ['blur', 'change'] },
  ],
})

const handleRegister = async () => {
  if (!ruleFormRef.value) return
  try {
    await ruleFormRef.value.validate()
    await register(ruleForm)
  } catch (error) {
    ElMessage.error('Please correct the errors before submitting')
  }
}

const goToLogin = async () => {
  router.push('/login')
}
</script>

<style scoped>
.card-header {
  font-size: 24px;
}
.register-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  width: 100vw;
  background: linear-gradient(135deg, #3a1c71, #d76d77, #c37745);
  position: fixed;
  top: 0;
  left: 0;
}
.register-card {
  width: auto;
  padding: 20px;
  text-align: center;
}
</style>
