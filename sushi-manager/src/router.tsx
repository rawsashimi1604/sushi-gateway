import { createBrowserRouter } from "react-router-dom";
import IndexModule from "./modules/index";

export const router = createBrowserRouter([
  { path: "/", element: <IndexModule /> },
]);
