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
	merkle_root []byte
}

// Block constructor
func NewBlock(txs []*Transaction) *Block {
	return &Block{
		timestamp: time.Now(),
		next: make([]*Block, 0),
		prev: nil,
		transactions: txs,
		hash: make([]byte, 0),
		merkle_root: NewMerkleTree(txs).GetRootHash(),
	}
}

//////// Getters \\\\\\\\

func (B *Block) GetTimestamp() time.Time { return B.timestamp }
func (B *Block) GetNext() []*Block { return B.next }
func (B *Block) GetPrev() *Block { return B.prev }
func (B *Block) GetTransactions() []*Transaction { return B.transactions }
func (B *Block) GetHash() []byte { return B.hash }
func (B *Block) GetMerkleRoot() []byte { return B.merkle_root }

//////// Utilities \\\\\\\\

// Determines the number of successors to this block
func (B *Block) NumBranches() int { return len(B.next) }

// Determines the number of transactions in a block
func (B *Block) NumTransactions() int { return len(B.transactions) }

// Calculates and stores the block's hash
func (B *Block) UpdateHash() *Block {
	B.hash = HashBlock(B)
	return B
}

// Gets the hash of the previous block
func (B *Block) GetPrevHash() []byte {
	if (B.prev == nil) {
		return nil
	}

	return B.prev.hash
}

// Formats and prints the block's contents
func (B *Block) Print() {
	prev_hash := B.GetPrevHash()
	var prev_hash_str, next_blocks string

	if (prev_hash == nil) {
		prev_hash_str = "None"
	} else {
		prev_hash_str = StringFromHash(prev_hash)
	}

	if (len(B.next) == 0) {
		next_blocks = "[]"
	} else {
		for idx, block := range B.next {
			next_blocks += fmt.Sprintf("\n  > #%v: %v", idx, StringFromHash(block.hash))
		}
	}

	fmt.Println("----- Block Details -----")
	fmt.Printf(
		"- Timestamp: %v\n- # Txs: %v\n- Next Blocks: %v\n- Previous Block Hash: %v\n- Hash: %v\n",
		B.timestamp, B.NumTransactions(), next_blocks, prev_hash_str, StringFromHash(B.hash),
	)
	fmt.Println("-------------------------")
}

/*func (B *Block) AddTransaction(tx Transaction) {} */
/*func (B *Block) IsValid() {} */
