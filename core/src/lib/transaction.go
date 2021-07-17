package lib

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"time"
)

type Transaction struct {
	timestamp time.Time
	amount uint64
	sender *rsa.PublicKey
	recipient *rsa.PublicKey
	//sender_prev_tx string    // Will become a TxPointer
	//recipient_prev_tx string // ^^
	// ptr TxPointer
	signature []byte
	hash []byte
}

func NewTransaction(
	amount uint64,
	sender *rsa.PublicKey,
	recipient *rsa.PublicKey,
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
	sig, err := rsa.SignPKCS1v15(rand.Reader, private_key, crypto.SHA256, tx_hash)

	if err != nil {
		panic(err)
	}

	T.signature = sig
}

func (T *Transaction) Hash() {
	T.hash = HashTransaction(T)
}

// Cryptographically verify a transaction's signature
// and check that transaction has correct values
func (T *Transaction) IsValid() (bool, string) {

	if (len(T.signature) == 0) {
		return false, "Transaction not signed"
	}

	if (len(T.hash) == 0) {
		return false, "Transaction not hashed"
	}

	err := rsa.VerifyPKCS1v15(T.sender, crypto.SHA256, T.hash, T.signature)

	if (err != nil) {
		return false, "Cryptographic verification failed"
	}

	return true, ""
}

// Display the transaction's contents
func (T * Transaction) Print() { /* TODO */ }
