package main

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

func signMessageSHA256(privateKey *ecdsa.PrivateKey, message []byte) ([]byte, error) {
	// Hash the message using SHA-256
	hash := sha256.Sum256(message)

	// Sign the hash
	signature, err := crypto.Sign(hash[:], privateKey)
	if err != nil {
		return nil, err
	}

	return signature, nil
}

func verifySignatureSHA256(publicKey *ecdsa.PublicKey, message []byte, signature []byte) bool {
	// Hash the original message
	hash := sha256.Sum256(message)

	// Verify the signature
	return crypto.VerifySignature(
		crypto.FromECDSAPub(publicKey),
		hash[:],
		signature[:len(signature)-1],
	)
}

func main() {
	// Generate a private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	// Get the public key
	publicKey := privateKey.PublicKey

	// Message to sign
	message := []byte("Hello, Ethereum Signing!")

	// Sign the message
	signature, err := signMessageSHA256(privateKey, message)
	if err != nil {
		log.Fatal(err)
	}

	// Verify the signature
	verified := verifySignatureSHA256(&publicKey, message, signature)

	fmt.Println("Original Message:", string(message))
	fmt.Println("Signature Verified:", verified)
}
