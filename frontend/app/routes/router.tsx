import Layout from "@/layouts/Layout";
import NotFound from "@/pages/404";
import Categories from "@/pages/Categories";
import ErrorBoundary from "@/pages/ErrorBoundary";
import Favorites from "@/pages/Favorites";
import Hero from "@/pages/Hero";
import { createBrowserRouter } from "react-router-dom";

export const mainRouter = createBrowserRouter([
  {
    element: <Layout />,
    errorElement: <ErrorBoundary />,
    children: [
      {
        path: "*",
        element: <NotFound />,
      },
      {
        path: "/",
        element: <Hero />,
      },
      {
        path: "/favorites",
        element: <Favorites />,
      },
      {
        path: "/categories",
        element: <Categories />,
      },
    ],
  },
]);


