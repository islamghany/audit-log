import { Outlet } from "react-router-dom";
import DashboardFooter from "./DashboardFooter";
import DashboardHeader from "./DashboardHeader";
import DashboardSidebar from "./DashboardSidebar";

export const DashboardLayout = () => {
  return (
    <div>
      <div>
        <DashboardHeader />
      </div>
      <div>
        <DashboardSidebar />
      </div>
      <main>
        <Outlet />
      </main>
      <div>
        <DashboardFooter />
      </div>
    </div>
  );
};
