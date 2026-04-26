import { useCallback, useEffect, useMemo, useState } from 'react';
import {
    KeyRound,
    Layers,
    RefreshCw,
    Search,
    ShieldCheck,
    Sparkles,
} from 'lucide-react';
import { toast } from 'sonner';
import api from '@/lib/api';
import { API } from '@/lib/api-routes';
import { Button } from '@/components/ui/button';

interface PermissionItem {
    id: number;
    name: string;
    guard_name: string;
}
interface PermissionsResponse {
    permissions: PermissionItem[];
    grouped: Record<string, string[]>;
}

interface RoleRow {
    role: { id: number; name: string };
    permissions: string[];
}

function StatTile({
    label,
    value,
    icon: Icon,
    accent,
    sub,
}: {
    label: string;
    value: number | string;
    icon: React.ComponentType<{ className?: string }>;
    accent: 'sky' | 'emerald' | 'violet';
    sub?: string;
}) {
    const tone = {
        sky: 'border-sky-600/20 bg-sky-600/15 text-sky-400',
        emerald: 'border-emerald-600/20 bg-emerald-600/15 text-emerald-400',
        violet: 'border-violet-600/20 bg-violet-600/15 text-violet-400',
    }[accent];
    return (
        <div className="rounded-xl border border-blue-900/40 bg-[#0d1530] p-5">
            <div className="flex items-start justify-between">
                <div>
                    <p className="text-xs font-medium tracking-wide text-slate-500 uppercase">
                        {label}
                    </p>
                    <p className="mt-1 text-2xl font-bold text-slate-100">
                        {value}
                    </p>
                    {sub && (
                        <p className="mt-0.5 text-[11px] text-slate-600">{sub}</p>
                    )}
                </div>
                <div
                    className={`flex h-9 w-9 items-center justify-center rounded-lg border ${tone}`}
                >
                    <Icon className="h-4 w-4" />
                </div>
            </div>
        </div>
    );
}

