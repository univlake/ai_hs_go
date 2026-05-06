import { Card } from './Card';
import { ZhejiangMap } from './ZhejiangMap';
import { PieChart } from '../charts/PieChart';
import type { IndustryRiskSlice, RiskLevelInfo, RiskPoint } from '../types';
import type { Density, ThemeColors } from '../theme';

interface MapCardProps {
  riskPoints: RiskPoint[];
  riskLevels: RiskLevelInfo[];
  industryRisk: IndustryRiskSlice[];
  density: Density;
  theme: ThemeColors;
  animate: boolean;
}

export function MapCard({
  riskPoints,
  riskLevels,
  industryRisk,
  density,
  theme,
  animate,
}: MapCardProps) {
  return (
    <Card
      density={density}
      style={{ height: '100%', padding: 0, overflow: 'hidden' }}
      hover={false}
    >
      <div
        style={{
          position: 'relative',
          flex: 1,
          padding: 12,
          display: 'flex',
          flexDirection: 'column',
          minHeight: 0,
        }}
      >
        {/* 标题 + 4 级风险图例 */}
        <div style={{ display: 'flex', alignItems: 'center', gap: 8, marginBottom: 6 }}>
          <span
            style={{
              width: 3,
              height: 12,
              background: 'var(--primary)',
              borderRadius: 1,
            }}
          />
          <span style={{ fontSize: 13, fontWeight: 500 }}>风险点分布 · 浙江</span>
          <div style={{ flex: 1 }} />
          <div
            style={{
              display: 'flex',
              gap: 12,
              fontSize: 10,
              flexWrap: 'wrap',
              justifyContent: 'flex-end',
            }}
          >
            {riskLevels.map((l) => (
              <div key={l.key} style={{ display: 'flex', alignItems: 'center', gap: 4 }}>
                <span
                  style={{
                    position: 'relative',
                    width: 10,
                    height: 10,
                    flexShrink: 0,
                  }}
                >
                  <span
                    style={{
                      position: 'absolute',
                      inset: 0,
                      background: l.color,
                      borderRadius: '50%',
                      opacity: 0.5,
                    }}
                  />
                  <span
                    style={{
                      position: 'absolute',
                      inset: 2,
                      background: l.color,
                      borderRadius: '50%',
                    }}
                  />
                </span>
                <span style={{ opacity: 0.85, whiteSpace: 'nowrap' }}>{l.label}</span>
              </div>
            ))}
          </div>
        </div>

        {/* 地图区 */}
        <div
          style={{
            flex: 1,
            position: 'relative',
            minHeight: 0,
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
          }}
        >
          <ZhejiangMap
            points={riskPoints}
            levels={riskLevels}
            accent={theme.primary}
            animate={animate}
          />
        </div>

        {/* 行业风险占比 */}
        <div
          style={{
            marginTop: 8,
            paddingTop: 8,
            borderTop: '1px solid var(--border)',
            display: 'flex',
            alignItems: 'center',
            gap: 12,
          }}
        >
          <div style={{ flexShrink: 0 }}>
            <div style={{ fontSize: 10, opacity: 0.65, marginBottom: 2 }}>行业风险占比</div>
            <div
              style={{
                fontSize: 9,
                opacity: 0.45,
                fontFamily: 'var(--font-num)',
              }}
            >
              {riskPoints.length} risk points
            </div>
          </div>
          <PieChart data={industryRisk} size={70} />
          <div
            style={{
              display: 'grid',
              gridTemplateColumns: '1fr 1fr',
              gap: '2px 14px',
              flex: 1,
            }}
          >
            {industryRisk.map((r) => (
              <div
                key={r.name}
                style={{
                  fontSize: 10,
                  display: 'flex',
                  alignItems: 'center',
                  gap: 5,
                }}
              >
                <span
                  style={{
                    width: 7,
                    height: 7,
                    background: r.color,
                    borderRadius: 1,
                    flexShrink: 0,
                  }}
                />
                <span style={{ opacity: 0.85, flex: 1 }}>{r.name}</span>
                <span style={{ fontFamily: 'var(--font-num)', opacity: 0.95 }}>{r.value}</span>
                <span
                  style={{
                    fontFamily: 'var(--font-num)',
                    opacity: 0.55,
                    width: 32,
                    textAlign: 'right',
                  }}
                >
                  {r.percent}%
                </span>
              </div>
            ))}
          </div>
        </div>
      </div>
    </Card>
  );
}
