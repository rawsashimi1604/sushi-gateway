import Container from "../../components/layout/Container";
import DashboardCard from "../../components/layout/DashboardCard";

import Header from "../../components/typography/Header";
import Subtitle from "../../components/typography/Subtitle";
import PluginModal from "./PluginModal";
import PluginTable from "./PluginTable";

function PluginsModule() {
  return (
    <Container>
      <PluginModal />
      <DashboardCard>
        <div className="p-6">
          <div className="mb-6">
            <Header text="plugins" align="left" size="sm" />
            <Subtitle text="Plugins are specific policies applied on the gateway at different levels - global, service and route." />
          </div>
          <PluginTable />
        </div>
      </DashboardCard>
    </Container>
  );
}

export default PluginsModule;
