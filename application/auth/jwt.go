package auth

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
)

// generate a jwt token when user login
func GenerateJWT(username string, password string, accessTokenMaxAge int) (string, error) {
	privateKeyData, err := ioutil.ReadFile("private.key")
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBlock, _ := pem.Decode(privateKeyData)
	if privateKeyBlock == nil || privateKeyBlock.Type != "RSA PRIVATE KEY" {
		log.Fatal("failed to decode PEM block containing RSA private key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	// getting private key from file
	jwk, err := jwk.ParseString(string(privateKeyData), jwk.WithPEM(true))
	if err != nil {
		panic(err)
	}

	// print jwk to verify (only for debuging)
	buf, _ := json.MarshalIndent(jwk, "", "  ")
	fmt.Println(string(buf))

	// generating jwt token
	token := jwt.New()
	token.Set(jwt.IssuerKey, "https://example.com")
	token.Set(jwt.ExpirationKey, time.Now().Add(time.Minute*time.Duration(accessTokenMaxAge)).Unix())
	token.Set(jwt.IssuedAtKey, time.Now().Unix())
	ciaims := map[string]interface{}{
		"username": username,
		"password": password,
	}
	token.Set("ciaims", ciaims)

	// sign token with private key
	signedToken, err := jwt.Sign(token, jwa.RS256, privateKey)
	if err != nil {
		return "", err
	}
	return string(signedToken), nil
}
