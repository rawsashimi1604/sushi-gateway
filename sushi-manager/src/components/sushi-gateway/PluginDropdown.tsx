import { useState } from "react";
import Subheader from "../typography/Subheader";
import ReactJson from "react-json-view";
import { IoIosArrowDown, IoIosArrowUp } from "react-icons/io";

interface PluginDropdownProps {
    name: string;
    data: object;
}

function PluginDropdown({ name, data }: PluginDropdownProps) {
    const [isClicked, setIsClicked] = useState(false);

    return (
        <div className="bg-neutral-100 px-4 py-2 rounded-lg shadow-sm w-[80%]">
            <div
                className="flex items-center justify-between"
                onClick={() => setIsClicked((prev) => !prev)}
            >
                <Subheader text={name} align="left" size="xs" />
                {isClicked ? <IoIosArrowUp /> : <IoIosArrowDown />}
            </div>
            {isClicked && (
                <div className="mt-4">
                    <ReactJson src={data} />
                </div>
            )}
        </div>
    );
}

export default PluginDropdown;
