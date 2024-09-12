import { IoMdInformationCircle } from "react-icons/io";
import DashboardCard from "../../components/layout/DashboardCard";
import Header from "../../components/typography/Header";

interface ConfigurationItemProps {
  item: string;
  value?: string;
  tooltip?: string;
}

function ConfigurationItem({ item, value, tooltip }: ConfigurationItemProps) {
  return (
    <div className="flex flex-row">
      <div className="w-[150px] flex gap-3 items-start justify-start">
        <h2 className="tracking-wider text-sm font-light">{item}</h2>
        {tooltip && <IoMdInformationCircle className="text-md mt-0.5" />}
      </div>
      <span className="tracking-wider text-sm font-serif pr-4">::</span>
      <span className="tracking-wide text-sm">{value}</span>
    </div>
  );
}

function Configuration() {
  return (
    <DashboardCard className="flex flex-col gap-2 p-6 ">
      <div className="">
        <Header text="configuration" align="left" size="sm" />
        <div className="flex flex-col gap-2">
          <ConfigurationItem item="state" value="stateless" />
          <ConfigurationItem item="version" value="0.0.1" />
          <ConfigurationItem item="config_format" value="json" />
        </div>
      </div>
    </DashboardCard>
  );
}

export default Configuration;
