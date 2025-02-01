# Declarative Configuration File

The declarative configuration uses a JSON file to define services, routes, upstreams, and plugins in a single, centralized document. This approach is ideal for environments that benefit from a version-controlled, GitOps-friendly setup.

## Structure of Config File

The configuration file is a hierarchical structure with the following key sections:

| Section    | Description                                                            |
| ---------- | ---------------------------------------------------------------------- |
| `global`   | Global settings for the gateway, such as gateway-wide applied plugins. |
| `services` | A list of services, each containing its upstreams and routes.          |

::: tip
Learn more about services in the **[Service Documentation](../entities/service.md)**.
:::

### Example `config.json`

```json
{
  "global": {
    "name": "example-gateway",
    "plugins": []
  },
  "services": [
    {
      "name": "example-service",
      "base_path": "/example",
      "protocol": "http",
      "load_balancing_strategy": "round_robin",
      "upstreams": [
        {
          "id": "upstream_1",
          "host": "example-app",
          "port": 3000
        }
      ],
      "routes": [
        {
          "name": "example-route",
          "path": "/v1/sushi",
          "methods": ["GET"],
          "plugins": [
            {
              "name": "rate_limit",
              "enabled": true,
              "config": {
                "limit_second": 10,
                "limit_min": 100
              }
            }
          ]
        }
      ]
    }
  ]
}
```

## Key Sections Explained

### 1. Global

Defines settings that apply to the entire gateway.

Example:

```json
"global": {
  "name": "example-gateway",
  "plugins": []
}
```

::: tip
For a deeper dive into plugins, see the **[Plugins Overview](../../plugins/index.md)**.
:::

### 2. Services

Defines the backend services and their configurations.

Example:

```json
"services": [
  {
    "name": "example-service",
    "base_path": "/example",
    "protocol": "http",
    "load_balancing_strategy": "round_robin",
    "upstreams": [
      {
        "id": "upstream_1",
        "host": "example-app",
        "port": 3000
      }
    ],
    "routes": [
      {
        "name": "example-route",
        "path": "/v1/sushi",
        "methods": ["GET"],
        "plugins": []
      }
    ]
  }
]
```

::: tip
Learn how to define services and upstreams in the **[Service Documentation](../entities/service.md)** and **[Upstream Documentation](../entities/upstream.md)**.
:::

## Tips for Using Declarative Configuration

::: tip
Use version control systems like Git to manage changes to your `config.json` file.
:::

::: tip
Reference the **[Route Documentation](../entities/route.md)** for defining and managing routes within your services.
:::
