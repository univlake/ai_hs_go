import {
    createContext,
    createElement,
    useCallback,
    useContext,
    useEffect,
    useMemo,
    useRef,
    useState,
} from 'react';
import type { ReactNode } from 'react';
import { authApi } from '@/lib/auth';
import type { AuthUser, Role } from '@/types/models';

// ── State ─────────────────────────────────────────────────────────

interface AuthState {
    user: AuthUser | null;
    company: { id: number; title: string } | null;
    roles: Role[];
    permissions: string[];
    isLoading: boolean;
    isAuthenticated: boolean;
}

interface AuthValue extends AuthState {
    hasPermission: (permission: string) => boolean;
    hasRole: (...roles: string[]) => boolean;
    login: (
        username: string,
        password: string,
    ) => Promise<{ access_token: string; token_type: 'Bearer'; expires_at: string }>;
    logout: () => Promise<void>;
    refetch: () => Promise<void>;
}

// ── Context ───────────────────────────────────────────────────────
// useAuth used to be self-contained (each call site kept its own state). That
// broke as soon as login() succeeded in one component but its sibling
// (ScopeProvider, Sidebar, etc.) kept the stale "unauthenticated" state. The
// fix is a single Provider that owns the state and exposes it via context, so
// every consumer sees the same user / company / roles / permissions.

const AuthContext = createContext<AuthValue | null>(null);

// Internal implementation — kept as a hook so the Provider can mount it once.
function useAuthInternal(): AuthValue {
    const [state, setState] = useState<AuthState>({
        user: null,
        company: null,
        roles: [],
        permissions: [],
        isLoading: true,
        isAuthenticated: false,
    });
    const fetched = useRef(false);

    const fetchUser = useCallback(async () => {
        const token = localStorage.getItem('access_token');

        if (!token) {
            setState((s) => ({
                ...s,
                isLoading: false,
                isAuthenticated: false,
            }));
            return;
        }

        try {
            const data = await authApi.userInfo();
            setState({
                user: data.user,
                company: data.company,
                roles: data.roles,
                permissions: data.permissions,
                isLoading: false,
                isAuthenticated: true,
            });
        } catch {
            localStorage.removeItem('access_token');
            setState({
                user: null,
                company: null,
                roles: [],
                permissions: [],
                isLoading: false,
                isAuthenticated: false,
            });
        }
    }, []);

    useEffect(() => {
        if (!fetched.current) {
            fetched.current = true;
            fetchUser();
        }
    }, [fetchUser]);

    const hasPermission = useCallback(
        (permission: string) =>
            state.permissions.includes(permission) ||
            state.user?.is_platform_admin === true ||
            state.roles.some((r) => r.name === 'super-admin'),
        [state.permissions, state.roles, state.user?.is_platform_admin],
    );

    const hasRole = useCallback(
        (...roles: string[]) => state.roles.some((r) => roles.includes(r.name)),
        [state.roles],
    );

    const login = useCallback(
        async (username: string, password: string) => {
            const data = await authApi.login(username, password);
            localStorage.setItem('access_token', data.access_token);
            await fetchUser();
            return data;
        },
        [fetchUser],
    );

    const logout = useCallback(async () => {
        await authApi.logout();
        setState({
            user: null,
            company: null,
            roles: [],
            permissions: [],
            isLoading: false,
            isAuthenticated: false,
        });
        window.location.href = '/login';
    }, []);

    return useMemo(
        () => ({
            ...state,
            hasPermission,
            hasRole,
            login,
            logout,
            refetch: fetchUser,
        }),
        [state, hasPermission, hasRole, login, logout, fetchUser],
    );
}

// ── Provider + hook ───────────────────────────────────────────────

export function AuthProvider({ children }: { children: ReactNode }) {
    const value = useAuthInternal();
    return createElement(AuthContext.Provider, { value }, children);
}

export function useAuth(): AuthValue {
    const ctx = useContext(AuthContext);
    if (!ctx) {
        throw new Error('useAuth must be used inside AuthProvider');
    }
    return ctx;
}
