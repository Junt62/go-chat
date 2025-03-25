<template>
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
</template>

<script setup lang="ts">
import { defineProps } from 'vue'

defineProps<{
  title: string
  content: string
  icon: any
}>()
</script>

<style scoped>
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
/* .box-item {
  background-color: #ffffff44;
  color: #ffffff;
} */
.rectangle {
  width: 56px;
  height: 56px;
  background-color: rgba(255, 255, 255, 0.2);
  border-radius: 24px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 12px;
  cursor: pointer;
  transition: background 0.3s;
}
.rectangle:hover {
  background-color: rgba(255, 255, 255, 0.4);
}
</style>
