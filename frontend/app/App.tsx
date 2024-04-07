import { Outlet } from "react-router-dom";
import Header from "./components/base/Header";

function Layout() {
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
      <div className="min-h-screen flex flex-col items-center px-16 pt-10 pb-20 bg-sky-100 max-md:px-5">
        <div className="flex flex-col w-full max-w-[1065px] max-md:max-w-full">
          <Header />
          <Outlet />
        </div>
      </div>
    </>
  );
}

export default Layout;
