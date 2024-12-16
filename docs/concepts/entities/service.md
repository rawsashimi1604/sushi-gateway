# Service

A **Service** in Sushi Gateway represents a backend application or a collection of upstream instances. Services define the core configuration for routing, load balancing, and managing API traffic to backend systems.

## Fields in a Service

A Service configuration consists of the following fields:

| Field                     | Type                                  | Description                                                                                              |
| ------------------------- | ------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `name`                    | String                                | A unique identifier for the service.                                                                     |
| `base_path`               | String                                | The base path for routing requests to this service.                                                      |
| `protocol`                | String                                | The communication protocol to use (`http` or `https`).                                                   |
| `load_balancing_strategy` | String                                | Determines how traffic is distributed across upstreams (`round_robin`, `least_connections`, `weighted`). |
| `upstreams`               | [Upstream](../entities/upstream.md)[] | A list of upstream instances for this service.                                                           |
| `routes`                  | [Route](../entities/route.md)[]       | A list of routes associated with the service.                                                            |
| `plugins`                 | [Plugin](../entities/plugin.md)[]     | A list of plugins applied to the service.                                                                |

## Example Configuration

Here’s an example of a complete service definition in `config.json`:

```json
{
  "name": "example-service",
  "base_path": "/example",
  "protocol": "http",
  "load_balancing_strategy": "round_robin",
  "upstreams": [
    {
      "id": "upstream_1",
      "host": "example-app",
      "port": 3000
    }
  ],
  "routes": [
    {
      "name": "example-route",
      "path": "/v1/sushi",
      "methods": ["GET"],
      "plugins": []
    }
  ],
  "plugins": [
    {
      "name": "basic_auth",
      "enabled": true,
      "config": {
        "username": "user",
        "password": "changeme"
      }
    }
  ]
}
```

### Key Fields Explained

1. **`name`**:

   - Identifies the service uniquely.
   - Example: `"example-service"`.

2. **`base_path`**:

   - Defines the base path for routing requests to this service.
   - Example: `"/example"` routes all requests with this prefix.

3. **`protocol`**:

   - Specifies the protocol for communication (`http` or `https`).
   - Example: `"http"`.

4. **`load_balancing_strategy`**:

   - Determines how traffic is distributed among upstreams.
   - Supported strategies:
     - `round_robin`
     - `least_connections`
     - `weighted`
   - Example: `"round_robin"`.

   ::: tip
   Learn more about these strategies in the **[Load Balancing Concepts](../load-balancing.md)** page.
   :::

5. **`upstreams`**:

   - Lists the backend instances.
   - Each upstream includes:
     - `id`: A unique identifier.
     - `host`: The hostname or IP address.
     - `port`: The service port.
   - Example:
     ```json
     {
       "id": "upstream_1",
       "host": "example-app",
       "port": 3000
     }
     ```

6. **`routes`**:

   - Defines the API paths and methods handled by the service.
   - Includes optional plugins for authentication, rate limiting, etc.
   - Example:
     ```json
     {
       "name": "example-route",
       "path": "/v1/sushi",
       "methods": ["GET"],
       "plugins": []
     }
     ```

7. **`plugins`**:

   - Applies additional functionalities to the service, such as Basic Authentication.
   - Example:
     ```json
     {
       "name": "basic_auth",
       "enabled": true,
       "config": {
         "username": "user",
         "password": "changeme"
       }
     }
     ```

   ::: tip
   Explore all available plugins in the **[Plugins Overview](../../plugins/index.md)** page to see the full list of supported features and configurations.
   :::

---

The Service entity forms the foundation for managing API traffic in Sushi Gateway. To understand related entities, see:

- **[Route](../entities/route.md)**
- **[Upstream](../entities/upstream.md)**
- **[Plugin](../entities/plugin.md)**
