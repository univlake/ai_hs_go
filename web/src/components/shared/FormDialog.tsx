import {
    Dialog,
    DialogContent,
    DialogTitle,
} from '@/components/ui/dialog';
import { Button } from '@/components/ui/button';
import { Loader2, X, type LucideIcon } from 'lucide-react';
import { cn } from '@/lib/utils';

type Tone = 'primary' | 'amber' | 'rose';

const toneClasses: Record<Tone, string> = {
    primary: 'bg-blue-600 hover:bg-blue-500',
    amber: 'bg-amber-600 hover:bg-amber-500',
    rose: 'bg-rose-600 hover:bg-rose-500',
};

interface FormDialogProps {
    open: boolean;
    onOpenChange: (v: boolean) => void;
    title: string;
    /** Optional title prefix icon. */
    icon?: LucideIcon;
    /** Optional subtitle / hint shown right under the title. */
    description?: React.ReactNode;
    children: React.ReactNode;
    /** Default `lg`. md ≈ 480, lg ≈ 640, xl ≈ 960. */
    size?: 'md' | 'lg' | 'xl';

    onCancel?: () => void;
    cancelLabel?: string;
    onConfirm: () => void;
    confirmLabel?: string;
    confirmTone?: Tone;
    loading?: boolean;
    /** Disable confirm without a busy spinner (e.g. invalid form). */
    confirmDisabled?: boolean;
}

/**
 * FormDialog wraps shadcn Dialog (built on @base-ui/react) but overrides the
 * default DialogContent layout so the form has:
 *   - dark techy palette matching Dashboard / Tenants
 *   - a fixed three-row layout: header → scrollable body → footer
 *   - no internal gap, no -m hacks, no overlapping default close button
 *
 * The shadcn DialogHeader/DialogFooter have built-in negative margins for
 * their default padded shell — we don't use them and write our own divs.
 */
export function FormDialog({
    open,
    onOpenChange,
    title,
    icon: Icon,
    description,
    children,
    size = 'lg',
    onCancel,
    cancelLabel = '取消',
    onConfirm,
    confirmLabel = '保存',
    confirmTone = 'primary',
    loading = false,
    confirmDisabled = false,
}: FormDialogProps) {
    const sizeCls = {
        md: 'sm:max-w-lg',
        lg: 'sm:max-w-2xl',
        xl: 'sm:max-w-4xl',
    }[size];

    return (
        <Dialog open={open} onOpenChange={onOpenChange}>
            <DialogContent
                showCloseButton={false}
                className={cn(
                    sizeCls,
                    // Override shadcn defaults: gap-0 + p-0 so we can fully
                    // control the three-row layout below; keep the rounded
                    // corners by repeating rounded-xl explicitly.
                    'flex max-h-[90vh] flex-col gap-0 overflow-hidden rounded-xl border-blue-900/50 bg-[#0d1530] p-0 text-slate-100 shadow-2xl shadow-black/60',
                )}
            >
                {/* Header */}
                <header className="flex shrink-0 items-start justify-between gap-3 border-b border-blue-900/30 px-6 py-4">
                    <div className="min-w-0 flex-1">
                        <DialogTitle className="flex items-center gap-2 text-base font-semibold text-slate-100">
                            {Icon && <Icon className="h-4 w-4 shrink-0 text-blue-400" />}
                            <span className="truncate">{title}</span>
                        </DialogTitle>
                        {description && (
                            <div className="mt-1 text-xs leading-relaxed text-slate-500">
                                {description}
                            </div>
                        )}
                    </div>
                    <button
                        type="button"
                        onClick={() => onOpenChange(false)}
                        aria-label="关闭"
                        className="flex h-7 w-7 shrink-0 items-center justify-center rounded-lg text-slate-600 transition-colors hover:bg-white/5 hover:text-slate-300"
                    >
                        <X className="h-4 w-4" />
                    </button>
                </header>

                {/* Body */}
                <div className="min-h-0 flex-1 overflow-y-auto px-6 py-5">
                    {children}
                </div>

                {/* Footer */}
                <footer className="flex shrink-0 items-center justify-end gap-2 border-t border-blue-900/30 bg-[#0a1428]/40 px-6 py-3">
                    {cancelLabel && (
                        <Button
                            variant="outline"
                            onClick={() => {
                                if (onCancel) onCancel();
                                else onOpenChange(false);
                            }}
                            className="border-blue-900/40 bg-transparent text-slate-300 hover:bg-blue-600/10 hover:text-slate-100"
                        >
                            {cancelLabel}
                        </Button>
                    )}
                    <Button
                        onClick={onConfirm}
                        disabled={loading || confirmDisabled}
                        className={cn('min-w-[88px]', toneClasses[confirmTone])}
                    >
                        {loading && (
                            <Loader2 className="mr-1.5 h-3.5 w-3.5 animate-spin" />
                        )}
                        {loading ? '提交中…' : confirmLabel}
                    </Button>
                </footer>
            </DialogContent>
        </Dialog>
    );
}

