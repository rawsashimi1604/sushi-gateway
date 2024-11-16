import Container from "../../components/layout/Container";
import DashboardCard from "../../components/layout/DashboardCard";

import Header from "../../components/typography/Header";
import Subtitle from "../../components/typography/Subtitle";
import { useGatewayData } from "../../hooks/useGatewayState";

import PluginTable from "./PluginTable";

export function getPluginAppliedToDisplayText(plugin: any) {
  if (plugin.scope === "global") {
    return "global";
  }

  if (plugin.scope === "service") {
    return plugin.service;
  }

  if (plugin.scope === "route") {
    return plugin.service + "::" + plugin.route;
  }
}

function PluginsModule() {
  const gatewayInfo = useGatewayData();

  function parseGatewayPlugins(): any[] {
    if (!gatewayInfo.gateway.global) {
      return [];
    }

    const parsed: any[] = [];

    // Get global scope plugins...
    gatewayInfo.gateway.global.plugins.forEach((plugin: any) => {
      parsed.push({ ...plugin, scope: "global" });
    });

    // Get service scope plugins...
    gatewayInfo.gateway.services.forEach((service: any) => {
      service.plugins.forEach((plugin: any) => {
        parsed.push({ ...plugin, scope: "service", service: service.name });
      });
    });

    // Get route scope plugins...
    gatewayInfo.gateway.services.forEach((service: any) => {
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

    return parsed;
  }

  return (
    <Container>
      {/* <PluginModal /> */}
      <DashboardCard>
        <div className="p-6">
          <div className="mb-6">
            <Header text="plugins" align="left" size="sm" />
            <Subtitle text="Plugins are specific policies applied on the gateway at different levels - global, service and route." />
          </div>
          <PluginTable plugins={parseGatewayPlugins()} />
        </div>
      </DashboardCard>
      <div className="mb-24" />
    </Container>
  );
}

export default PluginsModule;
