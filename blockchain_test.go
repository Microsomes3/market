package main

import (
	"fmt"
	"math/big"
	"testing"
)

func TestBc(t *testing.T) {
	fmt.Println("hello")
	crp := &CryptoHelper{}

	ex := crp.SHA256([]byte("hello"))

	fmt.Println(crp.ByteToHex(ex[:]))

	var a big.Int
	a.SetBytes(ex[:])

	fmt.Println(a.Int64())
}
