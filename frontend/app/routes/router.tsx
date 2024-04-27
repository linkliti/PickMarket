import NotFound from "@/components/base/404";
import Layout from "@/layouts/Layout";
import Categories from "@/pages/Categories";
import Favorites from "@/pages/Favorites";
import Hero from "@/pages/Hero";
import { createBrowserRouter } from "react-router-dom";

export default createBrowserRouter([
  {
    element: <Layout />,
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
