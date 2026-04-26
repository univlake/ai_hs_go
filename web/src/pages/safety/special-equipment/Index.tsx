import { useCallback, useEffect, useMemo, useRef, useState } from 'react';
import {
    Download,
    Edit3,
    Eye,
    FileSpreadsheet,
    Plus,
    RefreshCw,
    Search,
    Trash2,
    Upload,
    Wrench,
} from 'lucide-react';
import { toast } from 'sonner';
import api from '@/lib/api';
import { API } from '@/lib/api-routes';
import { Button } from '@/components/ui/button';
import { FormDialog, FormRow } from '@/components/shared/FormDialog';
import { ObjectDetailDrawer } from '@/components/shared/ObjectDetailDrawer';
import { DatePicker } from '@/components/shared/DatePicker';
import {
    NativeInput,
    NativeSelect,
    NativeTextarea,
} from '@/components/shared/FormControls';
import { useAuth } from '@/hooks/use-auth';

// ── Spec field config (mirrors Laravel project) ──────────────────
// Map special_type id → list of dynamic spec fields (norm_model1..4 with
// labels/units/select options). Types not listed get no dynamic fields.

interface SpecFieldDef {
    key: 'norm_model1' | 'norm_model2' | 'norm_model3' | 'norm_model4';
    label: string;
    unit?: string;
    inputType: 'text' | 'select';
    options?: { value: string; label: string }[];
}

const SPEC_CONFIGS: Record<number, SpecFieldDef[]> = {
    // 2 = 压力容器(含气瓶)
    2: [
        { key: 'norm_model1', label: '设计压力', unit: 'MPa', inputType: 'text' },
        { key: 'norm_model2', label: '容积', unit: 'm³', inputType: 'text' },
        { key: 'norm_model3', label: '盛放介质', inputType: 'text' },
    ],
    // 3 = 压力管道
    3: [
        { key: 'norm_model1', label: '设计压力', unit: 'MPa', inputType: 'text' },
        { key: 'norm_model2', label: '设计温度', unit: '°C', inputType: 'text' },
        { key: 'norm_model3', label: '管径', unit: 'cm', inputType: 'text' },
        { key: 'norm_model4', label: '输送介质', inputType: 'text' },
    ],
    // 4 = 电梯
    4: [
        {
            key: 'norm_model1',
            label: '用途',
            inputType: 'select',
            options: [
                { value: '载人', label: '载人' },
                { value: '载重', label: '载重' },
            ],
        },
        { key: 'norm_model2', label: '载重', unit: '吨', inputType: 'text' },
    ],
};

interface SpecialType {
    id: number;
    title: string;
}

// Static portion of the detail labels — `norm_model{1..4}` are filled in at
// render time using the active record's `special_type_id` so the detail
// shows e.g. "设计压力" / "容积" rather than "规格型号 1".
const baseSpecialEquipmentLabels: Record<string, string> = {
    equipment_name: '设备名称',
    special_type: '设备类型',
    title: '类型名称',
    utilized_location: '使用位置',
    serial_number: '使用证编号/场内机动车牌证',
    use_time: '投用时间',
    valid_period: '检验有效期',
    inspect_unit: '检验单位',
    maintain_info: '维护信息',
    comment: '备注',
};

const ALL_SPEC_KEYS: SpecFieldDef['key'][] = [
    'norm_model1',
    'norm_model2',
    'norm_model3',
    'norm_model4',
];

interface SpecialEquipment {
    id: number;
    equipment_name: string | null;
    special_type_id: number | null;
    utilized_location: string | null;
    norm_model1: string | null;
    norm_model2: string | null;
    norm_model3: string | null;
    norm_model4: string | null;
    serial_number: string | null;
    use_time: string | null;
    valid_period: string | null;
    inspect_unit: string | null;
    maintain_info: string | null;
    comment: string | null;
    company_id: number;
    created_at: string;
    updated_at: string;
}

type FormState = Omit<SpecialEquipment, 'id' | 'company_id' | 'created_at' | 'updated_at'>;

const emptyForm: FormState = {
    equipment_name: '',
    special_type_id: null,
    utilized_location: '',
    norm_model1: '',
    norm_model2: '',
    norm_model3: '',
    norm_model4: '',
    serial_number: '',
    use_time: '',
    valid_period: '',
    inspect_unit: '',
    maintain_info: '',
    comment: '',
};

