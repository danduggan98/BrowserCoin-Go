package lib

import (
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Hashes a string with SHA-256
func HashString(str string) []byte {
	hash := sha256.Sum256([]byte(str))
	return hash[:]
}

// Converts a hash to a hexadecimal string
func StringFromHash(hash []byte) string {
	return hex.EncodeToString(hash)
}

// Creates an address from a public key
func AddressFromKey(key *rsa.PublicKey) string {
	bytes := key.N.Bytes()
	hash := sha256.Sum256(bytes)

	return StringFromHash(hash[:])
}

// Concatenates and hashes a transaction's properties
func HashTransaction(tx* Transaction) []byte {
	tx_str := fmt.Sprintf("%v%v%v%v", tx.timestamp, tx.amount, tx.sender, tx.recipient)
	return HashString(tx_str)
}

func HashBlock (block* Block) { /* TODO */ }
