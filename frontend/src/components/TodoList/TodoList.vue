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
  <table class="w-full">
    <tr class="border-b">
      <th class="text-sm text-gray-400">タイトル</th>
      <th class="text-sm text-gray-400">期限</th>
      <th></th>
      <th></th>
    </tr>
    <tr v-for="task in tasks" :key="task.id" class="border-b">
      <td
        class="text-lg cursor-pointer"
        :class="{ 'line-through': task.completed }"
        @click="emit('updateTask', { ...task, completed: !task.completed })"
      >
        {{ task.title }}
      </td>
      <td v-if="task.limitedAt" class="w-1/10 text-lg">{{ task.limitedAt }}</td>
      <td v-if="!task.limitedAt" class="w-1/10 text-lg">なし</td>
      <td class="w-1/18 pl-2">
        <button
          class="text-center px-3 text-sm border-0 text-white bg-blue-700 h-10 w-full"
          @click="
            () => {
              openEditModal(task)
            }
          "
        >
          編集
        </button>
      </td>
      <td class="w-1/18 pl-3">
        <button
          class="text-center px-3 text-sm border-0 text-white bg-red-600 h-10 w-full my-4"
          @click="emit('deleteTask', task.id)"
        >
          削除
        </button>
      </td>
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
    </tr>
  </table>
</template>
