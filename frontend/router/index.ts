import { createRouter, createWebHistory } from "vue-router";
import Home from "../src/Home/home.vue";
// import About from "../views/About.vue";

const routes = [{ path: "/", name: "home", component: Home }];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
