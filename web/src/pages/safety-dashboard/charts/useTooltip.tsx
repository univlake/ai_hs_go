import { useRef, useState, type ReactNode } from 'react';

interface TipState {
  x: number;
  y: number;
  content: ReactNode;
}

// Shared tooltip primitive for SVG charts. The hook returns a wrapper ref to
// attach to the chart container, plus mouse handlers and a tooltip node to
// render at the end of the container's children. Positioning is relative to
// the wrapper so the tooltip follows the cursor without document offset math.
export function useTooltip() {
  const [tip, setTip] = useState<TipState | null>(null);
  const wrapRef = useRef<HTMLDivElement | null>(null);

  const onMove = (e: { clientX: number; clientY: number }, content: ReactNode) => {
    const wrap = wrapRef.current;
    if (!wrap) return;
    const rect = wrap.getBoundingClientRect();
    setTip({
      x: e.clientX - rect.left + 12,
      y: e.clientY - rect.top + 12,
      content,
    });
  };
  const onLeave = () => setTip(null);

  const tipNode = tip ? (
    <div
      style={{
        position: 'absolute',
        left: tip.x,
        top: tip.y,
        background: 'rgba(10, 14, 22, 0.95)',
        border: '1px solid rgba(255,255,255,0.12)',
        borderRadius: 6,
        padding: '6px 10px',
        fontSize: 11,
        color: '#e6edf5',
        pointerEvents: 'none',
        zIndex: 50,
        whiteSpace: 'nowrap',
        backdropFilter: 'blur(8px)',
        fontFamily: 'inherit',
      }}
    >
      {tip.content}
    </div>
  ) : null;

  return { wrapRef, onMove, onLeave, tipNode };
}
