# Upstream

An **Upstream** in Sushi Gateway represents a backend service or server to which API requests are forwarded. Upstreams are part of a service configuration and help distribute traffic across multiple instances.

## Fields in an Upstream

An Upstream configuration consists of the following fields:

| Field  | Type   | Description                                               |
| ------ | ------ | --------------------------------------------------------- |
| `id`   | String | A unique identifier for the upstream.                     |
| `host` | String | The hostname or IP address of the backend server.         |
| `port` | Number | The port number on which the backend server is listening. |

::: tip
Each upstream must have a unique `id` to avoid conflicts during configuration.
:::

## Example Configuration

Here’s an example of an upstream definition in `config.json`:

```json
{
  "id": "upstream_1",
  "host": "example-upstream",
  "port": 3000
}
```

### Key Fields Explained

1. **`id`**:

   - A unique identifier for the upstream.
   - Example: `"upstream_1"`.

2. **`host`**:

   - Specifies the hostname or IP address of the backend service.
   - Example: `"example-upstream"`.

3. **`port`**:
   - Defines the port number where the backend service is running.
   - Example: `3000`.

## Relationships with Other Entities

Upstreams are closely associated with the following entities:

- **[Service](../entities/service.md)**: Services define the upstreams used for routing traffic.
- **[Route](../entities/route.md)**: Routes determine how requests are directed to upstreams within a service.

## Load Balancing

Upstreams are a critical component in load balancing strategies, such as:

- **Round Robin**: Distributes traffic equally across all upstreams.
- **Weighted** _(in progress)_: Routes traffic based on predefined weights for each upstream.
- **IP Hash** _(in progress)_: Routes traffic based on the client IP address, ensuring consistent upstream selection.

::: tip
Learn more about load balancing strategies in the **[Load Balancing Concepts](../load-balancing.md)** page.
:::

---

The Upstream entity is essential for connecting Sushi Gateway to backend services. To understand other entities, see:

- **[Service](../entities/service.md)**
- **[Route](../entities/route.md)**
- **[Plugin](../entities/plugin.md)**
