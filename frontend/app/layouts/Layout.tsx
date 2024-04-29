import Header from "@/components/base/Header";
import ErrorBoundary from "@/pages/ErrorBoundary";
import { ReactElement } from "react";
import { Outlet } from "react-router-dom";

export default function Layout(): ReactElement {

  ErrorBoundary();

  return (
    <div className="mx-auto flex min-h-screen min-w-[320px] max-w-[1280px] flex-col items-center px-16 py-4 max-md:px-2">
      <Header />
      <Outlet />
    </div>
  );
}
