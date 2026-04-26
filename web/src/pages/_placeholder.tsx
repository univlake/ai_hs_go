import { useLocation } from 'react-router-dom';
import { Construction } from 'lucide-react';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';

// Default fallback for SPA routes that exist in the Laravel sidebar but
// haven't been wired into the new SPA router yet (most of the 172 pages).
// Shown so users get a clear "尚未迁移" message instead of being silently
// redirected to /dashboard.
export default function PlaceholderPage() {
    const location = useLocation();
    return (
        <div className="space-y-4">
            <Card>
                <CardHeader>
                    <CardTitle className="flex items-center gap-2 text-lg">
                        <Construction className="h-5 w-5 text-amber-500" />
                        模块尚未迁移
                    </CardTitle>
                </CardHeader>
                <CardContent className="space-y-2 text-sm text-[var(--color-text-muted)]">
                    <p>
                        路径{' '}
                        <code className="rounded bg-[var(--color-surface-muted)] px-1.5 py-0.5">
                            {location.pathname}
                        </code>{' '}
                        对应的页面文件已从 Laravel 项目复制到
                        <code className="ml-1">web/src/pages/</code>， 但还没接入 SPA
                        路由表（<code>web/src/router.tsx</code>），也尚未替换其内部的
                        Inertia / wayfinder 调用。
                    </p>
                    <p>
                        参见迁移计划{' '}
                        <code>/Users/zero/.claude/plans/...md</code>
                        第十章的功能清单。当前 MVP 范围只覆盖：
                    </p>
                    <ul className="list-inside list-disc pl-2">
                        <li>认证 / 账号中心 (login / first-login-setup / profile)</li>
                        <li>平台超管：一级企业管理 + 企业森林</li>
                        <li>安全台账 API（特种设备 / 危险设备 / 危险化学品 — UI 待迁）</li>
                    </ul>
                </CardContent>
            </Card>
        </div>
    );
}
