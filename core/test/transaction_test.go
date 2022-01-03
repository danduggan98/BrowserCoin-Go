package test

import (
	"testing"

	"github.com/danduggan98/BrowserCoin-Go/core/src/lib"
)

// Create, sign, and verify a transaction
func TestTx(t *testing.T) {

	// Create two users
	sender, sender_err := lib.NewRsaKey()

	if (sender_err != nil) {
		panic(sender_err)
	}

	recipient, recipient_err := lib.NewRsaKey()

	if (recipient_err != nil) {
		panic(recipient_err)
	}

	// Valid transaction
	tx1 := lib.NewTransaction(42, &sender.PublicKey, &recipient.PublicKey).SignAndHash(sender)
	tx1.Print()

	tx1_valid, tx1_invalid_reason := tx1.IsValid()

	if (tx1_valid == false) {
		t.Errorf("Transaction should be valid, but failed because: %v", tx1_invalid_reason)
	}

	// Invalid transaction - signed by wrong person
	tx2 := lib.NewTransaction(5000, &sender.PublicKey, &recipient.PublicKey)
	tx2.SignAndHash(recipient) // Attempt to steal funds
	tx2.Print()

	tx2_valid, _ := tx2.IsValid()

	if (tx2_valid == true) {
		t.Error("Transaction should be invalid, but succeeded")
	}

	// Invalid transaction - not signed
	tx3 := lib.NewTransaction(123, &sender.PublicKey, &recipient.PublicKey)
	tx3.Print()

	tx3_valid, _ := tx3.IsValid()

	if (tx3_valid == true) {
		t.Error("Unsigned transaction should be invalid, but succeeded")
	}
}
