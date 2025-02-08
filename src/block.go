package market

import (
	"bytes"
	"encoding/gob"
)

type Block struct {
	Hash [32]byte

	BlockSize uint64 //in bytes

	Tx []Transaction

	PrevHash [32]byte

	Nonce uint64

	MerkleRoot [32]byte

	Timestamp int64
}

func NewBlockTemplate() *Block {
	return &Block{
		Hash:       [32]byte{},
		BlockSize:  0,
		Tx:         []Transaction{},
		PrevHash:   [32]byte{},
		Nonce:      0,
		MerkleRoot: [32]byte{},
		Timestamp:  0,
	}
}

func (b *Block) Bytes() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(b); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (b *Block) HashIt() {
	b.Hash = [32]byte{} //reset to avoid hasing the hash

	bytes, err := b.Bytes()
	if err != nil {
		panic(err)
	}

	crp := &CryptoHelper{}
	b.Hash = crp.SHA256(bytes)
}

func (b *Block) CalculateMerkleRoot() {

	merc := MerkleTree{}

	for _, tx := range b.Tx {
		tx.UpdateHash()

		merc.PushHash(tx.Hash)
	}

	b.MerkleRoot = merc.Root()
}
