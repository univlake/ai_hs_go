// SPA entry — replaces the Inertia bootstrap from app.tsx.
//
// Notes for migration off Inertia:
// - All pages under src/pages/**/*.tsx are statically registered with React Router
//   in src/router.tsx. Old Inertia <Link> calls and `router.visit` should be
//   migrated to react-router's <Link>/<Navigate>/useNavigate.
// - Generated wayfinder helpers in src/actions/ and src/routes/ remain on disk
//   for now so pages still compile; they will be removed module-by-module as
//   pages migrate to centralised URL constants in src/lib/api-routes.ts.
// - Auth/session flow stays the same: REST + Bearer token + X-Scope-Company-Id.

import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { RouterProvider } from 'react-router-dom';
import { Toaster } from 'sonner';

import '@/styles/app.css';
import { router } from '@/router';
import { ThemeProvider } from '@/contexts/ThemeContext';
import { ScopeProvider } from '@/contexts/ScopeContext';
import { AuthProvider } from '@/hooks/use-auth';

const root = document.getElementById('app');
if (!root) {
  throw new Error('mount node #app missing');
}

createRoot(root).render(
  <StrictMode>
    <AuthProvider>
      <ThemeProvider>
        <ScopeProvider>
          <RouterProvider router={router} />
          <Toaster position="top-right" richColors closeButton />
        </ScopeProvider>
      </ThemeProvider>
    </AuthProvider>
  </StrictMode>,
);
