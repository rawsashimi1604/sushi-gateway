# Sushi Gateway

![High Level Design](./docs/images/design.png)

## TLS
https/tls support has been added to the proxy, add cert and key into environment variables to use.

## Plugins
- TODO: need to externalize plugin configurations ASAP, move credentials to plugin array. Remove credentials array 
```json
{
  "plugins": [
    {
      "name": "key-auth",
      "config": {
        "key": "someConfigurationHere"
      }
    }
  ] 
}
```


- Auth
  - **(DONE)** basic auth
  - **(DONE)** jwt (only Hs256, Rs256 tbd)
  - **(DONE)** key auth
  - oauth2
- Security
  - **(DONE)** bot_protection
  - cors
- Traffic Control
  - **(DONE)** acl
  - rate limit
  - **(DONE)** request size limit

    