// Field row helper — labels above inputs, matching the Laravel project's
// uppercase tracked typography. The label colour, asterisk colour and
// spacing all line up with `Field` / hand-rolled labels in the legacy code so
// FormDialog forms render visually 1:1 with the original.
export function FormRow({
    label,
    required,
    hint,
    error,
    children,
    cols = 1,
    colSpan,
    labelTrailing,
}: {
    label?: string;
    required?: boolean;
    hint?: string;
    error?: string;
    children: React.ReactNode;
    /** Render two columns inside this row (rare; prefer wrapping in a
     *  `grid grid-cols-2 gap-4` container instead). */
    cols?: 1 | 2;
    /** Span this many columns of the parent grid. Mirrors Laravel `colSpan`. */
    colSpan?: 1 | 2 | 3 | 4;
    /** Extra inline element rendered next to the label — used for amber
     *  validation hints ("请填写唯一编号") that the legacy forms place inside
     *  the label rather than below the input. */
    labelTrailing?: React.ReactNode;
}) {
    const spanCls =
        colSpan === 2
            ? 'col-span-2'
            : colSpan === 3
              ? 'col-span-3'
              : colSpan === 4
                ? 'col-span-4'
                : '';

    if (cols === 2) {
        return (
            <div className={cn('grid grid-cols-2 gap-4', spanCls)}>
                {children}
            </div>
        );
    }

    return (
        <div className={cn('space-y-1.5', spanCls)}>
            {label && (
                <label className="mb-1.5 block text-xs font-medium tracking-wide text-slate-500 uppercase">
                    {label}
                    {required && <span className="ml-1 text-blue-500">*</span>}
                    {labelTrailing && (
                        <span className="normal-case">{labelTrailing}</span>
                    )}
                </label>
            )}
            {children}
            {(hint || error) && (
                <p
                    className={cn(
                        'text-[11px] leading-tight',
                        error ? 'text-rose-400' : 'text-slate-600',
                    )}
                >
                    {error || hint}
                </p>
            )}
        </div>
    );
}

// FormSection groups related fields under a titled separator. Matches the
// Laravel `FormSection` helper used in hazardous-chemicals — title styled
// like a section header with an optional leading icon and a thin rule on
// the right.
export function FormSection({
    title,
    icon: Icon,
    children,
}: {
    title: string;
    icon?: LucideIcon;
    children: React.ReactNode;
}) {
    return (
        <section>
            <div className="mb-3 flex items-center gap-2">
                {Icon && <Icon className="h-3.5 w-3.5 text-blue-500" />}
                <h3 className="text-xs font-semibold tracking-widest text-slate-600 uppercase">
                    {title}
                </h3>
                <div className="flex-1 border-t border-blue-900/20" />
            </div>
            {children}
        </section>
    );
}
