import Container from "../../components/layout/Container";
import DashboardCard from "../../components/layout/DashboardCard";
import Header from "../../components/typography/Header";
import Subtitle from "../../components/typography/Subtitle";
import { useGatewayData } from "../../hooks/useGatewayState";
import ServiceTable from "./ServiceTable";

function ServicesModule() {
  const gatewayInfo = useGatewayData();

  return (
    <Container>
      <DashboardCard>
        <div className="p-6">
          <div className="mb-6">
            <Header text="services" align="left" size="sm" />
            <Subtitle text="Services define the upstream APIs for the proxy to forward requests to." />
          </div>
          <ServiceTable services={gatewayInfo?.gateway?.services} />
        </div>
      </DashboardCard>
      <div className="mb-24" />
    </Container>
  );
}

export default ServicesModule;
