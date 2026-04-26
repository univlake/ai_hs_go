import { useEffect, useMemo, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import {
    Building2,
    GitBranch,
    Plus,
    Power,
    PowerOff,
    Search,
    ShieldCheck,
    UserCheck,
    Users,
    Wrench,
    AlertTriangle,
    LogIn,
    RefreshCw,
} from 'lucide-react';
import { toast } from 'sonner';
import api from '@/lib/api';
import { API } from '@/lib/api-routes';
import { useScope } from '@/contexts/ScopeContext';
import { Button } from '@/components/ui/button';
import { FormDialog, FormRow } from '@/components/shared/FormDialog';
import { Input } from '@/components/ui/input';

interface Tenant {
    id: number;
    title: string;
    status: number;
    is_bloc: number;
    is_produce: number;
    root_id: number;
    created_at: string;
    address?: string | null;
    credit_code?: string | null;
}

interface CreatedAdmin {
    id: number;
    username: string;
    must_change_credentials: boolean;
}

// ── Stat tile ─────────────────────────────────────────────────────

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
    accent: 'sky' | 'emerald' | 'amber' | 'violet';
    sub?: string;
}) {
    const tone = {
        sky: {
            border: 'border-sky-600/20',
            bg: 'bg-sky-600/15',
            text: 'text-sky-400',
        },
        emerald: {
            border: 'border-emerald-600/20',
            bg: 'bg-emerald-600/15',
            text: 'text-emerald-400',
        },
        amber: {
            border: 'border-amber-600/20',
            bg: 'bg-amber-600/15',
            text: 'text-amber-400',
        },
        violet: {
            border: 'border-violet-600/20',
            bg: 'bg-violet-600/15',
            text: 'text-violet-400',
        },
    }[accent];

    return (
        <div className="rounded-xl border border-blue-900/40 bg-[#0d1530] p-5">
            <div className="flex items-start justify-between">
                <div>
                    <p className="text-xs font-medium tracking-wide text-slate-500 uppercase">
                        {label}
                    </p>
                    <p className="mt-1 text-2xl font-bold text-slate-100">{value}</p>
                    {sub && (
                        <p className="mt-0.5 text-[11px] text-slate-600">{sub}</p>
                    )}
                </div>
                <div
                    className={`flex h-9 w-9 items-center justify-center rounded-lg border ${tone.border} ${tone.bg}`}
                >
                    <Icon className={`h-4 w-4 ${tone.text}`} />
                </div>
            </div>
        </div>
    );
}

// ── Status badge ──────────────────────────────────────────────────

function StatusBadge({ status }: { status: number }) {
    if (status === 1) {
        return (
            <span className="inline-flex items-center gap-1 rounded-full border border-emerald-500/30 bg-emerald-600/15 px-2 py-0.5 text-[10px] font-medium text-emerald-400">
                <span className="h-1.5 w-1.5 rounded-full bg-emerald-400" />
                启用中
            </span>
        );
    }
    if (status === 2) {
        return (
            <span className="inline-flex items-center gap-1 rounded-full border border-amber-500/30 bg-amber-600/15 px-2 py-0.5 text-[10px] font-medium text-amber-400">
                <PowerOff className="h-2.5 w-2.5" />
                已禁用
            </span>
        );
    }
    return (
        <span className="inline-flex items-center gap-1 rounded-full border border-slate-500/30 bg-slate-600/15 px-2 py-0.5 text-[10px] font-medium text-slate-400">
            <AlertTriangle className="h-2.5 w-2.5" />
            待审
        </span>
    );
}

function CapabilityBadge({ tenant }: { tenant: Tenant }) {
    const tags: { label: string; cls: string }[] = [];
    if (tenant.is_bloc === 1)
        tags.push({
            label: '集团',
            cls: 'border-blue-500/30 bg-blue-600/15 text-blue-400',
        });
    if (tenant.is_produce === 1)
        tags.push({
            label: '生产',
            cls: 'border-violet-500/30 bg-violet-600/15 text-violet-400',
        });
    if (tags.length === 0) return null;
    return (
        <div className="flex gap-1">
            {tags.map((t) => (
                <span
                    key={t.label}
                    className={`rounded-md border px-1.5 py-0.5 text-[10px] font-semibold tracking-wide uppercase ${t.cls}`}
                >
                    {t.label}
                </span>
            ))}
        </div>
    );
}

// ── Page ──────────────────────────────────────────────────────────

