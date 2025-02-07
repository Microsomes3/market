package main

import "github.com/dgraph-io/badger/v3"

type Blockchain struct {
	DB               *badger.DB
	TailBlock        [32]byte //hash of the latest block
	Height           int
	TargetDifficulty int64
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
