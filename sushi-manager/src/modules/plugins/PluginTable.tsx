import { IoMdInformationCircle } from "react-icons/io";

function PluginTable() {
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
              <span>applied to</span>
              <IoMdInformationCircle className="text-lg mb-0.5" />
            </div>
          </th>
          <th className="px-6 py-3">
            <div className="flex flex-row items-center gap-2">
              <span>configuration</span>
              <IoMdInformationCircle className="text-lg mb-0.5" />
            </div>
          </th>
        </tr>
      </thead>
      <tbody className="font-lora tracking-wider">
        <tr className="border-b">
          <td
            scope="row"
            className="pl-0 px-6 py-4 font-medium  whitespace-nowrap"
          >
            http log
          </td>

          <td scope="row" className="px-6 py-4 font-medium whitespace-nowrap">
            true
          </td>

          <td scope="row" className="px-6 py-4 font-medium whitespace-nowrap">
            service:: SushiService
          </td>

          <td scope="row" className="px-6 py-4 font-medium whitespace-nowrap">
            view
          </td>
        </tr>
        <tr className="border-b">
          <td
            scope="row"
            className="pl-0 px-6 py-4 font-medium  whitespace-nowrap"
          >
            basic authentication
          </td>

          <td scope="row" className="px-6 py-4 font-medium whitespace-nowrap">
            true
          </td>

          <td scope="row" className="px-6 py-4 font-medium whitespace-nowrap">
            service:: SushiService
          </td>

          <td scope="row" className="px-6 py-4 font-medium whitespace-nowrap">
            view
          </td>
        </tr>
      </tbody>
    </table>
  );
}

export default PluginTable;
