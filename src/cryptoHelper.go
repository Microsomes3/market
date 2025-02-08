package market

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"

	gocryptoeth "github.com/ethereum/go-ethereum/crypto"
	sk1 "github.com/ethereum/go-ethereum/crypto/secp256k1"
)

type CryptoHelper struct{}

func (ci *CryptoHelper) SHA256(msg []byte) [32]byte {
	return sha256.Sum256(msg)
}

func (ci *CryptoHelper) ByteToHex(msg []byte) string {
	return hex.EncodeToString(msg)
}

func (ci *CryptoHelper) GenerateDeterministicKey(seed []byte) (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {

	sk1 := sk1.S256()
	rd := bytes.NewReader(seed)
	priv, err := ecdsa.GenerateKey(sk1, rd)
	if err != nil {
		return nil, nil, err
	}

	return priv, &priv.PublicKey, nil
}

func (ci *CryptoHelper) GeneratePrivateKey() (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	priv, err := gocryptoeth.GenerateKey()
	if err != nil {
		return nil, nil, err
	}

	return priv, &priv.PublicKey, nil

}

func (ci *CryptoHelper) GetPublicKeyBytes(pub *ecdsa.PublicKey) []byte {
	return gocryptoeth.FromECDSAPub(pub)
}

func (ci *CryptoHelper) SignMessage(msg []byte, priv *ecdsa.PrivateKey) (sig []byte, e error) {
	hashOfMsg := sha256.Sum256(msg)

	sig, err := gocryptoeth.Sign(hashOfMsg[:], priv)

	if err != nil {
		return nil, err
	}

	return sig, nil
}

func (ci *CryptoHelper) VerifyMessage(pubKey *ecdsa.PublicKey, msg []byte, sig []byte) bool {
	hashOfMsg := sha256.Sum256(msg)
	return gocryptoeth.VerifySignature(ci.GetPublicKeyBytes(pubKey), hashOfMsg[:], sig[:len(sig)-1])

}
