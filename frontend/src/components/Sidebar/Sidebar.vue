<script setup lang="ts">
  import { useRouter } from "vue-router"
  const router = useRouter()
</script>

<template>
  <div class="w-1/9 bg-gray-700 text-left px-3 py-8">
    <div @click="router.push('/')" class="hover:bg-gray-800 w-full px-2 py-4">Home</div>
  </div>
</template>
