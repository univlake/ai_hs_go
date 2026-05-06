// Data contract for the 集团安全风险看板. Mirrors `dashboard-data.js` from the
// design handoff verbatim — keep field names stable so a future backend
// endpoint can serve this shape with minimal mapping.

export type RiskLevel = 'severe' | 'major' | 'normal' | 'low';

export type ColorRole = 'primary' | 'accent' | 'warn';

export interface PersonnelComposition {
  label: string;
  value: number;
  color: ColorRole;
}

export interface PersonnelProfessional {
  label: string;
  value: number;
  percent: number;
}

export interface MonthlyTraining {
  month: string;
  sessions: number;
  people: number;
}

export interface MonthlyInvestment {
  month: string;
  value: number;
}

export interface HazardRow {
  industry: string;
  found: number;
  fixed: number;
  rate: number;
}

export interface RiskPoint {
  city: string;
  lng: number;
  lat: number;
  level: RiskLevel;
  name: string;
}

export interface RiskLevelInfo {
  key: RiskLevel;
  label: string;
  color: string;
}

export interface IndustryRiskSlice {
  name: string;
  value: number;
  percent: number;
  color: string;
}

export interface StandardLevelBreakdown {
  label: string;
  value: number;
  color: string;
}

export interface StandardLevel {
  total: number;
  current: string;
  breakdown: StandardLevelBreakdown[];
}

export interface DrillKpi {
  label: string;
  value: number;
  suffix: string;
  color: 'primary' | 'warn';
}

export interface OccupationalHealthRow {
  hazard: string;
  people: number;
}

export type AlertLevel = 'major' | 'normal' | 'low';

export interface AlertEntry {
  time: string;
  level: AlertLevel;
  text: string;
}

export interface DashboardData {
  meta: {
    title: string;
    safeDays: number;
    updatedAt: string;
  };
  personnel: {
    total: number;
    composition: PersonnelComposition[];
    professionals: PersonnelProfessional[];
  };
  training: MonthlyTraining[];
  investment: MonthlyInvestment[];
  hazardYear: HazardRow[];
  hazardMonth: HazardRow[];
  riskPoints: RiskPoint[];
  riskLevels: RiskLevelInfo[];
  industryRisk: IndustryRiskSlice[];
  standardLevel: StandardLevel;
  drills: DrillKpi[];
  occupationalHealth: OccupationalHealthRow[];
  alerts: AlertEntry[];
}
