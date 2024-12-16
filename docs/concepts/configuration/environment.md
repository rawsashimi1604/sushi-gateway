# Environment Variable Configuration

Environment variables in Sushi Gateway allow for dynamic configuration of the gateway without modifying files. This method is especially useful in containerized environments, enabling seamless updates and integration with external tools.

## Commonly Used Environment Variables

The following table lists the key environment variables supported by Sushi Gateway:

| Variable                    | Description                                                                     | Required | Example Value       | Default Value      |
| --------------------------- | ------------------------------------------------------------------------------- | -------- | ------------------- | ------------------ |
| `PERSISTENCE_CONFIG`        | Defines the persistence mode (`dbless` for stateless, `db` for stateful).       | Yes      | `dbless`            | `dbless`           |
| `PERSISTENCE_SYNC_INTERVAL` | The interval (in seconds) for syncing in-memory configurations to the database. | No       | `5`                 | `5`                |
| `DB_CONNECTION_HOST`        | Hostname or IP address of the database server.                                  | Yes      | `localhost`         | -                  |
| `DB_CONNECTION_PORT`        | Port number of the database server.                                             | Yes      | `5432`              | -                  |
| `DB_CONNECTION_NAME`        | Name of the database.                                                           | Yes      | `sushi`             | -                  |
| `DB_CONNECTION_USER`        | Username for the database.                                                      | Yes      | `postgres`          | -                  |
| `DB_CONNECTION_PASS`        | Password for the database user.                                                 | Yes      | `mysecretpassword`  | -                  |
| `CONFIG_FILE_PATH`          | Path to the `config.json` file for declarative configurations.                  | No       | `/app/config.json`  | `/app/config.json` |
| `ADMIN_USER`                | Username for the Admin API.                                                     | Yes      | `admin`             | `admin`            |
| `ADMIN_PASSWORD`            | Password for the Admin API user.                                                | Yes      | `changeme`          | `changeme`         |
| `SERVER_CERT_PATH`          | Path to the server certificate for HTTPS.                                       | No       | `/path/to/cert.crt` | -                  |
| `SERVER_KEY_PATH`           | Path to the server private key for HTTPS.                                       | No       | `/path/to/key.pem`  | -                  |
| `CA_CERT_PATH`              | Path to the Certificate Authority (CA) file for mutual TLS (mTLS).              | No       | `/path/to/ca.crt`   | -                  |

## Example Configuration

Here’s an example of using environment variables to configure Sushi Gateway in a Docker container:

```bash
docker run \
  --rm \
  -e PERSISTENCE_CONFIG=db \
  -e PERSISTENCE_SYNC_INTERVAL=5 \
  -e DB_CONNECTION_HOST=localhost \
  -e DB_CONNECTION_PORT=5432 \
  -e DB_CONNECTION_NAME=sushi \
  -e DB_CONNECTION_USER=postgres \
  -e DB_CONNECTION_PASS=mysecretpassword \
  -e ADMIN_USER=admin \
  -e ADMIN_PASSWORD=securepassword \
  -p 8008:8008 \
  -p 8443:8443 \
  -p 8081:8081 \
  rawsashimi/sushi-proxy:latest
```

## Tips for Using Environment Variables

::: tip
Use a `.env` file to manage environment variables for local development and testing. Load it using tools like `docker-compose` or `dotenv` libraries.
:::

::: tip
Ensure sensitive variables, such as database credentials, are stored securely using secret management tools (e.g., AWS Secrets Manager, HashiCorp Vault).
:::

::: tip
Combine environment variables with declarative configurations for maximum flexibility in hybrid setups.
:::

::: tip
Validate the values of required environment variables during startup to prevent misconfiguration.
:::

::: tip
For a deeper understanding of declarative configurations, visit the **[Declarative Configuration Guide](./files.md)**.
:::
