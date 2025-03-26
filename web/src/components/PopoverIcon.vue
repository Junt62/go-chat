<template>
  <div
    class="popover-icon"
    :class="{ active: isActive || selectedPage === chatPath }"
    @mousedown="setActive"
    @mouseup="clearActive"
    @mouseleave="clearActive"
    @click="handleClick"
  >
    <el-popover
      class="box-item"
      :title="title"
      :content="content"
      placement="right"
      popper-class="custom-popover"
      :hide-after="0"
      :persistent="false"
    >
      <template #reference>
        <div class="rectangle">
          <el-icon :size="28">
            <component :is="icon" />
          </el-icon>
        </div>
      </template>
    </el-popover>
  </div>
</template>

<script setup lang="ts">
import { defineProps, inject, ref, type Ref } from 'vue'
import { useRouter } from 'vue-router'

const props = defineProps<{
  title: string
  content: string
  icon: any
  chatPath: string
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
  setSelectedPage(props.chatPath)
  router.push(props.chatPath)
}
</script>

<style scoped>
.popover-icon {
  display: flex;
  align-items: center;
  justify-items: center;
  width: 56px;
  height: 56px;
  border-radius: 24px;
  margin-bottom: 12px;
  background-color: rgba(255, 255, 255, 0.15);
  transition: background-color 0.2s;
  cursor: pointer;
}
.popover-icon.active {
  background-color: rgba(255, 255, 255, 0.5);
}
.popover-icon:not(.active) {
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
.rectangle {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 56px;
  height: 56px;
}
</style>
