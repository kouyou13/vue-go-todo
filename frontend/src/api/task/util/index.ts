import { v4 } from "uuid"

import type { Task } from "../../../types"

export const emptyTask: Task = {
  id: v4(),
  title: "",
  completed: false,
  limitedAt: null,
  note: "",
  categoryId: null,
}
