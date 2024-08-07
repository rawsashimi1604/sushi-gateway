import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import "react18-json-view/src/style.css";
import { RouterProvider } from "react-router-dom";
import { router } from "./router";
import { RecoilRoot } from "recoil";

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <RecoilRoot>
      <RouterProvider router={router} />
    </RecoilRoot>
  </React.StrictMode>
);
