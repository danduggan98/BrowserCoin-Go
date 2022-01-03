//
// Merkle tree data structure
//

package lib

import (
	"github.com/danduggan98/BrowserCoin-Go/core/src/collections"
)

type (
	node struct {
		hash []byte
		left_child *node
		right_child *node
	}

	MerkleTree struct {
		root_hash []byte
	}
)

// Merkle tree constructor 
func NewMerkleTree(transactions []*Transaction) *MerkleTree {

	// Panic if no transactions provided
	if len(transactions) == 0 {
		panic("Merkle tree can not be constructed without transactions")
	}
	
	// Odd number of transactions -> duplicate the last one
	var tx_list []*Transaction = transactions[:]
	num_txs := len(tx_list)

	if (num_txs % 2 != 0) {
		last_tx := tx_list[num_txs - 1]
		tx_list = append(tx_list, last_tx)
	}

	node_queue := collections.NewQueue()

	// Create leaf nodes for each transaction
	for _, tx := range tx_list {
		leaf := &node{
			hash: tx.GetHash(),
			left_child: nil,
			right_child: nil,
		}

		node_queue.Push(leaf)
	}

	// Combine pairs of nodes together
	for node_queue.Size() > 1 {
		first := node_queue.Pop().(*node)
		second := node_queue.Pop().(*node)

		combined_hashes := append(first.hash, second.hash...)
		new_node := &node{
			hash: HashBytes(combined_hashes),
			left_child: first,
			right_child: second,
		}

		node_queue.Push(new_node)
	}

	root := node_queue.Pop().(*node)
	return &MerkleTree{ root.hash }
}

func (M *MerkleTree) GetRootHash () []byte { return M.root_hash }

// TODO
// Verify that a transcation is included in this Merkle tree
// func Proof()
