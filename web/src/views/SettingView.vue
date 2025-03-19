<script setup lang="ts">
import TopBar from '@/components/TopBar.vue'
</script>

<template>
  <div class="setting-container">
    <TopBar />
    <h1>This is an setting page</h1>
  </div>
</template>

<style>
.setting-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  width: 100vw;
  background: linear-gradient(135deg, #3a1c71, #d76d77, #c37745);
  position: fixed;
  top: 0;
  left: 0;
  color: white;
}
</style>
