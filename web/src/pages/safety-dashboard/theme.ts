// Theme tokens for the safety dashboard. The design ships 4 themes × 2 modes;
// for now we expose only `tech blue + dark` since that's the canonical big-screen
// look. The data shape is preserved so a future tweaks panel can flip themes
// without touching components.

export type ThemeName = 'blue' | 'cyan' | 'violet' | 'emerald';

export interface ThemeColors {
  primary: string;
  accent: string;
  warn: string;
  bg: string;
  card: string;
}

const THEMES_DARK: Record<ThemeName, ThemeColors> = {
  blue: { primary: '#3b82f6', accent: '#60a5fa', warn: '#f59e0b', bg: '#0a0e15', card: '#0f141d' },
  cyan: { primary: '#06b6d4', accent: '#22d3ee', warn: '#f59e0b', bg: '#08111a', card: '#0d1822' },
  violet: { primary: '#8b5cf6', accent: '#a78bfa', warn: '#f59e0b', bg: '#0d0a1a', card: '#141022' },
  emerald: { primary: '#10b981', accent: '#34d399', warn: '#f59e0b', bg: '#08120f', card: '#0f1a16' },
};

const THEMES_LIGHT: Record<ThemeName, ThemeColors> = {
  blue: { primary: '#2563eb', accent: '#3b82f6', warn: '#d97706', bg: '#f5f7fb', card: '#ffffff' },
  cyan: { primary: '#0891b2', accent: '#06b6d4', warn: '#d97706', bg: '#f3f8fa', card: '#ffffff' },
  violet: { primary: '#7c3aed', accent: '#8b5cf6', warn: '#d97706', bg: '#f7f5fb', card: '#ffffff' },
  emerald: { primary: '#059669', accent: '#10b981', warn: '#d97706', bg: '#f4faf7', card: '#ffffff' },
};

export function getTheme(name: ThemeName, dark: boolean): ThemeColors {
  return (dark ? THEMES_DARK : THEMES_LIGHT)[name];
}

function hexToRgba(hex: string, alpha: number): string {
  const r = parseInt(hex.slice(1, 3), 16);
  const g = parseInt(hex.slice(3, 5), 16);
  const b = parseInt(hex.slice(5, 7), 16);
  return `rgba(${r},${g},${b},${alpha})`;
}

// CSS variables consumed by every component via `var(--xxx)`. Returned as a
// React.CSSProperties so the dashboard root can spread them with the `style`
// prop and let components stay theme-agnostic.
export function buildThemeVars(name: ThemeName, dark: boolean): React.CSSProperties {
  const t = getTheme(name, dark);
  const fg = dark ? '#e6edf5' : '#0f172a';
  const fgDim = dark ? 'rgba(230,237,245,0.6)' : 'rgba(15,23,42,0.6)';
  const border = dark ? 'rgba(255,255,255,0.07)' : 'rgba(15,23,42,0.08)';
  const card2 = dark ? 'rgba(255,255,255,0.03)' : 'rgba(15,23,42,0.03)';
  return {
    ['--bg' as string]: t.bg,
    ['--card' as string]: t.card,
    ['--card-2' as string]: card2,
    ['--fg' as string]: fg,
    ['--fg-dim' as string]: fgDim,
    ['--border' as string]: border,
    ['--primary' as string]: t.primary,
    ['--accent' as string]: t.accent,
    ['--warn' as string]: t.warn,
    ['--primary-30' as string]: hexToRgba(t.primary, 0.3),
    ['--primary-15' as string]: hexToRgba(t.primary, 0.15),
    ['--font-zh' as string]:
      "'Noto Sans SC', -apple-system, BlinkMacSystemFont, 'PingFang SC', 'Microsoft YaHei', sans-serif",
    ['--font-num' as string]: "'DM Mono', 'SF Mono', Menlo, monospace",
  };
}

export type Density = 'compact' | 'cozy';
