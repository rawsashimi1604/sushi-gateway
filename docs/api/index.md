# Admin REST API Overview

The Sushi Gateway Admin REST API provides an internal interface for managing and configuring your API gateway. It allows administrators to create, update, delete, and retrieve information about services, routes, plugins, and more.

The Admin REST API is designed to automate gateway management tasks, making it easier to integrate with CI/CD pipelines or other external systems.

## Key Features

- **Comprehensive Management**: Manage all entities such as services, routes, plugins, and upstreams through a RESTful interface.
- **Automation-Ready**: Integrate with automation tools for seamless configuration updates.
- **Real-Time Changes**: Apply configuration changes instantly without needing to restart the gateway.

## Accessing the Admin API

The Admin REST API is hosted on **port 8081** and communicates over HTTP. Ensure that the API is accessible only from trusted networks or through secure tunnels (e.g., VPN or SSH).

### Base URL

```
http://<gateway-host>:8081
```

## Authentication

The Admin REST API is secured through **Basic Authentication (RFC 7617)** and **JWT (RFC 7519)**. These methods ensure that only authorized users can access and modify gateway configurations. By design, the Admin API is not exposed publicly and is intended for internal use only.

### Login Workflow

1. **Login Request**: Clients authenticate via `POST /login` by sending a Base64-encoded username and password in the `Authorization` header.
2. **JWT Provision**: Upon successful authentication, the API issues a JWT stored as an HttpOnly cookie.
3. **Subsequent Requests**: The JWT cookie is included in requests to authenticate against the Admin API.

## Endpoints

Here are the endpoints available in the Admin REST API:

| Method   | Endpoint          | Description                                              |
| -------- | ----------------- | -------------------------------------------------------- |
| `POST`   | `/login`          | Login and authenticate via Basic Authentication.         |
| `DELETE` | `/logout`         | Log out by clearing the JWT cookie.                      |
| `GET`    | `/gateway`        | Retrieve the global gateway configuration (ProxyConfig). |
| `GET`    | `/gateway/config` | Retrieve the gateway environment configuration.          |
| `GET`    | `/service`        | Retrieve all services.                                   |
| `POST`   | `/service`        | Add a new service.                                       |
| `DELETE` | `/service`        | Remove a service by its name.                            |
| `POST`   | `/route`          | Add a route to an existing service.                      |
| `DELETE` | `/route`          | Remove a route by its name.                              |
| `POST`   | `/plugin`         | Add a plugin at a global, service, or route level.       |
| `DELETE` | `/plugin`         | Remove a plugin at a global, service, or route level.    |
| `PUT`    | `/plugin`         | Update a plugin at a global, service, or route level.    |

## Use Cases

1. **Automated Deployments**: Manage services, routes, and plugins programmatically during deployment pipelines.
2. **Dynamic Configuration**: Update configurations in real-time without downtime.
3. **Monitoring and Auditing**: Retrieve details of the gateway configuration for reporting or troubleshooting.

::: tip
For more detailed information on available endpoints, refer to the **[Admin API Reference](../api/endpoints.md)**.
:::
