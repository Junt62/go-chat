<template>
  <el-aside width="240px">
    <div class="content">
      <div v-for="section in sections" :key="section.title">
        <div v-if="section.title" class="title">{{ section.title }}</div>
        <MsgBtn
          v-for="button in section.buttons"
          :key="button.index"
          :title="button.title"
          :icon="button.icon"
          :is-active="activeButton === button.index"
          @click="handlerButtonClick(button)"
        />
        <el-divider title="divider" />
      </div>
    </div>
  </el-aside>
</template>

<script setup lang="ts">
import MsgBtn from '@/components/MsgBtn.vue'
import { ref } from 'vue'

interface Button {
  index: number
  title: string
  icon?: any
  handler?: Function
}

interface Section {
  title?: string
  buttons: Button[]
}

const props = defineProps<{
  sections: Section[]
}>()

const activeButton = ref<number>(-1)

const emit = defineEmits(['button-click'])

const handlerButtonClick = (button: Button) => {
  activeButton.value = button.index

  emit('button-click', button.title)

  if (button.handler) {
    button.handler()
  }
}
</script>

<style scoped>
.content {
  height: 100vh;
  padding: 10px;
  overflow-y: auto;
  background-color: rgba(0, 0, 0, 0.2);
}
.title {
  margin-bottom: 12px;
  color: rgba(255, 255, 255, 0.8);
}
</style>
