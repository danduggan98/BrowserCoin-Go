//
// Basic ledger, implemented as a bi-directional linked list
//

package lib

import "time"

type Ledger struct {
	genesis *Block
}

func NewLedger() *Ledger {

	// Construct the genesis block
	return &Ledger{
		genesis: &Block{
			id: "genesis",
			timestamp: time.Now(),
			next: make([]*Block, 0),
			prev: nil,
		},
	}
}

// Connect a new block to an existing one
func (L *Ledger) AddBlock(new_block *Block, predecessor *Block) {
	new_block.prev = predecessor
	predecessor.next = append(predecessor.next, new_block)
}

func (L *Ledger) LongestValidChain() { /* TODO */ }
func (L *Ledger) FindBlockByID() { /* TODO */ }
