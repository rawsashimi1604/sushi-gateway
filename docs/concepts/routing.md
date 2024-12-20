# Routing Overview

Routing is a core feature of Sushi Gateway, enabling efficient handling of API requests by directing them to the appropriate backend services. This mechanism leverages a structured system of Services, Routes, and Upstreams to ensure precision, scalability, and reliability in API traffic management.

## How Routing Works

### Request Structure

Sushi Gateway employs a **path-based routing mechanism**, using the protocol, HTTP method, and path to direct traffic to the appropriate backend services (Upstreams). Here’s how it works with an example request:

```http
https://api.gateway.com:8443/sushi/restaurant
```

### Components of the Request

1. **Scheme**: `https`
   - Specifies the communication protocol. Since we are using https, we are hitting the HTTPS endpoint of the gateway.
2. **Domain**: `api.gateway.com`
   - Identifies the API Gateway host.
3. **Port**: `8443`
   - Secure HTTPS port where gateway is hosted.
4. **Service Path**: `/sushi`
   - Matches the `base_path` of a Service entity.
5. **Route Path**: `/restaurant`
   - Matches the `path` of a Route within the Service.

### Routing Workflow

1. **Receive Request**:

   - The gateway receives the incoming request at `https://api.gateway.com:8443/sushi/restaurant`.

2. **Match Service**:

   - The gateway matches `/sushi` to the `base_path` of a configured Service entity in memory.

3. **Match Route**:

   - Within the matched Service, the gateway identifies `/restaurant` as the `path` of a Route.

4. **Load Balancer**:

   - Based on the Service’s load balancing strategy, the gateway selects an appropriate upstream (e.g., `sushi.jp`).

5. **Forward Request**:

   - Constructs the upstream path: `https://sushi.jp/restaurant`.
   - Applies middleware plugins before forwarding the request.

6. **Process Response**:
   - Receives the response from the upstream, processes it through middleware plugins, and sends it back to the client.

## Example Configuration

### Service Definition

```json
{
  "services": [
    {
      "name": "sushi-service",
      "base_path": "/sushi",
      "protocol": "http",
      "load_balancing_strategy": "round_robin",
      "upstreams": [
        { "id": "upstream_1", "host": "localhost", "port": 8001 },
        { "id": "upstream_2", "host": "localhost", "port": 8002 }
      ],
      "routes": [
        {
          "name": "get-sushi",
          "path": "/restaurant",
          "methods": ["GET"],
          "plugins": []
        }
      ]
    }
  ]
}
```

### Request Flow

1. **Request**: `https://api.gateway.com:8443/sushi/restaurant`
2. **Service Match**:
   - `base_path`: `/sushi` → `sushi-service`
3. **Route Match**:
   - `path`: `/restaurant` → `get-sushi`
4. **Upstream Selection**:
   - Load balancing strategy: `round_robin`
   - Selected upstream: `localhost:8001`
5. **Final Upstream Path**:
   - `https://localhost:8001/restaurant`

## Middleware Chain

### Plugins Applied

1. **Global Plugins**: Configured at the gateway level.
2. **Service Plugins**: Configured for the matched Service entity.
3. **Route Plugins**: Configured for the matched Route entity.

The gateway executes plugins in the following order:

```bash
plugins = global_plugins + service_plugins + route_plugins
```

### Example Plugins

- **Rate Limiting**: Controls request limits per time window.
- **Authentication**: Enforces API key or JWT validation.

## Dynamic API Path Routing Matching

Sushi Gateway also comes with Dynamic API path routing matchin support, allowing for flexible and reusable routes by using parameterized paths.

### How It Works

In a route configuration, dynamic segments are defined using curly braces (`{}`) to represent placeholders for values. When an incoming request matches the dynamic segment, Sushi Gateway extracts the value and passes it to the upstream service or processes it as needed.

### Example Configuration

```json
{
  "routes": [
    {
      "name": "get-sushi-by-id",
      "path": "/sushi/{id}",
      "methods": ["GET"],
      "plugins": []
    }
  ]
}
```

### Example Request and Routing

#### Request:

```http
GET https://api.gateway.com:8443/sushi/123
```

#### Route Match:

- **Base Path**: `/sushi`
- **Dynamic Segment**: `{id}` → `123`

#### Upstream Request:

The dynamic value is included in the forwarded request to the upstream:

```http
GET https://upstream-service/sushi/123
```

Routing in Sushi Gateway ensures precise control over API traffic, allowing developers to create robust and efficient microservices architectures. For more details, refer to the **[Configuration Management Guide](../concepts/configuration/index.md)**.
