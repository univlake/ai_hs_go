import { Loader2, Plus, Trash2 } from 'lucide-react';
import { useCallback, useEffect, useState } from 'react';
import { useScope } from '@/contexts/ScopeContext';
import { usePermission } from '@/hooks/use-permission';
import api from '@/lib/api';
import type { ApiResponse } from '@/types/models';
import { YearSelect } from '../components';
import { TemplateScoringDrawer } from '../drawers/TemplateScoringDrawer';
import { StartScoringModal } from '../modals/StartScoringModal';
import type { PointsRow } from '../types';

export function PointsTab() {
    const { activeCompany } = useScope();
    const pid = activeCompany?.id;
    const { canWrite } = usePermission();
    const canCreate = canWrite('add-points.create');
    const canUpdate = canWrite('add-points.update');
    const canDelete = canWrite('add-points.delete');
    const [year, setYear] = useState(new Date().getFullYear());
    const [rows, setRows] = useState<PointsRow[]>([]);
    const [loading, setLoading] = useState(false);
    const [startOpen, setStartOpen] = useState(false);
    const [deleting, setDeleting] = useState<PointsRow | null>(null);
    const [deleteLoading, setDeleteLoading] = useState(false);
    const [scoring, setScoring] = useState<{
        companyId: number;
        companyName: string;
        year: number;
    } | null>(null);

    const fetchRows = useCallback(async () => {
        if (!pid) {
            return;
        }

        setLoading(true);

        try {
            const res = await api.get<ApiResponse<PointsRow[]>>(
                '/add-points/submissions',
                { params: { pid, year } },
            );
            setRows(res.data.data ?? []);
        } finally {
            setLoading(false);
        }
    }, [pid, year]);

    useEffect(() => {
        fetchRows();
    }, [fetchRows]);

    const handleDelete = useCallback(async () => {
        if (!deleting) {
            return;
        }
        setDeleteLoading(true);
        try {
            await api.delete('/add-points/submissions', {
                params: {
                    company_id: deleting.company_id,
                    year: deleting.year,
                },
            });
            setDeleting(null);
            await fetchRows();
        } finally {
            setDeleteLoading(false);
        }
    }, [deleting, fetchRows]);

    return (
        <>
            <div className="flex items-center justify-between gap-2">
                <div className="flex items-center gap-3">
                    <div className="flex items-center gap-2">
                        <label className="text-xs text-slate-400">年度</label>
                        <YearSelect value={year} onChange={setYear} />
                    </div>
                    <p className="text-xs text-slate-500">
                        加分项最终分 = 累计加分
                    </p>
                </div>
                {canCreate && (
                    <button
                        type="button"
                        onClick={() => setStartOpen(true)}
                        className="inline-flex min-w-[120px] items-center justify-center gap-2 rounded-lg border border-blue-500/30 bg-blue-600/20 px-4 py-2 text-sm font-medium text-blue-200 transition-colors hover:bg-blue-600/30"
                    >
                        <Plus className="h-4 w-4" />
                        开始考核
                    </button>
                )}
            </div>

            <div className="overflow-hidden rounded-xl border border-blue-900/40 bg-[#070d1f]">
                <table className="min-w-full text-sm">
                    <thead className="bg-blue-950/40 text-left text-xs tracking-wide text-slate-400 uppercase">
                        <tr>
                            <th className="px-4 py-3">企业</th>
                            <th className="w-20 px-4 py-3">年度</th>
                            <th className="w-24 px-4 py-3">累计加分</th>
                            <th className="w-32 px-4 py-3">更新时间</th>
                            <th className="w-32 px-4 py-3 text-right">操作</th>
                        </tr>
                    </thead>
                    <tbody className="divide-y divide-blue-900/30">
                        {loading && (
                            <tr>
                                <td
                                    colSpan={5}
                                    className="px-4 py-6 text-center text-slate-500"
                                >
                                    <Loader2 className="mx-auto h-4 w-4 animate-spin" />
                                </td>
                            </tr>
                        )}
                        {!loading && rows.length === 0 && (
                            <tr>
                                <td
                                    colSpan={5}
                                    className="px-4 py-6 text-center text-slate-500"
                                >
                                    暂无加分项记录，点击"开始考核"为下级企业评分。
                                </td>
                            </tr>
                        )}
                        {rows.map((r) => (
                            <tr
                                key={`${r.company_id}-${r.year}`}
                                className="text-slate-200 hover:bg-blue-950/30"
                            >
                                <td className="px-4 py-3 font-medium">
                                    {r.company_name}
                                </td>
                                <td className="px-4 py-3 text-slate-400">
                                    {r.year}
                                </td>
                                <td className="px-4 py-3 text-emerald-300">
                                    {r.total_number}
                                </td>
                                <td className="px-4 py-3 text-slate-400">
                                    {r.updated_at
                                        ? String(r.updated_at).slice(0, 10)
                                        : '—'}
                                </td>
                                <td className="px-4 py-3 text-right">
                                    <div className="flex items-center justify-end gap-1">
                                        {canUpdate && (
                                            <button
                                                type="button"
                                                onClick={() =>
                                                    setScoring({
                                                        companyId: r.company_id,
                                                        companyName:
                                                            r.company_name,
                                                        year: r.year,
                                                    })
                                                }
                                                className="inline-flex items-center gap-1 rounded-lg border border-blue-500/30 bg-blue-600/10 px-3 py-1 text-xs text-blue-300 transition-colors hover:bg-blue-600/20"
                                            >
                                                修改评分
                                            </button>
                                        )}
                                        {canDelete && (
                                            <button
                                                type="button"
                                                onClick={() => setDeleting(r)}
                                                title="删除该年评分"
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

            {startOpen && (
                <StartScoringModal
                    title="开始加分项考核"
                    onClose={() => setStartOpen(false)}
                    onStart={(y, id, name) => {
                        setStartOpen(false);
                        setScoring({
                            companyId: id,
                            companyName: name,
                            year: y,
                        });
                    }}
                />
            )}

            {scoring && (
                <TemplateScoringDrawer
                    kind="points"
                    companyId={scoring.companyId}
                    companyName={scoring.companyName}
                    year={scoring.year}
                    onClose={() => setScoring(null)}
                    onSaved={async () => {
                        setScoring(null);
                        await fetchRows();
                    }}
                />
            )}

            {deleting && (
                <div className="fixed inset-0 z-40 flex items-center justify-center bg-black/60 p-4">
                    <div className="w-full max-w-md rounded-xl border border-blue-900/40 bg-[#0a1428] p-5 shadow-2xl">
                        <h3 className="text-base font-semibold text-slate-100">
                            删除加分项评分
                        </h3>
                        <p className="mt-2 text-sm text-slate-400">
                            确认删除「{deleting.company_name}」{deleting.year}{' '}
                            年度的全部加分项评分？此操作不可恢复。
                        </p>
                        <div className="mt-5 flex justify-end gap-2">
                            <button
                                type="button"
                                onClick={() => setDeleting(null)}
                                disabled={deleteLoading}
                                className="rounded-lg border border-blue-900/40 bg-[#070d1f] px-4 py-2 text-sm text-slate-300 transition-colors hover:bg-blue-950/40 disabled:opacity-50"
                            >
                                取消
                            </button>
                            <button
                                type="button"
                                onClick={handleDelete}
                                disabled={deleteLoading}
                                className="inline-flex items-center gap-1 rounded-lg border border-red-500/30 bg-red-600/20 px-4 py-2 text-sm font-medium text-red-300 transition-colors hover:bg-red-600/30 disabled:opacity-50"
                            >
                                {deleteLoading && (
                                    <Loader2 className="h-3.5 w-3.5 animate-spin" />
                                )}
                                确认删除
                            </button>
                        </div>
                    </div>
                </div>
            )}
        </>
    );
}
