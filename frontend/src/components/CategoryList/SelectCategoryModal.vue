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
  <div class="modal-overlay" @click.self="emit('onClose')">
    <div class="modal-content">
      <h2>カテゴリー選択</h2>
      <div style="margin-bottom: 5px">
        <label>
          <input v-model="selectedCategoryId" type="radio" name="category" :value="null" />
          <span class="category-tag">なし</span>
        </label>
      </div>
      <div v-for="category in categories" :key="category.id" style="margin-bottom: 5px">
        <label>
          <input v-model="selectedCategoryId" type="radio" name="category" :value="category.id" />
          <span class="category-tag" style="background-color: blue">
            {{ category.name }}
          </span>
        </label>
      </div>
      <div class="modal-actions">
        <button class="cancel-btn" @click="emit('onClose')">キャンセル</button>
        <button class="save-btn" @click="emit('onSave', selectedCategoryId)">保存</button>
      </div>
    </div>
  </div>
</template>
