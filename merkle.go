package main

import (
	"crypto/sha256"
)

type MerkleTree struct {
	Leaves [][32]byte
}

func (m *MerkleTree) PushHash(hash [32]byte) {
	m.Leaves = append(m.Leaves, hash)
}

func (m *MerkleTree) PairHash(left, right [32]byte) [32]byte {
	// Always concatenate in order: left + right
	return sha256.Sum256(append(left[:], right[:]...))
}

func (m *MerkleTree) Root() [32]byte {
	hashes := m.Leaves

	// Handle empty tree
	if len(hashes) == 0 {
		return [32]byte{}
	}

	// Ensure even number of hashes
	if len(hashes)%2 != 0 {
		hashes = append(hashes, hashes[len(hashes)-1])
	}

	for len(hashes) > 1 {
		var newHashes [][32]byte
		for i := 0; i < len(hashes); i += 2 {
			newHashes = append(newHashes, m.PairHash(hashes[i], hashes[i+1]))
		}
		hashes = newHashes
	}

	return hashes[0]
}

func (m *MerkleTree) GetProof(target [32]byte) ([][32]byte, []bool, bool) {
	var proof [][32]byte
	var isLeftNode []bool // Track whether each proof element should be on the left
	hashes := m.Leaves

	index := -1
	// Find the index of the target leaf
	for i, h := range hashes {
		if h == target {
			index = i
			break
		}
	}

	if index == -1 {
		return nil, nil, false // Target not found
	}

	// Build the proof
	currentIndex := index
	for len(hashes) > 1 {
		var newHashes [][32]byte
		for i := 0; i < len(hashes); i += 2 {
			if i+1 >= len(hashes) {
				hashes = append(hashes, hashes[i]) // Duplicate last element if odd
			}

			// Add proof element
			if i == currentIndex {
				proof = append(proof, hashes[i+1])
				isLeftNode = append(isLeftNode, false) // Sibling goes on right
			} else if i+1 == currentIndex {
				proof = append(proof, hashes[i])
				isLeftNode = append(isLeftNode, true) // Sibling goes on left
			}

			newHashes = append(newHashes, m.PairHash(hashes[i], hashes[i+1]))
		}

		hashes = newHashes
		currentIndex = currentIndex / 2
	}

	return proof, isLeftNode, true
}

func calculateRoot(leaf [32]byte, proof [][32]byte, isLeftNode []bool) [32]byte {
	currentHash := leaf

	for i, proofElement := range proof {
		if isLeftNode[i] {
			// Proof element goes on the left
			currentHash = sha256.Sum256(append(proofElement[:], currentHash[:]...))
		} else {
			// Proof element goes on the right
			currentHash = sha256.Sum256(append(currentHash[:], proofElement[:]...))
		}
	}

	return currentHash
}

// func main() {
// 	// Create sample leaves
// 	leaf1 := sha256.Sum256([]byte("leaf1"))
// 	leaf2 := sha256.Sum256([]byte("leaf2"))
// 	leaf3 := sha256.Sum256([]byte("leaf3"))
// 	leaf4 := sha256.Sum256([]byte("leaf4"))

// 	// Initialize the Merkle tree
// 	m := MerkleTree{}
// 	m.PushHash(leaf1)
// 	m.PushHash(leaf2)
// 	m.PushHash(leaf3)
// 	m.PushHash(leaf4)

// 	// Compute the Merkle root
// 	root := m.Root()
// 	fmt.Println("Root:", hex.EncodeToString(root[:]))

// 	// Generate the proof for leaf1
// 	proof, isLeftNode, success := m.GetProof(leaf1)
// 	if !success {
// 		fmt.Println("❌ Failed to generate proof!")
// 		return
// 	}

// 	// Print the proof for verification
// 	fmt.Println("\nProof elements:")
// 	for i, p := range proof {
// 		side := "right"
// 		if isLeftNode[i] {
// 			side = "left"
// 		}
// 		fmt.Printf("Element %d (%s): %s\n", i+1, side, hex.EncodeToString(p[:]))
// 	}

// 	// Verify the proof
// 	verifyRoot := calculateRoot(leaf1, proof, isLeftNode)
// 	fmt.Println("\nVerify Root:", hex.EncodeToString(verifyRoot[:]))

// 	// Final verification
// 	if verifyRoot == root {
// 		fmt.Println("✅ Proof verified successfully!")
// 	} else {
// 		fmt.Println("❌ Proof verification failed!")
// 	}
// }
