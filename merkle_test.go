package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestCanGetMerkleRoot(t *testing.T) {

	hash1 := sha256.Sum256([]byte("hash1"))
	hash2 := sha256.Sum256([]byte("hash2"))

	hash1and2 := sha256.Sum256(append(hash1[:], hash2[:]...))

	m := MerkleTree{}
	m.PushHash(hash1)
	m.PushHash(hash2)

	root := m.Root()

	if hex.EncodeToString(root[:]) != hex.EncodeToString(hash1and2[:]) {
		t.Fail()
	}

}

func TestCanGetMerkleRootOdd(t *testing.T) {

	hash1 := sha256.Sum256([]byte("hash1"))
	hash2 := sha256.Sum256([]byte("hash2"))
	hash3 := sha256.Sum256([]byte("hash3"))

	hash1and2 := sha256.Sum256(append(hash1[:], hash2[:]...))
	hash3and3 := sha256.Sum256(append(hash3[:], hash3[:]...))

	hash1and2and3 := sha256.Sum256(append(hash1and2[:], hash3and3[:]...))

	m := MerkleTree{}
	m.PushHash(hash1)
	m.PushHash(hash2)
	m.PushHash(hash3)

	root := m.Root()

	if hex.EncodeToString(root[:]) != hex.EncodeToString(hash1and2and3[:]) {
		t.Fail()
	}

}

func TestCanVerifyProof(t *testing.T) {

	hash1 := sha256.Sum256([]byte("hash1"))
	hash2 := sha256.Sum256([]byte("hash2"))

	m := MerkleTree{}
	m.PushHash(hash1)
	m.PushHash(hash2)

	root := m.Root()

	proof, isLeftNode, success := m.GetProof(hash1)

	fmt.Println(proof)
	fmt.Println(isLeftNode)
	fmt.Println(success)

	fmt.Println(hex.EncodeToString(root[:]))
	fmt.Println(hex.EncodeToString(proof[0][:]))

	//todo verify proof later, for now im happy it produces a root

}
