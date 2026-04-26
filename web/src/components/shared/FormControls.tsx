// Lightweight form controls aligned with the Laravel project's visual style.
// Used inside FormDialog so the safety ledger pages match legacy 1:1.
//
//   <NativeInput>     wraps <input>            with the dark blue inputCls
//   <NativeSelect>    wraps <select>           with the same shell + chevron
//   <NativeTextarea>  wraps <textarea>         with the same shell
//   inputCls          exported for one-off cases
//
// Why native instead of shadcn Select: shadcn (base-ui) Select.Value defaults
// to rendering the raw value, so id→title mappings show "2" instead of
// "压力容器". Native <select> renders the matching <option>'s text correctly
// and matches the legacy project's pattern exactly.

import { ChevronDown } from 'lucide-react';
import { cn } from '@/lib/utils';
import type {
    InputHTMLAttributes,
    SelectHTMLAttributes,
    TextareaHTMLAttributes,
} from 'react';

export const inputCls =
    'w-full rounded-lg border border-blue-900/40 bg-[#070d1f] px-3 py-2 text-sm text-slate-200 placeholder-slate-700 outline-none transition-colors focus:border-blue-500/60 focus:ring-1 focus:ring-blue-500/20 disabled:cursor-not-allowed disabled:opacity-60';

export function NativeInput({
    className,
    ...rest
}: InputHTMLAttributes<HTMLInputElement>) {
    return <input className={cn(inputCls, className)} {...rest} />;
}

export function NativeTextarea({
    className,
    rows = 2,
    ...rest
}: TextareaHTMLAttributes<HTMLTextAreaElement>) {
    return (
        <textarea
            rows={rows}
            className={cn(inputCls, 'resize-y leading-relaxed', className)}
            {...rest}
        />
    );
}

export interface NativeSelectOption {
    value: string;
    label: string;
}

export function NativeSelect({
    className,
    options,
    placeholder,
    children,
    ...rest
}: SelectHTMLAttributes<HTMLSelectElement> & {
    options?: NativeSelectOption[] | string[];
    placeholder?: string;
}) {
    return (
        <div className="relative">
            <select
                className={cn(
                    inputCls,
                    'cursor-pointer appearance-none pr-8',
                    className,
                )}
                {...rest}
            >
                {placeholder !== undefined && (
                    <option value="" className="bg-[#0d1530]">
                        {placeholder}
                    </option>
                )}
                {options
                    ? options.map((opt) => {
                          const value =
                              typeof opt === 'string' ? opt : opt.value;
                          const label =
                              typeof opt === 'string' ? opt : opt.label;
                          return (
                              <option
                                  key={value}
                                  value={value}
                                  className="bg-[#0d1530]"
                              >
                                  {label}
                              </option>
                          );
                      })
                    : children}
            </select>
            <ChevronDown className="pointer-events-none absolute top-1/2 right-2.5 h-4 w-4 -translate-y-1/2 text-slate-600" />
        </div>
    );
}
