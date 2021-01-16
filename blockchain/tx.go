package blockchain

// TxOutput records transaction output
type TxOutput struct {
	Value  int
	PubKey string
}

// TxInput records transaction input
type TxInput struct {
	ID  []byte
	Out int
	Sig string
}

// CanUnlock checks the validity of transaction input
func (in *TxInput) CanUnlock(data string) bool {
	return in.Sig == data
}

// CanBeUnlocked checks the validity of transaction output
func (out *TxOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}
