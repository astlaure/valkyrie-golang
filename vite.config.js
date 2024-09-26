import { defineConfig } from 'vite';

export default defineConfig({
  root: 'web/static',
  build: {
    manifest: true,
    rollupOptions: {
      input: 'web/static/main.js',
    }
  }
})
