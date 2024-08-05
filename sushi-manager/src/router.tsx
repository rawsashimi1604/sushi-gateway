import { createBrowserRouter } from "react-router-dom";
import IndexModule from "./modules/index";
import LoginModule from "./modules/login";

export const router = createBrowserRouter([
  { path: "/login", element: <LoginModule /> },
  { path: "/", element: <IndexModule /> },
]);
