import type { CSSProperties } from 'react';

export interface LegendItem {
  name: string;
  color: string;
  shape?: 'square' | 'dot';
}

interface LegendProps {
  items: LegendItem[];
  style?: CSSProperties;
}

export function Legend({ items, style }: LegendProps) {
  return (
    <div style={{ display: 'flex', gap: 10, fontSize: 10, ...style }}>
      {items.map((i) => (
        <div key={i.name} style={{ display: 'flex', alignItems: 'center', gap: 4 }}>
          {i.shape === 'dot' ? (
            <span
              style={{
                width: 6,
                height: 6,
                background: i.color,
                borderRadius: '50%',
              }}
            />
          ) : (
            <span
              style={{
                width: 8,
                height: 8,
                background: i.color,
                borderRadius: 2,
              }}
            />
          )}
          <span style={{ opacity: 0.75 }}>{i.name}</span>
        </div>
      ))}
    </div>
  );
}
