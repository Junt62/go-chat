<template>
  <div class="register-container">
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
          <el-button type="primary" @click="handleRegister(ruleFormRef)">Sign Up</el-button>
          <el-button type="info" @click="handleToLogin">Have an account？Sign In</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { useUserStore } from '@/stores/user'
import { useRouter } from 'vue-router'
import { ElMessage, FormRules, FormInstance } from 'element-plus'

const userStore = useUserStore()

const router = useRouter()
interface RuleForm {
  username: string
  password: string
  email: string
}
const ruleFormRef = ref<FormInstance>()
const ruleForm = reactive<RuleForm>({
  username: '',
  password: '',
  email: '',
})
const rules = reactive<FormRules<RuleForm>>({
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

const handleRegister = async (formEl: FormInstance | undefined) => {
  if (!formEl) return
  await formEl.validate(async (valid) => {
    if (valid) {
      try {
        await userStore.register(ruleForm.username, ruleForm.password, ruleForm.email)
        ElMessage.success('Sign up successful')
        router.push('/login')
      } catch (error) {
        ElMessage.error('Sign up failed, please try again')
      }
    } else {
      ElMessage.error('Please check the form for errors')
    }
  })
}
const handleToLogin = async () => {
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
}
.register-card {
  width: auto;
  padding: 20px;
  text-align: center;
}
</style>
