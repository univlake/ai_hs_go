import { Card } from './Card';
import { BarChart } from '../charts/BarChart';
import { Legend } from './Legend';
import type { OccupationalHealthRow } from '../types';
import type { Density } from '../theme';

interface HealthCardProps {
  data: OccupationalHealthRow[];
  density: Density;
  animate: boolean;
}

export function HealthCard({ data, density, animate }: HealthCardProps) {
  return (
    <Card
      title="职业卫生"
      density={density}
      style={{ height: '100%' }}
      action={<Legend items={[{ name: '涉及人数', color: 'var(--primary)' }]} />}
    >
      <div style={{ flex: 1, minHeight: 0 }}>
        <BarChart
          data={data.map((h) => ({ label: h.hazard, people: h.people }))}
          series={[{ key: 'people', color: 'var(--primary)', name: '涉及人数' }]}
          height={140}
          animate={animate}
        />
      </div>
    </Card>
  );
}
