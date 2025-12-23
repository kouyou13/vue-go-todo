import { createRouter, createWebHistory } from "vue-router"

import TodoPage from "../pages/TodoPage.vue"

export default createRouter({
  history: createWebHistory(),
  routes: [{ path: "/", component: TodoPage }],
})
