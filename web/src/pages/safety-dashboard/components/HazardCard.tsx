import { useState } from 'react';
import { Card } from './Card';
import { BarChart } from '../charts/BarChart';
import { Legend } from './Legend';
import type { HazardRow } from '../types';
import type { Density } from '../theme';

type Tab = 'year' | 'month';

interface HazardCardProps {
  year: HazardRow[];
  month: HazardRow[];
  density: Density;
  animate: boolean;
}

export function HazardCard({ year, month, density, animate }: HazardCardProps) {
  const [tab, setTab] = useState<Tab>('year');
  const data = tab === 'year' ? year : month;

  return (
    <Card
      title="隐患排查整改情况"
      density={density}
      style={{ height: '100%' }}
      action={
        <div
          style={{
            display: 'flex',
            fontSize: 10,
            border: '1px solid var(--border)',
            borderRadius: 4,
            overflow: 'hidden',
          }}
        >
          {(
            [
              { k: 'year', l: '本年度' },
              { k: 'month', l: '前一月' },
            ] as const
          ).map((t) => (
            <button
              key={t.k}
              type="button"
              onClick={() => setTab(t.k)}
              style={{
                padding: '3px 10px',
                background: tab === t.k ? 'var(--primary)' : 'transparent',
                color: tab === t.k ? '#fff' : 'var(--fg)',
                border: 'none',
                cursor: 'pointer',
                fontSize: 10,
                fontFamily: 'inherit',
                transition: 'background 0.2s',
              }}
            >
              {t.l}
            </button>
          ))}
        </div>
      }
    >
      <div style={{ display: 'flex', gap: 8, flex: 1, minHeight: 0 }}>
        <div style={{ flex: 1.5, minWidth: 0 }}>
          <BarChart
            data={data.map((h) => ({
              label: h.industry,
              found: h.found,
              fixed: h.fixed,
            }))}
            series={[
              { key: 'found', color: 'var(--accent)', name: '隐患排查数' },
              { key: 'fixed', color: 'var(--primary)', name: '整改数' },
            ]}
            height={150}
            animate={animate}
          />
        </div>
        <div
          style={{
            width: 120,
            display: 'flex',
            flexDirection: 'column',
            gap: 6,
            justifyContent: 'center',
          }}
        >
          {data.map((h) => (
            <div key={h.industry} style={{ fontSize: 11 }}>
              <div
                style={{
                  display: 'flex',
                  justifyContent: 'space-between',
                  marginBottom: 2,
                }}
              >
                <span style={{ opacity: 0.8 }}>{h.industry.replace('行业', '')}</span>
                <span style={{ fontFamily: 'var(--font-num)', color: 'var(--accent)' }}>
                  {h.rate}%
                </span>
              </div>
              <div
                style={{
                  height: 2,
                  background: 'rgba(255,255,255,0.06)',
                  borderRadius: 1,
                  overflow: 'hidden',
                }}
              >
                <div
                  style={{
                    height: '100%',
                    width: `${h.rate}%`,
                    background: 'var(--accent)',
                  }}
                />
              </div>
            </div>
          ))}
        </div>
      </div>
      <Legend
        items={[
          { name: '隐患排查数', color: 'var(--accent)' },
          { name: '整改数', color: 'var(--primary)' },
          { name: '整改率', color: 'var(--accent)', shape: 'dot' },
        ]}
        style={{ marginTop: 6 }}
      />
    </Card>
  );
}
