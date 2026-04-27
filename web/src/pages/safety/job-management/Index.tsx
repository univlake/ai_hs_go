import { Head } from '@inertiajs/react';
import type { ColumnDef } from '@tanstack/react-table';
import { Plus, Pencil, Trash2, X, AlertTriangle, Loader2 } from 'lucide-react';
import { useState, useEffect, useCallback } from 'react';
import { DataTable } from '@/components/data-table/DataTable';
import { DatePicker } from '@/components/shared/DatePicker';
import { DetailButton } from '@/components/shared/DetailButton';
import { ObjectDetailDrawer } from '@/components/shared/ObjectDetailDrawer';
import { useScope } from '@/contexts/ScopeContext';
import { useApiList, useApiCrud } from '@/hooks/use-api';
import { useDetailDrawer } from '@/hooks/use-detail-drawer';
import { usePermission } from '@/hooks/use-permission';
import api from '@/lib/api';
import { cn } from '@/lib/utils';
import type { ApiResponse } from '@/types/models';

// ── Types ─────────────────────────────────────────────────────────

interface JobManagement {
    id: number;
    department: string | null;
    content: string | null;
    type: string | null;
    special_type: string | null;
    position: string | null;
    job_time: string | null;
    unit: string | null;
    is_outsourcing: number | null;
    evaluation_unit_id: number | null;
    job_people_number: number | null;
    job_people_name: string | null;
    risk_analysis: string | null;
    risk_protection: string | null;
    company_id: number;
    evaluation_unit?: { id: number; name: string } | null;
}

interface FormData {
    department: string;
    content: string;
    type: string;
    special_type: string;
    position: string;
    job_time: string;
    unit: string;
    is_outsourcing: '' | '1' | '2';
    evaluation_unit_id: string;
    job_people_number: string;
    job_people_name: string;
    risk_analysis: string;
    risk_protection: string;
}

const EMPTY_FORM: FormData = {
    department: '',
    content: '',
    type: '',
    special_type: '',
    position: '',
    job_time: '',
    unit: '',
    is_outsourcing: '',
    evaluation_unit_id: '',
    job_people_number: '',
    job_people_name: '',
    risk_analysis: '',
    risk_protection: '',
};

const JOB_TYPES = ['一般作业', '八大特殊作业'];

const OUTSOURCING_OPTIONS: { value: '1' | '2'; label: string }[] = [
    { value: '1', label: '是' },
    { value: '2', label: '否' },
];

// ── Page ──────────────────────────────────────────────────────────

