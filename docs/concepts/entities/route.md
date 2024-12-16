# Route

A **Route** in Sushi Gateway defines how requests are handled for a specific API path. Routes determine the flow of traffic, specifying the backend service and applying necessary plugins or rules.

## Fields in a Route

A Route configuration consists of the following fields:

| Field     | Type                              | Description                                                                     |
| --------- | --------------------------------- | ------------------------------------------------------------------------------- |
| `name`    | String                            | A unique identifier for the route.                                              |
| `path`    | String                            | The URL path that this route matches.                                           |
| `methods` | Array of Strings                  | The HTTP methods supported by this route (e.g., `GET`, `POST`).                 |
| `plugins` | [Plugin](../entities/plugin.md)[] | A list of plugins applied to this route for authentication, rate limiting, etc. |

## Example Configuration

Here’s an example of a route definition in `config.json`:

```json
{
  "name": "example-route",
  "path": "/v1/sushi",
  "methods": ["GET", "POST"],
  "plugins": [
    {
      "name": "rate_limit",
      "enabled": true,
      "config": {
        "limit_second": 10,
        "limit_min": 100,
        "limit_hour": 1000
      }
    }
  ]
}
```

### Key Fields Explained

1. **`name`**:

   - A unique identifier for the route.
   - Example: `"example-route"`.

2. **`path`**:

   - Specifies the URL path that the route matches.
   - Example: `"/v1/sushi"` matches requests to `/v1/sushi`.

3. **`methods`**:

   - Defines the HTTP methods that this route supports.
   - Example: `["GET", "POST"]` allows both GET and POST requests.

4. **`plugins`**:

   - Applies additional functionalities to the route, such as rate limiting or authentication.
   - Example:
     ```json
     {
       "name": "rate_limit",
       "enabled": true,
       "config": {
         "limit_second": 10,
         "limit_min": 100,
         "limit_hour": 1000
       }
     }
     ```

   ::: tip
   Explore all available plugins in the **[Plugins Overview](../../plugins/index.md)** page for more information.
   :::

## Relationships with Other Entities

Routes are associated with the following entities:

- **[Service](../entities/service.md)**: Routes belong to a service and determine how requests are forwarded to upstreams.
- **[Plugin](../entities/plugin.md)**: Plugins enhance route functionality by adding features like authentication or traffic shaping.

---

The Route entity is crucial for managing traffic flow in Sushi Gateway. To understand other entities, see:

- **[Service](../entities/service.md)**
- **[Upstream](../entities/upstream.md)**
- **[Plugin](../entities/plugin.md)**
