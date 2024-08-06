import React, { useState } from "react";
import Tag from "../typography/Tag";
import NormalText from "../typography/Text";
import PluginDropdown from "./PluginDropdown";
import { IoIosArrowDown, IoIosArrowUp } from "react-icons/io";

interface ServiceDropdownProps {
  data: any;
}

function ServiceDropdown({ data }: ServiceDropdownProps) {
  const [isClicked, setIsClicked] = useState(false);

  return (
    <div>
      {/* Service metadata */}
      <div className="bg-neutral-200 px-4 py-2 rounded-lg shadow-sm w-[80%]">
        <div
          className="flex items-center justify-between"
          onClick={() => setIsClicked((prev) => !prev)}
        >
          <div className="flex items-center gap-2 mb-2">
            <h1 className="text-md tracking-wide">{data?.name}</h1>
            <h1 className="text-xs italic text-neutral-800 mt-0.5">
              service_name
            </h1>
          </div>

          {isClicked ? <IoIosArrowUp /> : <IoIosArrowDown />}
        </div>
        {isClicked && (
          <div className="flex flex-col items-start gap-2 text-sm">
            <div className="flex items-center gap-2 text-sm">
              <Tag value="name" />
              <NormalText text={data?.name} />
            </div>
            <div className="flex items-center gap-2 text-sm">
              <Tag value="base_path" />
              <NormalText text={data?.base_path} />
            </div>

            <div className="flex items-center gap-2 text-sm">
              <Tag value="protocol" />
              <NormalText text={data?.protocol} />
            </div>

            <div className="flex flex-col justify-center items-start  text-sm">
              <Tag value="upstreams" />
              <div className="flex flex-col bg-neutral-100 px-4 py-2 mt-2">
                {data?.upstreams.map((upstream: any) => {
                  return (
                    <div
                      key={upstream?.host + upstream?.port}
                      className="flex items-center gap-2"
                    >
                      <NormalText text={"Host: " + upstream?.host} />
                      <NormalText text={"Port: " + upstream?.port} />
                    </div>
                  );
                })}
              </div>
            </div>

            <div className="w-full">
              <div className="mb-2">
                <Tag value="plugins" />
              </div>

              <div className="flex flex-col gap-3">
                {data?.plugins.length > 0 ? (
                  data?.plugins.map((plugin: any) => {
                    return (
                      <PluginDropdown
                        key={plugin?.name}
                        name={plugin?.name}
                        data={plugin}
                      />
                    );
                  })
                ) : (
                  <NormalText text="No plugins found" />
                )}
              </div>
            </div>
          </div>
        )}
      </div>
    </div>
  );
}

export default ServiceDropdown;
