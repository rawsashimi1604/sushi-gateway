# Plugin

A **Plugin** in Sushi Gateway is a modular extension that enhances gateway functionality by applying features such as authentication, rate limiting, and transformations. Plugins can be applied at various scopes, including global, service, and route levels, providing fine-grained control over API behavior.

## Fields in a Plugin

A Plugin configuration consists of the following fields:

| Field     | Type    | Description                                                |
| --------- | ------- | ---------------------------------------------------------- |
| `name`    | String  | The name of the plugin (e.g., `rate_limit`, `basic_auth`). |
| `enabled` | Boolean | Determines whether the plugin is active.                   |
| `config`  | Object  | Plugin-specific configuration options.                     |

::: tip
Use the `enabled` field to toggle plugin functionality without removing it from the configuration.
:::

## Example Configuration

Here’s an example of a plugin configuration in `config.json`:

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

### Key Fields Explained

1. **`name`**:

   - Specifies the type of plugin.
   - Example: `"rate_limit"` applies rate limiting rules.

2. **`enabled`**:

   - Toggles the plugin on or off.
   - Example: `true` activates the plugin.

3. **`config`**:
   - Provides plugin-specific settings.
   - Example: In the rate limit plugin, `limit_second` defines the number of allowed requests per second.

## Relationships with Other Entities

Plugins are closely associated with the following entities:

- **[Service](../entities/service.md)**: Plugins applied at the service level affect all routes within the service.
- **[Route](../entities/route.md)**: Plugins applied at the route level affect only the specific route.

## Available Plugins

Some of the commonly used plugins include:

- **Rate Limit**: Controls the rate of incoming requests.
- **Basic Auth**: Secures routes using basic authentication.
- **CORS**: Handles Cross-Origin Resource Sharing (CORS) policies.
- **JWT**: Validates JSON Web Tokens for secure access control.

::: tip
For a full list of available plugins and their configurations, refer to the **[Plugins Overview](../../plugins/index.md)** page.
:::

To explore other entities, see:

- **[Service](../entities/service.md)**
- **[Route](../entities/route.md)**
- **[Upstream](../entities/upstream.md)**
