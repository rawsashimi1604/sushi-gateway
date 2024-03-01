# Sushi Proxy

## Setting up TLS server

- Install OpenSSL
- Generate a private key and a self-signed certificate
- Example:
```bash
openssl genrsa -out key.pem 2048
openssl req -new -x509 -key key.pem -out cert.pem -days 365
```
- Set env variables to point to cert private key and cert path.



