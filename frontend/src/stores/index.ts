import { observer } from 'mobx-react-lite';
import AuthStore from './auth';
import ThemeStore from './theme';

export class RootStore {
  auth = new AuthStore(this);
  theme = new ThemeStore();
}

const store = new RootStore();

export default store;

export { observer };
