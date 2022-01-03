package test

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"testing"

	"github.com/danduggan98/BrowserCoin-Go/core/src/lib"
)

func TestMerkleTree(t *testing.T) {

	// Addresses
	sender, _ := rsa.GenerateKey(rand.Reader, 2048)
	recipient, _ := rsa.GenerateKey(rand.Reader, 2048)

	// Create some txs
	txs := make([]*lib.Transaction, 0)

	tx1 := lib.NewTransaction(111, &sender.PublicKey, &recipient.PublicKey)
	tx2 := lib.NewTransaction(222, &sender.PublicKey, &recipient.PublicKey)
	
	tx1.SignAndHash(sender)
	tx2.SignAndHash(sender)

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