export default function JobManagementIndex() {
    const { activeCompany } = useScope();
    const { canWrite } = usePermission();
    const canCreate = canWrite('job-management.create');
    const canUpdate = canWrite('job-management.update');
    const canDelete = canWrite('job-management.delete');
    const companyId = activeCompany?.id;

    const {
        data,
        loading,
        total,
        current_page,
        last_page,
        per_page,
        reload,
        goToPage,
        search,
    } = useApiList<JobManagement>({
        endpoint: '/job-management',
        defaultParams: companyId ? { company_id: companyId } : {},
        enabled: !!companyId,
    });
    const crud = useApiCrud('/job-management');

    const [dialogOpen, setDialogOpen] = useState(false);
    const [editingId, setEditingId] = useState<number | null>(null);
    const detail = useDetailDrawer<JobManagement>();
    const detailItem = detail.data;
    const [deleteId, setDeleteId] = useState<number | null>(null);
    const [deleteLabel, setDeleteLabel] = useState('');
    const [formLoading, setFormLoading] = useState(false);
    const [formError, setFormError] = useState('');
    const [searchValue, setSearchValue] = useState('');
    const [filterType, setFilterType] = useState('');
    const [filterOutsourcing, setFilterOutsourcing] = useState('');

    const [form, setForm] = useState<FormData>(EMPTY_FORM);
    const [evaluationUnits, setEvaluationUnits] = useState<
        { id: number; name: string }[]
    >([]);

    useEffect(() => {
        if (!companyId) {
            return;
        }

        api.get<ApiResponse<{ data: { id: number; name: string }[] }>>(
            '/evaluation-units',
            { params: { company_id: companyId, per_page: 200 } },
        )
            .then((r) => setEvaluationUnits(r.data.data?.data ?? []))
            .catch(() => {});
    }, [companyId]);

    const update = (key: keyof FormData, value: string) =>
        setForm((f) => ({ ...f, [key]: value }));

    const openCreate = () => {
        setFormError('');
        setForm(EMPTY_FORM);
        setEditingId(null);
        setDialogOpen(true);
    };

    const openEdit = (row: JobManagement) => {
        setFormError('');
        setForm({
            department: row.department ?? '',
            content: row.content ?? '',
            type: row.type ?? '',
            special_type: row.special_type ?? '',
            position: row.position ?? '',
            job_time: row.job_time ?? '',
            unit: row.unit ?? '',
            is_outsourcing:
                row.is_outsourcing === 1
                    ? '1'
                    : row.is_outsourcing === 2
                      ? '2'
                      : '',
            evaluation_unit_id: row.evaluation_unit_id
                ? String(row.evaluation_unit_id)
                : '',
            job_people_number:
                row.job_people_number != null
                    ? String(row.job_people_number)
                    : '',
            job_people_name: row.job_people_name ?? '',
            risk_analysis: row.risk_analysis ?? '',
            risk_protection: row.risk_protection ?? '',
        });
        setEditingId(row.id);
        setDialogOpen(true);
    };

    const handleSubmit = async () => {
        setFormLoading(true);
        setFormError('');

        try {
            const payload: Record<string, unknown> = {
                department: form.department || null,
                content: form.content || null,
                type: form.type || null,
                special_type: form.special_type || null,
                position: form.position || null,
                job_time: form.job_time || null,
                unit: form.unit || null,
                is_outsourcing: form.is_outsourcing
                    ? Number(form.is_outsourcing)
                    : null,
                evaluation_unit_id: form.evaluation_unit_id
                    ? Number(form.evaluation_unit_id)
                    : null,
                job_people_number: form.job_people_number
                    ? Number(form.job_people_number)
                    : null,
                job_people_name: form.job_people_name || null,
                risk_analysis: form.risk_analysis || null,
                risk_protection: form.risk_protection || null,
                company_id: companyId,
            };

            if (editingId) {
                await crud.update(editingId, payload);
            } else {
                await crud.create(payload);
            }

            setDialogOpen(false);
            reload();
        } catch (err: unknown) {
            const e = err as { response?: { data?: { message?: string } } };
            setFormError(e.response?.data?.message ?? '操作失败，请重试');
        } finally {
            setFormLoading(false);
        }
    };

    const handleDelete = async () => {
        if (!deleteId) {
            return;
        }

        await crud.remove(deleteId);
        setDeleteId(null);
        setDeleteLabel('');
        reload();
    };

    const applyFilters = useCallback(
        (next: { type?: string; out?: string; kw?: string }) => {
            const params: Record<string, unknown> = {};
            const t = next.type ?? filterType;
            const o = next.out ?? filterOutsourcing;
            const kw = next.kw ?? searchValue;

            if (t) {
                params.type = t;
            }

            if (o) {
                params.is_outsourcing = o;
            }

            if (kw) {
                params.keyword = kw;
            }

            search(params);
        },
        [filterType, filterOutsourcing, searchValue, search],
    );

    const handleSearch = useCallback(
        (v: string) => {
            setSearchValue(v);
            applyFilters({ kw: v });
        },
        [applyFilters],
    );

    // ── Table columns ──────────────────────────────────────────────

    const columns: ColumnDef<JobManagement>[] = [
        { accessorKey: 'id', header: 'ID', size: 60 },
        { accessorKey: 'department', header: '发起部门' },
        {
            accessorKey: 'type',
            header: '作业类型',
            size: 120,
            cell: ({ row }) =>
                row.original.type ? (
                    <span className="text-slate-300">{row.original.type}</span>
                ) : (
                    <span className="text-slate-700">—</span>
                ),
        },
        {
            accessorKey: 'special_type',
            header: '特殊作业类型',
            cell: ({ row }) =>
                row.original.special_type ? (
                    <span className="text-slate-300">
                        {row.original.special_type}
                    </span>
                ) : (
                    <span className="text-slate-700">—</span>
                ),
        },
        { accessorKey: 'position', header: '位置', size: 140 },
        {
            accessorKey: 'job_time',
            header: '作业时间',
            size: 110,
            cell: ({ row }) =>
                row.original.job_time ? (
                    String(row.original.job_time).slice(0, 10)
                ) : (
                    <span className="text-slate-700">—</span>
                ),
        },
        {
            accessorKey: 'is_outsourcing',
            header: '是否外协',
            size: 90,
            cell: ({ row }) => {
                const v = row.original.is_outsourcing;
                return v === 1 ? (
                    <span className="inline-flex items-center rounded-full bg-amber-600/15 px-2 py-0.5 text-xs font-medium text-amber-300">
                        外协
                    </span>
                ) : v === 2 ? (
                    <span className="inline-flex items-center rounded-full bg-blue-600/15 px-2 py-0.5 text-xs font-medium text-blue-300">
                        自营
                    </span>
                ) : (
                    <span className="text-slate-700">—</span>
                );
            },
        },
        {
            id: 'evaluation_unit',
            header: '外协单位',
            cell: ({ row }) =>
                row.original.evaluation_unit?.name ? (
                    <span className="text-slate-300">
                        {row.original.evaluation_unit.name}
                    </span>
                ) : (
                    <span className="text-slate-700">—</span>
                ),
        },
        {
            id: 'actions',
            header: '操作',
            size: 160,
            cell: ({ row }) => (
                <div className="flex items-center gap-0.5">
                    <DetailButton
                        row={row.original}
                        onOpen={detail.openDetail}
                    />
                    {canUpdate && (
                        <ActionBtn
                            onClick={() => openEdit(row.original)}
                            title="编辑"
                            variant="edit"
                        >
                            <Pencil className="h-3.5 w-3.5" />
                        </ActionBtn>
                    )}
                    {canDelete && (
                        <ActionBtn
                            onClick={() => {
                                setDeleteId(row.original.id);
                                setDeleteLabel(row.original.content ?? '');
                            }}
                            title="删除"
                            variant="delete"
                        >
                            <Trash2 className="h-3.5 w-3.5" />
                        </ActionBtn>
                    )}
                </div>
            ),
        },
    ];

    const detailData = detailItem
        ? {
              ...detailItem,
              evaluation_unit: detailItem.evaluation_unit?.name ?? null,
              is_outsourcing:
                  detailItem.is_outsourcing === 1
                      ? '是'
                      : detailItem.is_outsourcing === 2
                        ? '否'
                        : null,
          }
        : null;

    return (
        <>
            <Head title="作业管理" />

            <div className="mb-6 flex items-center justify-between">
                <div>
                    <h1 className="text-xl font-bold text-slate-100">
                        作业管理
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
                    <button onClick={openCreate} className={primaryBtn}>
                        <Plus className="h-4 w-4" />
                        新增作业
                    </button>
                )}
            </div>

            <div className="mb-3 flex items-center gap-2">
                <select
                    value={filterType}
                    onChange={(e) => {
                        setFilterType(e.target.value);
                        applyFilters({ type: e.target.value });
                    }}
                    className={cn(inputClass, 'w-40')}
                >
                    <option value="">作业类型（全部）</option>
                    {JOB_TYPES.map((t) => (
                        <option key={t} value={t}>
                            {t}
                        </option>
                    ))}
                </select>
                <select
                    value={filterOutsourcing}
                    onChange={(e) => {
                        setFilterOutsourcing(e.target.value);
                        applyFilters({ out: e.target.value });
                    }}
                    className={cn(inputClass, 'w-32')}
                >
                    <option value="">外协（全部）</option>
                    {OUTSOURCING_OPTIONS.map((o) => (
                        <option key={o.value} value={o.value}>
                            {o.label}
                        </option>
                    ))}
                </select>
            </div>

            <DataTable
                columns={columns}
                data={data}
                isLoading={loading}
                searchPlaceholder="搜索作业内容..."
                searchValue={searchValue}
                onSearchChange={handleSearch}
                pagination={{ current_page, last_page, per_page, total }}
                onPageChange={goToPage}
            />

            <ObjectDetailDrawer
                open={detail.open}
                title="作业详情"
                data={detailData as Record<string, unknown> | null}
                onClose={detail.close}
                labels={{
                    department: '发起部门',
                    content: '作业内容',
                    type: '作业类型',
                    special_type: '特殊作业类型',
                    position: '位置',
                    job_time: '作业时间',
                    unit: '作业单位',
                    is_outsourcing: '是否外协',
                    evaluation_unit: '外协单位',
                    job_people_number: '作业人数',
                    job_people_name: '作业负责人',
                    risk_analysis: '安全风险分析',
                    risk_protection: '安全保障措施',
                }}
            />

            {/* Create / Edit modal */}
            <Modal open={dialogOpen} onClose={() => setDialogOpen(false)}>
                <ModalHeader
                    title={editingId ? '编辑作业' : '新增作业'}
                    onClose={() => setDialogOpen(false)}
                />
                <div className="flex-1 space-y-4 overflow-y-auto p-6">
                    {formError && <ErrorAlert message={formError} />}

                    <div className="grid grid-cols-2 gap-4">
                        <Field label="发起部门">
                            <input
                                className={inputClass}
                                value={form.department}
                                onChange={(e) =>
                                    update('department', e.target.value)
                                }
                                placeholder="请输入发起部门"
                            />
                        </Field>
                        <Field label="作业类型">
                            <select
                                className={inputClass}
                                value={form.type}
                                onChange={(e) =>
                                    update('type', e.target.value)
                                }
                            >
                                <option value="">请选择</option>
                                {JOB_TYPES.map((t) => (
                                    <option key={t} value={t}>
                                        {t}
                                    </option>
                                ))}
                            </select>
                        </Field>
                        <Field label="特殊作业类型">
                            <input
                                className={inputClass}
                                value={form.special_type}
                                onChange={(e) =>
                                    update('special_type', e.target.value)
                                }
                                placeholder={
                                    form.type === '八大特殊作业'
                                        ? '如：动火作业 / 受限空间作业 / 高处作业'
                                        : '请输入特殊作业类型'
                                }
                            />
                        </Field>
                        <Field label="位置">
                            <input
                                className={inputClass}
                                value={form.position}
                                onChange={(e) =>
                                    update('position', e.target.value)
                                }
                                placeholder="请输入位置"
                            />
                        </Field>
                        <Field label="作业时间">
                            <DatePicker
                                value={form.job_time}
                                onChange={(v) => update('job_time', v)}
                            />
                        </Field>
                        <Field label="作业单位">
                            <input
                                className={inputClass}
                                value={form.unit}
                                onChange={(e) =>
                                    update('unit', e.target.value)
                                }
                                placeholder="请输入作业单位"
                            />
                        </Field>
                        <Field label="是否外协单位">
                            <select
                                className={inputClass}
                                value={form.is_outsourcing}
                                onChange={(e) =>
                                    update('is_outsourcing', e.target.value)
                                }
                            >
                                <option value="">请选择</option>
                                {OUTSOURCING_OPTIONS.map((o) => (
                                    <option key={o.value} value={o.value}>
                                        {o.label}
                                    </option>
                                ))}
                            </select>
                        </Field>
                        <Field label="外协单位">
                            <select
                                className={inputClass}
                                value={form.evaluation_unit_id}
                                onChange={(e) =>
                                    update(
                                        'evaluation_unit_id',
                                        e.target.value,
                                    )
                                }
                                disabled={form.is_outsourcing !== '1'}
                            >
                                <option value="">
                                    {form.is_outsourcing === '1'
                                        ? '请选择外协单位'
                                        : '仅在外协时可选'}
                                </option>
                                {evaluationUnits.map((u) => (
                                    <option key={u.id} value={String(u.id)}>
                                        {u.name}
                                    </option>
                                ))}
                            </select>
                        </Field>
                        <Field label="作业人数">
                            <input
                                type="number"
                                min={0}
                                className={inputClass}
                                value={form.job_people_number}
                                onChange={(e) =>
                                    update(
                                        'job_people_number',
                                        e.target.value,
                                    )
                                }
                                placeholder="请输入作业人数"
                            />
                        </Field>
                        <Field label="作业负责人">
                            <input
                                className={inputClass}
                                value={form.job_people_name}
                                onChange={(e) =>
                                    update('job_people_name', e.target.value)
                                }
                                placeholder="请输入负责人姓名"
                            />
                        </Field>
                    </div>

                    <Field label="作业内容">
                        <textarea
                            className={cn(inputClass, 'h-20 resize-y')}
                            value={form.content}
                            onChange={(e) =>
                                update('content', e.target.value)
                            }
                            placeholder="请描述作业内容"
                        />
                    </Field>

                    <Field label="作业安全风险分析">
                        <textarea
                            className={cn(inputClass, 'h-20 resize-y')}
                            value={form.risk_analysis}
                            onChange={(e) =>
                                update('risk_analysis', e.target.value)
                            }
                            placeholder="请描述作业安全风险"
                        />
                    </Field>

                    <Field label="作业安全保障措施">
                        <textarea
                            className={cn(inputClass, 'h-20 resize-y')}
                            value={form.risk_protection}
                            onChange={(e) =>
                                update('risk_protection', e.target.value)
                            }
                            placeholder="请描述安全保障措施"
                        />
                    </Field>
                </div>
                <ModalFooter
                    onCancel={() => setDialogOpen(false)}
                    onConfirm={handleSubmit}
                    loading={formLoading}
                />
            </Modal>

            {/* Delete confirm */}
            <Modal
                open={!!deleteId}
                onClose={() => {
                    setDeleteId(null);
                    setDeleteLabel('');
                }}
            >
                <div className="p-6">
                    <div className="mb-4 flex h-11 w-11 items-center justify-center rounded-xl border border-red-500/20 bg-red-600/10">
                        <Trash2 className="h-5 w-5 text-red-400" />
                    </div>
                    <h2 className="mb-1 text-base font-semibold text-slate-100">
                        删除作业
                    </h2>
                    <p className="text-sm text-slate-500">
                        确定要删除此作业记录吗？此操作不可撤销。
                        {deleteLabel && (
                            <span className="mt-1 block truncate text-xs text-slate-600">
                                内容：{deleteLabel}
                            </span>
                        )}
                    </p>
                </div>
                <div className="flex items-center justify-end gap-2 border-t border-blue-900/30 px-6 py-4">
                    <button
                        onClick={() => {
                            setDeleteId(null);
                            setDeleteLabel('');
                        }}
                        className={ghostBtn}
                    >
                        取消
                    </button>
                    <button
                        onClick={handleDelete}
                        disabled={crud.loading}
                        className={dangerBtn}
                    >
                        {crud.loading && (
                            <Loader2 className="h-3.5 w-3.5 animate-spin" />
                        )}
                        {crud.loading ? '删除中...' : '确认删除'}
                    </button>
                </div>
            </Modal>
        </>
    );
}

