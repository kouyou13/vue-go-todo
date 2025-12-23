import dayjs from "dayjs"
import { ref } from "vue"

import type { Task } from "../../types"

export const useTasks = () => {
  const API_URL: string = import.meta.env.VITE_API_URL
  const tasks = ref<Task[]>([])

  const fetchTasks = async (searchWord: string) => {
    try {
      const url =
        searchWord !== ""
          ? `${API_URL}/tasks?title=${encodeURIComponent(searchWord)}`
          : `${API_URL}/tasks`
      const response = await fetch(url)
      const data: Task[] = await response.json()
      tasks.value = data.map((t: Task) => ({
        ...t,
        limitedAt: t.limitedAt ? dayjs(t.limitedAt).format("YYYY-MM-DD") : null,
      }))
    } catch (error) {
      console.error("Error fetching tasks:", error)
    }
  }

  // タスクの追加 (Create)
  const addTask = async (newTask: Task) => {
    if (newTask.title === "") {
      alert("タスク名が未入力です")
      return
    }

    try {
      const payload = {
        ...newTask,
        limitedAt: newTask.limitedAt ? dayjs(newTask.limitedAt).toISOString() : null,
      }
      await fetch(`${API_URL}/tasks`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload),
      })
      tasks.value.push({
        ...newTask,
        limitedAt: newTask.limitedAt ? dayjs(newTask.limitedAt).format("YYYY-MM-DD") : null,
      }) // 配列に追加
    } catch (error) {
      console.error("Error adding task:", error)
    }
  }

  // タスクの更新 (完了状態の切替) (Update)
  const updateTask = async (task: Task) => {
    if (task.title === "") {
      alert("タスク名が未入力です")
      return
    }

    try {
      const response = await fetch(`${API_URL}/tasks/${task.id}`, {
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
        limitedAt: updatedTask.limitedAt ? dayjs(updatedTask.limitedAt).format("YYYY-MM-DD") : null,
      }
    } catch (error) {
      console.error("Error updating task:", error)
    }
  }

  // タスクの削除 (Delete)
  const deleteTask = async (id: string) => {
    if (!confirm("本当に削除しますか？")) return

    try {
      await fetch(`${API_URL}/tasks/${id}`, {
        method: "DELETE",
      })
      // 配列から該当タスクを削除
      tasks.value = tasks.value.filter((task) => task.id !== id)
    } catch (error) {
      console.error("Error deleting task:", error)
    }
  }

  return {
    tasks,
    fetchTasks,
    addTask,
    updateTask,
    deleteTask,
  }
}
