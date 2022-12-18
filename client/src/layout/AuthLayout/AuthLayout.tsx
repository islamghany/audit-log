import { Outlet } from "react-router-dom";
import AuthFooter from "./AuthFooter";
import AuthHeader from "./AuthHeader";
import blurCyanImage from "@/assets/images/blur-cyan.png";

export const AuthLayout = () => {
  return (
    <div className="overflow-hidden w-full min-h-screen bg-background flex flex-col">
      <div className="">
        <AuthHeader />
      </div>
      <main className="flex-1 flex">
        <Outlet />
      </main>
      <div>
        <AuthFooter />
      </div>
    </div>
  );
};
