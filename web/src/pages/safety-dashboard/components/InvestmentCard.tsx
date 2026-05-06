import { Card } from './Card';
import { LineChart } from '../charts/LineChart';
import type { MonthlyInvestment } from '../types';
import type { Density } from '../theme';

interface InvestmentCardProps {
  data: MonthlyInvestment[];
  density: Density;
  animate: boolean;
}

export function InvestmentCard({ data, density, animate }: InvestmentCardProps) {
  return (
    <Card
      title="安全投入"
      density={density}
      style={{ height: '100%' }}
      action={
        <span
          style={{
            fontSize: 10,
            opacity: 0.55,
            fontFamily: 'var(--font-num)',
          }}
        >
          单位：万元
        </span>
      }
    >
      <div style={{ flex: 1, minHeight: 0 }}>
        <LineChart
          data={data.map((d) => ({ month: d.month, value: d.value }))}
          valueKey="value"
          label="投入"
          color="var(--accent)"
          height={170}
          animate={animate}
        />
      </div>
    </Card>
  );
}
