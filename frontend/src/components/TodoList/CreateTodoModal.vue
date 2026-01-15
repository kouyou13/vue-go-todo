<script setup lang="ts">
  import { ref, onMounted } from "vue"

  import { useCategories } from "../../api/category"
  import { emptyTask } from "../../api/task/util"
  import type { Task } from "../../types"
  import SelectCategoryModal from "../CategoryList/SelectCategoryModal.vue"
  import { convertColorText } from "../CategoryList/utils/color"

  const emit = defineEmits(["onClose", "onSave"])

  const editingTask = ref<Task>({ ...emptyTask })
  const { categories, fetchCategories } = useCategories()
  const showSelectCategoryModal = ref<boolean>(false)

  const openSelectCategoryModal = () => {
    showSelectCategoryModal.value = true
  }
  const closeSelectCategoryModal = () => {
    showSelectCategoryModal.value = false
  }
  onMounted(() => {
    fetchCategories("")
  })
</script>

<template>
  <div
    class="fixed top-0 left-0 w-full h-full bg-black/60 bg-op flex justify-center items-center z-10"
    @click.self="emit('onClose')"
  >
    <div class="bg-gray-950 p-6 rounded-2xl w-1/4 shadow z-10">
      <h2 class="text-3xl">追加</h2>
      <p class="text-left my-1">タスク名</p>
      <input
        v-model="editingTask.title"
        type="text"
        class="w-11/12 mt-3 mb-6 border rounded-lg p-1"
        placeholder="タスク名を入力..."
      />
      <p class="text-left my-1">期限</p>
      <input
        v-model="editingTask.limitedAt"
        type="date"
        class="w-11/12 mt-3 mb-6 border rounded-lg p-1"
      />
      <p class="text-left my-1">カテゴリー</p>
      <div class="w-11/12 h-9 mt-3 mb-6 mx-auto" @click="openSelectCategoryModal">
        <div
          v-if="editingTask.categoryId == null"
          class="w-full py-1.5 rounded-lg hover:bg-gray-800"
        >
          なし
        </div>
        <div
          v-else
          class="w-full py-1.5 rounded-lg"
          :class="
            convertColorText(
              categories?.find((c) => c.id === editingTask.categoryId)?.color ?? 'blue',
            )
          "
        >
          {{ categories?.find((c) => c.id === editingTask.categoryId)?.name }}
        </div>
      </div>
      <p class="text-left my-1">備考</p>
      <textarea
        v-model="editingTask.note"
        placeholder="memo..."
        class="w-11/12 h-16 box-border mt-3 mb-6 border rounded-lg p-1"
      />

      <div class="flex justify-end gap-2.5 mt-2.5">
        <button class="bg-gray-100 border-0 text-black px-2 py-1.5" @click="emit('onClose')">
          キャンセル
        </button>
        <button
          class="bg-green-600 border-0 text-white px-2 py-1.5"
          @click="emit('onSave', editingTask)"
        >
          追加
        </button>
      </div>
    </div>

    <SelectCategoryModal
      v-if="showSelectCategoryModal"
      :categories="categories ?? []"
      :editing-task="editingTask"
      @on-close="closeSelectCategoryModal"
    />
  </div>
</template>
