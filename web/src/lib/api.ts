import axios from 'axios';
import { getScopeCompanyId } from '@/lib/scope-store';

// In dev the Vite proxy forwards /api/* to the Go server (see vite.config.ts).
// In prod the Go server hosts /api/v1 + the SPA dist together, so the relative
// prefix continues to work.
const baseURL = (import.meta.env.VITE_API_BASE_URL ?? '') + '/api/v1';

const api = axios.create({
    baseURL,
    headers: {
        'Content-Type': 'application/json',
        Accept: 'application/json',
    },
});

// ── Request: Bearer Token + Scope header ─────────────────────
api.interceptors.request.use((config) => {
    const token = localStorage.getItem('access_token');

    if (token) {
        config.headers.Authorization = `Bearer ${token}`;
        const scopeId = getScopeCompanyId();

        if (scopeId) {
            config.headers['X-Scope-Company-Id'] = String(scopeId);
        }
    }

    return config;
});

// ── Response: 统一错误处理 ────────────────────────────────────
api.interceptors.response.use(
    (response) => response,
    (error) => {
        if (error.response?.status === 401) {
            localStorage.removeItem('access_token');

            // 避免登录页循环重定向
            if (!window.location.pathname.startsWith('/login')) {
                window.location.href = '/login';
            }
        }

        return Promise.reject(error);
    },
);

export default api;
