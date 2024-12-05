import Container from "../../components/layout/Container";
// import DashboardCard from "../../components/layout/DashboardCard";
import { useGatewayData } from "../../hooks/useGatewayState";
import EnvConfiguration from "./EnvConfiguration";
import GatewayConfiguration from "./GatewayConfiguration";
import Summary from "./Summary";

function IndexModule() {
  // Get some information from Sushi proxy API, probably from global state.
  const gatewayInfo = useGatewayData();

  // TODO: add a loading state
  // TODO: add info storage as well as info bubble
  // TODO: add graph

  return (
    <Container>
      <div className="flex flex-col gap-6">
        {gatewayInfo?.gateway.global && <Summary config={gatewayInfo} />}

        <div className="grid grid-cols-2 gap-6">
          <EnvConfiguration config={gatewayInfo?.config} />
          <GatewayConfiguration config={gatewayInfo?.gateway} />
        </div>
        {/* <DashboardCard className="p-6">graph to be added...</DashboardCard> */}
      </div>
    </Container>
  );
}

export default IndexModule;
