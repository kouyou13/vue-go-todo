import { createRouter, createWebHistory } from "vue-router"

import Todo from "../pages/Todo.vue"

export default createRouter({
  history: createWebHistory(),
  routes: [{ path: "/", component: Todo }],
})
