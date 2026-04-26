import { Outlet, useLocation, useNavigate } from 'react-router-dom';
import { Shield, ArrowLeft, EyeOff } from 'lucide-react';
import { useEffect, useState } from 'react';
import { toast } from 'sonner';
import { ErrorBoundary } from '@/components/ErrorBoundary';
import { useScope } from '@/contexts/ScopeContext';
import { useAuth } from '@/hooks/use-auth';
import { Header } from './Header';
import { Sidebar, getRoutePermission } from './Sidebar';

// AppLayout is mounted as a route layout; children come from react-router's
// <Outlet />. ScopeProvider / ThemeProvider live in main.tsx so all pages
// share a single instance.
export default function AppLayout() {
    const { user, isLoading, isAuthenticated, hasPermission } = useAuth();
    const location = useLocation();
    const navigate = useNavigate();
    const [sidebarOpen, setSidebarOpen] = useState(true);
    const url = location.pathname + location.search;

    useEffect(() => {
        if (isLoading) {
            return;
        }

        if (!isAuthenticated) {
            navigate('/login', { replace: true });

            return;
        }

        // 首次登录必改账号信息 — 强制去 /first-login-setup
        if (user?.must_change_credentials) {
            navigate('/first-login-setup', { replace: true });
        }
    }, [isLoading, isAuthenticated, user, navigate]);

    // Rule 3 at the route level: if the user navigates (or pastes a URL) to a
    // page they don't have permission for, bounce them to /dashboard. Runs after
    // auth finishes loading so we don't redirect while permissions are empty.
    useEffect(() => {
        if (isLoading || !isAuthenticated) {
            return;
        }

        const pathname = location.pathname;
        const required = getRoutePermission(pathname);

        if (!required) {
            return;
        }

        const perms = Array.isArray(required) ? required : [required];
        const allowed = perms.some((p) => hasPermission(p));

        if (!allowed && pathname !== '/dashboard') {
            navigate('/dashboard', { replace: true });
        }
    }, [location.pathname, isLoading, isAuthenticated, hasPermission, navigate]);

    // Platform-admin path separation: a super admin who hasn't entered a
    // tenant via impersonation only has access to platform-layer routes.
    // Hitting any business path silently bounces them back to /platform/tenants
    // until they explicitly impersonate a tenant.
    useEffect(() => {
        if (isLoading || !isAuthenticated) {
            return;
        }
        const isPlatformAdmin = user?.is_platform_admin === true;
        if (!isPlatformAdmin) {
            return;
        }
        // Persisted impersonation lives in localStorage; ScopeContext reads
        // it but we duplicate here because AppLayout doesn't read scope.
        const isImpersonating = !!localStorage.getItem('platform_impersonation');
        if (isImpersonating) {
            return;
        }
        const path = location.pathname;
        const platformWhitelist =
            path === '/dashboard' ||
            path === '/profile' ||
            path === '/change-password' ||
            path === '/first-login-setup' ||
            path.startsWith('/platform');
        if (!platformWhitelist) {
            navigate('/platform/tenants', { replace: true });
        }
    }, [location.pathname, isLoading, isAuthenticated, user, navigate]);

    if (isLoading) {
        return (
            <div className="flex h-screen items-center justify-center bg-[#070d1f]">
                <div className="flex flex-col items-center gap-4">
                    <div className="flex h-14 w-14 items-center justify-center rounded-2xl border border-blue-500/30 bg-blue-600/15">
                        <Shield className="h-7 w-7 animate-pulse text-blue-400" />
                    </div>
                    <p className="text-sm text-slate-500">验证登录状态...</p>
                </div>
            </div>
        );
    }

    if (!isAuthenticated) {
        return null;
    }

    return (
        <AuthenticatedShell
            sidebarOpen={sidebarOpen}
            setSidebarOpen={setSidebarOpen}
            resetKey={url}
        />
    );
}

// ── Inner shell — uses ScopeProvider (mounted in main.tsx) ───────────────

