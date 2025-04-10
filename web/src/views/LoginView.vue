<template>
  <div class="login-container">
    <TopBar />

    <el-card class="login-card">
      <template #header>
        <div class="card-header">
          {{ $t('SignInTitle') }}
        </div>
      </template>

      <el-form ref="loginDataRef" :model="loginData" :rules="rules" label-width="80px">
        <el-form-item :label="$t('Username')" prop="username">
          <el-input v-model="loginData.username" />
        </el-form-item>
        <el-form-item :label="$t('Password')" prop="password">
          <el-input v-model="loginData.password" type="password" show-password />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleLogin">{{ $t('SignIn') }}</el-button>
          <el-button type="info" @click="goToRegister">{{ $t('NoAccountGoSignUp') }}</el-button>
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
import TopBar from '@/components/TopBar.vue'
import { useI18n } from 'vue-i18n'

const { login } = useAuth()
const { t } = useI18n()

const router = useRouter()

const loginDataRef = ref<FormInstance>()
const loginData = reactive<LoginData>({
  username: '',
  password: '',
})
const rules = reactive<FormRules<LoginData>>({
  username: [
    { required: true, message: t('UsernameRequireMsg'), trigger: 'blur' },
    { min: 3, message: t('UsernameMinMsg'), trigger: 'blur' },
  ],
  password: [
    { required: true, message: t('PasswordRequireMsg'), trigger: 'blur' },
    { min: 6, message: t('PasswordMinMsg'), trigger: 'blur' },
  ],
})

const handleLogin = async () => {
  if (!loginDataRef.value) return
  try {
    await loginDataRef.value.validate()
    await login(loginData)
  } catch (error) {
    ElMessage.error(t('LoginError'))
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
  width: 100vw;
  background: linear-gradient(135deg, #3a1c71, #d76d77, #c37745);
  position: fixed;
  top: 0;
  left: 0;
}
.login-card {
  width: auto;
  padding: 20px;
  text-align: center;
}
</style>
