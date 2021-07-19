package lib

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
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

//////// Getters \\\\\\\\

func (T *Transaction) GetTimestamp() time.Time { return T.timestamp }
func (T *Transaction) GetAmount() uint64 { return T.amount }
func (T *Transaction) GetSender() *rsa.PublicKey { return T.sender }
func (T *Transaction) GetRecipient() *rsa.PublicKey { return T.recipient }
func (T *Transaction) GetSignature() []byte { return T.signature }
func (T *Transaction) GetHash() []byte { return T.hash }

//////// Utilities \\\\\\\\

// Sign a transaction with a private key, then store its hash
func (T *Transaction) SignAndHash(private_key *rsa.PrivateKey) {
	tx_hash := HashTransaction(T)
	sig, err := rsa.SignPKCS1v15(rand.Reader, private_key, crypto.SHA256, tx_hash)

	if err != nil {
		panic(err)
	}

	T.signature = sig
	T.hash = tx_hash
}

// Verify a transaction's signature and properties
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
func (T * Transaction) Print() {
	var sig_formatted, hash_formatted string

	if (len(T.signature) == 0) {
		sig_formatted = "None"
	} else {
		sig_formatted = StringFromHash(T.signature)
	}

	if (len(T.hash) == 0) {
		hash_formatted = "None"
	} else {
		hash_formatted = StringFromHash(T.hash)
	}

	fmt.Println("----- Transaction Details -----")
	fmt.Printf(
		"- Timestamp: %v\n- Amount: %v\n- Sender: %v\n- Recipient: %v\n- Signature: %v\n- Hash: %v\n",
		T.timestamp, T.amount, AddressFromKey(T.sender), AddressFromKey(T.recipient), sig_formatted, hash_formatted,
	)
	fmt.Println("-------------------------------")
}
