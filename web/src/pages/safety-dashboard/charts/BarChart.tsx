import { useEffect, useRef, useState } from 'react';
import { useTooltip } from './useTooltip';

export interface BarRow {
  label: string;
  [key: string]: string | number | null | undefined;
}

export interface BarSeries {
  key: string;
  color: string;
  name: string;
}

interface BarChartProps {
  data: BarRow[];
  series: BarSeries[];
  height?: number;
  showAxis?: boolean;
  formatY?: (v: number) => string;
  animate?: boolean;
}

// Multi-series vertical bar chart. Width tracks the container via
// ResizeObserver so the chart stays sharp at any breakpoint. Each series is a
// separate rect inside a per-row group; gaps and bar widths are sized to keep
// 1–8 categories readable without per-call tuning.
export function BarChart({
  data,
  series,
  height = 140,
  showAxis = true,
  formatY = (v) => String(v),
  animate = true,
}: BarChartProps) {
  const { wrapRef, onMove, onLeave, tipNode } = useTooltip();
  const PAD = { l: 30, r: 8, t: 10, b: 22 };
  const [w, setW] = useState(400);
  const innerRef = useRef<HTMLDivElement | null>(null);

  useEffect(() => {
    const el = innerRef.current;
    if (!el) return;
    const ro = new ResizeObserver((entries) => {
      const e = entries[0];
      if (!e) return;
      setW(e.contentRect.width);
    });
    ro.observe(el);
    return () => ro.disconnect();
  }, []);

  const innerW = w - PAD.l - PAD.r;
  const innerH = height - PAD.t - PAD.b;
  const allVals = data
    .flatMap((d) => series.map((s) => Number(d[s.key])))
    .filter((v) => Number.isFinite(v));
  const max = Math.max(...allVals, 1);
  const niceMax = Math.ceil(max / 5) * 5 || max;
  const groupW = innerW / Math.max(1, data.length);
  const barW = Math.max(4, Math.min(14, (groupW - 6) / series.length));

  const yTicks = [0, 0.25, 0.5, 0.75, 1].map((t) => niceMax * t);

  return (
    <div
      ref={(el) => {
        innerRef.current = el;
        wrapRef.current = el;
      }}
      style={{ position: 'relative', width: '100%', height }}
    >
      <svg width={w} height={height}>
        {showAxis &&
          yTicks.map((v, i) => {
            const y = PAD.t + innerH - (v / niceMax) * innerH;
            return (
              <g key={i}>
                <line x1={PAD.l} y1={y} x2={w - PAD.r} y2={y} stroke="rgba(255,255,255,0.06)" />
                <text
                  x={PAD.l - 4}
                  y={y + 3}
                  textAnchor="end"
                  fontSize="9"
                  fill="currentColor"
                  fillOpacity="0.5"
                >
                  {formatY(v)}
                </text>
              </g>
            );
          })}
        {data.map((d, i) => {
          const cx = PAD.l + i * groupW + groupW / 2;
          return (
            <g key={i}>
              {series.map((s, si) => {
                const v = Number(d[s.key]) || 0;
                const h = (v / niceMax) * innerH;
                const x = cx - (series.length * barW) / 2 - 1 + si * (barW + 2);
                const y = PAD.t + innerH - h;
                return (
                  <rect
                    key={s.key}
                    x={x}
                    y={y}
                    width={barW}
                    height={h}
                    fill={s.color}
                    rx={1.5}
                    opacity={animate ? 0 : 1}
                    onMouseMove={(e) =>
                      onMove(e, `${d.label} · ${s.name}: ${v.toLocaleString()}`)
                    }
                    onMouseLeave={onLeave}
                    style={{ cursor: 'pointer' }}
                  >
                    {animate ? (
                      <animate
                        attributeName="opacity"
                        from="0"
                        to="1"
                        dur="0.5s"
                        begin={`${i * 0.04}s`}
                        fill="freeze"
                      />
                    ) : null}
                  </rect>
                );
              })}
              {showAxis ? (
                <text
                  x={cx}
                  y={height - 6}
                  textAnchor="middle"
                  fontSize="9"
                  fill="currentColor"
                  fillOpacity="0.5"
                >
                  {d.label}
                </text>
              ) : null}
            </g>
          );
        })}
      </svg>
      {tipNode}
    </div>
  );
}
