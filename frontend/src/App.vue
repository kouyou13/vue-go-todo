<script setup lang="ts">
import { ref, onMounted } from "vue"
import dayjs from "dayjs"
import type { Task } from "./types"

// データ定義
const tasks = ref<Task[]>([])
const newTaskTitle = ref<string>("")
const API_URL = "http://localhost:8080/api/tasks"

// --- モーダル関連のデータ定義 ---
const showEditModal = ref(false)
// 編集中のタスク情報を一時的に保持するオブジェクト
const editingTask = ref<Task>({
  id: "",
  title: "",
  completed: false,
  limitedAt: null,
})
console.log(editingTask.value)

// --- API通信用の関数 ---

// タスク一覧の取得 (Read)
const fetchTasks = async () => {
  try {
    const response = await fetch(API_URL)
    tasks.value = await response.json()
    console.log(tasks.value)
  } catch (error) {
    console.error("Error fetching tasks:", error)
  }
}

// タスクの追加 (Create)
const addTask = async () => {
  if (newTaskTitle.value.trim() === "") return

  try {
    const response = await fetch(API_URL, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ title: newTaskTitle.value, completed: false }),
    })
    const newTask = await response.json()
    tasks.value.push(newTask) // 配列に追加
    newTaskTitle.value = "" // 入力欄をクリア
  } catch (error) {
    console.error("Error adding task:", error)
  }
}

// タスクの更新 (完了状態の切替) (Update)
const updateTask = async (task: Task) => {
  try {
    const response = await fetch(`${API_URL}/${task.id}`, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        ...task,
        limitedAt: task.limitedAt ? dayjs(task.limitedAt).toISOString() : null,
      }),
    })
    const updatedTask = await response.json()
    // 配列内の該当タスクを更新
    const index = tasks.value.findIndex((t) => t.id === task.id)
    tasks.value[index] = {
      ...updatedTask,
      limitedAt: updatedTask.limitedAt
        ? dayjs(updatedTask.limitedAt).format("YYYY-MM-DD")
        : null,
    }
  } catch (error) {
    console.error("Error updating task:", error)
  }
}

// タスクの削除 (Delete)
const deleteTask = async (id: string) => {
  if (!confirm("本当に削除しますか？")) return

  try {
    await fetch(`${API_URL}/${id}`, {
      method: "DELETE",
    })
    // 配列から該当タスクを削除
    tasks.value = tasks.value.filter((task) => task.id !== id)
  } catch (error) {
    console.error("Error deleting task:", error)
  }
}

const openEditModal = (task: Task) => {
  editingTask.value = { ...task }
  showEditModal.value = true
}

const closeEditModal = () => {
  editingTask.value = {
    id: "",
    title: "",
    completed: false,
    limitedAt: null,
  }
  showEditModal.value = false
}

// コンポーネントがマウントされたらタスクを取得
onMounted(() => {
  fetchTasks()
})
</script>

<template>
  <div class="container">
    <h1>Vue + Go Todo App</h1>

    <div class="add-task">
      <input
        v-model="newTaskTitle"
        @keyup.enter="addTask"
        type="text"
        placeholder="新しいタスクを入力..."
      />
      <button @click="addTask">追加</button>
    </div>

    <ul>
      <li v-for="task in tasks" :key="task.id" class="task-item">
        <span
          :class="{ completed: task.completed }"
          @click="updateTask({ ...task, completed: !task.completed })"
        >
          {{ task.title }}
        </span>
        <button @click="openEditModal(task)" class="edit-btn">編集</button>
        <button @click="deleteTask(task.id.toString())" class="delete-btn">
          削除
        </button>
      </li>

      <div
        v-if="showEditModal"
        class="modal-overlay"
        @click.self="closeEditModal"
      >
        <div class="modal-content">
          <h2>編集</h2>
          <p>タスク名</p>
          <input v-model="editingTask.title" type="text" class="modal-input" />
          <p>期限</p>
          <input
            v-model="editingTask.limitedAt"
            type="date"
            class="modal-date-input"
          />
          <div class="modal-actions">
            <button @click="closeEditModal" class="cancel-btn">
              キャンセル
            </button>
            <button
              @click="
                () => {
                  updateTask(editingTask)
                  closeEditModal()
                }
              "
              class="save-btn"
            >
              保存
            </button>
          </div>
        </div>
      </div>
    </ul>
  </div>
</template>

<style scoped>
.container {
  max-width: 600px;
  margin: 0 auto;
  padding: 20px;
  font-family: sans-serif;
}

h1 {
  text-align: center;
}

.add-task {
  display: flex;
  gap: 10px;
  margin-bottom: 20px;
}

input {
  flex-grow: 1;
  padding: 10px;
  font-size: 16px;
}

button {
  padding: 10px 20px;
  font-size: 16px;
  cursor: pointer;
}

ul {
  list-style: none;
  padding: 0;
}

p {
  text-align: left;
  margin: 1% 0;
}

.task-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px;
  border-bottom: 1px solid #eee;
}

.task-item span {
  flex-grow: 1;
  cursor: pointer;
}

.completed {
  text-decoration: line-through;
  color: #888;
}

.edit-btn {
  background-color: #3030ff;
  color: white;
  border: none;
  padding: 5px 10px;
  font-size: 14px;
  margin-left: 10px;
}

.delete-btn {
  background-color: #ff4d4d;
  color: white;
  border: none;
  padding: 5px 10px;
  font-size: 14px;
  margin-left: 10px;
}

/* --- モーダル用のスタイルを追加 --- */
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5); /* 半透明の黒背景 */
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000; /* 最前面に表示 */
}

.modal-content {
  background-color: rgba(0, 0, 0);
  padding: 25px;
  border-radius: 8px;
  width: 90%;
  max-width: 500px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.modal-content h2 {
  margin-top: 0;
  margin-bottom: 20px;
  text-align: center;
}

.modal-input {
  width: 100%;
  box-sizing: border-box; /* paddingを含めた幅計算にする */
  margin-bottom: 20px;
}

.modal-date-input {
  width: 100%;
  box-sizing: border-box; /* paddingを含めた幅計算にする */
  margin-bottom: 20px;
}

.modal-actions {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
}

.cancel-btn {
  background-color: #ccc;
  border: none;
  color: black;
  border-radius: 4px;
}

.save-btn {
  background-color: #28a745; /* 緑色 */
  border: none;
  color: white;
  border-radius: 4px;
}
</style>
