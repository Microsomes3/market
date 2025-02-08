package main

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
)

type Transaction struct {
	Hash [32]byte

	Vin      []Vin
	Vout     []Vout
	Fee      uint64
	Locktime uint64 //when this tx can be spent
}

// Bytes serializes the transaction using gob encoding
func (tx *Transaction) Bytes() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(tx); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (tx *Transaction) TxHex() string {
	txbytes, _ := tx.Bytes()
	return hex.EncodeToString(txbytes)
}

func (tx *Transaction) HashIt() [32]byte {
	tx.Hash = [32]byte{}

	bytes, err := tx.Bytes()
	if err != nil {
		panic(err)
	}

	crp := &CryptoHelper{}
	return crp.SHA256(bytes)

}

func (tx *Transaction) UpdateHash() {
	tx.Hash = tx.HashIt()
}

func (tx *Transaction) PushInput(vin Vin) {
	tx.Vin = append(tx.Vin, vin)
}

func (tx *Transaction) PushOutput(vout Vout) {
	tx.Vout = append(tx.Vout, vout)
}