function StatTile({
    label,
    value,
    icon: Icon,
    accent,
}: {
    label: string;
    value: number | string;
    icon: React.ComponentType<{ className?: string }>;
    accent: 'sky' | 'amber' | 'emerald';
}) {
    const tone = {
        sky: 'border-sky-600/20 bg-sky-600/15 text-sky-400',
        amber: 'border-amber-600/20 bg-amber-600/15 text-amber-400',
        emerald: 'border-emerald-600/20 bg-emerald-600/15 text-emerald-400',
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

function fmtDate(d?: string | null) {
    if (!d) return '—';
    return d.slice(0, 10);
}

function isExpiringSoon(d?: string | null) {
    if (!d) return false;
    const date = new Date(d);
    const now = new Date();
    const days = (date.getTime() - now.getTime()) / (1000 * 60 * 60 * 24);
    return days >= 0 && days <= 30;
}

function isExpired(d?: string | null) {
    if (!d) return false;
    return new Date(d).getTime() < Date.now();
}

export default function SafetySpecialEquipment() {
    const { hasPermission } = useAuth();
    const [rows, setRows] = useState<SpecialEquipment[]>([]);
    const [loading, setLoading] = useState(true);
    const [page, setPage] = useState(1);
    const [perPage] = useState(20);
    const [total, setTotal] = useState(0);
    const [keyword, setKeyword] = useState('');

    const [creating, setCreating] = useState(false);
    const [editing, setEditing] = useState<SpecialEquipment | null>(null);
    const [viewing, setViewing] = useState<SpecialEquipment | null>(null);
    const [deleting, setDeleting] = useState<SpecialEquipment | null>(null);
    const [specialTypes, setSpecialTypes] = useState<SpecialType[]>([]);
    const fileRef = useRef<HTMLInputElement>(null);

    const canCreate = hasPermission('special-equipment.create');
    const canUpdate = hasPermission('special-equipment.update');
    const canDelete = hasPermission('special-equipment.delete');

    const refresh = useCallback(async () => {
        setLoading(true);
        try {
            const res = await api.get(API.specialEquipment(), {
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
            toast.error('加载特种设备失败');
            // eslint-disable-next-line no-console
            console.error(err);
        } finally {
            setLoading(false);
        }
    }, [page, perPage, keyword]);

    useEffect(() => {
        refresh();
    }, [refresh]);

    useEffect(() => {
        api.get(API.specialTypes())
            .then((res) =>
                setSpecialTypes((res.data?.data as SpecialType[]) ?? []),
            )
            .catch(() => {
                // dictionary lookup is non-critical — form just shows an
                // empty type select if it fails
            });
    }, []);

    const stats = useMemo(
        () => ({
            total,
            expiring: rows.filter((r) => isExpiringSoon(r.valid_period))
                .length,
            expired: rows.filter((r) => isExpired(r.valid_period)).length,
        }),
        [rows, total],
    );

    // ── Excel actions ──────────────────────────────────────────

    const downloadTemplate = async () => {
        try {
            const res = await api.get(API.specialEquipmentTemplate(), {
                responseType: 'blob',
            });
            triggerDownload(res.data, '特种设备导入模板.xlsx');
        } catch {
            toast.error('下载模板失败');
        }
    };

    const downloadExport = async () => {
        try {
            const res = await api.get(API.specialEquipmentExport(), {
                responseType: 'blob',
            });
            triggerDownload(res.data, '特种设备台账.xlsx');
        } catch {
            toast.error('导出失败');
        }
    };

    const triggerImport = () => fileRef.current?.click();

    const onImportSelected: React.ChangeEventHandler<HTMLInputElement> = async (
        e,
    ) => {
        const file = e.target.files?.[0];
        if (!file) return;
        const fd = new FormData();
        fd.append('file', file);
        try {
            // Force multipart Content-Type — axios instance has a default
            // application/json header that would serialize FormData as JSON
            // and cause Gin's c.FormFile("file") to fail with "missing file".
            const res = await api.post(API.specialEquipmentImport(), fd, {
                headers: { 'Content-Type': 'multipart/form-data' },
            });
            const d = res.data?.data ?? {};
            toast.success(
                `已导入 ${d.imported ?? 0} / ${d.total ?? 0} 条`,
            );
            refresh();
        } catch (err) {
            const msg =
                (err as { response?: { data?: { message?: string } } }).response
                    ?.data?.message ?? '导入失败';
            toast.error(msg);
        } finally {
            if (fileRef.current) fileRef.current.value = '';
        }
    };

    // ── Detail drawer wiring ──────────────────────────────────────
    //
    // Spec fields (norm_model1..4) get type-aware labels and we hide the
    // ones that don't apply to the equipment's special_type. We also synth
    // a `special_type` field that the drawer renders as plain text — Go's
    // backend only returns the numeric id, so the page-level component
    // owns the lookup against the already-loaded specialTypes list.
    const detailSpecFields: SpecFieldDef[] =
        viewing?.special_type_id != null
            ? SPEC_CONFIGS[viewing.special_type_id] ?? []
            : [];
    const detailSpecLabels: Record<string, string> = Object.fromEntries(
        detailSpecFields.map((f) => [
            f.key,
            f.unit ? `${f.label} (${f.unit})` : f.label,
        ]),
    );
    const detailLabels: Record<string, string> = {
        ...baseSpecialEquipmentLabels,
        ...detailSpecLabels,
    };
    const detailHiddenKeys =
        viewing?.special_type_id != null
            ? ALL_SPEC_KEYS.filter(
                  (k) => !detailSpecFields.some((f) => f.key === k),
              )
            : [];
    const detailData: Record<string, unknown> | null = useMemo(() => {
        if (!viewing) return null;
        const typeTitle = specialTypes.find(
            (t) => t.id === viewing.special_type_id,
        )?.title;
        const out: Record<string, unknown> = { ...viewing };
        if (typeTitle) out.special_type = typeTitle;
        return out;
    }, [viewing, specialTypes]);

    return (
        <div className="space-y-6">
            <div className="flex items-start justify-between">
                <div>
                    <h1 className="text-2xl font-bold text-slate-100">
                        特种设备
                    </h1>
                    <p className="mt-1 text-sm text-slate-500">
                        当前租户的特种设备台账 · 含使用证 / 投用 / 检验信息
                    </p>
                </div>
                <div className="flex gap-2">
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
                        variant="outline"
                        size="sm"
                        onClick={downloadTemplate}
                        className="border-blue-900/40 bg-[#0d1530]"
                    >
                        <FileSpreadsheet className="mr-1.5 h-3.5 w-3.5" />
                        模板
                    </Button>
                    <Button
                        variant="outline"
                        size="sm"
                        onClick={downloadExport}
                        className="border-blue-900/40 bg-[#0d1530]"
                    >
                        <Download className="mr-1.5 h-3.5 w-3.5" />
                        导出
                    </Button>
                    {canCreate && (
                        <>
                            <input
                                ref={fileRef}
                                type="file"
                                accept=".xlsx,.xls"
                                hidden
                                onChange={onImportSelected}
                            />
                            <Button
                                variant="outline"
                                size="sm"
                                onClick={triggerImport}
                                className="border-blue-900/40 bg-[#0d1530]"
                            >
                                <Upload className="mr-1.5 h-3.5 w-3.5" />
                                导入
                            </Button>
                            <Button
                                onClick={() => setCreating(true)}
                                className="bg-blue-600 hover:bg-blue-500"
                            >
                                <Plus className="mr-1.5 h-3.5 w-3.5" />
                                新建
                            </Button>
                        </>
                    )}
                </div>
            </div>

            <div className="grid gap-4 md:grid-cols-3">
                <StatTile
                    label="设备总数"
                    value={stats.total}
                    icon={Wrench}
                    accent="sky"
                />
                <StatTile
                    label="即将到期 (≤30天)"
                    value={stats.expiring}
                    icon={Wrench}
                    accent="amber"
                />
                <StatTile
                    label="已过期"
                    value={stats.expired}
                    icon={Wrench}
                    accent="emerald"
                />
            </div>

            <div className="flex items-center gap-3">
                <div className="relative max-w-sm flex-1">
                    <Search className="absolute top-1/2 left-3 h-4 w-4 -translate-y-1/2 text-slate-600" />
                    <input
                        value={keyword}
                        onChange={(e) => {
                            setKeyword(e.target.value);
                            setPage(1);
                        }}
                        placeholder="搜索设备名 / 使用证编号"
                        className="w-full rounded-lg border border-blue-900/40 bg-[#0d1530] py-2 pr-3 pl-9 text-sm text-slate-200 placeholder:text-slate-600 focus:border-blue-500/40 focus:outline-none"
                    />
                </div>
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
                        <Wrench className="mx-auto mb-3 h-10 w-10 text-slate-700" />
                        <p className="text-sm text-slate-400">暂无设备</p>
                        {canCreate && (
                            <p className="mt-1 text-xs text-slate-600">
                                点击右上角「新建」开始录入
                            </p>
                        )}
                    </div>
                ) : (
                    <table className="w-full text-sm">
                        <thead>
                            <tr className="border-b border-blue-900/30 bg-[#0a1428] text-[10px] tracking-widest text-slate-600 uppercase">
                                <th className="px-5 py-3 text-left">设备 / 使用证</th>
                                <th className="px-5 py-3 text-left">特种类型</th>
                                <th className="px-5 py-3 text-left">使用位置</th>
                                <th className="px-5 py-3 text-left">规格型号</th>
                                <th className="px-5 py-3 text-left">投用 / 有效期</th>
                                <th className="px-5 py-3 text-left">检验单位</th>
                                <th className="px-5 py-3 text-right">操作</th>
                            </tr>
                        </thead>
                        <tbody>
                            {rows.map((r) => {
                                const expired = isExpired(r.valid_period);
                                const expiring = !expired && isExpiringSoon(r.valid_period);
                                return (
                                    <tr
                                        key={r.id}
                                        className="border-b border-blue-900/20 transition-colors last:border-b-0 hover:bg-[#0a1428]"
                                    >
                                        <td className="px-5 py-3">
                                            <div className="font-medium text-slate-100">
                                                {r.equipment_name || '—'}
                                            </div>
                                            <div className="mt-0.5 font-mono text-[11px] text-slate-600">
                                                #{r.id}
                                                {r.serial_number && (
                                                    <> · {r.serial_number}</>
                                                )}
                                            </div>
                                        </td>
                                        <td className="px-5 py-3 text-xs">
                                            {r.special_type_id ? (
                                                <span className="rounded-md border border-blue-500/30 bg-blue-600/15 px-1.5 py-0.5 text-blue-300">
                                                    {specialTypes.find(
                                                        (t) =>
                                                            t.id ===
                                                            r.special_type_id,
                                                    )?.title ??
                                                        `#${r.special_type_id}`}
                                                </span>
                                            ) : (
                                                <span className="text-slate-600">
                                                    —
                                                </span>
                                            )}
                                        </td>
                                        <td className="px-5 py-3 text-xs text-slate-400">
                                            {r.utilized_location || '—'}
                                        </td>
                                        <td className="px-5 py-3 text-xs text-slate-400">
                                            <div className="space-y-0.5">
                                                {[
                                                    r.norm_model1,
                                                    r.norm_model2,
                                                    r.norm_model3,
                                                    r.norm_model4,
                                                ]
                                                    .filter(Boolean)
                                                    .map((m, i) => (
                                                        <div key={i}>{m}</div>
                                                    ))}
                                                {!r.norm_model1 && '—'}
                                            </div>
                                        </td>
                                        <td className="px-5 py-3 text-xs">
                                            <div className="text-slate-400">
                                                投用：{fmtDate(r.use_time)}
                                            </div>
                                            <div
                                                className={
                                                    expired
                                                        ? 'text-rose-400'
                                                        : expiring
                                                          ? 'text-amber-400'
                                                          : 'text-slate-500'
                                                }
                                            >
                                                有效期：{fmtDate(r.valid_period)}
                                                {expired && ' · 已过期'}
                                                {expiring && ' · 即将到期'}
                                            </div>
                                        </td>
                                        <td className="px-5 py-3 text-xs text-slate-400">
                                            {r.inspect_unit || '—'}
                                        </td>
                                        <td className="px-5 py-3 text-right">
                                            <div className="flex justify-end gap-1">
                                                <button
                                                    onClick={() => setViewing(r)}
                                                    title="查看详情"
                                                    className="rounded-md border border-blue-500/30 bg-blue-600/10 p-1.5 text-blue-400 hover:border-blue-500/50 hover:bg-blue-600/20"
                                                >
                                                    <Eye className="h-3.5 w-3.5" />
                                                </button>
                                                {canUpdate && (
                                                    <button
                                                        onClick={() => setEditing(r)}
                                                        title="编辑"
                                                        className="rounded-md border border-slate-700/50 bg-slate-800/30 p-1.5 text-slate-400 hover:border-slate-600 hover:text-slate-200"
                                                    >
                                                        <Edit3 className="h-3.5 w-3.5" />
                                                    </button>
                                                )}
                                                {canDelete && (
                                                    <button
                                                        onClick={() => setDeleting(r)}
                                                        title="删除"
                                                        className="rounded-md border border-rose-500/30 bg-rose-600/10 p-1.5 text-rose-400 hover:border-rose-500/50 hover:bg-rose-600/20"
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
                <SpecialEquipmentEditor
                    specialTypes={specialTypes}
                    onClose={() => setCreating(false)}
                    onSaved={() => {
                        setCreating(false);
                        refresh();
                    }}
                />
            )}
            {editing && (
                <SpecialEquipmentEditor
                    equipment={editing}
                    specialTypes={specialTypes}
                    onClose={() => setEditing(null)}
                    onSaved={() => {
                        setEditing(null);
                        refresh();
                    }}
                />
            )}
            {deleting && (
                <DeleteConfirm
                    equipment={deleting}
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
                title="特种设备详情"
                data={detailData}
                labels={detailLabels}
                hiddenKeys={detailHiddenKeys}
            />
        </div>
    );
}

// ── Editor ────────────────────────────────────────────────────────

function SpecialEquipmentEditor({
    equipment,
    specialTypes,
    onClose,
    onSaved,
}: {
    equipment?: SpecialEquipment;
    specialTypes: SpecialType[];
    onClose: () => void;
    onSaved: () => void;
}) {
    const isEdit = !!equipment;
    const [form, setForm] = useState<FormState>(() =>
        equipment
            ? {
                  equipment_name: equipment.equipment_name ?? '',
                  special_type_id: equipment.special_type_id ?? null,
                  utilized_location: equipment.utilized_location ?? '',
                  norm_model1: equipment.norm_model1 ?? '',
                  norm_model2: equipment.norm_model2 ?? '',
                  norm_model3: equipment.norm_model3 ?? '',
                  norm_model4: equipment.norm_model4 ?? '',
                  serial_number: equipment.serial_number ?? '',
                  use_time: (equipment.use_time ?? '').slice(0, 10),
                  valid_period: (equipment.valid_period ?? '').slice(0, 10),
                  inspect_unit: equipment.inspect_unit ?? '',
                  maintain_info: equipment.maintain_info ?? '',
                  comment: equipment.comment ?? '',
              }
            : { ...emptyForm },
    );
    const [busy, setBusy] = useState(false);

    const set = <K extends keyof FormState>(k: K, v: FormState[K]) =>
        setForm((f) => ({ ...f, [k]: v }));

    // Reset spec fields when type changes — irrelevant fields would otherwise
    // get persisted and surface in the next type's form by accident.
    const handleTypeChange = (idStr: string) => {
        const id = idStr ? Number(idStr) : null;
        setForm((f) => ({
            ...f,
            special_type_id: id,
            norm_model1: '',
            norm_model2: '',
            norm_model3: '',
            norm_model4: '',
        }));
    };

    const submit = async () => {
        if (!form.equipment_name?.toString().trim()) {
            toast.error('设备名称必填');
            return;
        }
        if (!form.special_type_id) {
            toast.error('请选择特种类型');
            return;
        }
        setBusy(true);
        try {
            const payload = {
                ...form,
                use_time: form.use_time || null,
                valid_period: form.valid_period || null,
            };
            if (isEdit) {
                await api.put(API.specialEquipmentById(equipment!.id), payload);
            } else {
                await api.post(API.specialEquipment(), payload);
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

    const specs = form.special_type_id
        ? SPEC_CONFIGS[form.special_type_id] ?? []
        : [];

    return (
        <FormDialog
            open
            onOpenChange={(v) => !v && onClose()}
            title={isEdit ? `编辑「${equipment!.equipment_name || `#${equipment!.id}`}」` : '新建特种设备'}
            description={
                isEdit
                    ? `#${equipment!.id} · 修改设备信息 / 检验状态`
                    : '记录一台新的特种设备到本租户台账'
            }
            size="lg"
            onConfirm={submit}
            confirmLabel={isEdit ? '保存' : '创建'}
            loading={busy}
        >
            <div className="grid grid-cols-2 gap-4">
                <FormRow label="设备名称" required colSpan={2}>
                    <NativeInput
                        value={form.equipment_name ?? ''}
                        onChange={(e) => set('equipment_name', e.target.value)}
                        placeholder="请输入设备名称"
                        autoFocus
                    />
                </FormRow>

                <FormRow label="设备类型">
                    <NativeSelect
                        value={
                            form.special_type_id
                                ? String(form.special_type_id)
                                : ''
                        }
                        onChange={(e) => handleTypeChange(e.target.value)}
                        placeholder="请选择类型"
                        options={specialTypes.map((t) => ({
                            value: String(t.id),
                            label: t.title,
                        }))}
                    />
                </FormRow>

                <FormRow label="使用位置">
                    <NativeInput
                        value={form.utilized_location ?? ''}
                        onChange={(e) =>
                            set('utilized_location', e.target.value)
                        }
                        placeholder="设备使用位置"
                    />
                </FormRow>

                <FormRow
                    label="使用证编号/场内机动车牌证"
                    colSpan={2}
                    labelTrailing={
                        isEdit && !form.serial_number ? (
                            <span className="ml-2 text-amber-500">
                                请填写唯一编号
                            </span>
                        ) : undefined
                    }
                >
                    <NativeInput
                        value={form.serial_number ?? ''}
                        onChange={(e) => set('serial_number', e.target.value)}
                        placeholder="请输入编号（每台设备唯一）"
                        className={
                            isEdit && !form.serial_number
                                ? 'border-amber-500/50 focus:border-amber-500/70'
                                : ''
                        }
                    />
                </FormRow>

                {/* Spec fields — driven by SPEC_CONFIGS[special_type_id]. */}
                {specs.map((s) => (
                    <FormRow
                        key={s.key}
                        label={s.unit ? `${s.label} (${s.unit})` : s.label}
                    >
                        {s.inputType === 'select' ? (
                            <NativeSelect
                                value={form[s.key] ?? ''}
                                onChange={(e) => set(s.key, e.target.value)}
                                placeholder="请选择"
                                options={s.options ?? []}
                            />
                        ) : (
                            <NativeInput
                                value={form[s.key] ?? ''}
                                onChange={(e) => set(s.key, e.target.value)}
                                placeholder={
                                    s.unit ? `单位: ${s.unit}` : s.label
                                }
                            />
                        )}
                    </FormRow>
                ))}

                <FormRow label="投用时间">
                    <DatePicker
                        value={form.use_time ?? ''}
                        onChange={(v) => set('use_time', v)}
                    />
                </FormRow>

                <FormRow label="检验有效期">
                    <DatePicker
                        value={form.valid_period ?? ''}
                        onChange={(v) => set('valid_period', v)}
                    />
                </FormRow>

                <FormRow label="检验单位">
                    <NativeInput
                        value={form.inspect_unit ?? ''}
                        onChange={(e) => set('inspect_unit', e.target.value)}
                        placeholder="检验机构名称"
                    />
                </FormRow>

                <FormRow label="维护信息">
                    <NativeInput
                        value={form.maintain_info ?? ''}
                        onChange={(e) => set('maintain_info', e.target.value)}
                        placeholder="维护情况简述"
                    />
                </FormRow>

                <FormRow label="备注" colSpan={2}>
                    <NativeTextarea
                        value={form.comment ?? ''}
                        onChange={(e) => set('comment', e.target.value)}
                        rows={3}
                    />
                </FormRow>
            </div>
        </FormDialog>
    );
}

// ── Delete confirm ────────────────────────────────────────────────

function DeleteConfirm({
    equipment,
    onClose,
    onDeleted,
}: {
    equipment: SpecialEquipment;
    onClose: () => void;
    onDeleted: () => void;
}) {
    const [busy, setBusy] = useState(false);
    const handle = async () => {
        setBusy(true);
        try {
            await api.delete(API.specialEquipmentById(equipment.id));
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
            title={`删除「${equipment.equipment_name || `#${equipment.id}`}」`}
            icon={Trash2}
            size="md"
            onConfirm={handle}
            confirmLabel="确认删除"
            confirmTone="rose"
            loading={busy}
        >
            <p className="text-sm leading-relaxed text-slate-400">
                软删除：记录被标记为已删除，不会真正从数据库移除。如有必要可联系运维恢复。
            </p>
        </FormDialog>
    );
}

// ── Helpers ───────────────────────────────────────────────────────

function triggerDownload(blob: Blob, filename: string) {
    const url = URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = filename;
    document.body.appendChild(a);
    a.click();
    a.remove();
    setTimeout(() => URL.revokeObjectURL(url), 5000);
}
