import { useCallback, useEffect, useMemo, useState } from 'react';
import {
    Edit3,
    KeyRound,
    Lock,
    Plus,
    RefreshCw,
    ShieldCheck,
    Trash2,
} from 'lucide-react';
import { toast } from 'sonner';
import api from '@/lib/api';
import { API } from '@/lib/api-routes';
import { ActionBtn } from '@/components/shared/ActionBtn';
import { FormDialog, FormRow } from '@/components/shared/FormDialog';
import { Input } from '@/components/ui/input';
import { useAuth } from '@/hooks/use-auth';

interface RoleRow {
    role: {
        id: number;
        name: string;
        guard_name: string;
        created_at: string;
    };
    permissions: string[];
}

interface PermissionItem {
    id: number;
    name: string;
}

interface PermissionsResponse {
    permissions: PermissionItem[];
    grouped: Record<string, string[]>;
}

// 跟 aifix_hs 项目页面（resources/js/pages/system/roles/Index.tsx）的按钮样式对齐：
//   primaryBtn — 主操作（新增、保存）边框 + 半透明蓝底
//   ghostBtn   — 次操作（取消、刷新）通透蓝灰边框
const primaryBtn =
    'flex items-center gap-2 rounded-lg border border-blue-500/30 bg-blue-600/15 px-4 py-2 text-sm font-medium text-blue-300 transition-all hover:border-blue-400/50 hover:bg-blue-600/25 hover:text-blue-200 active:scale-95';
const ghostBtn =
    'flex items-center gap-2 rounded-lg border border-blue-900/40 px-4 py-2 text-sm text-slate-400 transition-colors hover:border-blue-500/40 hover:text-slate-200';

const builtinRoles = new Set([
    'platform_admin',
    'tenant_root_admin',
    'tenant_admin',
    'viewer',
]);

