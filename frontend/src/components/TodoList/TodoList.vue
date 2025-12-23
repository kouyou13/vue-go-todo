<script setup lang="ts">
  import { type PropType, ref } from "vue"

  import EditTodoModal from "../../components/TodoList/EditTodoModal.vue"
  import type { Task } from "../../types"

  defineProps({
    tasks: {
      type: Array as PropType<Task[]>,
      default: () => [],
    },
  })
  const emit = defineEmits(["updateTask", "deleteTask"])

  const showEditModal = ref<boolean>(false)
  const currentTask = ref<Task | null>(null)
  const openEditModal = (task: Task) => {
    currentTask.value = task
    showEditModal.value = true
  }
  const closeEditModal = () => {
    currentTask.value = null
    showEditModal.value = false
  }
</script>

<template>
  <ul>
    <li v-for="task in tasks" :key="task.id" class="task-item">
      <span @click="emit('updateTask', { ...task, completed: !task.completed })">
        {{ task.title }}
      </span>
      <button
        class="edit-btn"
        @click="
          () => {
            openEditModal(task)
          }
        "
      >
        編集
      </button>
      <button class="delete-btn" @click="emit('deleteTask', task.id)">削除</button>

      <EditTodoModal
        v-if="showEditModal && currentTask"
        :task="currentTask"
        @on-close="closeEditModal"
        @on-save="
          (updated) => {
            emit('updateTask', updated)
            closeEditModal()
          }
        "
      />
    </li>
  </ul>
</template>
