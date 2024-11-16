export function checkGatewayPersistenceMode(gateway: any) {
  if (gateway?.config) {
    return gateway?.config?.PersistenceConfig;
  }
}
