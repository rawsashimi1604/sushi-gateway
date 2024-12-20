![Sushi Gateway](./docs/public/images/LogoWithText2RemoveBg.png)

![Stars](https://img.shields.io/github/stars/rawsashimi1604/sushi-gateway?style=flat-square) ![GitHub commit activity](https://img.shields.io/github/commit-activity/m/rawsashimi1604/sushi-gateway?style=flat-square) ![Docker Pulls](https://img.shields.io/docker/pulls/rawsashimi/sushi-proxy?style=flat-square) ![Version](https://img.shields.io/github/v/release/rawsashimi1604/sushi-gateway?color=green&label=Version&style=flat-square) ![License](https://img.shields.io/badge/License-Apache%202.0-blue?style=flat-square)

**Sushi Gateway** is a lightweight, high-performance, and extensible API Gateway designed to empower developers with seamless control, robust security, and dynamic adaptability.

By providing functionality such as routing, load balancing, authentication, and plugins, Sushi Gateway enables effortless orchestration of microservices and APIs.

---

[Installation Guide](https://rawsashimi1604.github.io/sushi-gateway/quick-start) | [Documentation](https://rawsashimi1604.github.io/sushi-gateway/) | [Releases](https://github.com/rawsashimi1604/sushi-gateway/releases)

---

## Features

- **Dynamic Routing**: Route traffic efficiently with support for dynamic paths and advanced match criteria.
- **Plugin System**: Extend functionality with modular plugins for security, rate limiting, logging, and more.
- **Load Balancing**: Built-in strategies like round robin, weighted (in progress), and IP hash (in progress).
- **Stateless and Stateful Modes**: Choose between declarative JSON configurations or database-backed persistence.
- **Secure API Management**: Features such as Mutual TLS, API key authentication, and JWT support.
- **Lightweight & Efficient**: Optimized for speed and scalability with a minimal footprint.

---

## Quick Start

View our quick start guide using Docker [here](https://rawsashimi1604.github.io/sushi-gateway/getting-started/docker.html).

---

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

---

## Contributing

We ❤️ contributions! Check out the [Contributing Guide](CONTRIBUTING.md) to get started.

- **Join the Community**: Share feedback and ask questions in our [Discussions](https://github.com/rawsashimi1604/sushi-gateway/discussions).
- **Report Issues**: Open issues directly in the [GitHub repository](https://github.com/rawsashimi1604/sushi-gateway/issues).

---

## License

```
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   https://www.apache.org/licenses/LICENSE-2.0
```

[sushi-url]: https://rawsashimi1604.github.io/sushi-gateway/

## TODO LIST

- [x] design database schema
- [x] provide database configuration options ( env variables to inject database env in )
- [x] sushi manager update logo
- [x] Logout endpoint to delete httponly cookie
- [x] sushi manager update gateway state to get types and also domain object to retrieve from gateway state.
- [x] sushi manager create screens
- [ ] Add gateway metadata (last configuration update, total requests)
  - [ ] gateway logs middleware
- [ ] Add stateful gateway configurations (good to have)
  - [x] Postgres DB (Externalise option)
  - [x] Postgres DB docker
  - [x] Admin API for CRUD operations
  - [x] global domain object config state retrieval from db.
  - [ ] add time created and time updated to schema.
  - [x] add simple gateway table to store gateway configurations.
- [x] Update readme with latest architecture diagrams and logos
- [x] Add CI github actions
- [x] Add CD for each release to push to dockerhub
- [ ] Add more tests for each plugin
  - [x] Acl
  - [x] Basic auth
  - [x] Bot protection
  - [x] Cors
  - [x] Jwt
  - [x] Key auth
  - [ ] Mtls
  - [x] Rate limit
  - [x] Request size limit
  - [ ] Http log
- [x] Finish up plugins
- [x] Add "enabled" field to plugins (common middleware for all plugins)
- [ ] Add validation schema for each plugin, that is validated at config file load time (good to have)
  - [ ] General architecture
  - [ ] Acl
  - [ ] Basic auth
  - [ ] Bot protection
  - [ ] Cors
  - [ ] Jwt
  - [ ] Key auth
  - [ ] Mtls
  - [ ] Rate limit
  - [ ] Request size limit
  - [ ] Http log
- [x] Add dynamic routing (route parameters like {id}, {anything})
- [ ] Add more tests for the proxy
- [x] Dockerizing proxy
- [ ] Create kube deployment (helm)
- [ ] Load balancing to upstreams
  - [x] Round robin
  - [ ] IP hash
  - [ ] Weighted
- [ ] Configure health checks for upstreams (good to have)
- [x] UI Portal development and design
  - [x] Update UI Portal to show services, routes, upstreams etc...
  - [ ] Update UI Portal to interface with Admin API
- [x] Flatten file structure, cyclic imports

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
