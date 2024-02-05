# Sushi Gateway Design Document

## Introduction and Overview

## High Level Design

![High Level Design](./images/design.png)

## API Request lifecycle

- Receive request from API client
- Forward to INGRESS HAProxy
  - Perform basic rate limiting
- INGRESS forward to sushi-proxy
  - Check service
  - Check route
  - Check consumer
  - Run plugins in configured order
    - get plugin priority (each plugin has preconfigured static priority)
      - https://docs.konghq.com/konnect/reference/plugins/
    - run the request into plugin
- Send request to EGRESS HAProxy

