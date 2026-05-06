import { useEffect, useRef, useState } from 'react';
import { useTooltip } from './useTooltip';

export interface LineRow {
  month?: string;
  label?: string;
  [key: string]: string | number | null | undefined;
}

interface LineChartProps {
  data: LineRow[];
  valueKey?: string;
  label?: string;
  color?: string;
  height?: number;
  animate?: boolean;
  formatY?: (v: number) => string;
}

// Single-series area + line chart. The area uses a vertical gradient (color
// at 35% → transparent) and the stroke animates in via stroke-dashoffset.
// Dots are interactive; the trailing label below each point uses month/label.
export function LineChart({
  data,
  valueKey = 'value',
  label = '安全投入',
  color = '#3b82f6',
  height = 140,
  animate = true,
  formatY = (v) => String(v),
}: LineChartProps) {
  const { wrapRef, onMove, onLeave, tipNode } = useTooltip();
  const PAD = { l: 36, r: 8, t: 10, b: 22 };
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
  const max = Math.max(...data.map((d) => Number(d[valueKey]) || 0), 1);
  const niceMax = Math.ceil(max * 1.1);
  const stepX = innerW / Math.max(1, data.length - 1);
  const points = data.map((d, i) => ({
    x: PAD.l + i * stepX,
    y: PAD.t + innerH - ((Number(d[valueKey]) || 0) / niceMax) * innerH,
    raw: d,
  }));
  const path = points.map((p, i) => `${i === 0 ? 'M' : 'L'} ${p.x} ${p.y}`).join(' ');
  const last = points[points.length - 1];
  const first = points[0];
  const areaPath =
    last && first
      ? `${path} L ${last.x} ${PAD.t + innerH} L ${first.x} ${PAD.t + innerH} Z`
      : path;

  const gid = `line-grad-${color.replace(/[^a-z0-9]/gi, '')}`;

  return (
    <div
      ref={(el) => {
        innerRef.current = el;
        wrapRef.current = el;
      }}
      style={{ position: 'relative', width: '100%', height }}
    >
      <svg width={w} height={height}>
        <defs>
          <linearGradient id={gid} x1="0" y1="0" x2="0" y2="1">
            <stop offset="0%" stopColor={color} stopOpacity="0.35" />
            <stop offset="100%" stopColor={color} stopOpacity="0" />
          </linearGradient>
        </defs>
        {[0, 0.5, 1].map((t, i) => {
          const y = PAD.t + innerH - innerH * t;
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
                {formatY(Math.round(niceMax * t))}
              </text>
            </g>
          );
        })}
        <path d={areaPath} fill={`url(#${gid})`} />
        <path
          d={path}
          fill="none"
          stroke={color}
          strokeWidth={1.5}
          strokeLinejoin="round"
          strokeLinecap="round"
          strokeDasharray={animate ? 1000 : 0}
          strokeDashoffset={animate ? 1000 : 0}
        >
          {animate ? (
            <animate
              attributeName="stroke-dashoffset"
              from="1000"
              to="0"
              dur="1.2s"
              fill="freeze"
            />
          ) : null}
        </path>
        {points.map((p, i) => {
          const text = String(p.raw.month ?? p.raw.label ?? '');
          const v = Number(p.raw[valueKey]) || 0;
          return (
            <g key={i}>
              <circle
                cx={p.x}
                cy={p.y}
                r={2.5}
                fill={color}
                onMouseMove={(e) => onMove(e, `${text} · ${label}: ${v.toLocaleString()}`)}
                onMouseLeave={onLeave}
                style={{ cursor: 'pointer' }}
              />
              <text
                x={p.x}
                y={height - 6}
                textAnchor="middle"
                fontSize="9"
                fill="currentColor"
                fillOpacity="0.5"
              >
                {text}
              </text>
            </g>
          );
        })}
      </svg>
      {tipNode}
    </div>
  );
}
