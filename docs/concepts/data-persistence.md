# Data Persistence Modes in Sushi Gateway

Sushi Gateway supports two primary data persistence modes: **Stateless** and **Stateful**. These modes determine how configurations are managed and stored, providing flexibility based on deployment needs.

## Stateless Mode

In Stateless Mode, configurations are provided as declarative files, such as `config.json`. This mode does not require a database, making it ideal for:

- Testing and development environments.
- Scenarios with fixed configurations that do not require frequent updates.

To configure Stateless Mode, set the following environment variable (`PERSISTENCE_CONFIG` will be set to `dbless` by default if not provided):

```bash
PERSISTENCE_CONFIG=dbless
```

::: tip
For step-by-step instructions on setting up Sushi Gateway in Stateless Mode, see the **[Quick Start with Docker Guide](../quick-start/docker.md)**.
:::

### Key Features

- Simple setup with configuration files.
- Lightweight and fast.
- GitOps-friendly workflows for managing version-controlled configurations.

## Stateful Mode

Stateful Mode leverages a database to persist configurations and runtime data. This mode is suitable for:

- Production environments.
- Deployments requiring dynamic updates to configurations.
- High availability and resilience.

### Database Connection Configuration

For Stateful Mode, you need to specify database connection settings and additional options using environment variables. Sushi Gateway supports PostgreSQL.

Set the following environment variables:

```bash
PERSISTENCE_CONFIG=db
PERSISTENCE_SYNC_INTERVAL=5
DB_CONNECTION_HOST=<database_host>
DB_CONNECTION_PORT=<database_port>
DB_CONNECTION_NAME=<database_name>
DB_CONNECTION_USER=<database_user>
DB_CONNECTION_PASS=<database_password>
```

- **`PERSISTENCE_SYNC_INTERVAL`**: Specifies how often (in seconds) the database should sync with the in-memory configuration.

::: tip
For step-by-step instructions on setting up Sushi Gateway with PostgreSQL, see the **[Installation Guide](../installation/install-with-postgres.md)**.
:::
