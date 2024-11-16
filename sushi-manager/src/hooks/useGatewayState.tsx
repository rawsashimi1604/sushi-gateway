import { useRecoilState } from "recoil";
import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { gatewayState } from "../states/GatewayState";
import AdminApiService from "../api/services/admin/AdminApiService";

export const useGatewayData = () => {
  const [gatewayInfo, setGatewayInfo] = useRecoilState<any>(gatewayState);
  const navigate = useNavigate();

  useEffect(() => {
    const fetchData = async () => {
      try {
        let data = await AdminApiService.getGatewayData();
        let config = await AdminApiService.getGatewayConfig();
        setGatewayInfo({
          gateway: data.data,
          config: config.data,
        });
      } catch (err: any) {
        if (err.response.status !== 401) {
          navigate("/login");
        }
      }
    };

    fetchData();
  }, [setGatewayInfo, navigate]);

  return gatewayInfo;
};
