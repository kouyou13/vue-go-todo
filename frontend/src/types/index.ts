export type Task = {
  id: string
  title: string
  completed: boolean
  limitedAt: string | null
  note: string
  categoryId: string | null
}

export type Category = {
  id: string
  name: string
}