// ── Shared primitives ─────────────────────────────────────────────

const inputClass =
    'w-full rounded-lg border border-blue-900/40 bg-[#070d1f] px-3 py-2 text-sm text-slate-200 placeholder-slate-700 outline-none transition-colors focus:border-blue-500/60 focus:ring-1 focus:ring-blue-500/20 appearance-none disabled:opacity-40 disabled:cursor-not-allowed';
const primaryBtn =
    'flex items-center gap-2 rounded-lg border border-blue-500/30 bg-blue-600/15 px-4 py-2 text-sm font-medium text-blue-300 transition-all hover:border-blue-400/50 hover:bg-blue-600/25 hover:text-blue-200 active:scale-95';
const ghostBtn =
    'rounded-lg border border-blue-900/40 px-4 py-2 text-sm text-slate-400 transition-colors hover:border-blue-500/40 hover:text-slate-200';
const dangerBtn =
    'flex items-center gap-2 rounded-lg border border-red-500/30 bg-red-600/15 px-4 py-2 text-sm font-medium text-red-400 transition-all hover:bg-red-600/25 hover:text-red-300 disabled:pointer-events-none disabled:opacity-50';

function Field({
    label,
    required,
    children,
}: {
    label: string;
    required?: boolean;
    children: React.ReactNode;
}) {
    return (
        <div>
            <label className="mb-1.5 block text-xs font-medium tracking-wide text-slate-500 uppercase">
                {label}
                {required && <span className="ml-1 text-blue-500">*</span>}
            </label>
            {children}
        </div>
    );
}

