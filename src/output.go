package market

type Vout struct {
	Value uint64
	N     uint64
	PK    []byte //who can spend this output
}
