import { useCallback, useEffect, useMemo, useState } from 'react';
import {
    Edit3,
    KeyRound,
    RefreshCw,
    Search,
    ShieldCheck,
    Trash2,
    UserCog,
    UserPlus,
    Users as UsersIcon,
} from 'lucide-react';
import { toast } from 'sonner';
import api from '@/lib/api';
import { API } from '@/lib/api-routes';
import { Button } from '@/components/ui/button';
import { ActionBtn } from '@/components/shared/ActionBtn';
import { FormDialog, FormRow } from '@/components/shared/FormDialog';
import { Input } from '@/components/ui/input';
import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from '@/components/ui/select';
import { useAuth } from '@/hooks/use-auth';

interface UserRow {
    id: number;
    username: string;
    nickname: string;
    company_id: number | null;
    phone: string | null;
    email: string | null;
    department_name: string | null;
    status: number;
    is_expert: number;
    must_change_credentials: boolean;
    created_at: string;
    roles: string[];
}

interface RoleRow {
    role: { id: number; name: string };
    permissions: string[];
}

// 跟 aifix_hs 项目页面（resources/js/pages/system/users/Index.tsx）的按钮样式对齐：
//   primaryBtn — 主操作（新增、保存）边框 + 半透明蓝底
//   ghostBtn   — 次操作（取消、刷新）通透蓝灰边框
const primaryBtn =
    'flex items-center gap-2 rounded-lg border border-blue-500/30 bg-blue-600/15 px-4 py-2 text-sm font-medium text-blue-300 transition-all hover:border-blue-400/50 hover:bg-blue-600/25 hover:text-blue-200 active:scale-95';
const ghostBtn =
    'flex items-center gap-2 rounded-lg border border-blue-900/40 px-4 py-2 text-sm text-slate-400 transition-colors hover:border-blue-500/40 hover:text-slate-200';

const statusOptions = [
    { value: '', label: '全部状态' },
    { value: '1', label: '启用' },
    { value: '2', label: '禁用' },
];

function StatusBadge({ status }: { status: number }) {
    if (status === 1) {
        return (
            <span className="inline-flex items-center gap-1 rounded-full border border-emerald-500/30 bg-emerald-600/15 px-2 py-0.5 text-[10px] text-emerald-400">
                <span className="h-1.5 w-1.5 rounded-full bg-emerald-400" />
                启用
            </span>
        );
    }
    return (
        <span className="inline-flex items-center gap-1 rounded-full border border-amber-500/30 bg-amber-600/15 px-2 py-0.5 text-[10px] text-amber-400">
            禁用
        </span>
    );
}

function RoleBadges({ roles }: { roles: string[] }) {
    if (!roles || roles.length === 0)
        return <span className="text-[11px] text-slate-600">未分配</span>;
    return (
        <div className="flex flex-wrap gap-1">
            {roles.map((r) => (
                <span
                    key={r}
                    className="rounded-md border border-blue-500/30 bg-blue-600/15 px-1.5 py-0.5 text-[10px] font-medium text-blue-300"
                >
                    {r}
                </span>
            ))}
        </div>
    );
}

