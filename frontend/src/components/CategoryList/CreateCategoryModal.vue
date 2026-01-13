<script setup lang="ts">
  import { ref } from "vue"

  import { emptyCategory } from "../../api/category/util"
  import type { Category } from "../../types"

  const emit = defineEmits(["onClose", "onSave"])

  const editingCategory = ref<Category>({ ...emptyCategory })
</script>

<template>
  <div
    class="fixed top-0 left-0 w-full h-full bg-black/60 bg-op flex justify-center items-center z-10"
    @click.self="emit('onClose')"
  >
    <div class="bg-gray-950 p-6 rounded-2xl w-1/4 shadow z-10">
      <h2 class="text-3xl">追加</h2>
      <p class="text-left my-1">カテゴリー名</p>
      <input
        v-model="editingCategory.name"
        type="text"
        class="w-11/12 my-3 border rounded-lg p-1"
        placeholder="カテゴリー名を入力..."
      />

      <div class="flex justify-end gap-2.5 mt-2.5">
        <button class="bg-gray-100 border-0 text-black px-2 py-1.5" @click="emit('onClose')">
          キャンセル
        </button>
        <button
          class="bg-green-600 border-0 text-white px-2 py-1.5"
          @click="emit('onSave', editingCategory)"
        >
          追加
        </button>
      </div>
    </div>
  </div>
</template>