export default function SystemRoles() {
    const { hasPermission } = useAuth();
    const [rows, setRows] = useState<RoleRow[]>([]);
    const [grouped, setGrouped] = useState<Record<string, string[]>>({});
    const [loading, setLoading] = useState(true);
    const [editing, setEditing] = useState<RoleRow | null>(null);
    const [creating, setCreating] = useState(false);

    const canCreate = hasPermission('roles.create');
    const canUpdate = hasPermission('roles.update');
    const canDelete = hasPermission('roles.delete');

    const refresh = useCallback(async () => {
        setLoading(true);
        try {
            const [rolesRes, permsRes] = await Promise.all([
                api.get(API.roles()),
                api.get(API.permissions()),
            ]);
            setRows((rolesRes.data?.data as RoleRow[]) ?? []);
            setGrouped(
                (permsRes.data?.data as PermissionsResponse)?.grouped ?? {},
            );
        } catch (err) {
            toast.error('加载角色失败');
            // eslint-disable-next-line no-console
            console.error(err);
        } finally {
            setLoading(false);
        }
    }, []);

    useEffect(() => {
        refresh();
    }, [refresh]);

    const totalPerms = useMemo(
        () =>
            Object.values(grouped).reduce(
                (s: number, arr: string[]) => s + arr.length,
                0,
            ),
        [grouped],
    );

    return (
        <div className="space-y-6">
            <div className="flex items-center justify-between">
                <div>
                    <h1 className="text-xl font-bold text-slate-100">角色管理</h1>
                    <p className="mt-0.5 text-xs text-slate-600">
                        共{' '}
                        <span className="font-medium text-slate-400">
                            {rows.length}
                        </span>{' '}
                        个角色
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
                            <Plus className="h-4 w-4" />
                            新增角色
                        </button>
                    )}
                </div>
            </div>

            <div className="grid gap-4 md:grid-cols-3">
                <Stat
                    label="角色总数"
                    value={rows.length}
                    icon={ShieldCheck}
                    accent="sky"
                />
                <Stat
                    label="内置角色"
                    value={
                        rows.filter((r) => builtinRoles.has(r.role.name)).length
                    }
                    icon={Lock}
                    accent="amber"
                />
                <Stat
                    label="可用权限总数"
                    value={totalPerms}
                    icon={KeyRound}
                    accent="violet"
                />
            </div>

            <div className="overflow-hidden rounded-xl border border-blue-900/40 bg-[#0d1530]">
                {loading ? (
                    <div className="px-5 py-12 text-center text-sm text-slate-500">
                        <RefreshCw className="mx-auto mb-3 h-5 w-5 animate-spin text-slate-600" />
                        加载中…
                    </div>
                ) : (
                    <table className="w-full text-sm">
                        <thead>
                            <tr className="border-b border-blue-900/30 bg-[#0a1428] text-[10px] tracking-widest text-slate-600 uppercase">
                                <th className="px-5 py-3 text-left">角色</th>
                                <th className="px-5 py-3 text-left">权限数</th>
                                <th className="px-5 py-3 text-left">权限预览</th>
                                <th className="px-5 py-3 text-right">操作</th>
                            </tr>
                        </thead>
                        <tbody>
                            {rows.map((r) => {
                                const isBuiltin = builtinRoles.has(r.role.name);
                                return (
                                    <tr
                                        key={r.role.id}
                                        className="border-b border-blue-900/20 transition-colors last:border-b-0 hover:bg-[#0a1428]"
                                    >
                                        <td className="px-5 py-3">
                                            <div className="flex items-center gap-3">
                                                <div
                                                    className={`flex h-9 w-9 shrink-0 items-center justify-center rounded-lg border ${
                                                        isBuiltin
                                                            ? 'border-amber-500/30 bg-amber-600/15 text-amber-400'
                                                            : 'border-blue-600/20 bg-blue-600/15 text-blue-400'
                                                    }`}
                                                >
                                                    {isBuiltin ? (
                                                        <Lock className="h-4 w-4" />
                                                    ) : (
                                                        <ShieldCheck className="h-4 w-4" />
                                                    )}
                                                </div>
                                                <div>
                                                    <div className="flex items-center gap-2">
                                                        <span className="font-mono font-medium text-slate-100">
                                                            {r.role.name}
                                                        </span>
                                                        {isBuiltin && (
                                                            <span className="rounded border border-amber-500/30 bg-amber-600/15 px-1 py-0.5 text-[9px] text-amber-400">
                                                                内置
                                                            </span>
                                                        )}
                                                    </div>
                                                    <div className="text-[11px] text-slate-600">
                                                        guard:{' '}
                                                        {r.role.guard_name}
                                                    </div>
                                                </div>
                                            </div>
                                        </td>
                                        <td className="px-5 py-3">
                                            <span className="rounded-md border border-blue-500/30 bg-blue-600/15 px-2 py-0.5 text-xs text-blue-300">
                                                {r.permissions.length}
                                            </span>
                                        </td>
                                        <td className="px-5 py-3">
                                            <div className="flex flex-wrap gap-1">
                                                {r.permissions.slice(0, 4).map((p) => (
                                                    <span
                                                        key={p}
                                                        className="rounded border border-slate-700/50 bg-slate-800/30 px-1.5 py-0.5 font-mono text-[10px] text-slate-400"
                                                    >
                                                        {p}
                                                    </span>
                                                ))}
                                                {r.permissions.length > 4 && (
                                                    <span className="text-[10px] text-slate-600">
                                                        +{r.permissions.length - 4}
                                                    </span>
                                                )}
                                            </div>
                                        </td>
                                        <td className="px-5 py-3 text-right">
                                            <div className="flex items-center justify-end gap-0.5">
                                                {canUpdate && (
                                                    <ActionBtn
                                                        onClick={() =>
                                                            setEditing(r)
                                                        }
                                                        title="编辑"
                                                        variant="edit"
                                                    >
                                                        <Edit3 className="h-3.5 w-3.5" />
                                                    </ActionBtn>
                                                )}
                                                {canDelete && !isBuiltin && (
                                                    <DeleteButton
                                                        role={r}
                                                        onDeleted={refresh}
                                                    />
                                                )}
                                            </div>
                                        </td>
                                    </tr>
                                );
                            })}
                        </tbody>
                    </table>
                )}
            </div>

            {creating && (
                <RoleEditorDialog
                    grouped={grouped}
                    onClose={() => setCreating(false)}
                    onSaved={() => {
                        setCreating(false);
                        refresh();
                    }}
                />
            )}
            {editing && (
                <RoleEditorDialog
                    role={editing}
                    grouped={grouped}
                    onClose={() => setEditing(null)}
                    onSaved={() => {
                        setEditing(null);
                        refresh();
                    }}
                />
            )}
        </div>
    );
}

