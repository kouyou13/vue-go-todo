import { createRouter, createWebHistory } from "vue-router"

import Categories from "../pages/Categories.vue"
import Todo from "../pages/Todo.vue"

export default createRouter({
  history: createWebHistory(),
  routes: [
    { path: "/", component: Todo },
    { path: "/categories", component: Categories },
  ],
})
