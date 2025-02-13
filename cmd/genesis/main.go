package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"time"

	market "microsomes.com/silky/src"
)

func main() {

	//this will eventually be a cli tool using it as a quick placeholder and to generate the main genesis block
	crp := market.CryptoHelper{}

	// _, pubk, _ := crp.GenerateDeterministicKey([]byte("burn"))
	priv, pubk, _ := crp.GeneratePrivateKey()

	pr, _ := os.Create("genesis_private")

	hexEncodedPrivateKey := hex.EncodeToString(crp.GetPrivateKeyBytes(priv))

	pr.WriteString(hexEncodedPrivateKey)
	pr.Close()

	var msg [64]byte

	copy(msg[:], "UK demands access to Apple users' encrypted data/7 February 2025")

	fmt.Println(string(msg[:]))

	vin := market.Vin{
		TXID:      [32]byte{},
		Vout:      0,
		Signature: msg,
	}

	vout := market.Vout{
		Value: 50 * 1_000_000,
		N:     0,
		PK:    crp.GetPublicKeyBytes(pubk),
	}

	fmt.Println(crp.GetPublicKeyBytes(pubk))

	tx1 := market.Transaction{
		Hash:     [32]byte{},
		Fee:      0,
		Locktime: 0,
		Vin:      []market.Vin{vin},
		Vout:     []market.Vout{vout},
	}
	tx1.UpdateHash()

	genesisBlock := &market.Block{
		Hash:       [32]byte{},
		PrevHash:   [32]byte{},
		BlockSize:  0,
		Tx:         []market.Transaction{tx1},
		Nonce:      0,
		MerkleRoot: [32]byte{},
		Timestamp:  0,
	}

	genesisBlock.CalculateMerkleRoot()

	fmt.Println("blockhash: ", genesisBlock.Hash)

	POW := market.NewPow(genesisBlock, 25)

	startTime := time.Now().Unix()

	POW.FindNonce()

	fmt.Println("hash", hex.EncodeToString(genesisBlock.Hash[:]))

	fmt.Println("nonce", genesisBlock.Nonce)
	fmt.Println("timestamp", genesisBlock.Timestamp)
	fmt.Println("merkleroot", genesisBlock.MerkleRoot)

	taken := time.Now().Unix() - startTime

	fmt.Println("seconds taken:", taken)

	fmt.Println("orgnonce", genesisBlock.Nonce)
	fmt.Println("orgnonce", genesisBlock.Hash)
	fmt.Println(hex.EncodeToString(genesisBlock.Hash[:]))

	genesisBlock.HashIt()

	fmt.Println("orgnonce", hex.EncodeToString(genesisBlock.Hash[:]))

	by, _ := genesisBlock.BytesWithoutHeader()

	sha2 := sha256.Sum256(by)

	fmt.Println(hex.EncodeToString(sha2[:]))

	genesisBlockF, err := os.Create("genesis.json")

	if err != nil {
		panic(err.Error())
	}
	json.NewEncoder(genesisBlockF).Encode(genesisBlock)

	bytesBlock, err := genesisBlock.Bytes()

	if err != nil {
		panic(err.Error())
	}

	fblock, err := os.Create("genesis_block2")

	if err != nil {
		panic(err.Error())
	}

	fblock.Write(bytesBlock)

	fblock.Close()

	isValid := genesisBlock.VerifyBlock()

	fmt.Println(isValid)

	//000000b5f68cf4df0587440fb66843afd0ebf2e88545c97a1bd6d0cfc01dea3c
	// 3239645

}
