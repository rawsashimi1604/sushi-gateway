import Container from "../../components/layout/Container";
import { useRecoilState } from "recoil";
import { gatewayState } from "../../states/GatewayState";
import Global from "./Global";
import Json from "./Json";
import Routes from "./Routes";
import Services from "./Services";
import { useNavigate } from "react-router-dom";

import { useEffect } from "react";
import AdminAuth from "../../api/services/admin/AdminAuth";

function IndexModule() {
  // Get some information from Sushi proxy API, probably from global state.
  const navigate = useNavigate();
  const [gatewayInfo, setGatewayInfo] = useRecoilState<any>(gatewayState);

  useEffect(() => {
    fetchData();
  }, []);

  async function fetchData() {
    try {
      const res = await AdminAuth.getGatewayData();
      setGatewayInfo(res.data);
    } catch (err: any) {
      if (err.response.status === 401) {
        return navigate("/login");
      }
    }
  }

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
