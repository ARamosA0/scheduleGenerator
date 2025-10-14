import { createApp } from "vue";
import "./style.css";
import App from "./App.vue";
import router from "../router";
import PrimeVue from "primevue/config";
import Aura from "@primeuix/themes/aura";
import { DataTable, Column } from "primevue";
import { createPinia } from "pinia";
// createApp(App).mount('#app')
const app = createApp(App);
const pinia = createPinia()
app.use(PrimeVue, {
  theme: {
    preset: Aura,
  },
});
app.use(router);
app.use(pinia)

app.component("DataTable", DataTable);
app.component("Column", Column);

app.mount("#app");
