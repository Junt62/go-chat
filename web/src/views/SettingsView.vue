<template>
  <el-container>
    <IconBar />

    <MsgBtnGroup :sections="sections" />

    <div class="content">管理偏好设置</div>
  </el-container>
</template>

<script setup lang="ts">
import IconBar from '@/components/bars/SideBar.vue'
import MsgBtnGroup from '@/components/bars/ChatBar.vue'
import { useAuth } from '@/composables/useAuth'
import { useUserStore } from '@/stores/user'
import { ElMessage } from 'element-plus'

const { logout } = useAuth()
const useStore = useUserStore()

let buttonIndex = 0

const sections = [
  {
    title: '用户设置',
    buttons: [
      { title: '我的账号' },
      { title: '个人资料' },
      { title: '内容与社交' },
      { title: '数据与隐私' },
      { title: '家庭中心' },
      { title: '已授权的APP' },
      { title: '设备' },
      { title: '连接' },
      { title: '剪辑片段' },
    ],
  },
  {
    title: '账单设置',
    buttons: [{ title: 'Nitro' }, { title: '服务器助力' }, { title: '订阅' }, { title: '礼物库' }, { title: '账单' }],
  },
  {
    title: 'APP设置',
    buttons: [
      { title: '外观' },
      { title: '可访问性' },
      { title: '语音和视频' },
      { title: '聊天' },
      { title: '通知' },
      { title: '快捷键' },
      { title: '语言' },
      { title: '主播模式' },
      { title: '高级设置' },
    ],
  },
  {
    title: '活动设置',
    buttons: [{ title: '当前状态隐私' }],
  },
  {
    buttons: [{ title: '新鲜事儿' }, { title: '周边商品' }, { title: 'HypeSquad' }],
  },
  {
    buttons: [
      {
        title: '登出',
        handler: () => {
          logout()
        },
      },
      {
        title: '查看token是否过期',
        handler: () => {
          const time = useStore.tokenSetTime
            ? (() => {
                const newTime = new Date(useStore.tokenSetTime)
                newTime.setHours(newTime.getHours() + 20)
                return newTime.toLocaleDateString() + ' ' + newTime.toLocaleTimeString()
              })()
            : '时间未设置'

          if (useStore.isTokenExpired()) {
            ElMessage.error('已过期：' + time)
          } else {
            ElMessage.success('未过期：' + time)
          }
        },
      },
    ],
  },
].map((section) => ({
  title: section.title,
  buttons: section.buttons.map((button) => ({
    index: buttonIndex++,
    ...button,
  })),
}))
</script>

<style scoped>
.el-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  width: 100vw;
  background: linear-gradient(135deg, #3a1c71, #d76d77, #c37745);
  position: fixed;
  top: 0;
  left: 0;
  color: rgba(255, 255, 255, 0.8);
}
.content {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  width: 100vw;
  background-color: rgba(0, 0, 0, 0.1);
  padding: 10px;
  font-size: 1.2rem;
}
</style>
