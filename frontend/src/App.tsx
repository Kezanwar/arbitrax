import AuthInitializer from './hocs/hocs/auth-initializer';
import Toast from './components/toast';
import Routes from './routes';
import LoadingOverlay from './components/loading-overlay';

function App() {
  return (
    <AuthInitializer>
      <Toast />
      <LoadingOverlay />
      <Routes />
    </AuthInitializer>
  );
}

export default App;
