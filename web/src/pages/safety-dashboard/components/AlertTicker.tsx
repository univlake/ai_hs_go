import { useEffect, useState } from 'react';
import type { AlertEntry } from '../types';

interface AlertTickerProps {
  alerts: AlertEntry[];
  /** ms between alert rotations. */
  interval?: number;
}

// Cycles through the alert list, one entry every `interval` ms. The text fades
// in via the keyframe in styles.css; the leading status dot pulses
// continuously with a color matching the alert level.
export function AlertTicker({ alerts, interval = 3500 }: AlertTickerProps) {
  const [idx, setIdx] = useState(0);

  useEffect(() => {
    if (alerts.length === 0) return;
    const id = setInterval(() => {
      setIdx((i) => (i + 1) % alerts.length);
    }, interval);
    return () => clearInterval(id);
  }, [alerts.length, interval]);

  if (alerts.length === 0) return null;
  const a = alerts[idx % alerts.length];
  if (!a) return null;

  const colors: Record<AlertEntry['level'], string> = {
    major: 'var(--warn)',
    normal: 'var(--accent)',
    low: 'var(--primary)',
  };

  return (
    <div
      style={{
        display: 'flex',
        alignItems: 'center',
        gap: 10,
        fontSize: 11,
        padding: '0 14px',
        height: 28,
        background: 'var(--card)',
        border: '1px solid var(--border)',
        borderRadius: 4,
        overflow: 'hidden',
        flex: 1,
        minWidth: 0,
      }}
    >
      <span
        style={{
          display: 'flex',
          alignItems: 'center',
          gap: 5,
          fontWeight: 500,
          color: colors[a.level],
          flexShrink: 0,
        }}
      >
        <span
          className="safety-dashboard-pulse-fast"
          style={{
            width: 5,
            height: 5,
            borderRadius: '50%',
            background: 'currentColor',
          }}
        />
        最新告警
      </span>
      <span style={{ opacity: 0.55, fontFamily: 'var(--font-num)', flexShrink: 0 }}>
        {a.time}
      </span>
      <span
        key={idx}
        className="safety-dashboard-tick-in"
        style={{
          opacity: 0.85,
          whiteSpace: 'nowrap',
          overflow: 'hidden',
          textOverflow: 'ellipsis',
          flex: 1,
          minWidth: 0,
        }}
      >
        {a.text}
      </span>
    </div>
  );
}
