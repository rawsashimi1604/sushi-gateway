# Admin API Endpoints

The Admin REST API in Sushi Gateway offers a comprehensive suite of endpoints to manage and configure the gateway's services, routes, plugins, and more. Below is an overview of the available endpoints and their functionality.

## Authentication Endpoints

### `POST /login`

Authenticate users via Basic Authentication and provision a JWT for subsequent requests.

#### Request

```http
POST /login HTTP/1.1
Host: localhost:8081
Authorization: Basic <base64-encoded-credentials>
```

#### Response

```http
HTTP/1.1 200 OK
Set-Cookie: token=<jwt-token>; HttpOnly; Secure
```

### `DELETE /logout`

Log out users by clearing the JWT stored in an `HttpOnly` cookie.

#### Request

```http
DELETE /logout HTTP/1.1
Host: localhost:8081
Cookie: token=<jwt-token>
```

## Gateway Configuration Endpoints

### `GET /gateway`

Retrieve the global gateway configuration (ProxyConfig), including services, routes, and plugins.

#### Request

```http
GET /gateway HTTP/1.1
Host: localhost:8081
Cookie: token=<jwt-token>
```

#### Response

```json
{
  "services": [
    {
      "name": "example-service",
      "routes": [{ "name": "example-route", "path": "/example" }]
    }
  ],
  "plugins": [{ "name": "rate_limit", "config": { "limit_second": 10 } }]
}
```

### `GET /gateway/config`

Retrieve detailed configuration of the gateway, including certificate paths and environment variables.

#### Request

```http
GET /gateway/config HTTP/1.1
Host: localhost:8081
Cookie: token=<jwt-token>
```

#### Response

```json
{
  "ServerCertPath": "./config/server.crt",
  "ServerKeyPath": "./config/server.key",
  "CACertPath": "./config/ca.crt",
  "AdminUser": "admin",
  "AdminPassword": "changeme",
  "ConfigFilePath": "./config/config.json",
}
```

## Tips for Using Endpoints

::: tip
Ensure the `token` cookie is securely stored and included in all requests requiring authentication.
:::

::: tip
Use tools like Postman, curl, or custom scripts to test and automate endpoint usage.
:::

For more information on the Admin API, refer to the **[Admin API Overview](./index.md)**.
