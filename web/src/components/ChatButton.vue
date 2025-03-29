<template>
  <div
    class="chat-button"
    :class="{ active: isActive || selectedPage === path }"
    @mousedown="setActive"
    @mouseup="clearActive"
    @mouseleave="clearActive"
    @click="handleClick"
  >
    <el-icon :size="28">
      <component :is="icon" />
    </el-icon>
    <span>{{ title }}</span>
  </div>
</template>

<script setup lang="ts">
import { defineProps, inject, ref, type Ref } from 'vue'
import { useRouter } from 'vue-router'

const props = defineProps<{
  title: string
  content: string
  icon: any
  path: string
}>()

const router = useRouter()
const isActive = ref(false)

const selectedPage = inject('selectedPage') as Ref<string>
const setSelectedPage = inject('setSelectedPage') as (path: string) => void

const setActive = () => {
  isActive.value = true
}
const clearActive = () => {
  isActive.value = false
}
const handleClick = () => {
  setSelectedPage(props.path)
  router.push(props.path)
}
</script>

<style scoped>
.chat-button {
  display: flex;
  align-items: center;
  justify-items: center;
  /* width: 20px; */
  height: 50px;
  border-radius: 24px;
  margin-bottom: 12px;
  background-color: rgba(255, 255, 255, 0.15);
  /* transition: background-color 0.2s; */
  cursor: pointer;
}
.chat-button.active {
  background-color: rgba(255, 255, 255, 0.5);
}
.chat-button:not(.active) {
  background-color: rgba(255, 255, 255, 0.15);
}
.custom-popover {
  background-color: transparent !important;
  box-shadow: none !important;
  border: none !important;
  color: white;
}
.custom-popover .el-popover__title {
  background-color: transparent !important;
  color: white;
}
</style>
