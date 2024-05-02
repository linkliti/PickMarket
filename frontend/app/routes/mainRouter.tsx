import Layout from "@/layouts/Layout";
import NotFound from "@/pages/404";
import Categories from "@/pages/Categories";
import ErrorBoundary from "@/pages/ErrorBoundary";
import Favorites from "@/pages/Favorites";
import Hero from "@/pages/Hero";
import ItemsPage from "@/pages/Items";
import Test from "@/pages/Test";
import { createBrowserRouter } from "react-router-dom";

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
        element: <ItemsPage />,
      },
      {
        path: "/categories",
        element: <Categories />,
      },
    ],
  },
  {
    element: <Test />,
    path: "/test",
  },
]);

export default mainRouter;
