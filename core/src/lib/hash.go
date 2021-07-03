package lib

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Hashes a string with SHA-256, returning both the
// raw bytes and the hex encoding of the hash.
func HashString(str string) ([]byte, string) {
	hash := sha256.New()
	hash.Write([]byte(str))

	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)

	return hashBytes, hashString
}

// Concatenate and hash the transaction's properties,
// returning both the raw bytes and the hex encoding of the hash.
func HashTransaction(tx* Transaction) ([]byte, string) {
	tx_str := fmt.Sprintf("%v%v%v%v%v", tx.timestamp, tx.amount, tx.sender, tx.recipient, tx.signature)
	return HashString(tx_str)
}

func HashBlock (block* Block) { /* TODO */ }
