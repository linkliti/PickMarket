import react from '@vitejs/plugin-react-swc'
import path from "path"
import { defineConfig } from 'vite'
import Terminal from 'vite-plugin-terminal'

// https://vitejs.dev/config/
export default defineConfig({
  server: {
    port: 5005,
    strictPort: true,
  },
  preview: {
    port: 5005,
    strictPort: true,
  },
  plugins: [react(), Terminal()],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./app"),
    },
  },
})
