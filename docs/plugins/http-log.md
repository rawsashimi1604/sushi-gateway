# HTTP Log Plugin

The HTTP Log plugin (`http_log`) enables logging of API traffic by forwarding request and response metadata to a specified HTTP endpoint. This plugin is essential for monitoring, analytics, and troubleshooting.

## How It Works

The HTTP Log plugin captures details of incoming requests and outgoing responses, then sends the information to a configured HTTP endpoint via a specified method (e.g., `POST`). The payload is customizable based on your logging requirements.

### Key Features

- Logs request and response metadata.
- Configurable HTTP endpoint, method, and content type.
- Supports JSON payloads for structured logging.

::: tip
Learn how to integrate this plugin into your setup in the **[Plugins Overview](../plugins/index.md)**.
:::

## Configuration Fields

| Field           | Type   | Description                                       | Example Value                  |
| --------------- | ------ | ------------------------------------------------- | ------------------------------ |
| `http_endpoint` | String | The URL of the endpoint to which logs are sent.   | `http://localhost:3000/v1/log` |
| `method`        | String | The HTTP method used to send logs (e.g., `POST`). | `POST`                         |
| `content_type`  | String | The content type of the log payload.              | `application/json`             |

::: tip
Ensure the `http_endpoint` is accessible and designed to handle incoming log data efficiently.
:::

## Example Configuration

Below is an example of configuring the HTTP Log plugin:

```json
{
  "name": "http_log",
  "enabled": true,
  "config": {
    "http_endpoint": "http://localhost:3000/v1/log",
    "method": "POST",
    "content_type": "application/json"
  }
}
```

### Explanation

- **`http_endpoint`**: The destination where log data is sent.
- **`method`**: Specifies `POST` as the HTTP method for sending logs.
- **`content_type`**: Defines the format of the log payload as JSON.

## Example Log Payload

Here is an example of a log payload in JSON format:

```json
{
  "service": {
    "name": "example-service",
    "protocol": "http",
    "host": "example-app",
    "port": 3000
  },
  "route": {
    "path": "/v1/sushi"
  },
  "request": {
    "protocol": "HTTP/1.1",
    "tls": false,
    "method": "GET",
    "path": "/v1/sushi",
    "url": "http://localhost:8008/v1/sushi",
    "uri": "/v1/sushi",
    "size": 123,
    "headers": {
      "Authorization": "Bearer <token>",
      "Content-Type": "application/json"
    }
  },
  "client_ip": "192.168.1.1",
  "started_at": "2024-12-15T12:34:56.789Z"
}
```

### Explanation of Fields

- **`service`**: Metadata about the service handling the request, including name, protocol, host, and port.
- **`route`**: Information about the route being accessed.
- **`request`**: Details of the incoming request, such as method, URL, headers, and size.
- **`client_ip`**: IP address of the client making the request.
- **`started_at`**: Timestamp indicating when the request was received.

## Applying the Plugin

The HTTP Log plugin can be applied at various levels:

1. **Global Level**: Logs all traffic passing through the gateway.
2. **Service Level**: Logs traffic for specific services.
3. **Route Level**: Logs traffic for individual routes.

Example of applying the plugin globally:

```json
{
  "name": "http_log",
  "enabled": true,
  "config": {
    "http_endpoint": "http://localhost:3000/v1/log",
    "method": "POST",
    "content_type": "application/json"
  }
}
```

::: tip
Use route-level logging for critical APIs requiring detailed monitoring.
:::

## Use Cases

1. **API Monitoring**: Track API usage and performance metrics.
2. **Troubleshooting**: Log errors and analyze traffic patterns to identify issues.
3. **Analytics**: Gather insights on API traffic for reporting and optimization.

## Tips for Using the HTTP Log Plugin

::: tip
Ensure that the logging endpoint can handle high volumes of traffic without impacting API performance.
:::

::: tip
Secure the logging endpoint using authentication and encryption to protect sensitive log data.
:::

::: tip
Monitor and rotate logs regularly to prevent storage overflow.
:::

For more plugins, visit the **[Plugins Overview](../plugins/index.md)**.
