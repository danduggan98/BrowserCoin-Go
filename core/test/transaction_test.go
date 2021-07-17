package test

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"

	"github.com/danduggan98/BrowserCoin-Go/core/src/lib"
)

// Create, sign, and verify a transaction
func TestTx(t *testing.T) {

	// Create two users
	sender, sender_err := rsa.GenerateKey(rand.Reader, 2048)

	if (sender_err != nil) {
		panic(sender_err)
	}

	recipient, recipient_err := rsa.GenerateKey(rand.Reader, 2048)

	if (recipient_err != nil) {
		panic(recipient_err)
	}

	// Valid transaction
	tx1 := lib.NewTransaction(50, &sender.PublicKey, &recipient.PublicKey)
	tx1.Sign(sender)
	tx1.Hash()

	tx1_valid, tx1_invalid_reason := tx1.IsValid()

	if (tx1_valid == false) {
		t.Errorf("Transaction should be valid, but failed because: %v", tx1_invalid_reason)
	}

	// Invalid transaction - signed by wrong person
	tx2 := lib.NewTransaction(5000, &sender.PublicKey, &recipient.PublicKey)
	tx2.Sign(recipient) // Attempt to steal funds
	tx2.Hash()

	tx2_valid, _ := tx2.IsValid()

	if (tx2_valid == true) {
		t.Error("Transaction should be invalid, but succeeded")
	}
}
