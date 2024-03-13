# Sushi Proxy

## Setting up TLS server for local development

- Install OpenSSL
- Generate the CA private key and certificate
```bash
    # Generate ca private key 
    openssl genrsa -out ca.key 4096
    # Generate self signed CA cert
    openssl req -new -x509 -days 3650 -key ca.key -out ca.crt
```

- Generate a server private key and certificate signing request (CSR)
```bash
    # Generate server private key
    openssl genrsa -out server.key 2048
    # Generate server CSR
    openssl req -new -key server.key -out server.csr
```

- Create Server certificate signed by your CA
```bash
openssl x509 -req -days 365 -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt \
-extensions req_ext -extfile <(printf "[req_ext]\nsubjectAltName=DNS:localhost")
```

- Verify certs
```bash
openssl verify -CAfile ca.crt server.crt
```

- Set env variables in `.env` file`...

