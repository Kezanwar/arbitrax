import Logo from '../logo';
import { Spinner } from '../ui/spinner';

const LoadingScreen = () => {
  return (
    <div className="bg-background flex h-[100vh] w-[100vw] flex-col items-center justify-center">
      <Spinner size={40} />
      <Logo />
    </div>
  );
};

export default LoadingScreen;
