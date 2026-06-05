import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueDevTools from 'vite-plugin-vue-devtools'

export default defineConfig({

  plugins: [
    vue(),
    vueDevTools(),
  ],

  resolve: {
    alias: {
      '@': fileURLToPath(
        new URL('./src', import.meta.url)
      )
    },
  },

  server: {

    // Permite acesso pela rede
    host: '0.0.0.0',

    // Mantém a porta atual
    port: 5173
  }
})