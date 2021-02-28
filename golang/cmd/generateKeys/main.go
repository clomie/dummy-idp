package main

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwk"
)

func main() {
	raw, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Printf("failed to generate new RSA privatre key: %s\n", err)
		return
	}

	privateKey, err := jwk.New(raw)
	if err != nil {
		fmt.Printf("failed to create symmetric key: %s\n", err)
		return
	}
	if _, ok := privateKey.(jwk.RSAPrivateKey); !ok {
		fmt.Printf("expected jwk.SymmetricKey, got %T\n", privateKey)
		return
	}

	privateKey.Set(jwk.KeyUsageKey, jwk.ForSignature.String())
	privateKey.Set(jwk.AlgorithmKey, jwa.RS256.String())
	jwk.AssignKeyID(privateKey)

	pkJSON, err := json.MarshalIndent(privateKey, "", "  ")
	if err != nil {
		fmt.Printf("failed to marshal private key: %s\n", err)
		return
	}

	err = ioutil.WriteFile("../privateKey.json", pkJSON, 0666)
	if err != nil {
		fmt.Printf("failed to read privateKey.json: %s\n", err)
		return
	}
	fmt.Println("Output: privateKey.json")

	publicKey, err := jwk.PublicKeyOf(privateKey)

	jwks := jwk.NewSet()
	jwks.Add(publicKey)

	jwksJSON, err := json.MarshalIndent(jwks, "", "  ")
	if err != nil {
		fmt.Printf("failed to marshal private key: %s\n", err)
		return
	}

	err = ioutil.WriteFile("../dummy-idp/.well-known/jwks.json", jwksJSON, 0666)
	if err != nil {
		fmt.Printf("failed to open jwks.json: %s\n", err)
		return
	}
	fmt.Println("Output: jwks.json")
}
