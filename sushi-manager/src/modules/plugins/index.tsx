import Container from "../../components/layout/Container";
import DashboardCard from "../../components/layout/DashboardCard";

import Header from "../../components/typography/Header";
import PluginModal from "./PluginModal";
import PluginTable from "./PluginTable";

function PluginsModule() {
  return (
    <Container>
      <PluginModal />
      <DashboardCard>
        <div className="p-6">
          <Header text="plugins" align="left" size="sm" />
          <PluginTable />
        </div>
      </DashboardCard>
    </Container>
  );
}

export default PluginsModule;
