import { createBrowserRouter } from "react-router-dom";
import IndexModule from "./modules/index";
import LoginModule from "./modules/login";
import ServicesModule from "./modules/services";
import RoutesModule from "./modules/routes";
import SushiAIModule from "./modules/sushi-ai";
import PluginsModule from "./modules/plugins";

export const router = createBrowserRouter([
  { path: "/login", element: <LoginModule /> },
  { path: "/", element: <IndexModule /> },
  { path: "/services", element: <ServicesModule /> },
  { path: "/routes", element: <RoutesModule /> },
  { path: "/plugins", element: <PluginsModule /> },
  { path: "/sushi-ai", element: <SushiAIModule /> },
]);
