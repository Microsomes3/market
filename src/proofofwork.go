package market

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"time"
)

type POW struct {
	Subject    *Block
	TargetDiff *big.Int
}

func VerifyPow(subject *Block, diff uint64) bool {
	td := big.NewInt(1)
	td.Lsh(td, uint(256-diff))

	var intHash big.Int
	intHash.SetBytes(subject.Hash[:])

	return intHash.Cmp(td) < 0
}

func NewPow(subject *Block, diff uint64) *POW {
	td := big.NewInt(1)
	td.Lsh(td, uint(256-diff))

	return &POW{
		Subject:    subject,
		TargetDiff: td,
	}
}

func (pow *POW) FindNonceWithPrefixedHash(suffix string) uint64 {
	nonce := uint64(0)

	for {
		pow.Subject.Nonce = nonce
		pow.Subject.Timestamp = time.Now().Unix() //must be careful not go set it too far ahread from previous 2 hours
		pow.Subject.HashIt()

		var intHash big.Int
		intHash.SetBytes(pow.Subject.Hash[:])

		if intHash.Cmp(pow.TargetDiff) < 0 {
			hashString := hex.EncodeToString(pow.Subject.Hash[:])
			lastChars := hashString[len(hashString)-len(suffix):]

			if lastChars == suffix {
				fmt.Println("Found hash:", hashString)
				fmt.Println("Nonce:", nonce)
				return nonce
			} else {
				fmt.Println(lastChars)
			}
		}

		nonce++
	}
}

func (pow *POW) FindNonce() uint64 {
	nonce := uint64(0)
	maxValue := 1844674407370955161
	for {

		pow.Subject.Nonce = nonce
		pow.Subject.Timestamp = time.Now().Unix() //must be careful not go set it too far ahread from previous 2 hours

		pow.Subject.HashIt()

		var intHash big.Int
		intHash.SetBytes(pow.Subject.Hash[:])

		if intHash.Cmp(pow.TargetDiff) < 0 {
			fmt.Println("Found hash: ", intHash)
			fmt.Println("Hash String", pow.Subject.Hash)
			break
		}

		fmt.Println(fmt.Sprintf("%d/%d", nonce, maxValue))

		nonce++

	}

	return nonce
	// return nonce
}
