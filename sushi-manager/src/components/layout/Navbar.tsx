import { RxDashboard } from "react-icons/rx";

import Breadcrumbs from "../typography/Breadcrumbs";

function Navbar() {
  return (
    <nav className="w-full flex flex-row items-center gap-1.5 pb-5 border-b border-gray-200">
      <RxDashboard className="w-4 h-4 mt-0.5" />
      <Breadcrumbs />
    </nav>
  );
}

export default Navbar;
