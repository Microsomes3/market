package main

import (
	"fmt"
	"testing"
)

func TestTX(t *testing.T) {
	tx := &Transaction{}
	fmt.Println(tx.HashIt())
	tx.UpdateHash()

	if tx.Hash == [32]byte{} {
		t.Error("Hash is empty")
	}
}

func TestCreateTr(t *testing.T) {
	tx := &Transaction{}

	input1 := Vin{
		TXID:      [32]byte{},
		Vout:      0,
		Signature: [32]byte{},
	}

	tx.PushInput(input1)
}
