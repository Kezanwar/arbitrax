import { useEffect } from 'react';
import { observer } from 'mobx-react-lite';
import store from '@app/stores';

export const ThemeProvider = observer(
  ({ children }: { children: React.ReactNode }) => {
    useEffect(() => {
      const root = document.documentElement;

      const applyTheme = (theme: 'light' | 'dark') => {
        root.classList.remove('light', 'dark');
        root.classList.add(theme);
        store.theme.setTheme(theme);
      };

      if (store.theme.theme === 'system') {
        const systemTheme = window.matchMedia('(prefers-color-scheme: dark)')
          .matches
          ? 'dark'
          : 'light';
        applyTheme(systemTheme);

        const media = window.matchMedia('(prefers-color-scheme: dark)');
        const listener = (e: MediaQueryListEvent) => {
          applyTheme(e.matches ? 'dark' : 'light');
        };

        media.addEventListener('change', listener);
        return () => media.removeEventListener('change', listener);
      } else {
        applyTheme(store.theme.theme);
      }
    }, [store.theme.theme]); // responds to observable change

    return <>{children}</>;
  }
);
