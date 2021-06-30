package lib

import (
	"time"

	"github.com/google/uuid"
)

type Block struct {
	id string
	timestamp time.Time
	next []*Block
	prev *Block
	//transactions []Transaction
}

func NewBlock() *Block {
	return &Block{
		id: uuid.NewString(),
		timestamp: time.Now(),
		next: make([]*Block, 0),
		prev: nil,
	}
}

func AddTransaction(
	amount float64,
	sender string,
	recipient string,
	sender_prev_tx string,    // Will become a tx pointer
	recipient_prev_tx string, // ^^
	// ptr TxPointer
	signature string,
	hash string,
) { /* TODO */ }
