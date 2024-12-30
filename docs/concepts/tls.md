# TLS Certificates

Sushi Gateway requires TLS certificates for secure communication on port `8443`. You can either use auto-generated certificates for development or provide your own certificates for production use.

## Auto-generated Certificates

For development environments, Sushi Gateway can automatically generate self-signed certificates. This is the easiest way to get started:

```env
# Don't set SERVER_CERT_PATH and SERVER_KEY_PATH
# Certificates will be auto-generated in the current directory
```

The auto-generated certificates:

- Are self-signed
- Valid for 365 days
- Include "localhost" and "sushi.gateway.local" as DNS names
- Are stored as `server.crt` and `server.key` in the current directory

## Manual Certificates

For production environments, you should provide your own certificates:

```env
# Required for TLS
SERVER_CERT_PATH=/path/to/server.crt
SERVER_KEY_PATH=/path/to/server.key
```

Both `SERVER_CERT_PATH` and `SERVER_KEY_PATH` must be provided together. If either one is set, the other must also be set.

## mTLS Support (Optional)

For mutual TLS authentication, you can provide a CA certificate:

```env
# Optional - only needed for mTLS
CA_CERT_PATH=/path/to/ca.crt
```

The CA certificate is only required if you want to validate client certificates using mTLS. If you're not using mTLS, you can omit this setting.

::: info
For more information regarding environment configuration, check out the [Environment Variable Guide](./configuration/environment.md).
:::
