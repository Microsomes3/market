import { defineConfig } from 'vite';
import tailwindcss from '@tailwindcss/vite'


export default defineConfig({
    plugins: [    tailwindcss(),  ],
  server: {
    host: true, // Listen on all available network interfaces
    port: 5173, // Optional: specify a port
  },
});