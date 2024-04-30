import terminal from "virtual:terminal";

export default function errorOverlay(): void {
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
      err.message = "Runtime error: " + err.message;
      terminal.error(
        err.message + " (" + err.stack?.split("\n", 2)[1].split(" (", 1)[0].trim() + ")",
      );
      const overlay = new ErrorOverlay(err);
      document.body.appendChild(overlay);
    };
  }
}