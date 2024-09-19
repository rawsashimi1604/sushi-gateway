import { useState } from "react";
import Modal from "../../components/layout/Modal";
import { IoMdInformationCircle } from "react-icons/io";
import JsonView from "react18-json-view";
import HttpMethodTag from "./HttpMethodTag";
import Tag from "../../components/typography/Tag";

function RouteModal() {
  const [isRouteModalOpen, setIsRouteModalOpen] = useState(true);
  const openModal = () => setIsRouteModalOpen(true);
  const closeModal = () => setIsRouteModalOpen(false);

  const data = {
    name: "get-sushi",
    path: "/v1/sushi",
    methods: ["GET"],
    plugins: [],
  };

  return (
    <Modal isOpen={isRouteModalOpen} onClose={closeModal} title="Route">
      <section className="flex flex-col gap-4 font-lora tracking-wider font-light text-sm">
        <div className="flex gap-2 items-center">
          <HttpMethodTag method="GET" />
          <span className="font-extralight font-sans tracking-widest text-md">
            /v1/sushi
          </span>
        </div>

        {/* Route Name */}
        <div className="flex gap-2">
          <div className="w-[110px] flex items-center gap-2">
            <span>name</span>
            <IoMdInformationCircle className="text-lg" />
          </div>
          <span>get-sushi</span>
        </div>

        {/* Route Service */}
        <div className="flex gap-2">
          <div className="w-[110px] flex items-center gap-2">
            <span>service</span>
            <IoMdInformationCircle className="text-lg" />
          </div>
          <span>SushiService</span>
        </div>

        {/* Route Plugins */}
        <div className="flex gap-2">
          <div className="w-[105px] flex items-center gap-2">
            <span>plugins</span>
            <IoMdInformationCircle className="text-lg" />
          </div>
          <ul className="flex gap-3">
            <li>
              <Tag value="Basic Auth" />
            </li>
            <li>
              <Tag value="JWT" />
            </li>
            <li>
              <Tag value="CORS" />
            </li>
          </ul>
        </div>

        {/* Route Configuration */}
        <div className="flex flex-col gap-2">
          <div className="flex items-center gap-2">
            <span>configuration json</span>
          </div>
          <JsonView style={{ fontSize: "11px" }} src={data} />
        </div>
      </section>
    </Modal>
  );
}

export default RouteModal;
