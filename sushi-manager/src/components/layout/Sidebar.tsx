import { useNavigate } from "react-router-dom";

import GatewayInfo from "../sushi-gateway/GatewayInfo";
import Logo from "../sushi-gateway/Logo";
import SidebarItem from "./SidebarItem";

function Sidebar() {
  const navigate = useNavigate();

  function handleLogout() {
    localStorage.removeItem("jwt-token");
    navigate("/login");
  }

  return (
    <nav className="relative w-full h-full p-6 pt-8 pr-8 flex flex-col justify-between z-0">
      {/* border */}
      <div className="absolute right-0 border-r border-gray-200 h-full w-full -z-10"></div>
      <div>
        <div className="mb-10">
          <Logo />
        </div>

        <div className="border-b border-gray-200 pb-5">
          {/* Reverse proxy name */}
          <GatewayInfo gateway="Sushi Gateway" user="admin" />
        </div>
        {/* Services, Routes */}
        <div className="mt-4">
          <h2 className="font-bold text-gray-800 tracking-tighter font-lora">
            gateway
          </h2>
          <ul className="flex flex-col gap-0">
            <SidebarItem item="Home" href="/" />
            <SidebarItem item="Services" href="/services" />
            <SidebarItem item="Routes" href="/routes" />
            <SidebarItem item="Consumers" href="/consumers" />
            <SidebarItem item="Logs" href="/logs" />
          </ul>
        </div>
        {/* Settings , Logout */}
        <div className="mt-4">
          <h2 className="font-bold text-gray-800 tracking-tighter font-lora">
            account
          </h2>
          <ul className="flex flex-col gap-0">
            <SidebarItem item="Account" href="/account" />
            <SidebarItem item="Settings" href="/settings" />
          </ul>
        </div>
      </div>

      <button
        className="w-full flex-end py-2 bg-blue-500 text-white shadow-md rounded-full font-sans border-0 tracking-widest duration-300 transition-all hover:-translate-y-1 hover:shadow-lg"
        onClick={handleLogout}
      >
        logout
      </button>
    </nav>
  );
}

export default Sidebar;
