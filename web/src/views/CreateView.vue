<template>
  <div class="create-container">
    <el-row>
      <FunctionBar />

      <el-col style="width: 240px; flex: 0 0 240px">
        <div class="column-content content-2">
          <i class="el-icon-star"></i>
          <p>secend column</p>
        </div>
      </el-col>

      <el-col>
        <div class="column-content content-3">
          <i class="el-icon-star"></i>
          <p>third column</p>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import FunctionBar from '@/components/FunctionBar.vue'
</script>

<style scoped>
.create-container {
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
.el-row {
  width: 100%;
  flex-wrap: nowrap;
}
.column-content {
  height: 100vh;
}
.content-1 {
  padding: 12px;
  background-color: rgba(0, 0, 0, 0.2);
}
.content-2 {
  padding: 10px;
  background-color: rgba(0, 0, 0, 0.15);
}
.content-3 {
  padding: 10px;
  background-color: rgba(0, 0, 0, 0.1);
}
</style>
