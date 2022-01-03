//
// Cryptography tools for hashing, addresses, etc.
//

package lib

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

//////// HASHING \\\\\\\\

// Converts a hash to a hexadecimal string
func StringFromHash(hash []byte) string {
	return hex.EncodeToString(hash)
}

// Hashes a string with SHA-256
func HashString(str string) []byte {
	hash := sha256.Sum256([]byte(str))
	return hash[:]
}

// Hashes a byte array with SHA-256
func HashBytes(bytes []byte) []byte {
	hash := sha256.Sum256(bytes)
	return hash[:]
}

// Concatenates and hashes a transaction's properties
func HashTransaction(tx* Transaction) []byte {
	tx_str := fmt.Sprintf("%v%v%v%v", tx.timestamp, tx.amount, tx.sender, tx.recipient)
	return HashString(tx_str)
}

// Hashes the combined contents of a list of transactions
func HashTransactionList (txs []*Transaction) []byte {
	list_str := make([]byte, 0)

	for _, tx := range txs {
		list_str = append(list_str, HashTransaction(tx)...)
	}

	return list_str
}

// Hashes the combined contents of a block
func HashBlock (block* Block) []byte {
	block_str := fmt.Sprintf("%v%v%v", block.timestamp, block.GetPrevHash(), HashTransactionList(block.transactions))
	return HashString(block_str)
}

//////// KEYS + ADDRESSES \\\\\\\\

// Generates a new RSA key
func NewRsaKey() (*rsa.PrivateKey, error) {
	const num_bits int = 2048
	return rsa.GenerateKey(rand.Reader, num_bits)
}

// Creates an address from a public key
func AddressFromKey(key *rsa.PublicKey) string {
	bytes := key.N.Bytes()
	hash := sha256.Sum256(bytes)

	return StringFromHash(hash[:])
}
