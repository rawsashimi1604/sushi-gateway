# Health Check

Health checking is a feature in Sushi Gateway that helps ensure traffic is only routed to healthy upstream services. When enabled, the gateway periodically checks the health of each upstream service and automatically removes unhealthy instances from the load balancing pools.

## Configuration

Health checks can be configured at the service level in your configuration:

```json
{
  "services": [
    {
      "name": "my-service",
      "health": {
        "enabled": true,
        "path": "/health"
      },
      "upstreams": [
        {
          "id": "upstream1",
          "host": "localhost",
          "port": 8081
        }
      ],
      ...
    }
  ]
}
```

### Health Check Properties

| Property | Type    | Required | Description                                             |
| -------- | ------- | -------- | ------------------------------------------------------- |
| enabled  | boolean | Yes      | Whether health checking is enabled for this service     |
| path     | string  | Yes      | The endpoint path to check for health (e.g., "/health") |

## How It Works

When health checking is enabled for a service:

1. The gateway initializes a health map for each service and its upstreams
2. Each upstream starts with a `NotAvailable` status
3. The gateway periodically sends HTTP GET requests to each upstream's health endpoint
4. Based on the response:
   - `200 OK` → Status set to `Healthy`
   - Non-200 response → Status set to `Unhealthy`
   - Connection error → Status set to `Unhealthy`

### Health Status Types

There are three possible health statuses:

- `Healthy`: The upstream is available and responding correctly
- `Unhealthy`: The upstream is not responding or returning errors
- `NotAvailable`: Initial state or health check is disabled

## Integration with Load Balancing

Health check status directly affects load balancing behavior:

1. When health check is disabled:

   - All upstreams are considered available
   - Load balancing works across all configured upstreams

2. When health check is enabled:
   - Only healthy upstreams receive traffic
   - Unhealthy upstreams are automatically removed from the load balancing pool
   - If all upstreams are unhealthy, requests will receive a "no upstreams available" response

### Example

```json
{
  "services": [
    {
      "name": "api-service",
      "protocol": "http",
      "health": {
        "enabled": true,
        "path": "/health"
      },
      "upstreams": [
        {
          "id": "api1",
          "host": "api1.example.com",
          "port": 8080
        },
        {
          "id": "api2",
          "host": "api2.example.com",
          "port": 8080
        }
      ],
      "load_balancing_strategy": "round_robin",
      ...
    }
  ]
}
```

In this example:

- Health checks are performed on `http://api1.example.com:8080/health` and `http://api2.example.com:8080/health`
- Round-robin load balancing only occurs between healthy upstreams
- If either upstream becomes unhealthy, it's automatically skipped in the rotation

## Best Practices

1. **Health Check Endpoint**:

   - Implement a lightweight health check endpoint
   - Return 200 OK only when the service is truly ready to handle requests
   - Include basic service dependency checks if relevant

2. **Monitoring**:

   - Monitor the health check status of your upstreams
   - Set up alerts for when upstreams become unhealthy
   - Keep track of health check history for troubleshooting

3. **Configuration**:
   - Choose an appropriate health check path that doesn't interfere with your service
   - Consider security implications of your health check endpoint
   - Use appropriate timeouts and intervals for your use case
