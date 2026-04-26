// Central registry of backend API URLs. Replaces wayfinder-generated helpers
// (`@/actions/*`, `@/routes/*`).
//
// IMPORTANT: paths are relative to the axios baseURL (`/api/v1`, set in
// lib/api.ts). Do NOT include the `/api/v1/` prefix here — it gets prepended
// by axios automatically.

export const API = {
  publicKey: () => '/public-key',
  login: () => '/login',
  logout: () => '/logout',
  userInfo: () => '/user-info',
  updateUserInfo: () => '/user-info',
  usernameAvailability: () => '/username-availability',
  firstLoginSetup: () => '/first-login-setup',
  resetPassword: () => '/reset-password',

  companiesPublic: () => '/companies-public',
  companies: () => '/companies',
  companyById: (id: number | string) => `/companies/${id}`,
  companiesTree: () => '/companies-tree',
  companiesForest: () => '/companies-forest',
  companyToggleStatus: (id: number | string) => `/companies/${id}/toggle-status`,
  companyReview: (id: number | string) => `/companies/${id}/review`,

  tenants: () => '/tenants',
  tenantToggleStatus: (id: number | string) => `/tenants/${id}/toggle-status`,

  // Platform-admin impersonation (path A — scope header switch + audit row)
  impersonate: () => '/platform/impersonate',
  impersonateExit: () => '/platform/impersonate/exit',

  users: () => '/users',
  userById: (id: number | string) => `/users/${id}`,
  userRoles: (id: number | string) => `/users/${id}/roles`,

  roles: () => '/roles',
  roleById: (id: number | string) => `/roles/${id}`,
  rolePermissions: (id: number | string) => `/roles/${id}/permissions`,
  permissions: () => '/permissions',

  specialTypes: () => '/special-types',

  specialEquipment: () => '/special-equipment',
  specialEquipmentById: (id: number | string) => `/special-equipment/${id}`,
  specialEquipmentTemplate: () => '/special-equipment/template',
  specialEquipmentExport: () => '/special-equipment/export',
  specialEquipmentImport: () => '/special-equipment/import',

  dangerousEquipment: () => '/dangerous-equipment',
  dangerousEquipmentById: (id: number | string) => `/dangerous-equipment/${id}`,
  dangerousEquipmentTemplate: () => '/dangerous-equipment/template',
  dangerousEquipmentExport: () => '/dangerous-equipment/export',
  dangerousEquipmentImport: () => '/dangerous-equipment/import',

  hazardousChemicals: () => '/hazardous-chemicals',
  hazardousChemicalById: (id: number | string) => `/hazardous-chemicals/${id}`,

  uploadOSS: {
    init: () => '/upload-oss/init',
    signPart: () => '/upload-oss/sign-part',
    complete: () => '/upload-oss/complete',
    abort: () => '/upload-oss/abort',
    signed: () => '/upload-oss/signed',
  },
} as const;
