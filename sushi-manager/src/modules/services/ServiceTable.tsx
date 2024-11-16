import { useState } from "react";
import { IoMdInformationCircle } from "react-icons/io";
import ServiceModal from "./ServiceModal";

interface ServiceTableProps {
  services: any;
}

function ServiceTable({ services }: ServiceTableProps) {
  return (
    <table className="w-full text-sm text-left rtl:text-right">
      <thead className="text-xs uppercase">
        <tr className="font-lora font-light tracking-widest">
          <th className="pl-0 px-6 py-3">
            <div className="flex flex-row items-center gap-2">
              <span>name</span>
              <IoMdInformationCircle className="text-lg mb-0.5" />
            </div>
          </th>
          <th className="px-6 py-3">
            <div className="flex flex-row items-center gap-2">
              <span>base path</span>
              <IoMdInformationCircle className="text-lg mb-0.5" />
            </div>
          </th>
          <th className="px-6 py-3">
            <div className="flex flex-row items-center gap-2">
              <span>protocol</span>
              <IoMdInformationCircle className="text-lg mb-0.5" />
            </div>
          </th>
          <th className="px-6 py-3">
            <div className="flex flex-row items-center gap-2">
              <span>load balancing strategy</span>
              <IoMdInformationCircle className="text-lg mb-0.5" />
            </div>
          </th>
        </tr>
      </thead>
      <tbody className="font-lora tracking-wider">
        {services?.map((service: any, i: number) => {
          return <ServiceTableRow key={i} service={service} />;
        })}
      </tbody>
    </table>
  );
}

interface ServiceTableRowProps {
  service: any;
}

function ServiceTableRow({ service }: ServiceTableRowProps) {
  const [showModal, setShowModal] = useState<boolean>(false);

  return (
    <>
      {showModal && (
        <ServiceModal
          showModal={showModal}
          onClose={() => setShowModal(false)}
          service={service}
        />
      )}
      <tr className="border-b" onClick={() => setShowModal(true)}>
        <td
          scope="row"
          className="pl-0 px-6 py-4 font-medium  whitespace-nowrap"
        >
          {!!service.name && service.name}
        </td>

        <td scope="row" className="px-6 py-4 font-medium whitespace-nowrap">
          {!!service.base_path && service.base_path}
        </td>

        <td scope="row" className="px-6 py-4 font-medium whitespace-nowrap">
          {!!service.protocol && service.protocol}
        </td>

        <td scope="row" className="px-6 py-4 font-medium whitespace-nowrap">
          {!!service.load_balancing_strategy && service.load_balancing_strategy}
        </td>
      </tr>
    </>
  );
}

export default ServiceTable;
