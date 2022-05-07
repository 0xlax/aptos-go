package aptostypes

import (
	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

func HashPrefix(name string) []byte {
	return Hash([]byte("APTOS::"), []byte(name))
}

func Hash(prefix []byte, bytes []byte) []byte {
	sha256 := sha3.New256()
	sha256.Write(prefix)
	sha256.Write(bytes)
	return sha256.Sum(nil)
}

func (t *SignedTransaction) TransactionHash() string {
	return hex.EncodeToString(Hash(
		HashPrefix("Transaction"),
		ToBCS(&Transaction__UserTransaction{*t}),
	))
}
