import { useNavigate } from "react-router-dom";

import GatewayInfo from "../sushi-gateway/GatewayInfo";
import Logo from "../sushi-gateway/Logo";
import SidebarItem from "./SidebarItem";
import AdminAuth from "../../api/services/admin/AdminApiService";
import { useGatewayData } from "../../hooks/useGatewayState";

function Sidebar() {
  const navigate = useNavigate();
  const gatewayInfo = useGatewayData();

  function handleLogout() {
    const logout = async () => {
      try {
        // Deletes the 'token' httpOnly cookie used for sessions
        await AdminAuth.logout();
        navigate("/login");
      } catch (err: any) {
        if (err.response && err.response.status === 401) {
          console.log("logout fail");
        }
        navigate("/login");
      }
    };
    logout();
  }

  return (
    <nav className="relative w-full h-full p-6 pt-8 pr-8 flex flex-col justify-between z-0">
      {/* border */}
      <div className="absolute right-0 border-r border-gray-200 h-full w-full -z-10"></div>
      <div>
        <div className="mb-6">
          <Logo />
        </div>

        <div className="border-b border-gray-200 pb-5">
          {/* TODO: Get reverse proxy_pass name */}
          <GatewayInfo
            gateway={gatewayInfo?.gateway?.global?.name || "loading..."}
            user="admin"
          />
        </div>
        {/* Services, Routes */}
        <div className="mt-4">
          <h2 className="font-bold text-gray-800 tracking-wider font-lora mb-2">
            GATEWAY
          </h2>
          <ul className="flex flex-col gap-2 ">
            <SidebarItem item="Home" href="/" />
            <SidebarItem item="Services" href="/services" />
            <SidebarItem item="Routes" href="/routes" />
            <SidebarItem item="Plugins" href="/plugins" />
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
