import { ref } from "vue"

import type { Category } from "../../types"

export const useCategories = () => {
  const API_URL: string = import.meta.env.VITE_API_URL
  const categories = ref<Category[]>([])

  const fetchCategories = async () => {
    try {
      const url = `${API_URL}/categories`
      const response = await fetch(url)
      categories.value = await response.json()
    } catch (error) {
      console.error("Error fetching categories:", error)
    }
  }

  const addCategory = async (newCategory: Category) => {
    if (newCategory.name === "") {
      alert("カテゴリー名が未入力です")
      return
    }

    try {
      await fetch(`${API_URL}/categories`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(newCategory),
      })
      categories.value.push(newCategory)
    } catch (error) {
      console.error("Error adding category:", error)
    }
  }

  const updateCategory = async (category: Category) => {
    if (category.name === "") {
      alert("タスク名が未入力です")
      return
    }

    try {
      await fetch(`${API_URL}/categories/${category.id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(category),
      })
      // 配列内の該当タスクを更新
      const index = categories.value.findIndex((c) => c.id === category.id)
      categories.value[index] = category
    } catch (error) {
      console.error("Error updating task:", error)
    }
  }

  const deleteCategory = async (id: string) => {
    if (!confirm("本当に削除しますか？")) return

    try {
      await fetch(`${API_URL}/categories/${id}`, {
        method: "DELETE",
      })
      // 配列から該当タスクを削除
      categories.value = categories.value.filter((c) => c.id !== id)
    } catch (error) {
      console.error("Error deleting task:", error)
    }
  }

  return {
    categories,
    fetchCategories,
    addCategory,
    updateCategory,
    deleteCategory,
  }
}
