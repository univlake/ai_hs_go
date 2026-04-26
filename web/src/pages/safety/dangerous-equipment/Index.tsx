import { useCallback, useEffect, useRef, useState } from 'react';
import {
    AlertTriangle,
    Copy,
    Download,
    FileSpreadsheet,
    Loader2,
    Pencil,
    Plus,
    RefreshCw,
    Search,
    Trash2,
    Upload,
} from 'lucide-react';
import { toast } from 'sonner';
import api from '@/lib/api';
import { API } from '@/lib/api-routes';
import { DetailButton } from '@/components/shared/DetailButton';
import { FormDialog, FormRow } from '@/components/shared/FormDialog';
import { ObjectDetailDrawer } from '@/components/shared/ObjectDetailDrawer';
import { DatePicker } from '@/components/shared/DatePicker';
import {
    NativeInput,
    NativeSelect,
    NativeTextarea,
} from '@/components/shared/FormControls';
import { useAuth } from '@/hooks/use-auth';

// Mirrors Laravel project's AUTOMATION_OPTIONS for 危险设备 自动化程度
const AUTOMATION_OPTIONS = [
    '全自动设备',
    '需要作业人员辅助操作设备',
    '作业人员长时间操作设备',
];

const dangerousEquipmentLabels: Record<string, string> = {
    equipment_name: '设备名称',
    utilized_location: '使用位置',
    norm_model: '规格型号',
    serial_number: '使用证编号',
    use_time: '投用时间',
    automation_type: '自动化程度',
    maintain_info: '维护信息',
    comment: '备注',
    created_at: '创建时间',
    updated_at: '更新时间',
};
const detailHiddenKeys = [
    'id',
    'company_id',
    'user_id',
    'last_user_id',
    'deleted_at',
];

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

interface DangerousEquipment {
    id: number;
    equipment_name: string | null;
    utilized_location: string | null;
    norm_model: string | null;
    serial_number: string | null;
    use_time: string | null;
    automation_type: string | null;
    maintain_info: string | null;
    comment: string | null;
    company_id: number;
    created_at: string;
    updated_at: string;
}

type FormState = Omit<
    DangerousEquipment,
    'id' | 'company_id' | 'created_at' | 'updated_at'
>;

const emptyForm: FormState = {
    equipment_name: '',
    utilized_location: '',
    norm_model: '',
    serial_number: '',
    use_time: '',
    automation_type: '',
    maintain_info: '',
    comment: '',
};

function fmtDate(d?: string | null) {
    if (!d) return '—';
    return d.slice(0, 10);
}

