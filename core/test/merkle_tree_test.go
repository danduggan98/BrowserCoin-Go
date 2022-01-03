package test

import (
	"bytes"
	"testing"

	"github.com/danduggan98/BrowserCoin-Go/core/src/lib"
)

func TestMerkleTree(t *testing.T) {

	// Addresses
	sender, _ := lib.NewRsaKey()
	recipient, _ := lib.NewRsaKey()

	// Create some txs
	txs := make([]*lib.Transaction, 0)

	tx1 := lib.NewTransaction(111, &sender.PublicKey, &recipient.PublicKey).SignAndHash(sender)
	tx2 := lib.NewTransaction(222, &sender.PublicKey, &recipient.PublicKey).SignAndHash(sender)

	txs = append(txs, tx1)
	txs = append(txs, tx2)

	tree := lib.NewMerkleTree(txs)

	// Check if root hash is combo of these two
	combined_hashes := lib.HashBytes(
		append(tx1.GetHash(), tx2.GetHash()...),
	)

	if (!bytes.Equal(combined_hashes, tree.GetRootHash())) {
		t.Errorf("Incorrect root hash - expected %v, got %v", combined_hashes, tree.GetRootHash())
	}
}