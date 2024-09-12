import Container from "../../components/layout/Container";
import DashboardCard from "../../components/layout/DashboardCard";
import Header from "../../components/typography/Header";
import Subtitle from "../../components/typography/Subtitle";
import RouteTable from "./RouteTable";

function RoutesModule() {
  return (
    <Container>
      <DashboardCard>
        <div className="p-6">
          <div className="mb-6">
            <Header text="routes  " align="left" size="sm" />
            <Subtitle text="Routes define the different endpoints for each Service." />
          </div>
          <RouteTable />
        </div>
      </DashboardCard>
    </Container>
  );
}

export default RoutesModule;
