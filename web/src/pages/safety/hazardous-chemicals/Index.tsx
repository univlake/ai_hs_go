import { useCallback, useEffect, useMemo, useState } from 'react';
import {
    FlaskConical,
    Flame,
    Pencil,
    Plus,
    RefreshCw,
    Search,
    ShieldAlert,
    Skull,
    Trash2,
    X,
} from 'lucide-react'; // FlaskConical / Flame / ShieldAlert / Skull 仍用于编辑器分组与空态图标
import { toast } from 'sonner';
import api from '@/lib/api';
import { API } from '@/lib/api-routes';
import { DetailButton } from '@/components/shared/DetailButton';
import {
    FormDialog,
    FormRow,
    FormSection,
} from '@/components/shared/FormDialog';
import { ObjectDetailDrawer } from '@/components/shared/ObjectDetailDrawer';
import {
    NativeInput,
    NativeSelect,
} from '@/components/shared/FormControls';
import { useAuth } from '@/hooks/use-auth';

interface Ingredient {
    id?: number;
    name: string | null;
    content: string | null;
    cas: string | null;
}

interface HazardousChemical {
    id: number;
    name: string | null;
    alias: string | null;
    is_supervision: number;
    is_toxin: number;
    is_explosion: number;
    is_easy_explode: number;
    is_narcotics: number;
    is_monitor: number;
    is_mixing: number;
    cas: string | null;
    position: string | null;
    temperature: string | null;
    storage: string | null;
    year_use: string | null;
    year_use_unit: string | null;
    storage_unit: string | null;
    pressure: string | null;
    pressure_unit: string | null;
    fire_class: string | null;
    matter_state: string | null;
    msds: string | null;
    company_id: number;
    ingredients: Ingredient[];
    created_at: string;
    updated_at: string;
}

type FlagKey =
    | 'is_supervision'
    | 'is_toxin'
    | 'is_explosion'
    | 'is_easy_explode'
    | 'is_narcotics'
    | 'is_monitor'
    | 'is_mixing';

// Labels mirror the Laravel project (FLAG_FIELDS in legacy).
const flagLabels: Record<FlagKey, { label: string; color: string }> = {
    is_supervision: {
        label: '重点监管',
        color: 'border-rose-500/30 bg-rose-600/15 text-rose-300',
    },
    is_toxin: {
        label: '毒性物质',
        color: 'border-violet-500/30 bg-violet-600/20 text-violet-300',
    },
    is_explosion: {
        label: '爆炸品',
        color: 'border-amber-500/30 bg-amber-600/15 text-amber-300',
    },
    is_easy_explode: {
        label: '易燃易爆',
        color: 'border-amber-500/30 bg-amber-600/15 text-amber-300',
    },
    is_narcotics: {
        label: '麻醉品',
        color: 'border-violet-500/30 bg-violet-600/15 text-violet-300',
    },
    is_monitor: {
        label: '监控化学品',
        color: 'border-blue-500/30 bg-blue-600/15 text-blue-300',
    },
    is_mixing: {
        label: '混合物',
        color: 'border-cyan-500/30 bg-cyan-600/15 text-cyan-300',
    },
};

// Mirrors Laravel project's MATTER_STATES / FIRE_CLASSES / *_UNITS.
const MATTER_STATES = ['液态', '固态', '气态', '液化气体', '溶液', '粉末'];
const FIRE_CLASSES = ['甲类', '乙类', '丙类', '丁类', '戊类', '不适用'];
const STORAGE_UNITS = ['吨', 'kg', 'L', 'm³'];
const YEAR_USE_UNITS = ['吨/年', 'kg/年', 'L/年', 'm³/年'];
const PRESSURE_UNITS = ['MPa', 'kPa', 'Pa', 'bar', 'atm'];

type FormState = {
    name: string;
    alias: string;
    cas: string;
    position: string;
    temperature: string;
    storage: string;
    storage_unit: string;
    year_use: string;
    year_use_unit: string;
    pressure: string;
    pressure_unit: string;
    fire_class: string;
    matter_state: string;
    msds: string;
    flags: Record<FlagKey, boolean>;
    ingredients: Ingredient[];
};

