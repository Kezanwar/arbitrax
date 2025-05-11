import AuthGuard from '@app/hocs/hocs/auth-guard';
import Home from '@app/pages/home';
// import HomePage from '@app/pages/home';
import Signin from '@app/pages/sign-in';
import React from 'react';
import { useRoutes } from 'react-router-dom';

// pages
// import Home from '@app/pages/Home';
// import Links from '@app/pages/Links';
// import File from '@app/pages/File';

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
  { path: '/register', element: <Signin /> }
];

const Routes: React.FC = () => {
  const elements = useRoutes(paths);
  return elements;
};

export default Routes;
