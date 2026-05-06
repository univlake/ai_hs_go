import type { DashboardData } from './types';

// Mock data ported from design_handoff_safety_dashboard/dashboard-data.js.
// All KPI top-line numbers preserved from the original dashboard screenshot;
// monthly breakdowns and risk-point coords are the design team's reconstruction.
//
// Replace this with a real fetch when the backend endpoint lands; the contract
// is in types.ts so that swap is mechanical.
export const mockDashboardData: DashboardData = {
  meta: {
    title: '杭实集团安全风险看板',
    safeDays: 1222,
    updatedAt: '2026-05-06 16:12:57',
  },

  personnel: {
    total: 26690,
    composition: [
      { label: '职工', value: 20367, color: 'primary' },
      { label: '管理人员', value: 3170, color: 'accent' },
      { label: '劳务派遣用工', value: 3153, color: 'warn' },
    ],
    professionals: [
      { label: '注册安全工程师', value: 39, percent: 15.9 },
      { label: '安全工程专业', value: 18, percent: 7.0 },
      { label: '注册消防工程师', value: 2, percent: 0.8 },
      { label: '其他', value: 198, percent: 77.0 },
    ],
  },

  training: [
    { month: '25/06', sessions: 128, people: 7540 },
    { month: '25/07', sessions: 101, people: 7124 },
    { month: '25/08', sessions: 109, people: 6667 },
    { month: '25/09', sessions: 94, people: 2674 },
    { month: '25/10', sessions: 97, people: 3871 },
    { month: '25/11', sessions: 134, people: 6128 },
    { month: '25/12', sessions: 211, people: 3331 },
    { month: '26/01', sessions: 60, people: 2998 },
    { month: '26/02', sessions: 48, people: 4250 },
    { month: '26/03', sessions: 84, people: 4487 },
    { month: '26/04', sessions: 139, people: 3956 },
    { month: '26/05', sessions: 92, people: 3402 },
  ],

  investment: [
    { month: '25/06', value: 309 },
    { month: '25/07', value: 286 },
    { month: '25/08', value: 30308 },
    { month: '25/09', value: 402 },
    { month: '25/10', value: 1692 },
    { month: '25/11', value: 333 },
    { month: '25/12', value: 457 },
    { month: '26/01', value: 127 },
    { month: '26/02', value: 280 },
    { month: '26/03', value: 196 },
    { month: '26/04', value: 9 },
    { month: '26/05', value: 145 },
  ],

  hazardYear: [
    { industry: '化工行业', found: 612, fixed: 598, rate: 97.7 },
    { industry: '机械行业', found: 88, fixed: 84, rate: 95.5 },
    { industry: '商贸租赁行业', found: 142, fixed: 138, rate: 97.2 },
    { industry: '其他行业', found: 305, fixed: 290, rate: 95.1 },
  ],

  hazardMonth: [
    { industry: '化工行业', found: 105, fixed: 102, rate: 97.1 },
    { industry: '机械行业', found: 22, fixed: 21, rate: 95.5 },
    { industry: '商贸租赁行业', found: 38, fixed: 36, rate: 94.7 },
    { industry: '其他行业', found: 168, fixed: 158, rate: 94.0 },
  ],

  // Real Zhejiang city centroids — ECharts geo projects these directly.
  riskPoints: [
    { city: '杭州', lng: 120.16, lat: 30.29, level: 'low', name: '杭州主厂区' },
    { city: '杭州', lng: 120.27, lat: 30.16, level: 'low', name: '萧山仓储' },
    { city: '杭州', lng: 120.30, lat: 30.42, level: 'low', name: '余杭物流' },
    { city: '杭州', lng: 119.96, lat: 30.05, level: 'low', name: '富阳分厂' },
    { city: '湖州', lng: 120.09, lat: 30.89, level: 'normal', name: '湖州化工' },
    { city: '嘉兴', lng: 120.76, lat: 30.77, level: 'normal', name: '嘉兴贸易' },
    { city: '嘉兴', lng: 120.69, lat: 30.51, level: 'low', name: '海宁仓库' },
    { city: '绍兴', lng: 120.58, lat: 30.03, level: 'normal', name: '绍兴机械' },
    { city: '绍兴', lng: 120.87, lat: 30.03, level: 'low', name: '上虞园区' },
    { city: '宁波', lng: 121.62, lat: 29.87, level: 'low', name: '宁波港务' },
    { city: '宁波', lng: 121.85, lat: 29.90, level: 'normal', name: '北仑储运' },
    { city: '舟山', lng: 122.21, lat: 30.00, level: 'low', name: '舟山基地' },
    { city: '台州', lng: 121.42, lat: 28.66, level: 'normal', name: '台州化工' },
    { city: '温州', lng: 120.69, lat: 28.00, level: 'low', name: '温州瓯江' },
    { city: '金华', lng: 119.65, lat: 29.08, level: 'low', name: '金华义乌' },
    { city: '丽水', lng: 119.92, lat: 28.46, level: 'low', name: '丽水山区' },
    { city: '衢州', lng: 118.87, lat: 28.94, level: 'low', name: '衢州工厂' },
  ],

  riskLevels: [
    { key: 'severe', label: '重大风险', color: '#ef4444' },
    { key: 'major', label: '较大风险', color: '#f97316' },
    { key: 'normal', label: '一般风险', color: '#eab308' },
    { key: 'low', label: '低风险', color: '#3b82f6' },
  ],

  industryRisk: [
    { name: '其他', value: 31, percent: 60, color: '#3b82f6' },
    { name: '机械', value: 12, percent: 23, color: '#a855f7' },
    { name: '化工', value: 6, percent: 12, color: '#ef4444' },
    { name: '商贸租赁', value: 3, percent: 6, color: '#64748b' },
  ],

  standardLevel: {
    total: 14,
    current: '三级',
    breakdown: [
      { label: '一级', value: 2, color: '#eab308' },
      { label: '二级', value: 3, color: '#3b82f6' },
      { label: '三级', value: 9, color: '#a855f7' },
    ],
  },

  drills: [
    { label: '安全演练人数', value: 413, suffix: '人', color: 'primary' },
    { label: '安全演练次数', value: 25, suffix: '次', color: 'primary' },
    { label: '消防演练人数', value: 2313, suffix: '人', color: 'warn' },
    { label: '消防演练次数', value: 261, suffix: '次', color: 'warn' },
  ],

  occupationalHealth: [
    { hazard: '高温焊电', people: 8 },
    { hazard: '使用甲乙类溶剂', people: 2 },
    { hazard: '噪音', people: 2 },
    { hazard: '冲击工艺', people: 2 },
    { hazard: '焊接', people: 3 },
    { hazard: '堆焊', people: 3 },
    { hazard: '其他工艺', people: 2 },
  ],

  alerts: [
    { time: '16:08', level: 'normal', text: '湖州化工区温度传感器读数偏高，已派员核查' },
    { time: '15:42', level: 'low', text: '杭州主厂区完成季度消防演练，参演 86 人' },
    { time: '15:21', level: 'major', text: '台州化工 2 号罐区压力波动，已联动应急响应' },
    { time: '14:55', level: 'low', text: '宁波港务区夜班交接班巡检完成，无异常' },
    { time: '14:33', level: 'normal', text: '绍兴机械车间 PM2.5 短时升高，启动通风' },
    { time: '13:50', level: 'low', text: '嘉兴贸易仓库新增智能烟感 12 台' },
    { time: '13:12', level: 'low', text: '舟山基地完成年度安全标准化二级评审' },
  ],
};
