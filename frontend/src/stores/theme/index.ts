import { observable, action, makeObservable } from 'mobx';

type Theme = 'light' | 'dark' | 'system';

class ThemeStore {
  theme: Theme = 'system';

  constructor() {
    makeObservable(this, {
      theme: observable,
      setTheme: action
    });

    const saved = localStorage.getItem('$MobX-theme') as Theme;
    if (saved === 'light' || saved === 'dark') {
      this.theme = saved;
    }
  }

  setTheme(theme: Theme) {
    this.theme = theme;
    localStorage.setItem('$MobX-theme', theme);
  }
}

export default ThemeStore;
