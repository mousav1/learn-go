package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"runtime/secret"
)

// simulate sensitive key generation
func generateSecretKey() []byte {

	key := make([]byte, 32)

	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}

	return key
}

// unsafe version (memory may keep sensitive data)
func unsafeExample() {

	fmt.Println("=== Unsafe Example ===")

	key := generateSecretKey()

	fmt.Println("Secret key:", hex.EncodeToString(key))

	// key may remain in memory after this function
}

// safe version using secret.Do
func safeExample() {

	fmt.Println("\n=== Safe Example using secret.Do ===")

	var publicResult string

	secret.Do(func() {

		key := generateSecretKey()

		// use key
		publicResult = hex.EncodeToString(key[:8])

		fmt.Println("Using secret key internally")

		// key will be wiped after function ends
	})

	fmt.Println("Public result:", publicResult)
}

// production example
func deriveSessionToken() string {

	var token string

	secret.Do(func() {

		key := generateSecretKey()

		token = hex.EncodeToString(key[:16])

	})

	return token
}

func main() {

	unsafeExample()

	safeExample()

	fmt.Println("\n=== Production Example ===")

	token := deriveSessionToken()

	fmt.Println("Session token:", token)
}
