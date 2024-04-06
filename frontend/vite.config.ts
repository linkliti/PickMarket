import react from '@vitejs/plugin-react-swc'
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
})
