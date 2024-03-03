# Sushi Gateway

![High Level Design](./docs/images/design.png)

## TODO LIST
- [ ] Add more tests for each plugin
- [ ] Finish up plugins
- [ ] Add "enabled" field to plugins (common middleware for all plugins)
- [ ] Add more tests for the proxy
- [ ] Dockerizing proxy
- [ ] Create kube deployment (helm)
- [ ] Admin API development and design
- [ ] UI Portal development and design
- [ ] AI Component development and design

## TLS
https/tls support has been added to the proxy, add cert and key into environment variables to use.

## Plugins
- Auth
  - **(DONE)** basic auth
  - **(DONE)** jwt (only Hs256, Rs256 tbd)
  - **(DONE)** key auth
  - oauth2
- Security
  - **(DONE)** bot_protection
  - cors
- Traffic Control
  - **(DONE)** acl
  - rate limit
  - **(DONE)** request size limit
- Logging and metrics
  - OpenTelemetry
  - http log
  - file log
  - kafka log
    
