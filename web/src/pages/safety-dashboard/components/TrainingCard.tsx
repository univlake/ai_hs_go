import { Card } from './Card';
import { BarChart } from '../charts/BarChart';
import { Legend } from './Legend';
import type { MonthlyTraining } from '../types';
import type { Density } from '../theme';

interface TrainingCardProps {
  data: MonthlyTraining[];
  density: Density;
  animate: boolean;
}

export function TrainingCard({ data, density, animate }: TrainingCardProps) {
  return (
    <Card
      title="培训情况"
      density={density}
      style={{ height: '100%' }}
      action={
        <Legend
          items={[
            { name: '培训人数', color: 'var(--primary)' },
            { name: '培训次数', color: 'var(--warn)' },
          ]}
        />
      }
    >
      <div style={{ flex: 1, minHeight: 0 }}>
        <BarChart
          data={data.map((t) => ({ label: t.month, people: t.people, sessions: t.sessions }))}
          series={[
            { key: 'people', color: 'var(--primary)', name: '培训人数' },
            { key: 'sessions', color: 'var(--warn)', name: '培训次数' },
          ]}
          height={170}
          animate={animate}
        />
      </div>
    </Card>
  );
}
