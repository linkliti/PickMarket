/* eslint-disable react-refresh/only-export-components */
import Layout from "@/layouts/Layout";
import NotFound from "@/pages/404";

import ErrorBoundary from "@/pages/ErrorBoundary";
import Favorites from "@/pages/Favorites";
import Hero from "@/pages/Hero";
import PageLoading from "@/routes/PageLoading";
import { lazy, Suspense } from "react";
import { createBrowserRouter } from "react-router-dom";

const ItemsPage = lazy(() => import("@/pages/Items"));
// import ItemsPage from "@/pages/Items";

const Categories = lazy(() => import("@/pages/Categories"));
// import Categories from "@/pages/Categories";

const mainRouter = createBrowserRouter([
  {
    element: <Layout />,
    errorElement: <ErrorBoundary />,
    children: [
      {
        path: "*",
        element: <NotFound />,
      },
      {
        index: true,
        element: <Hero />,
      },
      {
        path: "/favorites",
        element: <Favorites />,
      },
      {
        path: "/items",
        element: (
          <Suspense fallback={PageLoading("Загрузка предпочтений")}>
            <ItemsPage />
          </Suspense>
        ),
      },
      {
        path: "/categories",
        element: (
          <Suspense fallback={PageLoading("Загрузка категорий")}>
            <Categories />
          </Suspense>
        ),
      },
    ],
  },
]);

export default mainRouter;
