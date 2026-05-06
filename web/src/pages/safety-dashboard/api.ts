import api from '@/lib/api';
import type { DashboardData } from './types';

interface Envelope<T> {
  status: 'success' | 'error';
  code: number;
  data?: T;
  message?: string;
}

// Fetches the live safety-dashboard payload from the backend. Scope (which
// company subtree) is set via the X-Scope-Company-Id header that the shared
// axios interceptor already injects from useScope().
export async function fetchDashboardOverview(signal?: AbortSignal): Promise<DashboardData> {
  const res = await api.get<Envelope<DashboardData>>('/safety-dashboard/overview', { signal });
  const data = res.data?.data;
  if (!data) {
    throw new Error(res.data?.message ?? '后端未返回数据');
  }
  return data;
}
