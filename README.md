# JWT Token Verifier Using Envoy

Implementing JWT token verifier using Envoy Proxy.

## Description

This project is a JWT token verifier based on [Envoy](https://www.envoyproxy.io/) proxy. We use a randomly-generated 4096-bit key as a secret key to generate JWT tokens and RSA256 (public key/private key) to signing and verifing the token ([JWKS](https://auth0.com/docs/secure/tokens/json-web-tokens/json-web-key-sets)).

## Getting Started

### Dependencies

* This project is containerized and you can simply download and run it as follow.

### Download

```
# Clone this repository
$ git clone https://github.com/MoziiiKa/jwt-envoy-verifier
```

### Run

```
# Go into the repository
$ cd jwt-envoy-verifier

# Run the following bash script
$ ./Setup.bash
```
### Use (APIs)

To register a new user
```
POST http://localhost:9099/api/v1/user-management/registration HTTP/1.1
content-type: application/json
    
{
    "name": "foo",
    "username": "bar",
    "email": "foo@example.com",
    "password": "complex_password"
}
```

To login (generating a new token)
```
POST http://localhost:9099/api/v1/token-management/token HTTP/1.1
content-type: application/json
    
{
    "email": "foo@example.com",
    "password": "complex_password"
}
```

To authorized access
```
GET http://localhost:9099/api/v1/access-management/time-ir HTTP/1.1
content-type: application/text
authorization: <copy from the response of the login API and paste here>
```

## Author

Mozaffar Kazemi  

## Version History

* 0.1
    * Initial Release

## License

This project is licensed under the MIT License.