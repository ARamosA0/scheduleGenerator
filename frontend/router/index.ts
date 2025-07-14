import { createRouter, createWebHistory } from "vue-router";
import Home from "../src/Home/home.vue";
import Cursos from "../src/Cursos/cursos.vue";
import Profesores from "../src/Profesores/profesores.vue";
import Salones from "../src/Salones/salones.vue";
import Horarios from "../src/Horarios/horarios.vue";
import Historial from "../src/Historial/historial.vue";
import GenerateSchedule from "../src/Generate/generateSchedule.vue";

const routes = [
  { path: "/", name: "home", component: Home },
  { path: "/cursos", name: "cursos", component: Cursos },
  { path: "/profesores", name: "profesores", component: Profesores },
  { path: "/salones", name: "salones", component: Salones },
  { path: "/horarios", name: "horarios", component: Horarios },
  { path: "/generate", name: "generar", component: GenerateSchedule },
  { path: "/historial", name: "historial", component: Historial },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
