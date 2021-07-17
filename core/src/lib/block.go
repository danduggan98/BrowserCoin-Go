//
// A block containing a list of transactions and pointers to the next/last blocks
//

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
	// transactions []Transaction
	// hash string
	// prev_hash string
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

//////// Getters \\\\\\\\

func (B *Block) GetID() string { return B.id }
func (B *Block) GetTimestamp() time.Time { return B.timestamp }
func (B *Block) GetNext() []*Block { return B.next }
func (B *Block) GetPrev() *Block { return B.prev }

//////// Utilities \\\\\\\\

func (B *Block) NumBranches() int { return len(B.next) }

// Formats and prints a block's contents
func (B *Block) Print() {
	var prev_formatted, next_formatted string

	if (B.prev == nil) {
		prev_formatted = "None"
	} else {
		prev_formatted = B.prev.id
	}

	if (B.next == nil || len(B.next) == 0) {
		next_formatted = "[]"
	} else {
		for idx, block := range B.next {
			next_formatted += fmt.Sprintf("\n  > #%v: %v", idx, block.id)
		}
	}

	fmt.Println("----- Block Details -----")
	fmt.Printf(
		"- ID: %v\n- Timestamp: %v\n- Next Block: %v\n- Previous Block: %v\n",
		B.id, B.timestamp, next_formatted, prev_formatted,
	)
	fmt.Println("-------------------------")
}

/*func (B *Block) AddTransaction(
	amount float64,
	sender string,
	recipient string,
	sender_prev_tx string,    // Will become a tx pointer
	recipient_prev_tx string, // ^^
	// ptr TxPointer
	signature string,
	hash string,
) {}
*/

/*func (B *Block) AddTransaction(tx Transaction) {} */
