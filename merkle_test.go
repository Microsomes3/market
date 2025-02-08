package main

import (
	"fmt"
	"testing"
)

func TestCanGetMerkleRoot(t *testing.T) {
	crp := CryptoHelper{}
	m := NewMarkle()

	m.PushHash([32]byte{0})
	m.PushHash([32]byte{1})
	m.PushHash([32]byte{2})
	m.PushHash([32]byte{3})
	m.PushHash([32]byte{5})

	root := m.Root()

	fmt.Println(crp.ByteToHex(root[:]))

	if root == [32]byte{0} {
		t.Error("Root is not correct")
	}
}

func TestPairLoop(t *testing.T) {

	for i := 0; i < 10; i += 2 {

		fmt.Println(i)
		fmt.Println("[", i, i+1)
	}

}