export default function SystemUsers() {
    const { hasPermission } = useAuth();

    const [rows, setRows] = useState<UserRow[]>([]);
    const [roleOptions, setRoleOptions] = useState<RoleRow[]>([]);
    const [loading, setLoading] = useState(true);
    const [page, setPage] = useState(1);
    const [perPage] = useState(20);
    const [total, setTotal] = useState(0);

    const [keyword, setKeyword] = useState('');
    const [status, setStatus] = useState('');
    const [role, setRole] = useState('');

    const [editing, setEditing] = useState<UserRow | null>(null);
    const [creating, setCreating] = useState(false);
    const [resetTarget, setResetTarget] = useState<UserRow | null>(null);
    const [submitting, setSubmitting] = useState(false);

    const canCreate = hasPermission('users.create');
    const canUpdate = hasPermission('users.update');
    const canDelete = hasPermission('users.delete');
    const canAssignRole = hasPermission('roles.assign');

    const refresh = useCallback(async () => {
        setLoading(true);
        try {
            const res = await api.get(API.users(), {
                params: {
                    page,
                    per_page: perPage,
                    keyword: keyword || undefined,
                    status: status || undefined,
                    role: role || undefined,
                },
            });
            const d = res.data?.data ?? {};
            setRows(d.data ?? []);
            setTotal(d.total ?? 0);
        } catch (err) {
            toast.error('加载用户失败');
            // eslint-disable-next-line no-console
            console.error(err);
        } finally {
            setLoading(false);
        }
    }, [page, perPage, keyword, status, role]);

    const loadRoles = useCallback(async () => {
        try {
            const res = await api.get(API.roles());
            setRoleOptions((res.data?.data as RoleRow[]) ?? []);
        } catch {
            // ignore — assignment dropdown will fall back to empty
        }
    }, []);

    useEffect(() => {
        refresh();
    }, [refresh]);
    useEffect(() => {
        loadRoles();
    }, [loadRoles]);

    const stats = useMemo(
        () => ({
            total,
            enabled: rows.filter((r) => r.status === 1).length,
            withRole: rows.filter((r) => r.roles && r.roles.length > 0).length,
        }),
        [rows, total],
    );

    return (
        <div className="space-y-6">
            <div className="flex items-center justify-between">
                <div>
                    <h1 className="text-xl font-bold text-slate-100">账号管理</h1>
                    <p className="mt-0.5 text-xs text-slate-600">
                        共{' '}
                        <span className="font-medium text-slate-400">
                            {total}
                        </span>{' '}
                        个用户
                    </p>
                </div>
                <div className="flex gap-2">
                    <button onClick={refresh} className={ghostBtn}>
                        <RefreshCw className="h-4 w-4" />
                        刷新
                    </button>
                    {canCreate && (
                        <button
                            onClick={() => setCreating(true)}
                            className={primaryBtn}
                        >
                            <UserPlus className="h-4 w-4" />
                            新增用户
                        </button>
                    )}
                </div>
            </div>

            <div className="grid gap-4 md:grid-cols-3">
                <StatCard
                    label="总用户数"
                    value={stats.total}
                    icon={UsersIcon}
                    color="sky"
                />
                <StatCard
                    label="启用中（本页）"
                    value={stats.enabled}
                    icon={ShieldCheck}
                    color="emerald"
                />
                <StatCard
                    label="已分配角色（本页）"
                    value={stats.withRole}
                    icon={UserCog}
                    color="violet"
                />
            </div>

            <div className="flex flex-wrap items-center gap-3">
                <div className="relative max-w-sm min-w-[200px] flex-1">
                    <Search className="absolute top-1/2 left-3 h-4 w-4 -translate-y-1/2 text-slate-600" />
                    <input
                        value={keyword}
                        onChange={(e) => {
                            setKeyword(e.target.value);
                            setPage(1);
                        }}
                        placeholder="搜索用户名 / 昵称 / 手机"
                        className="w-full rounded-lg border border-blue-900/40 bg-[#0d1530] py-2 pr-3 pl-9 text-sm text-slate-200 placeholder:text-slate-600 focus:border-blue-500/40 focus:outline-none"
                    />
                </div>
                <select
                    value={status}
                    onChange={(e) => {
                        setStatus(e.target.value);
                        setPage(1);
                    }}
                    className="rounded-lg border border-blue-900/40 bg-[#0d1530] px-3 py-2 text-sm text-slate-200"
                >
                    {statusOptions.map((o) => (
                        <option key={o.value} value={o.value}>
                            {o.label}
                        </option>
                    ))}
                </select>
                <select
                    value={role}
                    onChange={(e) => {
                        setRole(e.target.value);
                        setPage(1);
                    }}
                    className="rounded-lg border border-blue-900/40 bg-[#0d1530] px-3 py-2 text-sm text-slate-200"
                >
                    <option value="">全部角色</option>
                    {roleOptions.map((r) => (
                        <option key={r.role.id} value={r.role.name}>
                            {r.role.name}
                        </option>
                    ))}
                </select>
                <span className="text-xs text-slate-500">共 {total} 条</span>
            </div>

            <div className="overflow-hidden rounded-xl border border-blue-900/40 bg-[#0d1530]">
                {loading ? (
                    <div className="px-5 py-12 text-center text-sm text-slate-500">
                        <RefreshCw className="mx-auto mb-3 h-5 w-5 animate-spin text-slate-600" />
                        加载中…
                    </div>
                ) : rows.length === 0 ? (
                    <div className="px-5 py-12 text-center">
                        <UsersIcon className="mx-auto mb-3 h-10 w-10 text-slate-700" />
                        <p className="text-sm text-slate-400">未找到用户</p>
                    </div>
                ) : (
                    <table className="w-full text-sm">
                        <thead>
                            <tr className="border-b border-blue-900/30 bg-[#0a1428] text-[10px] tracking-widest text-slate-600 uppercase">
                                <th className="px-5 py-3 text-left">用户</th>
                                <th className="px-5 py-3 text-left">联系方式</th>
                                <th className="px-5 py-3 text-left">部门</th>
                                <th className="px-5 py-3 text-left">角色</th>
                                <th className="px-5 py-3 text-left">状态</th>
                                <th className="px-5 py-3 text-right">操作</th>
                            </tr>
                        </thead>
                        <tbody>
                            {rows.map((u) => (
                                <tr
                                    key={u.id}
                                    className="border-b border-blue-900/20 transition-colors last:border-b-0 hover:bg-[#0a1428]"
                                >
                                    <td className="px-5 py-3">
                                        <div className="flex items-center gap-3">
                                            <div className="flex h-9 w-9 shrink-0 items-center justify-center rounded-lg border border-blue-600/20 bg-blue-600/15 text-xs font-semibold text-blue-300">
                                                {(u.nickname || u.username || '?')
                                                    .slice(0, 1)
                                                    .toUpperCase()}
                                            </div>
                                            <div className="min-w-0">
                                                <div className="flex items-center gap-2">
                                                    <span className="font-medium text-slate-100">
                                                        {u.nickname || u.username}
                                                    </span>
                                                    {u.must_change_credentials && (
                                                        <span className="rounded border border-amber-500/30 bg-amber-600/15 px-1 py-0.5 text-[9px] text-amber-400">
                                                            首登未完成
                                                        </span>
                                                    )}
                                                </div>
                                                <div className="text-[11px] text-slate-600">
                                                    @{u.username} · #{u.id}
                                                </div>
                                            </div>
                                        </div>
                                    </td>
                                    <td className="px-5 py-3 text-xs text-slate-400">
                                        <div>{u.phone || '—'}</div>
                                        <div className="text-slate-600">
                                            {u.email || '—'}
                                        </div>
                                    </td>
                                    <td className="px-5 py-3 text-xs text-slate-400">
                                        {u.department_name || '—'}
                                    </td>
                                    <td className="px-5 py-3">
                                        <RoleBadges roles={u.roles ?? []} />
                                    </td>
                                    <td className="px-5 py-3">
                                        <StatusBadge status={u.status} />
                                    </td>
                                    <td className="px-5 py-3 text-right">
                                        <div className="flex items-center justify-end gap-0.5">
                                            {canUpdate && (
                                                <ActionBtn
                                                    onClick={() => setEditing(u)}
                                                    title="编辑"
                                                    variant="edit"
                                                >
                                                    <Edit3 className="h-3.5 w-3.5" />
                                                </ActionBtn>
                                            )}
                                            {canUpdate && (
                                                <ActionBtn
                                                    onClick={() => setResetTarget(u)}
                                                    title="重置密码"
                                                    variant="edit"
                                                >
                                                    <KeyRound className="h-3.5 w-3.5" />
                                                </ActionBtn>
                                            )}
                                            {canDelete && (
                                                <DeleteButton
                                                    user={u}
                                                    onDeleted={refresh}
                                                />
                                            )}
                                        </div>
                                    </td>
                                </tr>
                            ))}
                        </tbody>
                    </table>
                )}
            </div>

            {total > perPage && (
                <div className="flex items-center justify-between text-xs text-slate-500">
                    <span>
                        第 {page} / {Math.ceil(total / perPage)} 页 · 共 {total} 条
                    </span>
                    <div className="flex gap-1">
                        <Button
                            variant="outline"
                            size="sm"
                            disabled={page <= 1}
                            onClick={() => setPage((p) => p - 1)}
                            className="border-blue-900/40 bg-[#0d1530]"
                        >
                            上一页
                        </Button>
                        <Button
                            variant="outline"
                            size="sm"
                            disabled={page >= Math.ceil(total / perPage)}
                            onClick={() => setPage((p) => p + 1)}
                            className="border-blue-900/40 bg-[#0d1530]"
                        >
                            下一页
                        </Button>
                    </div>
                </div>
            )}

            {creating && (
                <CreateUserDialog
                    roles={roleOptions}
                    onClose={() => setCreating(false)}
                    onCreated={() => {
                        setCreating(false);
                        refresh();
                    }}
                    submitting={submitting}
                    setSubmitting={setSubmitting}
                />
            )}
            {editing && (
                <EditUserDialog
                    user={editing}
                    roles={roleOptions}
                    canAssignRole={canAssignRole}
                    onClose={() => setEditing(null)}
                    onSaved={() => {
                        setEditing(null);
                        refresh();
                    }}
                />
            )}
            {resetTarget && (
                <ResetPasswordDialog
                    user={resetTarget}
                    onClose={() => setResetTarget(null)}
                    onSaved={() => setResetTarget(null)}
                />
            )}
        </div>
    );
}

