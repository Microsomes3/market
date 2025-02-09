package market

import (
	"encoding/hex"
	"encoding/json"
)

type Vin struct {
	TXID [32]byte //from which tx is it coming from
	Vout int      //which output of that tx

	Signature [64]byte //need to prove the right to spend this input
}

func (vin *Vin) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		TXID string `json:"txid"`
		Vout int    `json:"vout"`
		Sig  string `json:"sig"`
	}{
		TXID: hex.EncodeToString(vin.TXID[:]),
		Vout: vin.Vout,
		Sig:  hex.EncodeToString(vin.Signature[:]),
	})
}
