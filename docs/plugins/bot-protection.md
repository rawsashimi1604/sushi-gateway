# Bot Protection Plugin

The Bot Protection (`bot_protection`) plugin helps safeguard your APIs by blocking requests based on User Agent strings in the HTTP headers. This plugin is ideal for filtering out traffic from known bots or restricting access to specific clients.

## How It Works

The Bot Protection plugin inspects the `User-Agent` header in incoming requests and applies one of two rules:

1. **Blacklist**: Blocks requests from User Agents listed in the `blacklist`.
2. **Whitelist**: Allows only requests from User Agents listed in the `whitelist`.

::: warning
You cannot configure both `blacklist` and `whitelist` at the same time. Choose one based on your requirements.
:::

Requests that do not meet the configured criteria are rejected with a **403 Forbidden** response.

### Key Features

- Simple filtering based on `User-Agent` strings.
- Configurable as a blacklist or whitelist.

::: tip
Learn how to integrate this plugin into your setup in the **[Plugins Overview](../plugins/overview.md)**.
:::

## Configuration Fields

| Field       | Type  | Description                                                    | Example Value              |
| ----------- | ----- | -------------------------------------------------------------- | -------------------------- |
| `blacklist` | Array | List of User Agents to block. Cannot be used with `whitelist`. | `["googlebot", "bingbot"]` |
| `whitelist` | Array | List of User Agents to allow. Cannot be used with `blacklist`. | `["custom-client"]`        |

::: tip
Use `blacklist` for blocking unwanted bots, and `whitelist` for restricting access to specific clients.
:::

## Example Configuration

### Blacklist Example

```json
{
  "name": "bot_protection",
  "enabled": true,
  "config": {
    "blacklist": ["googlebot", "bingbot", "yahoobot"]
  }
}
```

### Whitelist Example

```json
{
  "name": "bot_protection",
  "enabled": true,
  "config": {
    "whitelist": ["custom-client", "trusted-agent"]
  }
}
```

### Explanation

- **`blacklist`**: Blocks requests from the specified User Agents.
- **`whitelist`**: Allows requests only from the specified User Agents.

## Applying the Plugin

The Bot Protection plugin can be applied at various levels:

1. **Global Level**: Filters all traffic at the gateway.
2. **Service Level**: Applies filtering to specific services.
3. **Route Level**: Restricts access to individual routes.

Example of applying the plugin globally:

```json
{
  "name": "bot_protection",
  "enabled": true,
  "config": {
    "blacklist": ["googlebot", "bingbot"]
  }
}
```

::: tip
Apply the plugin at the global level to protect all APIs from unwanted User Agents.
:::

## Use Cases

1. **Block Search Engine Crawlers**: Prevent bots like `googlebot` and `bingbot` from indexing sensitive APIs.
2. **Restrict API Access**: Allow only trusted clients using the `whitelist` feature.
3. **Enhance Security**: Reduce unwanted traffic and potential misuse by filtering suspicious User Agents.

## Tips for Using the Bot Protection Plugin

::: tip
Regularly update the `blacklist` or `whitelist` based on observed traffic patterns and new threats.
:::

::: tip
Combine with other plugins like Rate Limiting for layered protection.
:::

::: tip
Monitor blocked requests using logs or analytics to fine-tune your configuration.
:::

::: warning
This plugin provides basic filtering and may not prevent sophisticated attacks. Combine it with authentication and IP-based restrictions for robust security.
:::

For more plugins, visit the **[Plugins Overview](../plugins/overview.md)**.
