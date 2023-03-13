package main

import (
	"encoding/json"
	"fmt"
	"github.com/lestrrat-go/jwx/jwk"
	"io/ioutil"
	"log"
)

func main() {
	privateKeyData, err := ioutil.ReadFile("private.key")
	if err != nil {
		log.Fatal(err)
	}

	// getting private key from file
	jwk, err := jwk.ParseString(string(privateKeyData), jwk.WithPEM(true))
	if err != nil {
		panic(err)
	}

	// print jwk to verify
	buf, _ := json.MarshalIndent(jwk, "", "  ")
	fmt.Println(string(buf))
	
	// write this to jwks.json file
	file, _ := json.MarshalIndent(jwk, "", "  ")
	err = ioutil.WriteFile("jwks.json", file, 0644)
	if err == nil {
		fmt.Println("jwks.json file created")
	}
}