export default function PlatformPermissions() {
    const [data, setData] = useState<PermissionsResponse | null>(null);
    const [roles, setRoles] = useState<RoleRow[]>([]);
    const [loading, setLoading] = useState(true);
    const [keyword, setKeyword] = useState('');

    const refresh = useCallback(async () => {
        setLoading(true);
        try {
            const [permRes, roleRes] = await Promise.all([
                api.get(API.permissions()),
                api.get(API.roles()),
            ]);
            setData(permRes.data?.data as PermissionsResponse);
            setRoles((roleRes.data?.data as RoleRow[]) ?? []);
        } catch (err) {
            toast.error('加载权限目录失败');
            // eslint-disable-next-line no-console
            console.error(err);
        } finally {
            setLoading(false);
        }
    }, []);

    useEffect(() => {
        refresh();
    }, [refresh]);

    // Build a permission → roles index for the "used by" badges
    const usedBy = useMemo(() => {
        const map: Record<string, string[]> = {};
        for (const r of roles) {
            for (const p of r.permissions) {
                if (!map[p]) map[p] = [];
                map[p].push(r.role.name);
            }
        }
        return map;
    }, [roles]);

    const groupedFiltered = useMemo(() => {
        const grouped = data?.grouped ?? {};
        const kw = keyword.trim().toLowerCase();
        if (!kw) return grouped;
        const out: Record<string, string[]> = {};
        for (const [g, perms] of Object.entries(grouped)) {
            const filtered = perms.filter(
                (p) =>
                    p.toLowerCase().includes(kw) ||
                    g.toLowerCase().includes(kw),
            );
            if (filtered.length > 0) out[g] = filtered;
        }
        return out;
    }, [data?.grouped, keyword]);

    const groups = Object.keys(groupedFiltered).sort();
    const totalPerms = data?.permissions?.length ?? 0;
    const orphanCount =
        data?.permissions?.filter((p) => !usedBy[p.name]).length ?? 0;
    const visibleCount = Object.values(groupedFiltered).reduce(
        (s, arr) => s + arr.length,
        0,
    );

    return (
        <div className="space-y-6">
            <div className="flex items-start justify-between">
                <div>
                    <h1 className="text-2xl font-bold text-slate-100">权限目录</h1>
                    <p className="mt-1 text-sm text-slate-500">
                        系统提供的所有权限标识 ·
                        权限是代码侧定义的，要修改请到「角色管理」给角色分配
                    </p>
                </div>
                <Button
                    variant="outline"
                    size="sm"
                    onClick={refresh}
                    className="border-blue-900/40 bg-[#0d1530]"
                >
                    <RefreshCw className="mr-1.5 h-3.5 w-3.5" />
                    刷新
                </Button>
            </div>

            <div className="grid gap-4 md:grid-cols-3">
                <StatTile
                    label="权限总数"
                    value={totalPerms}
                    icon={KeyRound}
                    accent="sky"
                />
                <StatTile
                    label="模块（资源）数"
                    value={Object.keys(data?.grouped ?? {}).length}
                    icon={Layers}
                    accent="violet"
                    sub="按 resource.action 命名拆分"
                />
                <StatTile
                    label="未被任何角色使用"
                    value={orphanCount}
                    icon={Sparkles}
                    accent="emerald"
                    sub={
                        orphanCount === 0
                            ? '所有权限都已被分配'
                            : '可考虑分配给某个角色'
                    }
                />
            </div>

            {/* Search */}
            <div className="flex items-center gap-3">
                <div className="relative max-w-sm flex-1">
                    <Search className="absolute top-1/2 left-3 h-4 w-4 -translate-y-1/2 text-slate-600" />
                    <input
                        value={keyword}
                        onChange={(e) => setKeyword(e.target.value)}
                        placeholder="搜索权限名 / 模块名"
                        className="w-full rounded-lg border border-blue-900/40 bg-[#0d1530] py-2 pr-3 pl-9 text-sm text-slate-200 placeholder:text-slate-600 focus:border-blue-500/40 focus:outline-none"
                    />
                </div>
                <span className="text-xs text-slate-500">
                    共 {visibleCount} / {totalPerms} 项
                </span>
            </div>

            {/* Permission groups */}
            {loading ? (
                <div className="rounded-xl border border-blue-900/40 bg-[#0d1530] px-5 py-12 text-center text-sm text-slate-500">
                    <RefreshCw className="mx-auto mb-3 h-5 w-5 animate-spin text-slate-600" />
                    加载中…
                </div>
            ) : groups.length === 0 ? (
                <div className="rounded-xl border border-blue-900/40 bg-[#0d1530] px-5 py-12 text-center text-sm text-slate-500">
                    未找到匹配的权限
                </div>
            ) : (
                <div className="space-y-4">
                    {groups.map((g) => {
                        const perms = groupedFiltered[g];
                        return (
                            <div
                                key={g}
                                className="rounded-xl border border-blue-900/40 bg-[#0d1530]"
                            >
                                <div className="flex items-center justify-between border-b border-blue-900/30 bg-[#0a1428] px-5 py-3">
                                    <div className="flex items-center gap-2">
                                        <ShieldCheck className="h-4 w-4 text-blue-400" />
                                        <span className="font-mono font-semibold text-slate-200">
                                            {g}
                                        </span>
                                        <span className="rounded-md border border-blue-500/30 bg-blue-600/15 px-1.5 py-0.5 text-[10px] text-blue-300">
                                            {perms.length}
                                        </span>
                                    </div>
                                </div>
                                <div className="divide-y divide-blue-900/20">
                                    {perms.map((p) => {
                                        const usedRoles = usedBy[p] ?? [];
                                        const action = p.replace(g + '.', '');
                                        return (
                                            <div
                                                key={p}
                                                className="flex items-center gap-4 px-5 py-2.5 text-sm transition-colors hover:bg-[#0a1428]"
                                            >
                                                <div className="flex min-w-0 flex-1 items-center gap-2">
                                                    <code className="font-mono text-slate-300">
                                                        {p}
                                                    </code>
                                                    <span className="text-[10px] text-slate-700">
                                                        {action}
                                                    </span>
                                                </div>
                                                <div className="flex shrink-0 items-center gap-1.5">
                                                    {usedRoles.length === 0 ? (
                                                        <span className="rounded border border-slate-700/50 bg-slate-800/30 px-1.5 py-0.5 text-[10px] text-slate-600">
                                                            未分配
                                                        </span>
                                                    ) : (
                                                        usedRoles.map((r) => (
                                                            <span
                                                                key={r}
                                                                className="rounded border border-emerald-500/30 bg-emerald-600/15 px-1.5 py-0.5 text-[10px] font-medium text-emerald-400"
                                                            >
                                                                {r}
                                                            </span>
                                                        ))
                                                    )}
                                                </div>
                                            </div>
                                        );
                                    })}
                                </div>
                            </div>
                        );
                    })}
                </div>
            )}

            <div className="rounded-lg border border-blue-900/30 bg-[#0a1428] p-4">
                <div className="flex gap-3 text-xs text-slate-500">
                    <KeyRound className="h-4 w-4 shrink-0 text-slate-600" />
                    <p>
                        权限命名约定：<code>resource.action</code>
                        （例如 <code>users.list</code>、
                        <code>special-equipment.create</code>）。新增权限需要在
                        Go 后端的{' '}
                        <code>internal/service/seeder.go</code>
                        中加入并 <code>make seed</code> 重新写入。这里仅作目录展示。
                    </p>
                </div>
            </div>
        </div>
    );
}
