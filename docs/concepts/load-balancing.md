# Load Balancing

Load balancing in Sushi Gateway ensures that incoming API requests are distributed efficiently across upstream servers. It improves performance, reliability, and fault tolerance by managing how traffic is routed to backend services.

## Supported Load Balancing Algorithms

Sushi Gateway supports the following load balancing strategies:

### Round Robin

- **Description**: Distributes requests sequentially across all available upstreams.
- **Use Case**: Suitable for scenarios with equally capable upstreams and evenly distributed workloads.

### Weighted _(In Progress)_

- **Description**: Distributes requests based on predefined weights assigned to each upstream.
- **Use Case**: Ideal for scenarios where some upstreams have higher capacity or priority.

### IP Hash _(In Progress)_

- **Description**: Routes requests based on the client’s IP address, ensuring consistent upstream selection for the same client.
- **Use Case**: Useful for maintaining session persistence (sticky sessions).

## Example Configuration

Here’s an example of configuring load balancing in a `config.json` file:

```json
{
  "name": "example-service",
  "base_path": "/example",
  "protocol": "http",
  "load_balancing_strategy": "round_robin",
  "upstreams": [
    { "id": "upstream_1", "host": "example-app-1", "port": 3000 },
    { "id": "upstream_2", "host": "example-app-2", "port": 3001 }
  ]
}
```

### Explanation

- **`load_balancing_strategy`**: Defines the strategy to use (`round_robin`, `weighted`, or `ip_hash`).
- **`upstreams`**: Lists the backend servers to which requests are distributed.

::: tip
Use `round_robin` for balanced traffic distribution when all upstreams have similar capacity.
:::

## Choosing the Right Strategy

The appropriate load balancing strategy depends on your use case:

| Strategy    | Best For                                                   |
| ----------- | ---------------------------------------------------------- |
| Round Robin | Uniform upstreams with similar capabilities.               |
| Weighted    | Upstreams with varying capacity or priorities.             |
| IP Hash     | Scenarios requiring session persistence (sticky sessions). |
