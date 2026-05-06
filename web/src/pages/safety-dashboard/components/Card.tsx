import { useState, type CSSProperties, type ReactNode } from 'react';
import type { Density } from '../theme';

interface CardProps {
  title?: string;
  action?: ReactNode;
  children?: ReactNode;
  style?: CSSProperties;
  density?: Density;
  /** When false, the card never grows the hover border / shadow (e.g. the map). */
  hover?: boolean;
}

// Generic "section" card. The hover state is local, not tied to CSS :hover,
// because the design wants a soft shadow lift that's hard to express with
// pseudo-classes when the surrounding background is also a CSS variable.
export function Card({
  title,
  action,
  children,
  style,
  density = 'cozy',
  hover = true,
}: CardProps) {
  const [hov, setHov] = useState(false);
  const pad = density === 'compact' ? 10 : 14;

  return (
    <div
      onMouseEnter={() => setHov(true)}
      onMouseLeave={() => setHov(false)}
      style={{
        position: 'relative',
        background: 'var(--card)',
        border:
          hov && hover ? '1px solid var(--primary-30)' : '1px solid var(--border)',
        borderRadius: 8,
        padding: pad,
        display: 'flex',
        flexDirection: 'column',
        transition: 'border-color 0.25s, box-shadow 0.25s, transform 0.25s',
        boxShadow:
          hov && hover
            ? '0 0 0 1px var(--primary-15), 0 8px 24px rgba(0,0,0,0.3)'
            : 'none',
        ...style,
      }}
    >
      {title ? (
        <div
          style={{
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'space-between',
            marginBottom: density === 'compact' ? 6 : 10,
          }}
        >
          <div style={{ display: 'flex', alignItems: 'center', gap: 8 }}>
            <span
              style={{
                width: 3,
                height: 12,
                background: 'var(--primary)',
                borderRadius: 1,
              }}
            />
            <span
              style={{
                fontSize: 13,
                fontWeight: 500,
                letterSpacing: 0.5,
                color: 'var(--fg)',
              }}
            >
              {title}
            </span>
          </div>
          {action}
        </div>
      ) : null}
      {children}
    </div>
  );
}
