import { IoMdInformationCircle } from "react-icons/io";
import DashboardCard from "../../components/layout/DashboardCard";
import Header from "../../components/typography/Header";
import Tag from "../../components/typography/Tag";

interface EnvConfigurationItemProps {
  item: string;
  value?: string;
  tooltip?: string;
}

function EnvConfigurationItem({
  item,
  value,
  tooltip,
}: EnvConfigurationItemProps) {
  return (
    <div className="flex flex-row">
      <div className="w-[200px] flex gap-3 items-start justify-start">
        <h2 className="tracking-wider text-sm">{item}</h2>
        {tooltip && <IoMdInformationCircle className="text-md mt-0.5" />}
      </div>
      <span className="tracking-wider text-sm font-serif pr-4">::</span>
      <Tag value={value || ""} className="tracking-wide text-sm font-mono" />
    </div>
  );
}

interface EnvConfigurationProps {
  config: any;
}

function EnvConfiguration({ config }: EnvConfigurationProps) {
  return (
    <DashboardCard className="flex flex-col gap-2 p-6 ">
      <div className="">
        <Header text="environment configuration" align="left" size="sm" />
        <div className="mt-4 flex flex-col gap-2">
          {Object.keys(config).map((configKey, i) => {
            return (
              <EnvConfigurationItem
                key={i}
                item={configKey}
                value={config[configKey].toString()}
              />
            );
          })}
        </div>
      </div>
    </DashboardCard>
  );
}

export default EnvConfiguration;
