package market

type Vin struct {
	TXID [32]byte //from which tx is it coming from
	Vout int      //which output of that tx

	Signature [64]byte //need to prove the right to spend this input
}
