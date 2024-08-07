import { useState } from "react";
import { plugins } from "../../data/plugins";
import JsonView from "react18-json-view";
import { IoIosArrowDown, IoIosArrowUp } from "react-icons/io";

interface PluginDropdownProps {
  name: string;
  data: object;
}

function PluginDropdown({ name, data }: PluginDropdownProps) {
  const [isClicked, setIsClicked] = useState(false);

  return (
    <div className="bg-neutral-100 px-4 py-2 rounded-lg shadow-sm">
      <div
        className="flex items-center justify-between"
        onClick={() => setIsClicked((prev) => !prev)}
      >
        <div className="flex items-center gap-2">
          <h1 className="text-md tracking-wide">{plugins[name]}</h1>
          <h1 className="text-xs italic text-neutral-800 mt-0.5">{name}</h1>
        </div>

        {isClicked ? <IoIosArrowUp /> : <IoIosArrowDown />}
      </div>
      {isClicked && (
        <div className="mt-4">
          <JsonView src={data} />
        </div>
      )}
    </div>
  );
}

export default PluginDropdown;
