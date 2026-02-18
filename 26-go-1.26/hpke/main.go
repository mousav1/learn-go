package main

import (
	"crypto/hpke"
	"encoding/hex"
	"fmt"
)

func basicExample() {

	fmt.Println("=== HPKE Basic Example ===")

	kem := hpke.MLKEM768X25519()
	kdf := hpke.HKDFSHA256()
	aead := hpke.AES256GCM()

	// generate recipient key
	privateKey, err := kem.GenerateKey()
	if err != nil {
		panic(err)
	}

	publicKeyBytes := privateKey.PublicKey().Bytes()

	// sender reconstruct public key
	publicKey, err := kem.NewPublicKey(publicKeyBytes)
	if err != nil {
		panic(err)
	}

	message := []byte("hello secure world")

	ciphertext, err := hpke.Seal(
		publicKey,
		kdf,
		aead,
		nil,
		message,
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Ciphertext:", hex.EncodeToString(ciphertext))

	plaintext, err := hpke.Open(
		privateKey,
		kdf,
		aead,
		nil,
		ciphertext,
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Decrypted:", string(plaintext))
}

// production example
func encryptAPIKey(apiKey string) []byte {

	kem := hpke.MLKEM768X25519()
	kdf := hpke.HKDFSHA256()
	aead := hpke.AES256GCM()

	priv, _ := kem.GenerateKey()

	pub, _ := kem.NewPublicKey(priv.PublicKey().Bytes())

	ciphertext, _ := hpke.Seal(
		pub,
		kdf,
		aead,
		nil,
		[]byte(apiKey),
	)

	return ciphertext
}

func main() {

	basicExample()

	fmt.Println("\n=== Production Example ===")

	encrypted := encryptAPIKey("my-secret-api-key")

	fmt.Println("Encrypted API key:", hex.EncodeToString(encrypted))
}
