package market

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

	crp := &CryptoHelper{}

	_, pubk, _ := crp.GeneratePrivateKey()

	input1 := Vin{
		TXID:      [32]byte{},
		Vout:      0,
		Signature: [64]byte{},
	}

	tx.PushInput(input1)

	pubBytes := crp.GetPublicKeyBytes(pubk)

	output1 := Vout{
		Value: 50 * 1_000_000,
		N:     0,
		PK:    pubBytes,
	}

	tx.PushOutput(output1)

	tx.UpdateHash()

	fmt.Println(tx.Hash)
}
