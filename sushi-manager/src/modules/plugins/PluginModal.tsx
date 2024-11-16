import Modal from "../../components/layout/Modal";
import { IoMdInformationCircle } from "react-icons/io";
import JsonView from "react18-json-view";
import { getPluginAppliedToDisplayText } from ".";

interface PluginModalProps {
  showModal: boolean;
  onClose: () => void;
  plugin: any;
}

function PluginModal({ showModal, onClose, plugin }: PluginModalProps) {
  return (
    <Modal isOpen={showModal} onClose={onClose} title="Plugin">
      <section className="flex flex-col gap-4 font-lora tracking-wider font-light text-sm">
        {/* Plugin Name */}
        <div className="flex gap-2">
          <div className="w-[110px] flex items-center gap-2">
            <span>name</span>
            <IoMdInformationCircle className="text-lg" />
          </div>
          <span>{plugin && plugin.name}</span>
        </div>

        {/* Plugin Enabled */}
        <div className="flex gap-2">
          <div className="w-[110px] flex items-center gap-2">
            <span>enabled</span>
            <IoMdInformationCircle className="text-lg" />
          </div>
          <span>{plugin && plugin.enabled.toString()}</span>
        </div>

        {/* Plugin Level */}
        <div className="flex gap-2">
          <div className="w-[110px] flex items-center gap-2">
            <span>scope</span>
            <IoMdInformationCircle className="text-lg" />
          </div>
          <span>{plugin && plugin.scope}</span>
        </div>

        {/* Plugin Applied To */}
        <div className="flex gap-2">
          <div className="w-[110px] flex items-center gap-2">
            <span>applied to</span>
            <IoMdInformationCircle className="text-lg" />
          </div>
          <span>{plugin && getPluginAppliedToDisplayText(plugin)}</span>
        </div>

        {/* Plugin Configuration */}
        <div className="flex flex-col gap-2">
          <div className="flex items-center gap-2">
            <span>configuration json</span>
          </div>
          <JsonView style={{ fontSize: "11px" }} src={plugin} />
        </div>
      </section>
    </Modal>
  );
}

export default PluginModal;
