import React from 'react';
import { Navigate, useRoutes, type RouteObject } from 'react-router-dom';

import Register from '@app/pages/guest/register';
import Signin from '@app/pages/guest/sign-in';
import GuestGuard from '@app/hocs/guest-guard';
import DashboardLayout from '@app/layouts/dashboard';
import Home from '@app/pages/dashboard/home';
import Dummy from '@app/pages/dashboard/dummy';

const paths: RouteObject[] = [
  {
    path: '/',
    element: <DashboardLayout />,
    children: [
      {
        index: true,
        element: <Home />
      },
      {
        path: 'agents',
        element: <Navigate to={'/agents/all'} replace />,
        children: [
          {
            path: 'all',
            element: <Dummy page="All Agents" />
          },
          {
            path: 'deploy',
            element: <Dummy page="Deploy an Agent" />
          }
        ]
      }
    ]
  },
  {
    path: '/sign-in',
    element: (
      <GuestGuard>
        <Signin />
      </GuestGuard>
    )
  },
  {
    path: '/register',
    element: (
      <GuestGuard>
        <Register />
      </GuestGuard>
    )
  }
];

const Routes: React.FC = () => {
  const elements = useRoutes(paths);
  return elements;
};

export default Routes;
