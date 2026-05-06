import { useEffect, useRef, useState } from 'react';

interface CountUpProps {
  value: number;
  duration?: number;
  format?: (v: number) => string;
}

// Animates a number from 0 → value with cubic ease-out. Re-firing the same
// value (e.g. parent re-renders) skips the animation; only a value *change*
// kicks off a new tween, so polling-based dashboards don't replay forever.
export function CountUp({
  value,
  duration = 1100,
  format = (v) => Math.round(v).toLocaleString(),
}: CountUpProps) {
  const [n, setN] = useState(0);
  const startedRef = useRef(false);

  useEffect(() => {
    if (startedRef.current) {
      setN(value);
      return;
    }
    startedRef.current = true;
    const t0 = performance.now();
    let raf = 0;
    const tick = (t: number) => {
      const p = Math.min(1, (t - t0) / duration);
      const eased = 1 - Math.pow(1 - p, 3);
      setN(eased * value);
      if (p < 1) raf = requestAnimationFrame(tick);
    };
    raf = requestAnimationFrame(tick);
    return () => cancelAnimationFrame(raf);
  }, [value, duration]);

  return <>{format(n)}</>;
}
