import react from "@vitejs/plugin-react-swc";
import path from "path";
import { defineConfig } from "vite";
import terminal from "vite-plugin-terminal";

// https://vitejs.dev/config/

const frontHost: string = process.env.FRONTEND_ADDR ?? "127.0.0.1:1111";
const handlerHost: string = process.env.HANDLER_ADDR ?? "127.0.0.1:1111";

const host: string[] = frontHost.split(":");
const handlerAddr: string = `http://${handlerHost}`;

export default defineConfig({
  server: {
    port: Number.parseInt(host[1]),
    host: host[0],
    strictPort: true,
    proxy: {
      "/api": {
        target: handlerAddr,
        changeOrigin: true,
        secure: false,
        rewrite: (path: string):string => path.replace(/^\/api/, ""),
      },
    },
  },
  preview: {
    port: Number.parseInt(host[1]),
    host: host[0],
    strictPort: true,
    proxy: {
      "/api": {
        target: handlerAddr,
        changeOrigin: true,
        secure: false,
        rewrite: (path: string):string => path.replace(/^\/api/, ""),
      },
    },
  },
  // plugins: [react(), terminal({ console: "terminal" })],
  plugins: [react(), terminal()],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./app"),
    },
  },
});
