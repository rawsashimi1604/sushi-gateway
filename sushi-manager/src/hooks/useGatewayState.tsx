import { useRecoilState } from "recoil";
import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { gatewayState } from "../states/GatewayState";
import AdminAuth from "../api/services/admin/AdminAuth";

export const useGatewayData = () => {
  const [gatewayInfo, setGatewayInfo] = useRecoilState<any>(gatewayState);
  const navigate = useNavigate();

  useEffect(() => {
    const fetchData = async () => {
      try {
        const res = await AdminAuth.getGatewayData();
        setGatewayInfo(res.data);
      } catch (err: any) {
        if (err.response.status === 401) {
          navigate("/login");
        }
      }
    };

    fetchData();
  }, [setGatewayInfo, navigate]);

  return gatewayInfo;
};
