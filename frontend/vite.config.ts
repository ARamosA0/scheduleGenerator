import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import tailwindcss from "@tailwindcss/vite";

// https://vite.dev/config/
export default defineConfig({
  plugins: [vue(), tailwindcss()],
  server: {
    host: "0.0.0.0", // Escucha en todas las interfaces
    port: 5173, // Puerto definido (debe coincidir con el EXPOSE en Docker)
    strictPort: true, // Evita que Vite cambie el puerto si está ocupado
    watch: {
      usePolling: true,
      interval: 100,
    },
  },
});
