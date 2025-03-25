<template>
  <div class="about-container">
    <h1>This is an about page</h1>
  </div>
</template>

<script setup lang="ts"></script>

<style>
@media (min-width: 1024px) {
  .about-container {
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
}
</style>
