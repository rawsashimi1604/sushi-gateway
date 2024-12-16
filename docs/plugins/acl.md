# Access Control List (ACL) Plugin

The Access Control List (ACL) (`acl`) plugin provides a mechanism to control access to APIs based on client IP addresses. By defining whitelists and blacklists, this plugin allows you to restrict or permit traffic from specific IP ranges or addresses.

## How It Works

The ACL plugin inspects the client’s IP address from incoming requests and checks it against configured whitelists or blacklists:

1. **Whitelist**: Allows traffic only from IP addresses in the whitelist.
2. **Blacklist**: Blocks traffic from IP addresses in the blacklist.

::: warning
You cannot configure both `whitelist` and `blacklist` at the same time. Choose one based on your requirements.
:::

Requests that do not meet the criteria are rejected with a **403 Forbidden** response.

### Key Features

- IP-based access control for services and routes.
- Supports both whitelists and blacklists.
- Configurable at global, service, or route levels.

::: tip
Learn how to integrate this plugin into your setup in the **[Plugins Overview](../plugins/overview.md)**.
:::

## Configuration Fields

| Field       | Type  | Description                                     | Example Value                |
| ----------- | ----- | ----------------------------------------------- | ---------------------------- |
| `whitelist` | Array | List of IP addresses allowed access to the API. | `["127.0.0.1", "127.0.0.2"]` |
| `blacklist` | Array | List of IP addresses denied access to the API.  | `["192.168.0.1"]`            |

::: tip
Use `whitelist` to restrict access to trusted clients and `blacklist` to block known malicious IPs.
:::

## Example Configuration

### Whitelist Example

```json
{
  "name": "acl",
  "enabled": true,
  "config": {
    "whitelist": ["127.0.0.1", "127.0.0.2"]
  }
}
```

### Blacklist Example

```json
{
  "name": "acl",
  "enabled": true,
  "config": {
    "blacklist": ["192.168.0.1"]
  }
}
```

### Explanation

- **`whitelist`**: Only requests from `127.0.0.1` and `127.0.0.2` are allowed.
- **`blacklist`**: Requests from `192.168.0.1` are blocked.

## Applying the Plugin

The ACL plugin can be applied at various levels:

1. **Global Level**: Applies access control to all services and routes.
2. **Service Level**: Applies access control to all routes within a service.
3. **Route Level**: Applies access control to specific routes.

Example of applying the plugin globally:

```json
{
  "name": "acl",
  "enabled": true,
  "config": {
    "whitelist": ["127.0.0.1", "127.0.0.2"]
  }
}
```

::: tip
Apply the plugin at the route level for granular control of access restrictions.
:::

## Use Cases

1. **Restrict Internal APIs**: Allow access only to trusted IPs for sensitive endpoints.
2. **Block Malicious Traffic**: Deny access to known bad actors using a blacklist.
3. **Enhance API Security**: Layer access control with other security plugins like JWT or Basic Auth.

## Tips for Using the ACL Plugin

::: tip
Regularly review and update your IP lists to maintain effective access control.
:::

For more plugins, visit the **[Plugins Overview](../plugins/overview.md)**.