const emptyForm: FormState = {
    name: '',
    alias: '',
    cas: '',
    position: '',
    temperature: '',
    storage: '',
    storage_unit: 'kg',
    year_use: '',
    year_use_unit: 'kg/年',
    pressure: '',
    pressure_unit: 'MPa',
    fire_class: '',
    matter_state: '',
    msds: '',
    flags: {
        is_supervision: false,
        is_toxin: false,
        is_explosion: false,
        is_easy_explode: false,
        is_narcotics: false,
        is_monitor: false,
        is_mixing: false,
    },
    ingredients: [],
};

export default function SafetyHazardousChemicals() {
    const { hasPermission } = useAuth();
    const [rows, setRows] = useState<HazardousChemical[]>([]);
    const [loading, setLoading] = useState(true);
    const [page, setPage] = useState(1);
    const [perPage] = useState(20);
    const [total, setTotal] = useState(0);
    const [keyword, setKeyword] = useState('');

    const [creating, setCreating] = useState(false);
    const [editing, setEditing] = useState<HazardousChemical | null>(null);
    const [viewing, setViewing] = useState<HazardousChemical | null>(null);
    const [deleting, setDeleting] = useState<HazardousChemical | null>(null);

    const canCreate = hasPermission('hazardous-chemicals.create');
    const canUpdate = hasPermission('hazardous-chemicals.update');
    const canDelete = hasPermission('hazardous-chemicals.delete');

    const refresh = useCallback(async () => {
        setLoading(true);
        try {
            const res = await api.get(API.hazardousChemicals(), {
                params: {
                    page,
                    per_page: perPage,
                    keyword: keyword || undefined,
                },
            });
            const d = res.data?.data ?? {};
            setRows(d.data ?? []);
            setTotal(d.total ?? 0);
        } catch (err) {
            toast.error('加载化学品失败');
            // eslint-disable-next-line no-console
            console.error(err);
        } finally {
            setLoading(false);
        }
    }, [page, perPage, keyword]);

    useEffect(() => {
        refresh();
    }, [refresh]);

    // ── Detail drawer wiring ──────────────────────────────────────
    //
    // Mirror the Laravel project's preprocessing:
    //   • join value+unit columns (year_use / storage / pressure)
    //   • render flag ints (0 / 1) as 是 / 否
    //   • drop the leftover *_unit keys so they don't render twice
    //   • hide the ingredients list when this isn't a mixture
    const detailData = useMemo<Record<string, unknown> | null>(() => {
        if (!viewing) return null;
        const join = (v?: string | null, u?: string | null) =>
            v ? `${v}${u ? ` ${u}` : ''}` : '';
        const d: Record<string, unknown> = {
            ...(viewing as unknown as Record<string, unknown>),
            year_use: join(viewing.year_use, viewing.year_use_unit),
            storage: join(viewing.storage, viewing.storage_unit),
            pressure: join(viewing.pressure, viewing.pressure_unit),
        };
        for (const k of Object.keys(flagLabels) as FlagKey[]) {
            d[k] = viewing[k] ? '是' : '否';
        }
        delete d.year_use_unit;
        delete d.storage_unit;
        delete d.pressure_unit;
        return d;
    }, [viewing]);
    const detailHiddenKeys = viewing?.is_mixing ? [] : ['ingredients'];

    return (
        <div>
            <div className="mb-6 flex items-center justify-between">
                <div>
                    <h1 className="text-xl font-bold text-slate-100">
                        危险化学品
                    </h1>
                    <p className="mt-0.5 text-xs text-slate-600">
                        共{' '}
                        <span className="font-medium text-slate-400">
                            {total}
                        </span>{' '}
                        条记录
                    </p>
                </div>
                {canCreate && (
                    <button
                        onClick={() => setCreating(true)}
                        className="flex items-center gap-2 rounded-lg border border-blue-500/30 bg-blue-600/15 px-4 py-2 text-sm font-medium text-blue-300 transition-all hover:border-blue-400/50 hover:bg-blue-600/25 hover:text-blue-200 active:scale-95"
                    >
                        <Plus className="h-4 w-4" />
                        新增危险化学品
                    </button>
                )}
            </div>

            <div className="mb-3 flex items-center gap-3">
                <div className="relative max-w-xs flex-1">
                    <Search className="pointer-events-none absolute top-1/2 left-3 h-3.5 w-3.5 -translate-y-1/2 text-slate-600" />
                    <input
                        value={keyword}
                        onChange={(e) => {
                            setKeyword(e.target.value);
                            setPage(1);
                        }}
                        placeholder="搜索名称 / 别名 / CAS"
                        className="h-9 w-full rounded-lg border border-blue-900/40 bg-[#070d1f] pr-3 pl-9 text-sm text-slate-300 placeholder-slate-700 transition-colors outline-none focus:border-blue-500/60 focus:ring-1 focus:ring-blue-500/20"
                    />
                </div>
                <button
                    onClick={refresh}
                    className="flex items-center gap-1.5 rounded-lg border border-blue-900/40 bg-[#070d1f] px-3 py-1.5 text-xs text-slate-400 hover:border-blue-500/40 hover:text-slate-200"
                >
                    <RefreshCw className="h-3.5 w-3.5" />
                    刷新
                </button>
            </div>

            <div className="overflow-hidden rounded-xl border border-blue-900/30 bg-[#0d1530]">
                {loading ? (
                    <div className="py-20 text-center">
                        <div className="flex flex-col items-center gap-3">
                            <RefreshCw className="h-5 w-5 animate-spin text-slate-600" />
                            <p className="text-sm text-slate-700">加载中…</p>
                        </div>
                    </div>
                ) : rows.length === 0 ? (
                    <div className="py-20 text-center">
                        <div className="flex flex-col items-center gap-3">
                            <div className="flex h-12 w-12 items-center justify-center rounded-xl border border-blue-900/30 bg-blue-600/8">
                                <FlaskConical className="h-5 w-5 text-slate-700" />
                            </div>
                            <p className="text-sm text-slate-700">暂无数据</p>
                        </div>
                    </div>
                ) : (
                    <div className="overflow-x-auto">
                        <table className="w-full text-sm">
                            <thead>
                                <tr className="border-b border-blue-900/40 bg-[#070d1f]/60">
                                    <th
                                        style={{ width: '60px' }}
                                        className="h-10 px-4 text-left align-middle text-[10px] font-semibold tracking-widest text-slate-600 uppercase first:pl-5"
                                    >
                                        ID
                                    </th>
                                    <th className="h-10 px-4 text-left align-middle text-[10px] font-semibold tracking-widest text-slate-600 uppercase">
                                        名称
                                    </th>
                                    <th className="h-10 px-4 text-left align-middle text-[10px] font-semibold tracking-widest text-slate-600 uppercase">
                                        监管标志
                                    </th>
                                    <th className="h-10 px-4 text-left align-middle text-[10px] font-semibold tracking-widest text-slate-600 uppercase">
                                        位置 / 状态
                                    </th>
                                    <th className="h-10 px-4 text-left align-middle text-[10px] font-semibold tracking-widest text-slate-600 uppercase">
                                        储量 / 用量
                                    </th>
                                    <th className="h-10 px-4 text-left align-middle text-[10px] font-semibold tracking-widest text-slate-600 uppercase">
                                        成分
                                    </th>
                                    <th className="h-10 px-4 text-right align-middle text-[10px] font-semibold tracking-widest text-slate-600 uppercase last:pr-5">
                                        操作
                                    </th>
                                </tr>
                            </thead>
                            <tbody>
                                {rows.map((r, idx) => {
                                    const flags = (
                                        Object.keys(flagLabels) as FlagKey[]
                                    ).filter((k) => r[k] === 1);
                                    return (
                                        <tr
                                            key={r.id}
                                            className={`group border-b border-blue-900/15 transition-colors last:border-0 hover:bg-blue-600/5 ${
                                                idx % 2 === 1
                                                    ? 'bg-white/[0.015]'
                                                    : ''
                                            }`}
                                        >
                                            <td className="px-4 py-3 align-middle text-slate-300 first:pl-5">
                                                {(page - 1) * perPage +
                                                    idx +
                                                    1}
                                            </td>
                                            <td className="px-4 py-3 align-middle text-slate-300">
                                                <div className="font-medium text-slate-100">
                                                    {r.name || '—'}
                                                </div>
                                                <div className="mt-0.5 font-mono text-[11px] text-slate-600">
                                                    #{r.id}
                                                    {r.cas && (
                                                        <> · CAS: {r.cas}</>
                                                    )}
                                                    {r.alias && <> · {r.alias}</>}
                                                </div>
                                            </td>
                                            <td className="px-4 py-3 align-middle">
                                                {flags.length === 0 ? (
                                                    <span className="text-[11px] text-slate-600">
                                                        —
                                                    </span>
                                                ) : (
                                                    <div className="flex flex-wrap gap-1">
                                                        {flags.map((f) => (
                                                            <span
                                                                key={f}
                                                                className={`rounded-md border px-1.5 py-0.5 text-[10px] ${flagLabels[f].color}`}
                                                            >
                                                                {
                                                                    flagLabels[f]
                                                                        .label
                                                                }
                                                            </span>
                                                        ))}
                                                    </div>
                                                )}
                                            </td>
                                            <td className="px-4 py-3 align-middle text-xs text-slate-400">
                                                <div>{r.position || '—'}</div>
                                                <div className="text-slate-600">
                                                    {r.matter_state || ''}
                                                </div>
                                            </td>
                                            <td className="px-4 py-3 align-middle text-xs text-slate-400">
                                                <div>
                                                    储量：{r.storage || '—'}
                                                    {r.storage_unit &&
                                                        ` ${r.storage_unit}`}
                                                </div>
                                                <div className="text-slate-600">
                                                    用量：{r.year_use || '—'}
                                                    {r.year_use_unit &&
                                                        ` ${r.year_use_unit}`}
                                                </div>
                                            </td>
                                            <td className="px-4 py-3 align-middle text-xs">
                                                {(r.ingredients ?? []).length ===
                                                0 ? (
                                                    <span className="text-slate-600">
                                                        —
                                                    </span>
                                                ) : (
                                                    <span className="rounded border border-slate-700/50 bg-slate-800/30 px-1.5 py-0.5 text-slate-400">
                                                        {r.ingredients.length} 项
                                                    </span>
                                                )}
                                            </td>
                                            <td className="px-4 py-3 text-right align-middle last:pr-5">
                                                <div className="flex items-center justify-end gap-1">
                                                    <DetailButton
                                                        row={r}
                                                        onOpen={setViewing}
                                                    />
                                                    {canUpdate && (
                                                        <button
                                                            onClick={() =>
                                                                setEditing(r)
                                                            }
                                                            title="编辑"
                                                            className="flex h-7 w-7 items-center justify-center rounded-md text-slate-600 transition-colors hover:bg-blue-600/10 hover:text-blue-400"
                                                        >
                                                            <Pencil className="h-3.5 w-3.5" />
                                                        </button>
                                                    )}
                                                    {canDelete && (
                                                        <button
                                                            onClick={() =>
                                                                setDeleting(r)
                                                            }
                                                            title="删除"
                                                            className="flex h-7 w-7 items-center justify-center rounded-md text-slate-600 transition-colors hover:bg-red-600/10 hover:text-red-400"
                                                        >
                                                            <Trash2 className="h-3.5 w-3.5" />
                                                        </button>
                                                    )}
                                                </div>
                                            </td>
                                        </tr>
                                    );
                                })}
                            </tbody>
                        </table>
                    </div>
                )}

                {total > perPage && (
                    <div className="flex items-center justify-between border-t border-blue-900/30 px-5 py-3">
                        <p className="text-xs text-slate-600">
                            共{' '}
                            <span className="font-medium text-slate-400">
                                {total}
                            </span>{' '}
                            条 &nbsp;·&nbsp; 第{' '}
                            <span className="font-medium text-slate-400">
                                {page}
                            </span>{' '}
                            <span className="text-slate-700">/</span>{' '}
                            <span className="font-medium text-slate-400">
                                {Math.ceil(total / perPage)}
                            </span>{' '}
                            页
                        </p>
                        <div className="flex items-center gap-1">
                            <button
                                onClick={() =>
                                    setPage((p) => Math.max(1, p - 1))
                                }
                                disabled={page <= 1}
                                className="rounded-md border border-blue-900/40 bg-[#070d1f] px-3 py-1 text-xs text-slate-400 hover:border-blue-500/40 hover:text-slate-200 disabled:cursor-not-allowed disabled:opacity-40"
                            >
                                上一页
                            </button>
                            <button
                                onClick={() => setPage((p) => p + 1)}
                                disabled={page >= Math.ceil(total / perPage)}
                                className="rounded-md border border-blue-900/40 bg-[#070d1f] px-3 py-1 text-xs text-slate-400 hover:border-blue-500/40 hover:text-slate-200 disabled:cursor-not-allowed disabled:opacity-40"
                            >
                                下一页
                            </button>
                        </div>
                    </div>
                )}
            </div>

            {creating && (
                <Editor
                    onClose={() => setCreating(false)}
                    onSaved={() => {
                        setCreating(false);
                        refresh();
                    }}
                />
            )}
            {editing && (
                <Editor
                    item={editing}
                    onClose={() => setEditing(null)}
                    onSaved={() => {
                        setEditing(null);
                        refresh();
                    }}
                />
            )}
            {deleting && (
                <DeleteConfirm
                    item={deleting}
                    onClose={() => setDeleting(null)}
                    onDeleted={() => {
                        setDeleting(null);
                        refresh();
                    }}
                />
            )}
            <ObjectDetailDrawer
                open={viewing !== null}
                onClose={() => setViewing(null)}
                title="危险化学品详情"
                data={detailData}
                labels={{
                    name: '化学品名称',
                    alias: '别名',
                    cas: 'CAS号',
                    matter_state: '物质状态',
                    fire_class: '火灾危险性类别',
                    position: '存储位置',
                    temperature: '操作温度',
                    storage: '最大存储量',
                    year_use: '年使用量',
                    pressure: '操作压力',
                    msds: 'MSDS附件',
                    is_supervision: '重点监管',
                    is_toxin: '毒性物质',
                    is_explosion: '爆炸品',
                    is_easy_explode: '易燃易爆',
                    is_narcotics: '麻醉品',
                    is_monitor: '监控化学品',
                    is_mixing: '混合物',
                    ingredients: '化学成分',
                    // Nested ingredient field labels — drawer renders array
                    // items as cards and uses the same labels map for the
                    // keys inside each item.
                    name_ingredient: '成分名称',
                    content: '含量',
                }}
                hiddenKeys={detailHiddenKeys}
            />
        </div>
    );
}

