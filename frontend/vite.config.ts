import react from "@vitejs/plugin-react-swc";
import path from "path";
import { defineConfig } from "vite";
import Terminal from "vite-plugin-terminal";

// https://vitejs.dev/config/

const port = (process.env.FRONTEND_ADDR ?? "127.0.0.1:1111").split(":")[2];
const handlerAddr = process.env.HANDLER_ADDR ?? "http://127.0.0.1:1111";
console.log("handlerAddr", handlerAddr);

export default defineConfig({
  server: {
    port: Number.parseInt(port),
    strictPort: true,
    proxy: {
      "/api": {
        target: handlerAddr,
        changeOrigin: true,
        secure: false,
        rewrite: (path) => path.replace(/^\/api/, ""),
      },
    },
  },
  preview: {
    port: Number.parseInt(port),
    strictPort: true,
    proxy: {
      "/api": {
        target: handlerAddr,
        changeOrigin: true,
        secure: false,
        rewrite: (path) => path.replace(/^\/api/, ""),
      },
    },
  },
  plugins: [react(), Terminal({ console: "terminal" })],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./app"),
    },
  },
});
