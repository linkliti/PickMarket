// import "@/globals.css";
import "@/styles/index.css";
import React from "react";
import ReactDOM from "react-dom/client";
import { RouterProvider } from "react-router-dom";

import mainRouter from "@/routes/mainRouter";
import devErrorOverlay from "@/utilities/pmutils";
import { createSyncStoragePersister } from "@tanstack/query-sync-storage-persister";
import { QueryClient } from "@tanstack/react-query";
import { PersistQueryClientProvider } from "@tanstack/react-query-persist-client";

devErrorOverlay();

const persister = createSyncStoragePersister({
  storage: window.localStorage,
});

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      refetchOnWindowFocus: false,
      gcTime: 1000 * 60 * 60 * 24, // 24 hours
    },
  },
});

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <PersistQueryClientProvider
      client={queryClient}
      persistOptions={{ persister }}
    >
      <RouterProvider router={mainRouter} />
    </PersistQueryClientProvider>
  </React.StrictMode>,
);
