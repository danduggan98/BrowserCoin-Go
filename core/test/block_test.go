package test

import (
	"bytes"
	"testing"

	"github.com/danduggan98/BrowserCoin-Go/core/src/lib"
)

func TestBlock(t *testing.T) {

	// Create a list of txs for testing
	sender, _ := lib.NewRsaKey()
	recipient, _ := lib.NewRsaKey()

	tx1 := lib.NewTransaction(111, &sender.PublicKey, &recipient.PublicKey).SignAndHash(sender)
	tx2 := lib.NewTransaction(222, &recipient.PublicKey, &sender.PublicKey).SignAndHash(recipient)

	tx_list := make([]*lib.Transaction, 0)
	tx_list = append(tx_list, tx1)
	tx_list = append(tx_list, tx2)

	block := lib.NewBlock(tx_list)

	// Check that values are initialized correctly
	if (len(block.GetNext()) != 0) {
		t.Error("Block should have no successor initially")
	}

	if (block.GetPrev() != nil) {
		t.Error("Block should have no previous initially")
	}

	if (block.GetPrevHash() != nil) {
		t.Error("Block should have no prev hash initially")
	}

	if (block.NumTransactions() != 2) {
		t.Error("Block should have 2 txs initially")
	}

	if (!bytes.Equal(block.GetHash(), make([]byte, 0))) {
		t.Error("Block hash should be empty initially")
	}

	// Update the block hash and check that it worked
	block.UpdateHash()

	if (!bytes.Equal(block.GetHash(), lib.HashBlock(block))) {
		t.Errorf("Wrong block hash: expected %v, got %v", lib.HashBlock(block), block.GetHash())
	}
}