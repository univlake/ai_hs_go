import { useState, useEffect } from 'react';
import {
    AlertTriangle,
    Check,
    Copy,
    Download,
    KeyRound,
} from 'lucide-react';

// One-time admin credentials popup, shown after a sub-company / tenant is
// created. The plaintext password only exists on the server's response —
// once dismissed it can't be recovered, so the modal forces the operator
// to acknowledge they've copied or downloaded it before closing.
//
// Used by both 平台/企业森林 (Forest) and 集团/下级企业 (Subsidiaries).

export interface AdminCredentials {
    companyTitle: string;
    username: string;
    password: string;
}

interface Props {
    credentials: AdminCredentials | null;
    onClose: () => void;
}

export function CredentialsModal({ credentials, onClose }: Props) {
    const [copied, setCopied] = useState<'username' | 'password' | null>(null);
    const [acknowledged, setAcknowledged] = useState(false);

    useEffect(() => {
        if (!credentials) {
            setCopied(null);
            setAcknowledged(false);
        }
    }, [credentials]);

    if (!credentials) {
        return null;
    }

    const copy = async (text: string, which: 'username' | 'password') => {
        try {
            await navigator.clipboard.writeText(text);
            setCopied(which);
            setTimeout(
                () => setCopied((c) => (c === which ? null : c)),
                1500,
            );
        } catch {
            // clipboard unavailable — user can still select+copy manually
        }
    };

    const download = () => {
        const content = [
            `企业名称：${credentials.companyTitle}`,
            `管理员账号：${credentials.username}`,
            `初始密码：${credentials.password}`,
            '',
            '请妥善保管此凭据。建议登录后立即修改密码。',
            `生成时间：${new Date().toLocaleString('zh-CN')}`,
        ].join('\n');
        const blob = new Blob([content], { type: 'text/plain;charset=utf-8' });
        const url = URL.createObjectURL(blob);
        const a = document.createElement('a');
        a.href = url;
        a.download = `${credentials.username}.txt`;
        a.click();
        URL.revokeObjectURL(url);
    };

    return (
        <div className="fixed inset-0 z-50 flex items-center justify-center p-4">
            <div
                className="absolute inset-0 bg-black/60 backdrop-blur-sm"
                onClick={() => {}}
            />
            <div className="relative flex max-h-[92vh] w-full max-w-lg flex-col overflow-hidden rounded-2xl border border-blue-900/50 bg-[#0d1530] shadow-2xl shadow-black/60">
                <div className="p-6">
                    <div className="mb-4 flex h-11 w-11 items-center justify-center rounded-xl border border-emerald-500/20 bg-emerald-600/10">
                        <KeyRound className="h-5 w-5 text-emerald-400" />
                    </div>
                    <h2 className="mb-1 text-base font-semibold text-slate-100">
                        企业已创建
                    </h2>
                    <p className="mb-5 text-sm text-slate-500">
                        已为{' '}
                        <span className="font-medium text-slate-300">
                            {credentials.companyTitle}
                        </span>{' '}
                        自动生成管理员账号。
                    </p>

                    <div className="mb-4 flex items-start gap-2 rounded-lg border border-amber-500/20 bg-amber-600/8 px-3 py-2.5 text-xs text-amber-400">
                        <AlertTriangle className="mt-0.5 h-3.5 w-3.5 shrink-0" />
                        <span>
                            此密码只显示一次，关闭后无法再次查看。请立即复制或下载保存。
                        </span>
                    </div>

                    <div className="space-y-3">
                        <CredentialRow
                            label="账号"
                            value={credentials.username}
                            copied={copied === 'username'}
                            onCopy={() =>
                                copy(credentials.username, 'username')
                            }
                        />
                        <CredentialRow
                            label="密码"
                            value={credentials.password}
                            copied={copied === 'password'}
                            onCopy={() =>
                                copy(credentials.password, 'password')
                            }
                        />
                    </div>

                    <button
                        onClick={download}
                        className="mt-4 flex w-full items-center justify-center gap-2 rounded-lg border border-blue-500/30 bg-blue-600/10 px-4 py-2 text-sm font-medium text-blue-300 hover:bg-blue-600/20"
                    >
                        <Download className="h-4 w-4" />
                        下载为 .txt 文件
                    </button>

                    <label className="mt-5 flex cursor-pointer items-center gap-2 text-xs text-slate-400">
                        <input
                            type="checkbox"
                            checked={acknowledged}
                            onChange={(e) => setAcknowledged(e.target.checked)}
                            className="h-4 w-4 rounded border-blue-900/40 bg-[#070d1f] text-blue-500 focus:ring-1 focus:ring-blue-500/40"
                        />
                        我已保存账号密码
                    </label>
                </div>
                <div className="flex items-center justify-end gap-2 border-t border-blue-900/30 px-6 py-4">
                    <button
                        onClick={onClose}
                        disabled={!acknowledged}
                        className="rounded-lg border border-blue-500/30 bg-blue-600/15 px-4 py-2 text-sm font-medium text-blue-300 hover:bg-blue-600/25 disabled:pointer-events-none disabled:opacity-40"
                    >
                        完成
                    </button>
                </div>
            </div>
        </div>
    );
}

function CredentialRow({
    label,
    value,
    copied,
    onCopy,
}: {
    label: string;
    value: string;
    copied: boolean;
    onCopy: () => void;
}) {
    return (
        <div>
            <div className="mb-1 text-xs font-medium tracking-wide text-slate-500 uppercase">
                {label}
            </div>
            <div className="flex items-center gap-2 rounded-lg border border-blue-900/40 bg-[#070d1f] px-3 py-2">
                <code className="flex-1 font-mono text-sm text-slate-200 break-all select-all">
                    {value}
                </code>
                <button
                    type="button"
                    onClick={onCopy}
                    className="flex h-7 w-7 shrink-0 items-center justify-center rounded-md text-slate-500 transition-colors hover:bg-blue-600/10 hover:text-blue-300"
                    title="复制"
                >
                    {copied ? (
                        <Check className="h-3.5 w-3.5 text-emerald-400" />
                    ) : (
                        <Copy className="h-3.5 w-3.5" />
                    )}
                </button>
            </div>
        </div>
    );
}
