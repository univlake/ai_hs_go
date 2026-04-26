import { lazy, Suspense } from 'react';
import { createBrowserRouter, Navigate, Outlet, useLocation } from 'react-router-dom';

import AppLayout from '@/layouts/AppLayout';
import { useAuth } from '@/hooks/use-auth';

// Lazy-load every page bundle. Vite's `import.meta.glob` keeps Inertia-era
// page resolution working: the original conventions (`pages/auth/Login`,
// `pages/dashboard/Index`, etc.) are preserved.
const pageModules = import.meta.glob('@/pages/**/*.tsx');

function lazyPage(path: string) {
  const importer = pageModules[`/src/pages/${path}.tsx`];
  if (!importer) {
    throw new Error(`page not found: ${path}`);
  }
  return lazy(() => importer() as Promise<{ default: React.ComponentType }>);
}

const Welcome = lazy(
  () =>
    (pageModules['/src/pages/welcome.tsx']?.() as Promise<{
      default: React.ComponentType;
    }>) ?? Promise.reject(new Error('welcome page missing')),
);

function AuthGate() {
  const { user, isLoading, isAuthenticated } = useAuth();
  const location = useLocation();
  if (isLoading) {
    return <div className="p-8 text-sm text-[var(--color-text-muted)]">验证登录状态…</div>;
  }
  if (!isAuthenticated || !user) {
    return <Navigate to="/login" replace state={{ from: location }} />;
  }
  if (user.must_change_credentials && location.pathname !== '/first-login-setup') {
    return <Navigate to="/first-login-setup" replace />;
  }
  return <Outlet />;
}

function LoadingShell({ children }: { children: React.ReactNode }) {
  return <Suspense fallback={<div className="p-8 text-sm">加载中…</div>}>{children}</Suspense>;
}

// Lazy page registry — extend as more Laravel pages are migrated. Page paths
// are relative to web/src/pages/ without the .tsx suffix.
const Login = lazyPage('auth/Login');
const Register = lazyPage('auth/Register');
const Dashboard = lazyPage('dashboard/Dashboard');
const FirstLoginSetup = lazyPage('first-login-setup/Index');
const ChangePassword = lazyPage('change-password/Index');
const Profile = lazyPage('account/Index');
const PlatformTenants = lazyPage('platform/Tenants');
const PlatformForest = lazyPage('platform/Forest');
const PlatformPermissions = lazyPage('platform/Permissions');
const SystemUsers = lazyPage('system/users/Index');
const SystemRoles = lazyPage('system/roles/Index');
const SafetySpecialEquipment = lazyPage('safety/special-equipment/Index');
const SafetyDangerousEquipment = lazyPage('safety/dangerous-equipment/Index');
const SafetyHazardousChemicals = lazyPage('safety/hazardous-chemicals/Index');
const SafetyMajorHazards = lazyPage('safety/major-hazards/Index');
const SafetyFireSafeties = lazyPage('safety/fire-safeties/Index');
const HealthFiles = lazyPage('health/health-files/Index');
const HealthOccupationalHazards = lazyPage('health/occupational-hazards/Index');
const HealthCheckups = lazyPage('health/health-checkups/Index');
const HealthHygieneEvaluations = lazyPage('health/hygiene-evaluations/Index');
const HealthDiseaseReport = lazyPage('health/disease-report/Index');
const FireEquipment = lazyPage('fire/equipment/Index');
const FireFacilities = lazyPage('fire/facilities/Index');
const EnterpriseCurrent = lazyPage('enterprise/current/Index');
const EnterpriseSubsidiaries = lazyPage('enterprise/subsidiaries/Index');
const EnterpriseApproval = lazyPage('enterprise/approval/Index');
const EnterpriseGrouping = lazyPage('enterprise/grouping/Index');
const SystemEnterpriseTypes = lazyPage('system/enterprise-types/Index');
const SystemOperationalTypes = lazyPage('system/operational-types/Index');
const SystemRecordTypes = lazyPage('system/record-types/Index');
const SystemProcessTypes = lazyPage('system/process-types/Index');
const SystemSpecialTypes = lazyPage('system/special-types/Index');
const SystemSpaceWorkTypes = lazyPage('system/space-work-types/Index');
const SystemSafetyManagementTypes = lazyPage('system/safety-management-types/Index');
const SystemDangerousCraftTypes = lazyPage('system/dangerous-craft-types/Index');
const SystemHazards = lazyPage('system/hazards/Index');
const SystemRegulationLibraryTypes = lazyPage('system/regulation-library-types/Index');
const KnowledgeChecklists = lazyPage('knowledge/checklists/Index');
const SafetyArchOrgStructure = lazyPage('safety-archives/org-structure/Index');
const SafetyArchQualifications = lazyPage('safety-archives/qualifications/Index');
const SafetyArchInvestment = lazyPage('safety-archives/safety-investment/Index');
const SafetyPersonnels = lazyPage('safety/personnels/Index');
const SafetyBuildings = lazyPage('safety/buildings/Index');
const ProcessKeyHazardous = lazyPage('process/key-hazardous/Index');
const ProcessHazardous = lazyPage('process/hazardous/Index');
// 有限空间辨识页面文件实际位于 process/confined-space/，但 Sidebar 用
// /space-work/space-works 作为 href（permission space-works.list），这里挂载
// 同一组件而不动 Sidebar，避免误碰其他不在本次范围的链接。
const ProcessConfinedSpace = lazyPage('process/confined-space/Index');
const ThirdPartyDesignUnits = lazyPage('third-party/design-units/Index');
const ThirdPartyConstructionUnits = lazyPage('third-party/construction-units/Index');
const ThirdPartyTechnicalServices = lazyPage('third-party/technical-services/Index');
const MeetingInternalMeetings = lazyPage('meeting/internal-meetings/Index');
const TrainingEducation = lazyPage('training/education/Index');
const Placeholder = lazyPage('_placeholder');