// ── Editor ────────────────────────────────────────────────────────

function Editor({
    item,
    onClose,
    onSaved,
}: {
    item?: HazardousChemical;
    onClose: () => void;
    onSaved: () => void;
}) {
    const isEdit = !!item;
    const [form, setForm] = useState<FormState>(() => {
        if (!item) return { ...emptyForm, flags: { ...emptyForm.flags } };
        return {
            name: item.name ?? '',
            alias: item.alias ?? '',
            cas: item.cas ?? '',
            position: item.position ?? '',
            temperature: item.temperature ?? '',
            storage: item.storage ?? '',
            storage_unit: item.storage_unit ?? 'kg',
            year_use: item.year_use ?? '',
            year_use_unit: item.year_use_unit ?? 'kg/年',
            pressure: item.pressure ?? '',
            pressure_unit: item.pressure_unit ?? 'MPa',
            fire_class: item.fire_class ?? '',
            matter_state: item.matter_state ?? '',
            msds: item.msds ?? '',
            flags: {
                is_supervision: item.is_supervision === 1,
                is_toxin: item.is_toxin === 1,
                is_explosion: item.is_explosion === 1,
                is_easy_explode: item.is_easy_explode === 1,
                is_narcotics: item.is_narcotics === 1,
                is_monitor: item.is_monitor === 1,
                is_mixing: item.is_mixing === 1,
            },
            ingredients: (item.ingredients ?? []).map((i) => ({
                id: i.id,
                name: i.name,
                content: i.content,
                cas: i.cas,
            })),
        };
    });
    const [busy, setBusy] = useState(false);

    const set = <K extends keyof FormState>(k: K, v: FormState[K]) =>
        setForm((f) => ({ ...f, [k]: v }));
    const toggleFlag = (k: FlagKey) =>
        setForm((f) => ({ ...f, flags: { ...f.flags, [k]: !f.flags[k] } }));

    const addIngredient = () =>
        setForm((f) => ({
            ...f,
            ingredients: [
                ...f.ingredients,
                { name: '', content: '', cas: '' },
            ],
        }));
    const removeIngredient = (i: number) =>
        setForm((f) => ({
            ...f,
            ingredients: f.ingredients.filter((_, idx) => idx !== i),
        }));
    const updateIngredient = (i: number, k: keyof Ingredient, v: string) =>
        setForm((f) => ({
            ...f,
            ingredients: f.ingredients.map((ing, idx) =>
                idx === i ? { ...ing, [k]: v } : ing,
            ),
        }));

    const submit = async () => {
        if (!form.name.trim()) {
            toast.error('化学品名必填');
            return;
        }
        setBusy(true);
        try {
            const payload = {
                name: form.name,
                alias: form.alias || null,
                cas: form.cas || null,
                position: form.position || null,
                temperature: form.temperature || null,
                storage: form.storage || null,
                storage_unit: form.storage_unit || null,
                year_use: form.year_use || null,
                year_use_unit: form.year_use_unit || null,
                pressure: form.pressure || null,
                pressure_unit: form.pressure_unit || null,
                fire_class: form.fire_class || null,
                matter_state: form.matter_state || null,
                msds: form.msds || null,
                is_supervision: form.flags.is_supervision ? 1 : 0,
                is_toxin: form.flags.is_toxin ? 1 : 0,
                is_explosion: form.flags.is_explosion ? 1 : 0,
                is_easy_explode: form.flags.is_easy_explode ? 1 : 0,
                is_narcotics: form.flags.is_narcotics ? 1 : 0,
                is_monitor: form.flags.is_monitor ? 1 : 0,
                is_mixing: form.flags.is_mixing ? 1 : 0,
                ingredients: form.ingredients.map((i) => ({
                    name: i.name,
                    content: i.content,
                    cas: i.cas,
                })),
            };
            if (isEdit) {
                await api.put(API.hazardousChemicalById(item!.id), payload);
            } else {
                await api.post(API.hazardousChemicals(), payload);
            }
            toast.success(isEdit ? '已保存' : '已创建');
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
            title={
                isEdit
                    ? `编辑「${item!.name || `#${item!.id}`}」`
                    : '新建危险化学品'
            }
            description={
                isEdit ? `#${item!.id} · 修改信息 / 监管标志 / 成分` : ''
            }
            size="lg"
            onConfirm={submit}
            confirmLabel={isEdit ? '保存' : '创建'}
            loading={busy}
        >
            <div className="space-y-6">
                <FormSection title="基本信息" icon={FlaskConical}>
                    <div className="grid grid-cols-2 gap-4">
                        <FormRow label="化学品名称" required>
                            <NativeInput
                                value={form.name}
                                onChange={(e) => set('name', e.target.value)}
                                placeholder="化学品名称"
                                autoFocus
                            />
                        </FormRow>
                        <FormRow label="别名">
                            <NativeInput
                                value={form.alias}
                                onChange={(e) => set('alias', e.target.value)}
                                placeholder="别名"
                            />
                        </FormRow>
                        <FormRow label="CAS 号">
                            <NativeInput
                                value={form.cas}
                                onChange={(e) => set('cas', e.target.value)}
                                placeholder="如 74-82-8"
                                className="font-mono"
                            />
                        </FormRow>
                        <FormRow label="存储位置">
                            <NativeInput
                                value={form.position}
                                onChange={(e) =>
                                    set('position', e.target.value)
                                }
                                placeholder="存储位置/罐区/仓库编号"
                            />
                        </FormRow>
                    </div>
                </FormSection>

                <FormSection title="监管属性" icon={ShieldAlert}>
                    <div className="flex flex-wrap gap-2">
                        {(Object.keys(flagLabels) as FlagKey[]).map((k) => {
                            const on = form.flags[k];
                            const meta = flagLabels[k];
                            return (
                                <button
                                    key={k}
                                    type="button"
                                    onClick={() => toggleFlag(k)}
                                    className={`rounded-md border px-2.5 py-1 text-xs transition-all ${
                                        on
                                            ? meta.color
                                            : 'border-slate-700/50 bg-slate-800/30 text-slate-500 hover:border-slate-600 hover:text-slate-300'
                                    }`}
                                >
                                    {meta.label}
                                </button>
                            );
                        })}
                    </div>
                </FormSection>

                <FormSection title="储存与使用" icon={Flame}>
                    <div className="grid grid-cols-2 gap-4">
                        <FormRow label="操作温度">
                            <NativeInput
                                value={form.temperature}
                                onChange={(e) =>
                                    set('temperature', e.target.value)
                                }
                                placeholder="如 25"
                            />
                        </FormRow>
                        <FormRow label="物质状态">
                            <NativeSelect
                                value={form.matter_state}
                                onChange={(e) =>
                                    set('matter_state', e.target.value)
                                }
                                placeholder="请选择"
                                options={MATTER_STATES}
                            />
                        </FormRow>

                        <FormRow label="最大存储量">
                            <div className="flex gap-2">
                                <NativeInput
                                    value={form.storage}
                                    onChange={(e) =>
                                        set('storage', e.target.value)
                                    }
                                    placeholder="数值"
                                    className="flex-1"
                                />
                                <div className="w-28 shrink-0">
                                    <NativeSelect
                                        value={form.storage_unit}
                                        onChange={(e) =>
                                            set('storage_unit', e.target.value)
                                        }
                                        placeholder="单位"
                                        options={STORAGE_UNITS}
                                    />
                                </div>
                            </div>
                        </FormRow>

                        <FormRow label="年使用量">
                            <div className="flex gap-2">
                                <NativeInput
                                    value={form.year_use}
                                    onChange={(e) =>
                                        set('year_use', e.target.value)
                                    }
                                    placeholder="数值"
                                    className="flex-1"
                                />
                                <div className="w-28 shrink-0">
                                    <NativeSelect
                                        value={form.year_use_unit}
                                        onChange={(e) =>
                                            set('year_use_unit', e.target.value)
                                        }
                                        placeholder="单位"
                                        options={YEAR_USE_UNITS}
                                    />
                                </div>
                            </div>
                        </FormRow>

                        <FormRow label="操作压力">
                            <div className="flex gap-2">
                                <NativeInput
                                    value={form.pressure}
                                    onChange={(e) =>
                                        set('pressure', e.target.value)
                                    }
                                    placeholder="数值"
                                    className="flex-1"
                                />
                                <div className="w-28 shrink-0">
                                    <NativeSelect
                                        value={form.pressure_unit}
                                        onChange={(e) =>
                                            set('pressure_unit', e.target.value)
                                        }
                                        placeholder="单位"
                                        options={PRESSURE_UNITS}
                                    />
                                </div>
                            </div>
                        </FormRow>

                        <FormRow label="火灾危险性类别">
                            <NativeSelect
                                value={form.fire_class}
                                onChange={(e) =>
                                    set('fire_class', e.target.value)
                                }
                                placeholder="请选择"
                                options={FIRE_CLASSES}
                            />
                        </FormRow>

                        <FormRow label="MSDS 附件" colSpan={2}>
                            <NativeInput
                                value={form.msds}
                                onChange={(e) => set('msds', e.target.value)}
                                placeholder="可填 OSS 文件 URL"
                            />
                        </FormRow>
                    </div>
                </FormSection>

                {/* 化学成分仅在「混合物」标志开启时显示，跟 legacy 一致 */}
                {form.flags.is_mixing && (
                    <FormSection title="主要成分" icon={Skull}>
                        <div className="mb-3 flex items-center justify-between">
                            <span className="text-[11px] text-slate-600">
                                共 {form.ingredients.length} 项
                            </span>
                            <button
                                type="button"
                                onClick={addIngredient}
                                className="flex items-center gap-1 rounded-md border border-blue-500/30 bg-blue-600/15 px-2 py-1 text-[11px] text-blue-300 hover:border-blue-500/50 hover:bg-blue-600/25"
                            >
                                <Plus className="h-3 w-3" />
                                添加成分
                            </button>
                        </div>
                        {form.ingredients.length === 0 ? (
                            <p className="rounded-md border border-blue-900/30 bg-[#0a1428] px-3 py-3 text-center text-[11px] text-slate-600">
                                混合物登记每个组分的名称、占比与 CAS。点击右上角添加。
                            </p>
                        ) : (
                            <div className="space-y-2">
                                {form.ingredients.map((ing, i) => (
                                    <div
                                        key={i}
                                        className="grid grid-cols-12 items-start gap-2 rounded-lg border border-blue-900/30 bg-[#0a1428] p-2.5"
                                    >
                                        <div className="col-span-4">
                                            <NativeInput
                                                value={ing.name ?? ''}
                                                onChange={(e) =>
                                                    updateIngredient(
                                                        i,
                                                        'name',
                                                        e.target.value,
                                                    )
                                                }
                                                placeholder="成分名"
                                            />
                                        </div>
                                        <div className="col-span-3">
                                            <NativeInput
                                                value={ing.content ?? ''}
                                                onChange={(e) =>
                                                    updateIngredient(
                                                        i,
                                                        'content',
                                                        e.target.value,
                                                    )
                                                }
                                                placeholder="含量"
                                            />
                                        </div>
                                        <div className="col-span-4">
                                            <NativeInput
                                                value={ing.cas ?? ''}
                                                onChange={(e) =>
                                                    updateIngredient(
                                                        i,
                                                        'cas',
                                                        e.target.value,
                                                    )
                                                }
                                                placeholder="CAS"
                                            />
                                        </div>
                                        <div className="col-span-1 flex justify-end">
                                            <button
                                                type="button"
                                                onClick={() =>
                                                    removeIngredient(i)
                                                }
                                                className="flex h-8 w-8 items-center justify-center rounded-md text-slate-600 hover:bg-rose-600/15 hover:text-rose-400"
                                            >
                                                <X className="h-3.5 w-3.5" />
                                            </button>
                                        </div>
                                    </div>
                                ))}
                            </div>
                        )}
                    </FormSection>
                )}
            </div>
        </FormDialog>
    );
}

function DeleteConfirm({
    item,
    onClose,
    onDeleted,
}: {
    item: HazardousChemical;
    onClose: () => void;
    onDeleted: () => void;
}) {
    const [busy, setBusy] = useState(false);
    const handle = async () => {
        setBusy(true);
        try {
            await api.delete(API.hazardousChemicalById(item.id));
            toast.success('已删除');
            onDeleted();
        } catch (err) {
            toast.error('删除失败');
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
            title={`删除「${item.name || `#${item.id}`}」`}
            icon={Trash2}
            size="md"
            onConfirm={handle}
            confirmLabel="确认删除"
            confirmTone="rose"
            loading={busy}
        >
            <p className="text-sm leading-relaxed text-slate-400">
                软删除：化学品记录被标记为已删除（其下的成分会随之级联删除）。
                如有必要可联系运维恢复。
            </p>
        </FormDialog>
    );
}
