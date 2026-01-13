<script setup lang="ts">
  import { type PropType, ref } from "vue"

  import type { Category } from "../../types"

  import EditCategoryModal from "./EditCategoryModal.vue"

  defineProps({
    categories: {
      type: Array as PropType<Category[]>,
      default: () => [],
    },
  })
  const emit = defineEmits(["updateCategory", "deleteCategory"])

  const showEditModal = ref<boolean>(false)
  const currentCategory = ref<Category | null>(null)
  const openEditModal = (category: Category) => {
    currentCategory.value = category
    showEditModal.value = true
  }
  const closeEditModal = () => {
    currentCategory.value = null
    showEditModal.value = false
  }
</script>

<template>
  <table class="w-full">
    <tr v-for="category in categories" :key="category.id" class="border-b">
      <td class="text-lg">
        {{ category.name }}
      </td>
      <td class="w-1/18 pl-2">
        <button
          class="text-center px-3 text-sm border-0 text-white bg-blue-700 h-10 w-full"
          @click="
            () => {
              openEditModal(category)
            }
          "
        >
          編集
        </button>
      </td>
      <td class="w-1/18 pl-3">
        <button
          class="text-center px-3 text-sm border-0 text-white bg-red-600 h-10 w-full my-4"
          @click="emit('deleteCategory', category.id)"
        >
          削除
        </button>
      </td>
      <EditCategoryModal
        v-if="showEditModal && currentCategory"
        :category="currentCategory"
        @on-close="closeEditModal"
        @on-save="
          (updated) => {
            emit('updateCategory', updated)
            closeEditModal()
          }
        "
      />
    </tr>
  </table>
</template>
