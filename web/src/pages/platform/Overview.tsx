import { useCallback, useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import {
    Activity,
    Building,
    Building2,
    GitBranch,
    KeyRound,
    LogIn,
    PowerOff,
    RefreshCw,
    ShieldAlert,
    ShieldCheck,
    UserCog,
    Users,
} from 'lucide-react';
import { toast } from 'sonner';
import api from '@/lib/api';
import { Button } from '@/components/ui/button';
import { useAuth } from '@/hooks/use-auth';

interface Overview {
    tenants_total: number;
    tenants_enabled: number;
    tenants_disabled: number;
    companies_total: number;
    users_total: number;
    users_platform: number;
    roles_total: number;
    permissions_total: number;
    recent_events: string[];
}

function StatCard({
    label,
    value,
    icon: Icon,
    accent,
    sub,
}: {
    label: string;
    value: number | string;
    icon: React.ComponentType<{ className?: string }>;
    accent: 'sky' | 'emerald' | 'amber' | 'violet' | 'rose';
    sub?: string;
}) {
    const tone = {
        sky: 'border-sky-600/20 bg-sky-600/15 text-sky-400',
        emerald: 'border-emerald-600/20 bg-emerald-600/15 text-emerald-400',
        amber: 'border-amber-600/20 bg-amber-600/15 text-amber-400',
        violet: 'border-violet-600/20 bg-violet-600/15 text-violet-400',
        rose: 'border-rose-600/20 bg-rose-600/15 text-rose-400',
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

export default function PlatformOverview() {
    const { user } = useAuth();
    const [data, setData] = useState<Overview | null>(null);
    const [loading, setLoading] = useState(true);

    const refresh = useCallback(async () => {
        setLoading(true);
        try {
            const res = await api.get('/platform/overview');
            setData(res.data?.data as Overview);
        } catch (err) {
            toast.error('加载平台概览失败');
            // eslint-disable-next-line no-console
            console.error(err);
        } finally {
            setLoading(false);
        }
    }, []);

    useEffect(() => {
        refresh();
    }, [refresh]);

    return (
        <div className="space-y-6">
            <div className="flex items-start justify-between">
                <div>
                    <h1 className="text-2xl font-bold text-slate-100">
                        平台总览
                    </h1>
                    <p className="mt-1 text-sm text-slate-500">
                        欢迎回来，
                        <span className="text-slate-300">{user?.nickname}</span>
                        。这里展示跨租户的全局指标。要管理具体租户的业务，请到
                        「一级企业」选择并化身。
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

            {loading || !data ? (
                <div className="rounded-xl border border-blue-900/40 bg-[#0d1530] px-5 py-12 text-center text-sm text-slate-500">
                    <RefreshCw className="mx-auto mb-3 h-5 w-5 animate-spin text-slate-600" />
                    加载中…
                </div>
            ) : (
                <>
                    {/* Tenant stats */}
                    <div>
                        <h2 className="mb-3 text-xs font-semibold tracking-widest text-slate-500 uppercase">
                            租户
                        </h2>
                        <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
                            <StatCard
                                label="一级企业总数"
                                value={data.tenants_total}
                                icon={Building2}
                                accent="sky"
                                sub="森林中的所有树根"
                            />
                            <StatCard
                                label="启用中"
                                value={data.tenants_enabled}
                                icon={ShieldCheck}
                                accent="emerald"
                                sub="可正常登录使用"
                            />
                            <StatCard
                                label="已禁用"
                                value={data.tenants_disabled}
                                icon={PowerOff}
                                accent="amber"
                                sub="登录被拒，待启用"
                            />
                            <StatCard
                                label="企业总数（含子公司）"
                                value={data.companies_total}
                                icon={Building}
                                accent="violet"
                                sub="所有节点（含子树）"
                            />
                        </div>
                    </div>

                    {/* User & RBAC stats */}
                    <div>
                        <h2 className="mb-3 text-xs font-semibold tracking-widest text-slate-500 uppercase">
                            用户与权限
                        </h2>
                        <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
                            <StatCard
                                label="用户总数"
                                value={data.users_total}
                                icon={Users}
                                accent="sky"
                                sub="跨所有租户（含平台管理员）"
                            />
                            <StatCard
                                label="平台管理员"
                                value={data.users_platform}
                                icon={ShieldAlert}
                                accent="rose"
                                sub="is_platform_admin = true"
                            />
                            <StatCard
                                label="角色数量"
                                value={data.roles_total}
                                icon={UserCog}
                                accent="violet"
                                sub="含内置 + 自定义"
                            />
                            <StatCard
                                label="可用权限"
                                value={data.permissions_total}
                                icon={KeyRound}
                                accent="emerald"
                                sub="平台级权限目录"
                            />
                        </div>
                    </div>

                    {/* Quick links */}
                    <div className="grid gap-4 lg:grid-cols-2">
                        <div className="rounded-xl border border-blue-900/40 bg-[#0d1530] p-5">
                            <div className="mb-3 flex items-center justify-between">
                                <h3 className="text-sm font-semibold text-slate-200">
                                    快速操作
                                </h3>
                            </div>
                            <div className="space-y-2 text-sm">
                                <QuickLink
                                    to="/platform/tenants"
                                    icon={Building2}
                                    label="管理一级企业"
                                    desc="新建 / 启停 / 化身查看"
                                />
                                <QuickLink
                                    to="/platform/forest"
                                    icon={GitBranch}
                                    label="企业森林"
                                    desc="跨租户拓扑 + 子公司管理"
                                />
                                <QuickLink
                                    to="/platform/permissions"
                                    icon={KeyRound}
                                    label="权限目录"
                                    desc="查看系统提供的所有权限"
                                />
                            </div>
                        </div>

                        <div className="rounded-xl border border-blue-900/40 bg-[#0d1530] p-5">
                            <div className="mb-3 flex items-center gap-2">
                                <Activity className="h-4 w-4 text-blue-400" />
                                <h3 className="text-sm font-semibold text-slate-200">
                                    最近平台活动
                                </h3>
                                <span className="ml-auto text-[11px] text-slate-600">
                                    最近 5 条 platform.* 审计
                                </span>
                            </div>
                            {data.recent_events &&
                            data.recent_events.length > 0 ? (
                                <ul className="space-y-1.5 text-xs text-slate-400">
                                    {data.recent_events.map((ev, i) => (
                                        <li
                                            key={i}
                                            className="flex items-center gap-2 rounded border border-blue-900/20 bg-[#0a1428] px-2 py-1.5"
                                        >
                                            {ev.startsWith('platform.impersonate.enter') ? (
                                                <LogIn className="h-3 w-3 shrink-0 text-blue-400" />
                                            ) : (
                                                <Activity className="h-3 w-3 shrink-0 text-slate-500" />
                                            )}
                                            <span className="truncate font-mono">
                                                {ev}
                                            </span>
                                        </li>
                                    ))}
                                </ul>
                            ) : (
                                <p className="rounded border border-blue-900/20 bg-[#0a1428] px-3 py-4 text-center text-xs text-slate-600">
                                    暂无平台活动
                                </p>
                            )}
                        </div>
                    </div>
                </>
            )}
        </div>
    );
}

function QuickLink({
    to,
    icon: Icon,
    label,
    desc,
}: {
    to: string;
    icon: React.ComponentType<{ className?: string }>;
    label: string;
    desc: string;
}) {
    return (
        <Link
            to={to}
            className="group flex items-center gap-3 rounded-lg border border-blue-900/30 bg-[#0a1428] px-3 py-2.5 transition-colors hover:border-blue-500/40 hover:bg-blue-600/8"
        >
            <div className="flex h-8 w-8 shrink-0 items-center justify-center rounded-lg border border-blue-600/20 bg-blue-600/15">
                <Icon className="h-3.5 w-3.5 text-blue-400" />
            </div>
            <div className="flex-1">
                <div className="text-sm font-medium text-slate-200 group-hover:text-slate-100">
                    {label}
                </div>
                <div className="text-[11px] text-slate-600">{desc}</div>
            </div>
            <span className="text-slate-700 group-hover:text-slate-500">→</span>
        </Link>
    );
}
