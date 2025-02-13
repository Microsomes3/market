package market

import (
	"errors"
	"fmt"

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

	isvalid := block.VerifyBlock()

	if !isvalid {
		return errors.New("block is invalid")
	}

	fmt.Println("added block")
	return nil

}
