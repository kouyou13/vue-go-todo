<script setup lang="ts">
  import { useRouter } from "vue-router"
  const router = useRouter()
</script>

<template>
  <div class="mt-10 ml-8 mb-20">
    <h1 class="text-5xl font-bold text-left" @click="router.push('/')">Vue + Go Todo App</h1>
  </div>
</template>
