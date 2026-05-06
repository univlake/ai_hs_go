import { useTooltip } from './useTooltip';

export interface PieSlice {
  name: string;
  value: number;
  percent: number;
  color: string;
}

interface PieChartProps {
  data: PieSlice[];
  size?: number;
}

// Standard pie chart, no donut hole. Slices start at 12 o'clock and go
// clockwise. Hover lifts opacity to 1.0 for the active slice.
export function PieChart({ data, size = 180 }: PieChartProps) {
  const { wrapRef, onMove, onLeave, tipNode } = useTooltip();
  const total = data.reduce((s, d) => s + d.value, 0) || 1;
  const cx = size / 2;
  const cy = size / 2;
  const r = size / 2 - 6;
  let a0 = -Math.PI / 2;

  return (
    <div ref={wrapRef} style={{ position: 'relative', width: size, height: size }}>
      <svg width={size} height={size}>
        {data.map((d, i) => {
          const frac = d.value / total;
          const a1 = a0 + frac * Math.PI * 2;
          const x0 = cx + Math.cos(a0) * r;
          const y0 = cy + Math.sin(a0) * r;
          const x1 = cx + Math.cos(a1) * r;
          const y1 = cy + Math.sin(a1) * r;
          const large = a1 - a0 > Math.PI ? 1 : 0;
          const path = `M ${cx} ${cy} L ${x0} ${y0} A ${r} ${r} 0 ${large} 1 ${x1} ${y1} Z`;
          a0 = a1;
          return (
            <path
              key={i}
              d={path}
              fill={d.color}
              fillOpacity={0.85}
              stroke="rgba(0,0,0,0.4)"
              strokeWidth={0.5}
              onMouseMove={(e) => onMove(e, `${d.name} · ${d.value} (${d.percent}%)`)}
              onMouseLeave={onLeave}
              onMouseEnter={(e) => e.currentTarget.setAttribute('fill-opacity', '1')}
              style={{ cursor: 'pointer', transition: 'fill-opacity 0.2s' }}
            />
          );
        })}
      </svg>
      {tipNode}
    </div>
  );
}
