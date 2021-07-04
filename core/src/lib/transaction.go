package lib

import (
	"crypto"
	"crypto/rsa"
	"time"
)

type Transaction struct {
	timestamp time.Time
	amount uint64
	sender rsa.PublicKey
	recipient rsa.PublicKey
	//sender_prev_tx string    // Will become a TxPointer
	//recipient_prev_tx string // ^^
	// ptr TxPointer
	signature []byte
	hash []byte
}

func NewTransaction(
	amount uint64,
	sender rsa.PublicKey,
	recipient rsa.PublicKey,
) *Transaction {
	return &Transaction{
		timestamp: time.Now(),
		amount: amount,
		sender: sender,
		recipient: recipient,
		signature: make([]byte, 0),
		hash: make([]byte, 0),
	}
}

// Sign a transaction with a private key
func (T *Transaction) Sign(private_key *rsa.PrivateKey) {
	tx_hash := HashTransaction(T)
	sig, err := rsa.SignPKCS1v15(nil, private_key, crypto.SHA256, tx_hash)

	if err != nil {
		panic(err)
	}

	T.signature = sig
}

func (T *Transaction) Hash() {
	T.hash = HashTransaction(T)
}

// Cryptographically verify a transaction's signature
func (T *Transaction) IsValid() { /* TODO */ }

// Display the transaction's contents
func (T * Transaction) Print() { /* TODO */ }
