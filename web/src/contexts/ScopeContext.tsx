import {
    createContext,
    useContext,
    useState,
    useEffect,
    useCallback,
    useRef,
} from 'react';
import type { ReactNode } from 'react';
import { useAuth } from '@/hooks/use-auth';
import api from '@/lib/api';
import { API } from '@/lib/api-routes';
import { setScopeCompanyId } from '@/lib/scope-store';

// ── Types ─────────────────────────────────────────────────────────

export interface ScopeCompany {
    id: number;
    title: string;
    is_bloc: number;
    is_produce: number;
    pid: number;
    status: number;
}

export type ScopeMode = 'bloc' | 'produce';

interface PersistedScope {
    rootCompanyId: number;
    activeCompanyId: number;
    mode: ScopeMode;
}

interface ScopeContextType {
    /** Currently active (scoped) company */
    activeCompany: ScopeCompany | null;
    /** Current feature mode */
    mode: ScopeMode;
    /** True when the active company differs from the root company */
    isSubordinate: boolean;

    /** The logged-in user's own company (or the actively-impersonated tenant for platform admins) */
    rootCompany: ScopeCompany | null;
    /** Root company capabilities */
    canBloc: boolean;
    canProduce: boolean;

    /** Direct subordinates of the root company (for the switcher).
     *  For platform admins this is the list of every tenant. */
    subordinates: ScopeCompany[];
    loadingSubordinates: boolean;

    /** True when the active session is a platform super admin. */
    isPlatformMode: boolean;

    /**
     * True when a platform admin is currently acting on a specific tenant
     * (impersonation). `impersonatedTenant` holds that tenant. Different from
     * `isPlatformMode + activeCompany` so the UI can show a banner only when
     * the admin has explicitly entered a tenant rather than just defaulting
     * to one.
     */
    impersonating: boolean;
    impersonatedTenant: ScopeCompany | null;

    setMode: (mode: ScopeMode) => void;
    enterCompany: (company: ScopeCompany) => void;
    exitToRoot: () => void;

    /** Platform-only: enter / exit tenant impersonation. */
    enterImpersonation: (tenant: ScopeCompany) => Promise<void>;
    exitImpersonation: () => Promise<void>;

    isReady: boolean;
}

// ── Context ───────────────────────────────────────────────────────

const ScopeContext = createContext<ScopeContextType | null>(null);

const STORAGE_KEY = 'scope_state';
const IMPERSONATION_KEY = 'platform_impersonation';

function loadPersisted(): PersistedScope | null {
    try {
        const raw = localStorage.getItem(STORAGE_KEY);

        return raw ? JSON.parse(raw) : null;
    } catch {
        return null;
    }
}

function savePersisted(data: PersistedScope): void {
    localStorage.setItem(STORAGE_KEY, JSON.stringify(data));
}

function clearPersisted(): void {
    localStorage.removeItem(STORAGE_KEY);
}

// Impersonation persistence — survive a page refresh so the admin doesn't
// silently fall out of the tenant they were inspecting.
function loadImpersonation(): number | null {
    try {
        const raw = localStorage.getItem(IMPERSONATION_KEY);
        if (!raw) return null;
        const id = Number(raw);
        return Number.isFinite(id) && id > 0 ? id : null;
    } catch {
        return null;
    }
}
function saveImpersonation(id: number) {
    localStorage.setItem(IMPERSONATION_KEY, String(id));
}
function clearImpersonation() {
    localStorage.removeItem(IMPERSONATION_KEY);
}

// Backend tenant rows include parent_id (camelCase: ParentID) instead of pid;
// adapt to the legacy ScopeCompany shape so existing UI keeps working.
function adaptTenant(row: {
    id: number;
    title: string;
    is_bloc: number;
    is_produce: number;
    parent_id: number | null;
    status: number;
}): ScopeCompany {
    return {
        id: row.id,
        title: row.title,
        is_bloc: row.is_bloc,
        is_produce: row.is_produce,
        pid: row.parent_id ?? 0,
        status: row.status,
    };
}

