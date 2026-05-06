import { useEffect, useRef, useState } from 'react';
import * as echarts from 'echarts/core';
import { GeoComponent, TooltipComponent } from 'echarts/components';
import { EffectScatterChart, ScatterChart, MapChart } from 'echarts/charts';
import { CanvasRenderer } from 'echarts/renderers';

import type { RiskLevelInfo, RiskPoint } from '../types';

// Tree-shaken ECharts: only the geo / scatter / tooltip pieces we actually
// render. Pulls about 350KB into the bundle vs ~1MB for the full build.
echarts.use([
  GeoComponent,
  TooltipComponent,
  MapChart,
  ScatterChart,
  EffectScatterChart,
  CanvasRenderer,
]);

const ZHEJIANG_MAP_NAME = 'zhejiang';
// Authoritative provincial outline from Aliyun DataV — the recommended source
// for China admin boundaries. 330000 = 浙江省. The fetch happens once and the
// browser caches it across reloads.
const ZHEJIANG_GEOJSON_URL = 'https://geo.datav.aliyun.com/areas_v3/bound/330000_full.json';

let registerPromise: Promise<void> | null = null;

// registerZhejiangMap is idempotent across all instances of this component.
// The first caller fires the fetch + registerMap; subsequent callers await
// the same promise so we never double-fetch on a route revisit.
function registerZhejiangMap(): Promise<void> {
  if (registerPromise) return registerPromise;
  registerPromise = fetch(ZHEJIANG_GEOJSON_URL)
    .then((r) => {
      if (!r.ok) throw new Error(`fetch zhejiang geojson: ${r.status}`);
      return r.json();
    })
    .then((geoJson) => {
      echarts.registerMap(ZHEJIANG_MAP_NAME, geoJson);
    })
    .catch((err) => {
      // Reset so a future remount can retry; without this a transient network
      // glitch on first mount would permanently break the map.
      registerPromise = null;
      throw err;
    });
  return registerPromise;
}

interface ZhejiangMapProps {
  points: RiskPoint[];
  levels: RiskLevelInfo[];
  accent?: string;
  animate?: boolean;
}

export function ZhejiangMap({
  points,
  levels,
  accent = '#3b82f6',
  animate = true,
}: ZhejiangMapProps) {
  const containerRef = useRef<HTMLDivElement | null>(null);
  const chartRef = useRef<echarts.ECharts | null>(null);
  const [ready, setReady] = useState(false);
  const [error, setError] = useState<string | null>(null);

  // Register the map data once per page load.
  useEffect(() => {
    let cancelled = false;
    registerZhejiangMap()
      .then(() => {
        if (!cancelled) setReady(true);
      })
      .catch((err: Error) => {
        if (!cancelled) setError(err.message);
      });
    return () => {
      cancelled = true;
    };
  }, []);

  // Build the ECharts option once we have the map registered + the points.
  // Re-runs on points/levels change so the dashboard's polling refresh
  // re-renders the markers without rebuilding the geo layer.
  useEffect(() => {
    const el = containerRef.current;
    if (!ready || !el) return;

    if (!chartRef.current) {
      chartRef.current = echarts.init(el, undefined, { renderer: 'canvas' });
    }
    const chart = chartRef.current;

    const colorOf = (key: RiskPoint['level']) =>
      levels.find((l) => l.key === key)?.color ?? '#3b82f6';
    const labelOf = (key: RiskPoint['level']) =>
      levels.find((l) => l.key === key)?.label ?? '';

    // Build scatter data per risk level so we can color + size each tier
    // separately. severe/major use effectScatter (pulsing); normal/low use
    // plain scatter to keep frame cost reasonable when there are 100+ sites.
    const grouped: Record<RiskPoint['level'], Array<{ name: string; value: [number, number] }>> = {
      severe: [],
      major: [],
      normal: [],
      low: [],
    };
    for (const p of points) {
      grouped[p.level].push({ name: p.name, value: [p.lng, p.lat] });
    }

    const ringSize: Record<RiskPoint['level'], number> = {
      severe: 12,
      major: 10,
      normal: 8,
      low: 6,
    };

    const tooltipFormatter = (params: { name: string; data: { extra?: string } }): string => {
      const extra = params.data?.extra ?? '';
      return `${params.name}${extra ? ` · ${extra}` : ''}`;
    };

    chart.setOption(
      {
        backgroundColor: 'transparent',
        tooltip: {
          trigger: 'item',
          backgroundColor: 'rgba(10, 14, 22, 0.95)',
          borderColor: 'rgba(255,255,255,0.12)',
          textStyle: { color: '#e6edf5', fontSize: 11 },
          formatter: tooltipFormatter,
        },
        geo: {
          map: ZHEJIANG_MAP_NAME,
          roam: false,
          aspectScale: 0.92,
          itemStyle: {
            areaColor: 'rgba(59, 130, 246, 0.06)',
            borderColor: accent,
            borderWidth: 1,
            shadowColor: accent,
            shadowBlur: 12,
          },
          emphasis: {
            disabled: true,
          },
        },
        series: [
          // Pulsing markers for severe / major.
          {
            name: '重大风险',
            type: 'effectScatter',
            coordinateSystem: 'geo',
            data: grouped.severe.map((d) => ({ ...d, extra: labelOf('severe') })),
            symbolSize: ringSize.severe,
            rippleEffect: animate
              ? { period: 2.4, scale: 3.2, brushType: 'stroke' }
              : { period: 0, scale: 1 },
            itemStyle: { color: colorOf('severe') },
            zlevel: 4,
          },
          {
            name: '较大风险',
            type: 'effectScatter',
            coordinateSystem: 'geo',
            data: grouped.major.map((d) => ({ ...d, extra: labelOf('major') })),
            symbolSize: ringSize.major,
            rippleEffect: animate
              ? { period: 2.4, scale: 3.0, brushType: 'stroke' }
              : { period: 0, scale: 1 },
            itemStyle: { color: colorOf('major') },
            zlevel: 3,
          },
          // Static dots for normal / low — large fleets stay smooth.
          {
            name: '一般风险',
            type: 'scatter',
            coordinateSystem: 'geo',
            data: grouped.normal.map((d) => ({ ...d, extra: labelOf('normal') })),
            symbolSize: ringSize.normal,
            itemStyle: { color: colorOf('normal') },
            zlevel: 2,
          },
          {
            name: '低风险',
            type: 'scatter',
            coordinateSystem: 'geo',
            data: grouped.low.map((d) => ({ ...d, extra: labelOf('low') })),
            symbolSize: ringSize.low,
            itemStyle: { color: colorOf('low'), opacity: 0.85 },
            zlevel: 1,
          },
        ],
      },
      { notMerge: true },
    );
  }, [ready, points, levels, accent, animate]);

  // Resize on container resize. ECharts doesn't auto-resize without this.
  useEffect(() => {
    const el = containerRef.current;
    if (!el) return;
    const ro = new ResizeObserver(() => {
      chartRef.current?.resize();
    });
    ro.observe(el);
    return () => ro.disconnect();
  }, []);

  // Cleanup on unmount.
  useEffect(() => {
    return () => {
      chartRef.current?.dispose();
      chartRef.current = null;
    };
  }, []);

  return (
    <div style={{ width: '100%', height: '100%', position: 'relative' }}>
      <div ref={containerRef} style={{ width: '100%', height: '100%' }} />
      {error ? (
        <div
          style={{
            position: 'absolute',
            inset: 0,
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            color: 'var(--fg-dim)',
            fontSize: 11,
          }}
        >
          地图加载失败：{error}
        </div>
      ) : null}
    </div>
  );
}
