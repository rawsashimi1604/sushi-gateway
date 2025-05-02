# Gateway Health Checks

## Overview

**Sushi Gateway** exposes a health check endpoint on its **Admin API** (`:8081`) to monitor the status of the gateway. This endpoint is used by:

- **Kubernetes liveness/readiness probes**
- **Load balancers**
- **Monitoring tools**

## **Health Check Endpoint**

### **Request**

- **Method:** `GET`
- **Path:** `/healthz`
- **Port:** `8081` (Admin API)

### **Response**

- **Status Code:** `200 OK` (Healthy)
- **Content-Type:** `application/json`
- **Body:**

```json
{
  "status": "healthy"
}
```

### **Example**

```sh
curl http://localhost:8081/healthz
```

**Output:**

```json
{ "status": "healthy" }
```

---
