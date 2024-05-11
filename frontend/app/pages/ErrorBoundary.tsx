import { ReactElement, useEffect } from "react";

export default function ErrorBoundary(): ReactElement {
  useEffect(() => {
    document.title = "Произошла ошибка!";
  }, []);
  return (
    <div className="min-h-screen text-center">
      <h1 className="my-8 text-3xl font-bold">Произошла ошибка!</h1>
      <a
        href="/"
        className="bg-primary text-primary-foreground hover:bg-primary/70 rounded px-4 py-2 font-bold"
      >
        Перезапустить приложение
      </a>
    </div>
  );
}
