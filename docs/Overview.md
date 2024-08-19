# Sushi Gateway Design Document

## Introduction and Overview

## High Level Design

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

- Set of tasks maybe:
  - Api Gateway
    - non blocking io
    - protocol buffer
    - 40% of features, basic gateway
    - Logs aggregation
  - AI detection feature
    - get data
    - LLM model: llama
    - salt security
    - no name security
    - classify based on warning level : HIGH, MID, LOW
      - if HIGH, MID send to event service...
      - Compartmentalize data into different buckets
      - Block request if happens...
    - AI Plugin -> send to llama but in a non blocking way...

  