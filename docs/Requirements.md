# General Requirements

- Routing: Route incoming requests to appropriate upstream APIs.
- Load Balancing: Distribute incoming requests evenly across upstream servers.
  - Ability to toggle load balancing on/off.
  - Support for both round-robin and least-connections strategies.
- Protocol Support: Handle both HTTP and TCP traffic.
  - Option for TLS termination at the gateway or pass-through to the service.
- Logging: Capture and store logs for all requests and responses.
  - Exportable log files in common formats (e.g., JSON, CSV).
- Health Checks: Periodically verify the health of upstream APIs.
  - Configurable intervals and conditions for health checks.
- Access Control: Manage access through an Access Control List (ACL).
  - Support for both whitelist and blacklist configurations.

# Service Management

- Service Registry: Maintain a catalog of all available services.
- Service Modification: Add, update, or remove upstream API services.
- Support for service versioning.

# Route Management

- Route Registry: Maintain a catalog of all available routes.
- Route Modification: Add, update, or remove routes for upstream services.
- Define methods, paths, and supported protocols for each route.

# Consumer Management

- Consumer Registry: Maintain a list of consumers allowed to access the API.
- Consumer Provisioning: Add, update, or remove consumers.
- Assign API keys or credentials to consumers.

# Authentication and Authorization

- Authentication Toggle: Ability to enable or disable authentication.
- JWT Support: Provision private/public key pairs for JWT authentication.
  - Choice of algorithms (e.g., RS256, ES256).
- Secure exposure of public keys to consumers while protecting private keys.
- Credential Management: List, provision, and revoke JWT credentials.
- Authentication Plugin: Authenticate requests using JWT tokens.
- Support for receiving JWTs via cookies, query strings, or HTTP headers.

# Sushi Manager (API Portal UI)

- User Interface: Develop a web-based UI for managing the API gateway.
- Configuration: Provide interfaces for adjusting gateway settings.
- Documentation: Incorporate API documentation and interactive exploration tools.
- Monitoring: Display real-time metrics and logs for the gateway and services.
- Alerts and Notifications: Configure alerts for system incidents and threshold breaches.

# Security

- Encryption: Enforce SSL/TLS for secure communication.
- Rate Limiting: Throttle requests to prevent abuse.
- IP Filtering: Allow or block requests based on IP addresses.
- DDoS Protection: Implement measures to mitigate distributed denial-of-service attacks.

# Observability and Monitoring

- Real-time Metrics: Gather and display real-time data on traffic, performance, and errors.
- Alerting System: Trigger alerts based on predefined criteria or anomalies.
- Log Analysis: Tools for searching and analyzing log data.

# Deployment and Operations

- High Availability: Design the gateway for fault tolerance and zero downtime.
- Configuration Management: Provide mechanisms for dynamic configuration changes.
- Scalability: Ensure the gateway can scale horizontally to handle increased load.
- CI/CD Integration: Allow for continuous integration and deployment processes.

# Developer Experience

- CLI Tooling: Offer a command-line interface for developers to interact with the gateway.

# Documentation and Support

Technical Documentation: Provide comprehensive guides and API documentation.
