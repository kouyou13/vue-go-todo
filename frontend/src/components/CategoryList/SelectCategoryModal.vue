<script setup lang="ts">
  import { type PropType, ref } from "vue"

  import type { Task, Category } from "../../types"

  const props = defineProps({
    categories: { type: Array as PropType<Category[]>, default: () => [] },
    editingTask: { type: Object as PropType<Task>, default: () => null },
  })
  const emit = defineEmits(["onClose", "onSave"])
  const selectedCategoryId = ref<string | null>(props.editingTask?.categoryId || null)
</script>

<template>
  <div
    class="fixed top-0 left-0 w-full h-full bg-black/60 bg-op flex justify-center items-center z-20"
    @click.self="emit('onClose')"
  >
    <div class="bg-gray-950 p-6 rounded-2xl w-1/5 shadow z-20">
      <h2 class="text-3xl">カテゴリー選択</h2>
      <div class="my-2">
        <label>
          <input v-model="selectedCategoryId" type="radio" name="category" :value="null" />
          <span class="max-w-9/12 px-2 rounded-md cursor-pointer">なし</span>
        </label>
      </div>
      <div v-for="category in categories" :key="category.id" class="my-4">
        <label>
          <input
            v-model="selectedCategoryId"
            type="radio"
            name="category"
            :value="category.id"
            class="mx-2"
          />
          <span class="max-w-9/12 px-2 py-1 rounded-md cursor-pointer bg-blue-600">
            {{ category.name }}
          </span>
        </label>
      </div>
      <div class="flex justify-end gap-2.5 mt-2.5">
        <button class="bg-gray-100 border-0 text-black px-2 py-1.5" @click="emit('onClose')">
          キャンセル
        </button>
        <button
          class="bg-green-600 border-0 text-white px-2 py-1.5"
          @click="emit('onSave', selectedCategoryId)"
        >
          保存
        </button>
      </div>
    </div>
  </div>
</template>
