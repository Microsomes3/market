package market

import (
	"fmt"
	"testing"
	"time"
)

func TestCanFindPOWHash(t *testing.T) {

	b1 := &Block{
		Hash:       [32]byte{},
		BlockSize:  0,
		Tx:         []Transaction{},
		PrevHash:   [32]byte{},
		Nonce:      0,
		MerkleRoot: [32]byte{},
		Timestamp:  time.Now().Unix(),
	}

	pow := NewPow(b1, 10)

	nonce := pow.FindNonce()

	fmt.Println(b1.Nonce)

	if nonce < 0 {
		t.Error("Nonce is negative")
	}

	if nonce != b1.Nonce {
		t.Error("Nonce does not match")
	}

	isValid := VerifyPow(b1, 10)

	if !isValid {
		t.Fail()
	}
}
