package main

type Vout struct {
	Value uint64
	N     uint64
	PK    [32]byte //who can spend this output
}