// ── Stat tile ─────────────────────────────────────────────────────

function Stat({
    label,
    value,
    icon: Icon,
    accent,
}: {
    label: string;
    value: number;
    icon: React.ComponentType<{ className?: string }>;
    accent: 'sky' | 'amber' | 'violet';
}) {
    const tone = {
        sky: 'border-sky-600/20 bg-sky-600/15 text-sky-400',
        amber: 'border-amber-600/20 bg-amber-600/15 text-amber-400',
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
    role,
    onDeleted,
}: {
    role: RoleRow;
    onDeleted: () => void;
}) {
    const [open, setOpen] = useState(false);
    const [busy, setBusy] = useState(false);
    const handle = async () => {
        setBusy(true);
        try {
            await api.delete(API.roleById(role.role.id));
            toast.success('已删除');
            onDeleted();
        } catch (err) {
            const msg =
                (err as { response?: { data?: { message?: string } } }).response
                    ?.data?.message ?? '删除失败';
            toast.error(msg);
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
                title={`删除角色「${role.role.name}」`}
                icon={Trash2}
                size="md"
                onConfirm={handle}
                confirmLabel="确认删除"
                confirmTone="rose"
                loading={busy}
            >
                <p className="text-sm leading-relaxed text-slate-400">
                    如该角色下还有用户，删除将被拒绝。请先把对应用户改成其它角色再来。
                </p>
            </FormDialog>
        </>
    );
}

// ── Role editor (create + edit share) ─────────────────────────────

function RoleEditorDialog({
    role,
    grouped,
    onClose,
    onSaved,
}: {
    role?: RoleRow;
    grouped: Record<string, string[]>;
    onClose: () => void;
    onSaved: () => void;
}) {
    const isEdit = !!role;
    const isBuiltin = isEdit && builtinRoles.has(role!.role.name);

    const [name, setName] = useState(role?.role.name ?? '');
    const [selected, setSelected] = useState<Set<string>>(
        new Set(role?.permissions ?? []),
    );
    const [busy, setBusy] = useState(false);

    const togglePerm = (p: string) =>
        setSelected((prev) => {
            const next = new Set(prev);
            if (next.has(p)) next.delete(p);
            else next.add(p);
            return next;
        });

    const toggleGroup = (perms: string[]) => {
        const allOn = perms.every((p) => selected.has(p));
        setSelected((prev) => {
            const next = new Set(prev);
            if (allOn) {
                perms.forEach((p) => next.delete(p));
            } else {
                perms.forEach((p) => next.add(p));
            }
            return next;
        });
    };

    const submit = async () => {
        if (!name.trim()) {
            toast.error('角色名不能为空');
            return;
        }
        setBusy(true);
        try {
            const payload: { name?: string; permissions: string[] } = {
                permissions: Array.from(selected),
            };
            if (!isBuiltin) {
                payload.name = name.trim();
            }
            if (isEdit) {
                await api.put(API.roleById(role!.role.id), payload);
            } else {
                await api.post(API.roles(), {
                    name: name.trim(),
                    permissions: Array.from(selected),
                });
            }
            toast.success(isEdit ? '已保存' : '角色已创建');
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

    const groups = Object.keys(grouped).sort();

    return (
        <FormDialog
            open
            onOpenChange={(v) => !v && onClose()}
            title={isEdit ? `编辑角色「${role!.role.name}」` : '新建角色'}
            icon={isEdit ? Edit3 : Plus}
            description={
                isEdit
                    ? '调整角色名（如非内置）与该角色拥有的权限组合'
                    : '为新角色起名并选择它该拥有哪些权限'
            }
            size="xl"
            onConfirm={submit}
            confirmLabel={isEdit ? '保存' : '创建'}
            loading={busy}
        >
            <div className="space-y-4">
                <FormRow
                    label="角色名"
                    required
                    hint={
                        isBuiltin
                            ? '内置角色名称不可修改，但权限可以调整'
                            : '建议使用 snake_case，例如 tenant_inspector'
                    }
                >
                    <Input
                        value={name}
                        onChange={(e) => setName(e.target.value)}
                        disabled={isBuiltin}
                        placeholder="例如：tenant_inspector"
                    />
                </FormRow>
                <div>
                    <div className="mb-2 flex items-center justify-between">
                        <span className="text-xs font-medium text-slate-400">
                            权限分配
                            <span className="ml-2 text-[11px] font-normal text-slate-600">
                                已选 {selected.size} 项
                            </span>
                        </span>
                        <button
                            onClick={() => setSelected(new Set())}
                            className="text-[11px] text-slate-500 hover:text-slate-300"
                        >
                            清空
                        </button>
                    </div>
                    <div className="max-h-[42vh] space-y-3 overflow-y-auto rounded-lg border border-blue-900/30 bg-[#0a1428] p-3">
                            {groups.length === 0 ? (
                                <p className="text-xs text-slate-500">
                                    无可用权限
                                </p>
                            ) : (
                                groups.map((g) => {
                                    const perms = grouped[g];
                                    const allOn = perms.every((p) =>
                                        selected.has(p),
                                    );
                                    const someOn =
                                        !allOn &&
                                        perms.some((p) => selected.has(p));
                                    return (
                                        <div key={g}>
                                            <div className="mb-1.5 flex items-center justify-between">
                                                <button
                                                    onClick={() =>
                                                        toggleGroup(perms)
                                                    }
                                                    className="flex items-center gap-2 text-xs"
                                                >
                                                    <span
                                                        className={`flex h-3.5 w-3.5 shrink-0 items-center justify-center rounded border ${
                                                            allOn
                                                                ? 'border-blue-500/60 bg-blue-600/40'
                                                                : someOn
                                                                  ? 'border-blue-500/40 bg-blue-600/15'
                                                                  : 'border-slate-700'
                                                        }`}
                                                    >
                                                        {allOn && (
                                                            <span className="h-1.5 w-1.5 rounded-sm bg-blue-300" />
                                                        )}
                                                        {someOn && (
                                                            <span className="h-1 w-1 rounded-sm bg-blue-400" />
                                                        )}
                                                    </span>
                                                    <span className="font-mono font-semibold text-slate-300">
                                                        {g}
                                                    </span>
                                                    <span className="text-[10px] text-slate-600">
                                                        ({perms.filter((p) => selected.has(p)).length}/{perms.length})
                                                    </span>
                                                </button>
                                            </div>
                                            <div className="ml-5 flex flex-wrap gap-1.5">
                                                {perms.map((p) => {
                                                    const on =
                                                        selected.has(p);
                                                    return (
                                                        <button
                                                            key={p}
                                                            onClick={() =>
                                                                togglePerm(p)
                                                            }
                                                            className={`rounded-md border px-2 py-1 font-mono text-[10px] transition-colors ${
                                                                on
                                                                    ? 'border-blue-500/40 bg-blue-600/20 text-blue-300'
                                                                    : 'border-slate-700/50 bg-slate-800/30 text-slate-400 hover:border-slate-600 hover:text-slate-200'
                                                            }`}
                                                        >
                                                            {p.replace(g + '.', '')}
                                                        </button>
                                                    );
                                                })}
                                            </div>
                                        </div>
                                    );
                                })
                        )}
                    </div>
                </div>
            </div>
        </FormDialog>
    );
}
