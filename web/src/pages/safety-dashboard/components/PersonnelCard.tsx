import { Card } from './Card';
import { DonutChart } from '../charts/DonutChart';
import type { ColorRole, DashboardData } from '../types';
import type { Density } from '../theme';

const colorMap: Record<ColorRole, string> = {
  primary: 'var(--primary)',
  accent: 'var(--accent)',
  warn: 'var(--warn)',
};

interface PersonnelCardProps {
  data: DashboardData['personnel'];
  density: Density;
}

export function PersonnelCard({ data, density }: PersonnelCardProps) {
  return (
    <Card title="人员构成" density={density} style={{ height: '100%' }}>
      <div style={{ display: 'flex', gap: 16, alignItems: 'center', flex: 1 }}>
        <DonutChart
          size={130}
          thickness={18}
          data={data.composition.map((c) => ({
            label: c.label,
            value: c.value,
            color: c.color,
          }))}
          centerTop="总计"
          centerBottom={data.total.toLocaleString()}
          accentByLabel={(_label, hint) => colorMap[(hint as ColorRole) ?? 'primary']}
        />
        <div style={{ flex: 1, display: 'flex', flexDirection: 'column', gap: 8 }}>
          <div style={{ fontSize: 11, opacity: 0.6, marginBottom: 2 }}>专业人员情况</div>
          {data.professionals.map((p) => (
            <div key={p.label} style={{ fontSize: 11 }}>
              <div
                style={{
                  display: 'flex',
                  justifyContent: 'space-between',
                  marginBottom: 3,
                }}
              >
                <span style={{ opacity: 0.85 }}>{p.label}</span>
                <span style={{ fontFamily: 'var(--font-num)', opacity: 0.95 }}>
                  {p.value} <span style={{ opacity: 0.55 }}>({p.percent}%)</span>
                </span>
              </div>
              <div
                style={{
                  height: 3,
                  background: 'rgba(255,255,255,0.06)',
                  borderRadius: 2,
                  overflow: 'hidden',
                }}
              >
                <div
                  style={{
                    height: '100%',
                    width: `${p.percent}%`,
                    background: 'var(--primary)',
                    transition: 'width 1s ease-out',
                  }}
                />
              </div>
            </div>
          ))}
        </div>
      </div>
      <div style={{ display: 'flex', gap: 12, marginTop: 8, fontSize: 10 }}>
        {data.composition.map((c) => (
          <div key={c.label} style={{ display: 'flex', alignItems: 'center', gap: 4 }}>
            <span
              style={{
                width: 8,
                height: 8,
                borderRadius: 2,
                background: colorMap[c.color],
              }}
            />
            <span style={{ opacity: 0.75 }}>{c.label}</span>
            <span style={{ fontFamily: 'var(--font-num)', opacity: 0.55 }}>
              {c.value.toLocaleString()}
            </span>
          </div>
        ))}
      </div>
    </Card>
  );
}
