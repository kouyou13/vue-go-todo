<script setup lang="ts">
  import Header from "./components/Header/Header.vue"
  import Sidebar from "./components/Sidebar/Sidebar.vue"
</script>

<template>
  <div class="w-full h-full flex">
    <Sidebar />
    <div class="w-full h-full">
      <Header />
      <router-view />
    </div>
  </div>
</template>
