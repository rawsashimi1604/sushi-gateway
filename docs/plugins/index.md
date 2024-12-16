# Plugins Overview in Sushi Gateway

Plugins in Sushi Gateway are modular extensions that enhance the gateway's functionality. They can be used for tasks such as authentication, rate limiting, transformations, and more. Each plugin operates within a middleware chain, allowing precise control over how requests and responses are processed.

## What are Plugins?

Plugins are:

- Reusable components that add features to services and routes.
- Configurable to meet specific API requirements.
- Applied at different scopes (global, service, route) for fine-grained control.

::: tip
Learn about plugin fields and configurations in the **[Plugin Entity Documentation](../concepts/entities/plugin.md)**.
:::

## Plugin Middleware Chain

Plugins in Sushi Gateway operate in a defined middleware chain:

1. **Global Plugins**: Applied to all services and routes.
2. **Service-Level Plugins**: Applied to all routes within a specific service.
3. **Route-Level Plugins**: Applied to individual routes, overriding service and global plugins if applicable.

### Plugin Priority

The table below illustrates the priority of specific plugins in Sushi Gateway. Plugins with higher priority values are executed earlier in the middleware chain.

| Priority | Plugin                                     |
| -------- | ------------------------------------------ |
| 2500     | Bot Protection                             |
| 2000     | Cross Origin Resource Sharing (RFC 6454)   |
| 1600     | Mutual Transport Layer Security (RFC 8705) |
| 1450     | JSON Web Token (RFC 7519)                  |
| 1250     | API Key Authentication                     |
| 1100     | Basic Authentication (RFC 7617)            |
| 951      | Request Size Limit                         |
| 950      | Access Control List                        |
| 910      | Rate Limit                                 |
| 12       | HTTP Log                                   |

::: tip
Use route-level plugins for the highest level of specificity and ensure priority alignment with your gateway logic.
:::

## Available Plugins

Sushi Gateway supports several plugins. Currently, there are **10 plugins** available. The table below provides an overview:

| Plugin Name          | Description                                               | Documentation                                                 |
| -------------------- | --------------------------------------------------------- | ------------------------------------------------------------- |
| `bot_protection`     | Protects against automated bots.                          | [Bot Protection Plugin](../plugins/bot-protection.md)         |
| `cors`               | Manages CORS policies for APIs.                           | [CORS Plugin](../plugins/cors.md)                             |
| `mtls`               | Implements mutual TLS authentication.                     | [mTLS Plugin](../plugins/mtls.md)                             |
| `jwt`                | Validates JSON Web Tokens (JWT).                          | [JWT Plugin](../plugins/jwt.md)                               |
| `key_auth`           | Secures APIs using API Key Authentication.                | [API Key Plugin](../plugins/key-auth.md)                      |
| `basic_auth`         | Secures routes with basic authentication.                 | [Basic Auth Plugin](../plugins/basic-auth.md)                 |
| `request_size_limit` | Limits the size of incoming requests.                     | [Request Size Limit Plugin](../plugins/request-size-limit.md) |
| `acl`                | Manages access control lists for API consumers.           | [Access Control List Plugin](../plugins/acl.md)               |
| `rate_limit`         | Controls request rates for clients.                       | [Rate Limiting Plugin](../plugins/rate-limit.md)              |
| `http_log`           | Logs HTTP requests and responses for monitoring purposes. | [HTTP Log Plugin](../plugins/http-log.md)                     |

::: tip
Click on a plugin name to learn more about its configuration and use cases.
:::

## Example Plugin Configuration

Here’s how to configure a `rate_limit` plugin:

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

### Explanation

- **`name`**: The plugin type (e.g., `rate_limit`).
- **`enabled`**: Toggles the plugin on or off.
- **`config`**: Plugin-specific settings.

## Tips for Using Plugins

::: tip
Combine multiple plugins at the route level to customize behavior for specific APIs.
:::
