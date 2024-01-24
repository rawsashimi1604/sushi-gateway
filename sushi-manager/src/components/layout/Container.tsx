import React, { ReactNode } from "react";

import Navbar from "./Navbar";
import Sidebar from "./Sidebar";

interface ContainerProps {
  children: React.ReactElement | React.ReactElement[] | ReactNode;
}

function Container({ children }: ContainerProps) {
  return (
    <div className="flex h-screen w-screen ">
      <div className="min-w-[300px]">
        <Sidebar />
      </div>
      <main className="flex-grow p-6">
        <Navbar />
        <div className="mt-3 overflow-y-scroll">{children}</div>
        {/* <Footer /> */}
      </main>
    </div>
  );
}

export default Container;
