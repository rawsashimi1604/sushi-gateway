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
      <div className="w-[150px] flex gap-3 items-start justify-start">
        <h2 className="tracking-wider text-sm">{item}</h2>
        {tooltip && <IoMdInformationCircle className="text-md mt-0.5" />}
      </div>
      <span className="tracking-wider text-sm font-serif pr-4">::</span>
      <Tag value={value || ""} className="tracking-wide text-sm font-mono" />
    </div>
  );
}

function EnvConfiguration() {
  return (
    <DashboardCard className="flex flex-col gap-2 p-6 ">
      <div className="">
        <Header text="environment configuration" align="left" size="sm" />
        <div className="mt-4 flex flex-col gap-2">
          <EnvConfigurationItem item="proxy_version" value="0.1" />
          <EnvConfigurationItem
            item="manager_url"
            value="http://localhost:5173"
          />
          <EnvConfigurationItem
            item="admin_api_url"
            value="http://localhost:8081"
          />
          <EnvConfigurationItem
            item="proxy_http_url"
            value="http://localhost:8008"
          />
          <EnvConfigurationItem
            item="proxy_https_url"
            value="https://localhost:8443"
          />
          <EnvConfigurationItem item="proxy_version" value="0.1" />
          <EnvConfigurationItem item="data_model" value="stateless" />
          <EnvConfigurationItem
            item="config_file_path"
            value="./config/config.json"
          />
          <EnvConfigurationItem
            item="server_cert_path"
            value="./config/server.crt"
          />
          <EnvConfigurationItem
            item="server_key_path"
            value="./config/server.key"
          />
          <EnvConfigurationItem item="ca_cert_path" value="./config/ca.crt" />
          <EnvConfigurationItem item="admin_user" value="admin" />
          <EnvConfigurationItem item="admin_password" value="changeme" />
          <EnvConfigurationItem item="config_format" value="json" />
        </div>
      </div>
    </DashboardCard>
  );
}

export default EnvConfiguration;
