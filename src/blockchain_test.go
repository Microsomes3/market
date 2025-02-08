package market

import (
	"encoding/hex"
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

func TestGenerateGenesisBlock(t *testing.T) {

	crp := &CryptoHelper{}

	_, pubk, _ := crp.GenerateDeterministicKey([]byte("burn"))

	var msg [64]byte

	copy(msg[:], "UK demands access to Apple users' encrypted data/7 February 2025")

	fmt.Println(string(msg[:]))

	vin := Vin{
		TXID:      [32]byte{},
		Vout:      0,
		Signature: msg,
	}

	vout := Vout{
		Value: 50 * 1_000_000,
		N:     0,
		PK:    crp.GetPublicKeyBytes(pubk),
	}

	tx1 := Transaction{
		Hash:     [32]byte{},
		Fee:      0,
		Locktime: 0,
		Vin:      []Vin{vin},
		Vout:     []Vout{vout},
	}

	genesisBlock := &Block{
		Hash:       [32]byte{},
		PrevHash:   [32]byte{},
		BlockSize:  0,
		Tx:         []Transaction{tx1},
		Nonce:      0,
		MerkleRoot: [32]byte{},
		Timestamp:  0,
	}

	genesisBlock.CalculateMerkleRoot()

	fmt.Println("blockhash: ", genesisBlock.Hash)

	POW := NewPow(genesisBlock, 10)

	POW.FindNonce()

	fmt.Println(hex.EncodeToString(genesisBlock.Hash[:]))

	fmt.Println(genesisBlock.Nonce)

}

func TestNewBlockchai(t *testing.T) {

	blockchain := NewBlockchain("testnet")

	fmt.Println(blockchain.TailBlock)

	block1 := NewBlockTemplate()

	err := blockchain.AddBlock(block1)

	if err != nil {
		t.Fail()
	}

}

func TestGenesisBlockMatchesHash(t *testing.T) {

	bc := Block{}
	gensis := bc.GenesisBlock()

	isValid := VerifyPow(gensis, 21)

	if !isValid {
		t.Fail()
	}

}
