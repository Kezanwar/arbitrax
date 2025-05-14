import React from 'react';
import { useRoutes } from 'react-router-dom';

import AuthGuard from '@app/hocs/hocs/auth-guard';
import Home from '@app/pages/home';
import Register from '@app/pages/register';
import Signin from '@app/pages/sign-in';

const paths = [
  {
    path: '/',
    element: (
      <AuthGuard>
        <Home />
      </AuthGuard>
    )
  },
  { path: '/sign-in', element: <Signin /> },
  { path: '/register', element: <Register /> }
];

const Routes: React.FC = () => {
  const elements = useRoutes(paths);
  return elements;
};

export default Routes;
