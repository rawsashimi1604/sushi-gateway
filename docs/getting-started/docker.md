# Quick Start with Docker

Welcome to the **Getting Started** guide for Sushi Gateway! This guide will help you set up Sushi Gateway quickly using Docker. In just a few steps, you'll have a working API Gateway to manage and secure your APIs.

## Step 0: Install Docker

::: tip
Ensure that Docker is installed and running on your machine. Docker is required to run Sushi Gateway components.

- **Install Docker Desktop** (for Windows/Mac): [Docker Desktop Installation Guide](https://docs.docker.com/desktop/)
- **Install Docker Engine** (for Linux): [Docker Engine Installation Guide](https://docs.docker.com/engine/install/)
  :::

Verify that Docker is installed correctly:

```bash
docker --version
```

You should see output similar to:

```bash
Docker version 24.0.5, build ced0996
```

Now you're ready to proceed!

## Step 1: Pull the Sushi Proxy and Manager Images

::: info
Pull the latest Docker images for Sushi Proxy and Sushi Manager from Docker Hub to get started.
:::

```bash
docker pull rawsashimi/sushi-proxy:0.4.1
docker pull rawsashimi/sushi-manager:0.4.1
```

## Step 2: Create a Configuration File

::: info Declarative Configuration
Sushi Gateway's configuration is managed using a declarative JSON file. For more information, refer to our [Declarative Configuration Concepts](../concepts/declarative-config.md).
:::

Create a `config.json` file with the following example configuration:

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
        { "id": "upstream_1", "host": "example-app", "port": 3000 }
      ],
      "routes": [
        {
          "name": "example-route",
          "path": "/v1/sushi",
          "methods": ["GET"],
          "plugins": [
            {
              "id": "example-plugin",
              "name": "rate_limit",
              "enabled": true,
              "config": {
                "limit_second": 10,
                "limit_min": 10,
                "limit_hour": 100
              }
            }
          ]
        }
      ]
    }
  ]
}
```

::: tip Configuration Management
Learn more about managing configurations in our [Configuration Management Guide](../concepts/configuration/index.md), including environment variables and JSON files.
:::

## Step 3: Create a Docker Network

Set up a Docker network for your services:

```bash
docker network create sushi-network
```

## Step 4: Pull the Example Upstream Service Image

Pull the example Node.js service image:

```bash
docker pull rawsashimi/express-sushi-app:latest
```

## Step 5: Run the Upstream Service

Run the upstream service container with the required environment variables:

::: tip
The upstream service is the backend that Sushi Gateway will route to. In this guide, we are using a simple example Node.js docker image - **Sushi App** that provides a mock API for sushi data.
:::

```bash
docker run -d \
--name example-app \
--network sushi-network \
-e APP_ID=3000 \
-e JWT_ISSUER=someIssuerKey \
-e JWT_SECRET=123secret456 \
-p 3000:3000 \
rawsashimi/express-sushi-app:latest
```

## Step 6: Test the Service

Ensure the upstream service is working:

```bash
curl http://localhost:3000/v1/sushi | jq
```

Expected response:

```json
{
  "app_id": "3000",
  "data": [
    {
      "id": 1,
      "name": "California Roll",
      "ingredients": ["Crab", "Avocado", "Cucumber"]
    },
    {
      "id": 2,
      "name": "Tuna Roll",
      "ingredients": ["Tuna", "Rice", "Nori"]
    }
  ]
}
```

## Step 7: Run the Sushi Proxy

Launch the Sushi Proxy container with the configuration and certificates:

```bash
docker run \
--rm \
--name example-proxy \
--network sushi-network \
-v $(pwd)/config.json:/app/config.json \
-e CONFIG_FILE_PATH="/app/config.json" \
-e ADMIN_USER=admin \
-e ADMIN_PASSWORD=changeme \
-p 8008:8008 \
-p 8081:8081 \
-p 8443:8443 \
rawsashimi/sushi-proxy:0.4.1
```

## Step 8: Test the Proxy

Verify that the proxy works:

```bash
curl http://localhost:8008/example/v1/sushi | jq
```

Expected response:

```json
{
  "app_id": "3000",
  "data": [
    {
      "id": 1,
      "name": "California Roll",
      "ingredients": ["Crab", "Avocado", "Cucumber"]
    },
    {
      "id": 2,
      "name": "Tuna Roll",
      "ingredients": ["Tuna", "Rice", "Nori"]
    }
  ]
}
```

## Step 9: Run Sushi Manager

Launch the interactive UI at `http://localhost:5173` via the pulled UI docker image for managing and monitoring your gateway:

::: info
Sushi Manager provides a user-friendly interface for configuring and monitoring your gateway.
:::

::: tip
Use the credentials specified above - `ADMIN_USER` and `ADMIN_PASSWORD` to login! If you have followed the guide, the credentials will be `admin` and `changeme` respectively.
:::

```bash
docker run --rm -p 5173:5173 \
-e SUSHI_MANAGER_BACKEND_API_URL=http://localhost:8081 \
rawsashimi/sushi-manager:0.4.1
```

## Next Steps

Congratulations on setting up Sushi Gateway! Now that you have your gateway running, explore these additional features:

- **[Explore More Plugins](../plugins/index.md)**: Enhance your API management with powerful plugins for rate limiting, authentication, and more.
- **[Admin REST API](../api/index.md)**: Learn how to query your gateway state using the Admin REST API.

Dive deeper into Sushi Gateway’s capabilities and take full control of your API ecosystem!