export default function SafetyDangerousEquipment() {
    const { hasPermission } = useAuth();
    const [rows, setRows] = useState<DangerousEquipment[]>([]);
    const [loading, setLoading] = useState(true);
    const [page, setPage] = useState(1);
    const [perPage] = useState(20);
    const [total, setTotal] = useState(0);
    const [keyword, setKeyword] = useState('');

    const [creating, setCreating] = useState(false);
    const [editing, setEditing] = useState<DangerousEquipment | null>(null);
    const [viewing, setViewing] = useState<DangerousEquipment | null>(null);
    const [deleting, setDeleting] = useState<DangerousEquipment | null>(null);
    const [selectedIds, setSelectedIds] = useState<Set<number>>(new Set());
    const [copyLoading, setCopyLoading] = useState(false);
    const fileRef = useRef<HTMLInputElement>(null);

    const canCreate = hasPermission('dangerous-equipment.create');
    const canUpdate = hasPermission('dangerous-equipment.update');
    const canDelete = hasPermission('dangerous-equipment.delete');

    const refresh = useCallback(async () => {
        setLoading(true);
        try {
            const res = await api.get(API.dangerousEquipment(), {
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
            toast.error('加载危险设备失败');
            // eslint-disable-next-line no-console
            console.error(err);
        } finally {
            setLoading(false);
        }
    }, [page, perPage, keyword]);

    useEffect(() => {
        refresh();
    }, [refresh]);

    const downloadTemplate = async () => {
        try {
            const res = await api.get(API.dangerousEquipmentTemplate(), {
                responseType: 'blob',
            });
            triggerDownload(res.data, '重大危险设备导入模板.xlsx');
        } catch {
            toast.error('下载模板失败');
        }
    };
    const downloadExport = async () => {
        try {
            const res = await api.get(API.dangerousEquipmentExport(), {
                responseType: 'blob',
            });
            triggerDownload(res.data, '重大危险设备台账.xlsx');
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
            const res = await api.post(API.dangerousEquipmentImport(), fd, {
                headers: { 'Content-Type': 'multipart/form-data' },
            });
            const d = res.data?.data ?? {};
            toast.success(`已导入 ${d.imported ?? 0} / ${d.total ?? 0} 条`);
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

    // ── Multi-select copy（与 fire-equipment / CrudPage 一致）──────
    const allPageSelected =
        rows.length > 0 && rows.every((r) => selectedIds.has(r.id));
    const somePageSelected = rows.some((r) => selectedIds.has(r.id));
    const toggleRow = (id: number) =>
        setSelectedIds((prev) => {
            const next = new Set(prev);
            if (next.has(id)) {
                next.delete(id);
            } else {
                next.add(id);
            }
            return next;
        });
    const toggleAll = () =>
        setSelectedIds((prev) => {
            const next = new Set(prev);
            if (allPageSelected) {
                rows.forEach((r) => next.delete(r.id));
            } else {
                rows.forEach((r) => next.add(r.id));
            }
            return next;
        });

    const handleCopy = async () => {
        if (!selectedIds.size) return;
        setCopyLoading(true);
        try {
            await api.post('/dangerous-equipment/copy', {
                ids: Array.from(selectedIds),
            });
            toast.success(`已复制 ${selectedIds.size} 条`);
            setSelectedIds(new Set());
            refresh();
        } catch (err) {
            const msg =
                (err as { response?: { data?: { message?: string } } }).response
                    ?.data?.message ?? '复制失败';
            toast.error(msg);
        } finally {
            setCopyLoading(false);
        }
    };

    return (
        <div>
            <div className="mb-6 flex items-center justify-between">
                <div>
                    <h1 className="text-xl font-bold text-slate-100">
                        重大危险设备
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
                        新增重大危险设备
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
                        placeholder="搜索设备名 / 使用证编号"
                        className="h-9 w-full rounded-lg border border-blue-900/40 bg-[#070d1f] pr-3 pl-9 text-sm text-slate-300 placeholder-slate-700 transition-colors outline-none focus:border-blue-500/60 focus:ring-1 focus:ring-blue-500/20"
                    />
                </div>
                {canCreate && selectedIds.size > 0 && (
                    <button
                        onClick={handleCopy}
                        disabled={copyLoading}
                        className="flex items-center gap-1.5 rounded-lg border border-emerald-500/30 bg-emerald-600/10 px-3 py-1.5 text-xs font-medium text-emerald-400 hover:bg-emerald-600/20 disabled:pointer-events-none disabled:opacity-50"
                    >
                        {copyLoading ? (
                            <Loader2 className="h-3.5 w-3.5 animate-spin" />
                        ) : (
                            <Copy className="h-3.5 w-3.5" />
                        )}
                        复制已选 ({selectedIds.size})
                    </button>
                )}
                <button
                    onClick={downloadTemplate}
                    className="flex items-center gap-1.5 rounded-lg border border-blue-900/40 bg-[#070d1f] px-3 py-1.5 text-xs text-slate-400 hover:border-blue-500/40 hover:text-slate-200"
                >
                    <FileSpreadsheet className="h-3.5 w-3.5" />
                    下载模板
                </button>
                {canCreate && (
                    <>
                        <input
                            ref={fileRef}
                            type="file"
                            accept=".xlsx,.xls"
                            hidden
                            onChange={onImportSelected}
                        />
                        <button
                            onClick={triggerImport}
                            className="flex items-center gap-1.5 rounded-lg border border-blue-900/40 bg-[#070d1f] px-3 py-1.5 text-xs text-slate-400 hover:border-blue-500/40 hover:text-slate-200"
                        >
                            <Upload className="h-3.5 w-3.5" />
                            导入
                        </button>
                    </>
                )}
                <button
                    onClick={downloadExport}
                    className="flex items-center gap-1.5 rounded-lg border border-blue-900/40 bg-[#070d1f] px-3 py-1.5 text-xs text-slate-400 hover:border-blue-500/40 hover:text-slate-200"
                >
                    <Download className="h-3.5 w-3.5" />
                    导出
                </button>
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
                                <AlertTriangle className="h-5 w-5 text-slate-700" />
                            </div>
                            <p className="text-sm text-slate-700">暂无数据</p>
                        </div>
                    </div>
                ) : (
                    <div className="overflow-x-auto">
                        <table className="w-full text-sm">
                            <thead>
                                <tr className="border-b border-blue-900/40 bg-[#070d1f]/60">
                                    {canCreate && (
                                        <th className="h-10 w-10 px-4 align-middle first:pl-5">
                                            <input
                                                type="checkbox"
                                                checked={allPageSelected}
                                                ref={(el) => {
                                                    if (el) {
                                                        el.indeterminate =
                                                            somePageSelected &&
                                                            !allPageSelected;
                                                    }
                                                }}
                                                onChange={toggleAll}
                                                className="h-3.5 w-3.5 rounded border-blue-900/60 bg-[#070d1f] accent-blue-500"
                                            />
                                        </th>
                                    )}
                                    <th className="h-10 px-4 text-left align-middle text-[10px] font-semibold tracking-widest text-slate-600 uppercase first:pl-5">
                                        设备
                                    </th>
                                    <th className="h-10 px-4 text-left align-middle text-[10px] font-semibold tracking-widest text-slate-600 uppercase">
                                        使用位置
                                    </th>
                                    <th className="h-10 px-4 text-left align-middle text-[10px] font-semibold tracking-widest text-slate-600 uppercase">
                                        规格型号
                                    </th>
                                    <th className="h-10 px-4 text-left align-middle text-[10px] font-semibold tracking-widest text-slate-600 uppercase">
                                        投用时间
                                    </th>
                                    <th className="h-10 px-4 text-left align-middle text-[10px] font-semibold tracking-widest text-slate-600 uppercase">
                                        自动化程度
                                    </th>
                                    <th className="h-10 px-4 text-right align-middle text-[10px] font-semibold tracking-widest text-slate-600 uppercase last:pr-5">
                                        操作
                                    </th>
                                </tr>
                            </thead>
                            <tbody>
                                {rows.map((r, idx) => (
                                    <tr
                                        key={r.id}
                                        className={`group border-b border-blue-900/15 transition-colors last:border-0 hover:bg-blue-600/5 ${
                                            idx % 2 === 1
                                                ? 'bg-white/[0.015]'
                                                : ''
                                        }`}
                                    >
                                        {canCreate && (
                                            <td className="w-10 px-4 py-3 align-middle first:pl-5">
                                                <input
                                                    type="checkbox"
                                                    checked={selectedIds.has(
                                                        r.id,
                                                    )}
                                                    onChange={() =>
                                                        toggleRow(r.id)
                                                    }
                                                    className="h-3.5 w-3.5 rounded border-blue-900/60 bg-[#070d1f] accent-blue-500"
                                                />
                                            </td>
                                        )}
                                        <td className="px-4 py-3 align-middle text-slate-300 first:pl-5">
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
                                        <td className="px-4 py-3 align-middle text-xs text-slate-400">
                                            {r.utilized_location || '—'}
                                        </td>
                                        <td className="px-4 py-3 align-middle text-xs text-slate-400">
                                            {r.norm_model || '—'}
                                        </td>
                                        <td className="px-4 py-3 align-middle text-xs text-slate-400">
                                            {fmtDate(r.use_time)}
                                        </td>
                                        <td className="px-4 py-3 align-middle text-xs">
                                            {r.automation_type ? (
                                                <span className="rounded-md border border-violet-500/30 bg-violet-600/15 px-1.5 py-0.5 text-violet-300">
                                                    {r.automation_type}
                                                </span>
                                            ) : (
                                                <span className="text-slate-600">
                                                    —
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
                                                        onClick={() => setEditing(r)}
                                                        title="编辑"
                                                        className="flex h-7 w-7 items-center justify-center rounded-md text-slate-600 transition-colors hover:bg-blue-600/10 hover:text-blue-400"
                                                    >
                                                        <Pencil className="h-3.5 w-3.5" />
                                                    </button>
                                                )}
                                                {canDelete && (
                                                    <button
                                                        onClick={() => setDeleting(r)}
                                                        title="删除"
                                                        className="flex h-7 w-7 items-center justify-center rounded-md text-slate-600 transition-colors hover:bg-red-600/10 hover:text-red-400"
                                                    >
                                                        <Trash2 className="h-3.5 w-3.5" />
                                                    </button>
                                                )}
                                            </div>
                                        </td>
                                    </tr>
                                ))}
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
                title={viewing?.equipment_name || `重大危险设备 #${viewing?.id ?? ''}`}
                data={viewing as Record<string, unknown> | null}
                labels={dangerousEquipmentLabels}
                hiddenKeys={detailHiddenKeys}
            />
        </div>
    );
}

function Editor({
    item,
    onClose,
    onSaved,
}: {
    item?: DangerousEquipment;
    onClose: () => void;
    onSaved: () => void;
}) {
    const isEdit = !!item;
    const [form, setForm] = useState<FormState>(() =>
        item
            ? {
                  equipment_name: item.equipment_name ?? '',
                  utilized_location: item.utilized_location ?? '',
                  norm_model: item.norm_model ?? '',
                  serial_number: item.serial_number ?? '',
                  use_time: (item.use_time ?? '').slice(0, 10),
                  automation_type: item.automation_type ?? '',
                  maintain_info: item.maintain_info ?? '',
                  comment: item.comment ?? '',
              }
            : { ...emptyForm },
    );
    const [busy, setBusy] = useState(false);

    const set = <K extends keyof FormState>(k: K, v: FormState[K]) =>
        setForm((f) => ({ ...f, [k]: v }));

    const submit = async () => {
        if (!form.equipment_name?.toString().trim()) {
            toast.error('设备名称必填');
            return;
        }
        setBusy(true);
        try {
            const payload = {
                ...form,
                use_time: form.use_time || null,
            };
            if (isEdit) {
                await api.put(API.dangerousEquipmentById(item!.id), payload);
            } else {
                await api.post(API.dangerousEquipment(), payload);
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
                    ? `编辑「${item!.equipment_name || `#${item!.id}`}」`
                    : '新建重大危险设备'
            }
            description={
                isEdit ? `#${item!.id} · 修改设备信息` : '记录一台新的重大危险设备'
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

                <FormRow label="使用位置">
                    <NativeInput
                        value={form.utilized_location ?? ''}
                        onChange={(e) =>
                            set('utilized_location', e.target.value)
                        }
                        placeholder="设备使用位置"
                    />
                </FormRow>

                <FormRow label="规格型号">
                    <NativeInput
                        value={form.norm_model ?? ''}
                        onChange={(e) => set('norm_model', e.target.value)}
                        placeholder="规格型号"
                    />
                </FormRow>

                <FormRow
                    label="使用证编号"
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
                        placeholder="请输入使用证编号（每台设备唯一）"
                        className={
                            isEdit && !form.serial_number
                                ? 'border-amber-500/50 focus:border-amber-500/70'
                                : ''
                        }
                    />
                </FormRow>

                <FormRow label="投用时间">
                    <DatePicker
                        value={form.use_time ?? ''}
                        onChange={(v) => set('use_time', v)}
                    />
                </FormRow>

                <FormRow label="自动化程度">
                    <NativeSelect
                        value={form.automation_type ?? ''}
                        onChange={(e) =>
                            set('automation_type', e.target.value)
                        }
                        placeholder="请选择"
                        options={AUTOMATION_OPTIONS}
                    />
                </FormRow>

                <FormRow label="维护信息" colSpan={2}>
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

function DeleteConfirm({
    item,
    onClose,
    onDeleted,
}: {
    item: DangerousEquipment;
    onClose: () => void;
    onDeleted: () => void;
}) {
    const [busy, setBusy] = useState(false);
    const handle = async () => {
        setBusy(true);
        try {
            await api.delete(API.dangerousEquipmentById(item.id));
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
            title={`删除「${item.equipment_name || `#${item.id}`}」`}
            icon={Trash2}
            size="md"
            onConfirm={handle}
            confirmLabel="确认删除"
            confirmTone="rose"
            loading={busy}
        >
            <p className="text-sm leading-relaxed text-slate-400">
                软删除：记录被标记为已删除，不会真正从数据库移除。
            </p>
        </FormDialog>
    );
}
