import { v4 } from "uuid"

import type { Category } from "../../../types"

export const emptyCategory: Category = {
  id: v4(),
  name: "",
}
