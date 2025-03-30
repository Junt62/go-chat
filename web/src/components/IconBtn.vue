<template>
  <div class="icon-button" :class="{ active: isActive }" @click="handleClick">
    <el-popover :title="title" :content="content" placement="right" :hide-after="0" :persistent="false">
      <template #reference>
        <div class="icon">
          <el-icon :size="28">
            <component :is="icon" />
          </el-icon>
        </div>
      </template>
    </el-popover>
  </div>
</template>

<script setup lang="ts">
import { computed, defineProps } from 'vue'
import { useRoute, useRouter } from 'vue-router'

const props = defineProps<{
  title: string
  content: string
  icon: any
  path: string
}>()

const router = useRouter()
const route = useRoute()
const isActive = computed(() => {
  return route.path === props.path
})

const handleClick = () => {
  router.push(props.path)
}
</script>

<style scoped>
.icon-button {
  display: flex;
  align-items: center;
  justify-items: center;
  width: 56px;
  height: 56px;
  border-radius: 24px;
  margin-bottom: 12px;
  background-color: rgba(255, 255, 255, 0.1);
  transition: background-color 0.2s ease;
  cursor: pointer;
}
.icon-button:hover {
  background-color: rgba(255, 255, 255, 0.4);
}
.icon-button:active {
  background-color: rgba(255, 255, 255, 0.2);
}
.icon-button.active {
  background-color: rgba(255, 255, 255, 0.3);
}
.icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: 100%;
}
</style>
