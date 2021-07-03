package lib

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Hashes a string with SHA-256,
// returning the hexadecimal encoding of the hash.
func HashString(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))

	hashBytes := hash.Sum(nil)
	return hex.EncodeToString(hashBytes)
}

// Concatenate and hash the transaction's properties
func HashTransaction(tx* Transaction) string {
	tx_str := fmt.Sprintf("%v%v%v%v", tx.timestamp, tx.amount, tx.sender, tx.recipient)
	return HashString(tx_str)
}

func HashBlock (block* Block) { /* TODO */ }
