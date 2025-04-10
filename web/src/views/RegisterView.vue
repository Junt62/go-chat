<template>
  <div class="register-container">
    <TopBar />

    <el-card class="register-card">
      <template #header>
        <div class="card-header">
          {{ $t('SignUpTitle') }}
        </div>
      </template>

      <el-form ref="ruleFormRef" :model="ruleForm" :rules="rules" label-width="80px">
        <el-form-item :label="$t('Username')" prop="username">
          <el-input v-model="ruleForm.username" />
        </el-form-item>
        <el-form-item :label="$t('Password')" prop="password">
          <el-input v-model="ruleForm.password" type="password" show-password />
        </el-form-item>
        <el-form-item :label="$t('Email')" prop="email">
          <el-input v-model="ruleForm.email" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleRegister">{{ $t('SignUp') }}</el-button>
          <el-button type="info" @click="goToLogin">{{ $t('HaveAccountGoSignIn') }}</el-button>
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
import { useI18n } from 'vue-i18n'

const { t } = useI18n()

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
    { required: true, message: t('UsernameRequireMsg'), trigger: 'blur' },
    { min: 3, message: t('UsernameMinMsg'), trigger: 'blur' },
  ],
  password: [
    { required: true, message: t('PasswordRequireMsg'), trigger: 'blur' },
    { min: 6, message: t('PasswordMinMsg'), trigger: 'blur' },
  ],
  email: [
    { required: true, message: t('EmailRequireMsg'), trigger: 'blur' },
    { type: 'email', message: t('EmailMinMsg'), trigger: ['blur', 'change'] },
  ],
})

const handleRegister = async () => {
  if (!ruleFormRef.value) return
  try {
    await ruleFormRef.value.validate()
    await register(ruleForm)
  } catch (error) {
    ElMessage.error(t('RegisterError'))
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
