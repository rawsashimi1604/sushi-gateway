# JSON Web Token (JWT) Plugin

The JSON Web Token (JWT) (`jwt`) plugin provides a robust mechanism for securing APIs by validating signed tokens. This plugin ensures that only requests with valid JWTs can access your endpoints, enhancing authentication and authorization.

## How It Works

The JWT plugin inspects the `Authorization` header of incoming requests for a Bearer token. It validates the token’s signature and claims using the configured algorithm and secret or public key.

Currently, only **HS256** (HMAC-SHA256) and **RS256** (RSA-SHA256) is supported.

### Key Features

- Supports token validation for stateless authentication.
- Configurable issuer (`iss`) to ensure tokens come from trusted sources.
- Validates token signatures using a secret key.
- Validates token signatures against a public key.

::: tip
Learn how to integrate this plugin into your setup in the **[Plugins Overview](../plugins/index.md)**.
:::

## Configuration Fields

| Field       | Type   | Description                                               | Example Value              |
| ----------- | ------ | --------------------------------------------------------- | -------------------------- |
| `alg`       | String | Algorithm for token signature validation (e.g., `HS256`). | `HS256`                    |
| `iss`       | String | The expected issuer of the JWT.                           | `someIssuerKey`            |
| `secret`    | String | Secret key for validating the JWT signature.              | `123secret456`             |
| `publicKey` | String | Public key certificate.                                   | `some RSA public key here` |

::: warning
Currently, only **HS256** is supported. Future releases will include support for **RS256**.
:::

## Example Configuration

Below is an example of configuring the JWT plugin:

```json
{
  "name": "jwt",
  "enabled": true,
  "config": {
    "alg": "HS256",
    "iss": "someIssuer",
    "secret": "123secret456"
  }
}
```

As for and example for configuring the JWT plugin for `RS256` algorithms:

```json
{
  "name": "jwt",
  "enabled": true,
  "config": {
    "alg": "RS256",
    "iss": "someIssuer",
    "publicKey": "-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAnOAjYfgEIPkSuYR58sWr
K21TeuRU03O2SAM9MxY4ojuK61SiX5cAEycQB7JL8okgAu1Sja8PXebPxX3mo+Ea
1lCI+SZ27wBPpqeqpmDKxwbaacDFJYH7TaoctTX+pOuCtjS4h1BjrVaNn0SiLay2
KfSJt0AGTppidQC6Llg7aG1LLMkjISzic3IHUR1s9fAVj85qAEY889LYfwitKzln
xGO68PkV1dOCsJCquFOSRF92nA/DC0IMIXHGXEMP/GlXo2S4MB6e9z8Ti+R0HV96
99RxtH/GLZM+J3no+Jk5zbiZpSzTJCecmETYpn+jrozDnpwH/dc5GFQhMIROSXX9
RQIDAQAB
-----END PUBLIC KEY-----"
  }
}
```

### Explanation

- **`alg`**: Defines the algorithm for signature validation (only `HS256` and `RS256` is currently supported).
- **`iss`**: Specifies the trusted issuer of the token.
- **`secret`**: The secret key used for validating the token’s signature.
- **`publicKey`**: The public key used for validating the token’s signature.

## Example HTTP Header

To authenticate with a JWT, include it in the `Authorization` header as a Bearer token:

```http
Authorization: Bearer <your-jwt-token>
```

### Example:

If the JWT is `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c`, the header would look like:

```http
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
```

## Applying the Plugin

The JWT plugin can be applied at various levels:

1. **Global Level**: Secures all services and routes with JWT validation.
2. **Service Level**: Applies token validation to all routes under a specific service.
3. **Route Level**: Provides granular control for specific routes requiring JWT validation.

Example of applying the plugin globally:

```json
{
  "name": "jwt",
  "enabled": true,
  "config": {
    "alg": "HS256",
    "iss": "someIssuerKey",
    "secret": "123secret456"
  }
}
```

::: tip
Apply the plugin at the route level for APIs requiring fine-grained security.
:::

## Use Cases

1. **Stateless Authentication**: Secure APIs by validating JWTs without maintaining session state.
2. **Multi-Tenant APIs**: Use the `iss` claim to differentiate tenants or trusted sources.
3. **Enhance Security**: Combine with plugins like Rate Limiting for robust access control.

## Tips for Using the JWT Plugin

::: tip
Ensure the `secret` is securely stored and rotated periodically to maintain security.
:::

::: tip
Use HTTPS to protect the token in transit and prevent interception by attackers.
:::

## References

- **[RFC 7519: JSON Web Token (JWT)](https://datatracker.ietf.org/doc/html/rfc7519)**
- **[JWT.io](https://jwt.io/)**

For more plugins, visit the **[Plugins Overview](../plugins/index.md)**.
