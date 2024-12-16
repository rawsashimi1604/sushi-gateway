# Basic Authentication Plugin

The Basic Authentication (`basic_auth`) plugin secures APIs by requiring users to provide a username and password. This plugin is compliant with **[RFC 7617](https://datatracker.ietf.org/doc/html/rfc7617)** and ensures that only authorized users can access your services.

## How It Works

The Basic Authentication plugin validates incoming requests by checking the provided credentials against a predefined username and password. Requests with invalid or missing credentials are rejected with a **401 Unauthorized** response.

### Key Features

- Simple and lightweight authentication method.
- Configurable at global, service, or route levels.

::: tip
Learn how to integrate this plugin into your setup in the **[Plugins Overview](../plugins/overview.md)**.
:::

## Configuration Fields

| Field      | Type   | Description                                | Example Value |
| ---------- | ------ | ------------------------------------------ | ------------- |
| `username` | String | The username required for authentication.  | `admin`       |
| `password` | String | The password associated with the username. | `adminpass`   |

::: tip
Ensure passwords are strong and securely stored to enhance security.
:::

## Example Configuration

Below is an example of configuring the Basic Authentication plugin for a route:

```json
{
  "name": "basic_auth",
  "enabled": true,
  "config": {
    "username": "admin",
    "password": "adminpass"
  }
}
```

### Explanation

- **`username`**: The username required for accessing the route or service.
- **`password`**: The password associated with the username.

## Applying the Plugin

The Basic Authentication plugin can be applied at various levels:

1. **Global Level**: Secures all incoming requests to the gateway.
2. **Service Level**: Secures all routes within a specific service.
3. **Route Level**: Secures individual routes for precise control.

Example of applying the plugin to a specific route:

```json
{
  "name": "example-route",
  "path": "/v1/sushi",
  "methods": ["GET"],
  "plugins": [
    {
      "name": "basic_auth",
      "enabled": true,
      "config": {
        "username": "admin",
        "password": "adminpass"
      }
    }
  ]
}
```

::: tip
Use route-level Basic Authentication for APIs that require granular access control.
:::

## Connecting to the API

To authenticate requests using Basic Authentication:

1. Base64 encode the `username:password` pair.
2. Include the encoded value in the `Authorization` header as:
   ```http
   Authorization: Basic <base64-encoded-credentials>
   ```

### Example:

If the username is `admin` and the password is `adminpass`, the header would look like:

```http
Authorization: Basic YWRtaW46YWRtaW5wYXNz
```

::: warning
**Basic Authentication is insufficient to protect production environments.**

Reasons:

- **Lack of Encryption**: Credentials are transmitted in plaintext unless HTTPS is used, making them vulnerable to interception.
- **No Token Revocation**: Unlike modern authentication mechanisms (e.g., OAuth), credentials cannot be easily revoked or rotated.
- **Prone to Brute Force Attacks**: Without additional security measures like rate limiting, attackers can guess credentials.
- **Static Credentials**: Managing static username-password pairs at scale is cumbersome and insecure.

Consider using advanced authentication methods like JWT for production-grade security.
:::

## Use Cases

1. **Restrict Access**: Secure endpoints by limiting access to authorized users.
2. **Protect Development APIs**: Add a simple authentication layer for staging or development environments.

## Tips for Using the Basic Authentication Plugin

::: tip
Use HTTPS to encrypt traffic and protect credentials from being intercepted.
:::

::: tip
Combine Basic Authentication with plugins like Rate Limiting to prevent brute-force attacks.
:::

::: tip
Store and manage credentials securely using environment variables or secret management tools.
:::

For more plugins, visit the **[Plugins Overview](../plugins/overview.md)**.
