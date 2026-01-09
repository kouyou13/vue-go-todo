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
  <ul class="list-none p-0">
    <li
      v-for="task in tasks"
      :key="task.id"
      class="flex justify-between items-center p-2.5 border-b-2 border-solid border-white"
    >
      <span
        class="grow cursor-pointer"
        :class="{'line-through': task.completed}"
        @click="emit('updateTask', { ...task, completed: !task.completed })"
      >
        {{ task.title }}
      </span>
      <button
        class="text-center px-3 text-sm border-0 text-white bg-blue-700 h-8"
        @click="
          () => {
            openEditModal(task)
          }
        "
      >
        編集
      </button>
      <button
        class="text-center px-3 text-sm border-0 text-white bg-red-600 h-8 ml-2"
        @click="emit('deleteTask', task.id)"
      >
        削除
      </button>

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
