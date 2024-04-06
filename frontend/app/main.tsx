import React from "react";
import ReactDOM from "react-dom/client";
import { ErrorBoundary } from "react-error-boundary";
// @ts-expect-error no types
import terminal from "virtual:terminal";
import App from "./App.tsx";
import "./index.css";

// eslint-disable-next-line react-refresh/only-export-components
function DevErrorComponent({ error }: { error: Error }) {
  return (
    <>
      <p className="text-red-600">Something went wrong</p>
      <p className="text-wrap text-sm">
        <pre className="text-wrap">{error.message}</pre>
        <pre className="text-wrap">{error.stack}</pre>
      </p>
    </>
  );
}

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <ErrorBoundary
      FallbackComponent={DevErrorComponent}
      onError={(err, errInfo) => {
        terminal.error(`Error:`, err.message);
        terminal.error(errInfo.componentStack);
      }}
    >
      <App />
    </ErrorBoundary>
  </React.StrictMode>
);