function ErrorAlert({ message }: { message: string }) {
    return (
        <div className="flex items-center gap-2 rounded-lg border border-red-500/20 bg-red-600/8 px-3 py-2.5 text-xs text-red-400">
            <AlertTriangle className="h-3.5 w-3.5 shrink-0" />
            {message}
        </div>
    );
}

function Modal({
    open,
    onClose,
    children,
}: {
    open: boolean;
    onClose: () => void;
    children: React.ReactNode;
}) {
    if (!open) {
        return null;
    }

    return (
        <div className="fixed inset-0 z-50 flex items-center justify-center p-4">
            <div
                className="absolute inset-0 bg-black/60 backdrop-blur-sm"
                onClick={onClose}
            />
            <div className="relative flex max-h-[92vh] w-full max-w-2xl flex-col overflow-hidden rounded-2xl border border-blue-900/50 bg-[#0d1530] shadow-2xl shadow-black/60">
                {children}
            </div>
        </div>
    );
}

function ModalHeader({
    title,
    onClose,
}: {
    title: string;
    onClose: () => void;
}) {
    return (
        <div className="flex items-center justify-between border-b border-blue-900/30 px-6 py-4">
            <h2 className="text-base font-semibold text-slate-100">{title}</h2>
            <button
                onClick={onClose}
                className="flex h-7 w-7 items-center justify-center rounded-lg text-slate-600 transition-colors hover:bg-white/5 hover:text-slate-300"
            >
                <X className="h-4 w-4" />
            </button>
        </div>
    );
}