export default function PlatformTenants() {
    const navigate = useNavigate();
    const { enterImpersonation } = useScope();
    const [rows, setRows] = useState<Tenant[]>([]);
    const [loading, setLoading] = useState(true);
    const [open, setOpen] = useState(false);
    const [submitting, setSubmitting] = useState(false);
    const [title, setTitle] = useState('');
    const [createdAdmin, setCreatedAdmin] = useState<CreatedAdmin | null>(null);
    const [keyword, setKeyword] = useState('');

    const handleImpersonate = async (t: Tenant) => {
        if (t.status !== 1) {
            toast.error('该企业已禁用，无法化身');
            return;
        }
        try {
            await enterImpersonation({
                id: t.id,
                title: t.title,
                is_bloc: t.is_bloc,
                is_produce: t.is_produce,
                pid: 0,
                status: t.status,
            });
            toast.success(`已化身为「${t.title}」`);
            navigate('/dashboard');
        } catch (err) {
            const msg =
                (err as { response?: { data?: { message?: string } } }).response
                    ?.data?.message ?? '化身失败';
            toast.error(msg);
        }
    };

    const refresh = async () => {
        setLoading(true);
        try {
            const res = await api.get(API.tenants());
            setRows((res.data?.data as Tenant[]) ?? []);
        } catch (err) {
            toast.error('加载租户失败');
            // eslint-disable-next-line no-console
            console.error(err);
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        refresh();
    }, []);

    const create = async () => {
        if (!title.trim()) {
            toast.error('请输入企业名称');
            return;
        }
        setSubmitting(true);
        try {
            const res = await api.post(API.tenants(), {
                title: title.trim(),
                operational_type_id: 1,
                is_bloc: 1,
            });
            const data = res.data?.data ?? {};
            setCreatedAdmin(data.admin ?? null);
            setTitle('');
            setOpen(false);
            await refresh();
            toast.success('一级企业已创建');
        } catch (err) {
            const msg =
                (err as { response?: { data?: { message?: string } } }).response
                    ?.data?.message ?? '创建失败';
            toast.error(msg);
        } finally {
            setSubmitting(false);
        }
    };

    const toggleStatus = async (id: number) => {
        try {
            await api.patch(API.tenantToggleStatus(id));
            await refresh();
            toast.success('状态已切换');
        } catch (err) {
            toast.error('切换状态失败');
            // eslint-disable-next-line no-console
            console.error(err);
        }
    };

    // Stats
    const stats = useMemo(
        () => ({
            total: rows.length,
            enabled: rows.filter((t) => t.status === 1).length,
            disabled: rows.filter((t) => t.status === 2).length,
            bloc: rows.filter((t) => t.is_bloc === 1).length,
        }),
        [rows],
    );

    const filtered = useMemo(() => {
        const kw = keyword.trim().toLowerCase();
        if (!kw) return rows;
        return rows.filter(
            (r) =>
                r.title.toLowerCase().includes(kw) ||
                (r.credit_code?.toLowerCase().includes(kw) ?? false),
        );
    }, [rows, keyword]);

    return (
        <div className="space-y-6">
            {/* Header */}
            <div className="flex items-start justify-between">
                <div>
                    <h1 className="text-2xl font-bold text-slate-100">
                        一级企业管理
                    </h1>
                    <p className="mt-1 text-sm text-slate-500">
                        平台超管视角 · 管理所有一级企业（森林根）及其初始管理员
                    </p>
                </div>
                <div className="flex items-center gap-2">
                    <Button
                        variant="outline"
                        size="sm"
                        onClick={refresh}
                        className="border-blue-900/40 bg-[#0d1530]"
                    >
                        <RefreshCw className="mr-1.5 h-3.5 w-3.5" />
                        刷新
                    </Button>
                    <Button
                        onClick={() => setOpen(true)}
                        className="bg-blue-600 hover:bg-blue-500"
                    >
                        <Plus className="mr-1.5 h-3.5 w-3.5" />
                        新建一级企业
                    </Button>
                </div>
            </div>

            {/* Stat row */}
            <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
                <StatTile
                    label="一级企业总数"
                    value={stats.total}
                    icon={Building2}
                    accent="sky"
                    sub="森林中所有树根"
                />
                <StatTile
                    label="启用中"
                    value={stats.enabled}
                    icon={ShieldCheck}
                    accent="emerald"
                    sub="可正常登录使用"
                />
                <StatTile
                    label="已禁用"
                    value={stats.disabled}
                    icon={PowerOff}
                    accent="amber"
                    sub="登录被拒，仅超管可见"
                />
                <StatTile
                    label="集团类企业"
                    value={stats.bloc}
                    icon={GitBranch}
                    accent="violet"
                    sub="带有 is_bloc = 1 标记"
                />
            </div>

            {/* Search bar */}
            <div className="flex items-center gap-3">
                <div className="relative max-w-sm flex-1">
                    <Search className="absolute top-1/2 left-3 h-4 w-4 -translate-y-1/2 text-slate-600" />
                    <input
                        value={keyword}
                        onChange={(e) => setKeyword(e.target.value)}
                        placeholder="搜索企业名称 / 统一社会信用代码"
                        className="w-full rounded-lg border border-blue-900/40 bg-[#0d1530] py-2 pr-3 pl-9 text-sm text-slate-200 placeholder:text-slate-600 focus:border-blue-500/40 focus:outline-none"
                    />
                </div>
                <span className="text-xs text-slate-500">
                    共 {filtered.length} / {rows.length} 条
                </span>
            </div>

            {/* Table */}
            <div className="overflow-hidden rounded-xl border border-blue-900/40 bg-[#0d1530]">
                {loading ? (
                    <div className="px-5 py-12 text-center text-sm text-slate-500">
                        <RefreshCw className="mx-auto mb-3 h-5 w-5 animate-spin text-slate-600" />
                        加载中…
                    </div>
                ) : filtered.length === 0 ? (
                    <div className="px-5 py-12 text-center">
                        <Building2 className="mx-auto mb-3 h-10 w-10 text-slate-700" />
                        <p className="text-sm text-slate-400">
                            {rows.length === 0
                                ? '还没有一级企业'
                                : '未找到匹配的企业'}
                        </p>
                        {rows.length === 0 && (
                            <p className="mt-1 text-xs text-slate-600">
                                点击右上角「新建一级企业」开始
                            </p>
                        )}
                    </div>
                ) : (
                    <table className="w-full text-sm">
                        <thead>
                            <tr className="border-b border-blue-900/30 bg-[#0a1428] text-[10px] tracking-widest text-slate-600 uppercase">
                                <th className="px-5 py-3 text-left">企业</th>
                                <th className="px-5 py-3 text-left">能力</th>
                                <th className="px-5 py-3 text-left">状态</th>
                                <th className="px-5 py-3 text-left">root_id</th>
                                <th className="px-5 py-3 text-left">创建时间</th>
                                <th className="px-5 py-3 text-right">操作</th>
                            </tr>
                        </thead>
                        <tbody>
                            {filtered.map((t) => (
                                <tr
                                    key={t.id}
                                    className="border-b border-blue-900/20 transition-colors last:border-b-0 hover:bg-[#0a1428]"
                                >
                                    <td className="px-5 py-3">
                                        <div className="flex items-center gap-3">
                                            <div className="flex h-9 w-9 shrink-0 items-center justify-center rounded-lg border border-blue-600/20 bg-blue-600/15">
                                                <Building2 className="h-4 w-4 text-blue-400" />
                                            </div>
                                            <div className="min-w-0">
                                                <div className="truncate font-medium text-slate-100">
                                                    {t.title}
                                                </div>
                                                <div className="mt-0.5 flex items-center gap-2 text-[11px] text-slate-600">
                                                    <span>#{t.id}</span>
                                                    {t.credit_code && (
                                                        <>
                                                            <span>·</span>
                                                            <span className="font-mono">
                                                                {t.credit_code}
                                                            </span>
                                                        </>
                                                    )}
                                                </div>
                                            </div>
                                        </div>
                                    </td>
                                    <td className="px-5 py-3">
                                        <CapabilityBadge tenant={t} />
                                    </td>
                                    <td className="px-5 py-3">
                                        <StatusBadge status={t.status} />
                                    </td>
                                    <td className="px-5 py-3">
                                        <span className="rounded bg-[#0a1428] px-2 py-0.5 font-mono text-xs text-slate-400">
                                            {t.root_id}
                                        </span>
                                    </td>
                                    <td className="px-5 py-3 text-xs text-slate-500">
                                        {new Date(t.created_at).toLocaleString(
                                            'zh-CN',
                                            {
                                                year: 'numeric',
                                                month: '2-digit',
                                                day: '2-digit',
                                                hour: '2-digit',
                                                minute: '2-digit',
                                            },
                                        )}
                                    </td>
                                    <td className="px-5 py-3 text-right">
                                        <div className="flex items-center justify-end gap-1">
                                            <button
                                                onClick={() => handleImpersonate(t)}
                                                disabled={t.status !== 1}
                                                title={
                                                    t.status === 1
                                                        ? '以此企业身份查看（化身）'
                                                        : '企业已禁用，无法化身'
                                                }
                                                className="flex items-center gap-1 rounded-md border border-blue-500/20 bg-blue-600/10 px-2 py-1 text-[11px] text-blue-400 transition-colors hover:border-blue-500/40 hover:bg-blue-600/20 disabled:cursor-not-allowed disabled:opacity-40 disabled:hover:border-blue-500/20 disabled:hover:bg-blue-600/10"
                                            >
                                                <LogIn className="h-3 w-3" />
                                                以此身份查看
                                            </button>
                                            <button
                                                onClick={() => toggleStatus(t.id)}
                                                title={
                                                    t.status === 1
                                                        ? '禁用此企业'
                                                        : '启用此企业'
                                                }
                                                className={`flex items-center gap-1 rounded-md border px-2 py-1 text-[11px] transition-colors ${
                                                    t.status === 1
                                                        ? 'border-amber-500/20 bg-amber-600/10 text-amber-400 hover:border-amber-500/40 hover:bg-amber-600/20'
                                                        : 'border-emerald-500/20 bg-emerald-600/10 text-emerald-400 hover:border-emerald-500/40 hover:bg-emerald-600/20'
                                                }`}
                                            >
                                                {t.status === 1 ? (
                                                    <>
                                                        <PowerOff className="h-3 w-3" />
                                                        停用
                                                    </>
                                                ) : (
                                                    <>
                                                        <Power className="h-3 w-3" />
                                                        启用
                                                    </>
                                                )}
                                            </button>
                                        </div>
                                    </td>
                                </tr>
                            ))}
                        </tbody>
                    </table>
                )}
            </div>

            {/* Help footer */}
            <div className="rounded-lg border border-blue-900/30 bg-[#0a1428] p-4">
                <div className="flex gap-3 text-xs text-slate-500">
                    <Wrench className="h-4 w-4 shrink-0 text-slate-600" />
                    <div className="space-y-1">
                        <p>
                            创建一级企业会自动派生初始管理员账号
                            <code className="mx-1 rounded bg-[#0d1530] px-1.5 py-0.5 text-slate-300">
                                c{'{'}id{'}'}_admin
                            </code>
                            ，初始密码{' '}
                            <code className="rounded bg-[#0d1530] px-1.5 py-0.5 text-slate-300">
                                Aifix@2026
                            </code>
                            ，首次登录强制改用户名 + 密码。
                        </p>
                        <p>
                            <code>root_id</code> 由数据库触发器自动维护：一级企业的{' '}
                            <code>root_id = id</code>，子公司的{' '}
                            <code>root_id</code> = 它所在树根的 id。
                        </p>
                    </div>
                </div>
            </div>

            {/* Create dialog */}
            <FormDialog
                open={open}
                onOpenChange={setOpen}
                title="新建一级企业"
                size="md"
                onConfirm={create}
                confirmLabel="创建"
                loading={submitting}
            >
                <div className="space-y-4">
                    <FormRow
                        label="企业名称"
                        required
                        hint="将作为森林中一棵新树的根节点"
                    >
                        <Input
                            value={title}
                            onChange={(e) => setTitle(e.target.value)}
                            placeholder="例如：集团一公司"
                            autoFocus
                        />
                    </FormRow>
                    <div className="rounded-md border border-blue-900/30 bg-[#0a1428] p-3 text-[11px] leading-relaxed text-slate-500">
                        创建后自动派生初始管理员
                        <code className="mx-1 rounded bg-[#0d1530] px-1.5 py-0.5 text-slate-300">
                            c{'{'}id{'}'}_admin
                        </code>
                        ，初始密码{' '}
                        <code className="rounded bg-[#0d1530] px-1.5 py-0.5 text-slate-300">
                            Aifix@2026
                        </code>
                        ，强制首登修改。
                    </div>
                </div>
            </FormDialog>

            {/* Created admin reveal */}
            <FormDialog
                open={createdAdmin !== null}
                onOpenChange={(v) => !v && setCreatedAdmin(null)}
                title="初始管理员已创建"
                icon={UserCheck}
                size="md"
                cancelLabel=""
                onConfirm={() => setCreatedAdmin(null)}
                confirmLabel="知道了"
            >
                {createdAdmin && (
                    <div className="space-y-3 text-sm">
                        <div className="space-y-2 rounded-lg border border-emerald-500/20 bg-emerald-600/5 p-3">
                            <div className="flex items-center justify-between">
                                <span className="text-xs text-slate-500">用户名</span>
                                <code className="rounded bg-[#0d1530] px-2 py-0.5 text-emerald-400">
                                    {createdAdmin.username}
                                </code>
                            </div>
                            <div className="flex items-center justify-between">
                                <span className="text-xs text-slate-500">初始密码</span>
                                <code className="rounded bg-[#0d1530] px-2 py-0.5 text-emerald-400">
                                    Aifix@2026
                                </code>
                            </div>
                        </div>
                        <div className="flex gap-2 rounded-md border border-amber-500/20 bg-amber-600/5 p-3 text-xs text-amber-300/80">
                            <Users className="mt-0.5 h-3.5 w-3.5 shrink-0" />
                            <p>
                                请通知该管理员立即登录并修改用户名 +
                                密码。系统已设置{' '}
                                <code>must_change_credentials</code> flag，
                                首次登录会自动跳转到设置页。
                            </p>
                        </div>
                    </div>
                )}
            </FormDialog>
        </div>
    );
}
