import Container from "../../components/layout/Container";
import DashboardCard from "../../components/layout/DashboardCard";
import { useGatewayData } from "../../hooks/useGatewayState";
import EnvConfiguration from "./EnvConfiguration";
import GatewayConfiguration from "./GatewayConfiguration";
import Summary from "./Summary";

function IndexModule() {
  // Get some information from Sushi proxy API, probably from global state.
  const gatewayInfo = useGatewayData();

  function parseRouteData(): any[] {
    const routes: any[] = [];
    if (gatewayInfo.services?.length > 0) {
      gatewayInfo?.services.forEach((service: any) => {
        service?.routes.forEach((route: any) => {
          routes.push({ ...route, service: service?.name });
        });
      });
    }
    return routes;
  }

  // TODO: add a loading state
  // TODO: add info storage as well as info bubble
  // TODO: add graph
  return (
    <Container>
      <div className="flex flex-col gap-6">
        <Summary />
        <div className="grid grid-cols-2 gap-6">
          <EnvConfiguration />
          <GatewayConfiguration />
        </div>
        <DashboardCard className="p-6">graph to be added...</DashboardCard>
      </div>
    </Container>
  );
}

export default IndexModule;
