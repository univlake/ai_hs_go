import ExpertInspectionGroupView from '@/pages/risk-control/expert-inspection/Index';

// 生产视图：仅展示「本企业相关」的检查计划，复用集团端组件传 mode='production'。
// 跟 sidebar.tsx 里的 /risk/expert-inspection 入口对应。
export default function RiskExpertInspectionIndex() {
    return <ExpertInspectionGroupView mode="production" />;
}
