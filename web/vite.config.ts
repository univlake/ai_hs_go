import path from 'node:path';
import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import tailwindcss from '@tailwindcss/vite';

// SPA build — no Laravel/Inertia/Wayfinder plugins. The Go backend serves the
// REST API and statically hosts the dist/ output in production.
export default defineConfig({
  plugins: [
    react({
      babel: {
        plugins: [['babel-plugin-react-compiler', {}]],
      },
    }),
    tailwindcss(),
  ],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, 'src'),
      // Pages copied from the Laravel project still import from
      // `@inertiajs/react`. The shim re-exports the small Inertia surface
      // (Head, Link, router, usePage, useForm) on top of react-router-dom +
      // axios so legacy pages keep compiling. Migrate pages off the shim
      // incrementally; once unused we can drop both alias and shim.
      '@inertiajs/react': path.resolve(__dirname, 'src/lib/inertia-shim.tsx'),
    },
  },
  server: {
    host: true,
    port: 5173,
    proxy: {
      '/api': {
        target: process.env.VITE_API_PROXY ?? 'http://127.0.0.1:8080',
        changeOrigin: true,
      },
    },
  },
  build: {
    outDir: 'dist',
    sourcemap: true,
  },
});
