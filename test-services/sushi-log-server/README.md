# Test Log Consumer for sushi-gateway

## Pushing this image to dockerhub

Added instructions on how to push image to docker hub, keep forgetting!

```
docker build -t rawsashimi/express-sushi-http-consumer:latest .
docker push rawsashimi/express-sushi-http-consumer:latest
```

## Public and private keys for RSA256 JWT signature

- Private key:

```text
-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAnOAjYfgEIPkSuYR58sWrK21TeuRU03O2SAM9MxY4ojuK61Si
X5cAEycQB7JL8okgAu1Sja8PXebPxX3mo+Ea1lCI+SZ27wBPpqeqpmDKxwbaacDF
JYH7TaoctTX+pOuCtjS4h1BjrVaNn0SiLay2KfSJt0AGTppidQC6Llg7aG1LLMkj
ISzic3IHUR1s9fAVj85qAEY889LYfwitKzlnxGO68PkV1dOCsJCquFOSRF92nA/D
C0IMIXHGXEMP/GlXo2S4MB6e9z8Ti+R0HV9699RxtH/GLZM+J3no+Jk5zbiZpSzT
JCecmETYpn+jrozDnpwH/dc5GFQhMIROSXX9RQIDAQABAoIBADG9Exrd0xlNP3WY
nj2uFK5pNF2zhX2ho3rDpCTNd9lgPZSNugnKy5hG+1slWdUlSwQCkPkhNyMTqm64
k2cEWUj4MeWlH3J5y8dQQ9gKumBOOPPszdUtmCsws3d1Di7mCQSSRKuKfoAYAEEu
Nql6qCs3QG7cmuNTKcJlH6LQEM3z0SBqtNEJTggMPNcDNU9XMr6o1iAr5CPiFDnK
N5mpTzxFpyCQsKOBtcNBP/DPjRWEQl6igc6EDbJglbihSvoPkROkZaxz4mdDvWmw
nGIfKK+9NJDNWAAmrUfLo/50+HOgSns9ZxpINeIAT8+rPRm5h/ixUGCG5KYR7eft
TEu+oWECgYEA4vgd2DxKyBrlCC2Czx6LIHoztSc2gcKixxUOUAO4QG/R+RkxyjQN
V2DgDxwysSFNsXVjJeIYgoUjtZDBFtffDA2tTH+1J673KrN1Dd/8363RBG7DlJcA
cw1i+h1fUdgC6P+sC09yCONbj2i/+WgOftDUodBjpNwMQZs5QQmcR00CgYEAsPDf
nP+kG3Aj+AojPis22riXracgtFt0l1wsifn4bRRreT//KisEHIo12QAWgdcGfHz9
mdvUnL0JO1iV2bWwdhlZ61b/dx7PQArDOYI5ME1nGvL+hxiaos8r/aCCRpIw3YjV
5WpYgjjIlz19RljkSH9CWmntfA0vo9USbPnvQdkCgYEAn+IUc/yU2T1I1Wfp26ky
bGBpCFVlKidHr2H/wRG9u3aJvSWoUz4zn7fYXgyJEQnaxwVgIJGSnm7XZtFfk43h
y4Xe7CKSJDA2YNglvu5oHdE9ihfUoll0sZdef74tJWQ7OJLSSO1f8S7nkrBBe5l3
jJHjF1HKv5la8OQ9grkYY4kCgYEAj2JikuNGpUV2oGX2sUZrqUq0/2/TvNPv40g6
f4Uln59QiG0n5Y/+QPJvOG4tzwHkq7TN/YR7Apjdhk3/APGPEeTxTRiu5GT/JbKT
CWNR3KacyuXnBKsXhJ/F0j3j1DRbjOp6CvLmzoAdbRHTFtKqC2W063ezjzdQR78+
szjGfeECgYBgdEGKbo+f+73Imjo0++sISW+PtFcdr2/9BZ8fvrgPT3drwOAAfuBN
JmxSKRzjDpo2Z0l69z7vFVB2k23Iqh01ABT7xmSe45Q3FX+Xo5WKYrLPelj7hco3
GUmXrDZWFryzc8mH3NmDTzxzrsfWtd0jFUQepCuC6zZIdgR+DNyhGw==
-----END RSA PRIVATE KEY-----
```

- Public key: 
```text
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAnOAjYfgEIPkSuYR58sWr
K21TeuRU03O2SAM9MxY4ojuK61SiX5cAEycQB7JL8okgAu1Sja8PXebPxX3mo+Ea
1lCI+SZ27wBPpqeqpmDKxwbaacDFJYH7TaoctTX+pOuCtjS4h1BjrVaNn0SiLay2
KfSJt0AGTppidQC6Llg7aG1LLMkjISzic3IHUR1s9fAVj85qAEY889LYfwitKzln
xGO68PkV1dOCsJCquFOSRF92nA/DC0IMIXHGXEMP/GlXo2S4MB6e9z8Ti+R0HV96
99RxtH/GLZM+J3no+Jk5zbiZpSzTJCecmETYpn+jrozDnpwH/dc5GFQhMIROSXX9
RQIDAQAB
-----END PUBLIC KEY-----
```
