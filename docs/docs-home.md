# Documentation

Welcome to the **Sushi Gateway** documentation! It is meant to serve as a reference resource for all things Sushi! Find all the information you need to use Sushi Gateway here.

## What is Sushi Gateway?

Sushi Gateway serves as a **Layer 7 API Gateway** that simplifies API traffic management by handling:

::: info
An API Gateway acts as a reverse proxy that sits between clients and backend services, handling tasks such as request routing, security enforcement, and load balancing. For a deeper dive into what an API Gateway is, check out our [What is an API Gateway?](./concepts/what-is-api-gateway.md) page.
:::

- **Request Routing**: Route requests to upstream services based on defined configurations.
- **Security**: Enforce robust security policies, including authentication, authorization, and rate limiting.
- **Load Balancing**: Efficiently distribute traffic across multiple instances to ensure system reliability.
- **Extensibility**: Customize and extend the gateway using modular plugins tailored to your specific needs.

Sushi Gateway is designed to be simple, and flexible, making it an ideal solution for users who want to mordenize their API infrastructure.

## Why Use Sushi Gateway?

Modern applications often rely on APIs for communication between services, clients, and third-party integrations. However, managing APIs in complex systems can present several challenges:

- **Security Risks**: APIs can expose sensitive data and operations.
- **Scaling Issues**: Handling high-traffic scenarios without affecting performance.
- **Operational Overhead**: Configuring and monitoring APIs across distributed systems.

Sushi Gateway addresses these challenges with:

### Key Features

- **Modular Policy Architecture**
  - Configure API plugins/policies at the global, service, or route scope for fine grained tuning.
- **High Performance**
  - Built in **Golang**, optimized for concurrency and low latency.
- **Declarative Configuration**
  - Stateless in-memory configuration makes for easy configuration and startup, reducing dependencies on a database.
- **Comprehensive Security**
  - Includes authentication (JWT, Basic Auth, API Keys), rate limiting, and CORS policies.
- **Developer-Friendly**
  - RESTful Admin API and intuitive UI for managing configurations.

## Architecture Overview

### Components

Sushi Gateway comprises two primary components:

1. **Sushi Proxy**
   - Core gateway component handling request routing, load balancing, and applying plugins.
2. **Sushi Manager**
   - Web-based UI for monitoring and managing gateway configurations.


::: tip
Try Sushi Gateway witt docker for an easy and fast setup. It uses declarative configuration files, making it perfect for testing or quick deployment scenarios.
[Get started with our Quick Start Guide now](./getting-started/docker.md)!
:::

## Supported Platforms

- Docker
- Kubernetes via Helm (Coming soon!)