// ── Stat card ─────────────────────────────────────────────────────

function StatCard({
    label,
    value,
    icon: Icon,
    color,
}: {
    label: string;
    value: number;
    icon: React.ComponentType<{ className?: string }>;
    color: 'sky' | 'emerald' | 'violet';
}) {
    const tone = {
        sky: 'border-sky-600/20 bg-sky-600/15 text-sky-400',
        emerald: 'border-emerald-600/20 bg-emerald-600/15 text-emerald-400',
        violet: 'border-violet-600/20 bg-violet-600/15 text-violet-400',
    }[color];
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

// ── Delete confirm ────────────────────────────────────────────────

function DeleteButton({
    user,
    onDeleted,
}: {
    user: UserRow;
    onDeleted: () => void;
}) {
    const [open, setOpen] = useState(false);
    const [busy, setBusy] = useState(false);
    const handle = async () => {
        setBusy(true);
        try {
            await api.delete(API.userById(user.id));
            toast.success('已删除');
            onDeleted();
        } catch (err) {
            toast.error('删除失败');
            // eslint-disable-next-line no-console
            console.error(err);
        } finally {
            setBusy(false);
            setOpen(false);
        }
    };
    return (
        <>
            <ActionBtn onClick={() => setOpen(true)} title="删除" variant="delete">
                <Trash2 className="h-3.5 w-3.5" />
            </ActionBtn>
            <FormDialog
                open={open}
                onOpenChange={setOpen}
                title={`删除用户「${user.username}」`}
                icon={Trash2}
                size="md"
                onConfirm={handle}
                confirmLabel="确认删除"
                confirmTone="rose"
                loading={busy}
            >
                <p className="text-sm leading-relaxed text-slate-400">
                    软删除：用户将被标记为已删除，不会真正从数据库移除。
                    如有必要可联系运维恢复。
                </p>
            </FormDialog>
        </>
    );
}

// ── Create dialog ─────────────────────────────────────────────────

function CreateUserDialog({
    roles,
    onClose,
    onCreated,
    submitting,
    setSubmitting,
}: {
    roles: RoleRow[];
    onClose: () => void;
    onCreated: () => void;
    submitting: boolean;
    setSubmitting: (v: boolean) => void;
}) {
    const [form, setForm] = useState({
        username: '',
        password: '',
        nickname: '',
        phone: '',
        email: '',
        department_name: '',
        role: '',
    });

    const submit = async () => {
        if (!form.username.trim() || !form.password || !form.nickname.trim()) {
            toast.error('用户名、密码、昵称是必填');
            return;
        }
        setSubmitting(true);
        try {
            await api.post(API.users(), {
                ...form,
                phone: form.phone || null,
                email: form.email || null,
                department_name: form.department_name || null,
                role: form.role || undefined,
            });
            toast.success('用户已创建');
            onCreated();
        } catch (err) {
            const msg =
                (err as { response?: { data?: { message?: string } } }).response
                    ?.data?.message ?? '创建失败';
            toast.error(msg);
        } finally {
            setSubmitting(false);
        }
    };

    return (
        <FormDialog
            open
            onOpenChange={(v) => !v && onClose()}
            title="新建用户"
            description="新用户会被加入到当前租户，初始密码请告知本人，建议首次登录后立即修改"
            size="lg"
            onConfirm={submit}
            confirmLabel="创建"
            loading={submitting}
        >
            <div className="space-y-4">
                <div className="grid grid-cols-2 gap-3">
                    <FormRow label="用户名" required hint="3-30 字符">
                        <Input
                            value={form.username}
                            onChange={(e) =>
                                setForm({ ...form, username: e.target.value })
                            }
                            placeholder="例如：zhangsan"
                        />
                    </FormRow>
                    <FormRow label="昵称" required>
                        <Input
                            value={form.nickname}
                            onChange={(e) =>
                                setForm({ ...form, nickname: e.target.value })
                            }
                            placeholder="显示名称"
                        />
                    </FormRow>
                </div>
                <FormRow label="初始密码" required hint="至少 6 位">
                    <Input
                        type="password"
                        value={form.password}
                        onChange={(e) =>
                            setForm({ ...form, password: e.target.value })
                        }
                        placeholder="••••••"
                    />
                </FormRow>
                <div className="grid grid-cols-2 gap-3">
                    <FormRow label="手机">
                        <Input
                            value={form.phone}
                            onChange={(e) =>
                                setForm({ ...form, phone: e.target.value })
                            }
                            placeholder="可选"
                        />
                    </FormRow>
                    <FormRow label="邮箱">
                        <Input
                            value={form.email}
                            onChange={(e) =>
                                setForm({ ...form, email: e.target.value })
                            }
                            placeholder="可选"
                        />
                    </FormRow>
                </div>
                <FormRow label="部门">
                    <Input
                        value={form.department_name}
                        onChange={(e) =>
                            setForm({
                                ...form,
                                department_name: e.target.value,
                            })
                        }
                        placeholder="可选"
                    />
                </FormRow>
                <FormRow label="初始角色" hint="不选则不分配角色，可在编辑里再设置">
                    <Select
                        value={form.role}
                        onValueChange={(v) => setForm({ ...form, role: v ?? '' })}
                    >
                        <SelectTrigger>
                            <SelectValue placeholder="选择一个角色（可选）" />
                        </SelectTrigger>
                        <SelectContent>
                            {roles.map((r) => (
                                <SelectItem
                                    key={r.role.id}
                                    value={r.role.name}
                                >
                                    {r.role.name}
                                </SelectItem>
                            ))}
                        </SelectContent>
                    </Select>
                </FormRow>
            </div>
        </FormDialog>
    );
}

// ── Edit dialog ───────────────────────────────────────────────────

function EditUserDialog({
    user,
    roles,
    canAssignRole,
    onClose,
    onSaved,
}: {
    user: UserRow;
    roles: RoleRow[];
    canAssignRole: boolean;
    onClose: () => void;
    onSaved: () => void;
}) {
    const [form, setForm] = useState({
        nickname: user.nickname,
        phone: user.phone ?? '',
        email: user.email ?? '',
        department_name: user.department_name ?? '',
        status: String(user.status),
        is_expert: String(user.is_expert ?? 0),
    });
    const [role, setRole] = useState(user.roles?.[0] ?? '');
    const [busy, setBusy] = useState(false);

    const submit = async () => {
        setBusy(true);
        try {
            await api.put(API.userById(user.id), {
                nickname: form.nickname,
                phone: form.phone || null,
                email: form.email || null,
                department_name: form.department_name || null,
                status: Number(form.status),
                is_expert: Number(form.is_expert),
            });
            if (canAssignRole && role && role !== user.roles?.[0]) {
                await api.put(`${API.userById(user.id)}/role`, { role });
            }
            toast.success('已保存');
            onSaved();
        } catch (err) {
            const msg =
                (err as { response?: { data?: { message?: string } } }).response
                    ?.data?.message ?? '保存失败';
            toast.error(msg);
        } finally {
            setBusy(false);
        }
    };

    return (
        <FormDialog
            open
            onOpenChange={(v) => !v && onClose()}
            title={`编辑用户「${user.username}」`}
            description={`#${user.id} · 修改资料 / 状态 / 角色`}
            size="lg"
            onConfirm={submit}
            confirmLabel="保存"
            loading={busy}
        >
            <div className="space-y-4">
                <FormRow label="昵称" required>
                    <Input
                        value={form.nickname}
                        onChange={(e) =>
                            setForm({ ...form, nickname: e.target.value })
                        }
                    />
                </FormRow>
                <div className="grid grid-cols-2 gap-3">
                    <FormRow label="手机">
                        <Input
                            value={form.phone}
                            onChange={(e) =>
                                setForm({ ...form, phone: e.target.value })
                            }
                        />
                    </FormRow>
                    <FormRow label="邮箱">
                        <Input
                            value={form.email}
                            onChange={(e) =>
                                setForm({ ...form, email: e.target.value })
                            }
                        />
                    </FormRow>
                </div>
                <FormRow label="部门">
                    <Input
                        value={form.department_name}
                        onChange={(e) =>
                            setForm({
                                ...form,
                                department_name: e.target.value,
                            })
                        }
                    />
                </FormRow>
                <div className="grid grid-cols-2 gap-3">
                    <FormRow label="状态">
                        <Select
                            value={form.status}
                            onValueChange={(v) =>
                                setForm({ ...form, status: v ?? '' })
                            }
                        >
                            <SelectTrigger>
                                <SelectValue />
                            </SelectTrigger>
                            <SelectContent>
                                <SelectItem value="1">启用</SelectItem>
                                <SelectItem value="2">禁用</SelectItem>
                            </SelectContent>
                        </Select>
                    </FormRow>
                    <FormRow label="专家身份">
                        <Select
                            value={form.is_expert}
                            onValueChange={(v) =>
                                setForm({ ...form, is_expert: v ?? '' })
                            }
                        >
                            <SelectTrigger>
                                <SelectValue />
                            </SelectTrigger>
                            <SelectContent>
                                <SelectItem value="0">否</SelectItem>
                                <SelectItem value="1">是</SelectItem>
                            </SelectContent>
                        </Select>
                    </FormRow>
                </div>
                {canAssignRole && (
                    <FormRow label="角色" hint="只能选一个，会替换现有角色">
                        <Select value={role} onValueChange={(v) => setRole(v ?? '')}>
                            <SelectTrigger>
                                <SelectValue placeholder="选择角色" />
                            </SelectTrigger>
                            <SelectContent>
                                {roles.map((r) => (
                                    <SelectItem
                                        key={r.role.id}
                                        value={r.role.name}
                                    >
                                        {r.role.name}
                                    </SelectItem>
                                ))}
                            </SelectContent>
                        </Select>
                    </FormRow>
                )}
            </div>
        </FormDialog>
    );
}

// ── Reset password dialog ─────────────────────────────────────────

function ResetPasswordDialog({
    user,
    onClose,
    onSaved,
}: {
    user: UserRow;
    onClose: () => void;
    onSaved: () => void;
}) {
    const [pwd, setPwd] = useState('');
    const [busy, setBusy] = useState(false);

    const submit = async () => {
        if (pwd.length < 6) {
            toast.error('密码至少 6 位');
            return;
        }
        setBusy(true);
        try {
            await api.put(`${API.userById(user.id)}/password`, {
                password: pwd,
            });
            toast.success(`已重置「${user.username}」的密码`);
            onSaved();
        } catch (err) {
            toast.error('重置失败');
            // eslint-disable-next-line no-console
            console.error(err);
        } finally {
            setBusy(false);
        }
    };

    return (
        <FormDialog
            open
            onOpenChange={(v) => !v && onClose()}
            title={`重置「${user.username}」的密码`}
            icon={KeyRound}
            description="管理员代为重置；新密码请告知本人"
            size="md"
            onConfirm={submit}
            confirmLabel="确认重置"
            confirmTone="amber"
            loading={busy}
        >
            <div className="space-y-3">
                <FormRow label="新密码" required hint="至少 6 位">
                    <Input
                        type="password"
                        value={pwd}
                        onChange={(e) => setPwd(e.target.value)}
                        placeholder="••••••"
                        autoFocus
                    />
                </FormRow>
                <p className="rounded-md border border-amber-500/30 bg-amber-600/10 p-3 text-[11px] leading-relaxed text-amber-300/80">
                    此操作仅重置密码，<strong>不会</strong> 强制对方下次登录改密。
                    如需强制改密，请联系运维设置 must_change_credentials = true。
                </p>
            </div>
        </FormDialog>
    );
}

