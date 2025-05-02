# Sushi Gateway Helm Chart - Minikube Installation Guide

Welcome to the Getting Started guide for Sushi Gateway! This guide will help you set up Sushi Gateway quickly when deploying onto a Minikube cluster. In just a few steps, you'll have a working API Gateway to manage and secure your APIs.

## Prerequisites

Before installing Sushi Gateway on Minikube, ensure you have the following installed:

- [Minikube](https://minikube.sigs.k8s.io/docs/start/)
- [Helm](https://helm.sh/docs/intro/install/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/)

## Installation Steps

### 1. Start Minikube Cluster

First, start your Minikube cluster with adequate resources:

```bash
minikube start --cpus=4 --memory=8192 --driver=docker
```

Verify your cluster is running:

```bash
minikube status
kubectl get nodes
```

### 2. Install Sushi Gateway Helm Chart

Navigate to the chart directory and install the Sushi Gateway:

```bash
cd k8s/charts/minikube
helm upgrade --install --debug sushi-gateway .
```

This will:

- Deploy all Sushi Gateway components
- Create necessary Kubernetes resources
- Create configuration JSON file based on the content in `files/config.json`
- Start your cluster :)

### 3. Expose Services with Minikube Tunnel

In a new terminal, run the Minikube tunnel to expose services:

```bash
minikube tunnel
```

Keep this terminal open while using the gateway.

### 4. Verify Installation

Check deployed services:

```bash
kubectl get svc
```

Look for the `EXTERNAL-IP` column in the output. You should see services like:

- `sushi-gateway-proxy` (ports 8008, 8443, 8081)
- `sushi-gateway-manager` (port 5173)

### 5. Access Sushi Gateway Components

Use the following endpoints to access different components:

| Component   | URL Format                   | Default Ports |
| ----------- | ---------------------------- | ------------- |
| Manager UI  | `http://<EXTERNAL-IP>:5173`  | 5173          |
| Proxy HTTP  | `http://<EXTERNAL-IP>:8008`  | 8008          |
| Proxy HTTPS | `https://<EXTERNAL-IP>:8443` | 8443          |
| Admin API   | `http://<EXTERNAL-IP>:8081`  | 8081          |

### 6. Verify Connectivity

Test the manager endpoint:

```bash
curl http://<EXTERNAL-IP>:5173/login
```

And login using your credentials in `values.yaml`!. The default credentials should be `admin` and `changeme`.

## Cleanup

To remove the installation:

```bash
helm uninstall sushi-gateway
minikube stop
minikube delete
```
