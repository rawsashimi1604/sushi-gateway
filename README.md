<a href="https://rawsashimi1604.github.io/sushi-gateway">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="./docs/public/images/LogoWithText_Dark.png">
    <source media="(prefers-color-scheme: light)" srcset="./docs/public/images/LogoWithText_Light.png">
    <img 
      alt="Sushi Gateway Logo"
      src="./docs/public/images/LogoWithText_Light.png"
      width="450">
  </picture>
</a>

<br/>
<br/>

![Stars](https://img.shields.io/github/stars/rawsashimi1604/sushi-gateway?style=flat-square) ![GitHub commit activity](https://img.shields.io/github/commit-activity/m/rawsashimi1604/sushi-gateway?style=flat-square) ![Docker Pulls](https://img.shields.io/docker/pulls/rawsashimi/sushi-proxy?style=flat-square) ![Version](https://img.shields.io/github/v/release/rawsashimi1604/sushi-gateway?color=green&label=Version&style=flat-square) ![License](https://img.shields.io/badge/License-MIT-yellow?style=flat-square)

**Sushi Gateway** is a lightweight and extensible Layer 7 API Gateway designed to empower developers with seamless control, robust security, and dynamic adaptability.

By providing functionality such as routing, load balancing, authentication, and plugins, Sushi Gateway enables effortless orchestration of microservices and APIs.

---

[Installation Guide](https://rawsashimi1604.github.io/sushi-gateway/getting-started/docker.html) | [Documentation](https://rawsashimi1604.github.io/sushi-gateway/docs-home.html) | [Releases](https://github.com/rawsashimi1604/sushi-gateway/releases)

---

## Roadmap

See the [Roadmap](ROADMAP.md) for more information on the project's future direction.

## Features

- **Dynamic Routing**: Route traffic efficiently with support for dynamic paths and advanced match criteria.
- **Plugin System**: Extend functionality with modular plugins for security, rate limiting, logging, and more.
- **Load Balancing**: Built-in strategies like round robin, weighted (in progress), and IP hash.
- **Declarative Configuration**: Use declarative JSON configurations to configure the gateway.
- **Secure API Management**: Features such as Mutual TLS, API key authentication, and JWT support.
- **Lightweight & Efficient**: Optimized for speed and scalability with a minimal footprint.

## Quick Start

View our quick start guide using Docker [here](https://rawsashimi1604.github.io/sushi-gateway/getting-started/docker.html).

## Plugins

Sushi Gateway offers a wide range of plugins to enhance functionality, including:

| Plugin Name              | Description                           |
| ------------------------ | ------------------------------------- |
| **Rate Limit**           | Limit requests per user/service.      |
| **CORS**                 | Enable cross-origin resource sharing. |
| **JWT Authentication**   | Secure APIs with JWT tokens.          |
| **Basic Authentication** | Simplify API access control.          |
| **Bot Protection**       | Block malicious user agents.          |

Explore all available plugins in the **[Plugins Documentation](https://rawsashimi1604.github.io/sushi-gateway/plugins)**.

## Contributing

We ❤️ contributions! Check out the [Contributing Guide](CONTRIBUTING.md) to get started.

- **Join the Community**: Share feedback and ask questions in our [Discussions](https://github.com/rawsashimi1604/sushi-gateway/discussions).
- **Chat on Discord**: Collaborate and discuss on our [Discord Channel](https://discord.gg/aPv4QhQ6).
- **Report Issues**: Open issues directly in the [GitHub repository](https://github.com/rawsashimi1604/sushi-gateway/issues).

[sushi-url]: https://rawsashimi1604.github.io/sushi-gateway/

## Building Sushi Gateway

- Ensure you have Go installed with at least version 1.22.
- Look up the [Quick Start](https://rawsashimi1604.github.io/sushi-gateway/getting-started/docker.html) guide to get the config file, certs and keys required for the gateway to run.
- Create the certs and keys required for the gateway to run.
- Create the config file for the gateway to run.
- Build and run the gateway using the following commands:
  - `go run cmd/main.go`
- To run tests use the following command:
  - `go test ./...`
- Use the `docker-compose.yml` file to quickly start up dev servers to test proxy
  - `docker compose up -d`
