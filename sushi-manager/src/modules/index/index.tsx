import Container from "../../components/layout/Container";
import { useGatewayData } from "../../hooks/useGatewayState";
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
  return (
    <Container>
      <div>
        <Summary />
      </div>
    </Container>
  );
}

export default IndexModule;
