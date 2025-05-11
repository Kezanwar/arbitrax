import React, { type FC } from 'react';
import { AppSidebar } from './components/app-sidebar';
import DashboardHeader from '@app/layouts/dashboard/components/header';
import { SidebarInset, SidebarProvider } from '@app/components/ui/sidebar';

type Props = {
  children: React.ReactNode;
};

const DashboardLayout: FC<Props> = ({ children }) => {
  return (
    <SidebarProvider>
      <AppSidebar />
      <SidebarInset>
        <div className="bg-background text-foreground flex min-h-screen flex-col">
          <DashboardHeader />
          <main className="flex-1 overflow-y-auto p-4">{children}</main>
        </div>
      </SidebarInset>
    </SidebarProvider>
  );
};

export default DashboardLayout;
