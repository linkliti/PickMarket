import react from '@vitejs/plugin-react-swc'
import path from "path"
import { defineConfig } from 'vite'
import Terminal from 'vite-plugin-terminal'

// https://vitejs.dev/config/

const port = (process.env.FRONTEND_ADDR ?? '127.0.0.1:1111').split(':')[1]

export default defineConfig({
  server: {
    port: Number.parseInt(port),
    strictPort: true,
  },
  preview: {
    port: Number.parseInt(port),
    strictPort: true,
  },
  plugins: [react(), Terminal()],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./app"),
    },
  },
})
