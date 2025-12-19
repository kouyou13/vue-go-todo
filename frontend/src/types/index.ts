export type Task = {
  id: string
  title: string
  completed: Boolean
  limitedAt: string | null
  note: string
  categoryId: string | null
}

export type Category = {
  id: string
  name: string
}