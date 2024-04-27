import Header from "@/components/base/Header";
import { Outlet } from "react-router-dom";

export default function Layout() {
  if (import.meta.env.DEV) {
    // Runtime error overlay in DEV
    window.onerror = (_event, _source, _lineno, _colno, err) => {
      const ErrorOverlay = customElements.get("vite-error-overlay");
      if (!ErrorOverlay) {
        return;
      }
      if (!err) {
        err = new Error("Unknown error");
      }
      err.message = "Runtime Error: " + err.message;
      const overlay = new ErrorOverlay(err);
      document.body.appendChild(overlay);
    };
  }

  return (
    <>
      <div className="flex min-h-screen min-w-[320px] flex-col items-cente px-16 py-4 max-md:px-2 max-w-[1280px] mx-auto">
        {/* <div className="flex min-h-screen w-full max-w-[1065px] flex-col max-md:max-w-full"> */}
        <Header />
        <Outlet />
        {/* </div> */}
      </div>
    </>
  );
}
