<script setup>
import { ref, onMounted } from 'vue'

// データ定義
const tasks = ref([])
const newTaskTitle = ref('')
const API_URL = 'http://localhost:8080/api/tasks'

// --- API通信用の関数 ---

// タスク一覧の取得 (Read)
const fetchTasks = async () => {
  try {
    const response = await fetch(API_URL)
    tasks.value = await response.json()
  } catch (error) {
    console.error('Error fetching tasks:', error)
  }
}

// タスクの追加 (Create)
const addTask = async () => {
  if (newTaskTitle.value.trim() === '') return

  try {
    const response = await fetch(API_URL, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ title: newTaskTitle.value, completed: false })
    })
    const newTask = await response.json()
    tasks.value.push(newTask) // 配列に追加
    newTaskTitle.value = '' // 入力欄をクリア
  } catch (error) {
    console.error('Error adding task:', error)
  }
}

// タスクの更新 (完了状態の切替) (Update)
const toggleComplete = async (task) => {
  try {
    const response = await fetch(`${API_URL}/${task.id}`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ ...task, completed: !task.completed })
    })
    const updatedTask = await response.json()
    // 配列内の該当タスクを更新
    const index = tasks.value.findIndex(t => t.id === task.id)
    tasks.value[index] = updatedTask
  } catch (error) {
    console.error('Error updating task:', error)
  }
}

// タスクの削除 (Delete)
const deleteTask = async (id) => {
  if (!confirm('本当に削除しますか？')) return

  try {
    await fetch(`${API_URL}/${id}`, {
      method: 'DELETE'
    })
    // 配列から該当タスクを削除
    tasks.value = tasks.value.filter(task => task.id !== id)
  } catch (error) {
    console.error('Error deleting task:', error)
  }
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
          @click="toggleComplete(task)"
        >
          {{ task.title }}
        </span>
        <button @click="deleteTask(task.id)" class="delete-btn">削除</button>
      </li>
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

.delete-btn {
  background-color: #ff4d4d;
  color: white;
  border: none;
  padding: 5px 10px;
  font-size: 14px;
  margin-left: 10px;
}
</style>