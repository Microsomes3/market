package market

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
)

type Transaction struct {
	Hash [32]byte `json:"hash"`

	Vin      []Vin
	Vout     []Vout
	Fee      uint64
	Locktime uint64 //when this tx can be spent
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Hash     string `json:"hash"`
		Fee      uint64 `json:"fee"`
		Locktime uint64 `json:"locktime"`
		Vin      []Vin  `json:"vin"`
		Vout     []Vout `json:"vout"`
	}{
		Hash:     hex.EncodeToString(t.Hash[:]),
		Fee:      t.Fee,
		Locktime: t.Locktime,
		Vin:      t.Vin,
		Vout:     t.Vout,
	})
}

func (tx *Transaction) UnmarshalJSON(bdata []byte) error {

	var tempTX struct {
		Hash     string `json:"hash"`
		Fee      uint64 `json:"fee"`
		Locktime uint64 `json:"locktime"`
		Vin      []Vin  `json:"vin"`
		Vout     []Vout `json:"vout"`
	}

	if err := json.Unmarshal(bdata, &tempTX); err != nil {
		return err
	}

	hashE, _ := hex.DecodeString(tempTX.Hash)

	copy(tx.Hash[:], hashE)

	tx.Fee = tempTX.Fee
	tx.Locktime = tempTX.Locktime

	tx.Vin = tempTX.Vin

	tx.Vout = tempTX.Vout

	return nil
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
