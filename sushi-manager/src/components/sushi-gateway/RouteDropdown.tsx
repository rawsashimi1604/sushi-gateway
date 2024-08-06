import React, { useState } from "react";
import Tag from "../typography/Tag";
import NormalText from "../typography/Text";
import PluginDropdown from "./PluginDropdown";
import { IoIosArrowDown, IoIosArrowUp } from "react-icons/io";

function RouteDropdown() {
  const [isClicked, setIsClicked] = useState(false);

  return (
    <div>
      {/* Route metadata */}
      <div className="bg-neutral-200 px-4 py-2 rounded-lg shadow-sm w-[80%]">
        <div
          className="flex items-center justify-between"
          onClick={() => setIsClicked((prev) => !prev)}
        >
          <div className="flex items-center gap-2 mb-2">
            <h1 className="text-md tracking-wide">route_name</h1>
            <h1 className="text-xs italic text-neutral-800 mt-0.5">
              route_name
            </h1>
          </div>

          {isClicked ? <IoIosArrowUp /> : <IoIosArrowDown />}
        </div>
        {isClicked && (
          <div className="flex flex-col items-start gap-2 text-sm">
            <div className="flex items-center gap-2 text-sm">
              <Tag value="service" />
              <NormalText text="some_service_name" />
            </div>

            <div className="flex items-center gap-2 text-sm">
              <Tag value="name" />
              <NormalText text="some_route_name" />
            </div>

            <div className="flex items-center gap-2 text-sm">
              <Tag value="path" />
              <NormalText text="/some_base_path" />
            </div>

            <div className="flex flex-col justify-center items-start  text-sm">
              <Tag value="methods" />
              <div className="flex flex-col bg-neutral-100 px-4 py-2 mt-2">
                <NormalText text="GET" />
                <NormalText text="POST" />
              </div>
            </div>

            <div className="w-full">
              <div className="mb-2">
                <Tag value="plugins" />
              </div>

              {/* Some dropdown for plugin design */}
              <PluginDropdown
                name="http_log"
                data={{
                  name: "http_log",
                  enabled: true,
                  data: {
                    http_endpoint: "http://localhost:8003/v1/log",
                    method: "POST",
                    content_type: "application/json",
                  },
                }}
              />
            </div>
          </div>
        )}
      </div>
    </div>
  );
}

export default RouteDropdown;
