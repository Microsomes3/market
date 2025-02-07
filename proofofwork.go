package main

import (
	"fmt"
	"math/big"
)

type POW struct {
	Subject    *Block
	TargetDiff *big.Int
}

func VerifyPow(subject *Block, diff uint64, nonce uint64) bool {
	td := big.NewInt(1)
	td.Lsh(td, uint(256-diff))

	subject.Nonce = nonce
	subject.HashIt()

	var intHash big.Int
	intHash.SetBytes(subject.Hash[:])

	return intHash.Cmp(td) < 0
}

func NewPow(subject *Block, diff uint64) *POW {
	td := big.NewInt(1)
	td.Lsh(td, uint(256-diff))

	return &POW{
		Subject:    subject,
		TargetDiff: td,
	}
}

func (pow *POW) FindNonce() uint64 {
	nonce := uint64(0)
	for {

		pow.Subject.Nonce = nonce

		fmt.Println("Nonce: ", nonce)
		fmt.Println("Hash: ", pow.Subject.Hash)
		pow.Subject.HashIt()

		var intHash big.Int
		intHash.SetBytes(pow.Subject.Hash[:])

		if intHash.Cmp(pow.TargetDiff) < 0 {
			fmt.Println("Found hash: ", intHash)
			fmt.Println("Hash String", pow.Subject.Hash)
			break
		}

		nonce++

	}

	return nonce
	// return nonce
}