function ModalFooter({
    onCancel,
    onConfirm,
    loading,
}: {
    onCancel: () => void;
    onConfirm: () => void;
    loading: boolean;
}) {
    return (
        <div className="flex items-center justify-end gap-2 border-t border-blue-900/30 px-6 py-4">
            <button onClick={onCancel} className={ghostBtn}>
                取消
            </button>
            <button
                onClick={onConfirm}
                disabled={loading}
                className="flex items-center gap-2 rounded-lg border border-blue-500/30 bg-blue-600/20 px-4 py-2 text-sm font-medium text-blue-300 transition-all hover:bg-blue-600/30 hover:text-blue-200 disabled:pointer-events-none disabled:opacity-50"
            >
                {loading && <Loader2 className="h-3.5 w-3.5 animate-spin" />}
                {loading ? '保存中...' : '保存'}
            </button>
        </div>
    );
}

function ActionBtn({
    children,
    onClick,
    title,
    variant,
}: {
    children: React.ReactNode;
    onClick: () => void;
    title: string;
    variant: 'edit' | 'delete';
}) {
    return (
        <button
            onClick={onClick}
            title={title}
            className={cn(
                'flex h-7 w-7 items-center justify-center rounded-md transition-colors',
                variant === 'edit'
                    ? 'text-slate-600 hover:bg-blue-600/10 hover:text-blue-400'
                    : 'text-slate-600 hover:bg-red-600/10 hover:text-red-400',
            )}
        >
            {children}
        </button>
    );
}
