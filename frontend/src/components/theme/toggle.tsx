'use client';

import { useThemeStore } from '@app/lib/stores/theme';

export function ThemeToggle() {
  const { theme, setTheme } = useThemeStore();

  return (
    <div className="flex gap-2">
      <button onClick={() => setTheme('light')} className="border px-2 py-1">
        Light
      </button>
      <button onClick={() => setTheme('dark')} className="border px-2 py-1">
        Dark
      </button>
      <button onClick={() => setTheme('system')} className="border px-2 py-1">
        System
      </button>
    </div>
  );
}
