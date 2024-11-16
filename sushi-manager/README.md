# Sushi Manager

## Development

Ensure that sushi gateway is running with the Admin API up.

### Docker build

```bash
docker build -t sushi-manager:latest .
```

### Docker Run

```bash
docker run --rm -p 5173:5173 \
-e SUSHI_MANAGER_BACKEND_API_URL=http://localhost:8081 \
sushi-manager:latest
```
