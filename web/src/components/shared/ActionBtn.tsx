import { cn } from '@/lib/utils';

// 跟 aifix_hs 项目（resources/js/pages/system/{users,roles}/Index.tsx）的 ActionBtn 一致：
//   尺寸 7×7 圆角小方块，常态 text-slate-600，hover 时按 variant 转蓝色 / 红色。
//   edit   variant — 编辑 / 重置密码 / 重置等无害动作
//   delete variant — 删除 / 移除等危险动作
export function ActionBtn({
    children,
    onClick,
    title,
    variant = 'edit',
    disabled,
}: {
    children: React.ReactNode;
    onClick: () => void;
    title: string;
    variant?: 'edit' | 'delete';
    disabled?: boolean;
}) {
    return (
        <button
            onClick={onClick}
            title={title}
            disabled={disabled}
            className={cn(
                'flex h-7 w-7 items-center justify-center rounded-md transition-colors disabled:pointer-events-none disabled:opacity-40',
                variant === 'edit'
                    ? 'text-slate-600 hover:bg-blue-600/10 hover:text-blue-400'
                    : 'text-slate-600 hover:bg-red-600/10 hover:text-red-400',
            )}
        >
            {children}
        </button>
    );
}
