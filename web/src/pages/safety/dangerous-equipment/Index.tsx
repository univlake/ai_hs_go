import { useCallback, useEffect, useMemo, useRef, useState } from 'react';
import {
    AlertTriangle,
    Download,
    Edit3,
    Eye,
    FileSpreadsheet,
    Plus,
    RefreshCw,
    Search,
    Trash2,
    Upload,
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

    const stats = useMemo(
        () => ({
            total,
            byAutomation: rows.reduce(
                (m: Record<string, number>, r) => {
                    const k = r.automation_type || '未填';
                    m[k] = (m[k] || 0) + 1;
                    return m;
                },
                {},
            ),
        }),
        [rows, total],
    );

    return (
        <div className="space-y-6">
            <div className="flex items-start justify-between">
                <div>
                    <h1 className="text-2xl font-bold text-slate-100">
                        重大危险设备
                    </h1>
                    <p className="mt-1 text-sm text-slate-500">
                        当前租户的重大危险设备台账 · 含使用证 / 自动化程度 /
                        维护情况
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
                <div className="rounded-xl border border-blue-900/40 bg-[#0d1530] p-5">
                    <div className="flex items-start justify-between">
                        <div>
                            <p className="text-xs font-medium tracking-wide text-slate-500 uppercase">
                                设备总数
                            </p>
                            <p className="mt-1 text-2xl font-bold text-slate-100">
                                {stats.total}
                            </p>
                        </div>
                        <div className="flex h-9 w-9 items-center justify-center rounded-lg border border-amber-600/20 bg-amber-600/15">
                            <AlertTriangle className="h-4 w-4 text-amber-400" />
                        </div>
                    </div>
                </div>
                <div className="rounded-xl border border-blue-900/40 bg-[#0d1530] p-5 md:col-span-2">
                    <p className="text-xs font-medium tracking-wide text-slate-500 uppercase">
                        自动化程度分布（本页）
                    </p>
                    <div className="mt-2 flex flex-wrap gap-1.5">
                        {Object.entries(stats.byAutomation).map(([k, v]) => (
                            <span
                                key={k}
                                className="rounded-md border border-blue-500/30 bg-blue-600/15 px-2 py-1 text-[11px] text-blue-300"
                            >
                                {k}
                                <span className="ml-1.5 text-blue-400/60">
                                    {v}
                                </span>
                            </span>
                        ))}
                        {Object.keys(stats.byAutomation).length === 0 && (
                            <span className="text-[11px] text-slate-600">
                                暂无数据
                            </span>
                        )}
                    </div>
                </div>
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
                        <AlertTriangle className="mx-auto mb-3 h-10 w-10 text-slate-700" />
                        <p className="text-sm text-slate-400">暂无设备</p>
                    </div>
                ) : (
                    <table className="w-full text-sm">
                        <thead>
                            <tr className="border-b border-blue-900/30 bg-[#0a1428] text-[10px] tracking-widest text-slate-600 uppercase">
                                <th className="px-5 py-3 text-left">设备</th>
                                <th className="px-5 py-3 text-left">使用位置</th>
                                <th className="px-5 py-3 text-left">规格型号</th>
                                <th className="px-5 py-3 text-left">投用时间</th>
                                <th className="px-5 py-3 text-left">自动化程度</th>
                                <th className="px-5 py-3 text-right">操作</th>
                            </tr>
                        </thead>
                        <tbody>
                            {rows.map((r) => (
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
                                    <td className="px-5 py-3 text-xs text-slate-400">
                                        {r.utilized_location || '—'}
                                    </td>
                                    <td className="px-5 py-3 text-xs text-slate-400">
                                        {r.norm_model || '—'}
                                    </td>
                                    <td className="px-5 py-3 text-xs text-slate-400">
                                        {fmtDate(r.use_time)}
                                    </td>
                                    <td className="px-5 py-3 text-xs">
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
