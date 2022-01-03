package test

import (
	"fmt"
	"testing"

	"github.com/danduggan98/BrowserCoin-Go/core/src/lib"
)

//Creates a new ledger and adds some blocks
func TestLedger(t *testing.T) {
	fmt.Println("Creating new ledger")
	ledger := lib.NewLedger()
	genesis_block := ledger.Genesis()

	fmt.Println("\nGenesis block:")
	genesis_block.Print()

	// Create a list of txs for testing
	sender, _ := lib.NewRsaKey()
	recipient, _ := lib.NewRsaKey()

	tx1 := lib.NewTransaction(111, &sender.PublicKey, &recipient.PublicKey).SignAndHash(sender)
	tx_list := append(make([]*lib.Transaction, 0), tx1)

	fmt.Println("\nAdding new blocks:")
	block1 := lib.NewBlock(tx_list)
	block2 := lib.NewBlock(tx_list)

	ledger.AddBlock(block1, genesis_block)
	ledger.AddBlock(block2, genesis_block)

	block1.Print()
	block2.Print()

	fmt.Println("\nNew genesis block:")
	genesis_block.Print()

	numBranches := genesis_block.NumBranches()
	if (numBranches != 2) {
		t.Errorf("Should have 2 branches, not %d", numBranches)
	}
}
