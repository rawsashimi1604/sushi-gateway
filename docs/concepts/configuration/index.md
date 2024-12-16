# Configuration in Sushi Gateway

Sushi Gateway supports two primary types of configurations: **Declarative Configurations** using `config.json` files and **Environment Variable Configurations**. Each type serves specific use cases, providing flexibility for different deployment models.

## Declarative Configuration (`config.json`)

Declarative configurations allow you to define the gateway's behavior and settings in a JSON file. This approach is ideal for:

- Stateless deployments.
- GitOps workflows where configuration is managed through version control.

When using declarative configuration:

- All entities, such as services, routes, upstreams, and plugins, are defined in a single file.
- Easy to replicate and manage across environments.

::: tip
To learn more about declarative configurations, visit the **[Declarative Configuration Guide](../configuration/declarative.md)**.
:::

## Environment Variable Configuration

Environment variable configurations allow you to dynamically adjust gateway behavior on startup. Simply specify it at docker runtime as an environment variable to utilize them.

- Define database connections, persistence modes, and other runtime options.

::: tip
For a detailed list of supported environment variables, see the **[Environment Variable Configuration Guide](../configuration/environment.md)**.
:::
