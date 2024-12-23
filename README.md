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

![Stars](https://img.shields.io/github/stars/rawsashimi1604/sushi-gateway?style=flat-square) ![GitHub commit activity](https://img.shields.io/github/commit-activity/m/rawsashimi1604/sushi-gateway?style=flat-square) ![Docker Pulls](https://img.shields.io/docker/pulls/rawsashimi/sushi-proxy?style=flat-square) ![Version](https://img.shields.io/github/v/release/rawsashimi1604/sushi-gateway?color=green&label=Version&style=flat-square) ![License](https://img.shields.io/badge/License-Apache%202.0-blue?style=flat-square)

**Sushi Gateway** is a lightweight, high-performance, and extensible API Gateway designed to empower developers with seamless control, robust security, and dynamic adaptability.

By providing functionality such as routing, load balancing, authentication, and plugins, Sushi Gateway enables effortless orchestration of microservices and APIs.

---

[Installation Guide](https://rawsashimi1604.github.io/sushi-gateway/getting-started/docker.html) | [Documentation](https://rawsashimi1604.github.io/sushi-gateway/docs-home.html) | [Releases](https://github.com/rawsashimi1604/sushi-gateway/releases)

---

## Features

- **Dynamic Routing**: Route traffic efficiently with support for dynamic paths and advanced match criteria.
- **Plugin System**: Extend functionality with modular plugins for security, rate limiting, logging, and more.
- **Load Balancing**: Built-in strategies like round robin, weighted (in progress), and IP hash (in progress).
- **Stateless and Stateful Modes**: Choose between declarative JSON configurations or database-backed persistence.
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

## Start a database

create a postgres database.

- `docker run --name postgres-db -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres`
- `docker exec -it postgres-db psql -U postgres -d sushi`
- `CREATE DATABASE sushi;`
- `\c sushi`
- `docker cp init.sql postgres-db:/init.sql`
- `docker exec -it postgres-db psql -U postgres -d sushi -f /init.sql`
- `docker cp mock.sql postgres-db:/mock.sql`
- `docker exec -it postgres-db psql -U postgres -d sushi -f /mock.sql`
- `\c`
