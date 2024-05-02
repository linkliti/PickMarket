// import "@/globals.css";
import "@/styles/index.css";
import React from "react";
import ReactDOM from "react-dom/client";
import { RouterProvider } from "react-router-dom";

import mainRouter from "@/routes/mainRouter";
import devErrorOverlay from "@/utilities/pmutils";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

devErrorOverlay();
const queryClient = new QueryClient();

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <QueryClientProvider client={queryClient}>
      <RouterProvider router={mainRouter} />
    </QueryClientProvider>
  </React.StrictMode>,
);
