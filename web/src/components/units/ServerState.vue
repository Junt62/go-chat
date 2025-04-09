<template>
  <div class="server-state">
    <el-popover
      title="服务器延迟"
      content="点击重新测试"
      placement="top"
      popper-class="custom-popover"
      :hide-after="0"
      :persistent="false"
    >
      <template #reference>
        <div @click="checkServerStatus">
          <div class="delay-icon" :class="['status-dot', getDotColorClass()]"></div>
          <span class="delay-message">{{ serverDelay }}ms</span>
        </div>
      </template>
    </el-popover>
  </div>
</template>

<script setup lang="ts">
import { defineProps, inject, onMounted, ref, type Ref } from 'vue'
import { useRouter } from 'vue-router'

enum ServerStatus {
  Good = '好',
  Average = '一般',
  Bad = '差',
  Failed = '失败',
}

var serverStatus: ServerStatus = ServerStatus.Failed
const serverDelay = ref<number>(0)

const checkServerStatus = async () => {
  const startTime = performance.now()
  try {
    const serverUrl = import.meta.env.VITE_SERVER_ADDRESS + '/api/ping'
    const response = await fetch(serverUrl)
    const endTime = performance.now()
    serverDelay.value = Math.floor(endTime - startTime)

    if (response.ok) {
      if (serverDelay.value < 50) {
        serverStatus = ServerStatus.Good
      } else if (serverDelay.value < 200) {
        serverStatus = ServerStatus.Average
      } else {
        serverStatus = ServerStatus.Bad
      }
    } else {
      serverStatus = ServerStatus.Failed
    }
  } catch (error) {
    serverStatus = ServerStatus.Failed
    serverDelay.value = 0
  }
}

onMounted(() => {
  checkServerStatus()
  const intervalId = setInterval(checkServerStatus, 60 * 1000)
  return () => clearInterval(intervalId)
})

const getDotColorClass = (): string => {
  if (serverStatus == ServerStatus.Good) {
    return 'green-dot'
  } else if (serverStatus == ServerStatus.Average) {
    return 'yellow-dot'
  } else if (serverStatus == ServerStatus.Bad) {
    return 'red-dot'
  }

  return 'gray-dot'
}
</script>

<style scoped>
.server-state {
  display: flex;
  align-items: center;
  justify-items: center;
  border-radius: 24px;
  /* background-color: rgba(255, 255, 255, 0.15); */
  cursor: pointer;
}
.delay-icon {
  margin-left: 4px;
  margin-right: 4px;
}
.delay-message {
  font-size: 0.8rem;
}
.status-dot {
  display: inline-block;
  width: 10px;
  height: 10px;
  border-radius: 50%;
  /* margin-right: 5px; */
}
.green-dot {
  background-color: hsl(120, 100%, 50%);
}
.yellow-dot {
  background-color: hsl(60, 100%, 50%);
}
.red-dot {
  background-color: hsl(0, 100%, 50%);
}
.gray-dot {
  background-color: hsl(0, 0%, 50%);
}
</style>
