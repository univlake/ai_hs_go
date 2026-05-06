import { useState } from 'react';
import { CountUp } from '../charts/CountUp';

interface KpiPillProps {
  label: string;
  value: number;
  /** Trailing text after the number — e.g. "次 · 413人" or "项 · 三级". */
  unit: string;
  /** Resolves to a CSS color via var() or hex. */
  color: string;
  animate: boolean;
}

// Top-strip KPI tile used by the Innovative variant. Six of these line up
// across the dashboard above the main grid; each one mirrors the Card hover
// treatment (border color shift) without the full Card chrome.
export function KpiPill({ label, value, unit, color, animate }: KpiPillProps) {
  const [hov, setHov] = useState(false);
  return (
    <div
      onMouseEnter={() => setHov(true)}
      onMouseLeave={() => setHov(false)}
      style={{
        padding: '10px 14px',
        border: hov ? '1px solid var(--primary-30)' : '1px solid var(--border)',
        borderRadius: 6,
        background: 'var(--card)',
        display: 'flex',
        alignItems: 'center',
        gap: 12,
        transition: 'all 0.25s',
        position: 'relative',
        overflow: 'hidden',
      }}
    >
      <div style={{ width: 3, height: 28, background: color, borderRadius: 1 }} />
      <div style={{ flex: 1, minWidth: 0 }}>
        <div style={{ fontSize: 10, opacity: 0.65, marginBottom: 2 }}>{label}</div>
        <div style={{ display: 'flex', alignItems: 'baseline', gap: 4 }}>
          <span
            style={{
              fontFamily: 'var(--font-num)',
              fontSize: 22,
              fontWeight: 700,
              color,
              letterSpacing: 0.5,
            }}
          >
            {animate ? <CountUp value={value} /> : value.toLocaleString()}
          </span>
          <span style={{ fontSize: 10, opacity: 0.6, whiteSpace: 'nowrap' }}>{unit}</span>
        </div>
      </div>
    </div>
  );
}
