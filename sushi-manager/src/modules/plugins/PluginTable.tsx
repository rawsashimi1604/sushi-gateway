import { IoMdInformationCircle } from "react-icons/io";
import { getPluginAppliedToDisplayText } from ".";
import { useState } from "react";
import PluginModal from "./PluginModal";

interface PluginTableProps {
  plugins: any[];
}

function PluginTable({ plugins }: PluginTableProps) {
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
              <span>enabled</span>
              <IoMdInformationCircle className="text-lg mb-0.5" />
            </div>
          </th>
          <th className="px-6 py-3">
            <div className="flex flex-row items-center gap-2">
              <span>scope</span>
              <IoMdInformationCircle className="text-lg mb-0.5" />
            </div>
          </th>
          <th className="px-6 py-3">
            <div className="flex flex-row items-center gap-2">
              <span>applied to</span>
              <IoMdInformationCircle className="text-lg mb-0.5" />
            </div>
          </th>
        </tr>
      </thead>
      <tbody className="font-lora tracking-wider">
        {plugins &&
          plugins.map((plugin: any, i: number) => {
            return <PluginTableRow key={i} plugin={plugin} />;
          })}
      </tbody>
    </table>
  );
}

interface PluginTableRowProps {
  plugin: any;
}

function PluginTableRow({ plugin }: PluginTableRowProps) {
  const [showModal, setShowModal] = useState<boolean>(false);

  return (
    <>
      {showModal && (
        <PluginModal
          showModal={showModal}
          onClose={() => setShowModal(false)}
          plugin={plugin}
        />
      )}

      <tr
        onClick={() => setShowModal(true)}
        className="border-b cursor-pointer transition-all duration-75 hover:bg-gray-100"
      >
        <td
          scope="row"
          className="pl-0 px-6 py-4 font-medium  whitespace-nowrap"
        >
          {plugin && plugin.name}
        </td>

        <td scope="row" className="px-6 py-4 font-medium whitespace-nowrap">
          {plugin && plugin.enabled.toString()}
        </td>

        <td scope="row" className="px-6 py-4 font-medium whitespace-nowrap">
          {plugin && plugin.scope}
        </td>

        <td scope="row" className="px-6 py-4 font-medium whitespace-nowrap">
          {plugin && getPluginAppliedToDisplayText(plugin)}
        </td>
      </tr>
    </>
  );
}

export default PluginTable;
