import { Outlet } from "react-router-dom";
import AuthFooter from "./AuthFooter";
import AuthHeader from "./AuthHeader";

export const AuthLayout = () => {
  return (
    <div>
      <div>
        <AuthHeader />
      </div>
      <main>
        <Outlet />
      </main>
      <div>
        <AuthFooter />
      </div>
    </div>
  );
};
