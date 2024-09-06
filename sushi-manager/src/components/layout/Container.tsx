import React, { ReactNode } from "react";

import Navbar from "./Navbar";
import Sidebar from "./Sidebar";

interface ContainerProps {
  children: React.ReactElement | React.ReactElement[] | ReactNode;
}

function Container({ children }: ContainerProps) {
  return (
    <div className="flex w-screen h-screen fixed">
      <div className="min-w-[300px]">
        <Sidebar />
      </div>
      <main className="flex-grow flex flex-col h-full">
        <div className="p-6 pb-0">
          <Navbar />
        </div>
        <div className="p-6 bg-gray-100 overflow-y-scroll min-h-screen">
          {children}
        </div>
      </main>
    </div>
  );
}

export default Container;
