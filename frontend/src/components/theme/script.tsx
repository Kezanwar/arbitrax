'use client';

export function ThemeScript() {
  const script = `
    try {
      const theme = localStorage.getItem('theme')
      if (theme === 'dark' || (!theme && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
        document.documentElement.classList.add('dark')
      } else {
        document.documentElement.classList.remove('dark')
      }
    } catch (_) {}
  `;

  return <script dangerouslySetInnerHTML={{ __html: script }} />;
}
