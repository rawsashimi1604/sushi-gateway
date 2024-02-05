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
    - get plugin priority
    - run the request into plugin
- Send request to EGRESS HAProxy

