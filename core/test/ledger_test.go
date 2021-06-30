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
	ledger.Genesis().Print()
}
