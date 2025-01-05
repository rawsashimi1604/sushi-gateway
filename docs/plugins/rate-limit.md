# Rate Limiting Plugin

The Rate Limiting (`rate_limit`) plugin is used to control the number of API requests a client can make over a specified period. This ensures fair usage, protects upstream services from excessive load, and enhances security.

## How It Works

The Rate Limiting plugin applies rules to incoming requests based on a defined configuration using the [token bucket algorithm](https://en.wikipedia.org/wiki/Token_bucket). It tracks requests and rejects those that exceed the allowed limits with a **429 Too Many Requests** response.

### Key Features

- Limit requests per second, minute, or hour.
- Flexible configuration for global, service, or route-specific limits.

::: tip
Learn how to integrate this plugin into your setup in the **[Plugins Overview](../plugins/index.md)**.
:::

## Configuration Fields

| Field          | Type    | Description                                    | Example Value |
| -------------- | ------- | ---------------------------------------------- | ------------- |
| `limit_second` | Integer | Maximum number of requests allowed per second. | `10`          |
| `limit_min`    | Integer | Maximum number of requests allowed per minute. | `100`         |
| `limit_hour`   | Integer | Maximum number of requests allowed per hour.   | `1000`        |

## Example Configuration

Below is an example of configuring the Rate Limiting plugin for a route:

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

- **`limit_second`**: Allows up to 10 requests per second.
- **`limit_min`**: Allows up to 100 requests per minute.
- **`limit_hour`**: Allows up to 1000 requests per hour.
- **`message`**: Custom response message returned when limits are breached.

## Applying the Plugin

The Rate Limiting plugin can be applied at various levels:

1. **Global Level**: Affects all incoming requests to the gateway.
2. **Service Level**: Affects all routes under a specific service.
3. **Route Level**: Affects only the specified route.

Example of applying the plugin to a specific route:

```json
{
  "name": "example-route",
  "path": "/v1/sushi",
  "methods": ["GET"],
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

## Use Cases

1. **Prevent Abuse**: Protect APIs from excessive requests by malicious users.
2. **Traffic Shaping**: Ensure fair usage by distributing requests evenly across clients.
3. **Upstream Protection**: Prevent backend services from being overwhelmed by traffic.

For more plugins, visit the **[Plugins Overview](../plugins/index.md)**.
