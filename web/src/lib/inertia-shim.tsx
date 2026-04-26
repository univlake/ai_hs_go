// Inertia → React Router compatibility shim.
//
// vite.config.ts aliases the bare `@inertiajs/react` import to this module so
// pages copied from the Laravel project continue to compile without per-file
// rewrites. The shim exposes the Inertia surface area we actually use:
//
//   Head     — sets document.title; ignores meta children
//   Link     — wraps react-router-dom <Link>, mapping `href` → `to`
//   router   — visit / replace / get / post backed by react-router or fetch
//   usePage  — returns `{ url, props }` derived from window.location
//   useForm  — minimal wrapper around an internal state machine that mirrors
//              Inertia's `data / setData / post / put / errors / processing`
//              signature so existing forms keep working until rewritten.
//
// As pages are migrated to native react-router-dom + axios, drop their
// `@inertiajs/react` import and the shim usage decreases organically.

import {
    Link as RRLink,
    useLocation,
    useNavigate,
} from 'react-router-dom';
import {
    forwardRef,
    useCallback,
    useEffect,
    useMemo,
    useState,
    type AnchorHTMLAttributes,
    type ReactNode,
} from 'react';
import api from '@/lib/api';

// ── Head ─────────────────────────────────────────────────────────────────

export interface HeadProps {
    title?: string;
    children?: ReactNode;
}

export function Head({ title }: HeadProps) {
    useEffect(() => {
        if (title) {
            document.title = title;
        }
    }, [title]);
    return null;
}

// ── Link ─────────────────────────────────────────────────────────────────

interface InertiaLinkProps
    extends Omit<AnchorHTMLAttributes<HTMLAnchorElement>, 'href'> {
    href: string;
    method?: 'get' | 'post' | 'put' | 'patch' | 'delete';
    as?: 'a' | 'button';
    preserveScroll?: boolean;
    preserveState?: boolean;
    replace?: boolean;
    only?: string[];
    data?: Record<string, unknown>;
}

export const Link = forwardRef<HTMLAnchorElement, InertiaLinkProps>(
    function InertiaCompatLink(
        { href, method, as: _as, preserveScroll, preserveState, replace, only, data, ...rest },
        ref,
    ) {
        // Drop Inertia-only props that react-router-dom doesn't understand.
        void method;
        void _as;
        void preserveScroll;
        void preserveState;
        void only;
        void data;
        return <RRLink ref={ref} to={href} replace={replace} {...rest} />;
    },
);

// ── router ───────────────────────────────────────────────────────────────

type Method = 'get' | 'post' | 'put' | 'patch' | 'delete';
type VisitOptions = {
    replace?: boolean;
    method?: Method;
    data?: Record<string, unknown>;
    onSuccess?: (page: PageObject) => void;
    onError?: (errors: Record<string, string>) => void;
    onFinish?: () => void;
    preserveScroll?: boolean;
    preserveState?: boolean;
};

interface PageObject {
    component: string;
    props: Record<string, unknown>;
    url: string;
}

// Imperative navigation outside React components (e.g. event handlers in
// pages). Internal navigation goes through window.history; absolute URLs
// fall through to a hard navigation. POST/PUT/etc. are routed through axios
// against the same path.
export const router = {
    visit(url: string, options: VisitOptions = {}) {
        const method = options.method ?? 'get';
        if (method === 'get') {
            if (options.replace) {
                window.history.replaceState({}, '', url);
            } else {
                window.history.pushState({}, '', url);
            }
            window.dispatchEvent(new PopStateEvent('popstate'));
            options.onFinish?.();
            return;
        }
        api.request({
            url,
            method,
            data: options.data,
        })
            .then((res) =>
                options.onSuccess?.({
                    component: '',
                    props: res.data ?? {},
                    url,
                }),
            )
            .catch((err) =>
                options.onError?.(
                    (err.response?.data?.errors as Record<string, string>) ?? {},
                ),
            )
            .finally(() => options.onFinish?.());
    },
    get(url: string, data?: Record<string, unknown>, options: VisitOptions = {}) {
        return router.visit(url, { ...options, method: 'get', data });
    },
    post(url: string, data?: Record<string, unknown>, options: VisitOptions = {}) {
        return router.visit(url, { ...options, method: 'post', data });
    },
    put(url: string, data?: Record<string, unknown>, options: VisitOptions = {}) {
        return router.visit(url, { ...options, method: 'put', data });
    },
    patch(url: string, data?: Record<string, unknown>, options: VisitOptions = {}) {
        return router.visit(url, { ...options, method: 'patch', data });
    },
    delete(url: string, options: VisitOptions = {}) {
        return router.visit(url, { ...options, method: 'delete' });
    },
    reload(options?: { only?: string[] }) {
        void options;
        window.location.reload();
    },
    cancelAll() {
        // no-op — kept for source compat with Inertia v3 callers
    },
    on(_event: string, _cb: (...args: unknown[]) => void) {
        // no-op event registration; returns an unsubscribe function for parity
        return () => undefined;
    },
};

