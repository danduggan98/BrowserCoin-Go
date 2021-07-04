package lib

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Hashes a string with SHA-256
func HashString(str string) []byte {
	hash := sha256.New()
	hash.Write([]byte(str))

	return hash.Sum(nil)
}

// Converts a hash to a hexadecimal string
func StringFromHash(hash []byte) string {
	return hex.EncodeToString(hash)
}

// Concatenates and hashes a transaction's properties
func HashTransaction(tx* Transaction) []byte {
	tx_str := fmt.Sprintf("%v%v%v%v%v", tx.timestamp, tx.amount, tx.sender, tx.recipient, tx.signature)
	return HashString(tx_str)
}

func HashBlock (block* Block) { /* TODO */ }
