import { BsBrightnessHighFill } from "react-icons/bs";
import { RxDashboard } from "react-icons/rx";

import Breadcrumbs from "../typography/Breadcrumbs";

function Navbar() {
  return (
    <nav className="w-full flex flex-row items-center gap-1.5 pb-2 border-b border-gray-200">
      <RxDashboard className="w-4 h-4 mt-0.5" />
      <Breadcrumbs />
      <div className="grow justify-self-end flex justify-end">
        <BsBrightnessHighFill className="w-4 h-4 text-black hover:cursor-pointer hover:shadow-lg transition-all duration-150 hover:text-gray-500" />
      </div>
    </nav>
  );
}

export default Navbar;
