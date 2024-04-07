import Layout from "@/Layout";
import NotFound from "@/components/base/404";
import Categories from "@/pages/Categories";
import Favorites from "@/pages/Favorites";
import Hero from "@/pages/Hero";
import React from "react";
import ReactDOM from "react-dom/client";
import { RouterProvider, createBrowserRouter } from "react-router-dom";
import "./globals.css";
import "./index.css";

const router = createBrowserRouter([
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

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>,
);
