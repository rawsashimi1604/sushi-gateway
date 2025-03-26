# What is an API Gateway?

An **API Gateway** is a server that manages requests to your APIs. The gateway acts as the entry point for clients and ensures seamless communication between clients and backend services. By using an API Gateway, API management is simplified when handling key tasks like authentication, load balancing, and rate limiting.

## Key Features of an API Gateway

- **Request Routing**: Directs incoming requests to the configured backend services.
- **Security**: Enforces policies or plugins for authentication, authorization, rate limiting and more!
- **Load Balancing**: Distributes traffic across multiple service instances to improve availability and reliability. This ties into the concept of horizontal scaling.
- **Request/Response Transformation**: Modifies requests and responses as needed, such as converting legacy protocols like SOAP to REST.
- **Monitoring and Logging**: Help track API usage and logs traffic for analytics, auditing troubleshooting.

## Why Use an API Gateway?

Modern applications rely on multiple services and APIs that need to work together. An API Gateway simplifies this complexity by:

1. **Centralizing Management**: It provides a single layer to manage all API interactions.
2. **Enhancing Security**: It applies global security policies to protect your APIs from threats.
3. **Simplifying Client Interactions**: It consolidates multiple APIs into a single endpoint for clients.

## How Sushi Gateway Implements an API Gateway

Sushi Gateway is a lightweight and modular API Gateway that includes:

- Load configurations from a declarative configuration file.
- Configurable plugins for dynamic request processing.
- Support for authentication methods like JWT, API keys, and MTLS.

::: tip Try Sushi Gateway out now!
Get started quickly with Sushi Gateway by following the **[Quick Start Guide](../getting-started/docker.md)** to set up using Docker.
:::

::: details Architecture Overview
Learn more about Sushi Gateway's architecture and features in the **[Architecture Overview](../concepts/architecture.md)** section!
:::

---

We will continue to update this section with more examples and use cases for API Gateways.
