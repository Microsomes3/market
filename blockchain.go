package main

import "github.com/dgraph-io/badger/v3"

type Blockchain struct {
	DB        *badger.DB
	TailBlock [32]byte //hash of the latest block
	Height    int
}

func NewBlockchain() *Blockchain {
	return &Blockchain{}
}
