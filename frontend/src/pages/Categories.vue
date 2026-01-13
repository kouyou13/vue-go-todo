<script setup lang="ts">
  import { ref, onMounted, watch } from "vue"

  import { useCategories } from "../api/category"
  import CategoryList from "../components/CategoryList/CategoryList.vue"
  import CreateCategoryModal from "../components/CategoryList/CreateCategoryModal.vue"

  // データ定義
  const { categories, fetchCategories, addCategory, updateCategory, deleteCategory } =
    useCategories()

  // --- モーダル関連のデータ定義 ---
  const showCreateModal = ref<boolean>(false)
  const searchWord = ref<string>("") // 検索ワード

  const openCreateModal = () => {
    showCreateModal.value = true
  }

  const closeCreateModal = () => {
    showCreateModal.value = false
  }

  // コンポーネントがマウントされたらタスクを取得
  onMounted(() => {
    fetchCategories(searchWord.value)
  })
  // リアルタイムで変数を監視
  watch(searchWord, () => {
    fetchCategories(searchWord.value)
  })
</script>

<template>
  <div class="mx-auto font-sans w-8/10 mt-2">
    <div class="flex gap-2.5 mb-5 items-center">
      <input
        v-model="searchWord"
        type="text"
        placeholder="カテゴリーを検索..."
        class="grow p-2.5 size-11 border-white border rounded-lg"
      />
      <button
        class="px-2.5 py-2 size-1/15 cursor-pointer bg-gray-200 text-black"
        @click="openCreateModal"
      >
        ＋追加
      </button>
    </div>

    <CategoryList
      :categories="categories"
      @update-category="updateCategory"
      @delete-category="deleteCategory"
    />

    <CreateCategoryModal
      v-if="showCreateModal"
      @on-close="closeCreateModal"
      @on-save="
        (newCategory) => {
          addCategory(newCategory)
          closeCreateModal()
        }
      "
    />
  </div>
</template>
