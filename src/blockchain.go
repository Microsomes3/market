package market

import (
	"errors"

	"github.com/dgraph-io/badger/v3"
)

type Blockchain struct {
	DB               *badger.DB
	TailBlock        [32]byte //hash of the latest block
	Height           int
	TargetDifficulty int64
	OrphanBlocks     []Block //keep these around until they can be added to the chain, or discarded
}

func NewBlockchain(dbName string) *Blockchain {
	badgerDB, err := badger.Open(badger.DefaultOptions(dbName))
	if err != nil {
		panic(err)
	}
	return &Blockchain{
		DB:               badgerDB,
		TailBlock:        [32]byte{}, //empty, todo will be genesis first
		Height:           0,
		TargetDifficulty: 20, //target at 20
	}
}

func (bc *Blockchain) GetTargetDifficulty() int64 {
	return bc.TargetDifficulty
}

func (bc *Blockchain) AddBlock(block *Block) error {
	//verify block is legal
	//verify coinbase is present

	//if block has no prev hash reject

	if block.PrevHash == [32]byte{} {
		return errors.New("no prev hash")
	}

	err := bc.DB.Update(func(txn *badger.Txn) error {

		bcbytes, _ := block.Bytes()

		return txn.Set(block.Hash[:], bcbytes)

	})

	if err != nil {
		return err
	}

	return nil

}
