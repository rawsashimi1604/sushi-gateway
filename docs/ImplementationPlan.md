# Implementation Plan

- Develop proxy and routing capabilities

  - (DONE) Deploy HA Proxy
  - (DONE) Configure HA Proxy
  - (DONE) Load balancing
  - (DONE) Enable data plane api, in docker compose
    - https://www.haproxy.com/documentation/haproxy-data-plane-api/installation/install-on-haproxy/
    - To test: `curl -X GET --user admin:adminpwd http://localhost:5555/v2/info | jq`
  - (DONE) Setup test servers to take in app id

- Ingress
  - Send req to HA Proxy
  - Retrieve res from HA Proxy
  - Send back to client

- Plugin architecture
  - Create simple plugin to append hello world to request
  - Load plugin dynamically.

- Sushi Manager
