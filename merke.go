package main

import "crypto/sha256"

type Merkle struct {
	Hashes [][32]byte
}

func NewMarkle() *Merkle {
	return &Merkle{
		Hashes: [][32]byte{},
	}
}

func (m *Merkle) PushHash(h [32]byte) {
	m.Hashes = append(m.Hashes, h)
}

func hashPair(a, b [32]byte) [32]byte {
	mergeAb := append(a[:], b[:]...) // Convert arrays to slices

	return sha256.Sum256(mergeAb)
}

func (m *Merkle) Root() [32]byte {
	for len(m.Hashes) > 1 {
		if len(m.Hashes)%2 != 0 {
			m.Hashes = append(m.Hashes, m.Hashes[len(m.Hashes)-1])
		}

		var newHashes [][32]byte
		for i := 0; i < len(m.Hashes); i += 2 {
			newHashes = append(newHashes, hashPair(m.Hashes[i], m.Hashes[i+1]))
		}

		m.Hashes = newHashes
	}

	return m.Hashes[0]
}
