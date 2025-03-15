# Environment Variable Configuration

Environment variables in Sushi Gateway allow for dynamic configuration of the gateway without modifying files. This method is especially useful in containerized environments, enabling seamless updates and integration with external tools.

## Commonly Used Environment Variables

The following table lists the key environment variables supported by Sushi Gateway:

| Variable            | Description                                                                                                         | Required | Example Value         | Default Value      |
| ------------------- | ------------------------------------------------------------------------------------------------------------------- | -------- | --------------------- | ------------------ |
| `CONFIG_FILE_PATH`  | Path to the `config.json` file for declarative configurations.                                                      | No       | `/app/config.json`    | `/app/config.json` |
| `ADMIN_USER`        | Username for the Admin API.                                                                                         | Yes      | `admin`               | `admin`            |
| `ADMIN_PASSWORD`    | Password for the Admin API user.                                                                                    | Yes      | `changeme`            | `changeme`         |
| `ADMIN_CORS_ORIGIN` | CORS origin for the Admin API.                                                                                      | No       | `my.admin.api.domain` | `localhost:5173`   |
| `SERVER_CERT_PATH`  | Path to the server certificate for HTTPS. Server certs will be self generated when not provided.                    | No       | `/path/to/cert.crt`   | -                  |
| `SERVER_KEY_PATH`   | Path to the server private key for HTTPS. Server certs will be self generated when not provided                     | No       | `/path/to/key.pem`    | -                  |
| `CA_CERT_PATH`      | Path to the Certificate Authority (CA) file for mutual TLS (mTLS). Required for MTLS plugin related configurations. | No       | `/path/to/ca.crt`     | -                  |

## Example Configuration

Here’s an example of using environment variables to configure Sushi Gateway in a Docker container:

```bash
docker run \
  --rm \
  -e ADMIN_USER=admin \
  -e ADMIN_PASSWORD=securepassword \
  -p 8008:8008 \
  -p 8443:8443 \
  -p 8081:8081 \
  rawsashimi/sushi-proxy:0.4.0
```

::: tip
For a deeper understanding of declarative configurations, visit the **[Declarative Configuration Guide](./files.md)**.
:::

::: tip
Ensure sensitive variables are stored securely using secret management tools (e.g., AWS Secrets Manager, HashiCorp Vault).
:::

::: tip
For more information on TLS configuration, chewk out the **[Configuring TLS Guide](../tls.md)**.
:::
