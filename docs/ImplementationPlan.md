# Implementation Plan

- Develop proxy and routing capabilities

  - (DONE) Deploy HA Proxy
  - (DONE) Configure HA Proxy
  - (DONE) Load balancing
  - (DONE) Enable data plane api, in docker compose
    - https://www.haproxy.com/documentation/haproxy-data-plane-api/installation/install-on-haproxy/
    - To test: `curl -X GET --user admin:adminpwd http://localhost:5555/v2/info | jq`
    - docs: https://www.haproxy.com/documentation/dataplaneapi/community/
  - (DONE) Setup test servers to take in app id

- Ingress
  - (DONE) Send req to HA Proxy
  - (DONE) Retrieve res from HA Proxy
  - (DONE) Send back to client

- Plugin architecture
  - (DONE) Create simple plugin to append hello world to request
  - (DONE) Load plugin dynamically.
  - Add auth plugins
    - Basic auth
    - Sym JWT
    - Asym JWT
    - OAuth2
  - Add logs / analytics plugin
  - Add ratelimit

- Sushi Manager
