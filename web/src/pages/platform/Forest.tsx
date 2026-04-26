import { useCallback, useEffect, useMemo, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import {
    Building,
    Building2,
    ChevronDown,
    ChevronRight,
    GitBranch,
    Layers,
    LogIn,
    Plus,
    RefreshCw,
    Trees,
    Wrench,
} from 'lucide-react';
import { toast } from 'sonner';
import api from '@/lib/api';
import { API } from '@/lib/api-routes';
import { useScope } from '@/contexts/ScopeContext';
import { Button } from '@/components/ui/button';
import { FormDialog, FormRow } from '@/components/shared/FormDialog';
import {
    CredentialsModal,
    type AdminCredentials,
} from '@/components/shared/CredentialsModal';
import { Input } from '@/components/ui/input';

interface TreeNode {
    id: number;
    title: string;
    parent_id: number | null;
    root_id: number;
    status: number;
    is_bloc?: number;
    is_produce?: number;
    children?: TreeNode[];
}

// ── helpers ────────────────────────────────────────────────────────

function countNodes(node: TreeNode): number {
    let n = 1;
    for (const c of node.children ?? []) n += countNodes(c);
    return n;
}
function maxDepth(node: TreeNode, d = 0): number {
    if (!node.children || node.children.length === 0) return d;
    return Math.max(...node.children.map((c) => maxDepth(c, d + 1)));
}

// ── small UI bits ─────────────────────────────────────────────────

function StatusDot({ status }: { status: number }) {
    const cls =
        status === 1
            ? 'bg-emerald-400 shadow-[0_0_6px_rgba(52,211,153,0.7)]'
            : status === 2
              ? 'bg-amber-400'
              : 'bg-slate-500';
    return <span className={`h-1.5 w-1.5 rounded-full ${cls}`} />;
}

function StatusBadge({ status }: { status: number }) {
    const map: Record<number, { cls: string; label: string }> = {
        0: {
            cls: 'border-slate-500/30 bg-slate-600/15 text-slate-400',
            label: '待审',
        },
        1: {
            cls: 'border-emerald-500/30 bg-emerald-600/15 text-emerald-400',
            label: '启用',
        },
        2: {
            cls: 'border-amber-500/30 bg-amber-600/15 text-amber-400',
            label: '禁用',
        },
    };
    const m = map[status] ?? map[0];
    return (
        <span
            className={`rounded-full border px-1.5 py-0.5 text-[9px] font-medium tracking-wide uppercase ${m.cls}`}
        >
            {m.label}
        </span>
    );
}

function StatTile({
    label,
    value,
    icon: Icon,
    accent,
}: {
    label: string;
    value: number | string;
    icon: React.ComponentType<{ className?: string }>;
    accent: 'sky' | 'emerald' | 'violet' | 'amber';
}) {
    const tone = {
        sky: 'border-sky-600/20 bg-sky-600/15 text-sky-400',
        emerald: 'border-emerald-600/20 bg-emerald-600/15 text-emerald-400',
        violet: 'border-violet-600/20 bg-violet-600/15 text-violet-400',
        amber: 'border-amber-600/20 bg-amber-600/15 text-amber-400',
    }[accent];
    return (
        <div className="rounded-xl border border-blue-900/40 bg-[#0d1530] p-5">
            <div className="flex items-start justify-between">
                <div>
                    <p className="text-xs font-medium tracking-wide text-slate-500 uppercase">
                        {label}
                    </p>
                    <p className="mt-1 text-2xl font-bold text-slate-100">{value}</p>
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

// ── small UI bits — toggle chip used in the create-subsidiary form ──

function ToggleChip({
    active,
    onClick,
    label,
    accent,
}: {
    active: boolean;
    onClick: () => void;
    label: string;
    accent: 'blue' | 'violet';
}) {
    const activeClass = {
        blue: 'border-blue-500/40 bg-blue-600/15 text-blue-300',
        violet: 'border-violet-500/40 bg-violet-600/15 text-violet-300',
    }[accent];
    return (
        <button
            type="button"
            onClick={onClick}
            className={`flex flex-1 items-center justify-center gap-2 rounded-lg border px-3 py-2 text-sm font-medium transition-all ${
                active
                    ? activeClass
                    : 'border-blue-900/40 bg-[#0a1428] text-slate-500 hover:border-blue-500/30 hover:text-slate-300'
            }`}
        >
            <span
                className={`h-1.5 w-1.5 rounded-full ${active ? 'bg-current' : 'bg-slate-700'}`}
            />
            {label}
        </button>
    );
}

// ── recursive tree row ────────────────────────────────────────────

function TreeRow({
    node,
    depth,
    expanded,
    onToggle,
    onAddChild,
    onEnter,
    isLast,
    parentRail,
}: {
    node: TreeNode;
    depth: number;
    expanded: Set<number>;
    onToggle: (id: number) => void;
    onAddChild: (parent: TreeNode) => void;
    onEnter: (node: TreeNode) => void;
    isLast: boolean;
    parentRail: boolean[]; // for each ancestor depth: should we draw the vertical rail?
}) {
    const hasKids = (node.children ?? []).length > 0;
    const isExpanded = expanded.has(node.id);
    const isRoot = node.parent_id === null;

    return (
        <div>
            <div className="group relative flex items-center gap-2 py-2 transition-colors hover:bg-[#0a1428]">
                {/* vertical rails for ancestors */}
                {parentRail.map((draw, i) => (
                    <span
                        key={i}
                        className="absolute top-0 bottom-0 w-px"
                        style={{
                            left: 18 + i * 22,
                            background: draw ? 'rgba(59,130,246,0.18)' : 'transparent',
                        }}
                    />
                ))}
                {/* this row's elbow */}
                {depth > 0 && (
                    <>
                        <span
                            className="absolute w-px bg-blue-500/20"
                            style={{
                                left: 18 + (depth - 1) * 22,
                                top: 0,
                                height: '50%',
                            }}
                        />
                        <span
                            className="absolute h-px bg-blue-500/20"
                            style={{
                                left: 18 + (depth - 1) * 22,
                                top: '50%',
                                width: 14,
                            }}
                        />
                        {!isLast && (
                            <span
                                className="absolute w-px bg-blue-500/20"
                                style={{
                                    left: 18 + (depth - 1) * 22,
                                    top: '50%',
                                    bottom: 0,
                                }}
                            />
                        )}
                    </>
                )}

                {/* indent + chevron */}
                <div
                    className="flex shrink-0 items-center"
                    style={{ paddingLeft: depth * 22 }}
                >
                    {hasKids ? (
                        <button
                            onClick={() => onToggle(node.id)}
                            className="flex h-5 w-5 items-center justify-center rounded text-slate-500 hover:bg-blue-600/15 hover:text-slate-200"
                        >
                            {isExpanded ? (
                                <ChevronDown className="h-3.5 w-3.5" />
                            ) : (
                                <ChevronRight className="h-3.5 w-3.5" />
                            )}
                        </button>
                    ) : (
                        <span className="h-5 w-5" />
                    )}
                </div>

                {/* icon */}
                <div
                    className={`flex h-7 w-7 shrink-0 items-center justify-center rounded-lg border ${
                        isRoot
                            ? 'border-blue-600/30 bg-blue-600/15'
                            : 'border-slate-700/30 bg-slate-700/15'
                    }`}
                >
                    {isRoot ? (
                        <Building2 className="h-3.5 w-3.5 text-blue-400" />
                    ) : (
                        <Building className="h-3.5 w-3.5 text-slate-400" />
                    )}
                </div>

                {/* title + meta */}
                <div className="min-w-0 flex-1">
                    <div className="flex items-center gap-2">
                        <span className="truncate text-sm font-medium text-slate-100">
                            {node.title}
                        </span>
                        <StatusDot status={node.status} />
                        <StatusBadge status={node.status} />
                        {node.is_bloc === 1 && (
                            <span className="rounded-md border border-blue-500/30 bg-blue-600/15 px-1.5 py-0.5 text-[9px] font-semibold tracking-wider text-blue-400 uppercase">
                                集团
                            </span>
                        )}
                        {node.is_produce === 1 && (
                            <span className="rounded-md border border-violet-500/30 bg-violet-600/15 px-1.5 py-0.5 text-[9px] font-semibold tracking-wider text-violet-400 uppercase">
                                生产
                            </span>
                        )}
                    </div>
                    <div className="mt-0.5 flex items-center gap-2 text-[11px] text-slate-600">
                        <span>#{node.id}</span>
                        <span>·</span>
                        <span>
                            root_id:{' '}
                            <code className="text-slate-500">{node.root_id}</code>
                        </span>
                        {node.parent_id !== null && (
                            <>
                                <span>·</span>
                                <span>
                                    parent:{' '}
                                    <code className="text-slate-500">
                                        {node.parent_id}
                                    </code>
                                </span>
                            </>
                        )}
                        {hasKids && (
                            <>
                                <span>·</span>
                                <span className="text-blue-500">
                                    {node.children!.length} 直系子公司
                                </span>
                            </>
                        )}
                    </div>
                </div>

                {/* actions (visible on hover) */}
                <div className="flex shrink-0 gap-1 pr-3 opacity-0 transition-opacity group-hover:opacity-100">
                    <button
                        onClick={() => onEnter(node)}
                        title="以此企业身份查看（化身）"
                        className="flex items-center gap-1 rounded-md border border-blue-500/20 bg-blue-600/10 px-2 py-1 text-[10px] text-blue-400 hover:border-blue-500/40 hover:bg-blue-600/20"
                    >
                        <LogIn className="h-3 w-3" />
                        化身
                    </button>
                    <button
                        onClick={() => onAddChild(node)}
                        title="在此节点下新建子公司"
                        className="flex items-center gap-1 rounded-md border border-emerald-500/20 bg-emerald-600/10 px-2 py-1 text-[10px] text-emerald-400 hover:border-emerald-500/40 hover:bg-emerald-600/20"
                    >
                        <Plus className="h-3 w-3" />
                        子公司
                    </button>
                </div>
            </div>

            {/* children */}
            {isExpanded &&
                hasKids &&
                node.children!.map((c, idx) => (
                    <TreeRow
                        key={c.id}
                        node={c}
                        depth={depth + 1}
                        expanded={expanded}
                        onToggle={onToggle}
                        onAddChild={onAddChild}
                        onEnter={onEnter}
                        isLast={idx === node.children!.length - 1}
                        parentRail={[
                            ...parentRail,
                            !isLast || depth === 0 || idx < node.children!.length - 1,
                        ]}
                    />
                ))}
        </div>
    );
}

// ── Page ──────────────────────────────────────────────────────────

export default function PlatformForest() {
    const navigate = useNavigate();
    const { enterImpersonation } = useScope();
    const [forest, setForest] = useState<TreeNode[]>([]);
    const [loading, setLoading] = useState(true);
    const [parent, setParent] = useState<TreeNode | null>(null);
    const [newTitle, setNewTitle] = useState('');
    const [newIsBloc, setNewIsBloc] = useState(false);
    const [newIsProduce, setNewIsProduce] = useState(true);
    const [submitting, setSubmitting] = useState(false);
    const [expanded, setExpanded] = useState<Set<number>>(new Set());
    const [credentials, setCredentials] = useState<AdminCredentials | null>(
        null,
    );

    const refresh = useCallback(async () => {
        setLoading(true);
        try {
            const res = await api.get(API.companiesForest());
            const data = (res.data?.data as TreeNode[]) ?? [];
            setForest(data);
            // Default: expand every node so the full topology (third level
            // and below) is visible without manual chevron-clicking. The
            // "全部折叠" button gives users a quick reset.
            const allIds = new Set<number>();
            const walk = (n: TreeNode) => {
                allIds.add(n.id);
                (n.children ?? []).forEach(walk);
            };
            data.forEach(walk);
            setExpanded(allIds);
        } catch (err) {
            toast.error('加载森林失败');
            // eslint-disable-next-line no-console
            console.error(err);
        } finally {
            setLoading(false);
        }
    }, []);

    useEffect(() => {
        refresh();
    }, [refresh]);

    const handleToggle = (id: number) =>
        setExpanded((prev) => {
            const next = new Set(prev);
            if (next.has(id)) next.delete(id);
            else next.add(id);
            return next;
        });

    const expandAll = () => {
        const all = new Set<number>();
        const walk = (n: TreeNode) => {
            all.add(n.id);
            (n.children ?? []).forEach(walk);
        };
        forest.forEach(walk);
        setExpanded(all);
    };
    const collapseAll = () => setExpanded(new Set());

    const handleCreate = async () => {
        if (!parent || !newTitle.trim()) return;
        if (!newIsBloc && !newIsProduce) {
            toast.error('请至少勾选一项类型（集团 / 生产）');
            return;
        }
        const createdTitle = newTitle.trim();
        setSubmitting(true);
        try {
            const res = await api.post(API.companies(), {
                title: createdTitle,
                parent_id: parent.id,
                operational_type_id: 1,
                is_bloc: newIsBloc ? 1 : 0,
                is_produce: newIsProduce ? 1 : 0,
            });
            const created = res.data?.data ?? {};
            const creds = created.admin_credentials;
            toast.success(`已在「${parent.title}」下创建子公司`);
            setParent(null);
            setNewTitle('');
            setNewIsBloc(false);
            setNewIsProduce(true);
            await refresh();
            // Sub-company creation auto-provisions a `c{id}_admin` user
            // (matches Laravel CompanyController::store). Surface the
            // initial credentials in a one-time modal — the plaintext
            // password is otherwise unrecoverable after this response.
            if (creds?.username && creds?.password) {
                setCredentials({
                    companyTitle: createdTitle,
                    username: creds.username,
                    password: creds.password,
                });
            }
        } catch (err) {
            const msg =
                (err as { response?: { data?: { message?: string } } }).response
                    ?.data?.message ?? '创建失败';
            toast.error(msg);
        } finally {
            setSubmitting(false);
        }
    };

    // Impersonate the clicked node directly, at any depth. The platform admin
    // becomes a user of that exact company — X-Scope-Company-Id will be the
    // node's id, and subordinates resolve to that node's direct children.
    const handleEnter = async (node: TreeNode) => {
        if (node.status !== 1) {
            toast.error('该企业已禁用，无法化身');
            return;
        }
        try {
            await enterImpersonation({
                id: node.id,
                title: node.title,
                is_bloc: node.is_bloc ?? 0,
                is_produce: node.is_produce ?? 0,
                pid: node.parent_id ?? 0,
                status: node.status,
            });
            toast.success(`已化身为「${node.title}」`);
            navigate('/dashboard');
        } catch (err) {
            const msg =
                (err as { response?: { data?: { message?: string } } }).response
                    ?.data?.message ?? '化身失败';
            toast.error(msg);
        }
    };

    // Stats
    const stats = useMemo(() => {
        const totalCompanies = forest.reduce((s, t) => s + countNodes(t), 0);
        const totalRoots = forest.length;
        const deepest = forest.reduce((m, t) => Math.max(m, maxDepth(t)), 0);
        const subsidiaries = totalCompanies - totalRoots;
        return { totalCompanies, totalRoots, subsidiaries, deepest };
    }, [forest]);

    return (
        <div className="space-y-6">
            {/* Header */}
            <div className="flex items-start justify-between">
                <div>
                    <h1 className="text-2xl font-bold text-slate-100">企业森林</h1>
                    <p className="mt-1 text-sm text-slate-500">
                        所有一级企业及其子树的完整视图，用于直观查看跨租户拓扑
                    </p>
                </div>
                <div className="flex items-center gap-2">
                    <Button
                        variant="outline"
                        size="sm"
                        onClick={expandAll}
                        className="border-blue-900/40 bg-[#0d1530]"
                    >
                        全部展开
                    </Button>
                    <Button
                        variant="outline"
                        size="sm"
                        onClick={collapseAll}
                        className="border-blue-900/40 bg-[#0d1530]"
                    >
                        全部折叠
                    </Button>
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
            </div>

            {/* Stats */}
            <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
                <StatTile
                    label="树根（一级企业）"
                    value={stats.totalRoots}
                    icon={Trees}
                    accent="sky"
                />
                <StatTile
                    label="子公司总数"
                    value={stats.subsidiaries}
                    icon={Building}
                    accent="violet"
                />
                <StatTile
                    label="企业总数"
                    value={stats.totalCompanies}
                    icon={Building2}
                    accent="emerald"
                />
                <StatTile
                    label="最深层级"
                    value={stats.deepest === 0 ? '—' : `${stats.deepest + 1}`}
                    icon={Layers}
                    accent="amber"
                />
            </div>

            {/* Forest */}
            <div className="overflow-hidden rounded-xl border border-blue-900/40 bg-[#0d1530]">
                <div className="border-b border-blue-900/30 bg-[#0a1428] px-5 py-3">
                    <div className="flex items-center gap-2 text-sm">
                        <GitBranch className="h-4 w-4 text-blue-400" />
                        <span className="font-semibold text-slate-200">
                            森林拓扑
                        </span>
                        <span className="text-xs text-slate-600">
                            悬停节点查看操作 · 化身 / 新建子公司
                        </span>
                    </div>
                </div>
                {loading ? (
                    <div className="px-5 py-12 text-center text-sm text-slate-500">
                        <RefreshCw className="mx-auto mb-3 h-5 w-5 animate-spin text-slate-600" />
                        加载中…
                    </div>
                ) : forest.length === 0 ? (
                    <div className="px-5 py-12 text-center">
                        <Trees className="mx-auto mb-3 h-10 w-10 text-slate-700" />
                        <p className="text-sm text-slate-400">森林为空</p>
                        <p className="mt-1 text-xs text-slate-600">
                            前往「平台 → 一级企业」创建第一棵树
                        </p>
                    </div>
                ) : (
                    <div className="divide-y divide-blue-900/30">
                        {forest.map((tree) => (
                            <div key={tree.id} className="relative">
                                <TreeRow
                                    node={tree}
                                    depth={0}
                                    expanded={expanded}
                                    onToggle={handleToggle}
                                    onAddChild={(p) => {
                                        setParent(p);
                                        setNewTitle('');
                                    }}
                                    onEnter={handleEnter}
                                    isLast
                                    parentRail={[]}
                                />
                            </div>
                        ))}
                    </div>
                )}
            </div>

            {/* Hint */}
            <div className="rounded-lg border border-blue-900/30 bg-[#0a1428] p-4">
                <div className="flex gap-3 text-xs text-slate-500">
                    <Wrench className="h-4 w-4 shrink-0 text-slate-600" />
                    <p>
                        每个子公司的 <code>root_id</code> 由数据库触发器自动写入，
                        指向所在树的根 id。这是跨租户隔离的关键索引（
                        <code>WHERE root_id = X</code> 一步定位整棵树的全部数据）。
                    </p>
                </div>
            </div>

            {/* Create child dialog */}
            <FormDialog
                open={parent !== null}
                onOpenChange={(v) => !v && setParent(null)}
                title={
                    parent
                        ? `在「${parent.title}」下新建子公司`
                        : '新建子公司'
                }
                icon={Plus}
                size="md"
                onConfirm={handleCreate}
                confirmLabel="创建"
                loading={submitting}
            >
                {parent && (
                    <div className="space-y-4">
                        <div className="space-y-1.5 rounded-lg border border-blue-900/30 bg-[#0a1428] p-3 text-xs">
                            <div className="flex items-center justify-between">
                                <span className="text-slate-500">父企业</span>
                                <span className="text-slate-300">
                                    {parent.title}
                                </span>
                            </div>
                            <div className="flex items-center justify-between">
                                <span className="text-slate-500">父企业 ID</span>
                                <code className="text-slate-300">
                                    {parent.id}
                                </code>
                            </div>
                            <div className="flex items-center justify-between">
                                <span className="text-slate-500">
                                    父企业 root_id
                                </span>
                                <code className="text-slate-300">
                                    {parent.root_id}
                                </code>
                            </div>
                            <div className="mt-2 border-t border-blue-900/30 pt-2 text-[11px] text-emerald-400/80">
                                新子公司的 root_id 应自动 = {parent.root_id}
                            </div>
                        </div>
                        <FormRow label="子公司名称" required>
                            <Input
                                value={newTitle}
                                onChange={(e) => setNewTitle(e.target.value)}
                                placeholder="例如：北京分公司"
                                autoFocus
                            />
                        </FormRow>

                        <FormRow
                            label="企业类型"
                            required
                            hint="可同时勾选 — 集团企业用于风险/统计聚合，生产企业用于业务台账"
                        >
                            <div className="flex gap-2">
                                <ToggleChip
                                    active={newIsBloc}
                                    onClick={() => setNewIsBloc((v) => !v)}
                                    label="集团企业"
                                    accent="blue"
                                />
                                <ToggleChip
                                    active={newIsProduce}
                                    onClick={() => setNewIsProduce((v) => !v)}
                                    label="生产企业"
                                    accent="violet"
                                />
                            </div>
                        </FormRow>
                    </div>
                )}
            </FormDialog>

            {/* One-time admin credentials popup */}
            <CredentialsModal
                credentials={credentials}
                onClose={() => setCredentials(null)}
            />
        </div>
    );
}
