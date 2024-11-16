import Modal from "../../components/layout/Modal";
import { IoMdInformationCircle } from "react-icons/io";
import JsonView from "react18-json-view";
import Tag from "../../components/typography/Tag";
import HttpMethodTag from "../routes/HttpMethodTag";

interface ServiceModalProps {
  showModal: boolean;
  onClose: () => void;
  service: any;
}

function ServiceModal({ showModal, onClose, service }: ServiceModalProps) {
  return (
    <Modal isOpen={showModal} onClose={onClose} title="Service">
      <section className="flex flex-col gap-4 font-lora tracking-wider font-light text-sm">
        {/* Service Name */}
        <div className="flex gap-2">
          <div className="w-[110px] flex items-center gap-2">
            <span>name</span>
            <IoMdInformationCircle className="text-lg" />
          </div>
          <span>{service && service.name}</span>
        </div>

        {/* Service Path */}
        <div className="flex gap-2">
          <div className="w-[110px] flex items-center gap-2">
            <span>path</span>
            <IoMdInformationCircle className="text-lg" />
          </div>
          <span>{service && service.base_path}</span>
        </div>

        {/* Service upstreams*/}
        <div className="flex flex-col gap-2 mb-2">
          <div className="w-[110px] flex items-center gap-2">
            <span>upstreams</span>
            <IoMdInformationCircle className="text-lg" />
          </div>
          <ul className="flex flex-col gap-3">
            {service &&
              service.upstreams.map((upstream, i) => {
                const protocol = service.protocol;
                const host = upstream.host;
                const port = upstream.port;

                if (port) {
                  return <li key={i}>{`${protocol}://${host}:${port}`}</li>;
                } else {
                  return <li key={i}>{`${protocol}://${host}`}</li>;
                }
              })}
          </ul>
        </div>

        {/* Route Plugins */}
        <div className="flex gap-2">
          <div className="w-[105px] flex items-center gap-2">
            <span>plugins</span>
            <IoMdInformationCircle className="text-lg" />
          </div>
          <ul className="flex gap-3">
            {service && service.plugins.length > 0 ? (
              service.plugins.map((plugin: any, i: number) => {
                return (
                  <li key={i}>
                    <Tag value={plugin?.name} />
                  </li>
                );
              })
            ) : (
              <li>
                <Tag value="none" />
              </li>
            )}
          </ul>
        </div>

        {/* Service routes*/}
        <div className="flex flex-col gap-2 mb-2">
          <div className="w-[110px] flex items-center gap-2">
            <span>routes</span>
            <IoMdInformationCircle className="text-lg" />
          </div>
          <ul className="flex flex-col gap-3">
            {service &&
              service.routes.map((route: any, i: number) => {
                return (
                  <li className="font-sand tracking-widest" key={i}>
                    <div className="min-w-[70px] inline-block">
                      {route.methods.map((method: any, j: number) => (
                        <HttpMethodTag key={j} method={method} />
                      ))}
                    </div>
                    {route.path}
                  </li>
                );
              })}
          </ul>
        </div>

        {/* Service Configuration */}
        <div className="flex flex-col gap-2">
          <div className="flex items-center gap-2">
            <span>configuration json</span>
          </div>
          <JsonView style={{ fontSize: "11px" }} src={service} />
        </div>
      </section>
    </Modal>
  );
}

export default ServiceModal;