export const router = createBrowserRouter([
  {
    path: '/',
    element: (
      <LoadingShell>
        <Welcome />
      </LoadingShell>
    ),
  },
  {
    path: '/login',
    element: (
      <LoadingShell>
        <Login />
      </LoadingShell>
    ),
  },
  {
    path: '/register',
    element: (
      <LoadingShell>
        <Register />
      </LoadingShell>
    ),
  },
  {
    element: <AuthGate />,
    children: [
      {
        path: '/first-login-setup',
        element: (
          <LoadingShell>
            <FirstLoginSetup />
          </LoadingShell>
        ),
      },
      {
        element: <AppLayout />,
        children: [
          { path: '/dashboard', element: <LoadingShell><Dashboard /></LoadingShell> },
          { path: '/profile', element: <LoadingShell><Profile /></LoadingShell> },
          { path: '/change-password', element: <LoadingShell><ChangePassword /></LoadingShell> },
          { path: '/platform/tenants', element: <LoadingShell><PlatformTenants /></LoadingShell> },
          { path: '/platform/forest', element: <LoadingShell><PlatformForest /></LoadingShell> },
          { path: '/platform/permissions', element: <LoadingShell><PlatformPermissions /></LoadingShell> },
          { path: '/system/users', element: <LoadingShell><SystemUsers /></LoadingShell> },
          { path: '/system/roles', element: <LoadingShell><SystemRoles /></LoadingShell> },
          { path: '/safety/special-equipment', element: <LoadingShell><SafetySpecialEquipment /></LoadingShell> },
          { path: '/safety/dangerous-equipment', element: <LoadingShell><SafetyDangerousEquipment /></LoadingShell> },
          { path: '/safety/hazardous-chemicals', element: <LoadingShell><SafetyHazardousChemicals /></LoadingShell> },
          { path: '/safety/major-hazards', element: <LoadingShell><SafetyMajorHazards /></LoadingShell> },
          { path: '/safety/fire-safeties', element: <LoadingShell><SafetyFireSafeties /></LoadingShell> },
          { path: '/health/health-files', element: <LoadingShell><HealthFiles /></LoadingShell> },
          { path: '/health/occupational-hazards', element: <LoadingShell><HealthOccupationalHazards /></LoadingShell> },
          { path: '/health/health-checkups', element: <LoadingShell><HealthCheckups /></LoadingShell> },
          { path: '/health/hygiene-evaluations', element: <LoadingShell><HealthHygieneEvaluations /></LoadingShell> },
          { path: '/health/disease-report', element: <LoadingShell><HealthDiseaseReport /></LoadingShell> },
          { path: '/fire/equipment', element: <LoadingShell><FireEquipment /></LoadingShell> },
          { path: '/fire/facilities', element: <LoadingShell><FireFacilities /></LoadingShell> },
          { path: '/enterprise/current', element: <LoadingShell><EnterpriseCurrent /></LoadingShell> },
          { path: '/enterprise/subsidiaries', element: <LoadingShell><EnterpriseSubsidiaries /></LoadingShell> },
          { path: '/enterprise/approval', element: <LoadingShell><EnterpriseApproval /></LoadingShell> },
          { path: '/enterprise/grouping', element: <LoadingShell><EnterpriseGrouping /></LoadingShell> },
          { path: '/system/enterprise-types', element: <LoadingShell><SystemEnterpriseTypes /></LoadingShell> },
          { path: '/system/operational-types', element: <LoadingShell><SystemOperationalTypes /></LoadingShell> },
          { path: '/system/record-types', element: <LoadingShell><SystemRecordTypes /></LoadingShell> },
          { path: '/system/process-types', element: <LoadingShell><SystemProcessTypes /></LoadingShell> },
          { path: '/system/special-types', element: <LoadingShell><SystemSpecialTypes /></LoadingShell> },
          { path: '/system/space-work-types', element: <LoadingShell><SystemSpaceWorkTypes /></LoadingShell> },
          { path: '/system/safety-management-types', element: <LoadingShell><SystemSafetyManagementTypes /></LoadingShell> },
          { path: '/system/dangerous-craft-types', element: <LoadingShell><SystemDangerousCraftTypes /></LoadingShell> },
          { path: '/system/hazards', element: <LoadingShell><SystemHazards /></LoadingShell> },
          { path: '/system/regulation-library-types', element: <LoadingShell><SystemRegulationLibraryTypes /></LoadingShell> },
          { path: '/knowledge/checklists', element: <LoadingShell><KnowledgeChecklists /></LoadingShell> },
          { path: '/safety-archives/org-structure', element: <LoadingShell><SafetyArchOrgStructure /></LoadingShell> },
          { path: '/safety-archives/qualifications', element: <LoadingShell><SafetyArchQualifications /></LoadingShell> },
          { path: '/safety-archives/safety-investment', element: <LoadingShell><SafetyArchInvestment /></LoadingShell> },
          { path: '/safety/personnels', element: <LoadingShell><SafetyPersonnels /></LoadingShell> },
          { path: '/safety/buildings', element: <LoadingShell><SafetyBuildings /></LoadingShell> },
          { path: '/process/key-hazardous', element: <LoadingShell><ProcessKeyHazardous /></LoadingShell> },
          { path: '/process/hazardous', element: <LoadingShell><ProcessHazardous /></LoadingShell> },
          { path: '/space-work/space-works', element: <LoadingShell><ProcessConfinedSpace /></LoadingShell> },
          { path: '/third-party/design-units', element: <LoadingShell><ThirdPartyDesignUnits /></LoadingShell> },
          { path: '/third-party/construction-units', element: <LoadingShell><ThirdPartyConstructionUnits /></LoadingShell> },
          { path: '/third-party/technical-services', element: <LoadingShell><ThirdPartyTechnicalServices /></LoadingShell> },
          { path: '/meeting/internal-meetings', element: <LoadingShell><MeetingInternalMeetings /></LoadingShell> },
          { path: '/training/education', element: <LoadingShell><TrainingEducation /></LoadingShell> },
          // Sidebar 旧 href 兼容（之前指向 /equipment/major-hazardous）
          { path: '/equipment/major-hazardous', element: <LoadingShell><SafetyDangerousEquipment /></LoadingShell> },
          // Catch-all inside AppLayout: shows a friendly "尚未迁移" page
          // instead of bouncing to /dashboard so testers can see which
          // sidebar links still need wiring.
          { path: '*', element: <LoadingShell><Placeholder /></LoadingShell> },
        ],
      },
    ],
  },
]);
