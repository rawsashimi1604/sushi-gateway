import { useState } from "react";
import Modal from "../../components/layout/Modal";
import { IoMdInformationCircle } from "react-icons/io";
import JsonView from "react18-json-view";
import Tag from "../../components/typography/Tag";
import HttpMethodTag from "../routes/HttpMethodTag";

function ServiceModal() {
  const [isServiceModalOpen, setIsServiceModalOpen] = useState(true);
  const openModal = () => setIsServiceModalOpen(true);
  const closeModal = () => setIsServiceModalOpen(false);

  const data = {
    name: "sushi-svc",
    base_path: "/sushi-service",
    protocol: "http",
    load_balancing_strategy: "round_robin",
    upstreams: [
      { host: "sushi-svc-1", port: 3000 },
      { host: "sushi-svc-2", port: 3000 },
    ],
    plugins: [],
    routes: [
      {
        name: "get-sushi",
        path: "/v1/sushi",
        methods: ["GET"],
        plugins: [],
      },
      {
        name: "get-sushi-restaurants",
        path: "/v1/sushi/restaurant",
        methods: ["GET"],
        plugins: [],
      },
      {
        name: "sushi-provision-jwt",
        path: "/v1/token",
        methods: ["GET"],
        plugins: [],
      },
    ],
  };

  return (
    <Modal isOpen={isServiceModalOpen} onClose={closeModal} title="Service">
      <section className="flex flex-col gap-4 font-lora tracking-wider font-light text-sm">
        {/* Service Name */}
        <div className="flex gap-2">
          <div className="w-[110px] flex items-center gap-2">
            <span>name</span>
            <IoMdInformationCircle className="text-lg" />
          </div>
          <span>SushiService</span>
        </div>

        {/* Service Path */}
        <div className="flex gap-2">
          <div className="w-[110px] flex items-center gap-2">
            <span>path</span>
            <IoMdInformationCircle className="text-lg" />
          </div>
          <span>http://sushi-service.com</span>
        </div>

        {/* Service upstreams*/}
        <div className="flex flex-col gap-2 mb-2">
          <div className="w-[110px] flex items-center gap-2">
            <span>upstreams</span>
            <IoMdInformationCircle className="text-lg" />
          </div>
          <ul className="flex flex-col gap-3">
            <li>http://localhost:8080</li>
            <li>http://localhost:8081</li>
            <li>http://localhost:8082</li>
          </ul>
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

        {/* Service routes*/}
        <div className="flex flex-col gap-2 mb-2">
          <div className="w-[110px] flex items-center gap-2">
            <span>routes</span>
            <IoMdInformationCircle className="text-lg" />
          </div>
          <ul className="flex flex-col gap-3">
            <li className="font-sand tracking-widest">
              <HttpMethodTag method="GET" /> /v1/sushi
            </li>
            <li className="font-sand tracking-widest">
              <HttpMethodTag method="POST" /> /v1/sushi
            </li>
            <li className="font-sand tracking-widest">
              <HttpMethodTag method="GET" /> {"/v1/sushi/{id}"}
            </li>
          </ul>
        </div>

        {/* Service Configuration */}
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

export default ServiceModal;
