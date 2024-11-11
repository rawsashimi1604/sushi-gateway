# Sushi Proxy

## Build docker image
```bash
docker build -t rawsashimi/sushi-proxy:latest . && \ 
docker push rawsashimi/sushi-proxy:latest 
```

## Setting up TLS server for local development

- Please place these files config folder. It should look like this
```bash
cd ./sushi-proxy/gateway 
gateway
├── ca.crt
├── ca.key
├── extfile.cnf
├── server.crt
├── server.csr
└── server.key
└── gateway.json    
```

- Set env variables as well in `.env` file

- Install OpenSSL
- Generate the CA private key and certificate
```bash
    # Generate ca private key 
    openssl genrsa -out ca.key 4096
    # Generate self signed CA cert
    openssl req -new -x509 -days 3650 -key ca.key -out ca.crt
```

- Generate a server private key and certificate signing request (CSR)
- create challenge password: `123456`
```bash
    # Generate server private key
    openssl genrsa -out server.key 2048
    # Generate server CSR
    openssl req -new -key server.key -out server.csr
```

- Create Server certificate signed by your CA
```bash
printf "[req_ext]\nsubjectAltName=DNS:localhost" > extfile.cnf
openssl x509 -req -days 365 -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt \
-extensions req_ext -extfile extfile.cnf
```

- Verify certs
```bash
openssl verify -CAfile ca.crt server.crt
```

- Set env variables in `.env` file`...

## Creating client certificate for MTLS
- Generate the Client Key and CSR 
```
openssl genrsa -out client.key 2048
openssl req -new -key client.key -out client.csr
```

- Sign the client CSR with the CA
```
openssl x509 -req -days 365 -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt
```

- Verify the client cert
```
openssl verify -CAfile ca.crt client.crt
```

