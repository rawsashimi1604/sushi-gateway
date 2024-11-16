import Modal from "../../components/layout/Modal";
import { IoMdInformationCircle } from "react-icons/io";
import JsonView from "react18-json-view";
import HttpMethodTag from "./HttpMethodTag";
import Tag from "../../components/typography/Tag";

interface RouteModalProps {
  showModal: boolean;
  onClose: () => void;
  route: any;
}

function RouteModal({ showModal, onClose, route }: RouteModalProps) {
  return (
    <Modal isOpen={showModal} onClose={onClose} title="Route">
      <section className="flex flex-col gap-4 font-lora tracking-wider font-light text-sm">
        <div className="flex gap-2 items-center">
          {route &&
            route.methods.map((method: any, i: number) => {
              return <HttpMethodTag method={method} key={i} />;
            })}
          <span className="font-extralight font-sans tracking-widest text-md">
            {route && route.path}
          </span>
        </div>

        {/* Route Name */}
        <div className="flex gap-2">
          <div className="w-[110px] flex items-center gap-2">
            <span>name</span>
            <IoMdInformationCircle className="text-lg" />
          </div>
          <span>{route && route.name}</span>
        </div>

        {/* Route Service */}
        <div className="flex gap-2">
          <div className="w-[110px] flex items-center gap-2">
            <span>service</span>
            <IoMdInformationCircle className="text-lg" />
          </div>
          <span>{route && route.service}</span>
        </div>

        {/* Route Plugins */}
        <div className="flex gap-2">
          <div className="w-[105px] flex items-center gap-2">
            <span>plugins</span>
            <IoMdInformationCircle className="text-lg" />
          </div>
          <ul className="flex gap-3">
            {route && route.plugins.length > 0 ? (
              route.plugins.map((plugin: any, i: number) => {
                return (
                  <li>
                    <Tag value={plugin.name} />
                  </li>
                );
              })
            ) : (
              <li>
                <Tag value="None" />
              </li>
            )}
          </ul>
        </div>

        {/* Route Configuration */}
        <div className="flex flex-col gap-2">
          <div className="flex items-center gap-2">
            <span>configuration json</span>
          </div>
          <JsonView style={{ fontSize: "11px" }} src={route} />
        </div>
      </section>
    </Modal>
  );
}

export default RouteModal;
