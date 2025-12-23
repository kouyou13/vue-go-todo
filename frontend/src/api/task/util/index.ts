import type { Task } from "../../../types"

export const emptyTask: Task = {
  id: "",
  title: "",
  completed: false,
  limitedAt: null,
  note: "",
  categoryId: null,
}
