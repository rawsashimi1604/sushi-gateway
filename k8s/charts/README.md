# Sushi Gateway Helm Chart

## Installing

<!-- TODO: enhance this document and also add in the /docs directory -->

- Ensure in chart relative directory
- Install minikube
- Start minikube cluster
- Run the command `helm upgrade --install --debug sushi-gateway .`
- `minikube tunnel`
- Manager should be available in `YOUR_EXPOSED_URL:5173`
- Proxy should be available in `YOUR_EXPOSED_URL:8008`, `YOUR_EXPOSED_URL:8443`, `YOUR_EXPOSED_URL:8081`
