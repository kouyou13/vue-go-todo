<script setup lang="ts">
  import { ref, onMounted } from "vue"

  import { useCategories } from "../../api/category"
  import { emptyTask } from "../../api/task/util"
  import type { Task } from "../../types"
  import SelectCategoryModal from "../CategoryList/SelectCategoryModal.vue"

  const emit = defineEmits(["onClose", "onSave"])

  const editingTask = ref<Task>({ ...emptyTask })
  const { categories, fetchCategories } = useCategories()
  const showSelectCategoryModal = ref<boolean>(false)
  const selectedCategoryId = ref<string | null>(null)

  const openSelectCategoryModal = () => {
    selectedCategoryId.value = editingTask.value.categoryId
    showSelectCategoryModal.value = true
  }
  const closeSelectCategoryModal = () => {
    showSelectCategoryModal.value = false
  }
  onMounted(() => {
    fetchCategories()
  })
</script>

<template>
  <div class="modal-overlay" @click.self="emit('onClose')">
    <div class="modal-content">
      <h2>追加</h2>
      <p>タスク名</p>
      <input
        v-model="editingTask.title"
        type="text"
        class="modal-input"
        placeholder="タスク名を入力..."
      />
      <p>期限</p>
      <input v-model="editingTask.limitedAt" type="date" class="modal-date-input" />
      <p>カテゴリー</p>
      <div class="modal-category-input" @click="openSelectCategoryModal">
        <div v-if="editingTask.categoryId == null" class="category-tag">なし</div>
        <div v-else class="category-tag" style="background-color: blue">
          {{ categories?.find((c) => c.id === editingTask.categoryId)?.name }}
        </div>
      </div>
      <p>備考</p>
      <textarea v-model="editingTask.note" class="modal-textarea" />

      <div class="modal-actions">
        <button class="cancel-btn" @click="emit('onClose')">キャンセル</button>
        <button class="save-btn" @click="emit('onSave', editingTask)">保存</button>
      </div>
    </div>

    <SelectCategoryModal
      v-if="showSelectCategoryModal"
      :categories="categories ?? []"
      :editing-task="editingTask"
      @on-close="closeSelectCategoryModal"
      @on-save="
        () => {
          editingTask.categoryId = selectedCategoryId
          closeSelectCategoryModal()
        }
      "
    />
  </div>
</template>
