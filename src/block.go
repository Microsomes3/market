package market

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type Block struct {
	Hash [32]byte

	BlockSize uint64

	Tx []Transaction

	PrevHash [32]byte

	Nonce uint64

	MerkleRoot [32]byte

	Timestamp int64
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Hash       string        `json:"hash"`
		PrevHash   string        `json:"prev_hash"`
		MerkleRoot string        `json:"merkle_root"`
		Nonce      uint64        `json:"nonce"`
		BlockSize  uint64        `json:"size"`
		Timestamp  int64         `json:"timestamp"`
		Tx         []Transaction `json:"tx"`
	}{
		Hash:       hex.EncodeToString(b.Hash[:]),
		PrevHash:   hex.EncodeToString(b.PrevHash[:]),
		MerkleRoot: hex.EncodeToString(b.MerkleRoot[:]),
		Nonce:      b.Nonce,
		BlockSize:  b.BlockSize,
		Timestamp:  b.Timestamp,
		Tx:         b.Tx,
	})
}

func (b *Block) UnmarshalJSON(bdata []byte) error {

	var tempBlock struct {
		Hash       string        `json:"hash"`
		PrevHash   string        `json:"prev_hash"`
		MerkleRoot string        `json:"merkle_root"`
		Nonce      uint64        `json:"nonce"`
		BlockSize  uint64        `json:"size"`
		Timestamp  int64         `json:"timestamp"`
		Tx         []Transaction `json:"tx"`
	}

	if err := json.Unmarshal(bdata, &tempBlock); err != nil {
		return err
	}

	hash, err := hex.DecodeString(tempBlock.Hash)

	if err != nil {
		return err
	}

	copy(b.Hash[:], hash)

	prevHash, err := hex.DecodeString(tempBlock.PrevHash)

	if err != nil {
		return nil
	}

	copy(b.PrevHash[:], prevHash)

	merkleRoot, err := hex.DecodeString(tempBlock.MerkleRoot)

	if err != nil {
		return nil
	}

	copy(b.MerkleRoot[:], merkleRoot)

	b.Nonce = tempBlock.Nonce
	b.BlockSize = tempBlock.BlockSize
	b.Timestamp = tempBlock.Timestamp

	b.Tx = tempBlock.Tx

	return nil
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

func (b Block) BytesWithoutHeader() ([]byte, error) {
	b.Hash = [32]byte{}
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(b); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (b *Block) Bytes() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(b); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func BlockFromBytes(b []byte) (*Block, error) {
	var block Block

	dec := gob.NewDecoder(bytes.NewReader(b))
	if err := dec.Decode(&block); err != nil {
		return nil, err
	}

	return &block, nil
}

func (b *Block) HashIt() {
	b.Hash = [32]byte{} //reset to avoid hasing the hash

	bytes, err := b.Bytes()
	if err != nil {
		panic(err)
	}

	bx := len(bytes)

	b.BlockSize = uint64(bx)

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

func (b *Block) GenesisBlock() *Block {
	crp := &CryptoHelper{}

	_, pubk, _ := crp.GenerateDeterministicKey([]byte("burn"))

	var msg [64]byte

	copy(msg[:], "UK demands access to Apple users' encrypted data/7 February 2025")

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

	genHash, _ := hex.DecodeString("000005997cbc47eb78418a96c0213b4002fdbbcaa01d958a8de975f94eddab4d")

	var hashArray [32]byte

	copy(hashArray[:], genHash)

	genesisBlock := &Block{
		Hash:      hashArray,
		PrevHash:  [32]byte{},
		BlockSize: 0,
		Tx:        []Transaction{tx1},
		Nonce:     1177637,
		MerkleRoot: [32]byte{
			69, 199, 36, 161, 120, 154, 47, 207, 47, 48, 107, 160, 71, 228, 65, 63, 228, 186, 77, 158, 140, 117, 136, 16, 254, 117, 215, 229, 182, 198, 65, 136,
		},
		Timestamp: 1739047442,
	}

	return genesisBlock
}

func CalculateAllowedCoinbaseReward(blockHash []byte) uint64 {
	return 50 * 1_000_000
}

func (b Block) VerifyBlock() bool {

	fmt.Println("verifying hash:", b.Hash)

	bytesWithoutHeader, _ := b.BytesWithoutHeader()

	hashCal := sha256.Sum256(bytesWithoutHeader)

	if hex.EncodeToString(hashCal[:]) != hex.EncodeToString(b.Hash[:]) {
		return false
	}

	//if no transactions its not valid block

	if len(b.Tx) == 0 {
		return false
	}

	//first tx must be a coinbase, aka have no references to previous

	if b.Tx[0].Vin[0].TXID != [32]byte{} {

		//is not coinbase
		return false

	}

	if b.Tx[0].Vin[0].Vout != 0 {

		//should only be 0 as theirs technically no vout
		return false

	}

	//check if coinbase value is too big

	//check an output exists for coinbase

	if len(b.Tx[0].Vout) == 0 {
		return false
	}

	//if grather then coinbase reward
	if b.Tx[0].Vout[0].Value > CalculateAllowedCoinbaseReward(b.Hash[:]) {
		return false
	}

	//check if less than
	if b.Tx[0].Vout[0].Value < CalculateAllowedCoinbaseReward(b.Hash[:]) {
		return false
	}

	//todo will need to at somepoint verify the other transactions using the database, but for now ignore the check

	return true
}

func VerifyBlockNoContext(b *Block) bool {
	//this will perform some checks without using the database
	//so it cant check if the transactions are valid, if the input and outputs are valid
	//its used to perform basic checks
	//such as coinbase tx is present

	return true
}
