import { useEffect, useMemo, useState } from 'react';
import { fetchDashboardOverview } from './api';
import { mockDashboardData } from './mockData';
import { buildThemeVars, getTheme, type Density, type ThemeName } from './theme';
import type { DashboardData, DrillKpi } from './types';
import { TopBar } from './components/TopBar';
import { StatusBar } from './components/StatusBar';
import { PersonnelCard } from './components/PersonnelCard';
import { TrainingCard } from './components/TrainingCard';
import { InvestmentCard } from './components/InvestmentCard';
import { HazardCard } from './components/HazardCard';
import { MapCard } from './components/MapCard';
import { StandardCard } from './components/StandardCard';
import { DrillsCard } from './components/DrillsCard';
import { HealthCard } from './components/HealthCard';
import { KpiPill } from './components/KpiPill';
import './styles.css';

// Polling cadence for the dashboard. 60s is the README's recommendation —
// long enough that 17 cards' worth of aggregation queries don't pile up,
// short enough that the "数据更新" timestamp stays fresh.
const POLL_INTERVAL_MS = 60_000;

// 杭实集团安全风险看板 — Innovative variant from the design handoff.
//
// Layout:
//   TopBar (56px clock + title + safe-days KPI)
//   Top KPI strip — 6 pills across (员工总数 / 持续安全生产 / 标准化等级 /
//                                   安全演练 / 消防演练 / 风险点)
//   Mid grid (340px 1fr 340px × 2 rows): personnel + map(2 rows) + standard
//                                         training + (map) + health
//   Bottom grid (1fr × 3, height 200px): investment / hazard / drills
//   StatusBar (32px alert ticker)
//
// Data flow: poll /safety-dashboard/overview every 60s; fall back to mock
// data only when the very first request fails so the dashboard isn't blank
// during a transient API hiccup. Once a real payload lands, never fall back.
export default function SafetyDashboardIndex() {
  // Reserved for a future tweaks panel; not exposed in UI yet.
  const [themeName] = useState<ThemeName>('blue');
  const [dark] = useState(true);
  const [density] = useState<Density>('compact');
  const [animate] = useState(true);

  const [data, setData] = useState<DashboardData | null>(null);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const ctrl = new AbortController();
    let timer: ReturnType<typeof setTimeout> | null = null;

    const tick = async () => {
      try {
        const next = await fetchDashboardOverview(ctrl.signal);
        setData(next);
        setError(null);
      } catch (e) {
        if (ctrl.signal.aborted) return;
        const msg = e instanceof Error ? e.message : String(e);
        setError(msg);
      } finally {
        if (!ctrl.signal.aborted) {
          timer = setTimeout(tick, POLL_INTERVAL_MS);
        }
      }
    };
    tick();

    return () => {
      ctrl.abort();
      if (timer) clearTimeout(timer);
    };
  }, []);

  const view = data ?? mockDashboardData;
  const theme = useMemo(() => getTheme(themeName, dark), [themeName, dark]);
  const themeVars = useMemo(() => buildThemeVars(themeName, dark), [themeName, dark]);

  // Pull each drill KPI by its label so we don't depend on array order; the
  // backend returns them in a fixed sequence today but small reorders
  // shouldn't break the top strip.
  const drillByLabel = useMemo(() => {
    const m: Record<string, DrillKpi> = {};
    for (const d of view.drills) m[d.label] = d;
    return m;
  }, [view.drills]);

  const safetyDrillSessions = drillByLabel['安全演练次数']?.value ?? 0;
  const safetyDrillPeople = drillByLabel['安全演练人数']?.value ?? 0;
  const fireDrillSessions = drillByLabel['消防演练次数']?.value ?? 0;
  const fireDrillPeople = drillByLabel['消防演练人数']?.value ?? 0;

  return (
    <div
      data-safety-dashboard=""
      style={{
        ...themeVars,
        background: 'var(--bg)',
        color: 'var(--fg)',
        fontFamily: 'var(--font-zh)',
        display: 'flex',
        flexDirection: 'column',
      }}
    >
      <TopBar title={view.meta.title} safeDays={view.meta.safeDays} />

      {/* 顶部 KPI 长条 — 6 across */}
      <div
        style={{
          padding: '10px 12px 0',
          display: 'grid',
          gridTemplateColumns: 'repeat(6, 1fr)',
          gap: 10,
        }}
      >
        <KpiPill
          label="员工总数"
          value={view.personnel.total}
          unit="人"
          color="var(--primary)"
          animate={animate}
        />
        <KpiPill
          label="持续安全生产"
          value={view.meta.safeDays}
          unit="天"
          color="var(--accent)"
          animate={animate}
        />
        <KpiPill
          label="标准化等级"
          value={view.standardLevel.total}
          unit={`项 · ${view.standardLevel.current || '—'}`}
          color="var(--primary)"
          animate={animate}
        />
        <KpiPill
          label="安全演练"
          value={safetyDrillSessions}
          unit={`次 · ${safetyDrillPeople}人`}
          color="var(--accent)"
          animate={animate}
        />
        <KpiPill
          label="消防演练"
          value={fireDrillSessions}
          unit={`次 · ${fireDrillPeople}人`}
          color="var(--warn)"
          animate={animate}
        />
        <KpiPill
          label="风险点"
          value={view.riskPoints.length}
          unit="个 · 浙江"
          color="var(--primary)"
          animate={animate}
        />
      </div>

      {/* 中央 grid — personnel | map(2 rows) | standard / training | (map) | health */}
      <div
        style={{
          flex: 1,
          padding: 12,
          display: 'grid',
          gridTemplateColumns: '340px 1fr 340px',
          gridTemplateRows: '1fr 1fr',
          gap: 10,
          minHeight: 0,
        }}
      >
        <PersonnelCard data={view.personnel} density={density} />
        <div style={{ gridRow: 'span 2', minHeight: 0 }}>
          <MapCard
            riskPoints={view.riskPoints}
            riskLevels={view.riskLevels}
            industryRisk={view.industryRisk}
            density={density}
            theme={theme}
            animate={animate}
          />
        </div>
        <StandardCard data={view.standardLevel} density={density} />

        <TrainingCard data={view.training} density={density} animate={animate} />
        <HealthCard data={view.occupationalHealth} density={density} animate={animate} />
      </div>

      {/* 底部 grid — investment / hazard / drills */}
      <div
        style={{
          padding: '0 12px 12px',
          display: 'grid',
          gridTemplateColumns: '1fr 1fr 1fr',
          gap: 10,
          height: 200,
        }}
      >
        <InvestmentCard data={view.investment} density={density} animate={animate} />
        <HazardCard
          year={view.hazardYear}
          month={view.hazardMonth}
          density={density}
          animate={animate}
        />
        <DrillsCard data={view.drills} density={density} animate={animate} />
      </div>

      <StatusBar updatedAt={view.meta.updatedAt} alerts={view.alerts} />

      {/* Connection error toast — keeps the dashboard usable while flagging
          the stale state. Hidden once any fetch succeeds. */}
      {error && !data ? (
        <div
          style={{
            position: 'absolute',
            top: 70,
            right: 16,
            padding: '6px 12px',
            background: 'var(--card)',
            border: '1px solid var(--warn)',
            borderRadius: 6,
            fontSize: 11,
            color: 'var(--warn)',
            zIndex: 100,
          }}
        >
          后端连接失败，显示示例数据
        </div>
      ) : null}
    </div>
  );
}