function AuthenticatedShell({
    sidebarOpen,
    setSidebarOpen,
    resetKey,
}: {
    sidebarOpen: boolean;
    setSidebarOpen: (v: boolean) => void;
    resetKey: string;
}) {
    const {
        isSubordinate,
        activeCompany,
        rootCompany,
        exitToRoot,
        impersonating,
        impersonatedTenant,
        exitImpersonation,
    } = useScope();
    const navigate = useNavigate();

    const handleExitImpersonation = async () => {
        try {
            await exitImpersonation();
            toast.success('已退出化身');
            navigate('/platform/tenants', { replace: true });
        } catch (err) {
            toast.error('退出化身失败');
            // eslint-disable-next-line no-console
            console.error(err);
        }
    };

    return (
        <div className="flex h-screen overflow-hidden bg-[#070d1f]">
            <Sidebar open={sidebarOpen} onOpenChange={setSidebarOpen} />

            <div className="flex flex-1 flex-col overflow-hidden">
                <Header onMenuClick={() => setSidebarOpen(!sidebarOpen)} />

                {/* Platform impersonation banner */}
                {impersonating && impersonatedTenant && (
                    <div className="flex shrink-0 items-center gap-3 border-b border-orange-500/30 bg-gradient-to-r from-orange-950/40 to-amber-950/40 px-5 py-2.5">
                        <div className="flex items-center gap-2">
                            <span className="relative flex h-2 w-2">
                                <span className="absolute inline-flex h-full w-full animate-ping rounded-full bg-orange-400 opacity-60"></span>
                                <span className="relative inline-flex h-2 w-2 rounded-full bg-orange-400"></span>
                            </span>
                            <span className="text-xs font-semibold tracking-wide text-orange-300 uppercase">
                                平台管理员化身
                            </span>
                        </div>
                        <span className="text-slate-700">·</span>
                        <span className="text-sm text-slate-300">
                            正以{' '}
                            <strong className="text-orange-200">
                                「{impersonatedTenant.title}」
                            </strong>{' '}
                            身份查看，所有写操作都会以该企业为上下文
                        </span>
                        <button
                            onClick={handleExitImpersonation}
                            className="ml-auto flex items-center gap-1.5 rounded-md border border-orange-500/40 bg-orange-600/15 px-3 py-1 text-xs font-medium text-orange-300 transition-colors hover:border-orange-400/60 hover:bg-orange-600/25"
                        >
                            <EyeOff className="h-3 w-3" />
                            退出化身
                        </button>
                    </div>
                )}

                {/* Subordinate scope breadcrumb */}
                {!impersonating && isSubordinate && activeCompany && rootCompany && (
                    <div className="flex shrink-0 items-center gap-3 border-b border-amber-500/20 bg-amber-950/20 px-5 py-2.5">
                        <button
                            onClick={exitToRoot}
                            className="flex items-center gap-1.5 rounded-md border border-amber-500/25 bg-amber-600/10 px-2.5 py-1 text-xs font-medium text-amber-400 transition-colors hover:bg-amber-600/20"
                        >
                            <ArrowLeft className="h-3 w-3" />
                            返回集团
                        </button>
                        <div className="flex items-center gap-1.5 text-xs text-slate-500">
                            <span className="text-slate-600">
                                {rootCompany.title}
                            </span>
                            <span className="text-slate-700">›</span>
                            <span className="font-medium text-amber-400">
                                {activeCompany.title}
                            </span>
                        </div>
                        <span className="ml-auto rounded-full border border-amber-500/20 bg-amber-600/10 px-2 py-0.5 text-[10px] text-amber-500">
                            生产视图
                        </span>
                    </div>
                )}

                <main className="flex-1 overflow-y-auto p-6">
                    <ErrorBoundary resetKey={resetKey}>
                        <Outlet />
                    </ErrorBoundary>
                </main>
            </div>
        </div>
    );
}
