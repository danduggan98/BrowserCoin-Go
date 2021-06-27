package lib

// Basic ledger, implemented as a bi-directional linked list

type Ledger struct {
	genesis *Block

	//TODO - linked list
}

func NewLedger() *Ledger {
	blockchain := &Ledger{nil}

	return blockchain
}
