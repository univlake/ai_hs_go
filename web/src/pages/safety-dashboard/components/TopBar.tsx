import { useEffect, useState } from 'react';
import { CountUp } from '../charts/CountUp';

interface TopBarProps {
  title: string;
  safeDays: number;
}

// Top status bar: clock on the left, gradient title in the center, "已持续
// 安全生产 N 天" KPI on the right. The clock ticks every second; React's
// re-render budget is fine for a single useState every 1s.
export function TopBar({ title, safeDays }: TopBarProps) {
  const [now, setNow] = useState(() => new Date());
  useEffect(() => {
    const id = setInterval(() => setNow(new Date()), 1000);
    return () => clearInterval(id);
  }, []);

  const fmt = (n: number) => String(n).padStart(2, '0');
  const dateStr = `${now.getFullYear()}-${fmt(now.getMonth() + 1)}-${fmt(now.getDate())}`;
  const timeStr = `${fmt(now.getHours())}:${fmt(now.getMinutes())}:${fmt(now.getSeconds())}`;

  return (
    <div
      style={{
        height: 56,
        padding: '0 24px',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'space-between',
        position: 'relative',
        borderBottom: '1px solid var(--border)',
        background: 'linear-gradient(180deg, var(--card), transparent)',
      }}
    >
      <div style={{ display: 'flex', alignItems: 'center', gap: 12, flex: 1 }}>
        <div
          style={{
            display: 'flex',
            alignItems: 'center',
            gap: 6,
            fontSize: 12,
            opacity: 0.7,
          }}
        >
          <span
            className="safety-dashboard-pulse"
            style={{
              width: 6,
              height: 6,
              borderRadius: '50%',
              background: '#10b981',
              boxShadow: '0 0 8px #10b981',
            }}
          />
          <span style={{ fontFamily: 'var(--font-num)' }}>{dateStr}</span>
          <span style={{ fontFamily: 'var(--font-num)' }}>{timeStr}</span>
        </div>
      </div>

      <div
        style={{
          flex: 0,
          textAlign: 'center',
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
        }}
      >
        <div
          style={{
            fontSize: 22,
            fontWeight: 600,
            letterSpacing: '0.25em',
            paddingLeft: '0.25em',
            whiteSpace: 'nowrap',
            lineHeight: 1.2,
            paddingTop: 2,
            paddingBottom: 2,
            background: 'linear-gradient(180deg, var(--fg), var(--primary))',
            WebkitBackgroundClip: 'text',
            WebkitTextFillColor: 'transparent',
            backgroundClip: 'text',
          }}
        >
          {title}
        </div>
      </div>

      <div
        style={{
          flex: 1,
          display: 'flex',
          justifyContent: 'flex-end',
          alignItems: 'center',
          gap: 14,
        }}
      >
        <div style={{ fontSize: 12, opacity: 0.7 }}>已持续安全生产</div>
        <div
          style={{
            fontFamily: 'var(--font-num)',
            fontSize: 26,
            fontWeight: 700,
            color: 'var(--primary)',
            letterSpacing: 1,
          }}
        >
          <CountUp value={safeDays} />
        </div>
        <div style={{ fontSize: 12, opacity: 0.7 }}>天</div>
      </div>
    </div>
  );
}
