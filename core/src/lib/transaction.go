package lib

import (
	"crypto"
	"crypto/rsa"
	"time"
)

type Transaction struct {
	hash string
	timestamp time.Time
	amount uint64
	sender rsa.PublicKey
	recipient rsa.PublicKey
	//sender_prev_tx string    // Will become a TxPointer
	//recipient_prev_tx string // ^^
	// ptr TxPointer
	signature string
}

func NewTransaction(
	amount uint64,
	sender rsa.PublicKey,
	recipient rsa.PublicKey,
) *Transaction {
	return &Transaction{
		hash: "",
		timestamp: time.Now(),
		amount: amount,
		sender: sender,
		recipient: recipient,
		signature: "",
	}
}

// Sign a transaction with a private key
func (T *Transaction) Sign(pk *rsa.PrivateKey) {
	tx_hash, _ := HashTransaction(T)
	sig, err := rsa.SignPKCS1v15(nil, pk, crypto.SHA256, tx_hash)

	if err != nil {
		panic(err)
	}

	T.signature = string(sig)
}

func (T *Transaction) Hash() {
	_, tx_hash := HashTransaction(T)
	T.hash = tx_hash
}

// Cryptographically verify a transaction's signature
func (T *Transaction) IsValid() { /* TODO */ }

// Display the transaction's contents
func (T * Transaction) Print() { /* TODO */ }
