package main

type Market struct {
	Hash      [32]byte
	Name      [32]byte
	Owner     [32]byte // hash of their public key
	BlockHash [32]byte //what block was it created from
}

func NewMarket(marketName [32]byte, owner [32]byte, bh [32]byte) *Market {
	return &Market{
		Hash:      [32]byte{}, //will be calculated later
		Name:      marketName,
		Owner:     owner,
		BlockHash: bh,
	}
}
