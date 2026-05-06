import { Card } from './Card';
import { DonutChart } from '../charts/DonutChart';
import type { StandardLevel } from '../types';
import type { Density } from '../theme';

interface StandardCardProps {
  data: StandardLevel;
  density: Density;
}

export function StandardCard({ data, density }: StandardCardProps) {
  return (
    <Card title="标准化等级" density={density} style={{ height: '100%' }}>
      <div
        style={{
          flex: 1,
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
          gap: 16,
        }}
      >
        <DonutChart
          size={140}
          thickness={18}
          data={data.breakdown.map((b) => ({
            label: b.label,
            value: b.value,
            color: b.color,
          }))}
          centerTop={data.current}
          centerBottom={String(data.total)}
        />
      </div>
      <div
        style={{
          display: 'flex',
          gap: 10,
          justifyContent: 'center',
          fontSize: 10,
          marginTop: 4,
        }}
      >
        {data.breakdown.map((b) => (
          <div key={b.label} style={{ display: 'flex', alignItems: 'center', gap: 4 }}>
            <span
              style={{
                width: 8,
                height: 8,
                background: b.color,
                borderRadius: 2,
              }}
            />
            <span style={{ opacity: 0.75 }}>{b.label}</span>
            <span style={{ fontFamily: 'var(--font-num)', opacity: 0.55 }}>{b.value}</span>
          </div>
        ))}
      </div>
    </Card>
  );
}
