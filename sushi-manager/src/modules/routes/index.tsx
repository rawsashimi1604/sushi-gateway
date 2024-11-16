import { useEffect, useState } from "react";
import Container from "../../components/layout/Container";
import DashboardCard from "../../components/layout/DashboardCard";
import Header from "../../components/typography/Header";
import Subtitle from "../../components/typography/Subtitle";
import { useGatewayData } from "../../hooks/useGatewayState";
import RouteTable from "./RouteTable";

function RoutesModule() {
  const gatewayInfo = useGatewayData();
  const [routes, setRoutes] = useState<any>(null);

  useEffect(() => setRoutes(parseRouteData()), [gatewayInfo]);

  function parseRouteData(): any[] {
    const routes: any[] = [];
    if (gatewayInfo.gateway.services?.length > 0) {
      gatewayInfo?.gateway.services.forEach((service: any) => {
        service?.routes.forEach((route: any) => {
          routes.push({ ...route, service: service?.name });
        });
      });
    }
    return routes;
  }

  return (
    <Container>
      <DashboardCard>
        <div className="p-6">
          <div className="mb-6">
            <Header text="routes  " align="left" size="sm" />
            <Subtitle text="Routes define the different endpoints for each Service." />
          </div>
          <RouteTable routes={routes} />
        </div>
      </DashboardCard>
    </Container>
  );
}

export default RoutesModule;
