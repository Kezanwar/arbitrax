import { ThemeProvider } from './components/theme/provider';

function App() {
  return (
    <ThemeProvider>
      <div className="bg-background">
        <div className="bg-background p-4">
          {' '}
          <div className="bg-paper text-foreground p-6">
            <h1 className="text-primary text-3xl font-bold">
              Welcome to ArbitraX
            </h1>
            <p className="text-secondary mt-2">
              High-speed FX arbitrage in style.
            </p>
          </div>
        </div>
      </div>
    </ThemeProvider>
  );
}

export default App;
