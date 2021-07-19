//
// A block containing a list of transactions and pointers to the next/last blocks
//

package lib

import (
	"fmt"
	"time"
)

type Block struct {
	timestamp time.Time
	next []*Block
	prev *Block
	transactions []*Transaction
	hash []byte
	prev_hash []byte
}

// Block constructor
func NewBlock() *Block {
	block := &Block{
		timestamp: time.Now(),
		next: make([]*Block, 0),
		prev: nil,
		transactions: make([]*Transaction, 0),
		hash: make([]byte, 0),
		prev_hash: make([]byte, 0),
	}

	block.UpdateHash()
	return block
}

//////// Getters \\\\\\\\

func (B *Block) GetTimestamp() time.Time { return B.timestamp }
func (B *Block) GetNext() []*Block { return B.next }
func (B *Block) GetPrev() *Block { return B.prev }
func (B *Block) GetTransactions() []*Transaction { return B.transactions }
func (B *Block) GetHash() []byte { return B.hash }
func (B *Block) GetPrevHash() []byte { return B.prev_hash }

//////// Utilities \\\\\\\\

func (B *Block) NumBranches() int { return len(B.next) }

func (B *Block) UpdateHash() { B.hash = HashBlock(B) }

// Formats and prints a block's contents
func (B *Block) Print() {
	var prev_formatted, next_formatted string

	if (B.prev == nil) {
		prev_formatted = "None"
	} else {
		prev_formatted = StringFromHash(B.prev.hash)
	}

	if (B.next == nil || len(B.next) == 0) {
		next_formatted = "[]"
	} else {
		for idx, block := range B.next {
			next_formatted += fmt.Sprintf("\n  > #%v: %v", idx, StringFromHash(block.hash))
		}
	}

	fmt.Println("----- Block Details -----")
	fmt.Printf(
		"- Timestamp: %v\n- Next Block: %v\n- Previous Block: %v\n- Hash: %v\n",
		B.timestamp, next_formatted, prev_formatted, StringFromHash(B.hash),
	)
	fmt.Println("-------------------------")
}

/*func (B *Block) AddTransaction(tx Transaction) {} */
