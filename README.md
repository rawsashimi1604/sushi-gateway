# Sushi Gateway

![High Level Design](./docs/images/design.png)

## GETTING STARTED

### docker-compose

- Go to `docker` folder
- Ensure have required TLS certs
- Run `docker compose up`

## TODO LIST

- [ ] Add gateway metadata (last configuration update, total requests)
  - [ ] gateway logs middleware
- [ ] Update readme with latest architecture diagrams and logos
- [ ] Add CI github actions
- [ ] Add CD for each release to push to dockerhub
- [ ] Add more tests for each plugin
  - [x] Acl
  - [x] Basic auth
  - [x] Bot protection
  - [x] Cors
  - [ ] Jwt
  - [x] Key auth
  - [ ] Mtls
  - [x] Rate limit
  - [x] Request size limit
  - [ ] Http log
- [x] Finish up plugins
- [x] Add "enabled" field to plugins (common middleware for all plugins)
- [ ] Add validation schema for each plugin, that is validated at config file load time
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
- [x] Admin API development and design
- [ ] Load balancing to upstreams
  - [x] Round robin
  - [ ] IP hash
  - [ ] Weighted
- [ ] Configure health checks for upstreams (good to have)
- [x] UI Portal development and design
  - [ ] Update UI Portal to show services, routes, upstreams etc...
- [ ] AI Component development and design
- [x] Flatten file structure, cyclic imports

## TLS

https/tls support has been added to the proxy, add cert and key into environment variables to use.

## Plugins

- Auth
  - **(DONE)** basic auth
  - **(DONE)** jwt (only Hs256, Rs256 tbd)
  - **(DONE)** key auth
  - **(DONE)** mtls
- Security
  - **(DONE)** bot_protection
  - **(DONE)** cors
- Traffic Control
  - **(DONE)** acl
  - **(DONE)** rate limit
  - **(DONE)** request size limit
- Logging and metrics
  - **(DONE)** http log
