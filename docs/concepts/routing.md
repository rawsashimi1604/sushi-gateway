# Routing

Routing is a core feature of Sushi Gateway, enabling the efficient handling of incoming requests by directing them to the appropriate backend services. This mechanism provides flexibility, scalability, and precision in API traffic management.

## How Routing Works

Routing in Sushi Gateway ensures that requests are directed to the appropriate backend services based on predefined rules.

### Route Matching

Routes are matched using the following criteria:

- **Path**: The URL path specified in the route.
- **HTTP Methods**: Supported methods such as GET, POST, etc.
- **Headers**: Optional headers for advanced matching.

## Defining Services

Defining services is a critical step in Sushi Gateway, enabling the management of APIs and their associated upstreams and routes. This mechanism allows you to organize, secure, and control traffic flow effectively.

## What is a Service?

A **service** in Sushi Gateway represents a backend application or a group of upstreams. Each service includes configuration details for routing traffic, applying plugins, and balancing loads across upstream instances.

::: tip
Learn more about the core entities in Sushi Gateway, such as **[Services](../concepts/entities/service.md)**, **[Routes](../concepts/entities/route.md)**, and **[Plugins](../concepts/entities/plugin.md)**.
:::

## Example Service Configuration

Here’s an example of a complete service definition in `config.json`:

```json
{
  "global": {
    "name": "example-gateway",
    "plugins": []
  },
  "services": [
    {
      "name": "example-service",
      "base_path": "/example",
      "protocol": "http",
      "load_balancing_strategy": "round_robin",
      "upstreams": [
        { "id": "upstream_1", "host": "example-app", "port": 3000 }
      ],
      "routes": [
        {
          "name": "example-route",
          "path": "/v1/sushi",
          "methods": ["GET"],
          "plugins": [
            {
              "id": "example-plugin",
              "name": "rate_limit",
              "enabled": true,
              "config": {
                "limit_second": 10,
                "limit_min": 10,
                "limit_hour": 100
              }
            }
          ]
        }
      ]
    }
  ]
}
```

### Key Components of the Configuration

1. **Service-Level Configuration**:

   - **`name`**: A unique identifier for the service.
   - **`base_path`**: The base path for routing requests to the service.
   - **`protocol`**: The protocol (`http` or `https`) used to communicate with upstreams.
   - **`load_balancing_strategy`**: Determines how traffic is distributed across upstreams.

2. **Upstreams**:

   - Define the backend instances associated with the service.
   - Include properties like `id`, `host`, and `port` for each upstream.

3. **Routes**:
   - Define specific API paths and methods handled by the service.
   - Include optional plugins for additional processing.

### Load Balancing

Once a route is matched, Sushi Gateway uses the defined load-balancing strategy to distribute traffic to the upstreams.

| Strategy        | Description                                                                |
| --------------- | -------------------------------------------------------------------------- |
| **Round Robin** | Distributes requests sequentially among upstreams.                         |
| **IP Hash**     | (WIP) Directs requests to the upstream based on IP Hash (sticky session).  |
| **Weighted**    | (WIP) Assigns weights to upstreams for proportionate traffic distribution. |

Set the strategy in the service configuration:

```json
"load_balancing_strategy": "round_robin"
```

### Route-Specific Plugins

You can configure plugins at the route level to handle:

- **Authentication**: Enforce JWT or API key validation.
- **Rate Limiting**: Control the number of requests allowed within a time window.

Example:

```json
{
  "routes": [
    {
      "name": "example-route",
      "path": "/v1/sushi",
      "methods": ["GET"],
      "plugins": [
        {
          "id": "example-plugin",
          "name": "rate_limit",
          "enabled": true,
          "config": {
            "limit_second": 10
          }
        }
      ]
    }
  ]
}
```

## Advanced Features

### Dynamic Parameters

Routes can include dynamic parameters for greater flexibility:

```json
{
  "path": "/v1/sushi/:id"
}
```

This matches paths like `/v1/sushi/123` and `/v1/sushi/456`.

### CORS Handling

Add CORS plugins to services to manage cross-origin requests:

```json
{
  "plugins": [
    {
      "name": "cors",
      "config": {
        "allow_origins": "*",
        "allow_methods": ["GET", "POST"]
      }
    }
  ]
}
```

---

Defining services and understanding routing in Sushi Gateway allows you to precisely control API traffic and backend communication, ensuring a scalable and secure architecture.

::: tip
Routing in Sushi Gateway ensures precise control over API traffic, allowing developers to create robust and efficient microservices architectures. For more details, refer to the **[Configuration Management Guide](../concepts/configuration/index.md)**.
:::
