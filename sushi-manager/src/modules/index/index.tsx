import Container from "../../components/layout/Container";
import Global from "./Global";
import Json from "./Json";
import Routes from "./Routes";
import Services from "./Services";
import { useGatewayData } from "../../hooks/useGatewayState";

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

  return (
    <Container>
      <div className="flex flex-col gap-6">
        <Global data={gatewayInfo?.global} />
        <Services data={gatewayInfo?.services} />
        <Routes data={parseRouteData()} />
        <Json data={gatewayInfo} />
      </div>
    </Container>
  );
}

export default IndexModule;
