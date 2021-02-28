package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"time"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
)

func main() {
	keyBytes, err := ioutil.ReadFile("../privateKey.json")
	if err != nil {
		fmt.Printf("failed to read privateKeys.json: %s\n", err)
		return
	}

	key, err := jwk.ParseKey(keyBytes)
	if err != nil {
		fmt.Printf("failed to parse private key: %s\n", err)
		return
	}

	token := jwt.New()
	token.Set(jwt.SubjectKey, os.Args[1])
	token.Set(jwt.IssuerKey, "http://localhost")
	token.Set(jwt.AudienceKey, "dummy-client-id")
	token.Set(jwt.IssuedAtKey, time.Now().Unix())
	token.Set(jwt.ExpirationKey, math.MaxInt32)

	buf, err := json.MarshalIndent(token, "", "  ")
	if err != nil {
		fmt.Printf("failed to generate JSON: %s\n", err)
		return
	}
	fmt.Printf("%s\n", buf)

	// Signing a token (using raw rsa.PrivateKey)
	signed, err := jwt.Sign(token, jwa.RS256, key)
	if err != nil {
		log.Printf("failed to sign token: %s", err)
		return
	}
	fmt.Printf("%s\n", signed)
}
