import { IoMdInformationCircle } from "react-icons/io";
import HttpMethodTag from "./HttpMethodTag";
import RouteModal from "./RouteModal";
import { useState } from "react";

interface RouteTableProps {
  routes: any;
}

function RouteTable({ routes }: RouteTableProps) {
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
              <span>path</span>
              <IoMdInformationCircle className="text-lg mb-0.5" />
            </div>
          </th>
          <th className="px-6 py-3">
            <div className="flex flex-row items-center gap-2">
              <span>methods</span>
              <IoMdInformationCircle className="text-lg mb-0.5" />
            </div>
          </th>
          <th className="px-6 py-3">
            <div className="flex flex-row items-center gap-2">
              <span>service</span>
              <IoMdInformationCircle className="text-lg mb-0.5" />
            </div>
          </th>
        </tr>
      </thead>
      <tbody className="font-lora tracking-wider">
        {routes &&
          routes.map((route: any, i: number) => {
            return <RouteTableRow key={i} route={route} />;
          })}
      </tbody>
    </table>
  );
}

interface RouteTableRowProps {
  route: any;
}

function RouteTableRow({ route }: RouteTableRowProps) {
  const [showModal, setShowModal] = useState<boolean>(false);

  return (
    <>
      {showModal && (
        <RouteModal
          showModal={showModal}
          onClose={() => setShowModal(false)}
          route={route}
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
          {!!route.name && route.name}
        </td>

        <td scope="row" className="px-6 py-4 font-medium whitespace-nowrap">
          {!!route.path && route.path}
        </td>

        <td scope="row" className="px-6 py-4 font-medium whitespace-nowrap">
          {!!route.methods &&
            route.methods.map((method: any, i: number) => {
              return <HttpMethodTag method={method} key={i} />;
            })}
        </td>

        <td scope="row" className="px-6 py-4 font-medium whitespace-nowrap">
          {!!route.service && route.service}
        </td>
      </tr>
    </>
  );
}

export default RouteTable;
