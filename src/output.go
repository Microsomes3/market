package market

import (
	"encoding/hex"
	"encoding/json"
)

type Vout struct {
	Value uint64
	N     uint64
	PK    []byte //who can spend this output
}

func (vout *Vout) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Value uint64 `json:"value"`
		N     uint64 `json:"n"`
		PK    string `json:"publicKey"`
	}{
		Value: vout.Value,
		N:     vout.N,
		PK:    hex.EncodeToString(vout.PK[:]),
	})
}

func (vout *Vout) UnmarshalJSON(bdata []byte) error {

	var tempVout struct {
		Value uint64 `json:"value"`
		N     uint64 `json:"n"`
		PK    string `json:"publicKey"`
	}

	if err := json.Unmarshal(bdata, &tempVout); err != nil {
		return err
	}

	vout.Value = tempVout.Value

	vout.N = tempVout.N

	pkE, err := hex.DecodeString(tempVout.PK)

	if err != nil {
		panic(err)
	}

	vout.PK = pkE

	return nil
}
