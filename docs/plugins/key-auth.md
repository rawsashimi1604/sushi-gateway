# API Key Authentication Plugin

The API Key Authentication (`key_auth`) plugin provides a straightforward method to secure APIs by requiring clients to include a valid API key in their requests. This plugin validates the key against the configuration to grant or deny access.

## How It Works

1. The plugin checks for the API key in the following order:
   - **Query Parameter**: The key is expected as a query parameter named `apiKey`.
   - **HTTP Header**: If not found in the query parameter, the key is then checked in the HTTP header under the `apiKey` field.
2. If a valid API key is found, the request proceeds; otherwise, it is rejected with a **401 Unauthorized** response.

### Key Features

- Supports API key validation via query parameters and headers.
- Configurable at global, service, or route levels.
- Lightweight and easy to implement for simple authentication needs.

::: tip
Learn how to integrate this plugin into your setup in the **[Plugins Overview](../plugins/index.md)**.
:::

## Configuration Fields

| Field | Type   | Description                          | Example Value |
| ----- | ------ | ------------------------------------ | ------------- |
| `key` | String | The API key required for validation. | `123456`      |

::: tip
Use strong and unique API keys to enhance security.
:::

## Example Configuration

Below is an example of configuring the API Key Authentication plugin:

```json
{
  "name": "key_auth",
  "enabled": true,
  "config": {
    "key": "123456"
  }
}
```

### Explanation

- **`key`**: Defines the API key required for requests to pass authentication.

## Applying the Plugin

The API Key Authentication plugin can be applied at various levels:

1. **Global Level**: Secures all services and routes with API key validation.
2. **Service Level**: Requires API keys for all routes within a service.
3. **Route Level**: Validates API keys for specific routes only.

Example of applying the plugin globally:

```json
{
  "name": "key_auth",
  "enabled": true,
  "config": {
    "key": "123456"
  }
}
```

::: tip
Apply the plugin at the route level for APIs requiring granular authentication control.
:::

## Use Cases

1. **Simple API Security**: Protect endpoints with a lightweight authentication mechanism.
2. **Internal APIs**: Authenticate requests from internal services using predefined keys.
3. **Development and Testing**: Use static API keys for quick and easy authentication in non-production environments.

## Example Request

### Query Parameter Example

```http
GET /api/v1/resource?apiKey=123456 HTTP/1.1
Host: example.com
```

### HTTP Header Example

```http
GET /api/v1/resource HTTP/1.1
Host: example.com
apiKey: 123456
```

## Tips for Using the API Key Authentication Plugin

::: tip
Rotate API keys periodically to maintain security and reduce risks from compromised keys.
:::

::: tip
Combine with Rate Limiting to prevent abuse of API keys.
:::

::: tip
Log unauthorized attempts to identify potential misuse or attacks.
:::

::: warning
Relying solely on API keys is not recommended for production environments. Combine with additional security mechanisms such as IP whitelisting or authentication plugins for enhanced protection.
:::

For more plugins, visit the **[Plugins Overview](../plugins/index.md)**.