// ── usePage ──────────────────────────────────────────────────────────────

export function usePage<T extends Record<string, unknown> = Record<string, unknown>>(): {
    url: string;
    props: T;
    component: string;
    version: string | null;
} {
    const location = useLocation();
    return {
        url: location.pathname + location.search,
        props: {} as T,
        component: '',
        version: null,
    };
}

// ── useForm ──────────────────────────────────────────────────────────────

type FormData = Record<string, unknown>;
type FormErrors = Record<string, string>;

export interface InertiaFormReturn<T extends FormData> {
    data: T;
    setData: <K extends keyof T>(key: K | T, value?: T[K]) => void;
    errors: FormErrors;
    processing: boolean;
    progress: { percentage: number } | null;
    isDirty: boolean;
    wasSuccessful: boolean;
    recentlySuccessful: boolean;
    transform: (cb: (data: T) => T) => void;
    submit: (method: Method, url: string, options?: VisitOptions) => void;
    get: (url: string, options?: VisitOptions) => void;
    post: (url: string, options?: VisitOptions) => void;
    put: (url: string, options?: VisitOptions) => void;
    patch: (url: string, options?: VisitOptions) => void;
    delete: (url: string, options?: VisitOptions) => void;
    reset: (...fields: (keyof T)[]) => void;
    clearErrors: (...fields: (keyof T)[]) => void;
    setError: (key: keyof T | FormErrors, value?: string) => void;
}

export function useForm<T extends FormData>(initial: T): InertiaFormReturn<T> {
    const [data, setDataState] = useState<T>(initial);
    const [errors, setErrors] = useState<FormErrors>({});
    const [processing, setProcessing] = useState(false);
    const [wasSuccessful, setWasSuccessful] = useState(false);
    const [recentlySuccessful, setRecentlySuccessful] = useState(false);
    const initialRef = useMemo(() => ({ ...initial }), []); // eslint-disable-line react-hooks/exhaustive-deps

    const setData = useCallback(
        <K extends keyof T>(key: K | T, value?: T[K]) => {
            if (typeof key === 'object') {
                setDataState(key as T);
            } else {
                setDataState((prev) => ({ ...prev, [key]: value as T[K] }));
            }
        },
        [],
    );

    const submit = useCallback(
        (method: Method, url: string, options: VisitOptions = {}) => {
            setProcessing(true);
            setErrors({});
            api
                .request({ url, method, data })
                .then((res) => {
                    setWasSuccessful(true);
                    setRecentlySuccessful(true);
                    setTimeout(() => setRecentlySuccessful(false), 2000);
                    options.onSuccess?.({
                        component: '',
                        props: res.data ?? {},
                        url,
                    });
                })
                .catch((err) => {
                    const e = (err.response?.data?.errors as FormErrors) ?? {};
                    setErrors(e);
                    options.onError?.(e);
                })
                .finally(() => {
                    setProcessing(false);
                    options.onFinish?.();
                });
        },
        [data],
    );

    return {
        data,
        setData,
        errors,
        processing,
        progress: null,
        isDirty: JSON.stringify(data) !== JSON.stringify(initialRef),
        wasSuccessful,
        recentlySuccessful,
        transform: (cb) => setDataState((prev) => cb(prev)),
        submit,
        get: (url, options) => submit('get', url, options),
        post: (url, options) => submit('post', url, options),
        put: (url, options) => submit('put', url, options),
        patch: (url, options) => submit('patch', url, options),
        delete: (url, options) => submit('delete', url, options),
        reset: (...fields) => {
            if (fields.length === 0) {
                setDataState({ ...initialRef });
            } else {
                setDataState((prev) => {
                    const next = { ...prev };
                    for (const f of fields) {
                        next[f] = initialRef[f];
                    }
                    return next;
                });
            }
        },
        clearErrors: (...fields) => {
            if (fields.length === 0) {
                setErrors({});
            } else {
                setErrors((prev) => {
                    const next = { ...prev };
                    for (const f of fields) {
                        delete next[String(f)];
                    }
                    return next;
                });
            }
        },
        setError: (key, value) => {
            if (typeof key === 'object') {
                setErrors(key);
            } else {
                setErrors((prev) => ({ ...prev, [String(key)]: value ?? '' }));
            }
        },
    };
}

// Internal navigation helper for shim callers that need a hook directly.
export function useShimNavigate() {
    return useNavigate();
}

// Convenience: keep `default` export shape consistent with Inertia for
// `import inertia from '@inertiajs/react'` style callers (uncommon).
const inertia = { Head, Link, router, usePage, useForm };
export default inertia;
