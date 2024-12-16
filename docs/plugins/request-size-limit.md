# Request Size Limit

The Request Size Limit (`request_size_limit`) plugin ensures that incoming requests adhere to a defined maximum payload size. By enforcing size limits, this plugin helps protect APIs from large payload attacks, resource exhaustion, and unintentional misuse.

## How It Works

The plugin inspects the total request payload size in bytes and compares it against the configured `max_payload_size`. Requests exceeding the allowed size are rejected with a **413 Payload Too Large** response.

### Key Features

- Enforces maximum payload size restrictions.
- Protects against resource exhaustion and denial-of-service (DoS) attacks.
- Configurable at global, service, or route levels for granular control.

::: tip
Learn how to integrate this plugin into your setup in the **[Plugins Overview](../plugins/index.md)**.
:::

## Configuration Fields

| Field              | Type    | Description                                           | Example Value |
| ------------------ | ------- | ----------------------------------------------------- | ------------- |
| `max_payload_size` | Integer | Maximum allowed size of the request payload in bytes. | `10`          |

::: tip
Set the `max_payload_size` according to your API's expected payload size to optimize performance.
:::

## Example Configuration

Below is an example of configuring the Request Size Limit plugin:

```json
{
  "name": "request_size_limit",
  "enabled": true,
  "config": {
    "max_payload_size": 10
  }
}
```

### Explanation

- **`max_payload_size`**: Specifies the maximum allowable payload size in bytes. Requests exceeding this limit are rejected.

## Applying the Plugin

The Request Size Limit plugin can be applied at various levels:

1. **Global Level**: Applies size limits to all incoming requests at the gateway level.
2. **Service Level**: Restricts payload size for all routes within a service.
3. **Route Level**: Enforces size limits on specific routes.

Example of applying the plugin at the service level:

```json
{
  "name": "request_size_limit",
  "enabled": true,
  "config": {
    "max_payload_size": 1024
  }
}
```

::: tip
Use route-level size limits for endpoints that handle large payloads, such as file uploads.
:::

## Use Cases

1. **Prevent Resource Exhaustion**: Protect servers from processing overly large payloads.
2. **Improve API Stability**: Ensure consistent performance by rejecting oversized requests.
3. **Enhance Security**: Mitigate risks from malicious actors attempting to exploit large payload attacks.

## Tips for Using the Request Size Limit Plugin

::: tip
Combine this plugin with Rate Limiting to further secure your API against abusive traffic.
:::

For more plugins, visit the **[Plugins Overview](../plugins/index.md)**.
