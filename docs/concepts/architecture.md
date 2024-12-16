# Sushi Gateway Architecture

Sushi Gateway’s architecture is designed for flexibility, modularity, and scalability. It provides a robust API gateway solution that can handle diverse workloads and use cases. Below, we explore the core components, deployment modes, and networking of Sushi Gateway.

![Sushi Gateway Architecture](/images/architecture.png "Sushi Gateway Architecture Diagram")

## Core Components

### 1. Sushi Proxy

The Sushi Proxy is the core component responsible for:

- **Routing**: Directing requests to the appropriate upstream services based on defined configurations.
- **Load Balancing**: Distributing requests across multiple service instances using strategies like round-robin.
- **Plugin Management**: Executing configured plugins for authentication, rate limiting, and request transformation.

### 2. Sushi Manager

Sushi Manager is a web-based user interface that simplifies:

- **Configuration Management**: Visualize and modify gateway configurations.
- **Monitoring**: Track API usage and inspect logs in real-time.
- **Testing**: Quickly test routes and plugins through the interactive UI.

### 3. Data Storage (Optional in Stateful Mode)

For stateful deployments, Sushi Gateway relies on a database for persistent storage of configurations:

- **Supported Databases**: PostgreSQL, MySQL, or other compatible systems.
- **Purpose**: Ensure that gateway configurations and runtime data persist across restarts.

## Deployment Modes

### 1. Stateless Mode

Stateless mode uses declarative configuration files and does not require a database. It is ideal for:

- **Testing and Development**: Quickly prototype and validate configurations.
- **GitOps Workflows**: Manage configurations using version-controlled files.

### 2. Stateful Mode

Stateful mode persists configurations in a database and is suited for:

- **Production Deployments**: Handle large-scale systems with dynamic configurations.
- **High Availability**: Ensure data consistency and resilience.

::: tip
Learn more about the differences between Stateless and Stateful modes in the **[Data Persistence Guide](../concepts/data-persistence.md)**.
:::

## Plugins and Extensibility

Sushi Gateway adopts a modular approach with plugins to extend functionality. Plugins can:

- **Authenticate Requests**: Support JWT, API keys, MTLS, and more.
- **Rate Limit Traffic**: Control the number of requests per user or service.
- **Transform Requests and Responses**: Modify payloads, headers, or protocols.

### Plugin Scopes

Plugins can be applied at different scopes, allowing fine-grained customizability:

| Scope       | Description                                    |
| ----------- | ---------------------------------------------- |
| **Global**  | Applies policies across all services.          |
| **Service** | Applies policies to specific backend services. |
| **Route**   | Tailors policies for individual API routes.    |

::: info
Read about supported plugins and their configurations in the **[Plugins Guide](../plugins/index.md)**.
:::

## Networking and Ports

| Port | Protocol | Purpose                                                   |
| ---- | -------- | --------------------------------------------------------- |
| 8008 | HTTP     | Exposes the API Gateway for client requests.              |
| 8443 | HTTPS    | Provides secure communication using TLS.                  |
| 8081 | HTTP     | Hosts the Admin API for managing configurations.          |
| 5173 | HTTP     | Runs the web-based UI for interacting with Sushi Gateway. |

::: tip
Ensure the appropriate ports are open in your firewall or networking settings for smooth operation.
:::
