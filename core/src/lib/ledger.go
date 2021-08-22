//
// Basic ledger, implemented as a bi-directional linked list
//

package lib

import "time"

type Ledger struct {
	genesis *Block
}

// Ledger constructor - creates the genesis block
func NewLedger() *Ledger {
	genesis_block := &Block{
		timestamp: time.Now(),
		next: make([]*Block, 0),
		prev: nil,
		transactions: make([]*Transaction, 0),
	}

	genesis_block.UpdateHash()
	return &Ledger{ genesis_block }
}

// Returns the genesis block
func (L* Ledger) Genesis() *Block {
	return L.genesis
}

// Connects a new block to an existing one
func (L *Ledger) AddBlock(new_block *Block, predecessor *Block) {
	new_block.prev = predecessor
	predecessor.next = append(predecessor.next, new_block)
	new_block.UpdateHash()
}

func (L *Ledger) LongestValidChain() { /* TODO */ }
