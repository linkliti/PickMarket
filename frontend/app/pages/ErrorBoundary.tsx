import { ReactElement } from "react";

export default function ErrorBoundary(): ReactElement | undefined {
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
      err.message = "Browser runtime error: " + err.message;
      const overlay = new ErrorOverlay(err);
      document.body.appendChild(overlay);
    };
  } else {
    return <div> Ошибка</div>;
  }
}
