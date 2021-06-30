package lib

import (
	"fmt"
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

// Block constructor
func NewBlock() *Block {
	return &Block{
		id: uuid.NewString(),
		timestamp: time.Now(),
		next: make([]*Block, 0),
		prev: nil,
	}
}

// Formats and prints a block's contents
func (B *Block) Print() {
	var prev_formatted, next_formatted string

	if (B.prev == nil) {
		prev_formatted = "None"
	} else {
		prev_formatted = B.prev.id
	}

	if (B.next == nil || len(B.next) == 0) {
		next_formatted = "None"
	} else {
		for idx, block := range B.next {
			next_formatted += fmt.Sprintf("  > #%v: %v\n", idx, block.id)
		}
	}

	fmt.Println("----- Block Details -----")
	fmt.Printf(
		"- ID: %v\n- Timestamp: %v\n- Next Block: %v\n- Previous Block: %v\n",
		B.id, B.timestamp, next_formatted, prev_formatted,
	)
	fmt.Println("-------------------------")
}

func (B *Block) AddTransaction(
	amount float64,
	sender string,
	recipient string,
	sender_prev_tx string,    // Will become a tx pointer
	recipient_prev_tx string, // ^^
	// ptr TxPointer
	signature string,
	hash string,
) { /* TODO */ }
