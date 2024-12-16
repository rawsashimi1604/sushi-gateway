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

Log out users by clearing the JWT stored in an HttpOnly cookie.

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

Retrieve detailed configuration of the gateway, including certificate paths, database settings, and persistence configuration.

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
  "PersistenceConfig": "dbless",
  "PersistenceSyncInterval": 0,
  "ConfigFilePath": "./config/config.json",
  "DbConnectionHost": "localhost",
  "DbConnectionName": "sushi",
  "DbConnectionUser": "postgres",
  "DbConnectionPass": "mysecretpassword",
  "DbConnectionPort": "5432"
}
```

## Service Endpoints

### `GET /service`

Retrieve all services configured in the gateway.

#### Request

```http
GET /service HTTP/1.1
Host: localhost:8081
Cookie: token=<jwt-token>
```

### `POST /service`

Add a new service to the gateway.

#### Request

```http
POST /service HTTP/1.1
Host: localhost:8081
Content-Type: application/json
Cookie: token=<jwt-token>

{
  "name": "example-service",
  "protocol": "http",
  "host": "example.com",
  "port": 80
}
```

### `DELETE /service`

Remove an existing service by name.

#### Request

```http
DELETE /service?serviceName=example-service HTTP/1.1
Host: localhost:8081
Cookie: token=<jwt-token>
```

## Route Endpoints

### `POST /route`

Add a new route to an existing service.

#### Request

```http
POST /route HTTP/1.1
Host: localhost:8081
Content-Type: application/json
Cookie: token=<jwt-token>

{
  "service_name": "example-service",
  "route": {
    "name": "example-route",
    "path": "/example",
    "methods": ["GET"]
  }
}
```

### `DELETE /route`

Remove an existing route by name.

#### Request

```http
DELETE /route?name=example-route HTTP/1.1
Host: localhost:8081
Cookie: token=<jwt-token>
```

## Plugin Endpoints

### `POST /plugin`

Add a plugin to a global, service, or route scope.

#### Request

```http
POST /plugin HTTP/1.1
Host: localhost:8081
Content-Type: application/json
Cookie: token=<jwt-token>

{
  "scope": "service",
  "name": "example-service",
  "plugin": {
    "name": "rate_limit",
    "enabled": true,
    "config": {
      "limit_second": 10
    }
  }
}
```

### `PUT /plugin`

Update an existing plugin.

#### Request

```http
PUT /plugin HTTP/1.1
Host: localhost:8081
Content-Type: application/json
Cookie: token=<jwt-token>

{
  "scope": "service",
  "name": "example-service",
  "plugin": {
    "name": "rate_limit",
    "enabled": true,
    "config": {
      "limit_second": 20
    }
  }
}
```

### `DELETE /plugin`

Remove an existing plugin.

#### Request

```http
DELETE /plugin HTTP/1.1
Host: localhost:8081
Content-Type: application/json
Cookie: token=<jwt-token>

{
  "scope": "service",
  "name": "example-service",
  "plugin_name": "rate_limit"
}
```

## Tips for Using Endpoints

::: tip
Ensure the `token` cookie is securely stored and included in all requests requiring authentication.
:::

::: tip
Use tools like Postman, curl, or custom scripts to test and automate endpoint usage.
:::

For more information on the Admin API, refer to the **[Admin API Overview](../admin-api-overview.md)**.
