<script setup lang="ts">
  import { ref, onMounted, watch } from "vue"

  import { useTasks } from "../api/task"
  import CreateTodoModal from "../components/TodoList/CreateTodoModal.vue"
  import TodoList from "../components/TodoList/TodoList.vue"

  // データ定義
  const { tasks, fetchTasks, addTask, updateTask, deleteTask } = useTasks()

  // --- モーダル関連のデータ定義 ---
  const showCreateModal = ref<boolean>(false)
  const searchWord = ref<string>("") // Todo検索ワード

  const openCreateModal = () => {
    showCreateModal.value = true
  }

  const closeCreateModal = () => {
    showCreateModal.value = false
  }

  // コンポーネントがマウントされたらタスクを取得
  onMounted(() => {
    fetchTasks(searchWord.value)
  })
  // リアルタイムで変数を監視
  watch(searchWord, () => {
    fetchTasks(searchWord.value)
  })
</script>

<template>
  <div class="mx-auto p-5 font-sans">
    <h1 class="text-7xl font-bold text-center mb-16 mx-20">Vue + Go Todo App</h1>

    <div class="flex gap-2.5 mb-5 items-center">
      <input
        v-model="searchWord"
        type="text"
        placeholder="タスクを検索..."
        class="grow p-2.5 size-11"
      />
      <button
        @click="openCreateModal"
        class="px-2.5 py-2 size-2/12 cursor-pointer bg-gray-200 text-black"
      >
        ＋追加
      </button>
    </div>

    <TodoList :tasks="tasks" @update-task="updateTask" @delete-task="deleteTask" />

    <CreateTodoModal
      v-if="showCreateModal"
      @on-close="closeCreateModal"
      @on-save="
        (newTask) => {
          addTask(newTask)
          closeCreateModal()
        }
      "
    />
  </div>
</template>
