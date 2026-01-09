<script setup lang="ts">
  import Sidebar from "./components/Sidebar/Sidebar.vue"
</script>

<template>
  <div class="w-full h-full flex">
    <Sidebar />
    <router-view />
  </div>
</template>
