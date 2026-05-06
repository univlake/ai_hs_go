import { useState } from 'react';
import { Card } from './Card';
import { CountUp } from '../charts/CountUp';
import type { DrillKpi } from '../types';
import type { Density } from '../theme';

interface DrillsCardProps {
  data: DrillKpi[];
  density: Density;
  animate: boolean;
}

export function DrillsCard({ data, density, animate }: DrillsCardProps) {
  return (
    <Card title="应急演练" density={density} style={{ height: '100%' }}>
      <div
        style={{
          flex: 1,
          display: 'grid',
          gridTemplateColumns: '1fr 1fr',
          gap: 8,
          alignContent: 'center',
        }}
      >
        {data.map((d) => (
          <DrillKpiCell key={d.label} item={d} animate={animate} />
        ))}
      </div>
    </Card>
  );
}

function DrillKpiCell({ item, animate }: { item: DrillKpi; animate: boolean }) {
  const [hov, setHov] = useState(false);
  const colorMap = { primary: 'var(--primary)', warn: 'var(--warn)' } as const;
  return (
    <div
      onMouseEnter={() => setHov(true)}
      onMouseLeave={() => setHov(false)}
      style={{
        padding: '10px 12px',
        background: hov ? 'var(--card-2)' : 'transparent',
        border: '1px solid var(--border)',
        borderRadius: 6,
        transition: 'all 0.2s',
        position: 'relative',
        overflow: 'hidden',
      }}
    >
      <div style={{ fontSize: 10, opacity: 0.65, marginBottom: 4 }}>{item.label}</div>
      <div style={{ display: 'flex', alignItems: 'baseline', gap: 3 }}>
        <span
          style={{
            fontFamily: 'var(--font-num)',
            fontSize: 26,
            fontWeight: 700,
            color: colorMap[item.color],
            letterSpacing: 1,
          }}
        >
          {animate ? <CountUp value={item.value} /> : item.value.toLocaleString()}
        </span>
        <span style={{ fontSize: 11, opacity: 0.7 }}>{item.suffix}</span>
      </div>
      <div
        style={{
          position: 'absolute',
          left: 0,
          bottom: 0,
          height: 1,
          width: hov ? '100%' : '30%',
          background: colorMap[item.color],
          transition: 'width 0.4s',
        }}
      />
    </div>
  );
}
