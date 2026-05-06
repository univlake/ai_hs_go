import { useTooltip } from './useTooltip';

export interface DonutSlice {
  label: string;
  value: number;
  color?: string;
}

interface DonutChartProps {
  data: DonutSlice[];
  size?: number;
  thickness?: number;
  centerTop?: string;
  centerBottom?: string;
  /** Resolve a slice's stroke from its label / data-color hint. */
  accentByLabel?: (label: string, hint: string | undefined) => string;
}

// SVG donut chart. Slices are drawn as overlapping arcs via stroke-dasharray
// trickery, rotated -90° so the first slice starts at 12 o'clock. The hover
// tooltip uses the shared useTooltip primitive.
export function DonutChart({
  data,
  size = 160,
  thickness = 22,
  centerTop,
  centerBottom,
  accentByLabel,
}: DonutChartProps) {
  const total = data.reduce((s, d) => s + d.value, 0) || 1;
  const r = size / 2 - thickness / 2;
  const c = 2 * Math.PI * r;
  let acc = 0;
  const { wrapRef, onMove, onLeave, tipNode } = useTooltip();

  return (
    <div ref={wrapRef} style={{ position: 'relative', width: size, height: size }}>
      <svg width={size} height={size} style={{ transform: 'rotate(-90deg)' }}>
        <circle
          cx={size / 2}
          cy={size / 2}
          r={r}
          fill="none"
          stroke="rgba(255,255,255,0.06)"
          strokeWidth={thickness}
        />
        {data.map((d, i) => {
          const frac = d.value / total;
          const dash = frac * c;
          const offset = -acc * c;
          acc += frac;
          const color = accentByLabel?.(d.label, d.color) ?? d.color ?? '#3b82f6';
          return (
            <circle
              key={i}
              cx={size / 2}
              cy={size / 2}
              r={r}
              fill="none"
              stroke={color}
              strokeWidth={thickness}
              strokeDasharray={`${dash} ${c}`}
              strokeDashoffset={offset}
              strokeLinecap="butt"
              onMouseMove={(e) =>
                onMove(
                  e,
                  `${d.label} · ${d.value.toLocaleString()} (${(frac * 100).toFixed(1)}%)`,
                )
              }
              onMouseLeave={onLeave}
              style={{ transition: 'stroke 0.3s', cursor: 'pointer' }}
            />
          );
        })}
      </svg>
      <div
        style={{
          position: 'absolute',
          inset: 0,
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
          justifyContent: 'center',
          pointerEvents: 'none',
        }}
      >
        {centerTop ? (
          <div style={{ fontSize: 11, opacity: 0.55, letterSpacing: 1 }}>{centerTop}</div>
        ) : null}
        {centerBottom ? (
          <div style={{ fontSize: 22, fontFamily: 'var(--font-num)', fontWeight: 600 }}>
            {centerBottom}
          </div>
        ) : null}
      </div>
      {tipNode}
    </div>
  );
}
