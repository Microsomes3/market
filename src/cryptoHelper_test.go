package market

import (
	"testing"
)

func TestItCanGeneratePrivateKey(t *testing.T) {

	crp := CryptoHelper{}
	_, _, err := crp.GeneratePrivateKey()
	if err != nil {
		t.Fail()
	}

}

func TestItGetsPubKeyBytes(t *testing.T) {

	crp := CryptoHelper{}
	_, pubk, _ := crp.GeneratePrivateKey()

	bytesPub := crp.GetPublicKeyBytes(pubk)

	if len(bytesPub) != 65 {
		t.Fail()
	}

}

func TestItCanSignMessage(t *testing.T) {

	crp := CryptoHelper{}
	priv, _, _ := crp.GeneratePrivateKey()
	msg := []byte("hello world")

	sig, err := crp.SignMessage(msg, priv)

	if len(sig) != 65 {
		t.Fail()
	}

	if err != nil {
		t.Fail()
	}
}

func TestItCanVerify(t *testing.T) {

	crp := CryptoHelper{}
	priv, pubk, _ := crp.GeneratePrivateKey()
	msg := []byte("hello world")

	sig, err := crp.SignMessage(msg, priv)

	if len(sig) != 65 {
		t.Fail()
	}

	if err != nil {
		t.Fail()
	}

	isValid := crp.VerifyMessage(pubk, msg, sig)

	if isValid != true {
		t.Fail()
	}

}

func TestItCanTurnPrivKeyToBytesAndBack(t *testing.T) {

	crp := CryptoHelper{}
	privk, _, err := crp.GeneratePrivateKey()

	if err != nil {
		t.Fail()
	}

	privKBytes := crp.GetPrivateKeyBytes(privk)

	x := privk.X
	y := privk.Y

	if len(privKBytes) != 32 {
		t.Fail()
	}

	privk2, err := crp.GetPrivKeyFromBytes(privKBytes)

	if err != nil {
		t.Fail()
	}

	if x.Cmp(privk2.X) != 0 {
		t.Fail()
	}

	if y.Cmp(privk2.Y) != 0 {
		t.Fail()
	}

}
