import { useState } from "react";
import Modal from "../../components/layout/Modal";
import { IoMdInformationCircle } from "react-icons/io";
import JsonView from "react18-json-view";

function PluginModal() {
  const [isPluginModalOpen, setIsPluginModalOpen] = useState(true);
  const openModal = () => setIsPluginModalOpen(true);
  const closeModal = () => setIsPluginModalOpen(false);

  const data = {
    name: "basic_auth",
    enabled: false,
    data: {
      username: "admin",
      password: "changeme",
    },
  };

  return (
    <Modal isOpen={isPluginModalOpen} onClose={closeModal} title="Plugin Modal">
      <section className="flex flex-col gap-4 font-lora tracking-wider font-light text-sm">
        {/* Plugin Name */}
        <div className="flex gap-2">
          <div className="w-[110px] flex items-center gap-2">
            <span>name</span>
            <IoMdInformationCircle className="text-lg" />
          </div>
          <span>Basic Authentication</span>
        </div>

        {/* Plugin Enabled */}
        <div className="flex gap-2">
          <div className="w-[110px] flex items-center gap-2">
            <span>enabled</span>
            <IoMdInformationCircle className="text-lg" />
          </div>
          <span>true</span>
        </div>

        {/* Plugin Level */}
        <div className="flex gap-2">
          <div className="w-[110px] flex items-center gap-2">
            <span>level</span>
            <IoMdInformationCircle className="text-lg" />
          </div>
          <span>Route</span>
        </div>

        {/* Plugin Applied To */}
        <div className="flex gap-2">
          <div className="w-[110px] flex items-center gap-2">
            <span>applied to</span>
            <IoMdInformationCircle className="text-lg" />
          </div>
          <span>SushiService</span>
        </div>

        {/* Plugin Configuration */}
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

export default PluginModal;
