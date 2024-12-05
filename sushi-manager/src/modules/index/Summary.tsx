import { IoMdInformationCircle } from "react-icons/io";
import Header from "../../components/typography/Header";
import DashboardCard from "../../components/layout/DashboardCard";

interface SummaryProps {
  config: any;
}

function Summary({ config }: SummaryProps) {
  function configExists() {
    return config && Object.keys(config).length > 0;
  }
  function countServices() {
    if (!config && !configExists()) {
      return;
    }
    return config?.gateway?.services.length;
  }

  function countRoutes() {
    if (!config && !configExists()) {
      return;
    }
    const routes: any[] = [];
    if (config.gateway.services?.length > 0) {
      config?.gateway.services.forEach((service: any) => {
        service?.routes.forEach((route: any) => {
          routes.push({ ...route, service: service?.name });
        });
      });
    }
    return routes.length;
  }

  function countPlugins() {
    if (!config && !configExists()) {
      return;
    }
    if (!config.gateway.global) {
      return;
    }

    const parsed: any[] = [];

    // Get global scope plugins...
    config.gateway.global.plugins.forEach((plugin: any) => {
      parsed.push({ ...plugin, scope: "global" });
    });

    // Get service scope plugins...
    config.gateway.services.forEach((service: any) => {
      if (service.plugins && service.plugins.length > 0) {
        service.plugins.forEach((plugin: any) => {
          parsed.push({ ...plugin, scope: "service", service: service.name });
        });
      }
    });

    // Get route scope plugins...
    config.gateway.services.forEach((service: any) => {
      service.routes.forEach((route: any) => {
        route.plugins.forEach((plugin: any) => {
          parsed.push({
            ...plugin,
            scope: "route",
            service: service.name,
            route: route.name,
          });
        });
      });
    });

    return parsed.length;
  }

  return (
    <DashboardCard className="flex flex-col gap-2 p-6 ">
      <Header text="summary" align="left" size="sm" />
      <div className="flex items-center gap-8">
        {/* Services */}
        <div className="flex flex-col gap-2 border-r pr-8">
          <div className="flex gap-3 items-start justify-start">
            <h2 className="tracking-wider text-sm">services</h2>
            <IoMdInformationCircle className="text-md mt-0.5" />
          </div>
          <Header
            text={countServices() ? countServices().toString() : "loading"}
            align="left"
            size="lg"
          />
        </div>

        {/* Routes */}
        <div className="flex flex-col gap-2 border-r pr-8">
          <div className="flex gap-3 items-start justify-start">
            <h2 className="tracking-wider text-sm">routes</h2>
            <IoMdInformationCircle className="text-md mt-0.5" />
          </div>
          <Header
            text={
              countRoutes() ? (countRoutes() as number).toString() : "loading"
            }
            align="left"
            size="lg"
          />
        </div>

        {/* Plugins */}
        <div className="flex flex-col gap-2 border-r pr-8">
          <div className="flex gap-3 items-start justify-start">
            <h2 className="tracking-wider text-sm">plugins</h2>
            <IoMdInformationCircle className="text-md mt-0.5" />
          </div>
          <Header
            text={
              countPlugins() ? (countPlugins() as number).toString() : "loading"
            }
            align="left"
            size="lg"
          />
        </div>

        {/* Last Updated */}
        {/* <div className="flex flex-col gap-2 items-start justify-start">
          <div className="flex gap-3 items-start justify-start">
            <h2 className="tracking-wider text-sm">last updated</h2>
            <IoMdInformationCircle className="text-md mt-0.5" />
          </div>
          <Header text="24 Sept 2024 15:00:00" align="left" size="lg" />
        </div> */}
      </div>
    </DashboardCard>
  );
}

export default Summary;
