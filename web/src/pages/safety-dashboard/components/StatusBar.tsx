import { AlertTicker } from './AlertTicker';
import type { AlertEntry } from '../types';

interface StatusBarProps {
  updatedAt: string;
  alerts: AlertEntry[];
  /** Number of upstream data sources powering the dashboard. */
  sources?: number;
  version?: string;
}

export function StatusBar({
  updatedAt,
  alerts,
  sources = 17,
  version = 'v2.6.0',
}: StatusBarProps) {
  return (
    <div
      style={{
        height: 32,
        padding: '0 16px',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'space-between',
        borderTop: '1px solid var(--border)',
        fontSize: 11,
        opacity: 0.7,
        background: 'var(--card)',
      }}
    >
      <div style={{ display: 'flex', alignItems: 'center', gap: 16 }}>
        <span style={{ display: 'flex', alignItems: 'center', gap: 5 }}>
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
          系统在线
        </span>
        <span>
          数据更新：<span style={{ fontFamily: 'var(--font-num)' }}>{updatedAt}</span>
        </span>
        <span>
          采集源：
          <span style={{ fontFamily: 'var(--font-num)', color: 'var(--primary)' }}>
            {sources}
          </span>{' '}
          个
        </span>
      </div>
      <AlertTicker alerts={alerts} />
      <div style={{ display: 'flex', gap: 14, marginLeft: 14, flexShrink: 0 }}>
        <span>{version}</span>
        <span style={{ fontFamily: 'var(--font-num)' }}>© 杭实集团安全技术中心</span>
      </div>
    </div>
  );
}