// ── Provider ──────────────────────────────────────────────────────

export function ScopeProvider({ children }: { children: ReactNode }) {
    const { user, company: authCompany, isAuthenticated, isLoading } = useAuth();
    const isPlatformMode = user?.is_platform_admin === true;

    const [rootCompany, setRootCompany] = useState<ScopeCompany | null>(null);
    const [activeCompany, setActiveCompany] = useState<ScopeCompany | null>(
        null,
    );
    const [mode, setModeState] = useState<ScopeMode>('produce');
    const [subordinates, setSubordinates] = useState<ScopeCompany[]>([]);
    const [loadingSubordinates, setLoadingSubordinates] = useState(false);
    const [isReady, setIsReady] = useState(false);
    const [impersonatedTenant, setImpersonatedTenant] =
        useState<ScopeCompany | null>(null);

    const initialized = useRef(false);

    // ── Initialise scope ──────────────────────────────────────────
    useEffect(() => {
        if (isLoading || !isAuthenticated) {
            return;
        }

        if (initialized.current) {
            return;
        }

        // Platform admin: no own company. Use tenants list as the universe of
        // possible "scopes" and let the switcher pick which tenant to act on.
        if (isPlatformMode) {
            initialized.current = true;
            setLoadingSubordinates(true);
            api.get('/tenants')
                .then(async (res) => {
                    const tenants: ScopeCompany[] = (
                        res.data?.data ?? []
                    ).map(adaptTenant);
                    setSubordinates(tenants);

                    // Impersonation persistence wins over the regular
                    // last-active scope: if the platform admin was sudo'd
                    // into a company, restore that on refresh. Sub-company
                    // impersonations won't be in /tenants (forest roots
                    // only), so fall back to fetching the company directly.
                    const impersonatedId = loadImpersonation();
                    let impersonated: ScopeCompany | undefined;
                    if (impersonatedId) {
                        impersonated = tenants.find(
                            (t) => t.id === impersonatedId,
                        );
                        if (!impersonated) {
                            try {
                                const r = await api.get(
                                    `/companies/${impersonatedId}`,
                                );
                                impersonated = adaptTenant(r.data.data);
                            } catch {
                                // company gone / status flipped — drop the
                                // stale impersonation and fall through to
                                // the default flow below
                                clearImpersonation();
                            }
                        }
                    }

                    if (impersonated) {
                        setRootCompany(impersonated);
                        setActiveCompany(impersonated);
                        setImpersonatedTenant(impersonated);
                        setModeState(
                            impersonated.is_bloc && !impersonated.is_produce
                                ? 'bloc'
                                : 'produce',
                        );
                        setScopeCompanyId(impersonated.id);
                        setLoadingSubordinates(false);
                        setIsReady(true);
                        return;
                    }

                    // No active impersonation — fall back to last-active
                    // tenant or first available so the switcher has a
                    // sensible default.
                    const persisted = loadPersisted();
                    const remembered = tenants.find(
                        (t) => persisted && t.id === persisted.activeCompanyId,
                    );

                    if (remembered) {
                        setRootCompany(remembered);
                        setActiveCompany(remembered);
                        setModeState(persisted!.mode);
                        setScopeCompanyId(remembered.id);
                    } else if (tenants.length > 0) {
                        const first = tenants[0];
                        setRootCompany(first);
                        setActiveCompany(first);
                        setModeState(
                            first.is_bloc && !first.is_produce
                                ? 'bloc'
                                : 'produce',
                        );
                        setScopeCompanyId(first.id);
                    } else {
                        setRootCompany(null);
                        setActiveCompany(null);
                        setScopeCompanyId(null);
                    }
                })
                .catch(() => setSubordinates([]))
                .finally(() => {
                    setLoadingSubordinates(false);
                    setIsReady(true);
                });
            return;
        }

        // Regular user: must have a company
        if (!authCompany?.id) {
            return;
        }

        initialized.current = true;

        api.get(`/companies/${authCompany.id}`)
            .then((res) => {
                const full: ScopeCompany = res.data.data;
                setRootCompany(full);

                const defaultMode: ScopeMode =
                    full.is_bloc && !full.is_produce
                        ? 'bloc'
                        : full.is_produce && !full.is_bloc
                          ? 'produce'
                          : 'bloc';

                const persisted = loadPersisted();

                if (persisted && persisted.rootCompanyId === full.id) {
                    setModeState(persisted.mode);

                    if (persisted.activeCompanyId === full.id) {
                        setActiveCompany(full);
                        setScopeCompanyId(full.id);
                    } else {
                        api.get(`/companies/${persisted.activeCompanyId}`)
                            .then((r) => {
                                setActiveCompany(r.data.data);
                                setScopeCompanyId(r.data.data.id);
                            })
                            .catch(() => {
                                setActiveCompany(full);
                                setScopeCompanyId(full.id);
                            });
                    }
                } else {
                    clearPersisted();
                    setModeState(defaultMode);
                    setActiveCompany(full);
                    setScopeCompanyId(full.id);
                }

                setIsReady(true);
            })
            .catch(() => {
                const fallback: ScopeCompany = {
                    id: authCompany.id,
                    title: authCompany.title,
                    is_bloc: 0,
                    is_produce: 1,
                    pid: 0,
                    status: 1,
                };
                setRootCompany(fallback);
                setActiveCompany(fallback);
                setModeState('produce');
                setScopeCompanyId(fallback.id);
                setIsReady(true);
            });
    }, [isLoading, isAuthenticated, authCompany, isPlatformMode]);

    // ── Fetch direct subordinates ─────────────────────────────────
    //
    //   • Regular user → children of their own company.
    //   • Platform admin without impersonation → init effect already loaded
    //     the tenants list as `subordinates`; nothing to do here.
    //   • Platform admin while impersonating any company (root or sub) →
    //     load that company's direct children so the switcher behaves
    //     exactly like a regular user of the impersonated company.
    useEffect(() => {
        if (!rootCompany) return;
        if (isPlatformMode && !impersonatedTenant) return;

        setLoadingSubordinates(true);
        api.get('/companies', { params: { pid: rootCompany.id, per_page: 50 } })
            .then((res) => {
                const items: ScopeCompany[] = res.data.data?.data ?? [];
                setSubordinates(items);
            })
            .catch(() => setSubordinates([]))
            .finally(() => setLoadingSubordinates(false));
    }, [rootCompany?.id, isPlatformMode, impersonatedTenant?.id]);

    // ── Reset when user logs out ──────────────────────────────────
    useEffect(() => {
        if (!isAuthenticated && !isLoading) {
            setRootCompany(null);
            setActiveCompany(null);
            setSubordinates([]);
            setImpersonatedTenant(null);
            setIsReady(false);
            initialized.current = false;
            clearPersisted();
            clearImpersonation();
            setScopeCompanyId(null);
        }
    }, [isAuthenticated, isLoading]);

    // ── Actions ───────────────────────────────────────────────────

    const persist = useCallback(
        (active: ScopeCompany, newMode: ScopeMode, root: ScopeCompany) => {
            savePersisted({
                rootCompanyId: root.id,
                activeCompanyId: active.id,
                mode: newMode,
            });
            setScopeCompanyId(active.id);
        },
        [],
    );

    const setMode = useCallback(
        (newMode: ScopeMode) => {
            if (!rootCompany || !activeCompany) {
                return;
            }

            setModeState(newMode);
            setActiveCompany(rootCompany);
            persist(rootCompany, newMode, rootCompany);
        },
        [rootCompany, activeCompany, persist],
    );

    const enterCompany = useCallback(
        (company: ScopeCompany) => {
            // Drilling into a subordinate must NEVER overwrite rootCompany —
            // exitToRoot reads it to navigate back, and the impersonation
            // banner reads it to display "正以 ... 身份查看". Used to behave
            // differently for platform admin (replaced root, breaking exit
            // and banner); now both regular users and impersonating platform
            // admins follow the same path.
            if (!rootCompany) {
                return;
            }
            setActiveCompany(company);
            setModeState('produce');
            persist(company, 'produce', rootCompany);
        },
        [rootCompany, persist],
    );

    const exitToRoot = useCallback(() => {
        if (!rootCompany) {
            return;
        }

        const defaultMode: ScopeMode = rootCompany.is_bloc ? 'bloc' : 'produce';
        setActiveCompany(rootCompany);
        setModeState(defaultMode);
        persist(rootCompany, defaultMode, rootCompany);
    }, [rootCompany, persist]);

    // ── Impersonation ─────────────────────────────────────────────

    const enterImpersonation = useCallback(
        async (tenant: ScopeCompany) => {
            if (!isPlatformMode) {
                throw new Error('only platform admin can impersonate');
            }
            await api.post(API.impersonate(), { tenant_id: tenant.id });
            const newMode: ScopeMode =
                tenant.is_bloc && !tenant.is_produce ? 'bloc' : 'produce';
            setRootCompany(tenant);
            setActiveCompany(tenant);
            setImpersonatedTenant(tenant);
            setModeState(newMode);
            saveImpersonation(tenant.id);
            persist(tenant, newMode, tenant);
        },
        [isPlatformMode, persist],
    );

    const exitImpersonation = useCallback(async () => {
        const id = impersonatedTenant?.id;
        try {
            await api.post(API.impersonateExit(), id ? { tenant_id: id } : {});
        } catch {
            // best-effort — even if the audit row write fails, the admin
            // should still be able to drop out of impersonation.
        }
        clearImpersonation();
        setImpersonatedTenant(null);

        // Re-fetch the tenants list as the platform-admin's universe of
        // scopes. We can't trust `subordinates` here because, while
        // impersonating, it was overwritten with the impersonated company's
        // direct children — which would be the wrong fallback after exit.
        let tenants: ScopeCompany[] = [];
        try {
            const res = await api.get('/tenants');
            tenants = (res.data?.data ?? []).map(adaptTenant);
            setSubordinates(tenants);
        } catch {
            setSubordinates([]);
        }

        const fallback = tenants[0] ?? null;
        if (fallback) {
            const newMode: ScopeMode =
                fallback.is_bloc && !fallback.is_produce ? 'bloc' : 'produce';
            setRootCompany(fallback);
            setActiveCompany(fallback);
            setModeState(newMode);
            setScopeCompanyId(fallback.id);
        } else {
            setRootCompany(null);
            setActiveCompany(null);
            setScopeCompanyId(null);
        }
    }, [impersonatedTenant]);

    // ── Derived ───────────────────────────────────────────────────

    const canBloc = Boolean(rootCompany?.is_bloc);
    const canProduce = Boolean(rootCompany?.is_produce);
    const isSubordinate = Boolean(
        rootCompany && activeCompany && activeCompany.id !== rootCompany.id,
    );

    return (
        <ScopeContext.Provider
            value={{
                activeCompany,
                mode,
                isSubordinate,
                rootCompany,
                canBloc,
                canProduce,
                subordinates,
                loadingSubordinates,
                isPlatformMode,
                impersonating: impersonatedTenant !== null,
                impersonatedTenant,
                setMode,
                enterCompany,
                exitToRoot,
                enterImpersonation,
                exitImpersonation,
                isReady,
            }}
        >
            {children}
        </ScopeContext.Provider>
    );
}

// ── Hook ──────────────────────────────────────────────────────────

export function useScope(): ScopeContextType {
    const ctx = useContext(ScopeContext);

    if (!ctx) {
        throw new Error('useScope must be used inside ScopeProvider');
    }

    return ctx;
}
