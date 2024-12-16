# Mutual TLS (mTLS) Plugin

The Mutual TLS (mTLS) (`mtls`) plugin ensures secure communication by requiring both the client and server to authenticate each other during the TLS handshake. This plugin adds an additional layer of security by verifying client certificates against a trusted Certificate Authority (CA).

## How It Works

The mTLS plugin enforces a two-way TLS handshake:

1. The client presents a valid certificate to the server.
2. Sushi Gateway validates the client certificate against the configured CA certificate.
3. If validation succeeds, the connection is established; otherwise, it is rejected.

This mechanism ensures that only authorized clients can connect to the API.

### Key Features

- Verifies client certificates for secure communication.
- Enforces mTLS requirements at the gateway level.
- Uses a configurable CA certificate for validation.

::: tip
Learn how to integrate this plugin into your setup in the **[Plugins Overview](../plugins/overview.md)**.
:::

## Configuration Requirements

The mTLS plugin relies on the following environment variable:

| Environment Variable | Description                                                                     | Example Value     |
| -------------------- | ------------------------------------------------------------------------------- | ----------------- |
| `CA_CERT_PATH`       | Path to the Certificate Authority (CA) file for validating client certificates. | `/path/to/ca.crt` |

::: warning
Ensure that the CA certificate file provided in `CA_CERT_PATH` is valid and accessible by Sushi Gateway. For more information, refer to the **[Environment Variables Configuration](../configuration/environment.md)**.
:::

## Example Configuration

Below is an example of enabling the mTLS plugin:

```json
{
  "name": "mtls",
  "enabled": true,
  "config": {}
}
```

### Explanation

- **`enabled`**: Activates the plugin to enforce mTLS.
- **`config`**: Currently, no additional configuration fields are required.

## Steps to Configure mTLS

1. **Prepare a CA Certificate**: Generate or obtain a valid CA certificate used to issue client certificates.

2. **Set the Environment Variable**:

   ```bash
   export CA_CERT_PATH=/path/to/ca.crt
   ```

3. **Apply the Plugin**: Enable the mTLS plugin in your configuration as shown above.

4. **Distribute Client Certificates**: Provide authorized clients with certificates signed by the trusted CA.

5. **Test the Handshake**: Verify that only clients with valid certificates can connect to Sushi Gateway.

## Important Note

The mTLS plugin requires requests to be sent over HTTPS to Sushi Gateway's **secure endpoint (port 8443)**. Any requests made to the HTTP endpoint will fail because mTLS requires a secure TLS handshake.

## Use Cases

1. **Secure Communication**: Ensure end-to-end encryption and mutual authentication between clients and the gateway.
2. **Restrict Unauthorized Access**: Allow only clients with trusted certificates to connect.
3. **Compliance Requirements**: Meet industry standards and regulations for secure communication.

## Tips for Using the mTLS Plugin

::: tip
Use strong and unique client certificates to prevent unauthorized access.
:::

::: tip
Regularly update and rotate CA certificates to maintain a secure environment.
:::

For more plugins, visit the **[Plugins Overview](../plugins/overview.md)**.
