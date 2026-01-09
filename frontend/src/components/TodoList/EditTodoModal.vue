<script setup lang="ts">
  import { ref, type PropType, onMounted } from "vue"

  import { useCategories } from "../../api/category"
  import { emptyTask } from "../../api/task/util"
  import type { Task } from "../../types"
  import SelectCategoryModal from "../CategoryList/SelectCategoryModal.vue"

  const props = defineProps({
    task: {
      type: Object as PropType<Task>,
      default: () => null,
    },
  })
  const emit = defineEmits(["onClose", "onSave"])
  if (!props.task) {
    alert("タスクの取得に失敗しました")
  }
  const editingTask = ref<Task>(props.task ? { ...props.task } : emptyTask)
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
  const handleSelectCategory = (selectedCategoryId: string | null) => {
    editingTask.value.categoryId = selectedCategoryId
    closeSelectCategoryModal()
  }
  onMounted(() => {
    fetchCategories()
  })
</script>

<template>
  <div
    class="fixed top-0 left-0 w-full h-full bg-black/80 bg-op flex justify-center items-center z-10"
    @click.self="emit('onClose')"
  >
    <div class="bg-gray-950 p-6 rounded-2xl w-1/4 shadow z-10">
      <h2 class="text-3xl">編集</h2>
      <p class="text-left my-1">タスク名</p>
      <input
        v-model="editingTask.title"
        type="text"
        class="w-11/12 mb-5"
        placeholder="タスク名を入力..."
      />
      <p class="text-left my-1">期限</p>
      <input v-model="editingTask.limitedAt" type="date" class="w-11/12 mb-5" />
      <p class="text-left my-1">カテゴリー</p>
      <div class="w-11/12 h-9 mb-5 mx-auto hover:bg-gray-800" @click="openSelectCategoryModal">
        <div v-if="editingTask.categoryId == null" class="w-full mx-1.5">なし</div>
        <div v-else class="w-full mx-1.5" style="background-color: blue">
          {{ categories?.find((c) => c.id === editingTask.categoryId)?.name }}
        </div>
      </div>
      <p class="text-left my-1">備考</p>
      <textarea
        v-model="editingTask.note"
        placeholder="memo..."
        class="w-11/12 h-16 box-border mb-5"
      />

      <div class="flex justify-end gap-2.5 mt-2.5">
        <button class="bg-gray-100 border-0 text-black px-2 py-1.5" @click="emit('onClose')">
          キャンセル
        </button>
        <button
          class="bg-green-600 border-0 text-white px-2 py-1.5"
          @click="emit('onSave', editingTask)"
        >
          保存
        </button>
      </div>
    </div>

    <SelectCategoryModal
      v-if="showSelectCategoryModal"
      :categories="categories ?? []"
      :editing-task="editingTask"
      @on-close="closeSelectCategoryModal"
      @on-save="handleSelectCategory"
    />
  </div>
</template>
