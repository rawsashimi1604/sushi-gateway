# CORS Plugin

The CORS (Cross-Origin Resource Sharing) (`cors`) plugin enables APIs to manage requests from different origins, ensuring compliance with **[RFC 6454](https://datatracker.ietf.org/doc/html/rfc6454)**. This is essential for enabling secure cross-origin requests, particularly in browser-based applications.

## How It Works

The CORS plugin inspects incoming requests for the `Origin` header and validates them against the configured rules. If the request meets the specified criteria, the plugin adds appropriate CORS headers to the response.

Requests that do not meet the criteria are either denied or processed without CORS headers, depending on the configuration.

### Key Features

- Supports granular configuration for origins, methods, headers, and credentials.
- Helps APIs comply with modern web security standards.

::: tip
Learn how to integrate this plugin into your setup in the **[Plugins Overview](../plugins/overview.md)**.
:::

## Configuration Fields

| Field                   | Type    | Description                                                           | Example Value                       |
| ----------------------- | ------- | --------------------------------------------------------------------- | ----------------------------------- |
| `allow_origins`         | Array   | List of allowed origins. Use `"*"` to allow all origins.              | `["*"]`                             |
| `allow_methods`         | Array   | List of HTTP methods allowed for cross-origin requests.               | `["GET", "POST"]`                   |
| `allow_headers`         | Array   | List of headers allowed in cross-origin requests.                     | `["Authorization", "Content-Type"]` |
| `expose_headers`        | Array   | List of headers exposed in the response.                              | `["Authorization"]`                 |
| `allow_credentials`     | Boolean | Whether credentials (cookies, authorization headers) are allowed.     | `true`                              |
| `allow_private_network` | Boolean | Whether private network requests (e.g., local IPs) are allowed.       | `false`                             |
| `preflight_continue`    | Boolean | Whether the gateway should continue processing preflight requests.    | `false`                             |
| `max_age`               | Integer | Maximum time (in seconds) for which the preflight response is cached. | `3600`                              |

::: tip
For maximum security, use specific origins instead of `"*"` in production environments.
:::

## Example Configuration

Below is an example of configuring the CORS plugin:

```json
{
  "name": "cors",
  "enabled": true,
  "config": {
    "allow_origins": ["*"],
    "allow_methods": ["GET", "POST"],
    "allow_headers": ["Authorization", "Content-Type"],
    "expose_headers": ["Authorization"],
    "allow_credentials": true,
    "allow_private_network": false,
    "preflight_continue": false,
    "max_age": 3600
  }
}
```

### Explanation

- **`allow_origins`**: Allows all origins (`"*"`). Replace with specific origins for stricter control.
- **`allow_methods`**: Permits `GET` and `POST` requests from cross-origin sources.
- **`allow_headers`**: Enables `Authorization` and `Content-Type` headers in requests.
- **`expose_headers`**: Exposes `Authorization` in the response.
- **`allow_credentials`**: Allows cookies and authorization headers in requests.
- **`allow_private_network`**: Blocks private network requests for security.
- **`preflight_continue`**: Stops further processing of preflight requests.
- **`max_age`**: Sets a 1-hour cache duration for preflight responses.

## Applying the Plugin

The CORS plugin can be applied at various levels:

1. **Global Level**: Applies CORS rules to all services and routes.
2. **Service Level**: Applies CORS rules to all routes within a service.
3. **Route Level**: Customizes CORS rules for specific routes.

Example of applying the plugin globally:

```json
{
  "name": "cors",
  "enabled": true,
  "config": {
    "allow_origins": ["https://example.com"],
    "allow_methods": ["GET"],
    "allow_headers": ["Content-Type"]
  }
}
```

::: tip
Use route-level configurations for APIs with varying CORS requirements.
:::

::: warning
Misconfigured CORS policies can expose your API to vulnerabilities. Ensure `allow_origins` is restrictive in production environments.
:::

## Use Cases

1. **Enable Cross-Origin Requests**: Allow secure communication between frontend applications and APIs.
2. **Control Sensitive Endpoints**: Restrict origins, methods, and headers for specific endpoints.
3. **Optimize Performance**: Cache preflight responses to reduce latency.

## Tips for Using the CORS Plugin

::: tip
Combine the CORS plugin with authentication plugins like JWT to secure cross-origin requests.
:::

::: tip
Regularly audit allowed origins and headers to align with your security policies.
:::

For more plugins, visit the **[Plugins Overview](../plugins/overview.md)**.
