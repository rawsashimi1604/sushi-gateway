# Implementation Plan

- Develop proxy and routing capabilities

  - (DONE) Deploy HA Proxy
  - (DONE) Configure HA Proxy
  - (DONE) Load balancing
  - Enable data plane api, in docker compose
  - (DONE) Setup test servers to take in app id

- Ingress

  - Send req to HA Proxy
  - Retrieve res from HA Proxy
  - Send back to client

- Sushi Manager
