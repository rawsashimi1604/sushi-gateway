import { useState } from "react";
import DashboardCard from "../../components/layout/DashboardCard";
import Header from "../../components/typography/Header";
import JsonView from "react18-json-view";

interface GatewayConfigurationProps {
  config: any;
}

function GatewayConfiguration({ config }: GatewayConfigurationProps) {
  const [showConfig, setShowConfig] = useState(false);

  return (
    <DashboardCard className="flex flex-col gap-2 p-6 ">
      <Header text="gateway configuration" align="left" size="sm" />
      <button
        type="button"
        onClick={() => setShowConfig((prev) => !prev)}
        className="w-[80px] mt-2 text-xs py-1.5 px-2 pb-2 tlext-white bg-blue-500 shadow-md rounded-lg font-sans tracking-wider border-0 duration-300 transition-all hover:-translate-y-1 hover:shadow-lg text-white"
      >
        <span>{showConfig ? "hide" : "show"}</span>
      </button>
      {/* TODO: popup modal */}
      {showConfig && (
        <div className="overflow-y-scroll">
          <JsonView src={config} />
        </div>
      )}
    </DashboardCard>
  );
}

export default GatewayConfiguration;
